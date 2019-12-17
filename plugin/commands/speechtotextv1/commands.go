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
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/watson-developer-cloud/go-sdk/speechtotextv1"
	"os"
)


func getListModelsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-models",
		Short: "List models",
		Long: "Lists all language models that are available for use with the service. The information includes the name of the model and its minimum sampling rate in Hertz, among other things. **See also:** [Languages and models](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-models#models).",
		Run: ListModels,
	}

	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")


	return cmd
}

func ListModels(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.ListModelsOptions{}

	result, _, responseErr := speechToText.ListModels(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetModelModelID string

func getGetModelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-model",
		Short: "Get a model",
		Long: "Gets information for a single specified language model that is available for use with the service. The information includes the name of the model and its minimum sampling rate in Hertz, among other things. **See also:** [Languages and models](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-models#models).",
		Run: GetModel,
	}

	cmd.Flags().StringVarP(&GetModelModelID, "model_id", "", "", "The identifier of the model in the form of its name from the output of the **Get a model** method.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("model_id")

	return cmd
}

func GetModel(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.GetModelOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "model_id" {
			optionsModel.SetModelID(GetModelModelID)
		}
	})

	result, _, responseErr := speechToText.GetModel(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var RecognizeAudio string
var RecognizeContentType string
var RecognizeModel string
var RecognizeLanguageCustomizationID string
var RecognizeAcousticCustomizationID string
var RecognizeBaseModelVersion string
var RecognizeCustomizationWeight float64
var RecognizeInactivityTimeout int64
var RecognizeKeywords []string
var RecognizeKeywordsThreshold float32
var RecognizeMaxAlternatives int64
var RecognizeWordAlternativesThreshold float32
var RecognizeWordConfidence bool
var RecognizeTimestamps bool
var RecognizeProfanityFilter bool
var RecognizeSmartFormatting bool
var RecognizeSpeakerLabels bool
var RecognizeCustomizationID string
var RecognizeGrammarName string
var RecognizeRedaction bool
var RecognizeAudioMetrics bool

func getRecognizeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "recognize",
		Short: "Recognize audio",
		Long: "Sends audio and returns transcription results for a recognition request. You can pass a maximum of 100 MB and a minimum of 100 bytes of audio with a request. The service automatically detects the endianness of the incoming audio and, for audio that includes multiple channels, downmixes the audio to one-channel mono during transcoding. The method returns only final results; to enable interim results, use the WebSocket API. (With the `curl` command, use the `--data-binary` option to upload the file for the request.) **See also:** [Making a basic HTTP request](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-http#HTTP-basic). ### Streaming mode For requests to transcribe live audio as it becomes available, you must set the `Transfer-Encoding` header to `chunked` to use streaming mode. In streaming mode, the service closes the connection (status code 408) if it does not receive at least 15 seconds of audio (including silence) in any 30-second period. The service also closes the connection (status code 400) if it detects no speech for `inactivity_timeout` seconds of streaming audio; use the `inactivity_timeout` parameter to change the default of 30 seconds. **See also:*** [Audio transmission](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-input#transmission)* [Timeouts](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-input#timeouts) ### Audio formats (content types) The service accepts audio in the following formats (MIME types).* For formats that are labeled **Required**, you must use the `Content-Type` header with the request to specify the format of the audio.* For all other formats, you can omit the `Content-Type` header or specify `application/octet-stream` with the header to have the service automatically detect the format of the audio. (With the `curl` command, you can specify either `'Content-Type:'` or `'Content-Type: application/octet-stream'`.) Where indicated, the format that you specify must include the sampling rate and can optionally include the number of channels and the endianness of the audio.* `audio/alaw` (**Required.** Specify the sampling rate (`rate`) of the audio.)* `audio/basic` (**Required.** Use only with narrowband models.)* `audio/flac`* `audio/g729` (Use only with narrowband models.)* `audio/l16` (**Required.** Specify the sampling rate (`rate`) and optionally the number of channels (`channels`) and endianness (`endianness`) of the audio.)* `audio/mp3`* `audio/mpeg`* `audio/mulaw` (**Required.** Specify the sampling rate (`rate`) of the audio.)* `audio/ogg` (The service automatically detects the codec of the input audio.)* `audio/ogg;codecs=opus`* `audio/ogg;codecs=vorbis`* `audio/wav` (Provide audio with a maximum of nine channels.)* `audio/webm` (The service automatically detects the codec of the input audio.)* `audio/webm;codecs=opus`* `audio/webm;codecs=vorbis` The sampling rate of the audio must match the sampling rate of the model for the recognition request: for broadband models, at least 16 kHz; for narrowband models, at least 8 kHz. If the sampling rate of the audio is higher than the minimum required rate, the service down-samples the audio to the appropriate rate. If the sampling rate of the audio is lower than the minimum required rate, the request fails. **See also:** [Audio formats](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-audio-formats#audio-formats). ### Multipart speech recognition **Note:** The Watson SDKs do not support multipart speech recognition. The HTTP `POST` method of the service also supports multipart speech recognition. With multipart requests, you pass all audio data as multipart form data. You specify some parameters as request headers and query parameters, but you pass JSON metadata as form data to control most aspects of the transcription. You can use multipart recognition to pass multiple audio files with a single request. Use the multipart approach with browsers for which JavaScript is disabled or when the parameters used with the request are greater than the 8 KB limit imposed by most HTTP servers and proxies. You can encounter this limit, for example, if you want to spot a very large number of keywords. **See also:** [Making a multipart HTTP request](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-http#HTTP-multi).",
		Run: Recognize,
	}

	cmd.Flags().StringVarP(&RecognizeAudio, "audio", "", "", "The audio to transcribe.")
	cmd.Flags().StringVarP(&RecognizeContentType, "content_type", "", "", "The format (MIME type) of the audio. For more information about specifying an audio format, see **Audio formats (content types)** in the method description.")
	cmd.Flags().StringVarP(&RecognizeModel, "model", "", "", "The identifier of the model that is to be used for the recognition request. See [Languages and models](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-models#models).")
	cmd.Flags().StringVarP(&RecognizeLanguageCustomizationID, "language_customization_id", "", "", "The customization ID (GUID) of a custom language model that is to be used with the recognition request. The base model of the specified custom language model must match the model specified with the `model` parameter. You must make the request with credentials for the instance of the service that owns the custom model. By default, no custom language model is used. See [Custom models](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-input#custom-input). **Note:** Use this parameter instead of the deprecated `customization_id` parameter.")
	cmd.Flags().StringVarP(&RecognizeAcousticCustomizationID, "acoustic_customization_id", "", "", "The customization ID (GUID) of a custom acoustic model that is to be used with the recognition request. The base model of the specified custom acoustic model must match the model specified with the `model` parameter. You must make the request with credentials for the instance of the service that owns the custom model. By default, no custom acoustic model is used. See [Custom models](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-input#custom-input).")
	cmd.Flags().StringVarP(&RecognizeBaseModelVersion, "base_model_version", "", "", "The version of the specified base model that is to be used with the recognition request. Multiple versions of a base model can exist when a model is updated for internal improvements. The parameter is intended primarily for use with custom models that have been upgraded for a new base model. The default value depends on whether the parameter is used with or without a custom model. See [Base model version](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-input#version).")
	cmd.Flags().Float64VarP(&RecognizeCustomizationWeight, "customization_weight", "", 0, "If you specify the customization ID (GUID) of a custom language model with the recognition request, the customization weight tells the service how much weight to give to words from the custom language model compared to those from the base model for the current request. Specify a value between 0.0 and 1.0. Unless a different customization weight was specified for the custom model when it was trained, the default value is 0.3. A customization weight that you specify overrides a weight that was specified when the custom model was trained. The default value yields the best performance in general. Assign a higher value if your audio makes frequent use of OOV words from the custom model. Use caution when setting the weight: a higher value can improve the accuracy of phrases from the custom model's domain, but it can negatively affect performance on non-domain phrases. See [Custom models](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-input#custom-input).")
	cmd.Flags().Int64VarP(&RecognizeInactivityTimeout, "inactivity_timeout", "", 0, "The time in seconds after which, if only silence (no speech) is detected in streaming audio, the connection is closed with a 400 error. The parameter is useful for stopping audio submission from a live microphone when a user simply walks away. Use `-1` for infinity. See [Inactivity timeout](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-input#timeouts-inactivity).")
	cmd.Flags().StringSliceVarP(&RecognizeKeywords, "keywords", "", nil, "An array of keyword strings to spot in the audio. Each keyword string can include one or more string tokens. Keywords are spotted only in the final results, not in interim hypotheses. If you specify any keywords, you must also specify a keywords threshold. You can spot a maximum of 1000 keywords. Omit the parameter or specify an empty array if you do not need to spot keywords. See [Keyword spotting](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-output#keyword_spotting).")
	cmd.Flags().Float32VarP(&RecognizeKeywordsThreshold, "keywords_threshold", "", 0, "A confidence value that is the lower bound for spotting a keyword. A word is considered to match a keyword if its confidence is greater than or equal to the threshold. Specify a probability between 0.0 and 1.0. If you specify a threshold, you must also specify one or more keywords. The service performs no keyword spotting if you omit either parameter. See [Keyword spotting](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-output#keyword_spotting).")
	cmd.Flags().Int64VarP(&RecognizeMaxAlternatives, "max_alternatives", "", 0, "The maximum number of alternative transcripts that the service is to return. By default, the service returns a single transcript. If you specify a value of `0`, the service uses the default value, `1`. See [Maximum alternatives](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-output#max_alternatives).")
	cmd.Flags().Float32VarP(&RecognizeWordAlternativesThreshold, "word_alternatives_threshold", "", 0, "A confidence value that is the lower bound for identifying a hypothesis as a possible word alternative (also known as 'Confusion Networks'). An alternative word is considered if its confidence is greater than or equal to the threshold. Specify a probability between 0.0 and 1.0. By default, the service computes no alternative words. See [Word alternatives](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-output#word_alternatives).")
	cmd.Flags().BoolVarP(&RecognizeWordConfidence, "word_confidence", "", false, "If `true`, the service returns a confidence measure in the range of 0.0 to 1.0 for each word. By default, the service returns no word confidence scores. See [Word confidence](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-output#word_confidence).")
	cmd.Flags().BoolVarP(&RecognizeTimestamps, "timestamps", "", false, "If `true`, the service returns time alignment for each word. By default, no timestamps are returned. See [Word timestamps](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-output#word_timestamps).")
	cmd.Flags().BoolVarP(&RecognizeProfanityFilter, "profanity_filter", "", false, "If `true`, the service filters profanity from all output except for keyword results by replacing inappropriate words with a series of asterisks. Set the parameter to `false` to return results with no censoring. Applies to US English transcription only. See [Profanity filtering](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-output#profanity_filter).")
	cmd.Flags().BoolVarP(&RecognizeSmartFormatting, "smart_formatting", "", false, "If `true`, the service converts dates, times, series of digits and numbers, phone numbers, currency values, and internet addresses into more readable, conventional representations in the final transcript of a recognition request. For US English, the service also converts certain keyword strings to punctuation symbols. By default, the service performs no smart formatting. **Note:** Applies to US English, Japanese, and Spanish transcription only. See [Smart formatting](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-output#smart_formatting).")
	cmd.Flags().BoolVarP(&RecognizeSpeakerLabels, "speaker_labels", "", false, "If `true`, the response includes labels that identify which words were spoken by which participants in a multi-person exchange. By default, the service returns no speaker labels. Setting `speaker_labels` to `true` forces the `timestamps` parameter to be `true`, regardless of whether you specify `false` for the parameter. **Note:** Applies to US English, Japanese, and Spanish (both broadband and narrowband models) and UK English (narrowband model) transcription only. To determine whether a language model supports speaker labels, you can also use the **Get a model** method and check that the attribute `speaker_labels` is set to `true`. See [Speaker labels](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-output#speaker_labels).")
	cmd.Flags().StringVarP(&RecognizeCustomizationID, "customization_id", "", "", "**Deprecated.** Use the `language_customization_id` parameter to specify the customization ID (GUID) of a custom language model that is to be used with the recognition request. Do not specify both parameters with a request.")
	cmd.Flags().StringVarP(&RecognizeGrammarName, "grammar_name", "", "", "The name of a grammar that is to be used with the recognition request. If you specify a grammar, you must also use the `language_customization_id` parameter to specify the name of the custom language model for which the grammar is defined. The service recognizes only strings that are recognized by the specified grammar; it does not recognize other custom words from the model's words resource. See [Grammars](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-input#grammars-input).")
	cmd.Flags().BoolVarP(&RecognizeRedaction, "redaction", "", false, "If `true`, the service redacts, or masks, numeric data from final transcripts. The feature redacts any number that has three or more consecutive digits by replacing each digit with an `X` character. It is intended to redact sensitive numeric data, such as credit card numbers. By default, the service performs no redaction. When you enable redaction, the service automatically enables smart formatting, regardless of whether you explicitly disable that feature. To ensure maximum security, the service also disables keyword spotting (ignores the `keywords` and `keywords_threshold` parameters) and returns only a single final transcript (forces the `max_alternatives` parameter to be `1`). **Note:** Applies to US English, Japanese, and Korean transcription only. See [Numeric redaction](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-output#redaction).")
	cmd.Flags().BoolVarP(&RecognizeAudioMetrics, "audio_metrics", "", false, "If `true`, requests detailed information about the signal characteristics of the input audio. The service returns audio metrics with the final transcription results. By default, the service returns no audio metrics.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("audio")

	return cmd
}

