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
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/watson-developer-cloud/go-sdk/texttospeechv1"
	"io"
	"os"
)


func getListVoicesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-voices",
		Short: "List voices",
		Long: "Lists all voices available for use with the service. The information includes the name, language, gender, and other details about the voice. To see information about a specific voice, use the **Get a voice** method. **See also:** [Listing all available voices](https://cloud.ibm.com/docs/services/text-to-speech?topic=text-to-speech-voices#listVoices).",
		Run: ListVoices,
	}

	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")


	return cmd
}

func ListVoices(cmd *cobra.Command, args []string) {
	textToSpeech, textToSpeechErr := texttospeechv1.
		NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(textToSpeechErr)

	optionsModel := texttospeechv1.ListVoicesOptions{}

	result, _, responseErr := textToSpeech.ListVoices(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetVoiceVoice string
var GetVoiceCustomizationID string

func getGetVoiceCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-voice",
		Short: "Get a voice",
		Long: "Gets information about the specified voice. The information includes the name, language, gender, and other details about the voice. Specify a customization ID to obtain information for a custom voice model that is defined for the language of the specified voice. To list information about all available voices, use the **List voices** method. **See also:** [Listing a specific voice](https://cloud.ibm.com/docs/services/text-to-speech?topic=text-to-speech-voices#listVoice).",
		Run: GetVoice,
	}

	cmd.Flags().StringVarP(&GetVoiceVoice, "voice", "", "", "The voice for which information is to be returned.")
	cmd.Flags().StringVarP(&GetVoiceCustomizationID, "customization_id", "", "", "The customization ID (GUID) of a custom voice model for which information is to be returned. You must make the request with credentials for the instance of the service that owns the custom model. Omit the parameter to see information about the specified voice with no customization.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("voice")

	return cmd
}

func GetVoice(cmd *cobra.Command, args []string) {
	textToSpeech, textToSpeechErr := texttospeechv1.
		NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(textToSpeechErr)

	optionsModel := texttospeechv1.GetVoiceOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "voice" {
			optionsModel.SetVoice(GetVoiceVoice)
		}
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(GetVoiceCustomizationID)
		}
	})

	result, _, responseErr := textToSpeech.GetVoice(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var SynthesizeText string
var SynthesizeAccept string
var SynthesizeVoice string
var SynthesizeCustomizationID string

func getSynthesizeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "synthesize",
		Short: "Synthesize audio",
		Long: "Synthesizes text to audio that is spoken in the specified voice. The service bases its understanding of the language for the input text on the specified voice. Use a voice that matches the language of the input text. The method accepts a maximum of 5 KB of input text in the body of the request, and 8 KB for the URL and headers. The 5 KB limit includes any SSML tags that you specify. The service returns the synthesized audio stream as an array of bytes. **See also:** [The HTTP interface](https://cloud.ibm.com/docs/services/text-to-speech?topic=text-to-speech-usingHTTP#usingHTTP). ### Audio formats (accept types) The service can return audio in the following formats (MIME types).* Where indicated, you can optionally specify the sampling rate (`rate`) of the audio. You must specify a sampling rate for the `audio/l16` and `audio/mulaw` formats. A specified sampling rate must lie in the range of 8 kHz to 192 kHz. Some formats restrict the sampling rate to certain values, as noted.* For the `audio/l16` format, you can optionally specify the endianness (`endianness`) of the audio: `endianness=big-endian` or `endianness=little-endian`. Use the `Accept` header or the `accept` parameter to specify the requested format of the response audio. If you omit an audio format altogether, the service returns the audio in Ogg format with the Opus codec (`audio/ogg;codecs=opus`). The service always returns single-channel audio.* `audio/basic` - The service returns audio with a sampling rate of 8000 Hz.* `audio/flac` - You can optionally specify the `rate` of the audio. The default sampling rate is 22,050 Hz.* `audio/l16` - You must specify the `rate` of the audio. You can optionally specify the `endianness` of the audio. The default endianness is `little-endian`.* `audio/mp3` - You can optionally specify the `rate` of the audio. The default sampling rate is 22,050 Hz.* `audio/mpeg` - You can optionally specify the `rate` of the audio. The default sampling rate is 22,050 Hz.* `audio/mulaw` - You must specify the `rate` of the audio.* `audio/ogg` - The service returns the audio in the `vorbis` codec. You can optionally specify the `rate` of the audio. The default sampling rate is 22,050 Hz.* `audio/ogg;codecs=opus` - You can optionally specify the `rate` of the audio. Only the following values are valid sampling rates: `48000`, `24000`, `16000`, `12000`, or `8000`. If you specify a value other than one of these, the service returns an error. The default sampling rate is 48,000 Hz.* `audio/ogg;codecs=vorbis` - You can optionally specify the `rate` of the audio. The default sampling rate is 22,050 Hz.* `audio/wav` - You can optionally specify the `rate` of the audio. The default sampling rate is 22,050 Hz.* `audio/webm` - The service returns the audio in the `opus` codec. The service returns audio with a sampling rate of 48,000 Hz.* `audio/webm;codecs=opus` - The service returns audio with a sampling rate of 48,000 Hz.* `audio/webm;codecs=vorbis` - You can optionally specify the `rate` of the audio. The default sampling rate is 22,050 Hz. For more information about specifying an audio format, including additional details about some of the formats, see [Audio formats](https://cloud.ibm.com/docs/services/text-to-speech?topic=text-to-speech-audioFormats#audioFormats). ### Warning messages If a request includes invalid query parameters, the service returns a `Warnings` response header that provides messages about the invalid parameters. The warning includes a descriptive message and a list of invalid argument strings. For example, a message such as `'Unknown arguments:'` or `'Unknown url query arguments:'` followed by a list of the form `'{invalid_arg_1}, {invalid_arg_2}.'` The request succeeds despite the warnings.",
		Run: Synthesize,
	}

	cmd.Flags().StringVarP(&SynthesizeText, "text", "", "", "The text to synthesize.")
	cmd.Flags().StringVarP(&SynthesizeAccept, "accept", "", "", "The requested format (MIME type) of the audio. You can use the `Accept` header or the `accept` parameter to specify the audio format. For more information about specifying an audio format, see **Audio formats (accept types)** in the method description.")
	cmd.Flags().StringVarP(&SynthesizeVoice, "voice", "", "", "The voice to use for synthesis.")
	cmd.Flags().StringVarP(&SynthesizeCustomizationID, "customization_id", "", "", "The customization ID (GUID) of a custom voice model to use for the synthesis. If a custom voice model is specified, it is guaranteed to work only if it matches the language of the indicated voice. You must make the request with credentials for the instance of the service that owns the custom model. Omit the parameter to use the specified voice with no customization.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")
	cmd.Flags().StringVarP(&OutputFilename, "output_file", "", "", "Filename/path to write the resulting output to.")

	cmd.MarkFlagRequired("text")
	cmd.MarkFlagRequired("output_file")

	return cmd
}

