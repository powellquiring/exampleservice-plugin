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

// add a function to return the super-command
func GetToneAnalyzerV3Command() *cobra.Command {
	ui = terminal.NewStdUI()

	// initialize the authenticator
	ServiceAuthenticator, AuthFactoryErr := core.GetAuthenticatorFromEnvironment("tone_analyzer")
	utils.HandleError(AuthFactoryErr)

	Authenticator = ServiceAuthenticator

	serviceCommands := []*cobra.Command{
		getToneCommand(),
		getToneChatCommand(),
	}

	toneAnalyzerCommand := &cobra.Command{
		Use: "tone-analyzer-v3 [operation]",
		Aliases: []string{"ta-v3"},
		Short: "Parent command for Tone Analyzer",
		Long: "The IBM Watson&trade; Tone Analyzer service uses linguistic analysis to detect emotional and language tones in written text. The service can analyze tone at both the document and sentence levels. You can use the service to understand how your written communications are perceived and then to improve the tone of your communications. Businesses can use the service to learn the tone of their customers' communications and to respond to each customer appropriately, or to understand and improve their customer conversations.**Note:** Request logging is disabled for the Tone Analyzer service. Regardless of whether you set the `X-Watson-Learning-Opt-Out` request header, the service does not log or retain data from requests and responses.",
	}

	for _, cmd := range serviceCommands {
		toneAnalyzerCommand.AddCommand(cmd)
	}

	return toneAnalyzerCommand 
}