func Recognize(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.RecognizeOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "audio" {
			audio, fileErr := os.Open(RecognizeAudio)
			utils.HandleError(fileErr)

			optionsModel.SetAudio(audio)
		}
		if flag.Name == "content_type" {
			optionsModel.SetContentType(RecognizeContentType)
		}
		if flag.Name == "model" {
			optionsModel.SetModel(RecognizeModel)
		}
		if flag.Name == "language_customization_id" {
			optionsModel.SetLanguageCustomizationID(RecognizeLanguageCustomizationID)
		}
		if flag.Name == "acoustic_customization_id" {
			optionsModel.SetAcousticCustomizationID(RecognizeAcousticCustomizationID)
		}
		if flag.Name == "base_model_version" {
			optionsModel.SetBaseModelVersion(RecognizeBaseModelVersion)
		}
		if flag.Name == "customization_weight" {
			optionsModel.SetCustomizationWeight(RecognizeCustomizationWeight)
		}
		if flag.Name == "inactivity_timeout" {
			optionsModel.SetInactivityTimeout(RecognizeInactivityTimeout)
		}
		if flag.Name == "keywords" {
			optionsModel.SetKeywords(RecognizeKeywords)
		}
		if flag.Name == "keywords_threshold" {
			optionsModel.SetKeywordsThreshold(RecognizeKeywordsThreshold)
		}
		if flag.Name == "max_alternatives" {
			optionsModel.SetMaxAlternatives(RecognizeMaxAlternatives)
		}
		if flag.Name == "word_alternatives_threshold" {
			optionsModel.SetWordAlternativesThreshold(RecognizeWordAlternativesThreshold)
		}
		if flag.Name == "word_confidence" {
			optionsModel.SetWordConfidence(RecognizeWordConfidence)
		}
		if flag.Name == "timestamps" {
			optionsModel.SetTimestamps(RecognizeTimestamps)
		}
		if flag.Name == "profanity_filter" {
			optionsModel.SetProfanityFilter(RecognizeProfanityFilter)
		}
		if flag.Name == "smart_formatting" {
			optionsModel.SetSmartFormatting(RecognizeSmartFormatting)
		}
		if flag.Name == "speaker_labels" {
			optionsModel.SetSpeakerLabels(RecognizeSpeakerLabels)
		}
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(RecognizeCustomizationID)
		}
		if flag.Name == "grammar_name" {
			optionsModel.SetGrammarName(RecognizeGrammarName)
		}
		if flag.Name == "redaction" {
			optionsModel.SetRedaction(RecognizeRedaction)
		}
		if flag.Name == "audio_metrics" {
			optionsModel.SetAudioMetrics(RecognizeAudioMetrics)
		}
	})

	result, _, responseErr := speechToText.Recognize(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var RegisterCallbackCallbackURL string
var RegisterCallbackUserSecret string

func getRegisterCallbackCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "register-callback",
		Short: "Register a callback",
		Long: "Registers a callback URL with the service for use with subsequent asynchronous recognition requests. The service attempts to register, or white-list, the callback URL if it is not already registered by sending a `GET` request to the callback URL. The service passes a random alphanumeric challenge string via the `challenge_string` parameter of the request. The request includes an `Accept` header that specifies `text/plain` as the required response type. To be registered successfully, the callback URL must respond to the `GET` request from the service. The response must send status code 200 and must include the challenge string in its body. Set the `Content-Type` response header to `text/plain`. Upon receiving this response, the service responds to the original registration request with response code 201. The service sends only a single `GET` request to the callback URL. If the service does not receive a reply with a response code of 200 and a body that echoes the challenge string sent by the service within five seconds, it does not white-list the URL; it instead sends status code 400 in response to the **Register a callback** request. If the requested callback URL is already white-listed, the service responds to the initial registration request with response code 200. If you specify a user secret with the request, the service uses it as a key to calculate an HMAC-SHA1 signature of the challenge string in its response to the `POST` request. It sends this signature in the `X-Callback-Signature` header of its `GET` request to the URL during registration. It also uses the secret to calculate a signature over the payload of every callback notification that uses the URL. The signature provides authentication and data integrity for HTTP communications. After you successfully register a callback URL, you can use it with an indefinite number of recognition requests. You can register a maximum of 20 callback URLS in a one-hour span of time. **See also:** [Registering a callback URL](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-async#register).",
		Run: RegisterCallback,
	}

	cmd.Flags().StringVarP(&RegisterCallbackCallbackURL, "callback_url", "", "", "An HTTP or HTTPS URL to which callback notifications are to be sent. To be white-listed, the URL must successfully echo the challenge string during URL verification. During verification, the client can also check the signature that the service sends in the `X-Callback-Signature` header to verify the origin of the request.")
	cmd.Flags().StringVarP(&RegisterCallbackUserSecret, "user_secret", "", "", "A user-specified string that the service uses to generate the HMAC-SHA1 signature that it sends via the `X-Callback-Signature` header. The service includes the header during URL verification and with every notification sent to the callback URL. It calculates the signature over the payload of the notification. If you omit the parameter, the service does not send the header.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("callback_url")

	return cmd
}

func RegisterCallback(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.RegisterCallbackOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "callback_url" {
			optionsModel.SetCallbackURL(RegisterCallbackCallbackURL)
		}
		if flag.Name == "user_secret" {
			optionsModel.SetUserSecret(RegisterCallbackUserSecret)
		}
	})

	result, _, responseErr := speechToText.RegisterCallback(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var UnregisterCallbackCallbackURL string

func getUnregisterCallbackCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "unregister-callback",
		Short: "Unregister a callback",
		Long: "Unregisters a callback URL that was previously white-listed with a **Register a callback** request for use with the asynchronous interface. Once unregistered, the URL can no longer be used with asynchronous recognition requests. **See also:** [Unregistering a callback URL](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-async#unregister).",
		Run: UnregisterCallback,
	}

	cmd.Flags().StringVarP(&UnregisterCallbackCallbackURL, "callback_url", "", "", "The callback URL that is to be unregistered.")

	cmd.MarkFlagRequired("callback_url")

	return cmd
}

