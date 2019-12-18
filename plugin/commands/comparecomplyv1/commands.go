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

package comparecomplyv1

import (
	"cli-watson-plugin/utils"
	"encoding/json"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/watson-developer-cloud/go-sdk/comparecomplyv1"
	"os"
	"time"
)

var ConvertToHTMLFile string
var ConvertToHTMLFileContentType string
var ConvertToHTMLModel string

func getConvertToHTMLCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "convert-to-html",
		Short: "Convert document to HTML",
		Long: "Converts a document to HTML.",
		Run: ConvertToHTML,
	}

	cmd.Flags().StringVarP(&ConvertToHTMLFile, "file", "", "", "The document to convert.")
	cmd.Flags().StringVarP(&ConvertToHTMLFileContentType, "file_content_type", "", "", "The content type of File.")
	cmd.Flags().StringVarP(&ConvertToHTMLModel, "model", "", "", "The analysis model to be used by the service. For the **Element classification** and **Compare two documents** methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults apply to the standalone methods as well as to the methods' use in batch-processing requests.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("file")
	cmd.MarkFlagRequired("version")

	return cmd
}

func ConvertToHTML(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	compareComply, compareComplyErr := comparecomplyv1.
		NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(compareComplyErr)

	optionsModel := comparecomplyv1.ConvertToHTMLOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "file" {
			file, fileErr := os.Open(ConvertToHTMLFile)
			utils.HandleError(fileErr)

			optionsModel.SetFile(file)
		}
		if flag.Name == "file_content_type" {
			optionsModel.SetFileContentType(ConvertToHTMLFileContentType)
		}
		if flag.Name == "model" {
			optionsModel.SetModel(ConvertToHTMLModel)
		}
	})

	result, _, responseErr := compareComply.ConvertToHTML(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ClassifyElementsFile string
var ClassifyElementsFileContentType string
var ClassifyElementsModel string

func getClassifyElementsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "classify-elements",
		Short: "Classify the elements of a document",
		Long: "Analyzes the structural and semantic elements of a document.",
		Run: ClassifyElements,
	}

	cmd.Flags().StringVarP(&ClassifyElementsFile, "file", "", "", "The document to classify.")
	cmd.Flags().StringVarP(&ClassifyElementsFileContentType, "file_content_type", "", "", "The content type of File.")
	cmd.Flags().StringVarP(&ClassifyElementsModel, "model", "", "", "The analysis model to be used by the service. For the **Element classification** and **Compare two documents** methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults apply to the standalone methods as well as to the methods' use in batch-processing requests.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("file")
	cmd.MarkFlagRequired("version")

	return cmd
}

func ClassifyElements(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	compareComply, compareComplyErr := comparecomplyv1.
		NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(compareComplyErr)

	optionsModel := comparecomplyv1.ClassifyElementsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "file" {
			file, fileErr := os.Open(ClassifyElementsFile)
			utils.HandleError(fileErr)

			optionsModel.SetFile(file)
		}
		if flag.Name == "file_content_type" {
			optionsModel.SetFileContentType(ClassifyElementsFileContentType)
		}
		if flag.Name == "model" {
			optionsModel.SetModel(ClassifyElementsModel)
		}
	})

	result, _, responseErr := compareComply.ClassifyElements(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ExtractTablesFile string
var ExtractTablesFileContentType string
var ExtractTablesModel string

func getExtractTablesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "extract-tables",
		Short: "Extract a document's tables",
		Long: "Analyzes the tables in a document.",
		Run: ExtractTables,
	}

	cmd.Flags().StringVarP(&ExtractTablesFile, "file", "", "", "The document on which to run table extraction.")
	cmd.Flags().StringVarP(&ExtractTablesFileContentType, "file_content_type", "", "", "The content type of File.")
	cmd.Flags().StringVarP(&ExtractTablesModel, "model", "", "", "The analysis model to be used by the service. For the **Element classification** and **Compare two documents** methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults apply to the standalone methods as well as to the methods' use in batch-processing requests.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("file")
	cmd.MarkFlagRequired("version")

	return cmd
}

