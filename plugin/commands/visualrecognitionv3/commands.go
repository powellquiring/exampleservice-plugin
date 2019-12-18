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

package visualrecognitionv3

import (
	"cli-watson-plugin/utils"
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/watson-developer-cloud/go-sdk/visualrecognitionv3"
	"io"
	"os"
)

var ClassifyImagesFile string
var ClassifyImagesFilename string
var ClassifyImagesFileContentType string
var ClassifyURL string
var ClassifyThreshold float32
var ClassifyOwners []string
var ClassifyClassifierIds []string
var ClassifyAcceptLanguage string

func getClassifyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "classify",
		Short: "Classify images",
		Long: "Classify images with built-in or custom classifiers.",
		Run: Classify,
	}

	cmd.Flags().StringVarP(&ClassifyImagesFile, "images_file", "", "", "An image file (.gif, .jpg, .png, .tif) or .zip file with images. Maximum image size is 10 MB. Include no more than 20 images and limit the .zip file to 100 MB. Encode the image and .zip file names in UTF-8 if they contain non-ASCII characters. The service assumes UTF-8 encoding if it encounters non-ASCII characters.You can also include an image with the **url** parameter.")
	cmd.Flags().StringVarP(&ClassifyImagesFilename, "images_filename", "", "", "The filename for ImagesFile.")
	cmd.Flags().StringVarP(&ClassifyImagesFileContentType, "images_file_content_type", "", "", "The content type of ImagesFile.")
	cmd.Flags().StringVarP(&ClassifyURL, "url", "", "", "The URL of an image (.gif, .jpg, .png, .tif) to analyze. The minimum recommended pixel density is 32X32 pixels, but the service tends to perform better with images that are at least 224 x 224 pixels. The maximum image size is 10 MB.You can also include images with the **images_file** parameter.")
	cmd.Flags().Float32VarP(&ClassifyThreshold, "threshold", "", 0, "The minimum score a class must have to be displayed in the response. Set the threshold to `0.0` to return all identified classes.")
	cmd.Flags().StringSliceVarP(&ClassifyOwners, "owners", "", nil, "The categories of classifiers to apply. The **classifier_ids** parameter overrides **owners**, so make sure that **classifier_ids** is empty. - Use `IBM` to classify against the `default` general classifier. You get the same result if both **classifier_ids** and **owners** parameters are empty.- Use `me` to classify against all your custom classifiers. However, for better performance use **classifier_ids** to specify the specific custom classifiers to apply.- Use both `IBM` and `me` to analyze the image against both classifier categories.")
	cmd.Flags().StringSliceVarP(&ClassifyClassifierIds, "classifier_ids", "", nil, "Which classifiers to apply. Overrides the **owners** parameter. You can specify both custom and built-in classifier IDs. The built-in `default` classifier is used if both **classifier_ids** and **owners** parameters are empty.The following built-in classifier IDs require no training:- `default`: Returns classes from thousands of general tags.- `food`: Enhances specificity and accuracy for images of food items.- `explicit`: Evaluates whether the image might be pornographic.")
	cmd.Flags().StringVarP(&ClassifyAcceptLanguage, "accept_language", "", "", "The desired language of parts of the response. See the response for details.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("version")

	return cmd
}

