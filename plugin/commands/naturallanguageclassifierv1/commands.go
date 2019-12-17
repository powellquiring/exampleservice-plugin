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

package naturallanguageclassifierv1

import (
	"cli-watson-plugin/utils"
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/watson-developer-cloud/go-sdk/naturallanguageclassifierv1"
	"os"
)

var ClassifyClassifierID string
var ClassifyText string

func getClassifyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "classify",
		Short: "Classify a phrase",
		Long: "Returns label information for the input. The status must be `Available` before you can use the classifier to classify text.",
		Run: Classify,
	}

	cmd.Flags().StringVarP(&ClassifyClassifierID, "classifier_id", "", "", "Classifier ID to use.")
	cmd.Flags().StringVarP(&ClassifyText, "text", "", "", "The submitted phrase. The maximum length is 2048 characters.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("classifier_id")
	cmd.MarkFlagRequired("text")

	return cmd
}

func Classify(cmd *cobra.Command, args []string) {
	naturalLanguageClassifier, naturalLanguageClassifierErr := naturallanguageclassifierv1.
		NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(naturalLanguageClassifierErr)

	optionsModel := naturallanguageclassifierv1.ClassifyOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "classifier_id" {
			optionsModel.SetClassifierID(ClassifyClassifierID)
		}
		if flag.Name == "text" {
			optionsModel.SetText(ClassifyText)
		}
	})

	result, _, responseErr := naturalLanguageClassifier.Classify(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ClassifyCollectionClassifierID string
var ClassifyCollectionCollection string

func getClassifyCollectionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "classify-collection",
		Short: "Classify multiple phrases",
		Long: "Returns label information for multiple phrases. The status must be `Available` before you can use the classifier to classify text.Note that classifying Japanese texts is a beta feature.",
		Run: ClassifyCollection,
	}

	cmd.Flags().StringVarP(&ClassifyCollectionClassifierID, "classifier_id", "", "", "Classifier ID to use.")
	cmd.Flags().StringVarP(&ClassifyCollectionCollection, "collection", "", "", "The submitted phrases.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("classifier_id")
	cmd.MarkFlagRequired("collection")

	return cmd
}

func ClassifyCollection(cmd *cobra.Command, args []string) {
	naturalLanguageClassifier, naturalLanguageClassifierErr := naturallanguageclassifierv1.
		NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(naturalLanguageClassifierErr)

	optionsModel := naturallanguageclassifierv1.ClassifyCollectionOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "classifier_id" {
			optionsModel.SetClassifierID(ClassifyCollectionClassifierID)
		}
		if flag.Name == "collection" {
			var collection []naturallanguageclassifierv1.ClassifyInput
			decodeErr := json.Unmarshal([]byte(ClassifyCollectionCollection), &collection);
			utils.HandleError(decodeErr)

			optionsModel.SetCollection(collection)
		}
	})

	result, _, responseErr := naturalLanguageClassifier.ClassifyCollection(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var CreateClassifierTrainingMetadata string
var CreateClassifierTrainingData string

func getCreateClassifierCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-classifier",
		Short: "Create classifier",
		Long: "Sends data to create and train a classifier and returns information about the new classifier.",
		Run: CreateClassifier,
	}

	cmd.Flags().StringVarP(&CreateClassifierTrainingMetadata, "training_metadata", "", "", "Metadata in JSON format. The metadata identifies the language of the data, and an optional name to identify the classifier. Specify the language with the 2-letter primary language code as assigned in ISO standard 639.Supported languages are English (`en`), Arabic (`ar`), French (`fr`), German, (`de`), Italian (`it`), Japanese (`ja`), Korean (`ko`), Brazilian Portuguese (`pt`), and Spanish (`es`).")
	cmd.Flags().StringVarP(&CreateClassifierTrainingData, "training_data", "", "", "Training data in CSV format. Each text value must have at least one class. The data can include up to 3,000 classes and 20,000 records. For details, see [Data preparation](https://cloud.ibm.com/docs/services/natural-language-classifier?topic=natural-language-classifier-using-your-data).")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("training_metadata")
	cmd.MarkFlagRequired("training_data")

	return cmd
}

func CreateClassifier(cmd *cobra.Command, args []string) {
	naturalLanguageClassifier, naturalLanguageClassifierErr := naturallanguageclassifierv1.
		NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(naturalLanguageClassifierErr)

	optionsModel := naturallanguageclassifierv1.CreateClassifierOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "training_metadata" {
			training_metadata, fileErr := os.Open(CreateClassifierTrainingMetadata)
			utils.HandleError(fileErr)

			optionsModel.SetTrainingMetadata(training_metadata)
		}
		if flag.Name == "training_data" {
			training_data, fileErr := os.Open(CreateClassifierTrainingData)
			utils.HandleError(fileErr)

			optionsModel.SetTrainingData(training_data)
		}
	})

	result, _, responseErr := naturalLanguageClassifier.CreateClassifier(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}


func getListClassifiersCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-classifiers",
		Short: "List classifiers",
		Long: "Returns an empty array if no classifiers are available.",
		Run: ListClassifiers,
	}

	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")


	return cmd
}

func ListClassifiers(cmd *cobra.Command, args []string) {
	naturalLanguageClassifier, naturalLanguageClassifierErr := naturallanguageclassifierv1.
		NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(naturalLanguageClassifierErr)

	optionsModel := naturallanguageclassifierv1.ListClassifiersOptions{}

	result, _, responseErr := naturalLanguageClassifier.ListClassifiers(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetClassifierClassifierID string

func getGetClassifierCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-classifier",
		Short: "Get information about a classifier",
		Long: "Returns status and other information about a classifier.",
		Run: GetClassifier,
	}

	cmd.Flags().StringVarP(&GetClassifierClassifierID, "classifier_id", "", "", "Classifier ID to query.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("classifier_id")

	return cmd
}

func GetClassifier(cmd *cobra.Command, args []string) {
	naturalLanguageClassifier, naturalLanguageClassifierErr := naturallanguageclassifierv1.
		NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(naturalLanguageClassifierErr)

	optionsModel := naturallanguageclassifierv1.GetClassifierOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "classifier_id" {
			optionsModel.SetClassifierID(GetClassifierClassifierID)
		}
	})

	result, _, responseErr := naturalLanguageClassifier.GetClassifier(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteClassifierClassifierID string

func getDeleteClassifierCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-classifier",
		Short: "Delete classifier",
		Long: "",
		Run: DeleteClassifier,
	}

	cmd.Flags().StringVarP(&DeleteClassifierClassifierID, "classifier_id", "", "", "Classifier ID to delete.")

	cmd.MarkFlagRequired("classifier_id")

	return cmd
}

func DeleteClassifier(cmd *cobra.Command, args []string) {
	naturalLanguageClassifier, naturalLanguageClassifierErr := naturallanguageclassifierv1.
		NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
			Authenticator: Authenticator,
		})
	utils.HandleError(naturalLanguageClassifierErr)

	optionsModel := naturallanguageclassifierv1.DeleteClassifierOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "classifier_id" {
			optionsModel.SetClassifierID(DeleteClassifierClassifierID)
		}
	})

	_, responseErr := naturalLanguageClassifier.DeleteClassifier(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}