func Synthesize(cmd *cobra.Command, args []string) {
	textToSpeech, textToSpeechErr := texttospeechv1.
		NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(textToSpeechErr)

	optionsModel := texttospeechv1.SynthesizeOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "text" {
			optionsModel.SetText(SynthesizeText)
		}
		if flag.Name == "accept" {
			optionsModel.SetAccept(SynthesizeAccept)
		}
		if flag.Name == "voice" {
			optionsModel.SetVoice(SynthesizeVoice)
		}
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(SynthesizeCustomizationID)
		}
	})

	result, _, responseErr := textToSpeech.Synthesize(&optionsModel)
	utils.HandleError(responseErr)

	// open a new file
	outFile, outFileErr := os.Create(OutputFilename)
	utils.HandleError(outFileErr)
	defer outFile.Close()

	// write the binary data to the file
	_, fileWriteErr := io.Copy(outFile, result)
	utils.HandleError(fileWriteErr)

	ui.Ok()
	ui.Say("Output written to " + OutputFilename)
}

var GetPronunciationText string
var GetPronunciationVoice string
var GetPronunciationFormat string
var GetPronunciationCustomizationID string

func getGetPronunciationCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-pronunciation",
		Short: "Get pronunciation",
		Long: "Gets the phonetic pronunciation for the specified word. You can request the pronunciation for a specific format. You can also request the pronunciation for a specific voice to see the default translation for the language of that voice or for a specific custom voice model to see the translation for that voice model. **Note:** This method is currently a beta release. **See also:** [Querying a word from a language](https://cloud.ibm.com/docs/services/text-to-speech?topic=text-to-speech-customWords#cuWordsQueryLanguage).",
		Run: GetPronunciation,
	}

	cmd.Flags().StringVarP(&GetPronunciationText, "text", "", "", "The word for which the pronunciation is requested.")
	cmd.Flags().StringVarP(&GetPronunciationVoice, "voice", "", "", "A voice that specifies the language in which the pronunciation is to be returned. All voices for the same language (for example, `en-US`) return the same translation.")
	cmd.Flags().StringVarP(&GetPronunciationFormat, "format", "", "", "The phoneme format in which to return the pronunciation. Omit the parameter to obtain the pronunciation in the default format.")
	cmd.Flags().StringVarP(&GetPronunciationCustomizationID, "customization_id", "", "", "The customization ID (GUID) of a custom voice model for which the pronunciation is to be returned. The language of a specified custom model must match the language of the specified voice. If the word is not defined in the specified custom model, the service returns the default translation for the custom model's language. You must make the request with credentials for the instance of the service that owns the custom model. Omit the parameter to see the translation for the specified voice with no customization.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("text")

	return cmd
}

