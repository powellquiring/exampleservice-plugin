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
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/watson-developer-cloud/go-sdk/assistantv1"
)

var MessageWorkspaceID string
var MessageInput string
var MessageIntents string
var MessageEntities string
var MessageAlternateIntents bool
var MessageContext string
var MessageOutput string
var MessageNodesVisitedDetails bool

func getMessageCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "message",
		Short: "Get response to user input",
		Long: "Send user input to a workspace and receive a response.**Important:** This method has been superseded by the new v2 runtime API. The v2 API offers significant advantages, including ease of deployment, automatic state management, versioning, and search capabilities. For more information, see the [documentation](https://cloud.ibm.com/docs/services/assistant?topic=assistant-api-overview).There is no rate limit for this operation.",
		Run: Message,
	}

	cmd.Flags().StringVarP(&MessageWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&MessageInput, "input", "", "", "An input object that includes the input text.")
	cmd.Flags().StringVarP(&MessageIntents, "intents", "", "", "Intents to use when evaluating the user input. Include intents from the previous response to continue using those intents rather than trying to recognize intents in the new input.")
	cmd.Flags().StringVarP(&MessageEntities, "entities", "", "", "Entities to use when evaluating the message. Include entities from the previous response to continue using those entities rather than detecting entities in the new input.")
	cmd.Flags().BoolVarP(&MessageAlternateIntents, "alternate_intents", "", false, "Whether to return more than one intent. A value of `true` indicates that all matching intents are returned.")
	cmd.Flags().StringVarP(&MessageContext, "context", "", "", "State information for the conversation. To maintain state, include the context from the previous response.")
	cmd.Flags().StringVarP(&MessageOutput, "message_output", "", "", "An output object that includes the response to the user, the dialog nodes that were triggered, and messages from the log.")
	cmd.Flags().BoolVarP(&MessageNodesVisitedDetails, "nodes_visited_details", "", false, "Whether to include additional diagnostic information about the dialog nodes that were visited during processing of the message.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func Message(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.MessageOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(MessageWorkspaceID)
		}
		if flag.Name == "input" {
			var input assistantv1.MessageInput
			decodeErr := json.Unmarshal([]byte(MessageInput), &input);
			utils.HandleError(decodeErr)

			optionsModel.SetInput(&input)
		}
		if flag.Name == "intents" {
			var intents []assistantv1.RuntimeIntent
			decodeErr := json.Unmarshal([]byte(MessageIntents), &intents);
			utils.HandleError(decodeErr)

			optionsModel.SetIntents(intents)
		}
		if flag.Name == "entities" {
			var entities []assistantv1.RuntimeEntity
			decodeErr := json.Unmarshal([]byte(MessageEntities), &entities);
			utils.HandleError(decodeErr)

			optionsModel.SetEntities(entities)
		}
		if flag.Name == "alternate_intents" {
			optionsModel.SetAlternateIntents(MessageAlternateIntents)
		}
		if flag.Name == "context" {
			var context assistantv1.Context
			decodeErr := json.Unmarshal([]byte(MessageContext), &context);
			utils.HandleError(decodeErr)

			optionsModel.SetContext(&context)
		}
		if flag.Name == "message_output" {
			var message_output assistantv1.OutputData
			decodeErr := json.Unmarshal([]byte(MessageOutput), &message_output);
			utils.HandleError(decodeErr)

			optionsModel.SetOutput(&message_output)
		}
		if flag.Name == "nodes_visited_details" {
			optionsModel.SetNodesVisitedDetails(MessageNodesVisitedDetails)
		}
	})

	result, _, responseErr := assistant.Message(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ListWorkspacesPageLimit int64
var ListWorkspacesSort string
var ListWorkspacesCursor string
var ListWorkspacesIncludeAudit bool

func getListWorkspacesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-workspaces",
		Short: "List workspaces",
		Long: "List the workspaces associated with a Watson Assistant service instance.This operation is limited to 500 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: ListWorkspaces,
	}

	cmd.Flags().Int64VarP(&ListWorkspacesPageLimit, "page_limit", "", 0, "The number of records to return in each page of results.")
	cmd.Flags().StringVarP(&ListWorkspacesSort, "sort", "", "", "The attribute by which returned workspaces will be sorted. To reverse the sort order, prefix the value with a minus sign (`-`).")
	cmd.Flags().StringVarP(&ListWorkspacesCursor, "cursor", "", "", "A token identifying the page of results to retrieve.")
	cmd.Flags().BoolVarP(&ListWorkspacesIncludeAudit, "include_audit", "", false, "Whether to include the audit properties (`created` and `updated` timestamps) in the response.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("version")

	return cmd
}

func ListWorkspaces(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.ListWorkspacesOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "page_limit" {
			optionsModel.SetPageLimit(ListWorkspacesPageLimit)
		}
		if flag.Name == "sort" {
			optionsModel.SetSort(ListWorkspacesSort)
		}
		if flag.Name == "cursor" {
			optionsModel.SetCursor(ListWorkspacesCursor)
		}
		if flag.Name == "include_audit" {
			optionsModel.SetIncludeAudit(ListWorkspacesIncludeAudit)
		}
	})

	result, _, responseErr := assistant.ListWorkspaces(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var CreateWorkspaceName string
var CreateWorkspaceDescription string
var CreateWorkspaceLanguage string
var CreateWorkspaceMetadata string
var CreateWorkspaceLearningOptOut bool
var CreateWorkspaceSystemSettings string
var CreateWorkspaceIntents string
var CreateWorkspaceEntities string
var CreateWorkspaceDialogNodes string
var CreateWorkspaceCounterexamples string

func getCreateWorkspaceCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-workspace",
		Short: "Create workspace",
		Long: "Create a workspace based on component objects. You must provide workspace components defining the content of the new workspace.This operation is limited to 30 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: CreateWorkspace,
	}

	cmd.Flags().StringVarP(&CreateWorkspaceName, "name", "", "", "The name of the workspace. This string cannot contain carriage return, newline, or tab characters.")
	cmd.Flags().StringVarP(&CreateWorkspaceDescription, "description", "", "", "The description of the workspace. This string cannot contain carriage return, newline, or tab characters.")
	cmd.Flags().StringVarP(&CreateWorkspaceLanguage, "language", "", "", "The language of the workspace.")
	cmd.Flags().StringVarP(&CreateWorkspaceMetadata, "metadata", "", "", "Any metadata related to the workspace.")
	cmd.Flags().BoolVarP(&CreateWorkspaceLearningOptOut, "learning_opt_out", "", false, "Whether training data from the workspace (including artifacts such as intents and entities) can be used by IBM for general service improvements. `true` indicates that workspace training data is not to be used.")
	cmd.Flags().StringVarP(&CreateWorkspaceSystemSettings, "system_settings", "", "", "Global settings for the workspace.")
	cmd.Flags().StringVarP(&CreateWorkspaceIntents, "intents", "", "", "An array of objects defining the intents for the workspace.")
	cmd.Flags().StringVarP(&CreateWorkspaceEntities, "entities", "", "", "An array of objects describing the entities for the workspace.")
	cmd.Flags().StringVarP(&CreateWorkspaceDialogNodes, "dialog_nodes", "", "", "An array of objects describing the dialog nodes in the workspace.")
	cmd.Flags().StringVarP(&CreateWorkspaceCounterexamples, "counterexamples", "", "", "An array of objects defining input examples that have been marked as irrelevant input.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("version")

	return cmd
}

