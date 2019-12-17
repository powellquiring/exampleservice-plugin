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

package personalityinsightsv3

import (
	"cli-watson-plugin/utils"
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/watson-developer-cloud/go-sdk/personalityinsightsv3"
	"io"
	"os"
)

var ProfileContent string
var ProfileBody string
var ProfileContentType string
var ProfileContentLanguage string
var ProfileAcceptLanguage string
var ProfileRawScores bool
var ProfileCsvHeaders bool
var ProfileConsumptionPreferences bool

func getProfileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "profile",
		Short: "Get profile",
		Long: "Generates a personality profile for the author of the input text. The service accepts a maximum of 20 MB of input content, but it requires much less text to produce an accurate profile. The service can analyze text in Arabic, English, Japanese, Korean, or Spanish. It can return its results in a variety of languages. **See also:*** [Requesting a profile](https://cloud.ibm.com/docs/services/personality-insights?topic=personality-insights-input#input)* [Providing sufficient input](https://cloud.ibm.com/docs/services/personality-insights?topic=personality-insights-input#sufficient) ### Content types You can provide input content as plain text (`text/plain`), HTML (`text/html`), or JSON (`application/json`) by specifying the **Content-Type** parameter. The default is `text/plain`.* Per the JSON specification, the default character encoding for JSON content is effectively always UTF-8.* Per the HTTP specification, the default encoding for plain text and HTML is ISO-8859-1 (effectively, the ASCII character set). When specifying a content type of plain text or HTML, include the `charset` parameter to indicate the character encoding of the input text; for example, `Content-Type: text/plain;charset=utf-8`. **See also:** [Specifying request and response formats](https://cloud.ibm.com/docs/services/personality-insights?topic=personality-insights-input#formats) ### Accept types You must request a response as JSON (`application/json`) or comma-separated values (`text/csv`) by specifying the **Accept** parameter. CSV output includes a fixed number of columns. Set the **csv_headers** parameter to `true` to request optional column headers for CSV output. **See also:*** [Understanding a JSON profile](https://cloud.ibm.com/docs/services/personality-insights?topic=personality-insights-output#output)* [Understanding a CSV profile](https://cloud.ibm.com/docs/services/personality-insights?topic=personality-insights-outputCSV#outputCSV).",
		Run: Profile,
	}

	cmd.Flags().StringVarP(&ProfileContent, "content", "", "", "A maximum of 20 MB of content to analyze, though the service requires much less text; for more information, see [Providing sufficient input](https://cloud.ibm.com/docs/services/personality-insights?topic=personality-insights-input#sufficient). For JSON input, provide an object of type `Content`.")
	cmd.Flags().StringVarP(&ProfileBody, "body", "", "", "A maximum of 20 MB of content to analyze, though the service requires much less text; for more information, see [Providing sufficient input](https://cloud.ibm.com/docs/services/personality-insights?topic=personality-insights-input#sufficient). For JSON input, provide an object of type `Content`.")
	cmd.Flags().StringVarP(&ProfileContentType, "content_type", "", "", "The type of the input. For more information, see **Content types** in the method description.")
	cmd.Flags().StringVarP(&ProfileContentLanguage, "content_language", "", "", "The language of the input text for the request: Arabic, English, Japanese, Korean, or Spanish. Regional variants are treated as their parent language; for example, `en-US` is interpreted as `en`. The effect of the **Content-Language** parameter depends on the **Content-Type** parameter. When **Content-Type** is `text/plain` or `text/html`, **Content-Language** is the only way to specify the language. When **Content-Type** is `application/json`, **Content-Language** overrides a language specified with the `language` parameter of a `ContentItem` object, and content items that specify a different language are ignored; omit this parameter to base the language on the specification of the content items. You can specify any combination of languages for **Content-Language** and **Accept-Language**.")
	cmd.Flags().StringVarP(&ProfileAcceptLanguage, "accept_language", "", "", "The desired language of the response. For two-character arguments, regional variants are treated as their parent language; for example, `en-US` is interpreted as `en`. You can specify any combination of languages for the input and response content.")
	cmd.Flags().BoolVarP(&ProfileRawScores, "raw_scores", "", false, "Indicates whether a raw score in addition to a normalized percentile is returned for each characteristic; raw scores are not compared with a sample population. By default, only normalized percentiles are returned.")
	cmd.Flags().BoolVarP(&ProfileCsvHeaders, "csv_headers", "", false, "Indicates whether column labels are returned with a CSV response. By default, no column labels are returned. Applies only when the response type is CSV (`text/csv`).")
	cmd.Flags().BoolVarP(&ProfileConsumptionPreferences, "consumption_preferences", "", false, "Indicates whether consumption preferences are returned with the results. By default, no consumption preferences are returned.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("version")

	return cmd
}

