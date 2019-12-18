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
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/watson-developer-cloud/go-sdk/languagetranslatorv3"
	"io"
	"os"
)

var TranslateText []string
var TranslateModelID string
var TranslateSource string
var TranslateTarget string

func getTranslateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "translate",
		Short: "Translate",
		Long: "Translates the input text from the source language to the target language.",
		Run: Translate,
	}

	cmd.Flags().StringSliceVarP(&TranslateText, "text", "", nil, "Input text in UTF-8 encoding. Multiple entries will result in multiple translations in the response.")
	cmd.Flags().StringVarP(&TranslateModelID, "model_id", "", "", "A globally unique string that identifies the underlying model that is used for translation.")
	cmd.Flags().StringVarP(&TranslateSource, "source", "", "", "Translation source language code.")
	cmd.Flags().StringVarP(&TranslateTarget, "target", "", "", "Translation target language code.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("text")
	cmd.MarkFlagRequired("version")

	return cmd
}

func Translate(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	languageTranslator, languageTranslatorErr := languagetranslatorv3.
		NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(languageTranslatorErr)

	optionsModel := languagetranslatorv3.TranslateOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "text" {
			optionsModel.SetText(TranslateText)
		}
		if flag.Name == "model_id" {
			optionsModel.SetModelID(TranslateModelID)
		}
		if flag.Name == "source" {
			optionsModel.SetSource(TranslateSource)
		}
		if flag.Name == "target" {
			optionsModel.SetTarget(TranslateTarget)
		}
	})

	result, _, responseErr := languageTranslator.Translate(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}


func getListIdentifiableLanguagesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-identifiable-languages",
		Short: "List identifiable languages",
		Long: "Lists the languages that the service can identify. Returns the language code (for example, `en` for English or `es` for Spanish) and name of each language.",
		Run: ListIdentifiableLanguages,
	}

	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("version")

	return cmd
}

func ListIdentifiableLanguages(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	languageTranslator, languageTranslatorErr := languagetranslatorv3.
		NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(languageTranslatorErr)

	optionsModel := languagetranslatorv3.ListIdentifiableLanguagesOptions{}

	result, _, responseErr := languageTranslator.ListIdentifiableLanguages(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var IdentifyText string

func getIdentifyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "identify",
		Short: "Identify language",
		Long: "Identifies the language of the input text.",
		Run: Identify,
	}

	cmd.Flags().StringVarP(&IdentifyText, "text", "", "", "Input text in UTF-8 format.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("text")
	cmd.MarkFlagRequired("version")

	return cmd
}

func Identify(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	languageTranslator, languageTranslatorErr := languagetranslatorv3.
		NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(languageTranslatorErr)

	optionsModel := languagetranslatorv3.IdentifyOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "text" {
			optionsModel.SetText(IdentifyText)
		}
	})

	result, _, responseErr := languageTranslator.Identify(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ListModelsSource string
var ListModelsTarget string
var ListModelsDefault bool

func getListModelsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-models",
		Short: "List models",
		Long: "Lists available translation models.",
		Run: ListModels,
	}

	cmd.Flags().StringVarP(&ListModelsSource, "source", "", "", "Specify a language code to filter results by source language.")
	cmd.Flags().StringVarP(&ListModelsTarget, "target", "", "", "Specify a language code to filter results by target language.")
	cmd.Flags().BoolVarP(&ListModelsDefault, "default", "", false, "If the default parameter isn't specified, the service will return all models (default and non-default) for each language pair. To return only default models, set this to `true`. To return only non-default models, set this to `false`. There is exactly one default model per language pair, the IBM provided base model.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("version")

	return cmd
}