func GetPronunciation(cmd *cobra.Command, args []string) {
	textToSpeech, textToSpeechErr := texttospeechv1.
		NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(textToSpeechErr)

	optionsModel := texttospeechv1.GetPronunciationOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "text" {
			optionsModel.SetText(GetPronunciationText)
		}
		if flag.Name == "voice" {
			optionsModel.SetVoice(GetPronunciationVoice)
		}
		if flag.Name == "format" {
			optionsModel.SetFormat(GetPronunciationFormat)
		}
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(GetPronunciationCustomizationID)
		}
	})

	result, _, responseErr := textToSpeech.GetPronunciation(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var CreateVoiceModelName string
var CreateVoiceModelLanguage string
var CreateVoiceModelDescription string

func getCreateVoiceModelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-voice-model",
		Short: "Create a custom model",
		Long: "Creates a new empty custom voice model. You must specify a name for the new custom model. You can optionally specify the language and a description for the new model. The model is owned by the instance of the service whose credentials are used to create it. **Note:** This method is currently a beta release. **See also:** [Creating a custom model](https://cloud.ibm.com/docs/services/text-to-speech?topic=text-to-speech-customModels#cuModelsCreate).",
		Run: CreateVoiceModel,
	}

	cmd.Flags().StringVarP(&CreateVoiceModelName, "name", "", "", "The name of the new custom voice model.")
	cmd.Flags().StringVarP(&CreateVoiceModelLanguage, "language", "", "", "The language of the new custom voice model. Omit the parameter to use the the default language, `en-US`.")
	cmd.Flags().StringVarP(&CreateVoiceModelDescription, "description", "", "", "A description of the new custom voice model. Specifying a description is recommended.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("name")

	return cmd
}

func CreateVoiceModel(cmd *cobra.Command, args []string) {
	textToSpeech, textToSpeechErr := texttospeechv1.
		NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(textToSpeechErr)

	optionsModel := texttospeechv1.CreateVoiceModelOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "name" {
			optionsModel.SetName(CreateVoiceModelName)
		}
		if flag.Name == "language" {
			optionsModel.SetLanguage(CreateVoiceModelLanguage)
		}
		if flag.Name == "description" {
			optionsModel.SetDescription(CreateVoiceModelDescription)
		}
	})

	result, _, responseErr := textToSpeech.CreateVoiceModel(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ListVoiceModelsLanguage string

func getListVoiceModelsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-voice-models",
		Short: "List custom models",
		Long: "Lists metadata such as the name and description for all custom voice models that are owned by an instance of the service. Specify a language to list the voice models for that language only. To see the words in addition to the metadata for a specific voice model, use the **List a custom model** method. You must use credentials for the instance of the service that owns a model to list information about it. **Note:** This method is currently a beta release. **See also:** [Querying all custom models](https://cloud.ibm.com/docs/services/text-to-speech?topic=text-to-speech-customModels#cuModelsQueryAll).",
		Run: ListVoiceModels,
	}

	cmd.Flags().StringVarP(&ListVoiceModelsLanguage, "language", "", "", "The language for which custom voice models that are owned by the requesting credentials are to be returned. Omit the parameter to see all custom voice models that are owned by the requester.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")


	return cmd
}