func UnregisterCallback(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.UnregisterCallbackOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "callback_url" {
			optionsModel.SetCallbackURL(UnregisterCallbackCallbackURL)
		}
	})

	_, responseErr := speechToText.UnregisterCallback(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var CreateJobAudio string
var CreateJobContentType string
var CreateJobModel string
var CreateJobCallbackURL string
var CreateJobEvents string
var CreateJobUserToken string
var CreateJobResultsTTL int64
var CreateJobLanguageCustomizationID string
var CreateJobAcousticCustomizationID string
var CreateJobBaseModelVersion string
var CreateJobCustomizationWeight float64
var CreateJobInactivityTimeout int64
var CreateJobKeywords []string
var CreateJobKeywordsThreshold float32
var CreateJobMaxAlternatives int64
var CreateJobWordAlternativesThreshold float32
var CreateJobWordConfidence bool
var CreateJobTimestamps bool
var CreateJobProfanityFilter bool
var CreateJobSmartFormatting bool
var CreateJobSpeakerLabels bool
var CreateJobCustomizationID string
var CreateJobGrammarName string
var CreateJobRedaction bool
var CreateJobProcessingMetrics bool
var CreateJobProcessingMetricsInterval float32
var CreateJobAudioMetrics bool

func getCreateJobCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-job",
		Short: "Create a job",
		Long: "Creates a job for a new asynchronous recognition request. The job is owned by the instance of the service whose credentials are used to create it. How you learn the status and results of a job depends on the parameters you include with the job creation request:* By callback notification: Include the `callback_url` parameter to specify a URL to which the service is to send callback notifications when the status of the job changes. Optionally, you can also include the `events` and `user_token` parameters to subscribe to specific events and to specify a string that is to be included with each notification for the job.* By polling the service: Omit the `callback_url`, `events`, and `user_token` parameters. You must then use the **Check jobs** or **Check a job** methods to check the status of the job, using the latter to retrieve the results when the job is complete. The two approaches are not mutually exclusive. You can poll the service for job status or obtain results from the service manually even if you include a callback URL. In both cases, you can include the `results_ttl` parameter to specify how long the results are to remain available after the job is complete. Using the HTTPS **Check a job** method to retrieve results is more secure than receiving them via callback notification over HTTP because it provides confidentiality in addition to authentication and data integrity. The method supports the same basic parameters as other HTTP and WebSocket recognition requests. It also supports the following parameters specific to the asynchronous interface:* `callback_url`* `events`* `user_token`* `results_ttl` You can pass a maximum of 1 GB and a minimum of 100 bytes of audio with a request. The service automatically detects the endianness of the incoming audio and, for audio that includes multiple channels, downmixes the audio to one-channel mono during transcoding. The method returns only final results; to enable interim results, use the WebSocket API. (With the `curl` command, use the `--data-binary` option to upload the file for the request.) **See also:** [Creating a job](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-async#create). ### Streaming mode For requests to transcribe live audio as it becomes available, you must set the `Transfer-Encoding` header to `chunked` to use streaming mode. In streaming mode, the service closes the connection (status code 408) if it does not receive at least 15 seconds of audio (including silence) in any 30-second period. The service also closes the connection (status code 400) if it detects no speech for `inactivity_timeout` seconds of streaming audio; use the `inactivity_timeout` parameter to change the default of 30 seconds. **See also:*** [Audio transmission](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-input#transmission)* [Timeouts](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-input#timeouts) ### Audio formats (content types) The service accepts audio in the following formats (MIME types).* For formats that are labeled **Required**, you must use the `Content-Type` header with the request to specify the format of the audio.* For all other formats, you can omit the `Content-Type` header or specify `application/octet-stream` with the header to have the service automatically detect the format of the audio. (With the `curl` command, you can specify either `'Content-Type:'` or `'Content-Type: application/octet-stream'`.) Where indicated, the format that you specify must include the sampling rate and can optionally include the number of channels and the endianness of the audio.* `audio/alaw` (**Required.** Specify the sampling rate (`rate`) of the audio.)* `audio/basic` (**Required.** Use only with narrowband models.)* `audio/flac`* `audio/g729` (Use only with narrowband models.)* `audio/l16` (**Required.** Specify the sampling rate (`rate`) and optionally the number of channels (`channels`) and endianness (`endianness`) of the audio.)* `audio/mp3`* `audio/mpeg`* `audio/mulaw` (**Required.** Specify the sampling rate (`rate`) of the audio.)* `audio/ogg` (The service automatically detects the codec of the input audio.)* `audio/ogg;codecs=opus`* `audio/ogg;codecs=vorbis`* `audio/wav` (Provide audio with a maximum of nine channels.)* `audio/webm` (The service automatically detects the codec of the input audio.)* `audio/webm;codecs=opus`* `audio/webm;codecs=vorbis` The sampling rate of the audio must match the sampling rate of the model for the recognition request: for broadband models, at least 16 kHz; for narrowband models, at least 8 kHz. If the sampling rate of the audio is higher than the minimum required rate, the service down-samples the audio to the appropriate rate. If the sampling rate of the audio is lower than the minimum required rate, the request fails. **See also:** [Audio formats](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-audio-formats#audio-formats).",
		Run: CreateJob,
	}

	cmd.Flags().StringVarP(&CreateJobAudio, "audio", "", "", "The audio to transcribe.")
	cmd.Flags().StringVarP(&CreateJobContentType, "content_type", "", "", "The format (MIME type) of the audio. For more information about specifying an audio format, see **Audio formats (content types)** in the method description.")
	cmd.Flags().StringVarP(&CreateJobModel, "model", "", "", "The identifier of the model that is to be used for the recognition request. See [Languages and models](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-models#models).")
	cmd.Flags().StringVarP(&CreateJobCallbackURL, "callback_url", "", "", "A URL to which callback notifications are to be sent. The URL must already be successfully white-listed by using the **Register a callback** method. You can include the same callback URL with any number of job creation requests. Omit the parameter to poll the service for job completion and results. Use the `user_token` parameter to specify a unique user-specified string with each job to differentiate the callback notifications for the jobs.")
	cmd.Flags().StringVarP(&CreateJobEvents, "events", "", "", "If the job includes a callback URL, a comma-separated list of notification events to which to subscribe. Valid events are* `recognitions.started` generates a callback notification when the service begins to process the job.* `recognitions.completed` generates a callback notification when the job is complete. You must use the **Check a job** method to retrieve the results before they time out or are deleted.* `recognitions.completed_with_results` generates a callback notification when the job is complete. The notification includes the results of the request.* `recognitions.failed` generates a callback notification if the service experiences an error while processing the job. The `recognitions.completed` and `recognitions.completed_with_results` events are incompatible. You can specify only of the two events. If the job includes a callback URL, omit the parameter to subscribe to the default events: `recognitions.started`, `recognitions.completed`, and `recognitions.failed`. If the job does not include a callback URL, omit the parameter.")
	cmd.Flags().StringVarP(&CreateJobUserToken, "user_token", "", "", "If the job includes a callback URL, a user-specified string that the service is to include with each callback notification for the job; the token allows the user to maintain an internal mapping between jobs and notification events. If the job does not include a callback URL, omit the parameter.")
	cmd.Flags().Int64VarP(&CreateJobResultsTTL, "results_ttl", "", 0, "The number of minutes for which the results are to be available after the job has finished. If not delivered via a callback, the results must be retrieved within this time. Omit the parameter to use a time to live of one week. The parameter is valid with or without a callback URL.")
	cmd.Flags().StringVarP(&CreateJobLanguageCustomizationID, "language_customization_id", "", "", "The customization ID (GUID) of a custom language model that is to be used with the recognition request. The base model of the specified custom language model must match the model specified with the `model` parameter. You must make the request with credentials for the instance of the service that owns the custom model. By default, no custom language model is used. See [Custom models](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-input#custom-input). **Note:** Use this parameter instead of the deprecated `customization_id` parameter.")
	cmd.Flags().StringVarP(&CreateJobAcousticCustomizationID, "acoustic_customization_id", "", "", "The customization ID (GUID) of a custom acoustic model that is to be used with the recognition request. The base model of the specified custom acoustic model must match the model specified with the `model` parameter. You must make the request with credentials for the instance of the service that owns the custom model. By default, no custom acoustic model is used. See [Custom models](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-input#custom-input).")
	cmd.Flags().StringVarP(&CreateJobBaseModelVersion, "base_model_version", "", "", "The version of the specified base model that is to be used with the recognition request. Multiple versions of a base model can exist when a model is updated for internal improvements. The parameter is intended primarily for use with custom models that have been upgraded for a new base model. The default value depends on whether the parameter is used with or without a custom model. See [Base model version](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-input#version).")
	cmd.Flags().Float64VarP(&CreateJobCustomizationWeight, "customization_weight", "", 0, "If you specify the customization ID (GUID) of a custom language model with the recognition request, the customization weight tells the service how much weight to give to words from the custom language model compared to those from the base model for the current request. Specify a value between 0.0 and 1.0. Unless a different customization weight was specified for the custom model when it was trained, the default value is 0.3. A customization weight that you specify overrides a weight that was specified when the custom model was trained. The default value yields the best performance in general. Assign a higher value if your audio makes frequent use of OOV words from the custom model. Use caution when setting the weight: a higher value can improve the accuracy of phrases from the custom model's domain, but it can negatively affect performance on non-domain phrases. See [Custom models](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-input#custom-input).")
	cmd.Flags().Int64VarP(&CreateJobInactivityTimeout, "inactivity_timeout", "", 0, "The time in seconds after which, if only silence (no speech) is detected in streaming audio, the connection is closed with a 400 error. The parameter is useful for stopping audio submission from a live microphone when a user simply walks away. Use `-1` for infinity. See [Inactivity timeout](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-input#timeouts-inactivity).")
	cmd.Flags().StringSliceVarP(&CreateJobKeywords, "keywords", "", nil, "An array of keyword strings to spot in the audio. Each keyword string can include one or more string tokens. Keywords are spotted only in the final results, not in interim hypotheses. If you specify any keywords, you must also specify a keywords threshold. You can spot a maximum of 1000 keywords. Omit the parameter or specify an empty array if you do not need to spot keywords. See [Keyword spotting](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-output#keyword_spotting).")
	cmd.Flags().Float32VarP(&CreateJobKeywordsThreshold, "keywords_threshold", "", 0, "A confidence value that is the lower bound for spotting a keyword. A word is considered to match a keyword if its confidence is greater than or equal to the threshold. Specify a probability between 0.0 and 1.0. If you specify a threshold, you must also specify one or more keywords. The service performs no keyword spotting if you omit either parameter. See [Keyword spotting](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-output#keyword_spotting).")
	cmd.Flags().Int64VarP(&CreateJobMaxAlternatives, "max_alternatives", "", 0, "The maximum number of alternative transcripts that the service is to return. By default, the service returns a single transcript. If you specify a value of `0`, the service uses the default value, `1`. See [Maximum alternatives](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-output#max_alternatives).")
	cmd.Flags().Float32VarP(&CreateJobWordAlternativesThreshold, "word_alternatives_threshold", "", 0, "A confidence value that is the lower bound for identifying a hypothesis as a possible word alternative (also known as 'Confusion Networks'). An alternative word is considered if its confidence is greater than or equal to the threshold. Specify a probability between 0.0 and 1.0. By default, the service computes no alternative words. See [Word alternatives](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-output#word_alternatives).")
	cmd.Flags().BoolVarP(&CreateJobWordConfidence, "word_confidence", "", false, "If `true`, the service returns a confidence measure in the range of 0.0 to 1.0 for each word. By default, the service returns no word confidence scores. See [Word confidence](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-output#word_confidence).")
	cmd.Flags().BoolVarP(&CreateJobTimestamps, "timestamps", "", false, "If `true`, the service returns time alignment for each word. By default, no timestamps are returned. See [Word timestamps](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-output#word_timestamps).")
	cmd.Flags().BoolVarP(&CreateJobProfanityFilter, "profanity_filter", "", false, "If `true`, the service filters profanity from all output except for keyword results by replacing inappropriate words with a series of asterisks. Set the parameter to `false` to return results with no censoring. Applies to US English transcription only. See [Profanity filtering](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-output#profanity_filter).")
	cmd.Flags().BoolVarP(&CreateJobSmartFormatting, "smart_formatting", "", false, "If `true`, the service converts dates, times, series of digits and numbers, phone numbers, currency values, and internet addresses into more readable, conventional representations in the final transcript of a recognition request. For US English, the service also converts certain keyword strings to punctuation symbols. By default, the service performs no smart formatting. **Note:** Applies to US English, Japanese, and Spanish transcription only. See [Smart formatting](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-output#smart_formatting).")
	cmd.Flags().BoolVarP(&CreateJobSpeakerLabels, "speaker_labels", "", false, "If `true`, the response includes labels that identify which words were spoken by which participants in a multi-person exchange. By default, the service returns no speaker labels. Setting `speaker_labels` to `true` forces the `timestamps` parameter to be `true`, regardless of whether you specify `false` for the parameter. **Note:** Applies to US English, Japanese, and Spanish (both broadband and narrowband models) and UK English (narrowband model) transcription only. To determine whether a language model supports speaker labels, you can also use the **Get a model** method and check that the attribute `speaker_labels` is set to `true`. See [Speaker labels](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-output#speaker_labels).")
	cmd.Flags().StringVarP(&CreateJobCustomizationID, "customization_id", "", "", "**Deprecated.** Use the `language_customization_id` parameter to specify the customization ID (GUID) of a custom language model that is to be used with the recognition request. Do not specify both parameters with a request.")
	cmd.Flags().StringVarP(&CreateJobGrammarName, "grammar_name", "", "", "The name of a grammar that is to be used with the recognition request. If you specify a grammar, you must also use the `language_customization_id` parameter to specify the name of the custom language model for which the grammar is defined. The service recognizes only strings that are recognized by the specified grammar; it does not recognize other custom words from the model's words resource. See [Grammars](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-input#grammars-input).")
	cmd.Flags().BoolVarP(&CreateJobRedaction, "redaction", "", false, "If `true`, the service redacts, or masks, numeric data from final transcripts. The feature redacts any number that has three or more consecutive digits by replacing each digit with an `X` character. It is intended to redact sensitive numeric data, such as credit card numbers. By default, the service performs no redaction. When you enable redaction, the service automatically enables smart formatting, regardless of whether you explicitly disable that feature. To ensure maximum security, the service also disables keyword spotting (ignores the `keywords` and `keywords_threshold` parameters) and returns only a single final transcript (forces the `max_alternatives` parameter to be `1`). **Note:** Applies to US English, Japanese, and Korean transcription only. See [Numeric redaction](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-output#redaction).")
	cmd.Flags().BoolVarP(&CreateJobProcessingMetrics, "processing_metrics", "", false, "If `true`, requests processing metrics about the service's transcription of the input audio. The service returns processing metrics at the interval specified by the `processing_metrics_interval` parameter. It also returns processing metrics for transcription events, for example, for final and interim results. By default, the service returns no processing metrics.")
	cmd.Flags().Float32VarP(&CreateJobProcessingMetricsInterval, "processing_metrics_interval", "", 0, "Specifies the interval in real wall-clock seconds at which the service is to return processing metrics. The parameter is ignored unless the `processing_metrics` parameter is set to `true`. The parameter accepts a minimum value of 0.1 seconds. The level of precision is not restricted, so you can specify values such as 0.25 and 0.125. The service does not impose a maximum value. If you want to receive processing metrics only for transcription events instead of at periodic intervals, set the value to a large number. If the value is larger than the duration of the audio, the service returns processing metrics only for transcription events.")
	cmd.Flags().BoolVarP(&CreateJobAudioMetrics, "audio_metrics", "", false, "If `true`, requests detailed information about the signal characteristics of the input audio. The service returns audio metrics with the final transcription results. By default, the service returns no audio metrics.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("audio")

	return cmd
}

func CreateJob(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.CreateJobOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "audio" {
			audio, fileErr := os.Open(CreateJobAudio)
			utils.HandleError(fileErr)

			optionsModel.SetAudio(audio)
		}
		if flag.Name == "content_type" {
			optionsModel.SetContentType(CreateJobContentType)
		}
		if flag.Name == "model" {
			optionsModel.SetModel(CreateJobModel)
		}
		if flag.Name == "callback_url" {
			optionsModel.SetCallbackURL(CreateJobCallbackURL)
		}
		if flag.Name == "events" {
			optionsModel.SetEvents(CreateJobEvents)
		}
		if flag.Name == "user_token" {
			optionsModel.SetUserToken(CreateJobUserToken)
		}
		if flag.Name == "results_ttl" {
			optionsModel.SetResultsTTL(CreateJobResultsTTL)
		}
		if flag.Name == "language_customization_id" {
			optionsModel.SetLanguageCustomizationID(CreateJobLanguageCustomizationID)
		}
		if flag.Name == "acoustic_customization_id" {
			optionsModel.SetAcousticCustomizationID(CreateJobAcousticCustomizationID)
		}
		if flag.Name == "base_model_version" {
			optionsModel.SetBaseModelVersion(CreateJobBaseModelVersion)
		}
		if flag.Name == "customization_weight" {
			optionsModel.SetCustomizationWeight(CreateJobCustomizationWeight)
		}
		if flag.Name == "inactivity_timeout" {
			optionsModel.SetInactivityTimeout(CreateJobInactivityTimeout)
		}
		if flag.Name == "keywords" {
			optionsModel.SetKeywords(CreateJobKeywords)
		}
		if flag.Name == "keywords_threshold" {
			optionsModel.SetKeywordsThreshold(CreateJobKeywordsThreshold)
		}
		if flag.Name == "max_alternatives" {
			optionsModel.SetMaxAlternatives(CreateJobMaxAlternatives)
		}
		if flag.Name == "word_alternatives_threshold" {
			optionsModel.SetWordAlternativesThreshold(CreateJobWordAlternativesThreshold)
		}
		if flag.Name == "word_confidence" {
			optionsModel.SetWordConfidence(CreateJobWordConfidence)
		}
		if flag.Name == "timestamps" {
			optionsModel.SetTimestamps(CreateJobTimestamps)
		}
		if flag.Name == "profanity_filter" {
			optionsModel.SetProfanityFilter(CreateJobProfanityFilter)
		}
		if flag.Name == "smart_formatting" {
			optionsModel.SetSmartFormatting(CreateJobSmartFormatting)
		}
		if flag.Name == "speaker_labels" {
			optionsModel.SetSpeakerLabels(CreateJobSpeakerLabels)
		}
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(CreateJobCustomizationID)
		}
		if flag.Name == "grammar_name" {
			optionsModel.SetGrammarName(CreateJobGrammarName)
		}
		if flag.Name == "redaction" {
			optionsModel.SetRedaction(CreateJobRedaction)
		}
		if flag.Name == "processing_metrics" {
			optionsModel.SetProcessingMetrics(CreateJobProcessingMetrics)
		}
		if flag.Name == "processing_metrics_interval" {
			optionsModel.SetProcessingMetricsInterval(CreateJobProcessingMetricsInterval)
		}
		if flag.Name == "audio_metrics" {
			optionsModel.SetAudioMetrics(CreateJobAudioMetrics)
		}
	})

	result, _, responseErr := speechToText.CreateJob(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}


func getCheckJobsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "check-jobs",
		Short: "Check jobs",
		Long: "Returns the ID and status of the latest 100 outstanding jobs associated with the credentials with which it is called. The method also returns the creation and update times of each job, and, if a job was created with a callback URL and a user token, the user token for the job. To obtain the results for a job whose status is `completed` or not one of the latest 100 outstanding jobs, use the **Check a job** method. A job and its results remain available until you delete them with the **Delete a job** method or until the job's time to live expires, whichever comes first. **See also:** [Checking the status of the latest jobs](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-async#jobs).",
		Run: CheckJobs,
	}

	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")


	return cmd
}

func CheckJobs(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.CheckJobsOptions{}

	result, _, responseErr := speechToText.CheckJobs(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var CheckJobID string

func getCheckJobCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "check-job",
		Short: "Check a job",
		Long: "Returns information about the specified job. The response always includes the status of the job and its creation and update times. If the status is `completed`, the response includes the results of the recognition request. You must use credentials for the instance of the service that owns a job to list information about it. You can use the method to retrieve the results of any job, regardless of whether it was submitted with a callback URL and the `recognitions.completed_with_results` event, and you can retrieve the results multiple times for as long as they remain available. Use the **Check jobs** method to request information about the most recent jobs associated with the calling credentials. **See also:** [Checking the status and retrieving the results of a job](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-async#job).",
		Run: CheckJob,
	}

	cmd.Flags().StringVarP(&CheckJobID, "id", "", "", "The identifier of the asynchronous job that is to be used for the request. You must make the request with credentials for the instance of the service that owns the job.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("id")

	return cmd
}

func CheckJob(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.CheckJobOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "id" {
			optionsModel.SetID(CheckJobID)
		}
	})

	result, _, responseErr := speechToText.CheckJob(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteJobID string

func getDeleteJobCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-job",
		Short: "Delete a job",
		Long: "Deletes the specified job. You cannot delete a job that the service is actively processing. Once you delete a job, its results are no longer available. The service automatically deletes a job and its results when the time to live for the results expires. You must use credentials for the instance of the service that owns a job to delete it. **See also:** [Deleting a job](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-async#delete-async).",
		Run: DeleteJob,
	}

	cmd.Flags().StringVarP(&DeleteJobID, "id", "", "", "The identifier of the asynchronous job that is to be used for the request. You must make the request with credentials for the instance of the service that owns the job.")

	cmd.MarkFlagRequired("id")

	return cmd
}

func DeleteJob(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.DeleteJobOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "id" {
			optionsModel.SetID(DeleteJobID)
		}
	})

	_, responseErr := speechToText.DeleteJob(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var CreateLanguageModelName string
var CreateLanguageModelBaseModelName string
var CreateLanguageModelDialect string
var CreateLanguageModelDescription string

func getCreateLanguageModelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-language-model",
		Short: "Create a custom language model",
		Long: "Creates a new custom language model for a specified base model. The custom language model can be used only with the base model for which it is created. The model is owned by the instance of the service whose credentials are used to create it. **See also:** [Create a custom language model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-languageCreate#createModel-language).",
		Run: CreateLanguageModel,
	}

	cmd.Flags().StringVarP(&CreateLanguageModelName, "name", "", "", "A user-defined name for the new custom language model. Use a name that is unique among all custom language models that you own. Use a localized name that matches the language of the custom model. Use a name that describes the domain of the custom model, such as `Medical custom model` or `Legal custom model`.")
	cmd.Flags().StringVarP(&CreateLanguageModelBaseModelName, "base_model_name", "", "", "The name of the base language model that is to be customized by the new custom language model. The new custom model can be used only with the base model that it customizes. To determine whether a base model supports language model customization, use the **Get a model** method and check that the attribute `custom_language_model` is set to `true`. You can also refer to [Language support for customization](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-customization#languageSupport).")
	cmd.Flags().StringVarP(&CreateLanguageModelDialect, "dialect", "", "", "The dialect of the specified language that is to be used with the custom language model. For most languages, the dialect matches the language of the base model by default. For example, `en-US` is used for either of the US English language models. For a Spanish language, the service creates a custom language model that is suited for speech in one of the following dialects:* `es-ES` for Castilian Spanish (`es-ES` models)* `es-LA` for Latin American Spanish (`es-AR`, `es-CL`, `es-CO`, and `es-PE` models)* `es-US` for Mexican (North American) Spanish (`es-MX` models) The parameter is meaningful only for Spanish models, for which you can always safely omit the parameter to have the service create the correct mapping. If you specify the `dialect` parameter for non-Spanish language models, its value must match the language of the base model. If you specify the `dialect` for Spanish language models, its value must match one of the defined mappings as indicated (`es-ES`, `es-LA`, or `es-MX`). All dialect values are case-insensitive.")
	cmd.Flags().StringVarP(&CreateLanguageModelDescription, "description", "", "", "A description of the new custom language model. Use a localized description that matches the language of the custom model.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("base_model_name")

	return cmd
}

func CreateLanguageModel(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.CreateLanguageModelOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "name" {
			optionsModel.SetName(CreateLanguageModelName)
		}
		if flag.Name == "base_model_name" {
			optionsModel.SetBaseModelName(CreateLanguageModelBaseModelName)
		}
		if flag.Name == "dialect" {
			optionsModel.SetDialect(CreateLanguageModelDialect)
		}
		if flag.Name == "description" {
			optionsModel.SetDescription(CreateLanguageModelDescription)
		}
	})

	result, _, responseErr := speechToText.CreateLanguageModel(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ListLanguageModelsLanguage string

func getListLanguageModelsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-language-models",
		Short: "List custom language models",
		Long: "Lists information about all custom language models that are owned by an instance of the service. Use the `language` parameter to see all custom language models for the specified language. Omit the parameter to see all custom language models for all languages. You must use credentials for the instance of the service that owns a model to list information about it. **See also:** [Listing custom language models](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-manageLanguageModels#listModels-language).",
		Run: ListLanguageModels,
	}

	cmd.Flags().StringVarP(&ListLanguageModelsLanguage, "language", "", "", "The identifier of the language for which custom language or custom acoustic models are to be returned (for example, `en-US`). Omit the parameter to see all custom language or custom acoustic models that are owned by the requesting credentials.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")


	return cmd
}

func ListLanguageModels(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.ListLanguageModelsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "language" {
			optionsModel.SetLanguage(ListLanguageModelsLanguage)
		}
	})

	result, _, responseErr := speechToText.ListLanguageModels(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetLanguageModelCustomizationID string

func getGetLanguageModelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-language-model",
		Short: "Get a custom language model",
		Long: "Gets information about a specified custom language model. You must use credentials for the instance of the service that owns a model to list information about it. **See also:** [Listing custom language models](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-manageLanguageModels#listModels-language).",
		Run: GetLanguageModel,
	}

	cmd.Flags().StringVarP(&GetLanguageModelCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom language model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("customization_id")

	return cmd
}

func GetLanguageModel(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.GetLanguageModelOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(GetLanguageModelCustomizationID)
		}
	})

	result, _, responseErr := speechToText.GetLanguageModel(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteLanguageModelCustomizationID string

func getDeleteLanguageModelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-language-model",
		Short: "Delete a custom language model",
		Long: "Deletes an existing custom language model. The custom model cannot be deleted if another request, such as adding a corpus or grammar to the model, is currently being processed. You must use credentials for the instance of the service that owns a model to delete it. **See also:** [Deleting a custom language model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-manageLanguageModels#deleteModel-language).",
		Run: DeleteLanguageModel,
	}

	cmd.Flags().StringVarP(&DeleteLanguageModelCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom language model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")

	cmd.MarkFlagRequired("customization_id")

	return cmd
}

func DeleteLanguageModel(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.DeleteLanguageModelOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(DeleteLanguageModelCustomizationID)
		}
	})

	_, responseErr := speechToText.DeleteLanguageModel(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var TrainLanguageModelCustomizationID string
var TrainLanguageModelWordTypeToAdd string
var TrainLanguageModelCustomizationWeight float64

func getTrainLanguageModelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "train-language-model",
		Short: "Train a custom language model",
		Long: "Initiates the training of a custom language model with new resources such as corpora, grammars, and custom words. After adding, modifying, or deleting resources for a custom language model, use this method to begin the actual training of the model on the latest data. You can specify whether the custom language model is to be trained with all words from its words resource or only with words that were added or modified by the user directly. You must use credentials for the instance of the service that owns a model to train it. The training method is asynchronous. It can take on the order of minutes to complete depending on the amount of data on which the service is being trained and the current load on the service. The method returns an HTTP 200 response code to indicate that the training process has begun. You can monitor the status of the training by using the **Get a custom language model** method to poll the model's status. Use a loop to check the status every 10 seconds. The method returns a `LanguageModel` object that includes `status` and `progress` fields. A status of `available` means that the custom model is trained and ready to use. The service cannot accept subsequent training requests or requests to add new resources until the existing request completes. **See also:** [Train the custom language model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-languageCreate#trainModel-language). ### Training failures Training can fail to start for the following reasons:* The service is currently handling another request for the custom model, such as another training request or a request to add a corpus or grammar to the model.* No training data have been added to the custom model.* The custom model contains one or more invalid corpora, grammars, or words (for example, a custom word has an invalid sounds-like pronunciation). You can correct the invalid resources or set the `strict` parameter to `false` to exclude the invalid resources from the training. The model must contain at least one valid resource for training to succeed.",
		Run: TrainLanguageModel,
	}

	cmd.Flags().StringVarP(&TrainLanguageModelCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom language model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&TrainLanguageModelWordTypeToAdd, "word_type_to_add", "", "", "The type of words from the custom language model's words resource on which to train the model:* `all` (the default) trains the model on all new words, regardless of whether they were extracted from corpora or grammars or were added or modified by the user.* `user` trains the model only on new words that were added or modified by the user directly. The model is not trained on new words extracted from corpora or grammars.")
	cmd.Flags().Float64VarP(&TrainLanguageModelCustomizationWeight, "customization_weight", "", 0, "Specifies a customization weight for the custom language model. The customization weight tells the service how much weight to give to words from the custom language model compared to those from the base model for speech recognition. Specify a value between 0.0 and 1.0; the default is 0.3. The default value yields the best performance in general. Assign a higher value if your audio makes frequent use of OOV words from the custom model. Use caution when setting the weight: a higher value can improve the accuracy of phrases from the custom model's domain, but it can negatively affect performance on non-domain phrases. The value that you assign is used for all recognition requests that use the model. You can override it for any recognition request by specifying a customization weight for that request.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("customization_id")

	return cmd
}

func TrainLanguageModel(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.TrainLanguageModelOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(TrainLanguageModelCustomizationID)
		}
		if flag.Name == "word_type_to_add" {
			optionsModel.SetWordTypeToAdd(TrainLanguageModelWordTypeToAdd)
		}
		if flag.Name == "customization_weight" {
			optionsModel.SetCustomizationWeight(TrainLanguageModelCustomizationWeight)
		}
	})

	result, _, responseErr := speechToText.TrainLanguageModel(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ResetLanguageModelCustomizationID string

func getResetLanguageModelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "reset-language-model",
		Short: "Reset a custom language model",
		Long: "Resets a custom language model by removing all corpora, grammars, and words from the model. Resetting a custom language model initializes the model to its state when it was first created. Metadata such as the name and language of the model are preserved, but the model's words resource is removed and must be re-created. You must use credentials for the instance of the service that owns a model to reset it. **See also:** [Resetting a custom language model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-manageLanguageModels#resetModel-language).",
		Run: ResetLanguageModel,
	}

	cmd.Flags().StringVarP(&ResetLanguageModelCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom language model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")

	cmd.MarkFlagRequired("customization_id")

	return cmd
}