func CreateWorkspace(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.CreateWorkspaceOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "name" {
			optionsModel.SetName(CreateWorkspaceName)
		}
		if flag.Name == "description" {
			optionsModel.SetDescription(CreateWorkspaceDescription)
		}
		if flag.Name == "language" {
			optionsModel.SetLanguage(CreateWorkspaceLanguage)
		}
		if flag.Name == "metadata" {
			var metadata map[string]interface{}
			decodeErr := json.Unmarshal([]byte(CreateWorkspaceMetadata), &metadata);
			utils.HandleError(decodeErr)

			optionsModel.SetMetadata(metadata)
		}
		if flag.Name == "learning_opt_out" {
			optionsModel.SetLearningOptOut(CreateWorkspaceLearningOptOut)
		}
		if flag.Name == "system_settings" {
			var system_settings assistantv1.WorkspaceSystemSettings
			decodeErr := json.Unmarshal([]byte(CreateWorkspaceSystemSettings), &system_settings);
			utils.HandleError(decodeErr)

			optionsModel.SetSystemSettings(&system_settings)
		}
		if flag.Name == "intents" {
			var intents []assistantv1.CreateIntent
			decodeErr := json.Unmarshal([]byte(CreateWorkspaceIntents), &intents);
			utils.HandleError(decodeErr)

			optionsModel.SetIntents(intents)
		}
		if flag.Name == "entities" {
			var entities []assistantv1.CreateEntity
			decodeErr := json.Unmarshal([]byte(CreateWorkspaceEntities), &entities);
			utils.HandleError(decodeErr)

			optionsModel.SetEntities(entities)
		}
		if flag.Name == "dialog_nodes" {
			var dialog_nodes []assistantv1.DialogNode
			decodeErr := json.Unmarshal([]byte(CreateWorkspaceDialogNodes), &dialog_nodes);
			utils.HandleError(decodeErr)

			optionsModel.SetDialogNodes(dialog_nodes)
		}
		if flag.Name == "counterexamples" {
			var counterexamples []assistantv1.Counterexample
			decodeErr := json.Unmarshal([]byte(CreateWorkspaceCounterexamples), &counterexamples);
			utils.HandleError(decodeErr)

			optionsModel.SetCounterexamples(counterexamples)
		}
	})

	result, _, responseErr := assistant.CreateWorkspace(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetWorkspaceWorkspaceID string
var GetWorkspaceExport bool
var GetWorkspaceIncludeAudit bool
var GetWorkspaceSort string

func getGetWorkspaceCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-workspace",
		Short: "Get information about a workspace",
		Long: "Get information about a workspace, optionally including all workspace content.With **export**=`false`, this operation is limited to 6000 requests per 5 minutes. With **export**=`true`, the limit is 20 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: GetWorkspace,
	}

	cmd.Flags().StringVarP(&GetWorkspaceWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().BoolVarP(&GetWorkspaceExport, "export", "", false, "Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only information about the element itself. If **export**=`true`, all content, including subelements, is included.")
	cmd.Flags().BoolVarP(&GetWorkspaceIncludeAudit, "include_audit", "", false, "Whether to include the audit properties (`created` and `updated` timestamps) in the response.")
	cmd.Flags().StringVarP(&GetWorkspaceSort, "sort", "", "", "Indicates how the returned workspace data will be sorted. This parameter is valid only if **export**=`true`. Specify `sort=stable` to sort all workspace objects by unique identifier, in ascending alphabetical order.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetWorkspace(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.GetWorkspaceOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(GetWorkspaceWorkspaceID)
		}
		if flag.Name == "export" {
			optionsModel.SetExport(GetWorkspaceExport)
		}
		if flag.Name == "include_audit" {
			optionsModel.SetIncludeAudit(GetWorkspaceIncludeAudit)
		}
		if flag.Name == "sort" {
			optionsModel.SetSort(GetWorkspaceSort)
		}
	})

	result, _, responseErr := assistant.GetWorkspace(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var UpdateWorkspaceWorkspaceID string
var UpdateWorkspaceName string
var UpdateWorkspaceDescription string
var UpdateWorkspaceLanguage string
var UpdateWorkspaceMetadata string
var UpdateWorkspaceLearningOptOut bool
var UpdateWorkspaceSystemSettings string
var UpdateWorkspaceIntents string
var UpdateWorkspaceEntities string
var UpdateWorkspaceDialogNodes string
var UpdateWorkspaceCounterexamples string
var UpdateWorkspaceAppend bool

func getUpdateWorkspaceCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-workspace",
		Short: "Update workspace",
		Long: "Update an existing workspace with new or modified data. You must provide component objects defining the content of the updated workspace.This operation is limited to 30 request per 30 minutes. For more information, see **Rate limiting**.",
		Run: UpdateWorkspace,
	}

	cmd.Flags().StringVarP(&UpdateWorkspaceWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&UpdateWorkspaceName, "name", "", "", "The name of the workspace. This string cannot contain carriage return, newline, or tab characters.")
	cmd.Flags().StringVarP(&UpdateWorkspaceDescription, "description", "", "", "The description of the workspace. This string cannot contain carriage return, newline, or tab characters.")
	cmd.Flags().StringVarP(&UpdateWorkspaceLanguage, "language", "", "", "The language of the workspace.")
	cmd.Flags().StringVarP(&UpdateWorkspaceMetadata, "metadata", "", "", "Any metadata related to the workspace.")
	cmd.Flags().BoolVarP(&UpdateWorkspaceLearningOptOut, "learning_opt_out", "", false, "Whether training data from the workspace (including artifacts such as intents and entities) can be used by IBM for general service improvements. `true` indicates that workspace training data is not to be used.")
	cmd.Flags().StringVarP(&UpdateWorkspaceSystemSettings, "system_settings", "", "", "Global settings for the workspace.")
	cmd.Flags().StringVarP(&UpdateWorkspaceIntents, "intents", "", "", "An array of objects defining the intents for the workspace.")
	cmd.Flags().StringVarP(&UpdateWorkspaceEntities, "entities", "", "", "An array of objects describing the entities for the workspace.")
	cmd.Flags().StringVarP(&UpdateWorkspaceDialogNodes, "dialog_nodes", "", "", "An array of objects describing the dialog nodes in the workspace.")
	cmd.Flags().StringVarP(&UpdateWorkspaceCounterexamples, "counterexamples", "", "", "An array of objects defining input examples that have been marked as irrelevant input.")
	cmd.Flags().BoolVarP(&UpdateWorkspaceAppend, "append", "", false, "Whether the new data is to be appended to the existing data in the workspace. If **append**=`false`, elements included in the new data completely replace the corresponding existing elements, including all subelements. For example, if the new data includes **entities** and **append**=`false`, all existing entities in the workspace are discarded and replaced with the new entities.If **append**=`true`, existing elements are preserved, and the new elements are added. If any elements in the new data collide with existing elements, the update request fails.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func UpdateWorkspace(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.UpdateWorkspaceOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(UpdateWorkspaceWorkspaceID)
		}
		if flag.Name == "name" {
			optionsModel.SetName(UpdateWorkspaceName)
		}
		if flag.Name == "description" {
			optionsModel.SetDescription(UpdateWorkspaceDescription)
		}
		if flag.Name == "language" {
			optionsModel.SetLanguage(UpdateWorkspaceLanguage)
		}
		if flag.Name == "metadata" {
			var metadata map[string]interface{}
			decodeErr := json.Unmarshal([]byte(UpdateWorkspaceMetadata), &metadata);
			utils.HandleError(decodeErr)

			optionsModel.SetMetadata(metadata)
		}
		if flag.Name == "learning_opt_out" {
			optionsModel.SetLearningOptOut(UpdateWorkspaceLearningOptOut)
		}
		if flag.Name == "system_settings" {
			var system_settings assistantv1.WorkspaceSystemSettings
			decodeErr := json.Unmarshal([]byte(UpdateWorkspaceSystemSettings), &system_settings);
			utils.HandleError(decodeErr)

			optionsModel.SetSystemSettings(&system_settings)
		}
		if flag.Name == "intents" {
			var intents []assistantv1.CreateIntent
			decodeErr := json.Unmarshal([]byte(UpdateWorkspaceIntents), &intents);
			utils.HandleError(decodeErr)

			optionsModel.SetIntents(intents)
		}
		if flag.Name == "entities" {
			var entities []assistantv1.CreateEntity
			decodeErr := json.Unmarshal([]byte(UpdateWorkspaceEntities), &entities);
			utils.HandleError(decodeErr)

			optionsModel.SetEntities(entities)
		}
		if flag.Name == "dialog_nodes" {
			var dialog_nodes []assistantv1.DialogNode
			decodeErr := json.Unmarshal([]byte(UpdateWorkspaceDialogNodes), &dialog_nodes);
			utils.HandleError(decodeErr)

			optionsModel.SetDialogNodes(dialog_nodes)
		}
		if flag.Name == "counterexamples" {
			var counterexamples []assistantv1.Counterexample
			decodeErr := json.Unmarshal([]byte(UpdateWorkspaceCounterexamples), &counterexamples);
			utils.HandleError(decodeErr)

			optionsModel.SetCounterexamples(counterexamples)
		}
		if flag.Name == "append" {
			optionsModel.SetAppend(UpdateWorkspaceAppend)
		}
	})

	result, _, responseErr := assistant.UpdateWorkspace(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteWorkspaceWorkspaceID string

func getDeleteWorkspaceCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-workspace",
		Short: "Delete workspace",
		Long: "Delete a workspace from the service instance.This operation is limited to 30 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: DeleteWorkspace,
	}

	cmd.Flags().StringVarP(&DeleteWorkspaceWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteWorkspace(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.DeleteWorkspaceOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(DeleteWorkspaceWorkspaceID)
		}
	})

	_, responseErr := assistant.DeleteWorkspace(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var ListIntentsWorkspaceID string
var ListIntentsExport bool
var ListIntentsPageLimit int64
var ListIntentsSort string
var ListIntentsCursor string
var ListIntentsIncludeAudit bool

func getListIntentsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-intents",
		Short: "List intents",
		Long: "List the intents for a workspace.With **export**=`false`, this operation is limited to 2000 requests per 30 minutes. With **export**=`true`, the limit is 400 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: ListIntents,
	}

	cmd.Flags().StringVarP(&ListIntentsWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().BoolVarP(&ListIntentsExport, "export", "", false, "Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only information about the element itself. If **export**=`true`, all content, including subelements, is included.")
	cmd.Flags().Int64VarP(&ListIntentsPageLimit, "page_limit", "", 0, "The number of records to return in each page of results.")
	cmd.Flags().StringVarP(&ListIntentsSort, "sort", "", "", "The attribute by which returned intents will be sorted. To reverse the sort order, prefix the value with a minus sign (`-`).")
	cmd.Flags().StringVarP(&ListIntentsCursor, "cursor", "", "", "A token identifying the page of results to retrieve.")
	cmd.Flags().BoolVarP(&ListIntentsIncludeAudit, "include_audit", "", false, "Whether to include the audit properties (`created` and `updated` timestamps) in the response.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func ListIntents(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.ListIntentsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(ListIntentsWorkspaceID)
		}
		if flag.Name == "export" {
			optionsModel.SetExport(ListIntentsExport)
		}
		if flag.Name == "page_limit" {
			optionsModel.SetPageLimit(ListIntentsPageLimit)
		}
		if flag.Name == "sort" {
			optionsModel.SetSort(ListIntentsSort)
		}
		if flag.Name == "cursor" {
			optionsModel.SetCursor(ListIntentsCursor)
		}
		if flag.Name == "include_audit" {
			optionsModel.SetIncludeAudit(ListIntentsIncludeAudit)
		}
	})

	result, _, responseErr := assistant.ListIntents(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var CreateIntentWorkspaceID string
var CreateIntentIntent string
var CreateIntentDescription string
var CreateIntentExamples string

func getCreateIntentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-intent",
		Short: "Create intent",
		Long: "Create a new intent.If you want to create multiple intents with a single API call, consider using the **[Update workspace](#update-workspace)** method instead.This operation is limited to 2000 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: CreateIntent,
	}

	cmd.Flags().StringVarP(&CreateIntentWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&CreateIntentIntent, "intent", "", "", "The name of the intent. This string must conform to the following restrictions:- It can contain only Unicode alphanumeric, underscore, hyphen, and dot characters.- It cannot begin with the reserved prefix `sys-`.")
	cmd.Flags().StringVarP(&CreateIntentDescription, "description", "", "", "The description of the intent. This string cannot contain carriage return, newline, or tab characters.")
	cmd.Flags().StringVarP(&CreateIntentExamples, "examples", "", "", "An array of user input examples for the intent.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("intent")
	cmd.MarkFlagRequired("version")

	return cmd
}

func CreateIntent(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.CreateIntentOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(CreateIntentWorkspaceID)
		}
		if flag.Name == "intent" {
			optionsModel.SetIntent(CreateIntentIntent)
		}
		if flag.Name == "description" {
			optionsModel.SetDescription(CreateIntentDescription)
		}
		if flag.Name == "examples" {
			var examples []assistantv1.Example
			decodeErr := json.Unmarshal([]byte(CreateIntentExamples), &examples);
			utils.HandleError(decodeErr)

			optionsModel.SetExamples(examples)
		}
	})

	result, _, responseErr := assistant.CreateIntent(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetIntentWorkspaceID string
var GetIntentIntent string
var GetIntentExport bool
var GetIntentIncludeAudit bool

func getGetIntentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-intent",
		Short: "Get intent",
		Long: "Get information about an intent, optionally including all intent content.With **export**=`false`, this operation is limited to 6000 requests per 5 minutes. With **export**=`true`, the limit is 400 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: GetIntent,
	}

	cmd.Flags().StringVarP(&GetIntentWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&GetIntentIntent, "intent", "", "", "The intent name.")
	cmd.Flags().BoolVarP(&GetIntentExport, "export", "", false, "Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only information about the element itself. If **export**=`true`, all content, including subelements, is included.")
	cmd.Flags().BoolVarP(&GetIntentIncludeAudit, "include_audit", "", false, "Whether to include the audit properties (`created` and `updated` timestamps) in the response.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("intent")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetIntent(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.GetIntentOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(GetIntentWorkspaceID)
		}
		if flag.Name == "intent" {
			optionsModel.SetIntent(GetIntentIntent)
		}
		if flag.Name == "export" {
			optionsModel.SetExport(GetIntentExport)
		}
		if flag.Name == "include_audit" {
			optionsModel.SetIncludeAudit(GetIntentIncludeAudit)
		}
	})

	result, _, responseErr := assistant.GetIntent(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var UpdateIntentWorkspaceID string
var UpdateIntentIntent string
var UpdateIntentNewIntent string
var UpdateIntentNewDescription string
var UpdateIntentNewExamples string

func getUpdateIntentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-intent",
		Short: "Update intent",
		Long: "Update an existing intent with new or modified data. You must provide component objects defining the content of the updated intent.If you want to update multiple intents with a single API call, consider using the **[Update workspace](#update-workspace)** method instead.This operation is limited to 2000 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: UpdateIntent,
	}

	cmd.Flags().StringVarP(&UpdateIntentWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&UpdateIntentIntent, "intent", "", "", "The intent name.")
	cmd.Flags().StringVarP(&UpdateIntentNewIntent, "new_intent", "", "", "The name of the intent. This string must conform to the following restrictions:- It can contain only Unicode alphanumeric, underscore, hyphen, and dot characters.- It cannot begin with the reserved prefix `sys-`.")
	cmd.Flags().StringVarP(&UpdateIntentNewDescription, "new_description", "", "", "The description of the intent. This string cannot contain carriage return, newline, or tab characters.")
	cmd.Flags().StringVarP(&UpdateIntentNewExamples, "new_examples", "", "", "An array of user input examples for the intent.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("intent")
	cmd.MarkFlagRequired("version")

	return cmd
}

