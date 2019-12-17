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

package visualrecognitionv4

import (
	"cli-watson-plugin/utils"
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/watson-developer-cloud/go-sdk/visualrecognitionv4"
	"io"
	"os"
)

var AnalyzeCollectionIds []string
var AnalyzeFeatures []string
var AnalyzeImagesFile string
var AnalyzeImageURL []string
var AnalyzeThreshold float32

func getAnalyzeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "analyze",
		Short: "Analyze images",
		Long: "Analyze images by URL, by file, or both against your own collection. Make sure that **training_status.objects.ready** is `true` for the feature before you use a collection to analyze images.Encode the image and .zip file names in UTF-8 if they contain non-ASCII characters. The service assumes UTF-8 encoding if it encounters non-ASCII characters.",
		Run: Analyze,
	}

	cmd.Flags().StringSliceVarP(&AnalyzeCollectionIds, "collection_ids", "", nil, "The IDs of the collections to analyze.")
	cmd.Flags().StringSliceVarP(&AnalyzeFeatures, "features", "", nil, "The features to analyze.")
	cmd.Flags().StringVarP(&AnalyzeImagesFile, "images_file", "", "", "An array of image files (.jpg or .png) or .zip files with images.- Include a maximum of 20 images in a request.- Limit the .zip file to 100 MB.- Limit each image file to 10 MB.You can also include an image with the **image_url** parameter.")
	cmd.Flags().StringSliceVarP(&AnalyzeImageURL, "image_url", "", nil, "An array of URLs of image files (.jpg or .png).- Include a maximum of 20 images in a request.- Limit each image file to 10 MB.- Minimum width and height is 30 pixels, but the service tends to perform better with images that are at least 300 x 300 pixels. Maximum is 5400 pixels for either height or width.You can also include images with the **images_file** parameter.")
	cmd.Flags().Float32VarP(&AnalyzeThreshold, "threshold", "", 0, "The minimum score a feature must have to be returned.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("collection_ids")
	cmd.MarkFlagRequired("features")
	cmd.MarkFlagRequired("version")

	return cmd
}

func Analyze(cmd *cobra.Command, args []string) {
	visualRecognition, visualRecognitionErr := visualrecognitionv4.
		NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(visualRecognitionErr)

	optionsModel := visualrecognitionv4.AnalyzeOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "collection_ids" {
			optionsModel.SetCollectionIds(AnalyzeCollectionIds)
		}
		if flag.Name == "features" {
			optionsModel.SetFeatures(AnalyzeFeatures)
		}
		if flag.Name == "images_file" {
			images_file := make([]visualrecognitionv4.FileWithMetadata, 0)
			var userImagesFile []map[string]string
			decodeErr := json.Unmarshal([]byte(AnalyzeImagesFile), &userImagesFile);
			utils.HandleError(decodeErr)

			for _, fileWithMetadataMap := range userImagesFile {
					userFile, userFileErr := os.Open(fileWithMetadataMap["data"])
					utils.HandleError(userFileErr)

					// create new file with metadata
					newFile := visualrecognitionv4.FileWithMetadata{
						Data: userFile,
					}

					if fileWithMetadataMap["filename"] != "" {
						filename := fileWithMetadataMap["filename"]
						newFile.Filename = &filename
					}

					if fileWithMetadataMap["content_type"] != "" {
						contentType := fileWithMetadataMap["content_type"]
						newFile.ContentType = &contentType
					}

					images_file = append(images_file, newFile)
			}

			optionsModel.SetImagesFile(images_file)
		}
		if flag.Name == "image_url" {
			optionsModel.SetImageURL(AnalyzeImageURL)
		}
		if flag.Name == "threshold" {
			optionsModel.SetThreshold(AnalyzeThreshold)
		}
	})

	result, _, responseErr := visualRecognition.Analyze(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var CreateCollectionName string
var CreateCollectionDescription string

func getCreateCollectionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-collection",
		Short: "Create a collection",
		Long: "Create a collection that can be used to store images.To create a collection without specifying a name and description, include an empty JSON object in the request body.Encode the name and description in UTF-8 if they contain non-ASCII characters. The service assumes UTF-8 encoding if it encounters non-ASCII characters.",
		Run: CreateCollection,
	}

	cmd.Flags().StringVarP(&CreateCollectionName, "name", "", "", "The name of the collection. The name can contain alphanumeric, underscore, hyphen, and dot characters. It cannot begin with the reserved prefix `sys-`.")
	cmd.Flags().StringVarP(&CreateCollectionDescription, "description", "", "", "The description of the collection.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("version")

	return cmd
}