func ListModels(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	languageTranslator, languageTranslatorErr := languagetranslatorv3.
		NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(languageTranslatorErr)

	optionsModel := languagetranslatorv3.ListModelsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "source" {
			optionsModel.SetSource(ListModelsSource)
		}
		if flag.Name == "target" {
			optionsModel.SetTarget(ListModelsTarget)
		}
		if flag.Name == "default" {
			optionsModel.SetDefault(ListModelsDefault)
		}
	})

	result, _, responseErr := languageTranslator.ListModels(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var CreateModelBaseModelID string
var CreateModelForcedGlossary string
var CreateModelParallelCorpus string
var CreateModelName string

func getCreateModelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-model",
		Short: "Create model",
		Long: "Uploads Translation Memory eXchange (TMX) files to customize a translation model.You can either customize a model with a forced glossary or with a corpus that contains parallel sentences. To create a model that is customized with a parallel corpus <b>and</b> a forced glossary, proceed in two steps: customize with a parallel corpus first and then customize the resulting model with a glossary. Depending on the type of customization and the size of the uploaded corpora, training can range from minutes for a glossary to several hours for a large parallel corpus. You can upload a single forced glossary file and this file must be less than <b>10 MB</b>. You can upload multiple parallel corpora tmx files. The cumulative file size of all uploaded files is limited to <b>250 MB</b>. To successfully train with a parallel corpus you must have at least <b>5,000 parallel sentences</b> in your corpus.You can have a <b>maximum of 10 custom models per language pair</b>.",
		Run: CreateModel,
	}

	cmd.Flags().StringVarP(&CreateModelBaseModelID, "base_model_id", "", "", "The model ID of the model to use as the base for customization. To see available models, use the `List models` method. Usually all IBM provided models are customizable. In addition, all your models that have been created via parallel corpus customization, can be further customized with a forced glossary.")
	cmd.Flags().StringVarP(&CreateModelForcedGlossary, "forced_glossary", "", "", "A TMX file with your customizations. The customizations in the file completely overwrite the domain translaton data, including high frequency or high confidence phrase translations. You can upload only one glossary with a file size less than 10 MB per call. A forced glossary should contain single words or short phrases.")
	cmd.Flags().StringVarP(&CreateModelParallelCorpus, "parallel_corpus", "", "", "A TMX file with parallel sentences for source and target language. You can upload multiple parallel_corpus files in one request. All uploaded parallel_corpus files combined, your parallel corpus must contain at least 5,000 parallel sentences to train successfully.")
	cmd.Flags().StringVarP(&CreateModelName, "name", "", "", "An optional model name that you can use to identify the model. Valid characters are letters, numbers, dashes, underscores, spaces and apostrophes. The maximum length is 32 characters.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("base_model_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func CreateModel(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	languageTranslator, languageTranslatorErr := languagetranslatorv3.
		NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(languageTranslatorErr)

	optionsModel := languagetranslatorv3.CreateModelOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "base_model_id" {
			optionsModel.SetBaseModelID(CreateModelBaseModelID)
		}
		if flag.Name == "forced_glossary" {
			forced_glossary, fileErr := os.Open(CreateModelForcedGlossary)
			utils.HandleError(fileErr)

			optionsModel.SetForcedGlossary(forced_glossary)
		}
		if flag.Name == "parallel_corpus" {
			parallel_corpus, fileErr := os.Open(CreateModelParallelCorpus)
			utils.HandleError(fileErr)

			optionsModel.SetParallelCorpus(parallel_corpus)
		}
		if flag.Name == "name" {
			optionsModel.SetName(CreateModelName)
		}
	})

	result, _, responseErr := languageTranslator.CreateModel(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteModelModelID string

func getDeleteModelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-model",
		Short: "Delete model",
		Long: "Deletes a custom translation model.",
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
	languageTranslator, languageTranslatorErr := languagetranslatorv3.
		NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(languageTranslatorErr)

	optionsModel := languagetranslatorv3.DeleteModelOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "model_id" {
			optionsModel.SetModelID(DeleteModelModelID)
		}
	})

	result, _, responseErr := languageTranslator.DeleteModel(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetModelModelID string

func getGetModelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-model",
		Short: "Get model details",
		Long: "Gets information about a translation model, including training status for custom models. Use this API call to poll the status of your customization request. A successfully completed training will have a status of `available`.",
		Run: GetModel,
	}

	cmd.Flags().StringVarP(&GetModelModelID, "model_id", "", "", "Model ID of the model to get.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("model_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetModel(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	languageTranslator, languageTranslatorErr := languagetranslatorv3.
		NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(languageTranslatorErr)

	optionsModel := languagetranslatorv3.GetModelOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "model_id" {
			optionsModel.SetModelID(GetModelModelID)
		}
	})

	result, _, responseErr := languageTranslator.GetModel(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}


func getListDocumentsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-documents",
		Short: "List documents",
		Long: "Lists documents that have been submitted for translation.",
		Run: ListDocuments,
	}

	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("version")

	return cmd
}