func ExtractTables(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	compareComply, compareComplyErr := comparecomplyv1.
		NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(compareComplyErr)

	optionsModel := comparecomplyv1.ExtractTablesOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "file" {
			file, fileErr := os.Open(ExtractTablesFile)
			utils.HandleError(fileErr)

			optionsModel.SetFile(file)
		}
		if flag.Name == "file_content_type" {
			optionsModel.SetFileContentType(ExtractTablesFileContentType)
		}
		if flag.Name == "model" {
			optionsModel.SetModel(ExtractTablesModel)
		}
	})

	result, _, responseErr := compareComply.ExtractTables(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var CompareDocumentsFile1 string
var CompareDocumentsFile2 string
var CompareDocumentsFile1ContentType string
var CompareDocumentsFile2ContentType string
var CompareDocumentsFile1Label string
var CompareDocumentsFile2Label string
var CompareDocumentsModel string

func getCompareDocumentsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "compare-documents",
		Short: "Compare two documents",
		Long: "Compares two input documents. Documents must be in the same format.",
		Run: CompareDocuments,
	}

	cmd.Flags().StringVarP(&CompareDocumentsFile1, "file1", "", "", "The first document to compare.")
	cmd.Flags().StringVarP(&CompareDocumentsFile2, "file2", "", "", "The second document to compare.")
	cmd.Flags().StringVarP(&CompareDocumentsFile1ContentType, "file1_content_type", "", "", "The content type of File1.")
	cmd.Flags().StringVarP(&CompareDocumentsFile2ContentType, "file2_content_type", "", "", "The content type of File2.")
	cmd.Flags().StringVarP(&CompareDocumentsFile1Label, "file1_label", "", "", "A text label for the first document.")
	cmd.Flags().StringVarP(&CompareDocumentsFile2Label, "file2_label", "", "", "A text label for the second document.")
	cmd.Flags().StringVarP(&CompareDocumentsModel, "model", "", "", "The analysis model to be used by the service. For the **Element classification** and **Compare two documents** methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults apply to the standalone methods as well as to the methods' use in batch-processing requests.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("file1")
	cmd.MarkFlagRequired("file2")
	cmd.MarkFlagRequired("version")

	return cmd
}

func CompareDocuments(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	compareComply, compareComplyErr := comparecomplyv1.
		NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(compareComplyErr)

	optionsModel := comparecomplyv1.CompareDocumentsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "file1" {
			file1, fileErr := os.Open(CompareDocumentsFile1)
			utils.HandleError(fileErr)

			optionsModel.SetFile1(file1)
		}
		if flag.Name == "file2" {
			file2, fileErr := os.Open(CompareDocumentsFile2)
			utils.HandleError(fileErr)

			optionsModel.SetFile2(file2)
		}
		if flag.Name == "file1_content_type" {
			optionsModel.SetFile1ContentType(CompareDocumentsFile1ContentType)
		}
		if flag.Name == "file2_content_type" {
			optionsModel.SetFile2ContentType(CompareDocumentsFile2ContentType)
		}
		if flag.Name == "file1_label" {
			optionsModel.SetFile1Label(CompareDocumentsFile1Label)
		}
		if flag.Name == "file2_label" {
			optionsModel.SetFile2Label(CompareDocumentsFile2Label)
		}
		if flag.Name == "model" {
			optionsModel.SetModel(CompareDocumentsModel)
		}
	})

	result, _, responseErr := compareComply.CompareDocuments(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var AddFeedbackFeedbackData string
var AddFeedbackUserID string
var AddFeedbackComment string

func getAddFeedbackCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "add-feedback",
		Short: "Add feedback",
		Long: "Adds feedback in the form of _labels_ from a subject-matter expert (SME) to a governing document. **Important:** Feedback is not immediately incorporated into the training model, nor is it guaranteed to be incorporated at a later date. Instead, submitted feedback is used to suggest future updates to the training model.",
		Run: AddFeedback,
	}

	cmd.Flags().StringVarP(&AddFeedbackFeedbackData, "feedback_data", "", "", "Feedback data for submission.")
	cmd.Flags().StringVarP(&AddFeedbackUserID, "user_id", "", "", "An optional string identifying the user.")
	cmd.Flags().StringVarP(&AddFeedbackComment, "comment", "", "", "An optional comment on or description of the feedback.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("feedback_data")
	cmd.MarkFlagRequired("version")

	return cmd
}