func CreateCollection(cmd *cobra.Command, args []string) {
	visualRecognition, visualRecognitionErr := visualrecognitionv4.
		NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(visualRecognitionErr)

	optionsModel := visualrecognitionv4.CreateCollectionOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "name" {
			optionsModel.SetName(CreateCollectionName)
		}
		if flag.Name == "description" {
			optionsModel.SetDescription(CreateCollectionDescription)
		}
	})

	result, _, responseErr := visualRecognition.CreateCollection(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}


func getListCollectionsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-collections",
		Short: "List collections",
		Long: "Retrieves a list of collections for the service instance.",
		Run: ListCollections,
	}

	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("version")

	return cmd
}

func ListCollections(cmd *cobra.Command, args []string) {
	visualRecognition, visualRecognitionErr := visualrecognitionv4.
		NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(visualRecognitionErr)

	optionsModel := visualrecognitionv4.ListCollectionsOptions{}

	result, _, responseErr := visualRecognition.ListCollections(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetCollectionCollectionID string

func getGetCollectionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-collection",
		Short: "Get collection details",
		Long: "Get details of one collection.",
		Run: GetCollection,
	}

	cmd.Flags().StringVarP(&GetCollectionCollectionID, "collection_id", "", "", "The identifier of the collection.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetCollection(cmd *cobra.Command, args []string) {
	visualRecognition, visualRecognitionErr := visualrecognitionv4.
		NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(visualRecognitionErr)

	optionsModel := visualrecognitionv4.GetCollectionOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(GetCollectionCollectionID)
		}
	})

	result, _, responseErr := visualRecognition.GetCollection(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var UpdateCollectionCollectionID string
var UpdateCollectionName string
var UpdateCollectionDescription string

func getUpdateCollectionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-collection",
		Short: "Update a collection",
		Long: "Update the name or description of a collection.Encode the name and description in UTF-8 if they contain non-ASCII characters. The service assumes UTF-8 encoding if it encounters non-ASCII characters.",
		Run: UpdateCollection,
	}

	cmd.Flags().StringVarP(&UpdateCollectionCollectionID, "collection_id", "", "", "The identifier of the collection.")
	cmd.Flags().StringVarP(&UpdateCollectionName, "name", "", "", "The name of the collection. The name can contain alphanumeric, underscore, hyphen, and dot characters. It cannot begin with the reserved prefix `sys-`.")
	cmd.Flags().StringVarP(&UpdateCollectionDescription, "description", "", "", "The description of the collection.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func UpdateCollection(cmd *cobra.Command, args []string) {
	visualRecognition, visualRecognitionErr := visualrecognitionv4.
		NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(visualRecognitionErr)

	optionsModel := visualrecognitionv4.UpdateCollectionOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(UpdateCollectionCollectionID)
		}
		if flag.Name == "name" {
			optionsModel.SetName(UpdateCollectionName)
		}
		if flag.Name == "description" {
			optionsModel.SetDescription(UpdateCollectionDescription)
		}
	})

	result, _, responseErr := visualRecognition.UpdateCollection(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteCollectionCollectionID string

func getDeleteCollectionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-collection",
		Short: "Delete a collection",
		Long: "Delete a collection from the service instance.",
		Run: DeleteCollection,
	}

	cmd.Flags().StringVarP(&DeleteCollectionCollectionID, "collection_id", "", "", "The identifier of the collection.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")

	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteCollection(cmd *cobra.Command, args []string) {
	visualRecognition, visualRecognitionErr := visualrecognitionv4.
		NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(visualRecognitionErr)

	optionsModel := visualrecognitionv4.DeleteCollectionOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(DeleteCollectionCollectionID)
		}
	})

	_, responseErr := visualRecognition.DeleteCollection(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var AddImagesCollectionID string
var AddImagesImagesFile string
var AddImagesImageURL []string
var AddImagesTrainingData string

func getAddImagesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "add-images",
		Short: "Add images",
		Long: "Add images to a collection by URL, by file, or both.Encode the image and .zip file names in UTF-8 if they contain non-ASCII characters. The service assumes UTF-8 encoding if it encounters non-ASCII characters.",
		Run: AddImages,
	}

	cmd.Flags().StringVarP(&AddImagesCollectionID, "collection_id", "", "", "The identifier of the collection.")
	cmd.Flags().StringVarP(&AddImagesImagesFile, "images_file", "", "", "An array of image files (.jpg or .png) or .zip files with images.- Include a maximum of 20 images in a request.- Limit the .zip file to 100 MB.- Limit each image file to 10 MB.You can also include an image with the **image_url** parameter.")
	cmd.Flags().StringSliceVarP(&AddImagesImageURL, "image_url", "", nil, "The array of URLs of image files (.jpg or .png).- Include a maximum of 20 images in a request.- Limit each image file to 10 MB.- Minimum width and height is 30 pixels, but the service tends to perform better with images that are at least 300 x 300 pixels. Maximum is 5400 pixels for either height or width.You can also include images with the **images_file** parameter.")
	cmd.Flags().StringVarP(&AddImagesTrainingData, "training_data", "", "", "Training data for a single image. Include training data only if you add one image with the request.The `object` property can contain alphanumeric, underscore, hyphen, space, and dot characters. It cannot begin with the reserved prefix `sys-` and must be no longer than 32 characters.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func AddImages(cmd *cobra.Command, args []string) {
	visualRecognition, visualRecognitionErr := visualrecognitionv4.
		NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(visualRecognitionErr)

	optionsModel := visualrecognitionv4.AddImagesOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(AddImagesCollectionID)
		}
		if flag.Name == "images_file" {
			images_file := make([]visualrecognitionv4.FileWithMetadata, 0)
			var userImagesFile []map[string]string
			decodeErr := json.Unmarshal([]byte(AddImagesImagesFile), &userImagesFile);
			utils.HandleError(decodeErr)

			for _, fileWithMetadataMap := range userImagesFile {
					userFile, userFileErr := os.Open(fileWithMetadataMap["data"])
					utils.HandleError(userFileErr)

					// create new file with metadata
					newFile := visualrecognitionv4.FileWithMetadata{
						Data: userFile,
					}

					if fileWithMetadataMap["filename"] != "" {
						filename := fileWithMetadataMap["filename"]
						newFile.Filename = &filename
					}

					if fileWithMetadataMap["content_type"] != "" {
						contentType := fileWithMetadataMap["content_type"]
						newFile.ContentType = &contentType
					}

					images_file = append(images_file, newFile)
			}

			optionsModel.SetImagesFile(images_file)
		}
		if flag.Name == "image_url" {
			optionsModel.SetImageURL(AddImagesImageURL)
		}
		if flag.Name == "training_data" {
			optionsModel.SetTrainingData(AddImagesTrainingData)
		}
	})

	result, _, responseErr := visualRecognition.AddImages(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ListImagesCollectionID string

func getListImagesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-images",
		Short: "List images",
		Long: "Retrieves a list of images in a collection.",
		Run: ListImages,
	}

	cmd.Flags().StringVarP(&ListImagesCollectionID, "collection_id", "", "", "The identifier of the collection.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func ListImages(cmd *cobra.Command, args []string) {
	visualRecognition, visualRecognitionErr := visualrecognitionv4.
		NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(visualRecognitionErr)

	optionsModel := visualrecognitionv4.ListImagesOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(ListImagesCollectionID)
		}
	})

	result, _, responseErr := visualRecognition.ListImages(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetImageDetailsCollectionID string
var GetImageDetailsImageID string

func getGetImageDetailsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-image-details",
		Short: "Get image details",
		Long: "Get the details of an image in a collection.",
		Run: GetImageDetails,
	}

	cmd.Flags().StringVarP(&GetImageDetailsCollectionID, "collection_id", "", "", "The identifier of the collection.")
	cmd.Flags().StringVarP(&GetImageDetailsImageID, "image_id", "", "", "The identifier of the image.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("image_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetImageDetails(cmd *cobra.Command, args []string) {
	visualRecognition, visualRecognitionErr := visualrecognitionv4.
		NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(visualRecognitionErr)

	optionsModel := visualrecognitionv4.GetImageDetailsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(GetImageDetailsCollectionID)
		}
		if flag.Name == "image_id" {
			optionsModel.SetImageID(GetImageDetailsImageID)
		}
	})

	result, _, responseErr := visualRecognition.GetImageDetails(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteImageCollectionID string
var DeleteImageImageID string

func getDeleteImageCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-image",
		Short: "Delete an image",
		Long: "Delete one image from a collection.",
		Run: DeleteImage,
	}

	cmd.Flags().StringVarP(&DeleteImageCollectionID, "collection_id", "", "", "The identifier of the collection.")
	cmd.Flags().StringVarP(&DeleteImageImageID, "image_id", "", "", "The identifier of the image.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")

	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("image_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteImage(cmd *cobra.Command, args []string) {
	visualRecognition, visualRecognitionErr := visualrecognitionv4.
		NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(visualRecognitionErr)

	optionsModel := visualrecognitionv4.DeleteImageOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(DeleteImageCollectionID)
		}
		if flag.Name == "image_id" {
			optionsModel.SetImageID(DeleteImageImageID)
		}
	})

	_, responseErr := visualRecognition.DeleteImage(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var GetJpegImageCollectionID string
var GetJpegImageImageID string
var GetJpegImageSize string

func getGetJpegImageCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-jpeg-image",
		Short: "Get a JPEG file of an image",
		Long: "Download a JPEG representation of an image.",
		Run: GetJpegImage,
	}

	cmd.Flags().StringVarP(&GetJpegImageCollectionID, "collection_id", "", "", "The identifier of the collection.")
	cmd.Flags().StringVarP(&GetJpegImageImageID, "image_id", "", "", "The identifier of the image.")
	cmd.Flags().StringVarP(&GetJpegImageSize, "size", "", "", "Specify the image size.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")
	cmd.Flags().StringVarP(&OutputFilename, "output_file", "", "", "Filename/path to write the resulting output to.")

	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("image_id")
	cmd.MarkFlagRequired("version")
	cmd.MarkFlagRequired("output_file")

	return cmd
}

