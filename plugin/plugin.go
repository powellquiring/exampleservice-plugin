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

package plugin

import (
	"cli-watson-plugin/plugin/commands"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
)

const PluginName = "watson"

type Plugin struct {
	ui terminal.UI
}

func (watson *Plugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{ // a lot of things in here can be separated out as variables, etc
		Name: PluginName,
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 2,
		},
		MinCliVersion: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 1,
		},
		Namespaces: []plugin.Namespace{ // can be split out
			{
				Name: "watson",
				Description: "", // can be put into translations
			},
		},
		Commands: commands.GetMetadata(),
	}
}

func (watson *Plugin) Run(context plugin.PluginContext, args []string) {
	watson.ui = terminal.NewStdUI()
	commands.Execute()
}