func Classify(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	visualRecognition, visualRecognitionErr := visualrecognitionv3.
		NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(visualRecognitionErr)

	optionsModel := visualrecognitionv3.ClassifyOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "images_file" {
			images_file, fileErr := os.Open(ClassifyImagesFile)
			utils.HandleError(fileErr)

			optionsModel.SetImagesFile(images_file)
		}
		if flag.Name == "images_filename" {
			optionsModel.SetImagesFilename(ClassifyImagesFilename)
		}
		if flag.Name == "images_file_content_type" {
			optionsModel.SetImagesFileContentType(ClassifyImagesFileContentType)
		}
		if flag.Name == "url" {
			optionsModel.SetURL(ClassifyURL)
		}
		if flag.Name == "threshold" {
			optionsModel.SetThreshold(ClassifyThreshold)
		}
		if flag.Name == "owners" {
			optionsModel.SetOwners(ClassifyOwners)
		}
		if flag.Name == "classifier_ids" {
			optionsModel.SetClassifierIds(ClassifyClassifierIds)
		}
		if flag.Name == "accept_language" {
			optionsModel.SetAcceptLanguage(ClassifyAcceptLanguage)
		}
	})

	result, _, responseErr := visualRecognition.Classify(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var CreateClassifierName string
var CreateClassifierPositiveExamples string
var CreateClassifierNegativeExamples string
var CreateClassifierNegativeExamplesFilename string

func getCreateClassifierCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-classifier",
		Short: "Create a classifier",
		Long: "Train a new multi-faceted classifier on the uploaded image data. Create your custom classifier with positive or negative example training images. Include at least two sets of examples, either two positive example files or one positive and one negative file. You can upload a maximum of 256 MB per call.**Tips when creating:**- If you set the **X-Watson-Learning-Opt-Out** header parameter to `true` when you create a classifier, the example training images are not stored. Save your training images locally. For more information, see [Data collection](#data-collection).- Encode all names in UTF-8 if they contain non-ASCII characters (.zip and image file names, and classifier and class names). The service assumes UTF-8 encoding if it encounters non-ASCII characters.",
		Run: CreateClassifier,
	}

	cmd.Flags().StringVarP(&CreateClassifierName, "name", "", "", "The name of the new classifier. Encode special characters in UTF-8.")
	cmd.Flags().StringVarP(&CreateClassifierPositiveExamples, "positive_examples", "", "", "A .zip file of images that depict the visual subject of a class in the new classifier. You can include more than one positive example file in a call.Specify the parameter name by appending `_positive_examples` to the class name. For example, `goldenretriever_positive_examples` creates the class **goldenretriever**.Include at least 10 images in .jpg or .png format. The minimum recommended image resolution is 32X32 pixels. The maximum number of images is 10,000 images or 100 MB per .zip file.Encode special characters in the file name in UTF-8.")
	cmd.Flags().StringVarP(&CreateClassifierNegativeExamples, "negative_examples", "", "", "A .zip file of images that do not depict the visual subject of any of the classes of the new classifier. Must contain a minimum of 10 images.Encode special characters in the file name in UTF-8.")
	cmd.Flags().StringVarP(&CreateClassifierNegativeExamplesFilename, "negative_examples_filename", "", "", "The filename for NegativeExamples.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("positive_examples")
	cmd.MarkFlagRequired("version")

	return cmd
}

func CreateClassifier(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	visualRecognition, visualRecognitionErr := visualrecognitionv3.
		NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(visualRecognitionErr)

	optionsModel := visualrecognitionv3.CreateClassifierOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "name" {
			optionsModel.SetName(CreateClassifierName)
		}
		if flag.Name == "positive_examples" {
			var userPositiveExamples map[string]string
			decodeErr := json.Unmarshal([]byte(CreateClassifierPositiveExamples), &userPositiveExamples);
			utils.HandleError(decodeErr)

			for key, value := range userPositiveExamples {
					file, err := os.Open(value)
					utils.HandleError(err)
					optionsModel.AddPositiveExamples(key, file)
			}
		}
		if flag.Name == "negative_examples" {
			negative_examples, fileErr := os.Open(CreateClassifierNegativeExamples)
			utils.HandleError(fileErr)

			optionsModel.SetNegativeExamples(negative_examples)
		}
		if flag.Name == "negative_examples_filename" {
			optionsModel.SetNegativeExamplesFilename(CreateClassifierNegativeExamplesFilename)
		}
	})

	result, _, responseErr := visualRecognition.CreateClassifier(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ListClassifiersVerbose bool

func getListClassifiersCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-classifiers",
		Short: "Retrieve a list of classifiers",
		Long: "",
		Run: ListClassifiers,
	}

	cmd.Flags().BoolVarP(&ListClassifiersVerbose, "verbose", "", false, "Specify `true` to return details about the classifiers. Omit this parameter to return a brief list of classifiers.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("version")

	return cmd
}