func GetJpegImage(cmd *cobra.Command, args []string) {
	visualRecognition, visualRecognitionErr := visualrecognitionv4.
		NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(visualRecognitionErr)

	optionsModel := visualrecognitionv4.GetJpegImageOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(GetJpegImageCollectionID)
		}
		if flag.Name == "image_id" {
			optionsModel.SetImageID(GetJpegImageImageID)
		}
		if flag.Name == "size" {
			optionsModel.SetSize(GetJpegImageSize)
		}
	})

	result, _, responseErr := visualRecognition.GetJpegImage(&optionsModel)
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

var TrainCollectionID string

func getTrainCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "train",
		Short: "Train a collection",
		Long: "Start training on images in a collection. The collection must have enough training data and untrained data (the **training_status.objects.data_changed** is `true`). If training is in progress, the request queues the next training job.",
		Run: Train,
	}

	cmd.Flags().StringVarP(&TrainCollectionID, "collection_id", "", "", "The identifier of the collection.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func Train(cmd *cobra.Command, args []string) {
	visualRecognition, visualRecognitionErr := visualrecognitionv4.
		NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(visualRecognitionErr)

	optionsModel := visualrecognitionv4.TrainOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(TrainCollectionID)
		}
	})

	result, _, responseErr := visualRecognition.Train(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var AddImageTrainingDataCollectionID string
var AddImageTrainingDataImageID string
var AddImageTrainingDataObjects string

func getAddImageTrainingDataCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "add-image-training-data",
		Short: "Add training data to an image",
		Long: "Add, update, or delete training data for an image. Encode the object name in UTF-8 if it contains non-ASCII characters. The service assumes UTF-8 encoding if it encounters non-ASCII characters.Elements in the request replace the existing elements.- To update the training data, provide both the unchanged and the new or changed values.- To delete the training data, provide an empty value for the training data.",
		Run: AddImageTrainingData,
	}

	cmd.Flags().StringVarP(&AddImageTrainingDataCollectionID, "collection_id", "", "", "The identifier of the collection.")
	cmd.Flags().StringVarP(&AddImageTrainingDataImageID, "image_id", "", "", "The identifier of the image.")
	cmd.Flags().StringVarP(&AddImageTrainingDataObjects, "objects", "", "", "Training data for specific objects.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("image_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func AddImageTrainingData(cmd *cobra.Command, args []string) {
	visualRecognition, visualRecognitionErr := visualrecognitionv4.
		NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(visualRecognitionErr)

	optionsModel := visualrecognitionv4.AddImageTrainingDataOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(AddImageTrainingDataCollectionID)
		}
		if flag.Name == "image_id" {
			optionsModel.SetImageID(AddImageTrainingDataImageID)
		}
		if flag.Name == "objects" {
			var objects []visualrecognitionv4.TrainingDataObject
			decodeErr := json.Unmarshal([]byte(AddImageTrainingDataObjects), &objects);
			utils.HandleError(decodeErr)

			optionsModel.SetObjects(objects)
		}
	})

	result, _, responseErr := visualRecognition.AddImageTrainingData(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
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
	visualRecognition, visualRecognitionErr := visualrecognitionv4.
		NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(visualRecognitionErr)

	optionsModel := visualrecognitionv4.DeleteUserDataOptions{}

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
