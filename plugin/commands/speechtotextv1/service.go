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

package speechtotextv1

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
func GetSpeechToTextV1Command() *cobra.Command {
	ui = terminal.NewStdUI()

	// initialize the authenticator
	ServiceAuthenticator, AuthFactoryErr := core.GetAuthenticatorFromEnvironment("speech_to_text")
	utils.HandleError(AuthFactoryErr)

	Authenticator = ServiceAuthenticator

	serviceCommands := []*cobra.Command{
		getListModelsCommand(),
		getGetModelCommand(),
		getRecognizeCommand(),
		getRegisterCallbackCommand(),
		getUnregisterCallbackCommand(),
		getCreateJobCommand(),
		getCheckJobsCommand(),
		getCheckJobCommand(),
		getDeleteJobCommand(),
		getCreateLanguageModelCommand(),
		getListLanguageModelsCommand(),
		getGetLanguageModelCommand(),
		getDeleteLanguageModelCommand(),
		getTrainLanguageModelCommand(),
		getResetLanguageModelCommand(),
		getUpgradeLanguageModelCommand(),
		getListCorporaCommand(),
		getAddCorpusCommand(),
		getGetCorpusCommand(),
		getDeleteCorpusCommand(),
		getListWordsCommand(),
		getAddWordsCommand(),
		getAddWordCommand(),
		getGetWordCommand(),
		getDeleteWordCommand(),
		getListGrammarsCommand(),
		getAddGrammarCommand(),
		getGetGrammarCommand(),
		getDeleteGrammarCommand(),
		getCreateAcousticModelCommand(),
		getListAcousticModelsCommand(),
		getGetAcousticModelCommand(),
		getDeleteAcousticModelCommand(),
		getTrainAcousticModelCommand(),
		getResetAcousticModelCommand(),
		getUpgradeAcousticModelCommand(),
		getListAudioCommand(),
		getAddAudioCommand(),
		getGetAudioCommand(),
		getDeleteAudioCommand(),
		getDeleteUserDataCommand(),
	}

	speechToTextCommand := &cobra.Command{
		Use: "speech-to-text-v1 [operation]",
		Aliases: []string{"stt-v1"},
		Short: "Parent command for Speech to Text",
		Long: "The IBM&reg; Speech to Text service provides APIs that use IBM's speech-recognition capabilities to produce transcripts of spoken audio. The service can transcribe speech from various languages and audio formats. In addition to basic transcription, the service can produce detailed information about many different aspects of the audio. For most languages, the service supports two sampling rates, broadband and narrowband. It returns all JSON response content in the UTF-8 character set. For speech recognition, the service supports synchronous and asynchronous HTTP Representational State Transfer (REST) interfaces. It also supports a WebSocket interface that provides a full-duplex, low-latency communication channel: Clients send requests and audio to the service and receive results over a single connection asynchronously. The service also offers two customization interfaces. Use language model customization to expand the vocabulary of a base model with domain-specific terminology. Use acoustic model customization to adapt a base model for the acoustic characteristics of your audio. For language model customization, the service also supports grammars. A grammar is a formal language specification that lets you restrict the phrases that the service can recognize. Language model customization is generally available for production use with most supported languages. Acoustic model customization is beta functionality that is available for all supported languages. ",
	}

	for _, cmd := range serviceCommands {
		speechToTextCommand.AddCommand(cmd)
	}

	return speechToTextCommand 
}