func ResetLanguageModel(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.ResetLanguageModelOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(ResetLanguageModelCustomizationID)
		}
	})

	_, responseErr := speechToText.ResetLanguageModel(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var UpgradeLanguageModelCustomizationID string

func getUpgradeLanguageModelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "upgrade-language-model",
		Short: "Upgrade a custom language model",
		Long: "Initiates the upgrade of a custom language model to the latest version of its base language model. The upgrade method is asynchronous. It can take on the order of minutes to complete depending on the amount of data in the custom model and the current load on the service. A custom model must be in the `ready` or `available` state to be upgraded. You must use credentials for the instance of the service that owns a model to upgrade it. The method returns an HTTP 200 response code to indicate that the upgrade process has begun successfully. You can monitor the status of the upgrade by using the **Get a custom language model** method to poll the model's status. The method returns a `LanguageModel` object that includes `status` and `progress` fields. Use a loop to check the status every 10 seconds. While it is being upgraded, the custom model has the status `upgrading`. When the upgrade is complete, the model resumes the status that it had prior to upgrade. The service cannot accept subsequent requests for the model until the upgrade completes. **See also:** [Upgrading a custom language model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-customUpgrade#upgradeLanguage).",
		Run: UpgradeLanguageModel,
	}

	cmd.Flags().StringVarP(&UpgradeLanguageModelCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom language model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")

	cmd.MarkFlagRequired("customization_id")

	return cmd
}

func UpgradeLanguageModel(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.UpgradeLanguageModelOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(UpgradeLanguageModelCustomizationID)
		}
	})

	_, responseErr := speechToText.UpgradeLanguageModel(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var ListCorporaCustomizationID string

func getListCorporaCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-corpora",
		Short: "List corpora",
		Long: "Lists information about all corpora from a custom language model. The information includes the total number of words and out-of-vocabulary (OOV) words, name, and status of each corpus. You must use credentials for the instance of the service that owns a model to list its corpora. **See also:** [Listing corpora for a custom language model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-manageCorpora#listCorpora).",
		Run: ListCorpora,
	}

	cmd.Flags().StringVarP(&ListCorporaCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom language model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("customization_id")

	return cmd
}

func ListCorpora(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.ListCorporaOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(ListCorporaCustomizationID)
		}
	})

	result, _, responseErr := speechToText.ListCorpora(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var AddCorpusCustomizationID string
var AddCorpusCorpusName string
var AddCorpusCorpusFile string
var AddCorpusAllowOverwrite bool

func getAddCorpusCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "add-corpus",
		Short: "Add a corpus",
		Long: "Adds a single corpus text file of new training data to a custom language model. Use multiple requests to submit multiple corpus text files. You must use credentials for the instance of the service that owns a model to add a corpus to it. Adding a corpus does not affect the custom language model until you train the model for the new data by using the **Train a custom language model** method. Submit a plain text file that contains sample sentences from the domain of interest to enable the service to extract words in context. The more sentences you add that represent the context in which speakers use words from the domain, the better the service's recognition accuracy. The call returns an HTTP 201 response code if the corpus is valid. The service then asynchronously processes the contents of the corpus and automatically extracts new words that it finds. This can take on the order of a minute or two to complete depending on the total number of words and the number of new words in the corpus, as well as the current load on the service. You cannot submit requests to add additional resources to the custom model or to train the model until the service's analysis of the corpus for the current request completes. Use the **List a corpus** method to check the status of the analysis. The service auto-populates the model's words resource with words from the corpus that are not found in its base vocabulary. These are referred to as out-of-vocabulary (OOV) words. You can use the **List custom words** method to examine the words resource. You can use other words method to eliminate typos and modify how words are pronounced as needed. To add a corpus file that has the same name as an existing corpus, set the `allow_overwrite` parameter to `true`; otherwise, the request fails. Overwriting an existing corpus causes the service to process the corpus text file and extract OOV words anew. Before doing so, it removes any OOV words associated with the existing corpus from the model's words resource unless they were also added by another corpus or grammar, or they have been modified in some way with the **Add custom words** or **Add a custom word** method. The service limits the overall amount of data that you can add to a custom model to a maximum of 10 million total words from all sources combined. Also, you can add no more than 90 thousand custom (OOV) words to a model. This includes words that the service extracts from corpora and grammars, and words that you add directly. **See also:*** [Working with corpora](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-corporaWords#workingCorpora)* [Add a corpus to the custom language model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-languageCreate#addCorpus).",
		Run: AddCorpus,
	}

	cmd.Flags().StringVarP(&AddCorpusCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom language model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&AddCorpusCorpusName, "corpus_name", "", "", "The name of the new corpus for the custom language model. Use a localized name that matches the language of the custom model and reflects the contents of the corpus.* Include a maximum of 128 characters in the name.* Do not use characters that need to be URL-encoded. For example, do not use spaces, slashes, backslashes, colons, ampersands, double quotes, plus signs, equals signs, questions marks, and so on in the name. (The service does not prevent the use of these characters. But because they must be URL-encoded wherever used, their use is strongly discouraged.)* Do not use the name of an existing corpus or grammar that is already defined for the custom model.* Do not use the name `user`, which is reserved by the service to denote custom words that are added or modified by the user.* Do not use the name `base_lm` or `default_lm`. Both names are reserved for future use by the service.")
	cmd.Flags().StringVarP(&AddCorpusCorpusFile, "corpus_file", "", "", "A plain text file that contains the training data for the corpus. Encode the file in UTF-8 if it contains non-ASCII characters; the service assumes UTF-8 encoding if it encounters non-ASCII characters. Make sure that you know the character encoding of the file. You must use that encoding when working with the words in the custom language model. For more information, see [Character encoding](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-corporaWords#charEncoding). With the `curl` command, use the `--data-binary` option to upload the file for the request.")
	cmd.Flags().BoolVarP(&AddCorpusAllowOverwrite, "allow_overwrite", "", false, "If `true`, the specified corpus overwrites an existing corpus with the same name. If `false`, the request fails if a corpus with the same name already exists. The parameter has no effect if a corpus with the same name does not already exist.")

	cmd.MarkFlagRequired("customization_id")
	cmd.MarkFlagRequired("corpus_name")
	cmd.MarkFlagRequired("corpus_file")

	return cmd
}

func AddCorpus(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.AddCorpusOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(AddCorpusCustomizationID)
		}
		if flag.Name == "corpus_name" {
			optionsModel.SetCorpusName(AddCorpusCorpusName)
		}
		if flag.Name == "corpus_file" {
			corpus_file, fileErr := os.Open(AddCorpusCorpusFile)
			utils.HandleError(fileErr)

			optionsModel.SetCorpusFile(corpus_file)
		}
		if flag.Name == "allow_overwrite" {
			optionsModel.SetAllowOverwrite(AddCorpusAllowOverwrite)
		}
	})

	_, responseErr := speechToText.AddCorpus(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var GetCorpusCustomizationID string
var GetCorpusCorpusName string

func getGetCorpusCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-corpus",
		Short: "Get a corpus",
		Long: "Gets information about a corpus from a custom language model. The information includes the total number of words and out-of-vocabulary (OOV) words, name, and status of the corpus. You must use credentials for the instance of the service that owns a model to list its corpora. **See also:** [Listing corpora for a custom language model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-manageCorpora#listCorpora).",
		Run: GetCorpus,
	}

	cmd.Flags().StringVarP(&GetCorpusCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom language model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&GetCorpusCorpusName, "corpus_name", "", "", "The name of the corpus for the custom language model.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("customization_id")
	cmd.MarkFlagRequired("corpus_name")

	return cmd
}

func GetCorpus(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.GetCorpusOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(GetCorpusCustomizationID)
		}
		if flag.Name == "corpus_name" {
			optionsModel.SetCorpusName(GetCorpusCorpusName)
		}
	})

	result, _, responseErr := speechToText.GetCorpus(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteCorpusCustomizationID string
var DeleteCorpusCorpusName string

func getDeleteCorpusCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-corpus",
		Short: "Delete a corpus",
		Long: "Deletes an existing corpus from a custom language model. The service removes any out-of-vocabulary (OOV) words that are associated with the corpus from the custom model's words resource unless they were also added by another corpus or grammar, or they were modified in some way with the **Add custom words** or **Add a custom word** method. Removing a corpus does not affect the custom model until you train the model with the **Train a custom language model** method. You must use credentials for the instance of the service that owns a model to delete its corpora. **See also:** [Deleting a corpus from a custom language model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-manageCorpora#deleteCorpus).",
		Run: DeleteCorpus,
	}

	cmd.Flags().StringVarP(&DeleteCorpusCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom language model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&DeleteCorpusCorpusName, "corpus_name", "", "", "The name of the corpus for the custom language model.")

	cmd.MarkFlagRequired("customization_id")
	cmd.MarkFlagRequired("corpus_name")

	return cmd
}

func DeleteCorpus(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.DeleteCorpusOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(DeleteCorpusCustomizationID)
		}
		if flag.Name == "corpus_name" {
			optionsModel.SetCorpusName(DeleteCorpusCorpusName)
		}
	})

	_, responseErr := speechToText.DeleteCorpus(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var ListWordsCustomizationID string
var ListWordsWordType string
var ListWordsSort string

func getListWordsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-words",
		Short: "List custom words",
		Long: "Lists information about custom words from a custom language model. You can list all words from the custom model's words resource, only custom words that were added or modified by the user, or only out-of-vocabulary (OOV) words that were extracted from corpora or are recognized by grammars. You can also indicate the order in which the service is to return words; by default, the service lists words in ascending alphabetical order. You must use credentials for the instance of the service that owns a model to list information about its words. **See also:** [Listing words from a custom language model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-manageWords#listWords).",
		Run: ListWords,
	}

	cmd.Flags().StringVarP(&ListWordsCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom language model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&ListWordsWordType, "word_type", "", "", "The type of words to be listed from the custom language model's words resource:* `all` (the default) shows all words.* `user` shows only custom words that were added or modified by the user directly.* `corpora` shows only OOV that were extracted from corpora.* `grammars` shows only OOV words that are recognized by grammars.")
	cmd.Flags().StringVarP(&ListWordsSort, "sort", "", "", "Indicates the order in which the words are to be listed, `alphabetical` or by `count`. You can prepend an optional `+` or `-` to an argument to indicate whether the results are to be sorted in ascending or descending order. By default, words are sorted in ascending alphabetical order. For alphabetical ordering, the lexicographical precedence is numeric values, uppercase letters, and lowercase letters. For count ordering, values with the same count are ordered alphabetically. With the `curl` command, URL-encode the `+` symbol as `%2B`.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("customization_id")

	return cmd
}

func ListWords(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.ListWordsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(ListWordsCustomizationID)
		}
		if flag.Name == "word_type" {
			optionsModel.SetWordType(ListWordsWordType)
		}
		if flag.Name == "sort" {
			optionsModel.SetSort(ListWordsSort)
		}
	})

	result, _, responseErr := speechToText.ListWords(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var AddWordsCustomizationID string
var AddWordsWords string

func getAddWordsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "add-words",
		Short: "Add custom words",
		Long: "Adds one or more custom words to a custom language model. The service populates the words resource for a custom model with out-of-vocabulary (OOV) words from each corpus or grammar that is added to the model. You can use this method to add additional words or to modify existing words in the words resource. The words resource for a model can contain a maximum of 90 thousand custom (OOV) words. This includes words that the service extracts from corpora and grammars and words that you add directly. You must use credentials for the instance of the service that owns a model to add or modify custom words for the model. Adding or modifying custom words does not affect the custom model until you train the model for the new data by using the **Train a custom language model** method. You add custom words by providing a `CustomWords` object, which is an array of `CustomWord` objects, one per word. You must use the object's `word` parameter to identify the word that is to be added. You can also provide one or both of the optional `sounds_like` and `display_as` fields for each word.* The `sounds_like` field provides an array of one or more pronunciations for the word. Use the parameter to specify how the word can be pronounced by users. Use the parameter for words that are difficult to pronounce, foreign words, acronyms, and so on. For example, you might specify that the word `IEEE` can sound like `i triple e`. You can specify a maximum of five sounds-like pronunciations for a word.* The `display_as` field provides a different way of spelling the word in a transcript. Use the parameter when you want the word to appear different from its usual representation or from its spelling in training data. For example, you might indicate that the word `IBM(trademark)` is to be displayed as `IBM&trade;`. If you add a custom word that already exists in the words resource for the custom model, the new definition overwrites the existing data for the word. If the service encounters an error with the input data, it returns a failure code and does not add any of the words to the words resource. The call returns an HTTP 201 response code if the input data is valid. It then asynchronously processes the words to add them to the model's words resource. The time that it takes for the analysis to complete depends on the number of new words that you add but is generally faster than adding a corpus or grammar. You can monitor the status of the request by using the **List a custom language model** method to poll the model's status. Use a loop to check the status every 10 seconds. The method returns a `Customization` object that includes a `status` field. A status of `ready` means that the words have been added to the custom model. The service cannot accept requests to add new data or to train the model until the existing request completes. You can use the **List custom words** or **List a custom word** method to review the words that you add. Words with an invalid `sounds_like` field include an `error` field that describes the problem. You can use other words-related methods to correct errors, eliminate typos, and modify how words are pronounced as needed. **See also:*** [Working with custom words](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-corporaWords#workingWords)* [Add words to the custom language model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-languageCreate#addWords).",
		Run: AddWords,
	}

	cmd.Flags().StringVarP(&AddWordsCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom language model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&AddWordsWords, "words", "", "", "An array of `CustomWord` objects that provides information about each custom word that is to be added to or updated in the custom language model.")

	cmd.MarkFlagRequired("customization_id")
	cmd.MarkFlagRequired("words")

	return cmd
}