func Profile(cmd *cobra.Command, args []string) {
	personalityInsights, personalityInsightsErr := personalityinsightsv3.
		NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(personalityInsightsErr)

	optionsModel := personalityinsightsv3.ProfileOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "content" {
			var content personalityinsightsv3.Content
			decodeErr := json.Unmarshal([]byte(ProfileContent), &content);
			utils.HandleError(decodeErr)

			optionsModel.SetContent(&content)
		}
		if flag.Name == "body" {
			optionsModel.SetBody(ProfileBody)
		}
		if flag.Name == "content_type" {
			optionsModel.SetContentType(ProfileContentType)
		}
		if flag.Name == "content_language" {
			optionsModel.SetContentLanguage(ProfileContentLanguage)
		}
		if flag.Name == "accept_language" {
			optionsModel.SetAcceptLanguage(ProfileAcceptLanguage)
		}
		if flag.Name == "raw_scores" {
			optionsModel.SetRawScores(ProfileRawScores)
		}
		if flag.Name == "csv_headers" {
			optionsModel.SetCsvHeaders(ProfileCsvHeaders)
		}
		if flag.Name == "consumption_preferences" {
			optionsModel.SetConsumptionPreferences(ProfileConsumptionPreferences)
		}
	})

	result, _, responseErr := personalityInsights.Profile(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ProfileAsCsvContent string
var ProfileAsCsvBody string
var ProfileAsCsvContentType string
var ProfileAsCsvContentLanguage string
var ProfileAsCsvAcceptLanguage string
var ProfileAsCsvRawScores bool
var ProfileAsCsvCsvHeaders bool
var ProfileAsCsvConsumptionPreferences bool

func getProfileAsCsvCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "profile-as-csv",
		Short: "Get profile as csv",
		Long: "Generates a personality profile for the author of the input text. The service accepts a maximum of 20 MB of input content, but it requires much less text to produce an accurate profile. The service can analyze text in Arabic, English, Japanese, Korean, or Spanish. It can return its results in a variety of languages. **See also:*** [Requesting a profile](https://cloud.ibm.com/docs/services/personality-insights?topic=personality-insights-input#input)* [Providing sufficient input](https://cloud.ibm.com/docs/services/personality-insights?topic=personality-insights-input#sufficient) ### Content types You can provide input content as plain text (`text/plain`), HTML (`text/html`), or JSON (`application/json`) by specifying the **Content-Type** parameter. The default is `text/plain`.* Per the JSON specification, the default character encoding for JSON content is effectively always UTF-8.* Per the HTTP specification, the default encoding for plain text and HTML is ISO-8859-1 (effectively, the ASCII character set). When specifying a content type of plain text or HTML, include the `charset` parameter to indicate the character encoding of the input text; for example, `Content-Type: text/plain;charset=utf-8`. **See also:** [Specifying request and response formats](https://cloud.ibm.com/docs/services/personality-insights?topic=personality-insights-input#formats) ### Accept types You must request a response as JSON (`application/json`) or comma-separated values (`text/csv`) by specifying the **Accept** parameter. CSV output includes a fixed number of columns. Set the **csv_headers** parameter to `true` to request optional column headers for CSV output. **See also:*** [Understanding a JSON profile](https://cloud.ibm.com/docs/services/personality-insights?topic=personality-insights-output#output)* [Understanding a CSV profile](https://cloud.ibm.com/docs/services/personality-insights?topic=personality-insights-outputCSV#outputCSV).",
		Run: ProfileAsCsv,
	}

	cmd.Flags().StringVarP(&ProfileAsCsvContent, "content", "", "", "A maximum of 20 MB of content to analyze, though the service requires much less text; for more information, see [Providing sufficient input](https://cloud.ibm.com/docs/services/personality-insights?topic=personality-insights-input#sufficient). For JSON input, provide an object of type `Content`.")
	cmd.Flags().StringVarP(&ProfileAsCsvBody, "body", "", "", "A maximum of 20 MB of content to analyze, though the service requires much less text; for more information, see [Providing sufficient input](https://cloud.ibm.com/docs/services/personality-insights?topic=personality-insights-input#sufficient). For JSON input, provide an object of type `Content`.")
	cmd.Flags().StringVarP(&ProfileAsCsvContentType, "content_type", "", "", "The type of the input. For more information, see **Content types** in the method description.")
	cmd.Flags().StringVarP(&ProfileAsCsvContentLanguage, "content_language", "", "", "The language of the input text for the request: Arabic, English, Japanese, Korean, or Spanish. Regional variants are treated as their parent language; for example, `en-US` is interpreted as `en`. The effect of the **Content-Language** parameter depends on the **Content-Type** parameter. When **Content-Type** is `text/plain` or `text/html`, **Content-Language** is the only way to specify the language. When **Content-Type** is `application/json`, **Content-Language** overrides a language specified with the `language` parameter of a `ContentItem` object, and content items that specify a different language are ignored; omit this parameter to base the language on the specification of the content items. You can specify any combination of languages for **Content-Language** and **Accept-Language**.")
	cmd.Flags().StringVarP(&ProfileAsCsvAcceptLanguage, "accept_language", "", "", "The desired language of the response. For two-character arguments, regional variants are treated as their parent language; for example, `en-US` is interpreted as `en`. You can specify any combination of languages for the input and response content.")
	cmd.Flags().BoolVarP(&ProfileAsCsvRawScores, "raw_scores", "", false, "Indicates whether a raw score in addition to a normalized percentile is returned for each characteristic; raw scores are not compared with a sample population. By default, only normalized percentiles are returned.")
	cmd.Flags().BoolVarP(&ProfileAsCsvCsvHeaders, "csv_headers", "", false, "Indicates whether column labels are returned with a CSV response. By default, no column labels are returned. Applies only when the response type is CSV (`text/csv`).")
	cmd.Flags().BoolVarP(&ProfileAsCsvConsumptionPreferences, "consumption_preferences", "", false, "Indicates whether consumption preferences are returned with the results. By default, no consumption preferences are returned.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")
	cmd.Flags().StringVarP(&OutputFilename, "output_file", "", "", "Filename/path to write the resulting output to.")

	cmd.MarkFlagRequired("version")
	cmd.MarkFlagRequired("output_file")

	return cmd
}