func AddFeedback(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	compareComply, compareComplyErr := comparecomplyv1.
		NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(compareComplyErr)

	optionsModel := comparecomplyv1.AddFeedbackOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "feedback_data" {
			var feedback_data comparecomplyv1.FeedbackDataInput
			decodeErr := json.Unmarshal([]byte(AddFeedbackFeedbackData), &feedback_data);
			utils.HandleError(decodeErr)

			optionsModel.SetFeedbackData(&feedback_data)
		}
		if flag.Name == "user_id" {
			optionsModel.SetUserID(AddFeedbackUserID)
		}
		if flag.Name == "comment" {
			optionsModel.SetComment(AddFeedbackComment)
		}
	})

	result, _, responseErr := compareComply.AddFeedback(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ListFeedbackFeedbackType string
var ListFeedbackBefore string
var ListFeedbackAfter string
var ListFeedbackDocumentTitle string
var ListFeedbackModelID string
var ListFeedbackModelVersion string
var ListFeedbackCategoryRemoved string
var ListFeedbackCategoryAdded string
var ListFeedbackCategoryNotChanged string
var ListFeedbackTypeRemoved string
var ListFeedbackTypeAdded string
var ListFeedbackTypeNotChanged string
var ListFeedbackPageLimit int64
var ListFeedbackCursor string
var ListFeedbackSort string
var ListFeedbackIncludeTotal bool

func getListFeedbackCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-feedback",
		Short: "List the feedback in a document",
		Long: "Lists the feedback in a document.",
		Run: ListFeedback,
	}

	cmd.Flags().StringVarP(&ListFeedbackFeedbackType, "feedback_type", "", "", "An optional string that filters the output to include only feedback with the specified feedback type. The only permitted value is `element_classification`.")
	cmd.Flags().StringVarP(&ListFeedbackBefore, "before", "", "", "An optional string in the format `YYYY-MM-DD` that filters the output to include only feedback that was added before the specified date.")
	cmd.Flags().StringVarP(&ListFeedbackAfter, "after", "", "", "An optional string in the format `YYYY-MM-DD` that filters the output to include only feedback that was added after the specified date.")
	cmd.Flags().StringVarP(&ListFeedbackDocumentTitle, "document_title", "", "", "An optional string that filters the output to include only feedback from the document with the specified `document_title`.")
	cmd.Flags().StringVarP(&ListFeedbackModelID, "model_id", "", "", "An optional string that filters the output to include only feedback with the specified `model_id`. The only permitted value is `contracts`.")
	cmd.Flags().StringVarP(&ListFeedbackModelVersion, "model_version", "", "", "An optional string that filters the output to include only feedback with the specified `model_version`.")
	cmd.Flags().StringVarP(&ListFeedbackCategoryRemoved, "category_removed", "", "", "An optional string in the form of a comma-separated list of categories. If it is specified, the service filters the output to include only feedback that has at least one category from the list removed.")
	cmd.Flags().StringVarP(&ListFeedbackCategoryAdded, "category_added", "", "", "An optional string in the form of a comma-separated list of categories. If this is specified, the service filters the output to include only feedback that has at least one category from the list added.")
	cmd.Flags().StringVarP(&ListFeedbackCategoryNotChanged, "category_not_changed", "", "", "An optional string in the form of a comma-separated list of categories. If this is specified, the service filters the output to include only feedback that has at least one category from the list unchanged.")
	cmd.Flags().StringVarP(&ListFeedbackTypeRemoved, "type_removed", "", "", "An optional string of comma-separated `nature`:`party` pairs. If this is specified, the service filters the output to include only feedback that has at least one `nature`:`party` pair from the list removed.")
	cmd.Flags().StringVarP(&ListFeedbackTypeAdded, "type_added", "", "", "An optional string of comma-separated `nature`:`party` pairs. If this is specified, the service filters the output to include only feedback that has at least one `nature`:`party` pair from the list removed.")
	cmd.Flags().StringVarP(&ListFeedbackTypeNotChanged, "type_not_changed", "", "", "An optional string of comma-separated `nature`:`party` pairs. If this is specified, the service filters the output to include only feedback that has at least one `nature`:`party` pair from the list unchanged.")
	cmd.Flags().Int64VarP(&ListFeedbackPageLimit, "page_limit", "", 0, "An optional integer specifying the number of documents that you want the service to return.")
	cmd.Flags().StringVarP(&ListFeedbackCursor, "cursor", "", "", "An optional string that returns the set of documents after the previous set. Use this parameter with the `page_limit` parameter.")
	cmd.Flags().StringVarP(&ListFeedbackSort, "sort", "", "", "An optional comma-separated list of fields in the document to sort on. You can optionally specify the sort direction by prefixing the value of the field with `-` for descending order or `+` for ascending order (the default). Currently permitted sorting fields are `created`, `user_id`, and `document_title`.")
	cmd.Flags().BoolVarP(&ListFeedbackIncludeTotal, "include_total", "", false, "An optional boolean value. If specified as `true`, the `pagination` object in the output includes a value called `total` that gives the total count of feedback created.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("version")

	return cmd
}

