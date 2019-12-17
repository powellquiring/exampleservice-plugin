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
func GetVisualRecognitionV3Command() *cobra.Command {
	ui = terminal.NewStdUI()

	// initialize the authenticator
	ServiceAuthenticator, AuthFactoryErr := core.GetAuthenticatorFromEnvironment("visual_recognition")
	utils.HandleError(AuthFactoryErr)

	Authenticator = ServiceAuthenticator

	serviceCommands := []*cobra.Command{
		getClassifyCommand(),
		getCreateClassifierCommand(),
		getListClassifiersCommand(),
		getGetClassifierCommand(),
		getUpdateClassifierCommand(),
		getDeleteClassifierCommand(),
		getGetCoreMlModelCommand(),
		getDeleteUserDataCommand(),
	}

	visualRecognitionCommand := &cobra.Command{
		Use: "visual-recognition-v3 [operation]",
		Aliases: []string{"vr-v3"},
		Short: "Parent command for Visual Recognition",
		Long: "The IBM Watson&trade; Visual Recognition service uses deep learning algorithms to identify scenes and objects in images that you upload to the service. You can create and train a custom classifier to identify subjects that suit your needs.",
	}

	for _, cmd := range serviceCommands {
		visualRecognitionCommand.AddCommand(cmd)
	}

	return visualRecognitionCommand 
}
