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

package naturallanguageunderstandingv1

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
func GetNaturalLanguageUnderstandingV1Command() *cobra.Command {
	ui = terminal.NewStdUI()

	// initialize the authenticator
	ServiceAuthenticator, AuthFactoryErr := core.GetAuthenticatorFromEnvironment("natural_language_understanding")
	utils.HandleError(AuthFactoryErr)

	Authenticator = ServiceAuthenticator

	serviceCommands := []*cobra.Command{
		getAnalyzeCommand(),
		getListModelsCommand(),
		getDeleteModelCommand(),
	}

	naturalLanguageUnderstandingCommand := &cobra.Command{
		Use: "natural-language-understanding-v1 [operation]",
		Aliases: []string{"nlu-v1"},
		Short: "Parent command for Natural Language Understanding",
		Long: "Analyze various features of text content at scale. Provide text, raw HTML, or a public URL and IBM Watson Natural Language Understanding will give you results for the features you request. The service cleans HTML content before analysis by default, so the results can ignore most advertisements and other unwanted content.You can create [custom models](https://cloud.ibm.com/docs/services/natural-language-understanding?topic=natural-language-understanding-customizing) with Watson Knowledge Studio to detect custom entities, relations, and categories in Natural Language Understanding.",
	}

	for _, cmd := range serviceCommands {
		naturalLanguageUnderstandingCommand.AddCommand(cmd)
	}

	return naturalLanguageUnderstandingCommand 
}
