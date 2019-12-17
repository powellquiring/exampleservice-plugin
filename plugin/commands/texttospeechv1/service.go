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

package texttospeechv1

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
var OutputFilename string

// add a function to return the super-command
func GetTextToSpeechV1Command() *cobra.Command {
	ui = terminal.NewStdUI()

	// initialize the authenticator
	ServiceAuthenticator, AuthFactoryErr := core.GetAuthenticatorFromEnvironment("text_to_speech")
	utils.HandleError(AuthFactoryErr)

	Authenticator = ServiceAuthenticator

	serviceCommands := []*cobra.Command{
		getListVoicesCommand(),
		getGetVoiceCommand(),
		getSynthesizeCommand(),
		getGetPronunciationCommand(),
		getCreateVoiceModelCommand(),
		getListVoiceModelsCommand(),
		getUpdateVoiceModelCommand(),
		getGetVoiceModelCommand(),
		getDeleteVoiceModelCommand(),
		getAddWordsCommand(),
		getListWordsCommand(),
		getAddWordCommand(),
		getGetWordCommand(),
		getDeleteWordCommand(),
		getDeleteUserDataCommand(),
	}

	textToSpeechCommand := &cobra.Command{
		Use: "text-to-speech-v1 [operation]",
		Aliases: []string{"tts-v1"},
		Short: "Parent command for Text to Speech",
		Long: "The IBM&reg; Text to Speech service provides APIs that use IBM's speech-synthesis capabilities to synthesize text into natural-sounding speech in a variety of languages, dialects, and voices. The service supports at least one male or female voice, sometimes both, for each language. The audio is streamed back to the client with minimal delay. For speech synthesis, the service supports a synchronous HTTP Representational State Transfer (REST) interface. It also supports a WebSocket interface that provides both plain text and SSML input, including the SSML &lt;mark&gt; element and word timings. SSML is an XML-based markup language that provides text annotation for speech-synthesis applications. The service also offers a customization interface. You can use the interface to define sounds-like or phonetic translations for words. A sounds-like translation consists of one or more words that, when combined, sound like the word. A phonetic translation is based on the SSML phoneme format for representing a word. You can specify a phonetic translation in standard International Phonetic Alphabet (IPA) representation or in the proprietary IBM Symbolic Phonetic Representation (SPR).",
	}

	for _, cmd := range serviceCommands {
		textToSpeechCommand.AddCommand(cmd)
	}

	return textToSpeechCommand 
}
