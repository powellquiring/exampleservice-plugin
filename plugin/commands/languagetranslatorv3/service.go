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

package languagetranslatorv3

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
var OutputFilename string

// add a function to return the super-command
func GetLanguageTranslatorV3Command() *cobra.Command {
	ui = terminal.NewStdUI()

	// initialize the authenticator
	ServiceAuthenticator, AuthFactoryErr := core.GetAuthenticatorFromEnvironment("language_translator")
	utils.HandleError(AuthFactoryErr)

	Authenticator = ServiceAuthenticator

	serviceCommands := []*cobra.Command{
		getTranslateCommand(),
		getListIdentifiableLanguagesCommand(),
		getIdentifyCommand(),
		getListModelsCommand(),
		getCreateModelCommand(),
		getDeleteModelCommand(),
		getGetModelCommand(),
		getListDocumentsCommand(),
		getTranslateDocumentCommand(),
		getGetDocumentStatusCommand(),
		getDeleteDocumentCommand(),
		getGetTranslatedDocumentCommand(),
	}

	languageTranslatorCommand := &cobra.Command{
		Use: "language-translator-v3 [operation]",
		Aliases: []string{"lt-v3"},
		Short: "Parent command for Language Translator",
		Long: "IBM Watson&trade; Language Translator translates text from one language to another. The service offers multiple IBM provided translation models that you can customize based on your unique terminology and language. Use Language Translator to take news from across the globe and present it in your language, communicate with your customers in their own language, and more.",
	}

	for _, cmd := range serviceCommands {
		languageTranslatorCommand.AddCommand(cmd)
	}

	return languageTranslatorCommand 
}
