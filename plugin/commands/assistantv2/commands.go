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

package assistantv2

import (
	"cli-watson-plugin/utils"
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
)

var CreateSessionAssistantID string

func getCreateSessionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-session",
		Short: "Create a session",
		Long: "Create a new session. A session is used to send user input to a skill and receive responses. It also maintains the state of the conversation.",
		Run: CreateSession,
	}

	cmd.Flags().StringVarP(&CreateSessionAssistantID, "assistant_id", "", "", "Unique identifier of the assistant. To find the assistant ID in the Watson Assistant user interface, open the assistant settings and click **API Details**. For information about creating assistants, see the [documentation](https://cloud.ibm.com/docs/services/assistant?topic=assistant-assistant-add#assistant-add-task).**Note:** Currently, the v2 API does not support creating assistants.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("assistant_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func CreateSession(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv2.
		NewAssistantV2(&assistantv2.AssistantV2Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv2.CreateSessionOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "assistant_id" {
			optionsModel.SetAssistantID(CreateSessionAssistantID)
		}
	})

	result, _, responseErr := assistant.CreateSession(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteSessionAssistantID string
var DeleteSessionSessionID string

func getDeleteSessionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-session",
		Short: "Delete session",
		Long: "Deletes a session explicitly before it times out.",
		Run: DeleteSession,
	}

	cmd.Flags().StringVarP(&DeleteSessionAssistantID, "assistant_id", "", "", "Unique identifier of the assistant. To find the assistant ID in the Watson Assistant user interface, open the assistant settings and click **API Details**. For information about creating assistants, see the [documentation](https://cloud.ibm.com/docs/services/assistant?topic=assistant-assistant-add#assistant-add-task).**Note:** Currently, the v2 API does not support creating assistants.")
	cmd.Flags().StringVarP(&DeleteSessionSessionID, "session_id", "", "", "Unique identifier of the session.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")

	cmd.MarkFlagRequired("assistant_id")
	cmd.MarkFlagRequired("session_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteSession(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv2.
		NewAssistantV2(&assistantv2.AssistantV2Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv2.DeleteSessionOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "assistant_id" {
			optionsModel.SetAssistantID(DeleteSessionAssistantID)
		}
		if flag.Name == "session_id" {
			optionsModel.SetSessionID(DeleteSessionSessionID)
		}
	})

	_, responseErr := assistant.DeleteSession(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var MessageAssistantID string
var MessageSessionID string
var MessageInput string
var MessageContext string

func getMessageCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "message",
		Short: "Send user input to assistant",
		Long: "Send user input to an assistant and receive a response.There is no rate limit for this operation.",
		Run: Message,
	}

	cmd.Flags().StringVarP(&MessageAssistantID, "assistant_id", "", "", "Unique identifier of the assistant. To find the assistant ID in the Watson Assistant user interface, open the assistant settings and click **API Details**. For information about creating assistants, see the [documentation](https://cloud.ibm.com/docs/services/assistant?topic=assistant-assistant-add#assistant-add-task).**Note:** Currently, the v2 API does not support creating assistants.")
	cmd.Flags().StringVarP(&MessageSessionID, "session_id", "", "", "Unique identifier of the session.")
	cmd.Flags().StringVarP(&MessageInput, "input", "", "", "An input object that includes the input text.")
	cmd.Flags().StringVarP(&MessageContext, "context", "", "", "State information for the conversation. The context is stored by the assistant on a per-session basis. You can use this property to set or modify context variables, which can also be accessed by dialog nodes.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("assistant_id")
	cmd.MarkFlagRequired("session_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func Message(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	assistant, assistantErr := assistantv2.
		NewAssistantV2(&assistantv2.AssistantV2Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(assistantErr)

	optionsModel := assistantv2.MessageOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "assistant_id" {
			optionsModel.SetAssistantID(MessageAssistantID)
		}
		if flag.Name == "session_id" {
			optionsModel.SetSessionID(MessageSessionID)
		}
		if flag.Name == "input" {
			var input assistantv2.MessageInput
			decodeErr := json.Unmarshal([]byte(MessageInput), &input);
			utils.HandleError(decodeErr)

			optionsModel.SetInput(&input)
		}
		if flag.Name == "context" {
			var context assistantv2.MessageContext
			decodeErr := json.Unmarshal([]byte(MessageContext), &context);
			utils.HandleError(decodeErr)

			optionsModel.SetContext(&context)
		}
	})

	result, _, responseErr := assistant.Message(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}