func ListFeedback(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	compareComply, compareComplyErr := comparecomplyv1.
		NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(compareComplyErr)

	optionsModel := comparecomplyv1.ListFeedbackOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "feedback_type" {
			optionsModel.SetFeedbackType(ListFeedbackFeedbackType)
		}
		if flag.Name == "before" {
			beforeTime, timeParseErr := time.Parse(time.RFC3339, ListFeedbackBefore)
			utils.HandleError(timeParseErr)

			before := strfmt.Date(beforeTime)

			optionsModel.SetBefore(&before)
		}
		if flag.Name == "after" {
			afterTime, timeParseErr := time.Parse(time.RFC3339, ListFeedbackAfter)
			utils.HandleError(timeParseErr)

			after := strfmt.Date(afterTime)

			optionsModel.SetAfter(&after)
		}
		if flag.Name == "document_title" {
			optionsModel.SetDocumentTitle(ListFeedbackDocumentTitle)
		}
		if flag.Name == "model_id" {
			optionsModel.SetModelID(ListFeedbackModelID)
		}
		if flag.Name == "model_version" {
			optionsModel.SetModelVersion(ListFeedbackModelVersion)
		}
		if flag.Name == "category_removed" {
			optionsModel.SetCategoryRemoved(ListFeedbackCategoryRemoved)
		}
		if flag.Name == "category_added" {
			optionsModel.SetCategoryAdded(ListFeedbackCategoryAdded)
		}
		if flag.Name == "category_not_changed" {
			optionsModel.SetCategoryNotChanged(ListFeedbackCategoryNotChanged)
		}
		if flag.Name == "type_removed" {
			optionsModel.SetTypeRemoved(ListFeedbackTypeRemoved)
		}
		if flag.Name == "type_added" {
			optionsModel.SetTypeAdded(ListFeedbackTypeAdded)
		}
		if flag.Name == "type_not_changed" {
			optionsModel.SetTypeNotChanged(ListFeedbackTypeNotChanged)
		}
		if flag.Name == "page_limit" {
			optionsModel.SetPageLimit(ListFeedbackPageLimit)
		}
		if flag.Name == "cursor" {
			optionsModel.SetCursor(ListFeedbackCursor)
		}
		if flag.Name == "sort" {
			optionsModel.SetSort(ListFeedbackSort)
		}
		if flag.Name == "include_total" {
			optionsModel.SetIncludeTotal(ListFeedbackIncludeTotal)
		}
	})

	result, _, responseErr := compareComply.ListFeedback(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetFeedbackFeedbackID string
var GetFeedbackModel string

func getGetFeedbackCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-feedback",
		Short: "Get a specified feedback entry",
		Long: "Gets a feedback entry with a specified `feedback_id`.",
		Run: GetFeedback,
	}

	cmd.Flags().StringVarP(&GetFeedbackFeedbackID, "feedback_id", "", "", "A string that specifies the feedback entry to be included in the output.")
	cmd.Flags().StringVarP(&GetFeedbackModel, "model", "", "", "The analysis model to be used by the service. For the **Element classification** and **Compare two documents** methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults apply to the standalone methods as well as to the methods' use in batch-processing requests.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("feedback_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetFeedback(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	compareComply, compareComplyErr := comparecomplyv1.
		NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(compareComplyErr)

	optionsModel := comparecomplyv1.GetFeedbackOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "feedback_id" {
			optionsModel.SetFeedbackID(GetFeedbackFeedbackID)
		}
		if flag.Name == "model" {
			optionsModel.SetModel(GetFeedbackModel)
		}
	})

	result, _, responseErr := compareComply.GetFeedback(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteFeedbackFeedbackID string
var DeleteFeedbackModel string

func getDeleteFeedbackCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-feedback",
		Short: "Delete a specified feedback entry",
		Long: "Deletes a feedback entry with a specified `feedback_id`.",
		Run: DeleteFeedback,
	}

	cmd.Flags().StringVarP(&DeleteFeedbackFeedbackID, "feedback_id", "", "", "A string that specifies the feedback entry to be deleted from the document.")
	cmd.Flags().StringVarP(&DeleteFeedbackModel, "model", "", "", "The analysis model to be used by the service. For the **Element classification** and **Compare two documents** methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults apply to the standalone methods as well as to the methods' use in batch-processing requests.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("feedback_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteFeedback(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	compareComply, compareComplyErr := comparecomplyv1.
		NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(compareComplyErr)

	optionsModel := comparecomplyv1.DeleteFeedbackOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "feedback_id" {
			optionsModel.SetFeedbackID(DeleteFeedbackFeedbackID)
		}
		if flag.Name == "model" {
			optionsModel.SetModel(DeleteFeedbackModel)
		}
	})

	result, _, responseErr := compareComply.DeleteFeedback(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var CreateBatchFunction string
var CreateBatchInputCredentialsFile string
var CreateBatchInputBucketLocation string
var CreateBatchInputBucketName string
var CreateBatchOutputCredentialsFile string
var CreateBatchOutputBucketLocation string
var CreateBatchOutputBucketName string
var CreateBatchModel string

func getCreateBatchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-batch",
		Short: "Submit a batch-processing request",
		Long: "Run Compare and Comply methods over a collection of input documents.**Important:** Batch processing requires the use of the [IBM Cloud Object Storage service](https://cloud.ibm.com/docs/services/cloud-object-storage?topic=cloud-object-storage-about#about-ibm-cloud-object-storage). The use of IBM Cloud Object Storage with Compare and Comply is discussed at [Using batch processing](https://cloud.ibm.com/docs/services/compare-comply?topic=compare-comply-batching#before-you-batch).",
		Run: CreateBatch,
	}

	cmd.Flags().StringVarP(&CreateBatchFunction, "function", "", "", "The Compare and Comply method to run across the submitted input documents.")
	cmd.Flags().StringVarP(&CreateBatchInputCredentialsFile, "input_credentials_file", "", "", "A JSON file containing the input Cloud Object Storage credentials. At a minimum, the credentials must enable `READ` permissions on the bucket defined by the `input_bucket_name` parameter.")
	cmd.Flags().StringVarP(&CreateBatchInputBucketLocation, "input_bucket_location", "", "", "The geographical location of the Cloud Object Storage input bucket as listed on the **Endpoint** tab of your Cloud Object Storage instance; for example, `us-geo`, `eu-geo`, or `ap-geo`.")
	cmd.Flags().StringVarP(&CreateBatchInputBucketName, "input_bucket_name", "", "", "The name of the Cloud Object Storage input bucket.")
	cmd.Flags().StringVarP(&CreateBatchOutputCredentialsFile, "output_credentials_file", "", "", "A JSON file that lists the Cloud Object Storage output credentials. At a minimum, the credentials must enable `READ` and `WRITE` permissions on the bucket defined by the `output_bucket_name` parameter.")
	cmd.Flags().StringVarP(&CreateBatchOutputBucketLocation, "output_bucket_location", "", "", "The geographical location of the Cloud Object Storage output bucket as listed on the **Endpoint** tab of your Cloud Object Storage instance; for example, `us-geo`, `eu-geo`, or `ap-geo`.")
	cmd.Flags().StringVarP(&CreateBatchOutputBucketName, "output_bucket_name", "", "", "The name of the Cloud Object Storage output bucket.")
	cmd.Flags().StringVarP(&CreateBatchModel, "model", "", "", "The analysis model to be used by the service. For the **Element classification** and **Compare two documents** methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults apply to the standalone methods as well as to the methods' use in batch-processing requests.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("function")
	cmd.MarkFlagRequired("input_credentials_file")
	cmd.MarkFlagRequired("input_bucket_location")
	cmd.MarkFlagRequired("input_bucket_name")
	cmd.MarkFlagRequired("output_credentials_file")
	cmd.MarkFlagRequired("output_bucket_location")
	cmd.MarkFlagRequired("output_bucket_name")
	cmd.MarkFlagRequired("version")

	return cmd
}