func ListVoiceModels(cmd *cobra.Command, args []string) {
	textToSpeech, textToSpeechErr := texttospeechv1.
		NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(textToSpeechErr)

	optionsModel := texttospeechv1.ListVoiceModelsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "language" {
			optionsModel.SetLanguage(ListVoiceModelsLanguage)
		}
	})

	result, _, responseErr := textToSpeech.ListVoiceModels(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var UpdateVoiceModelCustomizationID string
var UpdateVoiceModelName string
var UpdateVoiceModelDescription string
var UpdateVoiceModelWords string

func getUpdateVoiceModelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-voice-model",
		Short: "Update a custom model",
		Long: "Updates information for the specified custom voice model. You can update metadata such as the name and description of the voice model. You can also update the words in the model and their translations. Adding a new translation for a word that already exists in a custom model overwrites the word's existing translation. A custom model can contain no more than 20,000 entries. You must use credentials for the instance of the service that owns a model to update it. You can define sounds-like or phonetic translations for words. A sounds-like translation consists of one or more words that, when combined, sound like the word. Phonetic translations are based on the SSML phoneme format for representing a word. You can specify them in standard International Phonetic Alphabet (IPA) representation  <code>&lt;phoneme alphabet='ipa' ph='t&#601;m&#712;&#593;to'&gt;&lt;/phoneme&gt;</code>  or in the proprietary IBM Symbolic Phonetic Representation (SPR)  <code>&lt;phoneme alphabet='ibm' ph='1gAstroEntxrYFXs'&gt;&lt;/phoneme&gt;</code> **Note:** This method is currently a beta release. **See also:*** [Updating a custom model](https://cloud.ibm.com/docs/services/text-to-speech?topic=text-to-speech-customModels#cuModelsUpdate)* [Adding words to a Japanese custom model](https://cloud.ibm.com/docs/services/text-to-speech?topic=text-to-speech-customWords#cuJapaneseAdd)* [Understanding customization](https://cloud.ibm.com/docs/services/text-to-speech?topic=text-to-speech-customIntro#customIntro).",
		Run: UpdateVoiceModel,
	}

	cmd.Flags().StringVarP(&UpdateVoiceModelCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom voice model. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&UpdateVoiceModelName, "name", "", "", "A new name for the custom voice model.")
	cmd.Flags().StringVarP(&UpdateVoiceModelDescription, "description", "", "", "A new description for the custom voice model.")
	cmd.Flags().StringVarP(&UpdateVoiceModelWords, "words", "", "", "An array of `Word` objects that provides the words and their translations that are to be added or updated for the custom voice model. Pass an empty array to make no additions or updates.")

	cmd.MarkFlagRequired("customization_id")

	return cmd
}

func UpdateVoiceModel(cmd *cobra.Command, args []string) {
	textToSpeech, textToSpeechErr := texttospeechv1.
		NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(textToSpeechErr)

	optionsModel := texttospeechv1.UpdateVoiceModelOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(UpdateVoiceModelCustomizationID)
		}
		if flag.Name == "name" {
			optionsModel.SetName(UpdateVoiceModelName)
		}
		if flag.Name == "description" {
			optionsModel.SetDescription(UpdateVoiceModelDescription)
		}
		if flag.Name == "words" {
			var words []texttospeechv1.Word
			decodeErr := json.Unmarshal([]byte(UpdateVoiceModelWords), &words);
			utils.HandleError(decodeErr)

			optionsModel.SetWords(words)
		}
	})

	_, responseErr := textToSpeech.UpdateVoiceModel(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var GetVoiceModelCustomizationID string

func getGetVoiceModelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-voice-model",
		Short: "Get a custom model",
		Long: "Gets all information about a specified custom voice model. In addition to metadata such as the name and description of the voice model, the output includes the words and their translations as defined in the model. To see just the metadata for a voice model, use the **List custom models** method. **Note:** This method is currently a beta release. **See also:** [Querying a custom model](https://cloud.ibm.com/docs/services/text-to-speech?topic=text-to-speech-customModels#cuModelsQuery).",
		Run: GetVoiceModel,
	}

	cmd.Flags().StringVarP(&GetVoiceModelCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom voice model. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("customization_id")

	return cmd
}