func AddWords(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.AddWordsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(AddWordsCustomizationID)
		}
		if flag.Name == "words" {
			var words []speechtotextv1.CustomWord
			decodeErr := json.Unmarshal([]byte(AddWordsWords), &words);
			utils.HandleError(decodeErr)

			optionsModel.SetWords(words)
		}
	})

	_, responseErr := speechToText.AddWords(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var AddWordCustomizationID string
var AddWordWordName string
var AddWordWord string
var AddWordSoundsLike []string
var AddWordDisplayAs string

func getAddWordCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "add-word",
		Short: "Add a custom word",
		Long: "Adds a custom word to a custom language model. The service populates the words resource for a custom model with out-of-vocabulary (OOV) words from each corpus or grammar that is added to the model. You can use this method to add a word or to modify an existing word in the words resource. The words resource for a model can contain a maximum of 90 thousand custom (OOV) words. This includes words that the service extracts from corpora and grammars and words that you add directly. You must use credentials for the instance of the service that owns a model to add or modify a custom word for the model. Adding or modifying a custom word does not affect the custom model until you train the model for the new data by using the **Train a custom language model** method. Use the `word_name` parameter to specify the custom word that is to be added or modified. Use the `CustomWord` object to provide one or both of the optional `sounds_like` and `display_as` fields for the word.* The `sounds_like` field provides an array of one or more pronunciations for the word. Use the parameter to specify how the word can be pronounced by users. Use the parameter for words that are difficult to pronounce, foreign words, acronyms, and so on. For example, you might specify that the word `IEEE` can sound like `i triple e`. You can specify a maximum of five sounds-like pronunciations for a word.* The `display_as` field provides a different way of spelling the word in a transcript. Use the parameter when you want the word to appear different from its usual representation or from its spelling in training data. For example, you might indicate that the word `IBM(trademark)` is to be displayed as `IBM&trade;`. If you add a custom word that already exists in the words resource for the custom model, the new definition overwrites the existing data for the word. If the service encounters an error, it does not add the word to the words resource. Use the **List a custom word** method to review the word that you add. **See also:*** [Working with custom words](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-corporaWords#workingWords)* [Add words to the custom language model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-languageCreate#addWords).",
		Run: AddWord,
	}

	cmd.Flags().StringVarP(&AddWordCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom language model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&AddWordWordName, "word_name", "", "", "The custom word that is to be added to or updated in the custom language model. Do not include spaces in the word. Use a `-` (dash) or `_` (underscore) to connect the tokens of compound words. URL-encode the word if it includes non-ASCII characters. For more information, see [Character encoding](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-corporaWords#charEncoding).")
	cmd.Flags().StringVarP(&AddWordWord, "word", "", "", "For the **Add custom words** method, you must specify the custom word that is to be added to or updated in the custom model. Do not include spaces in the word. Use a `-` (dash) or `_` (underscore) to connect the tokens of compound words. Omit this parameter for the **Add a custom word** method.")
	cmd.Flags().StringSliceVarP(&AddWordSoundsLike, "sounds_like", "", nil, "An array of sounds-like pronunciations for the custom word. Specify how words that are difficult to pronounce, foreign words, acronyms, and so on can be pronounced by users.* For a word that is not in the service's base vocabulary, omit the parameter to have the service automatically generate a sounds-like pronunciation for the word.* For a word that is in the service's base vocabulary, use the parameter to specify additional pronunciations for the word. You cannot override the default pronunciation of a word; pronunciations you add augment the pronunciation from the base vocabulary. A word can have at most five sounds-like pronunciations. A pronunciation can include at most 40 characters not including spaces.")
	cmd.Flags().StringVarP(&AddWordDisplayAs, "display_as", "", "", "An alternative spelling for the custom word when it appears in a transcript. Use the parameter when you want the word to have a spelling that is different from its usual representation or from its spelling in corpora training data.")

	cmd.MarkFlagRequired("customization_id")
	cmd.MarkFlagRequired("word_name")

	return cmd
}

func AddWord(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.AddWordOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(AddWordCustomizationID)
		}
		if flag.Name == "word_name" {
			optionsModel.SetWordName(AddWordWordName)
		}
		if flag.Name == "word" {
			optionsModel.SetWord(AddWordWord)
		}
		if flag.Name == "sounds_like" {
			optionsModel.SetSoundsLike(AddWordSoundsLike)
		}
		if flag.Name == "display_as" {
			optionsModel.SetDisplayAs(AddWordDisplayAs)
		}
	})

	_, responseErr := speechToText.AddWord(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var GetWordCustomizationID string
var GetWordWordName string

func getGetWordCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-word",
		Short: "Get a custom word",
		Long: "Gets information about a custom word from a custom language model. You must use credentials for the instance of the service that owns a model to list information about its words. **See also:** [Listing words from a custom language model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-manageWords#listWords).",
		Run: GetWord,
	}

	cmd.Flags().StringVarP(&GetWordCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom language model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&GetWordWordName, "word_name", "", "", "The custom word that is to be read from the custom language model. URL-encode the word if it includes non-ASCII characters. For more information, see [Character encoding](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-corporaWords#charEncoding).")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("customization_id")
	cmd.MarkFlagRequired("word_name")

	return cmd
}

func GetWord(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.GetWordOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(GetWordCustomizationID)
		}
		if flag.Name == "word_name" {
			optionsModel.SetWordName(GetWordWordName)
		}
	})

	result, _, responseErr := speechToText.GetWord(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteWordCustomizationID string
var DeleteWordWordName string

func getDeleteWordCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-word",
		Short: "Delete a custom word",
		Long: "Deletes a custom word from a custom language model. You can remove any word that you added to the custom model's words resource via any means. However, if the word also exists in the service's base vocabulary, the service removes only the custom pronunciation for the word; the word remains in the base vocabulary. Removing a custom word does not affect the custom model until you train the model with the **Train a custom language model** method. You must use credentials for the instance of the service that owns a model to delete its words. **See also:** [Deleting a word from a custom language model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-manageWords#deleteWord).",
		Run: DeleteWord,
	}

	cmd.Flags().StringVarP(&DeleteWordCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom language model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&DeleteWordWordName, "word_name", "", "", "The custom word that is to be deleted from the custom language model. URL-encode the word if it includes non-ASCII characters. For more information, see [Character encoding](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-corporaWords#charEncoding).")

	cmd.MarkFlagRequired("customization_id")
	cmd.MarkFlagRequired("word_name")

	return cmd
}

func DeleteWord(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.DeleteWordOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(DeleteWordCustomizationID)
		}
		if flag.Name == "word_name" {
			optionsModel.SetWordName(DeleteWordWordName)
		}
	})

	_, responseErr := speechToText.DeleteWord(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var ListGrammarsCustomizationID string

func getListGrammarsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-grammars",
		Short: "List grammars",
		Long: "Lists information about all grammars from a custom language model. The information includes the total number of out-of-vocabulary (OOV) words, name, and status of each grammar. You must use credentials for the instance of the service that owns a model to list its grammars. **See also:** [Listing grammars from a custom language model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-manageGrammars#listGrammars).",
		Run: ListGrammars,
	}

	cmd.Flags().StringVarP(&ListGrammarsCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom language model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("customization_id")

	return cmd
}

func ListGrammars(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.ListGrammarsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(ListGrammarsCustomizationID)
		}
	})

	result, _, responseErr := speechToText.ListGrammars(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var AddGrammarCustomizationID string
var AddGrammarGrammarName string
var AddGrammarGrammarFile string
var AddGrammarContentType string
var AddGrammarAllowOverwrite bool

func getAddGrammarCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "add-grammar",
		Short: "Add a grammar",
		Long: "Adds a single grammar file to a custom language model. Submit a plain text file in UTF-8 format that defines the grammar. Use multiple requests to submit multiple grammar files. You must use credentials for the instance of the service that owns a model to add a grammar to it. Adding a grammar does not affect the custom language model until you train the model for the new data by using the **Train a custom language model** method. The call returns an HTTP 201 response code if the grammar is valid. The service then asynchronously processes the contents of the grammar and automatically extracts new words that it finds. This can take a few seconds to complete depending on the size and complexity of the grammar, as well as the current load on the service. You cannot submit requests to add additional resources to the custom model or to train the model until the service's analysis of the grammar for the current request completes. Use the **Get a grammar** method to check the status of the analysis. The service populates the model's words resource with any word that is recognized by the grammar that is not found in the model's base vocabulary. These are referred to as out-of-vocabulary (OOV) words. You can use the **List custom words** method to examine the words resource and use other words-related methods to eliminate typos and modify how words are pronounced as needed. To add a grammar that has the same name as an existing grammar, set the `allow_overwrite` parameter to `true`; otherwise, the request fails. Overwriting an existing grammar causes the service to process the grammar file and extract OOV words anew. Before doing so, it removes any OOV words associated with the existing grammar from the model's words resource unless they were also added by another resource or they have been modified in some way with the **Add custom words** or **Add a custom word** method. The service limits the overall amount of data that you can add to a custom model to a maximum of 10 million total words from all sources combined. Also, you can add no more than 90 thousand OOV words to a model. This includes words that the service extracts from corpora and grammars and words that you add directly. **See also:*** [Understanding grammars](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-grammarUnderstand#grammarUnderstand)* [Add a grammar to the custom language model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-grammarAdd#addGrammar).",
		Run: AddGrammar,
	}

	cmd.Flags().StringVarP(&AddGrammarCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom language model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&AddGrammarGrammarName, "grammar_name", "", "", "The name of the new grammar for the custom language model. Use a localized name that matches the language of the custom model and reflects the contents of the grammar.* Include a maximum of 128 characters in the name.* Do not use characters that need to be URL-encoded. For example, do not use spaces, slashes, backslashes, colons, ampersands, double quotes, plus signs, equals signs, questions marks, and so on in the name. (The service does not prevent the use of these characters. But because they must be URL-encoded wherever used, their use is strongly discouraged.)* Do not use the name of an existing grammar or corpus that is already defined for the custom model.* Do not use the name `user`, which is reserved by the service to denote custom words that are added or modified by the user.* Do not use the name `base_lm` or `default_lm`. Both names are reserved for future use by the service.")
	cmd.Flags().StringVarP(&AddGrammarGrammarFile, "grammar_file", "", "", "A plain text file that contains the grammar in the format specified by the `Content-Type` header. Encode the file in UTF-8 (ASCII is a subset of UTF-8). Using any other encoding can lead to issues when compiling the grammar or to unexpected results in decoding. The service ignores an encoding that is specified in the header of the grammar. With the `curl` command, use the `--data-binary` option to upload the file for the request.")
	cmd.Flags().StringVarP(&AddGrammarContentType, "content_type", "", "", "The format (MIME type) of the grammar file:* `application/srgs` for Augmented Backus-Naur Form (ABNF), which uses a plain-text representation that is similar to traditional BNF grammars.* `application/srgs+xml` for XML Form, which uses XML elements to represent the grammar.")
	cmd.Flags().BoolVarP(&AddGrammarAllowOverwrite, "allow_overwrite", "", false, "If `true`, the specified grammar overwrites an existing grammar with the same name. If `false`, the request fails if a grammar with the same name already exists. The parameter has no effect if a grammar with the same name does not already exist.")

	cmd.MarkFlagRequired("customization_id")
	cmd.MarkFlagRequired("grammar_name")
	cmd.MarkFlagRequired("grammar_file")
	cmd.MarkFlagRequired("content_type")

	return cmd
}

func AddGrammar(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.AddGrammarOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(AddGrammarCustomizationID)
		}
		if flag.Name == "grammar_name" {
			optionsModel.SetGrammarName(AddGrammarGrammarName)
		}
		if flag.Name == "grammar_file" {
			grammar_file, fileErr := os.Open(AddGrammarGrammarFile)
			utils.HandleError(fileErr)

			optionsModel.SetGrammarFile(grammar_file)
		}
		if flag.Name == "content_type" {
			optionsModel.SetContentType(AddGrammarContentType)
		}
		if flag.Name == "allow_overwrite" {
			optionsModel.SetAllowOverwrite(AddGrammarAllowOverwrite)
		}
	})

	_, responseErr := speechToText.AddGrammar(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var GetGrammarCustomizationID string
var GetGrammarGrammarName string

func getGetGrammarCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-grammar",
		Short: "Get a grammar",
		Long: "Gets information about a grammar from a custom language model. The information includes the total number of out-of-vocabulary (OOV) words, name, and status of the grammar. You must use credentials for the instance of the service that owns a model to list its grammars. **See also:** [Listing grammars from a custom language model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-manageGrammars#listGrammars).",
		Run: GetGrammar,
	}

	cmd.Flags().StringVarP(&GetGrammarCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom language model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&GetGrammarGrammarName, "grammar_name", "", "", "The name of the grammar for the custom language model.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("customization_id")
	cmd.MarkFlagRequired("grammar_name")

	return cmd
}