func UpdateIntent(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.UpdateIntentOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(UpdateIntentWorkspaceID)
		}
		if flag.Name == "intent" {
			optionsModel.SetIntent(UpdateIntentIntent)
		}
		if flag.Name == "new_intent" {
			optionsModel.SetNewIntent(UpdateIntentNewIntent)
		}
		if flag.Name == "new_description" {
			optionsModel.SetNewDescription(UpdateIntentNewDescription)
		}
		if flag.Name == "new_examples" {
			var new_examples []assistantv1.Example
			decodeErr := json.Unmarshal([]byte(UpdateIntentNewExamples), &new_examples);
			utils.HandleError(decodeErr)

			optionsModel.SetNewExamples(new_examples)
		}
	})

	result, _, responseErr := assistant.UpdateIntent(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteIntentWorkspaceID string
var DeleteIntentIntent string

func getDeleteIntentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-intent",
		Short: "Delete intent",
		Long: "Delete an intent from a workspace.This operation is limited to 2000 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: DeleteIntent,
	}

	cmd.Flags().StringVarP(&DeleteIntentWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&DeleteIntentIntent, "intent", "", "", "The intent name.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("intent")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteIntent(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.DeleteIntentOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(DeleteIntentWorkspaceID)
		}
		if flag.Name == "intent" {
			optionsModel.SetIntent(DeleteIntentIntent)
		}
	})

	_, responseErr := assistant.DeleteIntent(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var ListExamplesWorkspaceID string
var ListExamplesIntent string
var ListExamplesPageLimit int64
var ListExamplesSort string
var ListExamplesCursor string
var ListExamplesIncludeAudit bool

func getListExamplesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-examples",
		Short: "List user input examples",
		Long: "List the user input examples for an intent, optionally including contextual entity mentions.This operation is limited to 2500 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: ListExamples,
	}

	cmd.Flags().StringVarP(&ListExamplesWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&ListExamplesIntent, "intent", "", "", "The intent name.")
	cmd.Flags().Int64VarP(&ListExamplesPageLimit, "page_limit", "", 0, "The number of records to return in each page of results.")
	cmd.Flags().StringVarP(&ListExamplesSort, "sort", "", "", "The attribute by which returned examples will be sorted. To reverse the sort order, prefix the value with a minus sign (`-`).")
	cmd.Flags().StringVarP(&ListExamplesCursor, "cursor", "", "", "A token identifying the page of results to retrieve.")
	cmd.Flags().BoolVarP(&ListExamplesIncludeAudit, "include_audit", "", false, "Whether to include the audit properties (`created` and `updated` timestamps) in the response.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("intent")
	cmd.MarkFlagRequired("version")

	return cmd
}

func ListExamples(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.ListExamplesOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(ListExamplesWorkspaceID)
		}
		if flag.Name == "intent" {
			optionsModel.SetIntent(ListExamplesIntent)
		}
		if flag.Name == "page_limit" {
			optionsModel.SetPageLimit(ListExamplesPageLimit)
		}
		if flag.Name == "sort" {
			optionsModel.SetSort(ListExamplesSort)
		}
		if flag.Name == "cursor" {
			optionsModel.SetCursor(ListExamplesCursor)
		}
		if flag.Name == "include_audit" {
			optionsModel.SetIncludeAudit(ListExamplesIncludeAudit)
		}
	})

	result, _, responseErr := assistant.ListExamples(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var CreateExampleWorkspaceID string
var CreateExampleIntent string
var CreateExampleText string
var CreateExampleMentions string

func getCreateExampleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-example",
		Short: "Create user input example",
		Long: "Add a new user input example to an intent.If you want to add multiple exaples with a single API call, consider using the **[Update intent](#update-intent)** method instead.This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: CreateExample,
	}

	cmd.Flags().StringVarP(&CreateExampleWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&CreateExampleIntent, "intent", "", "", "The intent name.")
	cmd.Flags().StringVarP(&CreateExampleText, "text", "", "", "The text of a user input example. This string must conform to the following restrictions:- It cannot contain carriage return, newline, or tab characters.- It cannot consist of only whitespace characters.")
	cmd.Flags().StringVarP(&CreateExampleMentions, "mentions", "", "", "An array of contextual entity mentions.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("intent")
	cmd.MarkFlagRequired("text")
	cmd.MarkFlagRequired("version")

	return cmd
}

func CreateExample(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.CreateExampleOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(CreateExampleWorkspaceID)
		}
		if flag.Name == "intent" {
			optionsModel.SetIntent(CreateExampleIntent)
		}
		if flag.Name == "text" {
			optionsModel.SetText(CreateExampleText)
		}
		if flag.Name == "mentions" {
			var mentions []assistantv1.Mention
			decodeErr := json.Unmarshal([]byte(CreateExampleMentions), &mentions);
			utils.HandleError(decodeErr)

			optionsModel.SetMentions(mentions)
		}
	})

	result, _, responseErr := assistant.CreateExample(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetExampleWorkspaceID string
var GetExampleIntent string
var GetExampleText string
var GetExampleIncludeAudit bool

func getGetExampleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-example",
		Short: "Get user input example",
		Long: "Get information about a user input example.This operation is limited to 6000 requests per 5 minutes. For more information, see **Rate limiting**.",
		Run: GetExample,
	}

	cmd.Flags().StringVarP(&GetExampleWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&GetExampleIntent, "intent", "", "", "The intent name.")
	cmd.Flags().StringVarP(&GetExampleText, "text", "", "", "The text of the user input example.")
	cmd.Flags().BoolVarP(&GetExampleIncludeAudit, "include_audit", "", false, "Whether to include the audit properties (`created` and `updated` timestamps) in the response.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("intent")
	cmd.MarkFlagRequired("text")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetExample(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.GetExampleOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(GetExampleWorkspaceID)
		}
		if flag.Name == "intent" {
			optionsModel.SetIntent(GetExampleIntent)
		}
		if flag.Name == "text" {
			optionsModel.SetText(GetExampleText)
		}
		if flag.Name == "include_audit" {
			optionsModel.SetIncludeAudit(GetExampleIncludeAudit)
		}
	})

	result, _, responseErr := assistant.GetExample(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var UpdateExampleWorkspaceID string
var UpdateExampleIntent string
var UpdateExampleText string
var UpdateExampleNewText string
var UpdateExampleNewMentions string

func getUpdateExampleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-example",
		Short: "Update user input example",
		Long: "Update the text of a user input example.If you want to update multiple examples with a single API call, consider using the **[Update intent](#update-intent)** method instead.This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: UpdateExample,
	}

	cmd.Flags().StringVarP(&UpdateExampleWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&UpdateExampleIntent, "intent", "", "", "The intent name.")
	cmd.Flags().StringVarP(&UpdateExampleText, "text", "", "", "The text of the user input example.")
	cmd.Flags().StringVarP(&UpdateExampleNewText, "new_text", "", "", "The text of the user input example. This string must conform to the following restrictions:- It cannot contain carriage return, newline, or tab characters.- It cannot consist of only whitespace characters.")
	cmd.Flags().StringVarP(&UpdateExampleNewMentions, "new_mentions", "", "", "An array of contextual entity mentions.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("intent")
	cmd.MarkFlagRequired("text")
	cmd.MarkFlagRequired("version")

	return cmd
}

func UpdateExample(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.UpdateExampleOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(UpdateExampleWorkspaceID)
		}
		if flag.Name == "intent" {
			optionsModel.SetIntent(UpdateExampleIntent)
		}
		if flag.Name == "text" {
			optionsModel.SetText(UpdateExampleText)
		}
		if flag.Name == "new_text" {
			optionsModel.SetNewText(UpdateExampleNewText)
		}
		if flag.Name == "new_mentions" {
			var new_mentions []assistantv1.Mention
			decodeErr := json.Unmarshal([]byte(UpdateExampleNewMentions), &new_mentions);
			utils.HandleError(decodeErr)

			optionsModel.SetNewMentions(new_mentions)
		}
	})

	result, _, responseErr := assistant.UpdateExample(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteExampleWorkspaceID string
var DeleteExampleIntent string
var DeleteExampleText string

func getDeleteExampleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-example",
		Short: "Delete user input example",
		Long: "Delete a user input example from an intent.This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: DeleteExample,
	}

	cmd.Flags().StringVarP(&DeleteExampleWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&DeleteExampleIntent, "intent", "", "", "The intent name.")
	cmd.Flags().StringVarP(&DeleteExampleText, "text", "", "", "The text of the user input example.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("intent")
	cmd.MarkFlagRequired("text")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteExample(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.DeleteExampleOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(DeleteExampleWorkspaceID)
		}
		if flag.Name == "intent" {
			optionsModel.SetIntent(DeleteExampleIntent)
		}
		if flag.Name == "text" {
			optionsModel.SetText(DeleteExampleText)
		}
	})

	_, responseErr := assistant.DeleteExample(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var ListCounterexamplesWorkspaceID string
var ListCounterexamplesPageLimit int64
var ListCounterexamplesSort string
var ListCounterexamplesCursor string
var ListCounterexamplesIncludeAudit bool

func getListCounterexamplesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-counterexamples",
		Short: "List counterexamples",
		Long: "List the counterexamples for a workspace. Counterexamples are examples that have been marked as irrelevant input.This operation is limited to 2500 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: ListCounterexamples,
	}

	cmd.Flags().StringVarP(&ListCounterexamplesWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().Int64VarP(&ListCounterexamplesPageLimit, "page_limit", "", 0, "The number of records to return in each page of results.")
	cmd.Flags().StringVarP(&ListCounterexamplesSort, "sort", "", "", "The attribute by which returned counterexamples will be sorted. To reverse the sort order, prefix the value with a minus sign (`-`).")
	cmd.Flags().StringVarP(&ListCounterexamplesCursor, "cursor", "", "", "A token identifying the page of results to retrieve.")
	cmd.Flags().BoolVarP(&ListCounterexamplesIncludeAudit, "include_audit", "", false, "Whether to include the audit properties (`created` and `updated` timestamps) in the response.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func ListCounterexamples(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.ListCounterexamplesOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(ListCounterexamplesWorkspaceID)
		}
		if flag.Name == "page_limit" {
			optionsModel.SetPageLimit(ListCounterexamplesPageLimit)
		}
		if flag.Name == "sort" {
			optionsModel.SetSort(ListCounterexamplesSort)
		}
		if flag.Name == "cursor" {
			optionsModel.SetCursor(ListCounterexamplesCursor)
		}
		if flag.Name == "include_audit" {
			optionsModel.SetIncludeAudit(ListCounterexamplesIncludeAudit)
		}
	})

	result, _, responseErr := assistant.ListCounterexamples(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var CreateCounterexampleWorkspaceID string
var CreateCounterexampleText string

func getCreateCounterexampleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-counterexample",
		Short: "Create counterexample",
		Long: "Add a new counterexample to a workspace. Counterexamples are examples that have been marked as irrelevant input.If you want to add multiple counterexamples with a single API call, consider using the **[Update workspace](#update-workspace)** method instead.This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: CreateCounterexample,
	}

	cmd.Flags().StringVarP(&CreateCounterexampleWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&CreateCounterexampleText, "text", "", "", "The text of a user input marked as irrelevant input. This string must conform to the following restrictions:- It cannot contain carriage return, newline, or tab characters.- It cannot consist of only whitespace characters.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("text")
	cmd.MarkFlagRequired("version")

	return cmd
}

