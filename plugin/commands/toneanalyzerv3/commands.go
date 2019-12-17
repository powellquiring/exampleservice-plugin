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

package toneanalyzerv3

import (
	"cli-watson-plugin/utils"
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/watson-developer-cloud/go-sdk/toneanalyzerv3"
)

var ToneToneInput string
var ToneBody string
var ToneContentType string
var ToneSentences bool
var ToneTones []string
var ToneContentLanguage string
var ToneAcceptLanguage string

func getToneCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "tone",
		Short: "Analyze general tone",
		Long: "Use the general-purpose endpoint to analyze the tone of your input content. The service analyzes the content for emotional and language tones. The method always analyzes the tone of the full document; by default, it also analyzes the tone of each individual sentence of the content. You can submit no more than 128 KB of total input content and no more than 1000 individual sentences in JSON, plain text, or HTML format. The service analyzes the first 1000 sentences for document-level analysis and only the first 100 sentences for sentence-level analysis. Per the JSON specification, the default character encoding for JSON content is effectively always UTF-8; per the HTTP specification, the default encoding for plain text and HTML is ISO-8859-1 (effectively, the ASCII character set). When specifying a content type of plain text or HTML, include the `charset` parameter to indicate the character encoding of the input text; for example: `Content-Type: text/plain;charset=utf-8`. For `text/html`, the service removes HTML tags and analyzes only the textual content. **See also:** [Using the general-purpose endpoint](https://cloud.ibm.com/docs/services/tone-analyzer?topic=tone-analyzer-utgpe#utgpe).",
		Run: Tone,
	}

	cmd.Flags().StringVarP(&ToneToneInput, "tone_input", "", "", "JSON, plain text, or HTML input that contains the content to be analyzed. For JSON input, provide an object of type `ToneInput`.")
	cmd.Flags().StringVarP(&ToneBody, "body", "", "", "JSON, plain text, or HTML input that contains the content to be analyzed. For JSON input, provide an object of type `ToneInput`.")
	cmd.Flags().StringVarP(&ToneContentType, "content_type", "", "", "The type of the input. A character encoding can be specified by including a `charset` parameter. For example, 'text/plain;charset=utf-8'.")
	cmd.Flags().BoolVarP(&ToneSentences, "sentences", "", false, "Indicates whether the service is to return an analysis of each individual sentence in addition to its analysis of the full document. If `true` (the default), the service returns results for each sentence.")
	cmd.Flags().StringSliceVarP(&ToneTones, "tones", "", nil, "**`2017-09-21`:** Deprecated. The service continues to accept the parameter for backward-compatibility, but the parameter no longer affects the response. **`2016-05-19`:** A comma-separated list of tones for which the service is to return its analysis of the input; the indicated tones apply both to the full document and to individual sentences of the document. You can specify one or more of the valid values. Omit the parameter to request results for all three tones.")
	cmd.Flags().StringVarP(&ToneContentLanguage, "content_language", "", "", "The language of the input text for the request: English or French. Regional variants are treated as their parent language; for example, `en-US` is interpreted as `en`. The input content must match the specified language. Do not submit content that contains both languages. You can use different languages for **Content-Language** and **Accept-Language**.* **`2017-09-21`:** Accepts `en` or `fr`.* **`2016-05-19`:** Accepts only `en`.")
	cmd.Flags().StringVarP(&ToneAcceptLanguage, "accept_language", "", "", "The desired language of the response. For two-character arguments, regional variants are treated as their parent language; for example, `en-US` is interpreted as `en`. You can use different languages for **Content-Language** and **Accept-Language**.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("version")

	return cmd
}

func Tone(cmd *cobra.Command, args []string) {
	toneAnalyzer, toneAnalyzerErr := toneanalyzerv3.
		NewToneAnalyzerV3(&toneanalyzerv3.ToneAnalyzerV3Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(toneAnalyzerErr)

	optionsModel := toneanalyzerv3.ToneOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "tone_input" {
			var tone_input toneanalyzerv3.ToneInput
			decodeErr := json.Unmarshal([]byte(ToneToneInput), &tone_input);
			utils.HandleError(decodeErr)

			optionsModel.SetToneInput(&tone_input)
		}
		if flag.Name == "body" {
			optionsModel.SetBody(ToneBody)
		}
		if flag.Name == "content_type" {
			optionsModel.SetContentType(ToneContentType)
		}
		if flag.Name == "sentences" {
			optionsModel.SetSentences(ToneSentences)
		}
		if flag.Name == "tones" {
			optionsModel.SetTones(ToneTones)
		}
		if flag.Name == "content_language" {
			optionsModel.SetContentLanguage(ToneContentLanguage)
		}
		if flag.Name == "accept_language" {
			optionsModel.SetAcceptLanguage(ToneAcceptLanguage)
		}
	})

	result, _, responseErr := toneAnalyzer.Tone(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ToneChatUtterances string
var ToneChatContentLanguage string
var ToneChatAcceptLanguage string

func getToneChatCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "tone-chat",
		Short: "Analyze customer-engagement tone",
		Long: "Use the customer-engagement endpoint to analyze the tone of customer service and customer support conversations. For each utterance of a conversation, the method reports the most prevalent subset of the following seven tones: sad, frustrated, satisfied, excited, polite, impolite, and sympathetic. If you submit more than 50 utterances, the service returns a warning for the overall content and analyzes only the first 50 utterances. If you submit a single utterance that contains more than 500 characters, the service returns an error for that utterance and does not analyze the utterance. The request fails if all utterances have more than 500 characters. Per the JSON specification, the default character encoding for JSON content is effectively always UTF-8. **See also:** [Using the customer-engagement endpoint](https://cloud.ibm.com/docs/services/tone-analyzer?topic=tone-analyzer-utco#utco).",
		Run: ToneChat,
	}

	cmd.Flags().StringVarP(&ToneChatUtterances, "utterances", "", "", "An array of `Utterance` objects that provides the input content that the service is to analyze.")
	cmd.Flags().StringVarP(&ToneChatContentLanguage, "content_language", "", "", "The language of the input text for the request: English or French. Regional variants are treated as their parent language; for example, `en-US` is interpreted as `en`. The input content must match the specified language. Do not submit content that contains both languages. You can use different languages for **Content-Language** and **Accept-Language**.* **`2017-09-21`:** Accepts `en` or `fr`.* **`2016-05-19`:** Accepts only `en`.")
	cmd.Flags().StringVarP(&ToneChatAcceptLanguage, "accept_language", "", "", "The desired language of the response. For two-character arguments, regional variants are treated as their parent language; for example, `en-US` is interpreted as `en`. You can use different languages for **Content-Language** and **Accept-Language**.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("utterances")
	cmd.MarkFlagRequired("version")

	return cmd
}

func ToneChat(cmd *cobra.Command, args []string) {
	toneAnalyzer, toneAnalyzerErr := toneanalyzerv3.
		NewToneAnalyzerV3(&toneanalyzerv3.ToneAnalyzerV3Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(toneAnalyzerErr)

	optionsModel := toneanalyzerv3.ToneChatOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "utterances" {
			var utterances []toneanalyzerv3.Utterance
			decodeErr := json.Unmarshal([]byte(ToneChatUtterances), &utterances);
			utils.HandleError(decodeErr)

			optionsModel.SetUtterances(utterances)
		}
		if flag.Name == "content_language" {
			optionsModel.SetContentLanguage(ToneChatContentLanguage)
		}
		if flag.Name == "accept_language" {
			optionsModel.SetAcceptLanguage(ToneChatAcceptLanguage)
		}
	})

	result, _, responseErr := toneAnalyzer.ToneChat(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}
