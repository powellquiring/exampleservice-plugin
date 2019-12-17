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

package assistantv1

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
func GetAssistantV1Command() *cobra.Command {
	ui = terminal.NewStdUI()

	// initialize the authenticator
	ServiceAuthenticator, AuthFactoryErr := core.GetAuthenticatorFromEnvironment("assistant")
	utils.HandleError(AuthFactoryErr)

	Authenticator = ServiceAuthenticator

	serviceCommands := []*cobra.Command{
		getMessageCommand(),
		getListWorkspacesCommand(),
		getCreateWorkspaceCommand(),
		getGetWorkspaceCommand(),
		getUpdateWorkspaceCommand(),
		getDeleteWorkspaceCommand(),
		getListIntentsCommand(),
		getCreateIntentCommand(),
		getGetIntentCommand(),
		getUpdateIntentCommand(),
		getDeleteIntentCommand(),
		getListExamplesCommand(),
		getCreateExampleCommand(),
		getGetExampleCommand(),
		getUpdateExampleCommand(),
		getDeleteExampleCommand(),
		getListCounterexamplesCommand(),
		getCreateCounterexampleCommand(),
		getGetCounterexampleCommand(),
		getUpdateCounterexampleCommand(),
		getDeleteCounterexampleCommand(),
		getListEntitiesCommand(),
		getCreateEntityCommand(),
		getGetEntityCommand(),
		getUpdateEntityCommand(),
		getDeleteEntityCommand(),
		getListMentionsCommand(),
		getListValuesCommand(),
		getCreateValueCommand(),
		getGetValueCommand(),
		getUpdateValueCommand(),
		getDeleteValueCommand(),
		getListSynonymsCommand(),
		getCreateSynonymCommand(),
		getGetSynonymCommand(),
		getUpdateSynonymCommand(),
		getDeleteSynonymCommand(),
		getListDialogNodesCommand(),
		getCreateDialogNodeCommand(),
		getGetDialogNodeCommand(),
		getUpdateDialogNodeCommand(),
		getDeleteDialogNodeCommand(),
		getListLogsCommand(),
		getListAllLogsCommand(),
		getDeleteUserDataCommand(),
	}

	assistantCommand := &cobra.Command{
		Use: "assistant-v1 [operation]",
		Short: "Parent command for Watson Assistant v1",
		Long: "The IBM Watson&trade; Assistant service combines machine learning, natural language understanding, and an integrated dialog editor to create conversation flows between your apps and your users.The Assistant v1 API provides authoring methods your application can use to create or update a workspace.",
	}

	for _, cmd := range serviceCommands {
		assistantCommand.AddCommand(cmd)
	}

	return assistantCommand 
}