func CreateCounterexample(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.CreateCounterexampleOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(CreateCounterexampleWorkspaceID)
		}
		if flag.Name == "text" {
			optionsModel.SetText(CreateCounterexampleText)
		}
	})

	result, _, responseErr := assistant.CreateCounterexample(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetCounterexampleWorkspaceID string
var GetCounterexampleText string
var GetCounterexampleIncludeAudit bool

func getGetCounterexampleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-counterexample",
		Short: "Get counterexample",
		Long: "Get information about a counterexample. Counterexamples are examples that have been marked as irrelevant input.This operation is limited to 6000 requests per 5 minutes. For more information, see **Rate limiting**.",
		Run: GetCounterexample,
	}

	cmd.Flags().StringVarP(&GetCounterexampleWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&GetCounterexampleText, "text", "", "", "The text of a user input counterexample (for example, `What are you wearing?`).")
	cmd.Flags().BoolVarP(&GetCounterexampleIncludeAudit, "include_audit", "", false, "Whether to include the audit properties (`created` and `updated` timestamps) in the response.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("text")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetCounterexample(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.GetCounterexampleOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(GetCounterexampleWorkspaceID)
		}
		if flag.Name == "text" {
			optionsModel.SetText(GetCounterexampleText)
		}
		if flag.Name == "include_audit" {
			optionsModel.SetIncludeAudit(GetCounterexampleIncludeAudit)
		}
	})

	result, _, responseErr := assistant.GetCounterexample(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var UpdateCounterexampleWorkspaceID string
var UpdateCounterexampleText string
var UpdateCounterexampleNewText string

func getUpdateCounterexampleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-counterexample",
		Short: "Update counterexample",
		Long: "Update the text of a counterexample. Counterexamples are examples that have been marked as irrelevant input.If you want to update multiple counterexamples with a single API call, consider using the **[Update workspace](#update-workspace)** method instead.This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: UpdateCounterexample,
	}

	cmd.Flags().StringVarP(&UpdateCounterexampleWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&UpdateCounterexampleText, "text", "", "", "The text of a user input counterexample (for example, `What are you wearing?`).")
	cmd.Flags().StringVarP(&UpdateCounterexampleNewText, "new_text", "", "", "The text of a user input marked as irrelevant input. This string must conform to the following restrictions:- It cannot contain carriage return, newline, or tab characters.- It cannot consist of only whitespace characters.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("text")
	cmd.MarkFlagRequired("version")

	return cmd
}

func UpdateCounterexample(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.UpdateCounterexampleOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(UpdateCounterexampleWorkspaceID)
		}
		if flag.Name == "text" {
			optionsModel.SetText(UpdateCounterexampleText)
		}
		if flag.Name == "new_text" {
			optionsModel.SetNewText(UpdateCounterexampleNewText)
		}
	})

	result, _, responseErr := assistant.UpdateCounterexample(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteCounterexampleWorkspaceID string
var DeleteCounterexampleText string

func getDeleteCounterexampleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-counterexample",
		Short: "Delete counterexample",
		Long: "Delete a counterexample from a workspace. Counterexamples are examples that have been marked as irrelevant input.This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: DeleteCounterexample,
	}

	cmd.Flags().StringVarP(&DeleteCounterexampleWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&DeleteCounterexampleText, "text", "", "", "The text of a user input counterexample (for example, `What are you wearing?`).")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("text")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteCounterexample(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.DeleteCounterexampleOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(DeleteCounterexampleWorkspaceID)
		}
		if flag.Name == "text" {
			optionsModel.SetText(DeleteCounterexampleText)
		}
	})

	_, responseErr := assistant.DeleteCounterexample(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var ListEntitiesWorkspaceID string
var ListEntitiesExport bool
var ListEntitiesPageLimit int64
var ListEntitiesSort string
var ListEntitiesCursor string
var ListEntitiesIncludeAudit bool

func getListEntitiesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-entities",
		Short: "List entities",
		Long: "List the entities for a workspace.With **export**=`false`, this operation is limited to 1000 requests per 30 minutes. With **export**=`true`, the limit is 200 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: ListEntities,
	}

	cmd.Flags().StringVarP(&ListEntitiesWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().BoolVarP(&ListEntitiesExport, "export", "", false, "Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only information about the element itself. If **export**=`true`, all content, including subelements, is included.")
	cmd.Flags().Int64VarP(&ListEntitiesPageLimit, "page_limit", "", 0, "The number of records to return in each page of results.")
	cmd.Flags().StringVarP(&ListEntitiesSort, "sort", "", "", "The attribute by which returned entities will be sorted. To reverse the sort order, prefix the value with a minus sign (`-`).")
	cmd.Flags().StringVarP(&ListEntitiesCursor, "cursor", "", "", "A token identifying the page of results to retrieve.")
	cmd.Flags().BoolVarP(&ListEntitiesIncludeAudit, "include_audit", "", false, "Whether to include the audit properties (`created` and `updated` timestamps) in the response.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func ListEntities(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.ListEntitiesOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(ListEntitiesWorkspaceID)
		}
		if flag.Name == "export" {
			optionsModel.SetExport(ListEntitiesExport)
		}
		if flag.Name == "page_limit" {
			optionsModel.SetPageLimit(ListEntitiesPageLimit)
		}
		if flag.Name == "sort" {
			optionsModel.SetSort(ListEntitiesSort)
		}
		if flag.Name == "cursor" {
			optionsModel.SetCursor(ListEntitiesCursor)
		}
		if flag.Name == "include_audit" {
			optionsModel.SetIncludeAudit(ListEntitiesIncludeAudit)
		}
	})

	result, _, responseErr := assistant.ListEntities(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var CreateEntityWorkspaceID string
var CreateEntityEntity string
var CreateEntityDescription string
var CreateEntityMetadata string
var CreateEntityFuzzyMatch bool
var CreateEntityValues string

func getCreateEntityCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-entity",
		Short: "Create entity",
		Long: "Create a new entity, or enable a system entity.If you want to create multiple entities with a single API call, consider using the **[Update workspace](#update-workspace)** method instead.This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: CreateEntity,
	}

	cmd.Flags().StringVarP(&CreateEntityWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&CreateEntityEntity, "entity", "", "", "The name of the entity. This string must conform to the following restrictions:- It can contain only Unicode alphanumeric, underscore, and hyphen characters.- If you specify an entity name beginning with the reserved prefix `sys-`, it must be the name of a system entity that you want to enable. (Any entity content specified with the request is ignored.).")
	cmd.Flags().StringVarP(&CreateEntityDescription, "description", "", "", "The description of the entity. This string cannot contain carriage return, newline, or tab characters.")
	cmd.Flags().StringVarP(&CreateEntityMetadata, "metadata", "", "", "Any metadata related to the entity.")
	cmd.Flags().BoolVarP(&CreateEntityFuzzyMatch, "fuzzy_match", "", false, "Whether to use fuzzy matching for the entity.")
	cmd.Flags().StringVarP(&CreateEntityValues, "values", "", "", "An array of objects describing the entity values.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("entity")
	cmd.MarkFlagRequired("version")

	return cmd
}

func CreateEntity(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.CreateEntityOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(CreateEntityWorkspaceID)
		}
		if flag.Name == "entity" {
			optionsModel.SetEntity(CreateEntityEntity)
		}
		if flag.Name == "description" {
			optionsModel.SetDescription(CreateEntityDescription)
		}
		if flag.Name == "metadata" {
			var metadata map[string]interface{}
			decodeErr := json.Unmarshal([]byte(CreateEntityMetadata), &metadata);
			utils.HandleError(decodeErr)

			optionsModel.SetMetadata(metadata)
		}
		if flag.Name == "fuzzy_match" {
			optionsModel.SetFuzzyMatch(CreateEntityFuzzyMatch)
		}
		if flag.Name == "values" {
			var values []assistantv1.CreateValue
			decodeErr := json.Unmarshal([]byte(CreateEntityValues), &values);
			utils.HandleError(decodeErr)

			optionsModel.SetValues(values)
		}
	})

	result, _, responseErr := assistant.CreateEntity(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetEntityWorkspaceID string
var GetEntityEntity string
var GetEntityExport bool
var GetEntityIncludeAudit bool

func getGetEntityCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-entity",
		Short: "Get entity",
		Long: "Get information about an entity, optionally including all entity content.With **export**=`false`, this operation is limited to 6000 requests per 5 minutes. With **export**=`true`, the limit is 200 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: GetEntity,
	}

	cmd.Flags().StringVarP(&GetEntityWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&GetEntityEntity, "entity", "", "", "The name of the entity.")
	cmd.Flags().BoolVarP(&GetEntityExport, "export", "", false, "Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only information about the element itself. If **export**=`true`, all content, including subelements, is included.")
	cmd.Flags().BoolVarP(&GetEntityIncludeAudit, "include_audit", "", false, "Whether to include the audit properties (`created` and `updated` timestamps) in the response.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("entity")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetEntity(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.GetEntityOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(GetEntityWorkspaceID)
		}
		if flag.Name == "entity" {
			optionsModel.SetEntity(GetEntityEntity)
		}
		if flag.Name == "export" {
			optionsModel.SetExport(GetEntityExport)
		}
		if flag.Name == "include_audit" {
			optionsModel.SetIncludeAudit(GetEntityIncludeAudit)
		}
	})

	result, _, responseErr := assistant.GetEntity(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var UpdateEntityWorkspaceID string
var UpdateEntityEntity string
var UpdateEntityNewEntity string
var UpdateEntityNewDescription string
var UpdateEntityNewMetadata string
var UpdateEntityNewFuzzyMatch bool
var UpdateEntityNewValues string

func getUpdateEntityCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-entity",
		Short: "Update entity",
		Long: "Update an existing entity with new or modified data. You must provide component objects defining the content of the updated entity.If you want to update multiple entities with a single API call, consider using the **[Update workspace](#update-workspace)** method instead.This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: UpdateEntity,
	}

	cmd.Flags().StringVarP(&UpdateEntityWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&UpdateEntityEntity, "entity", "", "", "The name of the entity.")
	cmd.Flags().StringVarP(&UpdateEntityNewEntity, "new_entity", "", "", "The name of the entity. This string must conform to the following restrictions:- It can contain only Unicode alphanumeric, underscore, and hyphen characters.- It cannot begin with the reserved prefix `sys-`.")
	cmd.Flags().StringVarP(&UpdateEntityNewDescription, "new_description", "", "", "The description of the entity. This string cannot contain carriage return, newline, or tab characters.")
	cmd.Flags().StringVarP(&UpdateEntityNewMetadata, "new_metadata", "", "", "Any metadata related to the entity.")
	cmd.Flags().BoolVarP(&UpdateEntityNewFuzzyMatch, "new_fuzzy_match", "", false, "Whether to use fuzzy matching for the entity.")
	cmd.Flags().StringVarP(&UpdateEntityNewValues, "new_values", "", "", "An array of objects describing the entity values.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("entity")
	cmd.MarkFlagRequired("version")

	return cmd
}