func ProfileAsCsv(cmd *cobra.Command, args []string) {
	personalityInsights, personalityInsightsErr := personalityinsightsv3.
		NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(personalityInsightsErr)

	optionsModel := personalityinsightsv3.ProfileOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "content" {
			var content personalityinsightsv3.Content
			decodeErr := json.Unmarshal([]byte(ProfileAsCsvContent), &content);
			utils.HandleError(decodeErr)

			optionsModel.SetContent(&content)
		}
		if flag.Name == "body" {
			optionsModel.SetBody(ProfileAsCsvBody)
		}
		if flag.Name == "content_type" {
			optionsModel.SetContentType(ProfileAsCsvContentType)
		}
		if flag.Name == "content_language" {
			optionsModel.SetContentLanguage(ProfileAsCsvContentLanguage)
		}
		if flag.Name == "accept_language" {
			optionsModel.SetAcceptLanguage(ProfileAsCsvAcceptLanguage)
		}
		if flag.Name == "raw_scores" {
			optionsModel.SetRawScores(ProfileAsCsvRawScores)
		}
		if flag.Name == "csv_headers" {
			optionsModel.SetCsvHeaders(ProfileAsCsvCsvHeaders)
		}
		if flag.Name == "consumption_preferences" {
			optionsModel.SetConsumptionPreferences(ProfileAsCsvConsumptionPreferences)
		}
	})

	result, _, responseErr := personalityInsights.ProfileAsCsv(&optionsModel)
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
