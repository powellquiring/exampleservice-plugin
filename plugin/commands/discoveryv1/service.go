/**
 * (C) Copyright IBM Corp. 2019.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package discoveryv1

import (
	"cli-watson-plugin/utils"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/spf13/cobra"
)

var ui terminal.UI

// declare the Authenticator for the service
var Authenticator core.Authenticator
var Version string
var OutputFormat string
var JMESQuery string

// add a function to return the super-command
func GetDiscoveryV1Command() *cobra.Command {
	ui = terminal.NewStdUI()

	// initialize the authenticator
	ServiceAuthenticator, AuthFactoryErr := core.GetAuthenticatorFromEnvironment("discovery")
	utils.HandleError(AuthFactoryErr)

	Authenticator = ServiceAuthenticator

	serviceCommands := []*cobra.Command{
		getCreateEnvironmentCommand(),
		getListEnvironmentsCommand(),
		getGetEnvironmentCommand(),
		getUpdateEnvironmentCommand(),
		getDeleteEnvironmentCommand(),
		getListFieldsCommand(),
		getCreateConfigurationCommand(),
		getListConfigurationsCommand(),
		getGetConfigurationCommand(),
		getUpdateConfigurationCommand(),
		getDeleteConfigurationCommand(),
		getCreateCollectionCommand(),
		getListCollectionsCommand(),
		getGetCollectionCommand(),
		getUpdateCollectionCommand(),
		getDeleteCollectionCommand(),
		getListCollectionFieldsCommand(),
		getListExpansionsCommand(),
		getCreateExpansionsCommand(),
		getDeleteExpansionsCommand(),
		getGetTokenizationDictionaryStatusCommand(),
		getCreateTokenizationDictionaryCommand(),
		getDeleteTokenizationDictionaryCommand(),
		getGetStopwordListStatusCommand(),
		getCreateStopwordListCommand(),
		getDeleteStopwordListCommand(),
		getAddDocumentCommand(),
		getGetDocumentStatusCommand(),
		getUpdateDocumentCommand(),
		getDeleteDocumentCommand(),
		getQueryCommand(),
		getQueryNoticesCommand(),
		getFederatedQueryCommand(),
		getFederatedQueryNoticesCommand(),
		getGetAutocompletionCommand(),
		getListTrainingDataCommand(),
		getAddTrainingDataCommand(),
		getDeleteAllTrainingDataCommand(),
		getGetTrainingDataCommand(),
		getDeleteTrainingDataCommand(),
		getListTrainingExamplesCommand(),
		getCreateTrainingExampleCommand(),
		getDeleteTrainingExampleCommand(),
		getUpdateTrainingExampleCommand(),
		getGetTrainingExampleCommand(),
		getDeleteUserDataCommand(),
		getCreateEventCommand(),
		getQueryLogCommand(),
		getGetMetricsQueryCommand(),
		getGetMetricsQueryEventCommand(),
		getGetMetricsQueryNoResultsCommand(),
		getGetMetricsEventRateCommand(),
		getGetMetricsQueryTokenEventCommand(),
		getListCredentialsCommand(),
		getCreateCredentialsCommand(),
		getGetCredentialsCommand(),
		getUpdateCredentialsCommand(),
		getDeleteCredentialsCommand(),
		getListGatewaysCommand(),
		getCreateGatewayCommand(),
		getGetGatewayCommand(),
		getDeleteGatewayCommand(),
	}

	discoveryCommand := &cobra.Command{
		Use: "discovery-v1 [operation]",
		Short: "Parent command for Discovery",
		Long: "IBM Watson&trade; Discovery is a cognitive search and content analytics engine that you can add to applications to identify patterns, trends and actionable insights to drive better decision-making. Securely unify structured and unstructured data with pre-enriched content, and use a simplified query language to eliminate the need for manual filtering of results.",
	}

	for _, cmd := range serviceCommands {
		discoveryCommand.AddCommand(cmd)
	}

	return discoveryCommand 
}
