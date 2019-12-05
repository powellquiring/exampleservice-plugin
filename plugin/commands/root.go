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
  "github.ibm.com/dustinpopp/ibm-generated-cli-plugin-template/utils"
  "github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
  "github.com/IBM/sandbox-for-cli/plugin/commands/exampleservicev1"
  "github.com/spf13/cobra"
)

// root command to hold all commands
var rootCommand = &cobra.Command{
  Use:   "exampleservicev1 [service] [operation]",
  Short: "",
  Long: "",
}

func Init() {
  // compile service commands into one slice
  serviceCommands := []*cobra.Command{
    exampleservicev1.GetExampleServiceV1Command(),
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
    // fmt.Println(i)
    // process this metadata and add to array
    // fmt.Println("service", service.Name())
    convertedCommand := getPluginCommand(service)
    pluginCommands = append(pluginCommands, convertedCommand)

    for _, operation := range service.Commands() {
      // process this metadata
      // fmt.Println("  operation", operation.Name())
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