func GetVoiceModel(cmd *cobra.Command, args []string) {
	textToSpeech, textToSpeechErr := texttospeechv1.
		NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(textToSpeechErr)

	optionsModel := texttospeechv1.GetVoiceModelOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(GetVoiceModelCustomizationID)
		}
	})

	result, _, responseErr := textToSpeech.GetVoiceModel(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteVoiceModelCustomizationID string

func getDeleteVoiceModelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-voice-model",
		Short: "Delete a custom model",
		Long: "Deletes the specified custom voice model. You must use credentials for the instance of the service that owns a model to delete it. **Note:** This method is currently a beta release. **See also:** [Deleting a custom model](https://cloud.ibm.com/docs/services/text-to-speech?topic=text-to-speech-customModels#cuModelsDelete).",
		Run: DeleteVoiceModel,
	}

	cmd.Flags().StringVarP(&DeleteVoiceModelCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom voice model. You must make the request with credentials for the instance of the service that owns the custom model.")

	cmd.MarkFlagRequired("customization_id")

	return cmd
}

func DeleteVoiceModel(cmd *cobra.Command, args []string) {
	textToSpeech, textToSpeechErr := texttospeechv1.
		NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(textToSpeechErr)

	optionsModel := texttospeechv1.DeleteVoiceModelOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(DeleteVoiceModelCustomizationID)
		}
	})

	_, responseErr := textToSpeech.DeleteVoiceModel(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var AddWordsCustomizationID string
var AddWordsWords string

func getAddWordsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "add-words",
		Short: "Add custom words",
		Long: "Adds one or more words and their translations to the specified custom voice model. Adding a new translation for a word that already exists in a custom model overwrites the word's existing translation. A custom model can contain no more than 20,000 entries. You must use credentials for the instance of the service that owns a model to add words to it. You can define sounds-like or phonetic translations for words. A sounds-like translation consists of one or more words that, when combined, sound like the word. Phonetic translations are based on the SSML phoneme format for representing a word. You can specify them in standard International Phonetic Alphabet (IPA) representation  <code>&lt;phoneme alphabet='ipa' ph='t&#601;m&#712;&#593;to'&gt;&lt;/phoneme&gt;</code>  or in the proprietary IBM Symbolic Phonetic Representation (SPR)  <code>&lt;phoneme alphabet='ibm' ph='1gAstroEntxrYFXs'&gt;&lt;/phoneme&gt;</code> **Note:** This method is currently a beta release. **See also:*** [Adding multiple words to a custom model](https://cloud.ibm.com/docs/services/text-to-speech?topic=text-to-speech-customWords#cuWordsAdd)* [Adding words to a Japanese custom model](https://cloud.ibm.com/docs/services/text-to-speech?topic=text-to-speech-customWords#cuJapaneseAdd)* [Understanding customization](https://cloud.ibm.com/docs/services/text-to-speech?topic=text-to-speech-customIntro#customIntro).",
		Run: AddWords,
	}

	cmd.Flags().StringVarP(&AddWordsCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom voice model. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&AddWordsWords, "words", "", "", "The **Add custom words** method accepts an array of `Word` objects. Each object provides a word that is to be added or updated for the custom voice model and the word's translation. The **List custom words** method returns an array of `Word` objects. Each object shows a word and its translation from the custom voice model. The words are listed in alphabetical order, with uppercase letters listed before lowercase letters. The array is empty if the custom model contains no words.")

	cmd.MarkFlagRequired("customization_id")
	cmd.MarkFlagRequired("words")

	return cmd
}