func GetGrammar(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.GetGrammarOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(GetGrammarCustomizationID)
		}
		if flag.Name == "grammar_name" {
			optionsModel.SetGrammarName(GetGrammarGrammarName)
		}
	})

	result, _, responseErr := speechToText.GetGrammar(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteGrammarCustomizationID string
var DeleteGrammarGrammarName string

func getDeleteGrammarCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-grammar",
		Short: "Delete a grammar",
		Long: "Deletes an existing grammar from a custom language model. The service removes any out-of-vocabulary (OOV) words associated with the grammar from the custom model's words resource unless they were also added by another resource or they were modified in some way with the **Add custom words** or **Add a custom word** method. Removing a grammar does not affect the custom model until you train the model with the **Train a custom language model** method. You must use credentials for the instance of the service that owns a model to delete its grammar. **See also:** [Deleting a grammar from a custom language model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-manageGrammars#deleteGrammar).",
		Run: DeleteGrammar,
	}

	cmd.Flags().StringVarP(&DeleteGrammarCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom language model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&DeleteGrammarGrammarName, "grammar_name", "", "", "The name of the grammar for the custom language model.")

	cmd.MarkFlagRequired("customization_id")
	cmd.MarkFlagRequired("grammar_name")

	return cmd
}

func DeleteGrammar(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.DeleteGrammarOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(DeleteGrammarCustomizationID)
		}
		if flag.Name == "grammar_name" {
			optionsModel.SetGrammarName(DeleteGrammarGrammarName)
		}
	})

	_, responseErr := speechToText.DeleteGrammar(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var CreateAcousticModelName string
var CreateAcousticModelBaseModelName string
var CreateAcousticModelDescription string

func getCreateAcousticModelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-acoustic-model",
		Short: "Create a custom acoustic model",
		Long: "Creates a new custom acoustic model for a specified base model. The custom acoustic model can be used only with the base model for which it is created. The model is owned by the instance of the service whose credentials are used to create it. **See also:** [Create a custom acoustic model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-acoustic#createModel-acoustic).",
		Run: CreateAcousticModel,
	}

	cmd.Flags().StringVarP(&CreateAcousticModelName, "name", "", "", "A user-defined name for the new custom acoustic model. Use a name that is unique among all custom acoustic models that you own. Use a localized name that matches the language of the custom model. Use a name that describes the acoustic environment of the custom model, such as `Mobile custom model` or `Noisy car custom model`.")
	cmd.Flags().StringVarP(&CreateAcousticModelBaseModelName, "base_model_name", "", "", "The name of the base language model that is to be customized by the new custom acoustic model. The new custom model can be used only with the base model that it customizes. To determine whether a base model supports acoustic model customization, refer to [Language support for customization](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-customization#languageSupport).")
	cmd.Flags().StringVarP(&CreateAcousticModelDescription, "description", "", "", "A description of the new custom acoustic model. Use a localized description that matches the language of the custom model.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("base_model_name")

	return cmd
}

func CreateAcousticModel(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.CreateAcousticModelOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "name" {
			optionsModel.SetName(CreateAcousticModelName)
		}
		if flag.Name == "base_model_name" {
			optionsModel.SetBaseModelName(CreateAcousticModelBaseModelName)
		}
		if flag.Name == "description" {
			optionsModel.SetDescription(CreateAcousticModelDescription)
		}
	})

	result, _, responseErr := speechToText.CreateAcousticModel(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ListAcousticModelsLanguage string

func getListAcousticModelsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-acoustic-models",
		Short: "List custom acoustic models",
		Long: "Lists information about all custom acoustic models that are owned by an instance of the service. Use the `language` parameter to see all custom acoustic models for the specified language. Omit the parameter to see all custom acoustic models for all languages. You must use credentials for the instance of the service that owns a model to list information about it. **See also:** [Listing custom acoustic models](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-manageAcousticModels#listModels-acoustic).",
		Run: ListAcousticModels,
	}

	cmd.Flags().StringVarP(&ListAcousticModelsLanguage, "language", "", "", "The identifier of the language for which custom language or custom acoustic models are to be returned (for example, `en-US`). Omit the parameter to see all custom language or custom acoustic models that are owned by the requesting credentials.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")


	return cmd
}

func ListAcousticModels(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.ListAcousticModelsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "language" {
			optionsModel.SetLanguage(ListAcousticModelsLanguage)
		}
	})

	result, _, responseErr := speechToText.ListAcousticModels(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetAcousticModelCustomizationID string

func getGetAcousticModelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-acoustic-model",
		Short: "Get a custom acoustic model",
		Long: "Gets information about a specified custom acoustic model. You must use credentials for the instance of the service that owns a model to list information about it. **See also:** [Listing custom acoustic models](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-manageAcousticModels#listModels-acoustic).",
		Run: GetAcousticModel,
	}

	cmd.Flags().StringVarP(&GetAcousticModelCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("customization_id")

	return cmd
}

func GetAcousticModel(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.GetAcousticModelOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(GetAcousticModelCustomizationID)
		}
	})

	result, _, responseErr := speechToText.GetAcousticModel(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteAcousticModelCustomizationID string

func getDeleteAcousticModelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-acoustic-model",
		Short: "Delete a custom acoustic model",
		Long: "Deletes an existing custom acoustic model. The custom model cannot be deleted if another request, such as adding an audio resource to the model, is currently being processed. You must use credentials for the instance of the service that owns a model to delete it. **See also:** [Deleting a custom acoustic model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-manageAcousticModels#deleteModel-acoustic).",
		Run: DeleteAcousticModel,
	}

	cmd.Flags().StringVarP(&DeleteAcousticModelCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")

	cmd.MarkFlagRequired("customization_id")

	return cmd
}

func DeleteAcousticModel(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.DeleteAcousticModelOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(DeleteAcousticModelCustomizationID)
		}
	})

	_, responseErr := speechToText.DeleteAcousticModel(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var TrainAcousticModelCustomizationID string
var TrainAcousticModelCustomLanguageModelID string

func getTrainAcousticModelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "train-acoustic-model",
		Short: "Train a custom acoustic model",
		Long: "Initiates the training of a custom acoustic model with new or changed audio resources. After adding or deleting audio resources for a custom acoustic model, use this method to begin the actual training of the model on the latest audio data. The custom acoustic model does not reflect its changed data until you train it. You must use credentials for the instance of the service that owns a model to train it. The training method is asynchronous. It can take on the order of minutes or hours to complete depending on the total amount of audio data on which the custom acoustic model is being trained and the current load on the service. Typically, training a custom acoustic model takes approximately two to four times the length of its audio data. The range of time depends on the model being trained and the nature of the audio, such as whether the audio is clean or noisy. The method returns an HTTP 200 response code to indicate that the training process has begun. You can monitor the status of the training by using the **Get a custom acoustic model** method to poll the model's status. Use a loop to check the status once a minute. The method returns an `AcousticModel` object that includes `status` and `progress` fields. A status of `available` indicates that the custom model is trained and ready to use. The service cannot train a model while it is handling another request for the model. The service cannot accept subsequent training requests, or requests to add new audio resources, until the existing training request completes. You can use the optional `custom_language_model_id` parameter to specify the GUID of a separately created custom language model that is to be used during training. Train with a custom language model if you have verbatim transcriptions of the audio files that you have added to the custom model or you have either corpora (text files) or a list of words that are relevant to the contents of the audio files. Both of the custom models must be based on the same version of the same base model for training to succeed. **See also:*** [Train the custom acoustic model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-acoustic#trainModel-acoustic)* [Using custom acoustic and custom language models together](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-useBoth#useBoth) ### Training failures Training can fail to start for the following reasons:* The service is currently handling another request for the custom model, such as another training request or a request to add audio resources to the model.* The custom model contains less than 10 minutes or more than 200 hours of audio data.* You passed an incompatible custom language model with the `custom_language_model_id` query parameter. Both custom models must be based on the same version of the same base model.* The custom model contains one or more invalid audio resources. You can correct the invalid audio resources or set the `strict` parameter to `false` to exclude the invalid resources from the training. The model must contain at least one valid resource for training to succeed.",
		Run: TrainAcousticModel,
	}

	cmd.Flags().StringVarP(&TrainAcousticModelCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&TrainAcousticModelCustomLanguageModelID, "custom_language_model_id", "", "", "The customization ID (GUID) of a custom language model that is to be used during training of the custom acoustic model. Specify a custom language model that has been trained with verbatim transcriptions of the audio resources or that contains words that are relevant to the contents of the audio resources. The custom language model must be based on the same version of the same base model as the custom acoustic model. The credentials specified with the request must own both custom models.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("customization_id")

	return cmd
}

func TrainAcousticModel(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.TrainAcousticModelOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(TrainAcousticModelCustomizationID)
		}
		if flag.Name == "custom_language_model_id" {
			optionsModel.SetCustomLanguageModelID(TrainAcousticModelCustomLanguageModelID)
		}
	})

	result, _, responseErr := speechToText.TrainAcousticModel(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ResetAcousticModelCustomizationID string

func getResetAcousticModelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "reset-acoustic-model",
		Short: "Reset a custom acoustic model",
		Long: "Resets a custom acoustic model by removing all audio resources from the model. Resetting a custom acoustic model initializes the model to its state when it was first created. Metadata such as the name and language of the model are preserved, but the model's audio resources are removed and must be re-created. The service cannot reset a model while it is handling another request for the model. The service cannot accept subsequent requests for the model until the existing reset request completes. You must use credentials for the instance of the service that owns a model to reset it. **See also:** [Resetting a custom acoustic model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-manageAcousticModels#resetModel-acoustic).",
		Run: ResetAcousticModel,
	}

	cmd.Flags().StringVarP(&ResetAcousticModelCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")

	cmd.MarkFlagRequired("customization_id")

	return cmd
}

func ResetAcousticModel(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.ResetAcousticModelOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(ResetAcousticModelCustomizationID)
		}
	})

	_, responseErr := speechToText.ResetAcousticModel(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var UpgradeAcousticModelCustomizationID string
var UpgradeAcousticModelCustomLanguageModelID string
var UpgradeAcousticModelForce bool

func getUpgradeAcousticModelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "upgrade-acoustic-model",
		Short: "Upgrade a custom acoustic model",
		Long: "Initiates the upgrade of a custom acoustic model to the latest version of its base language model. The upgrade method is asynchronous. It can take on the order of minutes or hours to complete depending on the amount of data in the custom model and the current load on the service; typically, upgrade takes approximately twice the length of the total audio contained in the custom model. A custom model must be in the `ready` or `available` state to be upgraded. You must use credentials for the instance of the service that owns a model to upgrade it. The method returns an HTTP 200 response code to indicate that the upgrade process has begun successfully. You can monitor the status of the upgrade by using the **Get a custom acoustic model** method to poll the model's status. The method returns an `AcousticModel` object that includes `status` and `progress` fields. Use a loop to check the status once a minute. While it is being upgraded, the custom model has the status `upgrading`. When the upgrade is complete, the model resumes the status that it had prior to upgrade. The service cannot upgrade a model while it is handling another request for the model. The service cannot accept subsequent requests for the model until the existing upgrade request completes. If the custom acoustic model was trained with a separately created custom language model, you must use the `custom_language_model_id` parameter to specify the GUID of that custom language model. The custom language model must be upgraded before the custom acoustic model can be upgraded. Omit the parameter if the custom acoustic model was not trained with a custom language model. **See also:** [Upgrading a custom acoustic model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-customUpgrade#upgradeAcoustic).",
		Run: UpgradeAcousticModel,
	}

	cmd.Flags().StringVarP(&UpgradeAcousticModelCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&UpgradeAcousticModelCustomLanguageModelID, "custom_language_model_id", "", "", "If the custom acoustic model was trained with a custom language model, the customization ID (GUID) of that custom language model. The custom language model must be upgraded before the custom acoustic model can be upgraded. The credentials specified with the request must own both custom models.")
	cmd.Flags().BoolVarP(&UpgradeAcousticModelForce, "force", "", false, "If `true`, forces the upgrade of a custom acoustic model for which no input data has been modified since it was last trained. Use this parameter only to force the upgrade of a custom acoustic model that is trained with a custom language model, and only if you receive a 400 response code and the message `No input data modified since last training`. See [Upgrading a custom acoustic model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-customUpgrade#upgradeAcoustic).")

	cmd.MarkFlagRequired("customization_id")

	return cmd
}