func ListClassifiers(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	visualRecognition, visualRecognitionErr := visualrecognitionv3.
		NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(visualRecognitionErr)

	optionsModel := visualrecognitionv3.ListClassifiersOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "verbose" {
			optionsModel.SetVerbose(ListClassifiersVerbose)
		}
	})

	result, _, responseErr := visualRecognition.ListClassifiers(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetClassifierClassifierID string

func getGetClassifierCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-classifier",
		Short: "Retrieve classifier details",
		Long: "Retrieve information about a custom classifier.",
		Run: GetClassifier,
	}

	cmd.Flags().StringVarP(&GetClassifierClassifierID, "classifier_id", "", "", "The ID of the classifier.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("classifier_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetClassifier(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	visualRecognition, visualRecognitionErr := visualrecognitionv3.
		NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(visualRecognitionErr)

	optionsModel := visualrecognitionv3.GetClassifierOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "classifier_id" {
			optionsModel.SetClassifierID(GetClassifierClassifierID)
		}
	})

	result, _, responseErr := visualRecognition.GetClassifier(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var UpdateClassifierClassifierID string
var UpdateClassifierPositiveExamples string
var UpdateClassifierNegativeExamples string
var UpdateClassifierNegativeExamplesFilename string

func getUpdateClassifierCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-classifier",
		Short: "Update a classifier",
		Long: "Update a custom classifier by adding new positive or negative classes or by adding new images to existing classes. You must supply at least one set of positive or negative examples. For details, see [Updating custom classifiers](https://cloud.ibm.com/docs/services/visual-recognition?topic=visual-recognition-customizing#updating-custom-classifiers).Encode all names in UTF-8 if they contain non-ASCII characters (.zip and image file names, and classifier and class names). The service assumes UTF-8 encoding if it encounters non-ASCII characters.**Tips about retraining:**- You can't update the classifier if the **X-Watson-Learning-Opt-Out** header parameter was set to `true` when the classifier was created. Training images are not stored in that case. Instead, create another classifier. For more information, see [Data collection](#data-collection).- Don't make retraining calls on a classifier until the status is ready. When you submit retraining requests in parallel, the last request overwrites the previous requests. The `retrained` property shows the last time the classifier retraining finished.",
		Run: UpdateClassifier,
	}

	cmd.Flags().StringVarP(&UpdateClassifierClassifierID, "classifier_id", "", "", "The ID of the classifier.")
	cmd.Flags().StringVarP(&UpdateClassifierPositiveExamples, "positive_examples", "", "", "A .zip file of images that depict the visual subject of a class in the classifier. The positive examples create or update classes in the classifier. You can include more than one positive example file in a call.Specify the parameter name by appending `_positive_examples` to the class name. For example, `goldenretriever_positive_examples` creates the class `goldenretriever`.Include at least 10 images in .jpg or .png format. The minimum recommended image resolution is 32X32 pixels. The maximum number of images is 10,000 images or 100 MB per .zip file.Encode special characters in the file name in UTF-8.")
	cmd.Flags().StringVarP(&UpdateClassifierNegativeExamples, "negative_examples", "", "", "A .zip file of images that do not depict the visual subject of any of the classes of the new classifier. Must contain a minimum of 10 images.Encode special characters in the file name in UTF-8.")
	cmd.Flags().StringVarP(&UpdateClassifierNegativeExamplesFilename, "negative_examples_filename", "", "", "The filename for NegativeExamples.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("classifier_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func UpdateClassifier(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	visualRecognition, visualRecognitionErr := visualrecognitionv3.
		NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(visualRecognitionErr)

	optionsModel := visualrecognitionv3.UpdateClassifierOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "classifier_id" {
			optionsModel.SetClassifierID(UpdateClassifierClassifierID)
		}
		if flag.Name == "positive_examples" {
			var userPositiveExamples map[string]string
			decodeErr := json.Unmarshal([]byte(UpdateClassifierPositiveExamples), &userPositiveExamples);
			utils.HandleError(decodeErr)

			for key, value := range userPositiveExamples {
					file, err := os.Open(value)
					utils.HandleError(err)
					optionsModel.AddPositiveExamples(key, file)
			}
		}
		if flag.Name == "negative_examples" {
			negative_examples, fileErr := os.Open(UpdateClassifierNegativeExamples)
			utils.HandleError(fileErr)

			optionsModel.SetNegativeExamples(negative_examples)
		}
		if flag.Name == "negative_examples_filename" {
			optionsModel.SetNegativeExamplesFilename(UpdateClassifierNegativeExamplesFilename)
		}
	})

	result, _, responseErr := visualRecognition.UpdateClassifier(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteClassifierClassifierID string

func getDeleteClassifierCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-classifier",
		Short: "Delete a classifier",
		Long: "",
		Run: DeleteClassifier,
	}

	cmd.Flags().StringVarP(&DeleteClassifierClassifierID, "classifier_id", "", "", "The ID of the classifier.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")

	cmd.MarkFlagRequired("classifier_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteClassifier(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	visualRecognition, visualRecognitionErr := visualrecognitionv3.
		NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(visualRecognitionErr)

	optionsModel := visualrecognitionv3.DeleteClassifierOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "classifier_id" {
			optionsModel.SetClassifierID(DeleteClassifierClassifierID)
		}
	})

	_, responseErr := visualRecognition.DeleteClassifier(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var GetCoreMlModelClassifierID string

func getGetCoreMlModelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-core-ml-model",
		Short: "Retrieve a Core ML model of a classifier",
		Long: "Download a Core ML model file (.mlmodel) of a custom classifier that returns <tt>'core_ml_enabled': true</tt> in the classifier details.",
		Run: GetCoreMlModel,
	}

	cmd.Flags().StringVarP(&GetCoreMlModelClassifierID, "classifier_id", "", "", "The ID of the classifier.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")
	cmd.Flags().StringVarP(&OutputFilename, "output_file", "", "", "Filename/path to write the resulting output to.")

	cmd.MarkFlagRequired("classifier_id")
	cmd.MarkFlagRequired("version")
	cmd.MarkFlagRequired("output_file")

	return cmd
}