func AddWords(cmd *cobra.Command, args []string) {
	textToSpeech, textToSpeechErr := texttospeechv1.
		NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(textToSpeechErr)

	optionsModel := texttospeechv1.AddWordsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(AddWordsCustomizationID)
		}
		if flag.Name == "words" {
			var words []texttospeechv1.Word
			decodeErr := json.Unmarshal([]byte(AddWordsWords), &words);
			utils.HandleError(decodeErr)

			optionsModel.SetWords(words)
		}
	})

	_, responseErr := textToSpeech.AddWords(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var ListWordsCustomizationID string

func getListWordsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-words",
		Short: "List custom words",
		Long: "Lists all of the words and their translations for the specified custom voice model. The output shows the translations as they are defined in the model. You must use credentials for the instance of the service that owns a model to list its words. **Note:** This method is currently a beta release. **See also:** [Querying all words from a custom model](https://cloud.ibm.com/docs/services/text-to-speech?topic=text-to-speech-customWords#cuWordsQueryModel).",
		Run: ListWords,
	}

	cmd.Flags().StringVarP(&ListWordsCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom voice model. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("customization_id")

	return cmd
}

func ListWords(cmd *cobra.Command, args []string) {
	textToSpeech, textToSpeechErr := texttospeechv1.
		NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(textToSpeechErr)

	optionsModel := texttospeechv1.ListWordsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(ListWordsCustomizationID)
		}
	})

	result, _, responseErr := textToSpeech.ListWords(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var AddWordCustomizationID string
var AddWordWord string
var AddWordTranslation string
var AddWordPartOfSpeech string

func getAddWordCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "add-word",
		Short: "Add a custom word",
		Long: "Adds a single word and its translation to the specified custom voice model. Adding a new translation for a word that already exists in a custom model overwrites the word's existing translation. A custom model can contain no more than 20,000 entries. You must use credentials for the instance of the service that owns a model to add a word to it. You can define sounds-like or phonetic translations for words. A sounds-like translation consists of one or more words that, when combined, sound like the word. Phonetic translations are based on the SSML phoneme format for representing a word. You can specify them in standard International Phonetic Alphabet (IPA) representation  <code>&lt;phoneme alphabet='ipa' ph='t&#601;m&#712;&#593;to'&gt;&lt;/phoneme&gt;</code>  or in the proprietary IBM Symbolic Phonetic Representation (SPR)  <code>&lt;phoneme alphabet='ibm' ph='1gAstroEntxrYFXs'&gt;&lt;/phoneme&gt;</code> **Note:** This method is currently a beta release. **See also:*** [Adding a single word to a custom model](https://cloud.ibm.com/docs/services/text-to-speech?topic=text-to-speech-customWords#cuWordAdd)* [Adding words to a Japanese custom model](https://cloud.ibm.com/docs/services/text-to-speech?topic=text-to-speech-customWords#cuJapaneseAdd)* [Understanding customization](https://cloud.ibm.com/docs/services/text-to-speech?topic=text-to-speech-customIntro#customIntro).",
		Run: AddWord,
	}

	cmd.Flags().StringVarP(&AddWordCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom voice model. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&AddWordWord, "word", "", "", "The word that is to be added or updated for the custom voice model.")
	cmd.Flags().StringVarP(&AddWordTranslation, "translation", "", "", "The phonetic or sounds-like translation for the word. A phonetic translation is based on the SSML format for representing the phonetic string of a word either as an IPA translation or as an IBM SPR translation. A sounds-like is one or more words that, when combined, sound like the word.")
	cmd.Flags().StringVarP(&AddWordPartOfSpeech, "part_of_speech", "", "", "**Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot create multiple entries with different parts of speech for the same word. For more information, see [Working with Japanese entries](https://cloud.ibm.com/docs/services/text-to-speech?topic=text-to-speech-rules#jaNotes).")

	cmd.MarkFlagRequired("customization_id")
	cmd.MarkFlagRequired("word")
	cmd.MarkFlagRequired("translation")

	return cmd
}