func UpdateEntity(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.UpdateEntityOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(UpdateEntityWorkspaceID)
		}
		if flag.Name == "entity" {
			optionsModel.SetEntity(UpdateEntityEntity)
		}
		if flag.Name == "new_entity" {
			optionsModel.SetNewEntity(UpdateEntityNewEntity)
		}
		if flag.Name == "new_description" {
			optionsModel.SetNewDescription(UpdateEntityNewDescription)
		}
		if flag.Name == "new_metadata" {
			var new_metadata map[string]interface{}
			decodeErr := json.Unmarshal([]byte(UpdateEntityNewMetadata), &new_metadata);
			utils.HandleError(decodeErr)

			optionsModel.SetNewMetadata(new_metadata)
		}
		if flag.Name == "new_fuzzy_match" {
			optionsModel.SetNewFuzzyMatch(UpdateEntityNewFuzzyMatch)
		}
		if flag.Name == "new_values" {
			var new_values []assistantv1.CreateValue
			decodeErr := json.Unmarshal([]byte(UpdateEntityNewValues), &new_values);
			utils.HandleError(decodeErr)

			optionsModel.SetNewValues(new_values)
		}
	})

	result, _, responseErr := assistant.UpdateEntity(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteEntityWorkspaceID string
var DeleteEntityEntity string

func getDeleteEntityCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-entity",
		Short: "Delete entity",
		Long: "Delete an entity from a workspace, or disable a system entity.This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: DeleteEntity,
	}

	cmd.Flags().StringVarP(&DeleteEntityWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&DeleteEntityEntity, "entity", "", "", "The name of the entity.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("entity")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteEntity(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.DeleteEntityOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(DeleteEntityWorkspaceID)
		}
		if flag.Name == "entity" {
			optionsModel.SetEntity(DeleteEntityEntity)
		}
	})

	_, responseErr := assistant.DeleteEntity(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var ListMentionsWorkspaceID string
var ListMentionsEntity string
var ListMentionsExport bool
var ListMentionsIncludeAudit bool

func getListMentionsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-mentions",
		Short: "List entity mentions",
		Long: "List mentions for a contextual entity. An entity mention is an occurrence of a contextual entity in the context of an intent user input example.This operation is limited to 200 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: ListMentions,
	}

	cmd.Flags().StringVarP(&ListMentionsWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&ListMentionsEntity, "entity", "", "", "The name of the entity.")
	cmd.Flags().BoolVarP(&ListMentionsExport, "export", "", false, "Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only information about the element itself. If **export**=`true`, all content, including subelements, is included.")
	cmd.Flags().BoolVarP(&ListMentionsIncludeAudit, "include_audit", "", false, "Whether to include the audit properties (`created` and `updated` timestamps) in the response.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("entity")
	cmd.MarkFlagRequired("version")

	return cmd
}

func ListMentions(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.ListMentionsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(ListMentionsWorkspaceID)
		}
		if flag.Name == "entity" {
			optionsModel.SetEntity(ListMentionsEntity)
		}
		if flag.Name == "export" {
			optionsModel.SetExport(ListMentionsExport)
		}
		if flag.Name == "include_audit" {
			optionsModel.SetIncludeAudit(ListMentionsIncludeAudit)
		}
	})

	result, _, responseErr := assistant.ListMentions(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ListValuesWorkspaceID string
var ListValuesEntity string
var ListValuesExport bool
var ListValuesPageLimit int64
var ListValuesSort string
var ListValuesCursor string
var ListValuesIncludeAudit bool

func getListValuesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-values",
		Short: "List entity values",
		Long: "List the values for an entity.This operation is limited to 2500 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: ListValues,
	}

	cmd.Flags().StringVarP(&ListValuesWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&ListValuesEntity, "entity", "", "", "The name of the entity.")
	cmd.Flags().BoolVarP(&ListValuesExport, "export", "", false, "Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only information about the element itself. If **export**=`true`, all content, including subelements, is included.")
	cmd.Flags().Int64VarP(&ListValuesPageLimit, "page_limit", "", 0, "The number of records to return in each page of results.")
	cmd.Flags().StringVarP(&ListValuesSort, "sort", "", "", "The attribute by which returned entity values will be sorted. To reverse the sort order, prefix the value with a minus sign (`-`).")
	cmd.Flags().StringVarP(&ListValuesCursor, "cursor", "", "", "A token identifying the page of results to retrieve.")
	cmd.Flags().BoolVarP(&ListValuesIncludeAudit, "include_audit", "", false, "Whether to include the audit properties (`created` and `updated` timestamps) in the response.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("entity")
	cmd.MarkFlagRequired("version")

	return cmd
}

func ListValues(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.ListValuesOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(ListValuesWorkspaceID)
		}
		if flag.Name == "entity" {
			optionsModel.SetEntity(ListValuesEntity)
		}
		if flag.Name == "export" {
			optionsModel.SetExport(ListValuesExport)
		}
		if flag.Name == "page_limit" {
			optionsModel.SetPageLimit(ListValuesPageLimit)
		}
		if flag.Name == "sort" {
			optionsModel.SetSort(ListValuesSort)
		}
		if flag.Name == "cursor" {
			optionsModel.SetCursor(ListValuesCursor)
		}
		if flag.Name == "include_audit" {
			optionsModel.SetIncludeAudit(ListValuesIncludeAudit)
		}
	})

	result, _, responseErr := assistant.ListValues(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var CreateValueWorkspaceID string
var CreateValueEntity string
var CreateValueValue string
var CreateValueMetadata string
var CreateValueType string
var CreateValueSynonyms []string
var CreateValuePatterns []string

func getCreateValueCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-value",
		Short: "Create entity value",
		Long: "Create a new value for an entity.If you want to create multiple entity values with a single API call, consider using the **[Update entity](#update-entity)** method instead.This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: CreateValue,
	}

	cmd.Flags().StringVarP(&CreateValueWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&CreateValueEntity, "entity", "", "", "The name of the entity.")
	cmd.Flags().StringVarP(&CreateValueValue, "value", "", "", "The text of the entity value. This string must conform to the following restrictions:- It cannot contain carriage return, newline, or tab characters.- It cannot consist of only whitespace characters.")
	cmd.Flags().StringVarP(&CreateValueMetadata, "metadata", "", "", "Any metadata related to the entity value.")
	cmd.Flags().StringVarP(&CreateValueType, "type", "", "", "Specifies the type of entity value.")
	cmd.Flags().StringSliceVarP(&CreateValueSynonyms, "synonyms", "", nil, "An array of synonyms for the entity value. A value can specify either synonyms or patterns (depending on the value type), but not both. A synonym must conform to the following resrictions:- It cannot contain carriage return, newline, or tab characters.- It cannot consist of only whitespace characters.")
	cmd.Flags().StringSliceVarP(&CreateValuePatterns, "patterns", "", nil, "An array of patterns for the entity value. A value can specify either synonyms or patterns (depending on the value type), but not both. A pattern is a regular expression; for more information about how to specify a pattern, see the [documentation](https://cloud.ibm.com/docs/services/assistant?topic=assistant-entities#entities-create-dictionary-based).")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("entity")
	cmd.MarkFlagRequired("value")
	cmd.MarkFlagRequired("version")

	return cmd
}

func CreateValue(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.CreateValueOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(CreateValueWorkspaceID)
		}
		if flag.Name == "entity" {
			optionsModel.SetEntity(CreateValueEntity)
		}
		if flag.Name == "value" {
			optionsModel.SetValue(CreateValueValue)
		}
		if flag.Name == "metadata" {
			var metadata map[string]interface{}
			decodeErr := json.Unmarshal([]byte(CreateValueMetadata), &metadata);
			utils.HandleError(decodeErr)

			optionsModel.SetMetadata(metadata)
		}
		if flag.Name == "type" {
			optionsModel.SetType(CreateValueType)
		}
		if flag.Name == "synonyms" {
			optionsModel.SetSynonyms(CreateValueSynonyms)
		}
		if flag.Name == "patterns" {
			optionsModel.SetPatterns(CreateValuePatterns)
		}
	})

	result, _, responseErr := assistant.CreateValue(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetValueWorkspaceID string
var GetValueEntity string
var GetValueValue string
var GetValueExport bool
var GetValueIncludeAudit bool

func getGetValueCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-value",
		Short: "Get entity value",
		Long: "Get information about an entity value.This operation is limited to 6000 requests per 5 minutes. For more information, see **Rate limiting**.",
		Run: GetValue,
	}

	cmd.Flags().StringVarP(&GetValueWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&GetValueEntity, "entity", "", "", "The name of the entity.")
	cmd.Flags().StringVarP(&GetValueValue, "value", "", "", "The text of the entity value.")
	cmd.Flags().BoolVarP(&GetValueExport, "export", "", false, "Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only information about the element itself. If **export**=`true`, all content, including subelements, is included.")
	cmd.Flags().BoolVarP(&GetValueIncludeAudit, "include_audit", "", false, "Whether to include the audit properties (`created` and `updated` timestamps) in the response.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("entity")
	cmd.MarkFlagRequired("value")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetValue(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.GetValueOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(GetValueWorkspaceID)
		}
		if flag.Name == "entity" {
			optionsModel.SetEntity(GetValueEntity)
		}
		if flag.Name == "value" {
			optionsModel.SetValue(GetValueValue)
		}
		if flag.Name == "export" {
			optionsModel.SetExport(GetValueExport)
		}
		if flag.Name == "include_audit" {
			optionsModel.SetIncludeAudit(GetValueIncludeAudit)
		}
	})

	result, _, responseErr := assistant.GetValue(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var UpdateValueWorkspaceID string
var UpdateValueEntity string
var UpdateValueValue string
var UpdateValueNewValue string
var UpdateValueNewMetadata string
var UpdateValueNewType string
var UpdateValueNewSynonyms []string
var UpdateValueNewPatterns []string

func getUpdateValueCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-value",
		Short: "Update entity value",
		Long: "Update an existing entity value with new or modified data. You must provide component objects defining the content of the updated entity value.If you want to update multiple entity values with a single API call, consider using the **[Update entity](#update-entity)** method instead.This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: UpdateValue,
	}

	cmd.Flags().StringVarP(&UpdateValueWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&UpdateValueEntity, "entity", "", "", "The name of the entity.")
	cmd.Flags().StringVarP(&UpdateValueValue, "value", "", "", "The text of the entity value.")
	cmd.Flags().StringVarP(&UpdateValueNewValue, "new_value", "", "", "The text of the entity value. This string must conform to the following restrictions:- It cannot contain carriage return, newline, or tab characters.- It cannot consist of only whitespace characters.")
	cmd.Flags().StringVarP(&UpdateValueNewMetadata, "new_metadata", "", "", "Any metadata related to the entity value.")
	cmd.Flags().StringVarP(&UpdateValueNewType, "new_type", "", "", "Specifies the type of entity value.")
	cmd.Flags().StringSliceVarP(&UpdateValueNewSynonyms, "new_synonyms", "", nil, "An array of synonyms for the entity value. A value can specify either synonyms or patterns (depending on the value type), but not both. A synonym must conform to the following resrictions:- It cannot contain carriage return, newline, or tab characters.- It cannot consist of only whitespace characters.")
	cmd.Flags().StringSliceVarP(&UpdateValueNewPatterns, "new_patterns", "", nil, "An array of patterns for the entity value. A value can specify either synonyms or patterns (depending on the value type), but not both. A pattern is a regular expression; for more information about how to specify a pattern, see the [documentation](https://cloud.ibm.com/docs/services/assistant?topic=assistant-entities#entities-create-dictionary-based).")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("entity")
	cmd.MarkFlagRequired("value")
	cmd.MarkFlagRequired("version")

	return cmd
}