func ListDocuments(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	languageTranslator, languageTranslatorErr := languagetranslatorv3.
		NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(languageTranslatorErr)

	optionsModel := languagetranslatorv3.ListDocumentsOptions{}

	result, _, responseErr := languageTranslator.ListDocuments(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var TranslateDocumentFile string
var TranslateDocumentFilename string
var TranslateDocumentFileContentType string
var TranslateDocumentModelID string
var TranslateDocumentSource string
var TranslateDocumentTarget string
var TranslateDocumentDocumentID string

func getTranslateDocumentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "translate-document",
		Short: "Translate document",
		Long: "Submit a document for translation. You can submit the document contents in the `file` parameter, or you can reference a previously submitted document by document ID.",
		Run: TranslateDocument,
	}

	cmd.Flags().StringVarP(&TranslateDocumentFile, "file", "", "", "The source file to translate.[Supported file types](https://cloud.ibm.com/docs/services/language-translator?topic=language-translator-document-translator-tutorial#supported-file-formats)Maximum file size: **20 MB**.")
	cmd.Flags().StringVarP(&TranslateDocumentFilename, "filename", "", "", "The filename for File.")
	cmd.Flags().StringVarP(&TranslateDocumentFileContentType, "file_content_type", "", "", "The content type of File.")
	cmd.Flags().StringVarP(&TranslateDocumentModelID, "model_id", "", "", "The model to use for translation. `model_id` or both `source` and `target` are required.")
	cmd.Flags().StringVarP(&TranslateDocumentSource, "source", "", "", "Language code that specifies the language of the source document.")
	cmd.Flags().StringVarP(&TranslateDocumentTarget, "target", "", "", "Language code that specifies the target language for translation.")
	cmd.Flags().StringVarP(&TranslateDocumentDocumentID, "document_id", "", "", "To use a previously submitted document as the source for a new translation, enter the `document_id` of the document.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("file")
	cmd.MarkFlagRequired("filename")
	cmd.MarkFlagRequired("version")

	return cmd
}

func TranslateDocument(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	languageTranslator, languageTranslatorErr := languagetranslatorv3.
		NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(languageTranslatorErr)

	optionsModel := languagetranslatorv3.TranslateDocumentOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "file" {
			file, fileErr := os.Open(TranslateDocumentFile)
			utils.HandleError(fileErr)

			optionsModel.SetFile(file)
		}
		if flag.Name == "filename" {
			optionsModel.SetFilename(TranslateDocumentFilename)
		}
		if flag.Name == "file_content_type" {
			optionsModel.SetFileContentType(TranslateDocumentFileContentType)
		}
		if flag.Name == "model_id" {
			optionsModel.SetModelID(TranslateDocumentModelID)
		}
		if flag.Name == "source" {
			optionsModel.SetSource(TranslateDocumentSource)
		}
		if flag.Name == "target" {
			optionsModel.SetTarget(TranslateDocumentTarget)
		}
		if flag.Name == "document_id" {
			optionsModel.SetDocumentID(TranslateDocumentDocumentID)
		}
	})

	result, _, responseErr := languageTranslator.TranslateDocument(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetDocumentStatusDocumentID string

func getGetDocumentStatusCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-document-status",
		Short: "Get document status",
		Long: "Gets the translation status of a document.",
		Run: GetDocumentStatus,
	}

	cmd.Flags().StringVarP(&GetDocumentStatusDocumentID, "document_id", "", "", "The document ID of the document.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("document_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetDocumentStatus(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	languageTranslator, languageTranslatorErr := languagetranslatorv3.
		NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(languageTranslatorErr)

	optionsModel := languagetranslatorv3.GetDocumentStatusOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "document_id" {
			optionsModel.SetDocumentID(GetDocumentStatusDocumentID)
		}
	})

	result, _, responseErr := languageTranslator.GetDocumentStatus(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteDocumentDocumentID string

func getDeleteDocumentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-document",
		Short: "Delete document",
		Long: "Deletes a document.",
		Run: DeleteDocument,
	}

	cmd.Flags().StringVarP(&DeleteDocumentDocumentID, "document_id", "", "", "Document ID of the document to delete.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")

	cmd.MarkFlagRequired("document_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteDocument(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	languageTranslator, languageTranslatorErr := languagetranslatorv3.
		NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(languageTranslatorErr)

	optionsModel := languagetranslatorv3.DeleteDocumentOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "document_id" {
			optionsModel.SetDocumentID(DeleteDocumentDocumentID)
		}
	})

	_, responseErr := languageTranslator.DeleteDocument(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var GetTranslatedDocumentDocumentID string
var GetTranslatedDocumentAccept string

func getGetTranslatedDocumentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-translated-document",
		Short: "Get translated document",
		Long: "Gets the translated document associated with the given document ID.",
		Run: GetTranslatedDocument,
	}

	cmd.Flags().StringVarP(&GetTranslatedDocumentDocumentID, "document_id", "", "", "The document ID of the document that was submitted for translation.")
	cmd.Flags().StringVarP(&GetTranslatedDocumentAccept, "accept", "", "", "The type of the response: application/powerpoint, application/mspowerpoint, application/x-rtf, application/json, application/xml, application/vnd.ms-excel, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet, application/vnd.ms-powerpoint, application/vnd.openxmlformats-officedocument.presentationml.presentation, application/msword, application/vnd.openxmlformats-officedocument.wordprocessingml.document, application/vnd.oasis.opendocument.spreadsheet, application/vnd.oasis.opendocument.presentation, application/vnd.oasis.opendocument.text, application/pdf, application/rtf, text/html, text/json, text/plain, text/richtext, text/rtf, or text/xml. A character encoding can be specified by including a `charset` parameter. For example, 'text/html;charset=utf-8'.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")
	cmd.Flags().StringVarP(&OutputFilename, "output_file", "", "", "Filename/path to write the resulting output to.")

	cmd.MarkFlagRequired("document_id")
	cmd.MarkFlagRequired("version")
	cmd.MarkFlagRequired("output_file")

	return cmd
}

func GetTranslatedDocument(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	languageTranslator, languageTranslatorErr := languagetranslatorv3.
		NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(languageTranslatorErr)

	optionsModel := languagetranslatorv3.GetTranslatedDocumentOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "document_id" {
			optionsModel.SetDocumentID(GetTranslatedDocumentDocumentID)
		}
		if flag.Name == "accept" {
			optionsModel.SetAccept(GetTranslatedDocumentAccept)
		}
	})

	result, _, responseErr := languageTranslator.GetTranslatedDocument(&optionsModel)
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
