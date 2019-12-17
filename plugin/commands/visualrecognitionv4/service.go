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
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/spf13/cobra"
)

var ui terminal.UI

// declare the Authenticator for the service
var Authenticator core.Authenticator
var Version string
var OutputFormat string
var JMESQuery string
var OutputFilename string

// add a function to return the super-command
func GetVisualRecognitionV4Command() *cobra.Command {
	ui = terminal.NewStdUI()

	// initialize the authenticator
	ServiceAuthenticator, AuthFactoryErr := core.GetAuthenticatorFromEnvironment("visual_recognition")
	utils.HandleError(AuthFactoryErr)

	Authenticator = ServiceAuthenticator

	serviceCommands := []*cobra.Command{
		getAnalyzeCommand(),
		getCreateCollectionCommand(),
		getListCollectionsCommand(),
		getGetCollectionCommand(),
		getUpdateCollectionCommand(),
		getDeleteCollectionCommand(),
		getAddImagesCommand(),
		getListImagesCommand(),
		getGetImageDetailsCommand(),
		getDeleteImageCommand(),
		getGetJpegImageCommand(),
		getTrainCommand(),
		getAddImageTrainingDataCommand(),
		getDeleteUserDataCommand(),
	}

	visualRecognitionCommand := &cobra.Command{
		Use: "visual-recognition-v4 [operation]",
		Aliases: []string{"vr-v4"},
		Short: "Parent command for Visual Recognition v4",
		Long: "Provide images to the IBM Watson&trade; Visual Recognition service for analysis. The service detects objects based on a set of images with training data.**Beta:** The Visual Recognition v4 API and Object Detection model are beta features. For more information about beta features, see the [Release notes](https://cloud.ibm.com/docs/services/visual-recognition?topic=visual-recognition-release-notes#beta).{: important}",
	}

	for _, cmd := range serviceCommands {
		visualRecognitionCommand.AddCommand(cmd)
	}

	return visualRecognitionCommand 
}