func UpdateValue(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.UpdateValueOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(UpdateValueWorkspaceID)
		}
		if flag.Name == "entity" {
			optionsModel.SetEntity(UpdateValueEntity)
		}
		if flag.Name == "value" {
			optionsModel.SetValue(UpdateValueValue)
		}
		if flag.Name == "new_value" {
			optionsModel.SetNewValue(UpdateValueNewValue)
		}
		if flag.Name == "new_metadata" {
			var new_metadata map[string]interface{}
			decodeErr := json.Unmarshal([]byte(UpdateValueNewMetadata), &new_metadata);
			utils.HandleError(decodeErr)

			optionsModel.SetNewMetadata(new_metadata)
		}
		if flag.Name == "new_type" {
			optionsModel.SetNewType(UpdateValueNewType)
		}
		if flag.Name == "new_synonyms" {
			optionsModel.SetNewSynonyms(UpdateValueNewSynonyms)
		}
		if flag.Name == "new_patterns" {
			optionsModel.SetNewPatterns(UpdateValueNewPatterns)
		}
	})

	result, _, responseErr := assistant.UpdateValue(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteValueWorkspaceID string
var DeleteValueEntity string
var DeleteValueValue string

func getDeleteValueCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-value",
		Short: "Delete entity value",
		Long: "Delete a value from an entity.This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: DeleteValue,
	}

	cmd.Flags().StringVarP(&DeleteValueWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&DeleteValueEntity, "entity", "", "", "The name of the entity.")
	cmd.Flags().StringVarP(&DeleteValueValue, "value", "", "", "The text of the entity value.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("entity")
	cmd.MarkFlagRequired("value")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteValue(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.DeleteValueOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(DeleteValueWorkspaceID)
		}
		if flag.Name == "entity" {
			optionsModel.SetEntity(DeleteValueEntity)
		}
		if flag.Name == "value" {
			optionsModel.SetValue(DeleteValueValue)
		}
	})

	_, responseErr := assistant.DeleteValue(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var ListSynonymsWorkspaceID string
var ListSynonymsEntity string
var ListSynonymsValue string
var ListSynonymsPageLimit int64
var ListSynonymsSort string
var ListSynonymsCursor string
var ListSynonymsIncludeAudit bool

func getListSynonymsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-synonyms",
		Short: "List entity value synonyms",
		Long: "List the synonyms for an entity value.This operation is limited to 2500 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: ListSynonyms,
	}

	cmd.Flags().StringVarP(&ListSynonymsWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&ListSynonymsEntity, "entity", "", "", "The name of the entity.")
	cmd.Flags().StringVarP(&ListSynonymsValue, "value", "", "", "The text of the entity value.")
	cmd.Flags().Int64VarP(&ListSynonymsPageLimit, "page_limit", "", 0, "The number of records to return in each page of results.")
	cmd.Flags().StringVarP(&ListSynonymsSort, "sort", "", "", "The attribute by which returned entity value synonyms will be sorted. To reverse the sort order, prefix the value with a minus sign (`-`).")
	cmd.Flags().StringVarP(&ListSynonymsCursor, "cursor", "", "", "A token identifying the page of results to retrieve.")
	cmd.Flags().BoolVarP(&ListSynonymsIncludeAudit, "include_audit", "", false, "Whether to include the audit properties (`created` and `updated` timestamps) in the response.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("entity")
	cmd.MarkFlagRequired("value")
	cmd.MarkFlagRequired("version")

	return cmd
}

func ListSynonyms(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.ListSynonymsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(ListSynonymsWorkspaceID)
		}
		if flag.Name == "entity" {
			optionsModel.SetEntity(ListSynonymsEntity)
		}
		if flag.Name == "value" {
			optionsModel.SetValue(ListSynonymsValue)
		}
		if flag.Name == "page_limit" {
			optionsModel.SetPageLimit(ListSynonymsPageLimit)
		}
		if flag.Name == "sort" {
			optionsModel.SetSort(ListSynonymsSort)
		}
		if flag.Name == "cursor" {
			optionsModel.SetCursor(ListSynonymsCursor)
		}
		if flag.Name == "include_audit" {
			optionsModel.SetIncludeAudit(ListSynonymsIncludeAudit)
		}
	})

	result, _, responseErr := assistant.ListSynonyms(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var CreateSynonymWorkspaceID string
var CreateSynonymEntity string
var CreateSynonymValue string
var CreateSynonymSynonym string

func getCreateSynonymCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-synonym",
		Short: "Create entity value synonym",
		Long: "Add a new synonym to an entity value.If you want to create multiple synonyms with a single API call, consider using the **[Update entity](#update-entity)** or **[Update entity value](#update-entity-value)** method instead.This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: CreateSynonym,
	}

	cmd.Flags().StringVarP(&CreateSynonymWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&CreateSynonymEntity, "entity", "", "", "The name of the entity.")
	cmd.Flags().StringVarP(&CreateSynonymValue, "value", "", "", "The text of the entity value.")
	cmd.Flags().StringVarP(&CreateSynonymSynonym, "synonym", "", "", "The text of the synonym. This string must conform to the following restrictions:- It cannot contain carriage return, newline, or tab characters.- It cannot consist of only whitespace characters.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("entity")
	cmd.MarkFlagRequired("value")
	cmd.MarkFlagRequired("synonym")
	cmd.MarkFlagRequired("version")

	return cmd
}

func CreateSynonym(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.CreateSynonymOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(CreateSynonymWorkspaceID)
		}
		if flag.Name == "entity" {
			optionsModel.SetEntity(CreateSynonymEntity)
		}
		if flag.Name == "value" {
			optionsModel.SetValue(CreateSynonymValue)
		}
		if flag.Name == "synonym" {
			optionsModel.SetSynonym(CreateSynonymSynonym)
		}
	})

	result, _, responseErr := assistant.CreateSynonym(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetSynonymWorkspaceID string
var GetSynonymEntity string
var GetSynonymValue string
var GetSynonymSynonym string
var GetSynonymIncludeAudit bool

func getGetSynonymCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-synonym",
		Short: "Get entity value synonym",
		Long: "Get information about a synonym of an entity value.This operation is limited to 6000 requests per 5 minutes. For more information, see **Rate limiting**.",
		Run: GetSynonym,
	}

	cmd.Flags().StringVarP(&GetSynonymWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&GetSynonymEntity, "entity", "", "", "The name of the entity.")
	cmd.Flags().StringVarP(&GetSynonymValue, "value", "", "", "The text of the entity value.")
	cmd.Flags().StringVarP(&GetSynonymSynonym, "synonym", "", "", "The text of the synonym.")
	cmd.Flags().BoolVarP(&GetSynonymIncludeAudit, "include_audit", "", false, "Whether to include the audit properties (`created` and `updated` timestamps) in the response.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("entity")
	cmd.MarkFlagRequired("value")
	cmd.MarkFlagRequired("synonym")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetSynonym(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.GetSynonymOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(GetSynonymWorkspaceID)
		}
		if flag.Name == "entity" {
			optionsModel.SetEntity(GetSynonymEntity)
		}
		if flag.Name == "value" {
			optionsModel.SetValue(GetSynonymValue)
		}
		if flag.Name == "synonym" {
			optionsModel.SetSynonym(GetSynonymSynonym)
		}
		if flag.Name == "include_audit" {
			optionsModel.SetIncludeAudit(GetSynonymIncludeAudit)
		}
	})

	result, _, responseErr := assistant.GetSynonym(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var UpdateSynonymWorkspaceID string
var UpdateSynonymEntity string
var UpdateSynonymValue string
var UpdateSynonymSynonym string
var UpdateSynonymNewSynonym string

func getUpdateSynonymCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-synonym",
		Short: "Update entity value synonym",
		Long: "Update an existing entity value synonym with new text.If you want to update multiple synonyms with a single API call, consider using the **[Update entity](#update-entity)** or **[Update entity value](#update-entity-value)** method instead.This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: UpdateSynonym,
	}

	cmd.Flags().StringVarP(&UpdateSynonymWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&UpdateSynonymEntity, "entity", "", "", "The name of the entity.")
	cmd.Flags().StringVarP(&UpdateSynonymValue, "value", "", "", "The text of the entity value.")
	cmd.Flags().StringVarP(&UpdateSynonymSynonym, "synonym", "", "", "The text of the synonym.")
	cmd.Flags().StringVarP(&UpdateSynonymNewSynonym, "new_synonym", "", "", "The text of the synonym. This string must conform to the following restrictions:- It cannot contain carriage return, newline, or tab characters.- It cannot consist of only whitespace characters.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("entity")
	cmd.MarkFlagRequired("value")
	cmd.MarkFlagRequired("synonym")
	cmd.MarkFlagRequired("version")

	return cmd
}