func AddWord(cmd *cobra.Command, args []string) {
	textToSpeech, textToSpeechErr := texttospeechv1.
		NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(textToSpeechErr)

	optionsModel := texttospeechv1.AddWordOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(AddWordCustomizationID)
		}
		if flag.Name == "word" {
			optionsModel.SetWord(AddWordWord)
		}
		if flag.Name == "translation" {
			optionsModel.SetTranslation(AddWordTranslation)
		}
		if flag.Name == "part_of_speech" {
			optionsModel.SetPartOfSpeech(AddWordPartOfSpeech)
		}
	})

	_, responseErr := textToSpeech.AddWord(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var GetWordCustomizationID string
var GetWordWord string

func getGetWordCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-word",
		Short: "Get a custom word",
		Long: "Gets the translation for a single word from the specified custom model. The output shows the translation as it is defined in the model. You must use credentials for the instance of the service that owns a model to list its words. **Note:** This method is currently a beta release. **See also:** [Querying a single word from a custom model](https://cloud.ibm.com/docs/services/text-to-speech?topic=text-to-speech-customWords#cuWordQueryModel).",
		Run: GetWord,
	}

	cmd.Flags().StringVarP(&GetWordCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom voice model. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&GetWordWord, "word", "", "", "The word that is to be queried from the custom voice model.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("customization_id")
	cmd.MarkFlagRequired("word")

	return cmd
}

func GetWord(cmd *cobra.Command, args []string) {
	textToSpeech, textToSpeechErr := texttospeechv1.
		NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(textToSpeechErr)

	optionsModel := texttospeechv1.GetWordOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(GetWordCustomizationID)
		}
		if flag.Name == "word" {
			optionsModel.SetWord(GetWordWord)
		}
	})

	result, _, responseErr := textToSpeech.GetWord(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteWordCustomizationID string
var DeleteWordWord string

func getDeleteWordCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-word",
		Short: "Delete a custom word",
		Long: "Deletes a single word from the specified custom voice model. You must use credentials for the instance of the service that owns a model to delete its words. **Note:** This method is currently a beta release. **See also:** [Deleting a word from a custom model](https://cloud.ibm.com/docs/services/text-to-speech?topic=text-to-speech-customWords#cuWordDelete).",
		Run: DeleteWord,
	}

	cmd.Flags().StringVarP(&DeleteWordCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom voice model. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&DeleteWordWord, "word", "", "", "The word that is to be deleted from the custom voice model.")

	cmd.MarkFlagRequired("customization_id")
	cmd.MarkFlagRequired("word")

	return cmd
}

func DeleteWord(cmd *cobra.Command, args []string) {
	textToSpeech, textToSpeechErr := texttospeechv1.
		NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(textToSpeechErr)

	optionsModel := texttospeechv1.DeleteWordOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(DeleteWordCustomizationID)
		}
		if flag.Name == "word" {
			optionsModel.SetWord(DeleteWordWord)
		}
	})

	_, responseErr := textToSpeech.DeleteWord(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var DeleteUserDataCustomerID string

func getDeleteUserDataCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-user-data",
		Short: "Delete labeled data",
		Long: "Deletes all data that is associated with a specified customer ID. The method deletes all data for the customer ID, regardless of the method by which the information was added. The method has no effect if no data is associated with the customer ID. You must issue the request with credentials for the same instance of the service that was used to associate the customer ID with the data. You associate a customer ID with data by passing the `X-Watson-Metadata` header with a request that passes the data. **See also:** [Information security](https://cloud.ibm.com/docs/services/text-to-speech?topic=text-to-speech-information-security#information-security).",
		Run: DeleteUserData,
	}

	cmd.Flags().StringVarP(&DeleteUserDataCustomerID, "customer_id", "", "", "The customer ID for which all data is to be deleted.")

	cmd.MarkFlagRequired("customer_id")

	return cmd
}

func DeleteUserData(cmd *cobra.Command, args []string) {
	textToSpeech, textToSpeechErr := texttospeechv1.
		NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(textToSpeechErr)

	optionsModel := texttospeechv1.DeleteUserDataOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customer_id" {
			optionsModel.SetCustomerID(DeleteUserDataCustomerID)
		}
	})

	_, responseErr := textToSpeech.DeleteUserData(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}
