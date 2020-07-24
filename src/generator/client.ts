/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { camelCase } from '@azure-tools/codegen';
import { CodeModel, Parameter } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { contentPreamble, formatParameterTypeName, sortParametersByRequired } from './helpers';
import { ImportManager } from './imports';

// generates content for client.go
export async function generateClient(session: Session<CodeModel>): Promise<string> {
  // add standard imports
  imports.add('fmt');
  const isARM = session.model.language.go!.openApiType === 'arm';
  if (isARM) {
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/armcore');
  }
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
  imports.add('strings');
  imports.add('net/url');

  let text = await contentPreamble(session);
  text += imports.text();

  if (session.model.security.authenticationRequired) {
    const scope = await session.getValue('credential-scope');
    text += `const scope = "${scope}"\n`;
  }
  text += `const telemetryInfo = "azsdk-go-${session.model.language.go!.packageName}/<version>"\n`;
  const exportClient = await session.getValue('export-client', true);
  let clientOptions = 'ClientOptions';
  let defaultClientOptions = 'DefaultClientOptions';
  if (!exportClient) {
    clientOptions = 'clientOptions';
    defaultClientOptions = 'defaultClientOptions';
  }
  text += `// ${clientOptions} contains configuration settings for the default client's pipeline.\n`;
  text += `type ${clientOptions} struct {\n`;
  text += '\t// HTTPClient sets the transport for making HTTP requests.\n';
  text += '\tHTTPClient azcore.Transport\n';
  text += '\t// LogOptions configures the built-in request logging policy behavior.\n';
  text += '\tLogOptions azcore.RequestLogOptions\n';
  text += '\t// Retry configures the built-in retry policy behavior.\n';
  text += '\tRetry azcore.RetryOptions\n';
  text += '\t// Telemetry configures the built-in telemetry policy behavior.\n';
  text += '\tTelemetry azcore.TelemetryOptions\n';
  text += '\t// ApplicationID is an application-specific identification string used in telemetry.\n';
  text += '\t// It has a maximum length of 24 characters and must not contain any spaces.\n';
  text += '\tApplicationID string\n';
  if (isARM) {
    text += '\t// DisableRPRegistration controls if an unregistered resource provider should\n';
    text += '\t// automatically be registered. See https://aka.ms/rps-not-found for more information.\n';
    text += '\t// The default value is false, meaning registration will be attempted.\n';
    text += '\tDisableRPRegistration bool\n';
  }
  text += '}\n\n';
  text += `// ${defaultClientOptions} creates a ${clientOptions} type initialized with default values.\n`;
  text += `func ${defaultClientOptions}() ${clientOptions} {\n`;
  text += `\treturn ${clientOptions}{\n`;
  text += '\t\tHTTPClient: azcore.DefaultHTTPClientTransport(),\n';
  text += '\t\tRetry: azcore.DefaultRetryOptions(),\n';
  text += '\t}\n';
  text += '}\n\n';

  text += `func (c *${clientOptions}) telemetryOptions() azcore.TelemetryOptions {\n`;
  text += '\tt := telemetryInfo\n';
  text += '\tif c.ApplicationID != "" {\n';
  text += '\t\ta := strings.ReplaceAll(c.ApplicationID, " ", "/")\n';
  text += '\t\tif len(a) > 24 {\n';
  text += '\t\t\ta = a[:24]\n';
  text += '\t\t}\n';
  text += '\t\tt = fmt.Sprintf("%s %s", a, telemetryInfo)\n';
  text += '\t}\n';
  text += '\tif c.Telemetry.Value == "" {\n';
  text += '\t\treturn azcore.TelemetryOptions{Value: t}\n';
  text += '\t}\n';
  text += '\treturn azcore.TelemetryOptions{Value: fmt.Sprintf("%s %s", c.Telemetry.Value, t)}\n';
  text += '}\n\n';

  // Client
  let client = 'Client';
  let defaultEndpoint = 'DefaultEndpoint';
  let newDefaultClient = 'NewDefaultClient';
  let newClient = 'NewClient';
  let newClientWithPipeline = 'NewClientWithPipeline';
  if (!exportClient) {
    client = 'client';
    defaultEndpoint = 'defaultEndpoint';
    newDefaultClient = 'newDefaultClient';
    newClient = 'newClient';
    newClientWithPipeline = 'newClientWithPipeline';
  }
  if (session.model.info.description) {
    text += `// ${client} - ${session.model.info.description}\n`;
  }

  text += `type ${client} struct {\n`;
  text += `\tu *url.URL\n`;
  text += `\tp azcore.Pipeline\n`;
  text += '}\n\n';

  let credParam = 'cred azcore.Credential, ';
  if (!session.model.security.authenticationRequired) {
    credParam = '';
  }
  const endpoint = getDefaultEndpoint(session.model.globalParameters);
  if (endpoint) {
    text += `// ${defaultEndpoint} is the default service endpoint.\n`;
    text += `const ${defaultEndpoint} = "${endpoint}"\n\n`;
    text += `// ${newDefaultClient} creates an instance of the ${client} type using the ${defaultEndpoint}.\n`;
    text += `func ${newDefaultClient}(${credParam}options *${clientOptions}) (*${client}, error) {\n`;
    let cred = 'cred, ';
    if (!session.model.security.authenticationRequired) {
      cred = '';
    }
    // when a parameterized host is not set, the default endpoint is set for the url on the client
    text += `\treturn ${newClient}(${defaultEndpoint}, ${cred}options)\n`;
    text += '}\n\n';
  }

  const hostParam = getHostParam(session.model.globalParameters);
  const hostParamName = camelCase(hostParam!.language.go!.name).toLowerCase();
  text += `// ${newClient} creates an instance of the ${client} type with the specified endpoint.\n`;
  text += `func ${newClient}(${hostParamName} ${hostParam!.schema.language.go!.name}, ${credParam}options *${clientOptions}) (*${client}, error) {\n`;
  text += '\tif options == nil {\n';
  text += `\t\to := ${defaultClientOptions}()\n`;
  text += '\t\toptions = &o\n';
  text += '\t}\n';
  const telemetryPolicy = 'azcore.NewTelemetryPolicy(options.telemetryOptions())';
  const reqIDPolicy = 'azcore.NewUniqueRequestIDPolicy()';
  const retryPolicy = 'azcore.NewRetryPolicy(&options.Retry)';
  const credPolicy = 'cred.AuthenticationPolicy(azcore.AuthenticationPolicyOptions{Options: azcore.TokenRequestOptions{Scopes: []string{scope}}})';
  const logPolicy = 'azcore.NewRequestLogPolicy(options.LogOptions))';
  // ARM will optionally inject the RP registration policy into the pipeline
  if (isARM) {
    text += '\tpolicies := []azcore.Policy{\n';
    text += `\t\t${telemetryPolicy},\n`;
    text += `\t\t${reqIDPolicy},\n`;
    text += '\t}\n';
    // RP registration policy must appear before the retry policy
    text += '\tif !options.DisableRPRegistration {\n';
    text += '\t\trpOpts := armcore.DefaultRegistrationOptions()\n';
    text += '\t\trpOpts.HTTPClient = options.HTTPClient\n';
    text += '\t\trpOpts.LogOptions = options.LogOptions\n';
    text += '\t\trpOpts.Retry = options.Retry\n';
    text += '\t\tpolicies = append(policies, armcore.NewRPRegistrationPolicy(cred, &rpOpts))\n';
    text += '\t}\n';
    text += '\tpolicies = append(policies,\n';
    text += `\t\t${retryPolicy},\n`;
    // ARM implies authentication is required
    text += `\t\t${credPolicy},\n`;
    text += `\t\t${logPolicy}\n`;
    text += '\tp := azcore.NewPipeline(options.HTTPClient, policies...)\n';
  } else {
    text += '\tp := azcore.NewPipeline(options.HTTPClient,\n';
    text += `\t\t${telemetryPolicy},\n`;
    text += `\t\t${reqIDPolicy},\n`;
    text += `\t\t${retryPolicy},\n`;
    if (session.model.security.authenticationRequired) {
      text += `\t\t${credPolicy},\n`;
    }
    text += `\t\t${logPolicy}\n`;
  }
  text += `\treturn ${newClientWithPipeline}(${hostParamName}, p)\n`;
  text += '}\n\n';

  text += `// ${newClientWithPipeline} creates an instance of the ${client} type with the specified endpoint and pipeline.\n`;
  text += `func ${newClientWithPipeline}(${hostParamName} ${hostParam!.schema.language.go!.name}, p azcore.Pipeline) (*${client}, error) {\n`;
  text += `\tu, err := url.Parse(${hostParamName})\n`;
  text += '\tif err != nil {\n';
  text += '\t\treturn nil, err\n';
  text += '\t}\n';
  text += '\tif u.Scheme == "" {\n';
  text += `\t\treturn nil, fmt.Errorf("no scheme detected in endpoint %s", ${hostParamName})\n`;
  text += '\t}\n';
  text += `\treturn &${client}{u: u, p: p}, nil\n`;

  text += '}\n\n';

  for (const group of values(session.model.operationGroups)) {
    const clientLiterals = [`${client}: client`];
    const methodParams = new Array<string>();
    // add client params to the operation group getter method
    if (group.language.go!.clientParams) {
      const clientParams = <Array<Parameter>>group.language.go!.clientParams;
      clientParams.sort(sortParametersByRequired);
      for (const clientParam of values(clientParams)) {
        clientLiterals.push(`${clientParam.language.go!.name}: ${clientParam.language.go!.name}`);
        methodParams.push(`${clientParam.language.go!.name} ${formatParameterTypeName(clientParam)}`);
      }
    }
    text += `// ${group.language.go!.clientName} returns the ${group.language.go!.clientName} associated with this client.\n`;
    text += `func (client *${client}) ${group.language.go!.clientName}(${methodParams}) ${group.language.go!.clientName} {\n`;
    text += `\treturn &${group.operations[0].language.go!.clientName}{${clientLiterals.join(', ')}}\n`;
    text += '}\n\n';
  }

  return text;
}

function getDefaultEndpoint(params?: Parameter[]) {
  for (const param of values(params)) {
    if (param.language.default.name === '$host') {
      return param.clientDefaultValue;
    }
  }
}

function getHostParam(params?: Parameter[]): Parameter | undefined {
  for (const param of values(params)) {
    if (param.language.default.name === '$host') {
      return param;
    }
  }
  return undefined;
}

// the list of packages to import
const imports = new ImportManager();
