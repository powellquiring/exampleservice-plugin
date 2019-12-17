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

package naturallanguageclassifierv1

import (
	"cli-watson-plugin/utils"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/spf13/cobra"
)

var ui terminal.UI

// declare the Authenticator for the service
var Authenticator core.Authenticator
var OutputFormat string
var JMESQuery string

// add a function to return the super-command
func GetNaturalLanguageClassifierV1Command() *cobra.Command {
	ui = terminal.NewStdUI()

	// initialize the authenticator
	ServiceAuthenticator, AuthFactoryErr := core.GetAuthenticatorFromEnvironment("natural_language_classifier")
	utils.HandleError(AuthFactoryErr)

	Authenticator = ServiceAuthenticator

	serviceCommands := []*cobra.Command{
		getClassifyCommand(),
		getClassifyCollectionCommand(),
		getCreateClassifierCommand(),
		getListClassifiersCommand(),
		getGetClassifierCommand(),
		getDeleteClassifierCommand(),
	}

	naturalLanguageClassifierCommand := &cobra.Command{
		Use: "natural-language-classifier-v1 [operation]",
		Aliases: []string{"nlc-v1"},
		Short: "Parent command for Natural Language Classifier",
		Long: "IBM Watson&trade; Natural Language Classifier uses machine learning algorithms to return the top matching predefined classes for short text input. You create and train a classifier to connect predefined classes to example texts so that the service can apply those classes to new inputs.",
	}

	for _, cmd := range serviceCommands {
		naturalLanguageClassifierCommand.AddCommand(cmd)
	}

	return naturalLanguageClassifierCommand 
}
