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

package utils

import (
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	JmesPath "github.com/jmespath/go-jmespath"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"os"
	"strings"
)

var ui terminal.UI

func PrintOutput(result interface{}, outputFormat string, jmesQuery string) {
	ui = terminal.NewStdUI()

	// the jmes query applies to everything, so maybe should do it first
	if (jmesQuery != "") {
		jmes, jErr := JmesPath.Compile(jmesQuery)
		HandleError(jErr)

		newJson, searchErr := jmes.Search(result)
		HandleError(searchErr)

		result = newJson;
	}

	// print something based on the output format
	switch strings.ToLower(outputFormat) {
	case "yaml":
		yamlified, yErr := yaml.Marshal(result)
		HandleError(yErr)

		fmt.Println(string(yamlified))

	case "json":
		// this will print raw json
		b, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(b))

	default:
		// default to "table" - this will a dynamically generated table
		ui.Say("...")
		DoTheTable(result, jmesQuery)
	}
}

func HandleError(err error) {
	ui = terminal.NewStdUI()
	if err != nil {
		ui.Failed(err.Error())
		os.Exit(1)
	}
}