func UpdateSynonym(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.UpdateSynonymOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(UpdateSynonymWorkspaceID)
		}
		if flag.Name == "entity" {
			optionsModel.SetEntity(UpdateSynonymEntity)
		}
		if flag.Name == "value" {
			optionsModel.SetValue(UpdateSynonymValue)
		}
		if flag.Name == "synonym" {
			optionsModel.SetSynonym(UpdateSynonymSynonym)
		}
		if flag.Name == "new_synonym" {
			optionsModel.SetNewSynonym(UpdateSynonymNewSynonym)
		}
	})

	result, _, responseErr := assistant.UpdateSynonym(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteSynonymWorkspaceID string
var DeleteSynonymEntity string
var DeleteSynonymValue string
var DeleteSynonymSynonym string

func getDeleteSynonymCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-synonym",
		Short: "Delete entity value synonym",
		Long: "Delete a synonym from an entity value.This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: DeleteSynonym,
	}

	cmd.Flags().StringVarP(&DeleteSynonymWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&DeleteSynonymEntity, "entity", "", "", "The name of the entity.")
	cmd.Flags().StringVarP(&DeleteSynonymValue, "value", "", "", "The text of the entity value.")
	cmd.Flags().StringVarP(&DeleteSynonymSynonym, "synonym", "", "", "The text of the synonym.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("entity")
	cmd.MarkFlagRequired("value")
	cmd.MarkFlagRequired("synonym")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteSynonym(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.DeleteSynonymOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(DeleteSynonymWorkspaceID)
		}
		if flag.Name == "entity" {
			optionsModel.SetEntity(DeleteSynonymEntity)
		}
		if flag.Name == "value" {
			optionsModel.SetValue(DeleteSynonymValue)
		}
		if flag.Name == "synonym" {
			optionsModel.SetSynonym(DeleteSynonymSynonym)
		}
	})

	_, responseErr := assistant.DeleteSynonym(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var ListDialogNodesWorkspaceID string
var ListDialogNodesPageLimit int64
var ListDialogNodesSort string
var ListDialogNodesCursor string
var ListDialogNodesIncludeAudit bool

func getListDialogNodesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-dialog-nodes",
		Short: "List dialog nodes",
		Long: "List the dialog nodes for a workspace.This operation is limited to 2500 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: ListDialogNodes,
	}

	cmd.Flags().StringVarP(&ListDialogNodesWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().Int64VarP(&ListDialogNodesPageLimit, "page_limit", "", 0, "The number of records to return in each page of results.")
	cmd.Flags().StringVarP(&ListDialogNodesSort, "sort", "", "", "The attribute by which returned dialog nodes will be sorted. To reverse the sort order, prefix the value with a minus sign (`-`).")
	cmd.Flags().StringVarP(&ListDialogNodesCursor, "cursor", "", "", "A token identifying the page of results to retrieve.")
	cmd.Flags().BoolVarP(&ListDialogNodesIncludeAudit, "include_audit", "", false, "Whether to include the audit properties (`created` and `updated` timestamps) in the response.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func ListDialogNodes(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.ListDialogNodesOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(ListDialogNodesWorkspaceID)
		}
		if flag.Name == "page_limit" {
			optionsModel.SetPageLimit(ListDialogNodesPageLimit)
		}
		if flag.Name == "sort" {
			optionsModel.SetSort(ListDialogNodesSort)
		}
		if flag.Name == "cursor" {
			optionsModel.SetCursor(ListDialogNodesCursor)
		}
		if flag.Name == "include_audit" {
			optionsModel.SetIncludeAudit(ListDialogNodesIncludeAudit)
		}
	})

	result, _, responseErr := assistant.ListDialogNodes(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var CreateDialogNodeWorkspaceID string
var CreateDialogNodeDialogNode string
var CreateDialogNodeDescription string
var CreateDialogNodeConditions string
var CreateDialogNodeParent string
var CreateDialogNodePreviousSibling string
var CreateDialogNodeOutput string
var CreateDialogNodeContext string
var CreateDialogNodeMetadata string
var CreateDialogNodeNextStep string
var CreateDialogNodeTitle string
var CreateDialogNodeType string
var CreateDialogNodeEventName string
var CreateDialogNodeVariable string
var CreateDialogNodeActions string
var CreateDialogNodeDigressIn string
var CreateDialogNodeDigressOut string
var CreateDialogNodeDigressOutSlots string
var CreateDialogNodeUserLabel string

func getCreateDialogNodeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-dialog-node",
		Short: "Create dialog node",
		Long: "Create a new dialog node.If you want to create multiple dialog nodes with a single API call, consider using the **[Update workspace](#update-workspace)** method instead.This operation is limited to 500 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: CreateDialogNode,
	}

	cmd.Flags().StringVarP(&CreateDialogNodeWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&CreateDialogNodeDialogNode, "dialog_node", "", "", "The dialog node ID. This string must conform to the following restrictions:- It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters.")
	cmd.Flags().StringVarP(&CreateDialogNodeDescription, "description", "", "", "The description of the dialog node. This string cannot contain carriage return, newline, or tab characters.")
	cmd.Flags().StringVarP(&CreateDialogNodeConditions, "conditions", "", "", "The condition that will trigger the dialog node. This string cannot contain carriage return, newline, or tab characters.")
	cmd.Flags().StringVarP(&CreateDialogNodeParent, "parent", "", "", "The ID of the parent dialog node. This property is omitted if the dialog node has no parent.")
	cmd.Flags().StringVarP(&CreateDialogNodePreviousSibling, "previous_sibling", "", "", "The ID of the previous sibling dialog node. This property is omitted if the dialog node has no previous sibling.")
	cmd.Flags().StringVarP(&CreateDialogNodeOutput, "create_dialog_node_output", "", "", "The output of the dialog node. For more information about how to specify dialog node output, see the [documentation](https://cloud.ibm.com/docs/services/assistant?topic=assistant-dialog-overview#dialog-overview-responses).")
	cmd.Flags().StringVarP(&CreateDialogNodeContext, "context", "", "", "The context for the dialog node.")
	cmd.Flags().StringVarP(&CreateDialogNodeMetadata, "metadata", "", "", "The metadata for the dialog node.")
	cmd.Flags().StringVarP(&CreateDialogNodeNextStep, "next_step", "", "", "The next step to execute following this dialog node.")
	cmd.Flags().StringVarP(&CreateDialogNodeTitle, "title", "", "", "The alias used to identify the dialog node. This string must conform to the following restrictions:- It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters.")
	cmd.Flags().StringVarP(&CreateDialogNodeType, "type", "", "", "How the dialog node is processed.")
	cmd.Flags().StringVarP(&CreateDialogNodeEventName, "event_name", "", "", "How an `event_handler` node is processed.")
	cmd.Flags().StringVarP(&CreateDialogNodeVariable, "variable", "", "", "The location in the dialog context where output is stored.")
	cmd.Flags().StringVarP(&CreateDialogNodeActions, "actions", "", "", "An array of objects describing any actions to be invoked by the dialog node.")
	cmd.Flags().StringVarP(&CreateDialogNodeDigressIn, "digress_in", "", "", "Whether this top-level dialog node can be digressed into.")
	cmd.Flags().StringVarP(&CreateDialogNodeDigressOut, "digress_out", "", "", "Whether this dialog node can be returned to after a digression.")
	cmd.Flags().StringVarP(&CreateDialogNodeDigressOutSlots, "digress_out_slots", "", "", "Whether the user can digress to top-level nodes while filling out slots.")
	cmd.Flags().StringVarP(&CreateDialogNodeUserLabel, "user_label", "", "", "A label that can be displayed externally to describe the purpose of the node to users.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("dialog_node")
	cmd.MarkFlagRequired("version")

	return cmd
}

func CreateDialogNode(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.CreateDialogNodeOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(CreateDialogNodeWorkspaceID)
		}
		if flag.Name == "dialog_node" {
			optionsModel.SetDialogNode(CreateDialogNodeDialogNode)
		}
		if flag.Name == "description" {
			optionsModel.SetDescription(CreateDialogNodeDescription)
		}
		if flag.Name == "conditions" {
			optionsModel.SetConditions(CreateDialogNodeConditions)
		}
		if flag.Name == "parent" {
			optionsModel.SetParent(CreateDialogNodeParent)
		}
		if flag.Name == "previous_sibling" {
			optionsModel.SetPreviousSibling(CreateDialogNodePreviousSibling)
		}
		if flag.Name == "create_dialog_node_output" {
			var create_dialog_node_output assistantv1.DialogNodeOutput
			decodeErr := json.Unmarshal([]byte(CreateDialogNodeOutput), &create_dialog_node_output);
			utils.HandleError(decodeErr)

			optionsModel.SetOutput(&create_dialog_node_output)
		}
		if flag.Name == "context" {
			var context map[string]interface{}
			decodeErr := json.Unmarshal([]byte(CreateDialogNodeContext), &context);
			utils.HandleError(decodeErr)

			optionsModel.SetContext(context)
		}
		if flag.Name == "metadata" {
			var metadata map[string]interface{}
			decodeErr := json.Unmarshal([]byte(CreateDialogNodeMetadata), &metadata);
			utils.HandleError(decodeErr)

			optionsModel.SetMetadata(metadata)
		}
		if flag.Name == "next_step" {
			var next_step assistantv1.DialogNodeNextStep
			decodeErr := json.Unmarshal([]byte(CreateDialogNodeNextStep), &next_step);
			utils.HandleError(decodeErr)

			optionsModel.SetNextStep(&next_step)
		}
		if flag.Name == "title" {
			optionsModel.SetTitle(CreateDialogNodeTitle)
		}
		if flag.Name == "type" {
			optionsModel.SetType(CreateDialogNodeType)
		}
		if flag.Name == "event_name" {
			optionsModel.SetEventName(CreateDialogNodeEventName)
		}
		if flag.Name == "variable" {
			optionsModel.SetVariable(CreateDialogNodeVariable)
		}
		if flag.Name == "actions" {
			var actions []assistantv1.DialogNodeAction
			decodeErr := json.Unmarshal([]byte(CreateDialogNodeActions), &actions);
			utils.HandleError(decodeErr)

			optionsModel.SetActions(actions)
		}
		if flag.Name == "digress_in" {
			optionsModel.SetDigressIn(CreateDialogNodeDigressIn)
		}
		if flag.Name == "digress_out" {
			optionsModel.SetDigressOut(CreateDialogNodeDigressOut)
		}
		if flag.Name == "digress_out_slots" {
			optionsModel.SetDigressOutSlots(CreateDialogNodeDigressOutSlots)
		}
		if flag.Name == "user_label" {
			optionsModel.SetUserLabel(CreateDialogNodeUserLabel)
		}
	})

	result, _, responseErr := assistant.CreateDialogNode(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetDialogNodeWorkspaceID string
var GetDialogNodeDialogNode string
var GetDialogNodeIncludeAudit bool

func getGetDialogNodeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-dialog-node",
		Short: "Get dialog node",
		Long: "Get information about a dialog node.This operation is limited to 6000 requests per 5 minutes. For more information, see **Rate limiting**.",
		Run: GetDialogNode,
	}

	cmd.Flags().StringVarP(&GetDialogNodeWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&GetDialogNodeDialogNode, "dialog_node", "", "", "The dialog node ID (for example, `get_order`).")
	cmd.Flags().BoolVarP(&GetDialogNodeIncludeAudit, "include_audit", "", false, "Whether to include the audit properties (`created` and `updated` timestamps) in the response.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("dialog_node")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetDialogNode(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.GetDialogNodeOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(GetDialogNodeWorkspaceID)
		}
		if flag.Name == "dialog_node" {
			optionsModel.SetDialogNode(GetDialogNodeDialogNode)
		}
		if flag.Name == "include_audit" {
			optionsModel.SetIncludeAudit(GetDialogNodeIncludeAudit)
		}
	})

	result, _, responseErr := assistant.GetDialogNode(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var UpdateDialogNodeWorkspaceID string
var UpdateDialogNodeDialogNode string
var UpdateDialogNodeNewDialogNode string
var UpdateDialogNodeNewDescription string
var UpdateDialogNodeNewConditions string
var UpdateDialogNodeNewParent string
var UpdateDialogNodeNewPreviousSibling string
var UpdateDialogNodeNewOutput string
var UpdateDialogNodeNewContext string
var UpdateDialogNodeNewMetadata string
var UpdateDialogNodeNewNextStep string
var UpdateDialogNodeNewTitle string
var UpdateDialogNodeNewType string
var UpdateDialogNodeNewEventName string
var UpdateDialogNodeNewVariable string
var UpdateDialogNodeNewActions string
var UpdateDialogNodeNewDigressIn string
var UpdateDialogNodeNewDigressOut string
var UpdateDialogNodeNewDigressOutSlots string
var UpdateDialogNodeNewUserLabel string

func getUpdateDialogNodeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-dialog-node",
		Short: "Update dialog node",
		Long: "Update an existing dialog node with new or modified data.If you want to update multiple dialog nodes with a single API call, consider using the **[Update workspace](#update-workspace)** method instead.This operation is limited to 500 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: UpdateDialogNode,
	}

	cmd.Flags().StringVarP(&UpdateDialogNodeWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&UpdateDialogNodeDialogNode, "dialog_node", "", "", "The dialog node ID (for example, `get_order`).")
	cmd.Flags().StringVarP(&UpdateDialogNodeNewDialogNode, "new_dialog_node", "", "", "The dialog node ID. This string must conform to the following restrictions:- It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters.")
	cmd.Flags().StringVarP(&UpdateDialogNodeNewDescription, "new_description", "", "", "The description of the dialog node. This string cannot contain carriage return, newline, or tab characters.")
	cmd.Flags().StringVarP(&UpdateDialogNodeNewConditions, "new_conditions", "", "", "The condition that will trigger the dialog node. This string cannot contain carriage return, newline, or tab characters.")
	cmd.Flags().StringVarP(&UpdateDialogNodeNewParent, "new_parent", "", "", "The ID of the parent dialog node. This property is omitted if the dialog node has no parent.")
	cmd.Flags().StringVarP(&UpdateDialogNodeNewPreviousSibling, "new_previous_sibling", "", "", "The ID of the previous sibling dialog node. This property is omitted if the dialog node has no previous sibling.")
	cmd.Flags().StringVarP(&UpdateDialogNodeNewOutput, "new_output", "", "", "The output of the dialog node. For more information about how to specify dialog node output, see the [documentation](https://cloud.ibm.com/docs/services/assistant?topic=assistant-dialog-overview#dialog-overview-responses).")
	cmd.Flags().StringVarP(&UpdateDialogNodeNewContext, "new_context", "", "", "The context for the dialog node.")
	cmd.Flags().StringVarP(&UpdateDialogNodeNewMetadata, "new_metadata", "", "", "The metadata for the dialog node.")
	cmd.Flags().StringVarP(&UpdateDialogNodeNewNextStep, "new_next_step", "", "", "The next step to execute following this dialog node.")
	cmd.Flags().StringVarP(&UpdateDialogNodeNewTitle, "new_title", "", "", "The alias used to identify the dialog node. This string must conform to the following restrictions:- It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters.")
	cmd.Flags().StringVarP(&UpdateDialogNodeNewType, "new_type", "", "", "How the dialog node is processed.")
	cmd.Flags().StringVarP(&UpdateDialogNodeNewEventName, "new_event_name", "", "", "How an `event_handler` node is processed.")
	cmd.Flags().StringVarP(&UpdateDialogNodeNewVariable, "new_variable", "", "", "The location in the dialog context where output is stored.")
	cmd.Flags().StringVarP(&UpdateDialogNodeNewActions, "new_actions", "", "", "An array of objects describing any actions to be invoked by the dialog node.")
	cmd.Flags().StringVarP(&UpdateDialogNodeNewDigressIn, "new_digress_in", "", "", "Whether this top-level dialog node can be digressed into.")
	cmd.Flags().StringVarP(&UpdateDialogNodeNewDigressOut, "new_digress_out", "", "", "Whether this dialog node can be returned to after a digression.")
	cmd.Flags().StringVarP(&UpdateDialogNodeNewDigressOutSlots, "new_digress_out_slots", "", "", "Whether the user can digress to top-level nodes while filling out slots.")
	cmd.Flags().StringVarP(&UpdateDialogNodeNewUserLabel, "new_user_label", "", "", "A label that can be displayed externally to describe the purpose of the node to users.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("dialog_node")
	cmd.MarkFlagRequired("version")

	return cmd
}

