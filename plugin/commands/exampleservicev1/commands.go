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

package exampleservicev1

import (
  "github.ibm.com/dustinpopp/ibm-generated-cli-plugin-template/utils"
  "github.com/ibm/mysdk/exampleservicev1"
  "github.com/spf13/cobra"
  "github.com/spf13/pflag"
)

var ListResourcesLimit int64

func getListResourcesCommand() *cobra.Command {
  cmd := &cobra.Command{
    Use:   "list-resources",
    Short: "List all resources",
    Long: "",
    Run: ListResources,
  }

  cmd.Flags().Int64VarP(&ListResourcesLimit, "limit", "", 0, "How many items to return at one time (max 100).")
  cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
  cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")


  return cmd
}

func ListResources(cmd *cobra.Command, args []string) {
  ui.Say("...")

  exampleService, exampleServiceErr := exampleservicev1.
    NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
      Authenticator: Authenticator,
    })
  utils.HandleError(exampleServiceErr)

  optionsModel := exampleservicev1.ListResourcesOptions{}

  // optional params should only be set when they are explicitly passed by the user
  // otherwise, the default type values will be sent to the service
  flagSet := cmd.Flags()
  flagSet.Visit(func(flag *pflag.Flag) {
    if flag.Name == "limit" {
      optionsModel.SetLimit(ListResourcesLimit)
    }
  })

  result, _, responseErr := exampleService.ListResources(&optionsModel)
  utils.HandleError(responseErr)

  utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var CreateResourceResourceID int64
var CreateResourceName string
var CreateResourceTag string

func getCreateResourceCommand() *cobra.Command {
  cmd := &cobra.Command{
    Use:   "create-resource",
    Short: "Create a resource",
    Long: "",
    Run: CreateResource,
  }

  cmd.Flags().Int64VarP(&CreateResourceResourceID, "resource_id", "", 0, "The id of the resource.")
  cmd.Flags().StringVarP(&CreateResourceName, "name", "", "", "The name of the resource.")
  cmd.Flags().StringVarP(&CreateResourceTag, "tag", "", "", "A tag value for the resource.")
  cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
  cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")


  return cmd
}

func CreateResource(cmd *cobra.Command, args []string) {
  ui.Say("...")

  exampleService, exampleServiceErr := exampleservicev1.
    NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
      Authenticator: Authenticator,
    })
  utils.HandleError(exampleServiceErr)

  optionsModel := exampleservicev1.CreateResourceOptions{}

  // optional params should only be set when they are explicitly passed by the user
  // otherwise, the default type values will be sent to the service
  flagSet := cmd.Flags()
  flagSet.Visit(func(flag *pflag.Flag) {
    if flag.Name == "resource_id" {
      optionsModel.SetResourceID(CreateResourceResourceID)
    }
    if flag.Name == "name" {
      optionsModel.SetName(CreateResourceName)
    }
    if flag.Name == "tag" {
      optionsModel.SetTag(CreateResourceTag)
    }
  })

  result, _, responseErr := exampleService.CreateResource(&optionsModel)
  utils.HandleError(responseErr)

  utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetResourceResourceID string

func getGetResourceCommand() *cobra.Command {
  cmd := &cobra.Command{
    Use:   "get-resource",
    Short: "Info for a specific resource",
    Long: "",
    Run: GetResource,
  }

  cmd.Flags().StringVarP(&GetResourceResourceID, "resource_id", "", "", "The id of the resource to retrieve.")
  cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
  cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

  cmd.MarkFlagRequired("resource_id")

  return cmd
}

func GetResource(cmd *cobra.Command, args []string) {
  ui.Say("...")

  exampleService, exampleServiceErr := exampleservicev1.
    NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
      Authenticator: Authenticator,
    })
  utils.HandleError(exampleServiceErr)

  optionsModel := exampleservicev1.GetResourceOptions{}

  // optional params should only be set when they are explicitly passed by the user
  // otherwise, the default type values will be sent to the service
  flagSet := cmd.Flags()
  flagSet.Visit(func(flag *pflag.Flag) {
    if flag.Name == "resource_id" {
      optionsModel.SetResourceID(GetResourceResourceID)
    }
  })

  result, _, responseErr := exampleService.GetResource(&optionsModel)
  utils.HandleError(responseErr)

  utils.PrintOutput(result, OutputFormat, JMESQuery)
}
