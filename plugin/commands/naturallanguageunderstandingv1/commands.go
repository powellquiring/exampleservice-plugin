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
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/watson-developer-cloud/go-sdk/naturallanguageunderstandingv1"
)

var AnalyzeFeatures string
var AnalyzeText string
var AnalyzeHTML string
var AnalyzeURL string
var AnalyzeClean bool
var AnalyzeXpath string
var AnalyzeFallbackToRaw bool
var AnalyzeReturnAnalyzedText bool
var AnalyzeLanguage string
var AnalyzeLimitTextCharacters int64

func getAnalyzeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "analyze",
		Short: "Analyze text",
		Long: "Analyzes text, HTML, or a public webpage for the following features:- Categories- Concepts- Emotion- Entities- Keywords- Metadata- Relations- Semantic roles- Sentiment- Syntax (Experimental).",
		Run: Analyze,
	}

	cmd.Flags().StringVarP(&AnalyzeFeatures, "features", "", "", "Specific features to analyze the document for.")
	cmd.Flags().StringVarP(&AnalyzeText, "text", "", "", "The plain text to analyze. One of the `text`, `html`, or `url` parameters is required.")
	cmd.Flags().StringVarP(&AnalyzeHTML, "html", "", "", "The HTML file to analyze. One of the `text`, `html`, or `url` parameters is required.")
	cmd.Flags().StringVarP(&AnalyzeURL, "url", "", "", "The webpage to analyze. One of the `text`, `html`, or `url` parameters is required.")
	cmd.Flags().BoolVarP(&AnalyzeClean, "clean", "", false, "Set this to `false` to disable webpage cleaning. To learn more about webpage cleaning, see the [Analyzing webpages](https://cloud.ibm.com/docs/services/natural-language-understanding?topic=natural-language-understanding-analyzing-webpages) documentation.")
	cmd.Flags().StringVarP(&AnalyzeXpath, "xpath", "", "", "An [XPath query](https://cloud.ibm.com/docs/services/natural-language-understanding?topic=natural-language-understanding-analyzing-webpages#xpath) to perform on `html` or `url` input. Results of the query will be appended to the cleaned webpage text before it is analyzed. To analyze only the results of the XPath query, set the `clean` parameter to `false`.")
	cmd.Flags().BoolVarP(&AnalyzeFallbackToRaw, "fallback_to_raw", "", false, "Whether to use raw HTML content if text cleaning fails.")
	cmd.Flags().BoolVarP(&AnalyzeReturnAnalyzedText, "return_analyzed_text", "", false, "Whether or not to return the analyzed text.")
	cmd.Flags().StringVarP(&AnalyzeLanguage, "language", "", "", "ISO 639-1 code that specifies the language of your text. This overrides automatic language detection. Language support differs depending on the features you include in your analysis. See [Language support](https://cloud.ibm.com/docs/services/natural-language-understanding?topic=natural-language-understanding-language-support) for more information.")
	cmd.Flags().Int64VarP(&AnalyzeLimitTextCharacters, "limit_text_characters", "", 0, "Sets the maximum number of characters that are processed by the service.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("features")
	cmd.MarkFlagRequired("version")

	return cmd
}

func Analyze(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	naturalLanguageUnderstanding, naturalLanguageUnderstandingErr := naturallanguageunderstandingv1.
		NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(naturalLanguageUnderstandingErr)

	optionsModel := naturallanguageunderstandingv1.AnalyzeOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "features" {
			var features naturallanguageunderstandingv1.Features
			decodeErr := json.Unmarshal([]byte(AnalyzeFeatures), &features);
			utils.HandleError(decodeErr)

			optionsModel.SetFeatures(&features)
		}
		if flag.Name == "text" {
			optionsModel.SetText(AnalyzeText)
		}
		if flag.Name == "html" {
			optionsModel.SetHTML(AnalyzeHTML)
		}
		if flag.Name == "url" {
			optionsModel.SetURL(AnalyzeURL)
		}
		if flag.Name == "clean" {
			optionsModel.SetClean(AnalyzeClean)
		}
		if flag.Name == "xpath" {
			optionsModel.SetXpath(AnalyzeXpath)
		}
		if flag.Name == "fallback_to_raw" {
			optionsModel.SetFallbackToRaw(AnalyzeFallbackToRaw)
		}
		if flag.Name == "return_analyzed_text" {
			optionsModel.SetReturnAnalyzedText(AnalyzeReturnAnalyzedText)
		}
		if flag.Name == "language" {
			optionsModel.SetLanguage(AnalyzeLanguage)
		}
		if flag.Name == "limit_text_characters" {
			optionsModel.SetLimitTextCharacters(AnalyzeLimitTextCharacters)
		}
	})

	result, _, responseErr := naturalLanguageUnderstanding.Analyze(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}


func getListModelsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-models",
		Short: "List models",
		Long: "Lists Watson Knowledge Studio [custom entities and relations models](https://cloud.ibm.com/docs/services/natural-language-understanding?topic=natural-language-understanding-customizing) that are deployed to your Natural Language Understanding service.",
		Run: ListModels,
	}

	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("version")

	return cmd
}

func ListModels(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	naturalLanguageUnderstanding, naturalLanguageUnderstandingErr := naturallanguageunderstandingv1.
		NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(naturalLanguageUnderstandingErr)

	optionsModel := naturallanguageunderstandingv1.ListModelsOptions{}

	result, _, responseErr := naturalLanguageUnderstanding.ListModels(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteModelModelID string

func getDeleteModelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-model",
		Short: "Delete model",
		Long: "Deletes a custom model.",
		Run: DeleteModel,
	}

	cmd.Flags().StringVarP(&DeleteModelModelID, "model_id", "", "", "Model ID of the model to delete.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("model_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteModel(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	naturalLanguageUnderstanding, naturalLanguageUnderstandingErr := naturallanguageunderstandingv1.
		NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(naturalLanguageUnderstandingErr)

	optionsModel := naturallanguageunderstandingv1.DeleteModelOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "model_id" {
			optionsModel.SetModelID(DeleteModelModelID)
		}
	})

	result, _, responseErr := naturalLanguageUnderstanding.DeleteModel(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}