func GetCoreMlModel(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	visualRecognition, visualRecognitionErr := visualrecognitionv3.
		NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(visualRecognitionErr)

	optionsModel := visualrecognitionv3.GetCoreMlModelOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "classifier_id" {
			optionsModel.SetClassifierID(GetCoreMlModelClassifierID)
		}
	})

	result, _, responseErr := visualRecognition.GetCoreMlModel(&optionsModel)
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

var DeleteUserDataCustomerID string

func getDeleteUserDataCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-user-data",
		Short: "Delete labeled data",
		Long: "Deletes all data associated with a specified customer ID. The method has no effect if no data is associated with the customer ID. You associate a customer ID with data by passing the `X-Watson-Metadata` header with a request that passes data. For more information about personal data and customer IDs, see [Information security](https://cloud.ibm.com/docs/services/visual-recognition?topic=visual-recognition-information-security).",
		Run: DeleteUserData,
	}

	cmd.Flags().StringVarP(&DeleteUserDataCustomerID, "customer_id", "", "", "The customer ID for which all data is to be deleted.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")

	cmd.MarkFlagRequired("customer_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteUserData(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	visualRecognition, visualRecognitionErr := visualrecognitionv3.
		NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(visualRecognitionErr)

	optionsModel := visualrecognitionv3.DeleteUserDataOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customer_id" {
			optionsModel.SetCustomerID(DeleteUserDataCustomerID)
		}
	})

	_, responseErr := visualRecognition.DeleteUserData(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}