func UpgradeAcousticModel(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.UpgradeAcousticModelOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(UpgradeAcousticModelCustomizationID)
		}
		if flag.Name == "custom_language_model_id" {
			optionsModel.SetCustomLanguageModelID(UpgradeAcousticModelCustomLanguageModelID)
		}
		if flag.Name == "force" {
			optionsModel.SetForce(UpgradeAcousticModelForce)
		}
	})

	_, responseErr := speechToText.UpgradeAcousticModel(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var ListAudioCustomizationID string

func getListAudioCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-audio",
		Short: "List audio resources",
		Long: "Lists information about all audio resources from a custom acoustic model. The information includes the name of the resource and information about its audio data, such as its duration. It also includes the status of the audio resource, which is important for checking the service's analysis of the resource in response to a request to add it to the custom acoustic model. You must use credentials for the instance of the service that owns a model to list its audio resources. **See also:** [Listing audio resources for a custom acoustic model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-manageAudio#listAudio).",
		Run: ListAudio,
	}

	cmd.Flags().StringVarP(&ListAudioCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("customization_id")

	return cmd
}

func ListAudio(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.ListAudioOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(ListAudioCustomizationID)
		}
	})

	result, _, responseErr := speechToText.ListAudio(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var AddAudioCustomizationID string
var AddAudioAudioName string
var AddAudioAudioResource string
var AddAudioContentType string
var AddAudioContainedContentType string
var AddAudioAllowOverwrite bool

func getAddAudioCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "add-audio",
		Short: "Add an audio resource",
		Long: "Adds an audio resource to a custom acoustic model. Add audio content that reflects the acoustic characteristics of the audio that you plan to transcribe. You must use credentials for the instance of the service that owns a model to add an audio resource to it. Adding audio data does not affect the custom acoustic model until you train the model for the new data by using the **Train a custom acoustic model** method. You can add individual audio files or an archive file that contains multiple audio files. Adding multiple audio files via a single archive file is significantly more efficient than adding each file individually. You can add audio resources in any format that the service supports for speech recognition. You can use this method to add any number of audio resources to a custom model by calling the method once for each audio or archive file. You can add multiple different audio resources at the same time. You must add a minimum of 10 minutes and a maximum of 200 hours of audio that includes speech, not just silence, to a custom acoustic model before you can train it. No audio resource, audio- or archive-type, can be larger than 100 MB. To add an audio resource that has the same name as an existing audio resource, set the `allow_overwrite` parameter to `true`; otherwise, the request fails. The method is asynchronous. It can take several seconds to complete depending on the duration of the audio and, in the case of an archive file, the total number of audio files being processed. The service returns a 201 response code if the audio is valid. It then asynchronously analyzes the contents of the audio file or files and automatically extracts information about the audio such as its length, sampling rate, and encoding. You cannot submit requests to train or upgrade the model until the service's analysis of all audio resources for current requests completes. To determine the status of the service's analysis of the audio, use the **Get an audio resource** method to poll the status of the audio. The method accepts the customization ID of the custom model and the name of the audio resource, and it returns the status of the resource. Use a loop to check the status of the audio every few seconds until it becomes `ok`. **See also:** [Add audio to the custom acoustic model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-acoustic#addAudio). ### Content types for audio-type resources You can add an individual audio file in any format that the service supports for speech recognition. For an audio-type resource, use the `Content-Type` parameter to specify the audio format (MIME type) of the audio file, including specifying the sampling rate, channels, and endianness where indicated.* `audio/alaw` (Specify the sampling rate (`rate`) of the audio.)* `audio/basic` (Use only with narrowband models.)* `audio/flac`* `audio/g729` (Use only with narrowband models.)* `audio/l16` (Specify the sampling rate (`rate`) and optionally the number of channels (`channels`) and endianness (`endianness`) of the audio.)* `audio/mp3`* `audio/mpeg`* `audio/mulaw` (Specify the sampling rate (`rate`) of the audio.)* `audio/ogg` (The service automatically detects the codec of the input audio.)* `audio/ogg;codecs=opus`* `audio/ogg;codecs=vorbis`* `audio/wav` (Provide audio with a maximum of nine channels.)* `audio/webm` (The service automatically detects the codec of the input audio.)* `audio/webm;codecs=opus`* `audio/webm;codecs=vorbis` The sampling rate of an audio file must match the sampling rate of the base model for the custom model: for broadband models, at least 16 kHz; for narrowband models, at least 8 kHz. If the sampling rate of the audio is higher than the minimum required rate, the service down-samples the audio to the appropriate rate. If the sampling rate of the audio is lower than the minimum required rate, the service labels the audio file as `invalid`. **See also:** [Audio formats](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-audio-formats#audio-formats). ### Content types for archive-type resources You can add an archive file (**.zip** or **.tar.gz** file) that contains audio files in any format that the service supports for speech recognition. For an archive-type resource, use the `Content-Type` parameter to specify the media type of the archive file:* `application/zip` for a **.zip** file* `application/gzip` for a **.tar.gz** file. When you add an archive-type resource, the `Contained-Content-Type` header is optional depending on the format of the files that you are adding: * For audio files of type `audio/alaw`, `audio/basic`, `audio/l16`, or `audio/mulaw`, you must use the `Contained-Content-Type` header to specify the format of the contained audio files. Include the `rate`, `channels`, and `endianness` parameters where necessary. In this case, all audio files contained in the archive file must have the same audio format. * For audio files of all other types, you can omit the `Contained-Content-Type` header. In this case, the audio files contained in the archive file can have any of the formats not listed in the previous bullet. The audio files do not need to have the same format. Do not use the `Contained-Content-Type` header when adding an audio-type resource. ### Naming restrictions for embedded audio files The name of an audio file that is contained in an archive-type resource can include a maximum of 128 characters. This includes the file extension and all elements of the name (for example, slashes).",
		Run: AddAudio,
	}

	cmd.Flags().StringVarP(&AddAudioCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&AddAudioAudioName, "audio_name", "", "", "The name of the new audio resource for the custom acoustic model. Use a localized name that matches the language of the custom model and reflects the contents of the resource.* Include a maximum of 128 characters in the name.* Do not use characters that need to be URL-encoded. For example, do not use spaces, slashes, backslashes, colons, ampersands, double quotes, plus signs, equals signs, questions marks, and so on in the name. (The service does not prevent the use of these characters. But because they must be URL-encoded wherever used, their use is strongly discouraged.)* Do not use the name of an audio resource that has already been added to the custom model.")
	cmd.Flags().StringVarP(&AddAudioAudioResource, "audio_resource", "", "", "The audio resource that is to be added to the custom acoustic model, an individual audio file or an archive file. With the `curl` command, use the `--data-binary` option to upload the file for the request.")
	cmd.Flags().StringVarP(&AddAudioContentType, "content_type", "", "", "For an audio-type resource, the format (MIME type) of the audio. For more information, see **Content types for audio-type resources** in the method description. For an archive-type resource, the media type of the archive file. For more information, see **Content types for archive-type resources** in the method description.")
	cmd.Flags().StringVarP(&AddAudioContainedContentType, "contained_content_type", "", "", "**For an archive-type resource,** specify the format of the audio files that are contained in the archive file if they are of type `audio/alaw`, `audio/basic`, `audio/l16`, or `audio/mulaw`. Include the `rate`, `channels`, and `endianness` parameters where necessary. In this case, all audio files that are contained in the archive file must be of the indicated type. For all other audio formats, you can omit the header. In this case, the audio files can be of multiple types as long as they are not of the types listed in the previous paragraph. The parameter accepts all of the audio formats that are supported for use with speech recognition. For more information, see **Content types for audio-type resources** in the method description. **For an audio-type resource,** omit the header.")
	cmd.Flags().BoolVarP(&AddAudioAllowOverwrite, "allow_overwrite", "", false, "If `true`, the specified audio resource overwrites an existing audio resource with the same name. If `false`, the request fails if an audio resource with the same name already exists. The parameter has no effect if an audio resource with the same name does not already exist.")

	cmd.MarkFlagRequired("customization_id")
	cmd.MarkFlagRequired("audio_name")
	cmd.MarkFlagRequired("audio_resource")

	return cmd
}

func AddAudio(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.AddAudioOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(AddAudioCustomizationID)
		}
		if flag.Name == "audio_name" {
			optionsModel.SetAudioName(AddAudioAudioName)
		}
		if flag.Name == "audio_resource" {
			audio_resource, fileErr := os.Open(AddAudioAudioResource)
			utils.HandleError(fileErr)

			optionsModel.SetAudioResource(audio_resource)
		}
		if flag.Name == "content_type" {
			optionsModel.SetContentType(AddAudioContentType)
		}
		if flag.Name == "contained_content_type" {
			optionsModel.SetContainedContentType(AddAudioContainedContentType)
		}
		if flag.Name == "allow_overwrite" {
			optionsModel.SetAllowOverwrite(AddAudioAllowOverwrite)
		}
	})

	_, responseErr := speechToText.AddAudio(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var GetAudioCustomizationID string
var GetAudioAudioName string

func getGetAudioCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-audio",
		Short: "Get an audio resource",
		Long: "Gets information about an audio resource from a custom acoustic model. The method returns an `AudioListing` object whose fields depend on the type of audio resource that you specify with the method's `audio_name` parameter:* **For an audio-type resource,** the object's fields match those of an `AudioResource` object: `duration`, `name`, `details`, and `status`.* **For an archive-type resource,** the object includes a `container` field whose fields match those of an `AudioResource` object. It also includes an `audio` field, which contains an array of `AudioResource` objects that provides information about the audio files that are contained in the archive. The information includes the status of the specified audio resource. The status is important for checking the service's analysis of a resource that you add to the custom model.* For an audio-type resource, the `status` field is located in the `AudioListing` object.* For an archive-type resource, the `status` field is located in the `AudioResource` object that is returned in the `container` field. You must use credentials for the instance of the service that owns a model to list its audio resources. **See also:** [Listing audio resources for a custom acoustic model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-manageAudio#listAudio).",
		Run: GetAudio,
	}

	cmd.Flags().StringVarP(&GetAudioCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&GetAudioAudioName, "audio_name", "", "", "The name of the audio resource for the custom acoustic model.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("customization_id")
	cmd.MarkFlagRequired("audio_name")

	return cmd
}

func GetAudio(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.GetAudioOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(GetAudioCustomizationID)
		}
		if flag.Name == "audio_name" {
			optionsModel.SetAudioName(GetAudioAudioName)
		}
	})

	result, _, responseErr := speechToText.GetAudio(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteAudioCustomizationID string
var DeleteAudioAudioName string

func getDeleteAudioCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-audio",
		Short: "Delete an audio resource",
		Long: "Deletes an existing audio resource from a custom acoustic model. Deleting an archive-type audio resource removes the entire archive of files. The service does not allow deletion of individual files from an archive resource. Removing an audio resource does not affect the custom model until you train the model on its updated data by using the **Train a custom acoustic model** method. You can delete an existing audio resource from a model while a different resource is being added to the model. You must use credentials for the instance of the service that owns a model to delete its audio resources. **See also:** [Deleting an audio resource from a custom acoustic model](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-manageAudio#deleteAudio).",
		Run: DeleteAudio,
	}

	cmd.Flags().StringVarP(&DeleteAudioCustomizationID, "customization_id", "", "", "The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the request with credentials for the instance of the service that owns the custom model.")
	cmd.Flags().StringVarP(&DeleteAudioAudioName, "audio_name", "", "", "The name of the audio resource for the custom acoustic model.")

	cmd.MarkFlagRequired("customization_id")
	cmd.MarkFlagRequired("audio_name")

	return cmd
}

func DeleteAudio(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.DeleteAudioOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customization_id" {
			optionsModel.SetCustomizationID(DeleteAudioCustomizationID)
		}
		if flag.Name == "audio_name" {
			optionsModel.SetAudioName(DeleteAudioAudioName)
		}
	})

	_, responseErr := speechToText.DeleteAudio(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var DeleteUserDataCustomerID string

func getDeleteUserDataCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-user-data",
		Short: "Delete labeled data",
		Long: "Deletes all data that is associated with a specified customer ID. The method deletes all data for the customer ID, regardless of the method by which the information was added. The method has no effect if no data is associated with the customer ID. You must issue the request with credentials for the same instance of the service that was used to associate the customer ID with the data. You associate a customer ID with data by passing the `X-Watson-Metadata` header with a request that passes the data. **See also:** [Information security](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-information-security#information-security).",
		Run: DeleteUserData,
	}

	cmd.Flags().StringVarP(&DeleteUserDataCustomerID, "customer_id", "", "", "The customer ID for which all data is to be deleted.")

	cmd.MarkFlagRequired("customer_id")

	return cmd
}

func DeleteUserData(cmd *cobra.Command, args []string) {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(speechToTextErr)

	optionsModel := speechtotextv1.DeleteUserDataOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customer_id" {
			optionsModel.SetCustomerID(DeleteUserDataCustomerID)
		}
	})

	_, responseErr := speechToText.DeleteUserData(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}