func UpdateDialogNode(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.UpdateDialogNodeOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(UpdateDialogNodeWorkspaceID)
		}
		if flag.Name == "dialog_node" {
			optionsModel.SetDialogNode(UpdateDialogNodeDialogNode)
		}
		if flag.Name == "new_dialog_node" {
			optionsModel.SetNewDialogNode(UpdateDialogNodeNewDialogNode)
		}
		if flag.Name == "new_description" {
			optionsModel.SetNewDescription(UpdateDialogNodeNewDescription)
		}
		if flag.Name == "new_conditions" {
			optionsModel.SetNewConditions(UpdateDialogNodeNewConditions)
		}
		if flag.Name == "new_parent" {
			optionsModel.SetNewParent(UpdateDialogNodeNewParent)
		}
		if flag.Name == "new_previous_sibling" {
			optionsModel.SetNewPreviousSibling(UpdateDialogNodeNewPreviousSibling)
		}
		if flag.Name == "new_output" {
			var new_output assistantv1.DialogNodeOutput
			decodeErr := json.Unmarshal([]byte(UpdateDialogNodeNewOutput), &new_output);
			utils.HandleError(decodeErr)

			optionsModel.SetNewOutput(&new_output)
		}
		if flag.Name == "new_context" {
			var new_context map[string]interface{}
			decodeErr := json.Unmarshal([]byte(UpdateDialogNodeNewContext), &new_context);
			utils.HandleError(decodeErr)

			optionsModel.SetNewContext(new_context)
		}
		if flag.Name == "new_metadata" {
			var new_metadata map[string]interface{}
			decodeErr := json.Unmarshal([]byte(UpdateDialogNodeNewMetadata), &new_metadata);
			utils.HandleError(decodeErr)

			optionsModel.SetNewMetadata(new_metadata)
		}
		if flag.Name == "new_next_step" {
			var new_next_step assistantv1.DialogNodeNextStep
			decodeErr := json.Unmarshal([]byte(UpdateDialogNodeNewNextStep), &new_next_step);
			utils.HandleError(decodeErr)

			optionsModel.SetNewNextStep(&new_next_step)
		}
		if flag.Name == "new_title" {
			optionsModel.SetNewTitle(UpdateDialogNodeNewTitle)
		}
		if flag.Name == "new_type" {
			optionsModel.SetNewType(UpdateDialogNodeNewType)
		}
		if flag.Name == "new_event_name" {
			optionsModel.SetNewEventName(UpdateDialogNodeNewEventName)
		}
		if flag.Name == "new_variable" {
			optionsModel.SetNewVariable(UpdateDialogNodeNewVariable)
		}
		if flag.Name == "new_actions" {
			var new_actions []assistantv1.DialogNodeAction
			decodeErr := json.Unmarshal([]byte(UpdateDialogNodeNewActions), &new_actions);
			utils.HandleError(decodeErr)

			optionsModel.SetNewActions(new_actions)
		}
		if flag.Name == "new_digress_in" {
			optionsModel.SetNewDigressIn(UpdateDialogNodeNewDigressIn)
		}
		if flag.Name == "new_digress_out" {
			optionsModel.SetNewDigressOut(UpdateDialogNodeNewDigressOut)
		}
		if flag.Name == "new_digress_out_slots" {
			optionsModel.SetNewDigressOutSlots(UpdateDialogNodeNewDigressOutSlots)
		}
		if flag.Name == "new_user_label" {
			optionsModel.SetNewUserLabel(UpdateDialogNodeNewUserLabel)
		}
	})

	result, _, responseErr := assistant.UpdateDialogNode(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteDialogNodeWorkspaceID string
var DeleteDialogNodeDialogNode string

func getDeleteDialogNodeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-dialog-node",
		Short: "Delete dialog node",
		Long: "Delete a dialog node from a workspace.This operation is limited to 500 requests per 30 minutes. For more information, see **Rate limiting**.",
		Run: DeleteDialogNode,
	}

	cmd.Flags().StringVarP(&DeleteDialogNodeWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&DeleteDialogNodeDialogNode, "dialog_node", "", "", "The dialog node ID (for example, `get_order`).")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("dialog_node")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteDialogNode(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.DeleteDialogNodeOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(DeleteDialogNodeWorkspaceID)
		}
		if flag.Name == "dialog_node" {
			optionsModel.SetDialogNode(DeleteDialogNodeDialogNode)
		}
	})

	_, responseErr := assistant.DeleteDialogNode(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var ListLogsWorkspaceID string
var ListLogsSort string
var ListLogsFilter string
var ListLogsPageLimit int64
var ListLogsCursor string

func getListLogsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-logs",
		Short: "List log events in a workspace",
		Long: "List the events from the log of a specific workspace.If **cursor** is not specified, this operation is limited to 40 requests per 30 minutes. If **cursor** is specified, the limit is 120 requests per minute. For more information, see **Rate limiting**.",
		Run: ListLogs,
	}

	cmd.Flags().StringVarP(&ListLogsWorkspaceID, "workspace_id", "", "", "Unique identifier of the workspace.")
	cmd.Flags().StringVarP(&ListLogsSort, "sort", "", "", "How to sort the returned log events. You can sort by **request_timestamp**. To reverse the sort order, prefix the parameter value with a minus sign (`-`).")
	cmd.Flags().StringVarP(&ListLogsFilter, "filter", "", "", "A cacheable parameter that limits the results to those matching the specified filter. For more information, see the [documentation](https://cloud.ibm.com/docs/services/assistant?topic=assistant-filter-reference#filter-reference).")
	cmd.Flags().Int64VarP(&ListLogsPageLimit, "page_limit", "", 0, "The number of records to return in each page of results.")
	cmd.Flags().StringVarP(&ListLogsCursor, "cursor", "", "", "A token identifying the page of results to retrieve.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("workspace_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func ListLogs(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.ListLogsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "workspace_id" {
			optionsModel.SetWorkspaceID(ListLogsWorkspaceID)
		}
		if flag.Name == "sort" {
			optionsModel.SetSort(ListLogsSort)
		}
		if flag.Name == "filter" {
			optionsModel.SetFilter(ListLogsFilter)
		}
		if flag.Name == "page_limit" {
			optionsModel.SetPageLimit(ListLogsPageLimit)
		}
		if flag.Name == "cursor" {
			optionsModel.SetCursor(ListLogsCursor)
		}
	})

	result, _, responseErr := assistant.ListLogs(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ListAllLogsFilter string
var ListAllLogsSort string
var ListAllLogsPageLimit int64
var ListAllLogsCursor string

func getListAllLogsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-all-logs",
		Short: "List log events in all workspaces",
		Long: "List the events from the logs of all workspaces in the service instance.If **cursor** is not specified, this operation is limited to 40 requests per 30 minutes. If **cursor** is specified, the limit is 120 requests per minute. For more information, see **Rate limiting**.",
		Run: ListAllLogs,
	}

	cmd.Flags().StringVarP(&ListAllLogsFilter, "filter", "", "", "A cacheable parameter that limits the results to those matching the specified filter. You must specify a filter query that includes a value for `language`, as well as a value for `workspace_id` or `request.context.metadata.deployment`. For more information, see the [documentation](https://cloud.ibm.com/docs/services/assistant?topic=assistant-filter-reference#filter-reference).")
	cmd.Flags().StringVarP(&ListAllLogsSort, "sort", "", "", "How to sort the returned log events. You can sort by **request_timestamp**. To reverse the sort order, prefix the parameter value with a minus sign (`-`).")
	cmd.Flags().Int64VarP(&ListAllLogsPageLimit, "page_limit", "", 0, "The number of records to return in each page of results.")
	cmd.Flags().StringVarP(&ListAllLogsCursor, "cursor", "", "", "A token identifying the page of results to retrieve.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("filter")
	cmd.MarkFlagRequired("version")

	return cmd
}

func ListAllLogs(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.ListAllLogsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "filter" {
			optionsModel.SetFilter(ListAllLogsFilter)
		}
		if flag.Name == "sort" {
			optionsModel.SetSort(ListAllLogsSort)
		}
		if flag.Name == "page_limit" {
			optionsModel.SetPageLimit(ListAllLogsPageLimit)
		}
		if flag.Name == "cursor" {
			optionsModel.SetCursor(ListAllLogsCursor)
		}
	})

	result, _, responseErr := assistant.ListAllLogs(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteUserDataCustomerID string

func getDeleteUserDataCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-user-data",
		Short: "Delete labeled data",
		Long: "Deletes all data associated with a specified customer ID. The method has no effect if no data is associated with the customer ID. You associate a customer ID with data by passing the `X-Watson-Metadata` header with a request that passes data. For more information about personal data and customer IDs, see [Information security](https://cloud.ibm.com/docs/services/assistant?topic=assistant-information-security#information-security).",
		Run: DeleteUserData,
	}

	cmd.Flags().StringVarP(&DeleteUserDataCustomerID, "customer_id", "", "", "The customer ID for which all data is to be deleted.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")

	cmd.MarkFlagRequired("customer_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteUserData(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv1.
		NewAssistantV1(&assistantv1.AssistantV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv1.DeleteUserDataOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customer_id" {
			optionsModel.SetCustomerID(DeleteUserDataCustomerID)
		}
	})

	_, responseErr := assistant.DeleteUserData(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}
