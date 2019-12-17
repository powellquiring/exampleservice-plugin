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

package commands

import (
	"cli-watson-plugin/plugin/commands/assistantv1"
	"cli-watson-plugin/plugin/commands/assistantv2"
	"cli-watson-plugin/plugin/commands/comparecomplyv1"
	"cli-watson-plugin/plugin/commands/discoveryv1"
	"cli-watson-plugin/plugin/commands/languagetranslatorv3"
	"cli-watson-plugin/plugin/commands/naturallanguageclassifierv1"
	"cli-watson-plugin/plugin/commands/naturallanguageunderstandingv1"
	"cli-watson-plugin/plugin/commands/personalityinsightsv3"
	"cli-watson-plugin/plugin/commands/speechtotextv1"
	"cli-watson-plugin/plugin/commands/texttospeechv1"
	"cli-watson-plugin/plugin/commands/toneanalyzerv3"
	"cli-watson-plugin/plugin/commands/visualrecognitionv3"
	"cli-watson-plugin/plugin/commands/visualrecognitionv4"
	"cli-watson-plugin/utils"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.com/spf13/cobra"
)

// root command to hold all commands
var rootCommand = &cobra.Command{
	Use: "watson [service] [operation]",
	Short: "",
	Long: "",
}

func Init() {
	// compile service commands into one slice
	serviceCommands := []*cobra.Command{
		assistantv1.GetAssistantV1Command(),
		assistantv2.GetAssistantV2Command(),
		comparecomplyv1.GetCompareComplyV1Command(),
		discoveryv1.GetDiscoveryV1Command(),
		languagetranslatorv3.GetLanguageTranslatorV3Command(),
		naturallanguageclassifierv1.GetNaturalLanguageClassifierV1Command(),
		naturallanguageunderstandingv1.GetNaturalLanguageUnderstandingV1Command(),
		personalityinsightsv3.GetPersonalityInsightsV3Command(),
		speechtotextv1.GetSpeechToTextV1Command(),
		texttospeechv1.GetTextToSpeechV1Command(),
		toneanalyzerv3.GetToneAnalyzerV3Command(),
		visualrecognitionv3.GetVisualRecognitionV3Command(),
		visualrecognitionv4.GetVisualRecognitionV4Command(),
	}

	// add all the service commands to the root command
	// possibly could be split out to an init style command
	for _, cmd := range serviceCommands {
		rootCommand.AddCommand(cmd)
	}
}

func Execute() {
	err := rootCommand.Execute()
	utils.HandleError(err)
}

func GetMetadata() []plugin.Command {
	pluginCommands := make([]plugin.Command, 0)

	pluginCommands = append(pluginCommands, getPluginCommand(rootCommand))

	for _, service := range rootCommand.Commands() {
		// process this metadata and add to array
		convertedCommand := getPluginCommand(service)
		pluginCommands = append(pluginCommands, convertedCommand)

		for _, operation := range service.Commands() {
			// process this metadata
			convertedCommand := getPluginCommand(operation)
			pluginCommands = append(pluginCommands, convertedCommand)
		}
	}

	return pluginCommands
}

func getPluginCommand(cmd *cobra.Command) plugin.Command {
	rootName := cmd.Root().Name()
	isRoot := cmd.Name() == rootName

	parentIsRoot := cmd.HasParent() && cmd.Parent().Name() == rootName

	// process command name
	var name string
	if isRoot || parentIsRoot {
		name = cmd.Name()
	} else {
		// should have parent here - making that assumption
		name = cmd.Parent().Name() + " " + cmd.Name()
	}

	return plugin.Command{
		Namespace:   rootName,
		Name:        name,
		Description: cmd.Long,
		Usage:       cmd.Use,
		Aliases:     cmd.Aliases,
	}
}