func CreateBatch(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	compareComply, compareComplyErr := comparecomplyv1.
		NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(compareComplyErr)

	optionsModel := comparecomplyv1.CreateBatchOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "function" {
			optionsModel.SetFunction(CreateBatchFunction)
		}
		if flag.Name == "input_credentials_file" {
			input_credentials_file, fileErr := os.Open(CreateBatchInputCredentialsFile)
			utils.HandleError(fileErr)

			optionsModel.SetInputCredentialsFile(input_credentials_file)
		}
		if flag.Name == "input_bucket_location" {
			optionsModel.SetInputBucketLocation(CreateBatchInputBucketLocation)
		}
		if flag.Name == "input_bucket_name" {
			optionsModel.SetInputBucketName(CreateBatchInputBucketName)
		}
		if flag.Name == "output_credentials_file" {
			output_credentials_file, fileErr := os.Open(CreateBatchOutputCredentialsFile)
			utils.HandleError(fileErr)

			optionsModel.SetOutputCredentialsFile(output_credentials_file)
		}
		if flag.Name == "output_bucket_location" {
			optionsModel.SetOutputBucketLocation(CreateBatchOutputBucketLocation)
		}
		if flag.Name == "output_bucket_name" {
			optionsModel.SetOutputBucketName(CreateBatchOutputBucketName)
		}
		if flag.Name == "model" {
			optionsModel.SetModel(CreateBatchModel)
		}
	})

	result, _, responseErr := compareComply.CreateBatch(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}


func getListBatchesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-batches",
		Short: "List submitted batch-processing jobs",
		Long: "Lists batch-processing jobs submitted by users.",
		Run: ListBatches,
	}

	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("version")

	return cmd
}

func ListBatches(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	compareComply, compareComplyErr := comparecomplyv1.
		NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(compareComplyErr)

	optionsModel := comparecomplyv1.ListBatchesOptions{}

	result, _, responseErr := compareComply.ListBatches(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetBatchBatchID string

func getGetBatchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-batch",
		Short: "Get information about a specific batch-processing job",
		Long: "Gets information about a batch-processing job with a specified ID.",
		Run: GetBatch,
	}

	cmd.Flags().StringVarP(&GetBatchBatchID, "batch_id", "", "", "The ID of the batch-processing job whose information you want to retrieve.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("batch_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetBatch(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	compareComply, compareComplyErr := comparecomplyv1.
		NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(compareComplyErr)

	optionsModel := comparecomplyv1.GetBatchOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "batch_id" {
			optionsModel.SetBatchID(GetBatchBatchID)
		}
	})

	result, _, responseErr := compareComply.GetBatch(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var UpdateBatchBatchID string
var UpdateBatchAction string
var UpdateBatchModel string

func getUpdateBatchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-batch",
		Short: "Update a pending or active batch-processing job",
		Long: "Updates a pending or active batch-processing job. You can rescan the input bucket to check for new documents or cancel a job.",
		Run: UpdateBatch,
	}

	cmd.Flags().StringVarP(&UpdateBatchBatchID, "batch_id", "", "", "The ID of the batch-processing job you want to update.")
	cmd.Flags().StringVarP(&UpdateBatchAction, "action", "", "", "The action you want to perform on the specified batch-processing job.")
	cmd.Flags().StringVarP(&UpdateBatchModel, "model", "", "", "The analysis model to be used by the service. For the **Element classification** and **Compare two documents** methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults apply to the standalone methods as well as to the methods' use in batch-processing requests.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("batch_id")
	cmd.MarkFlagRequired("action")
	cmd.MarkFlagRequired("version")

	return cmd
}

func UpdateBatch(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	compareComply, compareComplyErr := comparecomplyv1.
		NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(compareComplyErr)

	optionsModel := comparecomplyv1.UpdateBatchOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "batch_id" {
			optionsModel.SetBatchID(UpdateBatchBatchID)
		}
		if flag.Name == "action" {
			optionsModel.SetAction(UpdateBatchAction)
		}
		if flag.Name == "model" {
			optionsModel.SetModel(UpdateBatchModel)
		}
	})

	result, _, responseErr := compareComply.UpdateBatch(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}
