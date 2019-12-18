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

package discoveryv1

import (
	"cli-watson-plugin/utils"
	"encoding/json"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/watson-developer-cloud/go-sdk/discoveryv1"
	"os"
)

var CreateEnvironmentName string
var CreateEnvironmentDescription string
var CreateEnvironmentSize string

func getCreateEnvironmentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-environment",
		Short: "Create an environment",
		Long: "Creates a new environment for private data. An environment must be created before collections can be created. **Note**: You can create only one environment for private data per service instance. An attempt to create another environment results in an error.",
		Run: CreateEnvironment,
	}

	cmd.Flags().StringVarP(&CreateEnvironmentName, "name", "", "", "Name that identifies the environment.")
	cmd.Flags().StringVarP(&CreateEnvironmentDescription, "description", "", "", "Description of the environment.")
	cmd.Flags().StringVarP(&CreateEnvironmentSize, "size", "", "", "Size of the environment. In the Lite plan the default and only accepted value is `LT`, in all other plans the default is `S`.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("version")

	return cmd
}

func CreateEnvironment(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.CreateEnvironmentOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "name" {
			optionsModel.SetName(CreateEnvironmentName)
		}
		if flag.Name == "description" {
			optionsModel.SetDescription(CreateEnvironmentDescription)
		}
		if flag.Name == "size" {
			optionsModel.SetSize(CreateEnvironmentSize)
		}
	})

	result, _, responseErr := discovery.CreateEnvironment(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ListEnvironmentsName string

func getListEnvironmentsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-environments",
		Short: "List environments",
		Long: "List existing environments for the service instance.",
		Run: ListEnvironments,
	}

	cmd.Flags().StringVarP(&ListEnvironmentsName, "name", "", "", "Show only the environment with the given name.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("version")

	return cmd
}

func ListEnvironments(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.ListEnvironmentsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "name" {
			optionsModel.SetName(ListEnvironmentsName)
		}
	})

	result, _, responseErr := discovery.ListEnvironments(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetEnvironmentEnvironmentID string

func getGetEnvironmentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-environment",
		Short: "Get environment info",
		Long: "",
		Run: GetEnvironment,
	}

	cmd.Flags().StringVarP(&GetEnvironmentEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetEnvironment(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.GetEnvironmentOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(GetEnvironmentEnvironmentID)
		}
	})

	result, _, responseErr := discovery.GetEnvironment(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var UpdateEnvironmentEnvironmentID string
var UpdateEnvironmentName string
var UpdateEnvironmentDescription string
var UpdateEnvironmentSize string

func getUpdateEnvironmentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-environment",
		Short: "Update an environment",
		Long: "Updates an environment. The environment's **name** and  **description** parameters can be changed. You must specify a **name** for the environment.",
		Run: UpdateEnvironment,
	}

	cmd.Flags().StringVarP(&UpdateEnvironmentEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&UpdateEnvironmentName, "name", "", "", "Name that identifies the environment.")
	cmd.Flags().StringVarP(&UpdateEnvironmentDescription, "description", "", "", "Description of the environment.")
	cmd.Flags().StringVarP(&UpdateEnvironmentSize, "size", "", "", "Size that the environment should be increased to. Environment size cannot be modified when using a Lite plan. Environment size can only increased and not decreased.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func UpdateEnvironment(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.UpdateEnvironmentOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(UpdateEnvironmentEnvironmentID)
		}
		if flag.Name == "name" {
			optionsModel.SetName(UpdateEnvironmentName)
		}
		if flag.Name == "description" {
			optionsModel.SetDescription(UpdateEnvironmentDescription)
		}
		if flag.Name == "size" {
			optionsModel.SetSize(UpdateEnvironmentSize)
		}
	})

	result, _, responseErr := discovery.UpdateEnvironment(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteEnvironmentEnvironmentID string

func getDeleteEnvironmentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-environment",
		Short: "Delete environment",
		Long: "",
		Run: DeleteEnvironment,
	}

	cmd.Flags().StringVarP(&DeleteEnvironmentEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteEnvironment(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.DeleteEnvironmentOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(DeleteEnvironmentEnvironmentID)
		}
	})

	result, _, responseErr := discovery.DeleteEnvironment(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ListFieldsEnvironmentID string
var ListFieldsCollectionIds []string

func getListFieldsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-fields",
		Short: "List fields across collections",
		Long: "Gets a list of the unique fields (and their types) stored in the indexes of the specified collections.",
		Run: ListFields,
	}

	cmd.Flags().StringVarP(&ListFieldsEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringSliceVarP(&ListFieldsCollectionIds, "collection_ids", "", nil, "A comma-separated list of collection IDs to be queried against.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_ids")
	cmd.MarkFlagRequired("version")

	return cmd
}

func ListFields(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.ListFieldsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(ListFieldsEnvironmentID)
		}
		if flag.Name == "collection_ids" {
			optionsModel.SetCollectionIds(ListFieldsCollectionIds)
		}
	})

	result, _, responseErr := discovery.ListFields(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var CreateConfigurationEnvironmentID string
var CreateConfigurationName string
var CreateConfigurationDescription string
var CreateConfigurationConversions string
var CreateConfigurationEnrichments string
var CreateConfigurationNormalizations string
var CreateConfigurationSource string

func getCreateConfigurationCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-configuration",
		Short: "Add configuration",
		Long: "Creates a new configuration.If the input configuration contains the **configuration_id**, **created**, or **updated** properties, then they are ignored and overridden by the system, and an error is not returned so that the overridden fields do not need to be removed when copying a configuration.The configuration can contain unrecognized JSON fields. Any such fields are ignored and do not generate an error. This makes it easier to use newer configuration files with older versions of the API and the service. It also makes it possible for the tooling to add additional metadata and information to the configuration.",
		Run: CreateConfiguration,
	}

	cmd.Flags().StringVarP(&CreateConfigurationEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&CreateConfigurationName, "name", "", "", "The name of the configuration.")
	cmd.Flags().StringVarP(&CreateConfigurationDescription, "description", "", "", "The description of the configuration, if available.")
	cmd.Flags().StringVarP(&CreateConfigurationConversions, "conversions", "", "", "Document conversion settings.")
	cmd.Flags().StringVarP(&CreateConfigurationEnrichments, "enrichments", "", "", "An array of document enrichment settings for the configuration.")
	cmd.Flags().StringVarP(&CreateConfigurationNormalizations, "normalizations", "", "", "Defines operations that can be used to transform the final output JSON into a normalized form. Operations are executed in the order that they appear in the array.")
	cmd.Flags().StringVarP(&CreateConfigurationSource, "source", "", "", "Object containing source parameters for the configuration.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("version")

	return cmd
}

func CreateConfiguration(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.CreateConfigurationOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(CreateConfigurationEnvironmentID)
		}
		if flag.Name == "name" {
			optionsModel.SetName(CreateConfigurationName)
		}
		if flag.Name == "description" {
			optionsModel.SetDescription(CreateConfigurationDescription)
		}
		if flag.Name == "conversions" {
			var conversions discoveryv1.Conversions
			decodeErr := json.Unmarshal([]byte(CreateConfigurationConversions), &conversions);
			utils.HandleError(decodeErr)

			optionsModel.SetConversions(&conversions)
		}
		if flag.Name == "enrichments" {
			var enrichments []discoveryv1.Enrichment
			decodeErr := json.Unmarshal([]byte(CreateConfigurationEnrichments), &enrichments);
			utils.HandleError(decodeErr)

			optionsModel.SetEnrichments(enrichments)
		}
		if flag.Name == "normalizations" {
			var normalizations []discoveryv1.NormalizationOperation
			decodeErr := json.Unmarshal([]byte(CreateConfigurationNormalizations), &normalizations);
			utils.HandleError(decodeErr)

			optionsModel.SetNormalizations(normalizations)
		}
		if flag.Name == "source" {
			var source discoveryv1.Source
			decodeErr := json.Unmarshal([]byte(CreateConfigurationSource), &source);
			utils.HandleError(decodeErr)

			optionsModel.SetSource(&source)
		}
	})

	result, _, responseErr := discovery.CreateConfiguration(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ListConfigurationsEnvironmentID string
var ListConfigurationsName string

func getListConfigurationsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-configurations",
		Short: "List configurations",
		Long: "Lists existing configurations for the service instance.",
		Run: ListConfigurations,
	}

	cmd.Flags().StringVarP(&ListConfigurationsEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&ListConfigurationsName, "name", "", "", "Find configurations with the given name.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func ListConfigurations(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.ListConfigurationsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(ListConfigurationsEnvironmentID)
		}
		if flag.Name == "name" {
			optionsModel.SetName(ListConfigurationsName)
		}
	})

	result, _, responseErr := discovery.ListConfigurations(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetConfigurationEnvironmentID string
var GetConfigurationConfigurationID string

func getGetConfigurationCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-configuration",
		Short: "Get configuration details",
		Long: "",
		Run: GetConfiguration,
	}

	cmd.Flags().StringVarP(&GetConfigurationEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&GetConfigurationConfigurationID, "configuration_id", "", "", "The ID of the configuration.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("configuration_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetConfiguration(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.GetConfigurationOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(GetConfigurationEnvironmentID)
		}
		if flag.Name == "configuration_id" {
			optionsModel.SetConfigurationID(GetConfigurationConfigurationID)
		}
	})

	result, _, responseErr := discovery.GetConfiguration(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var UpdateConfigurationEnvironmentID string
var UpdateConfigurationConfigurationID string
var UpdateConfigurationName string
var UpdateConfigurationDescription string
var UpdateConfigurationConversions string
var UpdateConfigurationEnrichments string
var UpdateConfigurationNormalizations string
var UpdateConfigurationSource string

func getUpdateConfigurationCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-configuration",
		Short: "Update a configuration",
		Long: "Replaces an existing configuration.  * Completely replaces the original configuration.  * The **configuration_id**, **updated**, and **created** fields are accepted in the request, but they are ignored, and an error is not generated. It is also acceptable for users to submit an updated configuration with none of the three properties.  * Documents are processed with a snapshot of the configuration as it was at the time the document was submitted to be ingested. This means that already submitted documents will not see any updates made to the configuration.",
		Run: UpdateConfiguration,
	}

	cmd.Flags().StringVarP(&UpdateConfigurationEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&UpdateConfigurationConfigurationID, "configuration_id", "", "", "The ID of the configuration.")
	cmd.Flags().StringVarP(&UpdateConfigurationName, "name", "", "", "The name of the configuration.")
	cmd.Flags().StringVarP(&UpdateConfigurationDescription, "description", "", "", "The description of the configuration, if available.")
	cmd.Flags().StringVarP(&UpdateConfigurationConversions, "conversions", "", "", "Document conversion settings.")
	cmd.Flags().StringVarP(&UpdateConfigurationEnrichments, "enrichments", "", "", "An array of document enrichment settings for the configuration.")
	cmd.Flags().StringVarP(&UpdateConfigurationNormalizations, "normalizations", "", "", "Defines operations that can be used to transform the final output JSON into a normalized form. Operations are executed in the order that they appear in the array.")
	cmd.Flags().StringVarP(&UpdateConfigurationSource, "source", "", "", "Object containing source parameters for the configuration.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("configuration_id")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("version")

	return cmd
}

func UpdateConfiguration(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.UpdateConfigurationOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(UpdateConfigurationEnvironmentID)
		}
		if flag.Name == "configuration_id" {
			optionsModel.SetConfigurationID(UpdateConfigurationConfigurationID)
		}
		if flag.Name == "name" {
			optionsModel.SetName(UpdateConfigurationName)
		}
		if flag.Name == "description" {
			optionsModel.SetDescription(UpdateConfigurationDescription)
		}
		if flag.Name == "conversions" {
			var conversions discoveryv1.Conversions
			decodeErr := json.Unmarshal([]byte(UpdateConfigurationConversions), &conversions);
			utils.HandleError(decodeErr)

			optionsModel.SetConversions(&conversions)
		}
		if flag.Name == "enrichments" {
			var enrichments []discoveryv1.Enrichment
			decodeErr := json.Unmarshal([]byte(UpdateConfigurationEnrichments), &enrichments);
			utils.HandleError(decodeErr)

			optionsModel.SetEnrichments(enrichments)
		}
		if flag.Name == "normalizations" {
			var normalizations []discoveryv1.NormalizationOperation
			decodeErr := json.Unmarshal([]byte(UpdateConfigurationNormalizations), &normalizations);
			utils.HandleError(decodeErr)

			optionsModel.SetNormalizations(normalizations)
		}
		if flag.Name == "source" {
			var source discoveryv1.Source
			decodeErr := json.Unmarshal([]byte(UpdateConfigurationSource), &source);
			utils.HandleError(decodeErr)

			optionsModel.SetSource(&source)
		}
	})

	result, _, responseErr := discovery.UpdateConfiguration(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteConfigurationEnvironmentID string
var DeleteConfigurationConfigurationID string

func getDeleteConfigurationCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-configuration",
		Short: "Delete a configuration",
		Long: "The deletion is performed unconditionally. A configuration deletion request succeeds even if the configuration is referenced by a collection or document ingestion. However, documents that have already been submitted for processing continue to use the deleted configuration. Documents are always processed with a snapshot of the configuration as it existed at the time the document was submitted.",
		Run: DeleteConfiguration,
	}

	cmd.Flags().StringVarP(&DeleteConfigurationEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&DeleteConfigurationConfigurationID, "configuration_id", "", "", "The ID of the configuration.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("configuration_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteConfiguration(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.DeleteConfigurationOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(DeleteConfigurationEnvironmentID)
		}
		if flag.Name == "configuration_id" {
			optionsModel.SetConfigurationID(DeleteConfigurationConfigurationID)
		}
	})

	result, _, responseErr := discovery.DeleteConfiguration(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var CreateCollectionEnvironmentID string
var CreateCollectionName string
var CreateCollectionDescription string
var CreateCollectionConfigurationID string
var CreateCollectionLanguage string

func getCreateCollectionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-collection",
		Short: "Create a collection",
		Long: "",
		Run: CreateCollection,
	}

	cmd.Flags().StringVarP(&CreateCollectionEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&CreateCollectionName, "name", "", "", "The name of the collection to be created.")
	cmd.Flags().StringVarP(&CreateCollectionDescription, "description", "", "", "A description of the collection.")
	cmd.Flags().StringVarP(&CreateCollectionConfigurationID, "configuration_id", "", "", "The ID of the configuration in which the collection is to be created.")
	cmd.Flags().StringVarP(&CreateCollectionLanguage, "language", "", "", "The language of the documents stored in the collection, in the form of an ISO 639-1 language code.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("version")

	return cmd
}

func CreateCollection(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.CreateCollectionOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(CreateCollectionEnvironmentID)
		}
		if flag.Name == "name" {
			optionsModel.SetName(CreateCollectionName)
		}
		if flag.Name == "description" {
			optionsModel.SetDescription(CreateCollectionDescription)
		}
		if flag.Name == "configuration_id" {
			optionsModel.SetConfigurationID(CreateCollectionConfigurationID)
		}
		if flag.Name == "language" {
			optionsModel.SetLanguage(CreateCollectionLanguage)
		}
	})

	result, _, responseErr := discovery.CreateCollection(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ListCollectionsEnvironmentID string
var ListCollectionsName string

func getListCollectionsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-collections",
		Short: "List collections",
		Long: "Lists existing collections for the service instance.",
		Run: ListCollections,
	}

	cmd.Flags().StringVarP(&ListCollectionsEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&ListCollectionsName, "name", "", "", "Find collections with the given name.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func ListCollections(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.ListCollectionsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(ListCollectionsEnvironmentID)
		}
		if flag.Name == "name" {
			optionsModel.SetName(ListCollectionsName)
		}
	})

	result, _, responseErr := discovery.ListCollections(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetCollectionEnvironmentID string
var GetCollectionCollectionID string

func getGetCollectionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-collection",
		Short: "Get collection details",
		Long: "",
		Run: GetCollection,
	}

	cmd.Flags().StringVarP(&GetCollectionEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&GetCollectionCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetCollection(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.GetCollectionOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(GetCollectionEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(GetCollectionCollectionID)
		}
	})

	result, _, responseErr := discovery.GetCollection(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var UpdateCollectionEnvironmentID string
var UpdateCollectionCollectionID string
var UpdateCollectionName string
var UpdateCollectionDescription string
var UpdateCollectionConfigurationID string

func getUpdateCollectionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-collection",
		Short: "Update a collection",
		Long: "",
		Run: UpdateCollection,
	}

	cmd.Flags().StringVarP(&UpdateCollectionEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&UpdateCollectionCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&UpdateCollectionName, "name", "", "", "The name of the collection.")
	cmd.Flags().StringVarP(&UpdateCollectionDescription, "description", "", "", "A description of the collection.")
	cmd.Flags().StringVarP(&UpdateCollectionConfigurationID, "configuration_id", "", "", "The ID of the configuration in which the collection is to be updated.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func UpdateCollection(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.UpdateCollectionOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(UpdateCollectionEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(UpdateCollectionCollectionID)
		}
		if flag.Name == "name" {
			optionsModel.SetName(UpdateCollectionName)
		}
		if flag.Name == "description" {
			optionsModel.SetDescription(UpdateCollectionDescription)
		}
		if flag.Name == "configuration_id" {
			optionsModel.SetConfigurationID(UpdateCollectionConfigurationID)
		}
	})

	result, _, responseErr := discovery.UpdateCollection(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteCollectionEnvironmentID string
var DeleteCollectionCollectionID string

func getDeleteCollectionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-collection",
		Short: "Delete a collection",
		Long: "",
		Run: DeleteCollection,
	}

	cmd.Flags().StringVarP(&DeleteCollectionEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&DeleteCollectionCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteCollection(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.DeleteCollectionOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(DeleteCollectionEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(DeleteCollectionCollectionID)
		}
	})

	result, _, responseErr := discovery.DeleteCollection(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ListCollectionFieldsEnvironmentID string
var ListCollectionFieldsCollectionID string

func getListCollectionFieldsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-collection-fields",
		Short: "List collection fields",
		Long: "Gets a list of the unique fields (and their types) stored in the index.",
		Run: ListCollectionFields,
	}

	cmd.Flags().StringVarP(&ListCollectionFieldsEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&ListCollectionFieldsCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func ListCollectionFields(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.ListCollectionFieldsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(ListCollectionFieldsEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(ListCollectionFieldsCollectionID)
		}
	})

	result, _, responseErr := discovery.ListCollectionFields(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ListExpansionsEnvironmentID string
var ListExpansionsCollectionID string

func getListExpansionsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-expansions",
		Short: "Get the expansion list",
		Long: "Returns the current expansion list for the specified collection. If an expansion list is not specified, an object with empty expansion arrays is returned.",
		Run: ListExpansions,
	}

	cmd.Flags().StringVarP(&ListExpansionsEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&ListExpansionsCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func ListExpansions(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.ListExpansionsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(ListExpansionsEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(ListExpansionsCollectionID)
		}
	})

	result, _, responseErr := discovery.ListExpansions(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var CreateExpansionsEnvironmentID string
var CreateExpansionsCollectionID string
var CreateExpansionsExpansions string

func getCreateExpansionsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-expansions",
		Short: "Create or update expansion list",
		Long: "Create or replace the Expansion list for this collection. The maximum number of expanded terms per collection is `500`. The current expansion list is replaced with the uploaded content.",
		Run: CreateExpansions,
	}

	cmd.Flags().StringVarP(&CreateExpansionsEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&CreateExpansionsCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&CreateExpansionsExpansions, "expansions", "", "", "An array of query expansion definitions.  Each object in the **expansions** array represents a term or set of terms that will be expanded into other terms. Each expansion object can be configured as bidirectional or unidirectional. Bidirectional means that all terms are expanded to all other terms in the object. Unidirectional means that a set list of terms can be expanded into a second list of terms. To create a bi-directional expansion specify an **expanded_terms** array. When found in a query, all items in the **expanded_terms** array are then expanded to the other items in the same array. To create a uni-directional expansion, specify both an array of **input_terms** and an array of **expanded_terms**. When items in the **input_terms** array are present in a query, they are expanded using the items listed in the **expanded_terms** array.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("expansions")
	cmd.MarkFlagRequired("version")

	return cmd
}

func CreateExpansions(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.CreateExpansionsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(CreateExpansionsEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(CreateExpansionsCollectionID)
		}
		if flag.Name == "expansions" {
			var expansions []discoveryv1.Expansion
			decodeErr := json.Unmarshal([]byte(CreateExpansionsExpansions), &expansions);
			utils.HandleError(decodeErr)

			optionsModel.SetExpansions(expansions)
		}
	})

	result, _, responseErr := discovery.CreateExpansions(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteExpansionsEnvironmentID string
var DeleteExpansionsCollectionID string

func getDeleteExpansionsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-expansions",
		Short: "Delete the expansion list",
		Long: "Remove the expansion information for this collection. The expansion list must be deleted to disable query expansion for a collection.",
		Run: DeleteExpansions,
	}

	cmd.Flags().StringVarP(&DeleteExpansionsEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&DeleteExpansionsCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteExpansions(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.DeleteExpansionsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(DeleteExpansionsEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(DeleteExpansionsCollectionID)
		}
	})

	_, responseErr := discovery.DeleteExpansions(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var GetTokenizationDictionaryStatusEnvironmentID string
var GetTokenizationDictionaryStatusCollectionID string

func getGetTokenizationDictionaryStatusCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-tokenization-dictionary-status",
		Short: "Get tokenization dictionary status",
		Long: "Returns the current status of the tokenization dictionary for the specified collection.",
		Run: GetTokenizationDictionaryStatus,
	}

	cmd.Flags().StringVarP(&GetTokenizationDictionaryStatusEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&GetTokenizationDictionaryStatusCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetTokenizationDictionaryStatus(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.GetTokenizationDictionaryStatusOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(GetTokenizationDictionaryStatusEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(GetTokenizationDictionaryStatusCollectionID)
		}
	})

	result, _, responseErr := discovery.GetTokenizationDictionaryStatus(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var CreateTokenizationDictionaryEnvironmentID string
var CreateTokenizationDictionaryCollectionID string
var CreateTokenizationDictionaryTokenizationRules string

func getCreateTokenizationDictionaryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-tokenization-dictionary",
		Short: "Create tokenization dictionary",
		Long: "Upload a custom tokenization dictionary to use with the specified collection.",
		Run: CreateTokenizationDictionary,
	}

	cmd.Flags().StringVarP(&CreateTokenizationDictionaryEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&CreateTokenizationDictionaryCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&CreateTokenizationDictionaryTokenizationRules, "tokenization_rules", "", "", "An array of tokenization rules. Each rule contains, the original `text` string, component `tokens`, any alternate character set `readings`, and which `part_of_speech` the text is from.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func CreateTokenizationDictionary(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.CreateTokenizationDictionaryOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(CreateTokenizationDictionaryEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(CreateTokenizationDictionaryCollectionID)
		}
		if flag.Name == "tokenization_rules" {
			var tokenization_rules []discoveryv1.TokenDictRule
			decodeErr := json.Unmarshal([]byte(CreateTokenizationDictionaryTokenizationRules), &tokenization_rules);
			utils.HandleError(decodeErr)

			optionsModel.SetTokenizationRules(tokenization_rules)
		}
	})

	result, _, responseErr := discovery.CreateTokenizationDictionary(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteTokenizationDictionaryEnvironmentID string
var DeleteTokenizationDictionaryCollectionID string

func getDeleteTokenizationDictionaryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-tokenization-dictionary",
		Short: "Delete tokenization dictionary",
		Long: "Delete the tokenization dictionary from the collection.",
		Run: DeleteTokenizationDictionary,
	}

	cmd.Flags().StringVarP(&DeleteTokenizationDictionaryEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&DeleteTokenizationDictionaryCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteTokenizationDictionary(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.DeleteTokenizationDictionaryOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(DeleteTokenizationDictionaryEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(DeleteTokenizationDictionaryCollectionID)
		}
	})

	_, responseErr := discovery.DeleteTokenizationDictionary(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var GetStopwordListStatusEnvironmentID string
var GetStopwordListStatusCollectionID string

func getGetStopwordListStatusCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-stopword-list-status",
		Short: "Get stopword list status",
		Long: "Returns the current status of the stopword list for the specified collection.",
		Run: GetStopwordListStatus,
	}

	cmd.Flags().StringVarP(&GetStopwordListStatusEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&GetStopwordListStatusCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetStopwordListStatus(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.GetStopwordListStatusOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(GetStopwordListStatusEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(GetStopwordListStatusCollectionID)
		}
	})

	result, _, responseErr := discovery.GetStopwordListStatus(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var CreateStopwordListEnvironmentID string
var CreateStopwordListCollectionID string
var CreateStopwordListStopwordFile string
var CreateStopwordListStopwordFilename string

func getCreateStopwordListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-stopword-list",
		Short: "Create stopword list",
		Long: "Upload a custom stopword list to use with the specified collection.",
		Run: CreateStopwordList,
	}

	cmd.Flags().StringVarP(&CreateStopwordListEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&CreateStopwordListCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&CreateStopwordListStopwordFile, "stopword_file", "", "", "The content of the stopword list to ingest.")
	cmd.Flags().StringVarP(&CreateStopwordListStopwordFilename, "stopword_filename", "", "", "The filename for StopwordFile.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("stopword_file")
	cmd.MarkFlagRequired("stopword_filename")
	cmd.MarkFlagRequired("version")

	return cmd
}

func CreateStopwordList(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.CreateStopwordListOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(CreateStopwordListEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(CreateStopwordListCollectionID)
		}
		if flag.Name == "stopword_file" {
			stopword_file, fileErr := os.Open(CreateStopwordListStopwordFile)
			utils.HandleError(fileErr)

			optionsModel.SetStopwordFile(stopword_file)
		}
		if flag.Name == "stopword_filename" {
			optionsModel.SetStopwordFilename(CreateStopwordListStopwordFilename)
		}
	})

	result, _, responseErr := discovery.CreateStopwordList(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteStopwordListEnvironmentID string
var DeleteStopwordListCollectionID string

func getDeleteStopwordListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-stopword-list",
		Short: "Delete a custom stopword list",
		Long: "Delete a custom stopword list from the collection. After a custom stopword list is deleted, the default list is used for the collection.",
		Run: DeleteStopwordList,
	}

	cmd.Flags().StringVarP(&DeleteStopwordListEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&DeleteStopwordListCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteStopwordList(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.DeleteStopwordListOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(DeleteStopwordListEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(DeleteStopwordListCollectionID)
		}
	})

	_, responseErr := discovery.DeleteStopwordList(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var AddDocumentEnvironmentID string
var AddDocumentCollectionID string
var AddDocumentFile string
var AddDocumentFilename string
var AddDocumentFileContentType string
var AddDocumentMetadata string

func getAddDocumentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "add-document",
		Short: "Add a document",
		Long: "Add a document to a collection with optional metadata.  * The **version** query parameter is still required.  * Returns immediately after the system has accepted the document for processing.  * The user must provide document content, metadata, or both. If the request is missing both document content and metadata, it is rejected.  * The user can set the **Content-Type** parameter on the **file** part to indicate the media type of the document. If the **Content-Type** parameter is missing or is one of the generic media types (for example, `application/octet-stream`), then the service attempts to automatically detect the document's media type.  * The following field names are reserved and will be filtered out if present after normalization: `id`, `score`, `highlight`, and any field with the prefix of: `_`, `+`, or `-`  * Fields with empty name values after normalization are filtered out before indexing.  * Fields containing the following characters after normalization are filtered out before indexing: `#` and `,` **Note:** Documents can be added with a specific **document_id** by using the **_/v1/environments/{environment_id}/collections/{collection_id}/documents** method.",
		Run: AddDocument,
	}

	cmd.Flags().StringVarP(&AddDocumentEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&AddDocumentCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&AddDocumentFile, "file", "", "", "The content of the document to ingest. The maximum supported file size when adding a file to a collection is 50 megabytes, the maximum supported file size when testing a confiruration is 1 megabyte. Files larger than the supported size are rejected.")
	cmd.Flags().StringVarP(&AddDocumentFilename, "filename", "", "", "The filename for File.")
	cmd.Flags().StringVarP(&AddDocumentFileContentType, "file_content_type", "", "", "The content type of File.")
	cmd.Flags().StringVarP(&AddDocumentMetadata, "metadata", "", "", "The maximum supported metadata file size is 1 MB. Metadata parts larger than 1 MB are rejected. Example:  ``` {  'Creator': 'Johnny Appleseed',  'Subject': 'Apples'} ```.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func AddDocument(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.AddDocumentOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(AddDocumentEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(AddDocumentCollectionID)
		}
		if flag.Name == "file" {
			file, fileErr := os.Open(AddDocumentFile)
			utils.HandleError(fileErr)

			optionsModel.SetFile(file)
		}
		if flag.Name == "filename" {
			optionsModel.SetFilename(AddDocumentFilename)
		}
		if flag.Name == "file_content_type" {
			optionsModel.SetFileContentType(AddDocumentFileContentType)
		}
		if flag.Name == "metadata" {
			optionsModel.SetMetadata(AddDocumentMetadata)
		}
	})

	result, _, responseErr := discovery.AddDocument(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetDocumentStatusEnvironmentID string
var GetDocumentStatusCollectionID string
var GetDocumentStatusDocumentID string

func getGetDocumentStatusCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-document-status",
		Short: "Get document details",
		Long: "Fetch status details about a submitted document. **Note:** this operation does not return the document itself. Instead, it returns only the document's processing status and any notices (warnings or errors) that were generated when the document was ingested. Use the query API to retrieve the actual document content.",
		Run: GetDocumentStatus,
	}

	cmd.Flags().StringVarP(&GetDocumentStatusEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&GetDocumentStatusCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&GetDocumentStatusDocumentID, "document_id", "", "", "The ID of the document.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("document_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetDocumentStatus(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.GetDocumentStatusOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(GetDocumentStatusEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(GetDocumentStatusCollectionID)
		}
		if flag.Name == "document_id" {
			optionsModel.SetDocumentID(GetDocumentStatusDocumentID)
		}
	})

	result, _, responseErr := discovery.GetDocumentStatus(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var UpdateDocumentEnvironmentID string
var UpdateDocumentCollectionID string
var UpdateDocumentDocumentID string
var UpdateDocumentFile string
var UpdateDocumentFilename string
var UpdateDocumentFileContentType string
var UpdateDocumentMetadata string

func getUpdateDocumentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-document",
		Short: "Update a document",
		Long: "Replace an existing document or add a document with a specified **document_id**. Starts ingesting a document with optional metadata.**Note:** When uploading a new document with this method it automatically replaces any document stored with the same **document_id** if it exists.",
		Run: UpdateDocument,
	}

	cmd.Flags().StringVarP(&UpdateDocumentEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&UpdateDocumentCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&UpdateDocumentDocumentID, "document_id", "", "", "The ID of the document.")
	cmd.Flags().StringVarP(&UpdateDocumentFile, "file", "", "", "The content of the document to ingest. The maximum supported file size when adding a file to a collection is 50 megabytes, the maximum supported file size when testing a confiruration is 1 megabyte. Files larger than the supported size are rejected.")
	cmd.Flags().StringVarP(&UpdateDocumentFilename, "filename", "", "", "The filename for File.")
	cmd.Flags().StringVarP(&UpdateDocumentFileContentType, "file_content_type", "", "", "The content type of File.")
	cmd.Flags().StringVarP(&UpdateDocumentMetadata, "metadata", "", "", "The maximum supported metadata file size is 1 MB. Metadata parts larger than 1 MB are rejected. Example:  ``` {  'Creator': 'Johnny Appleseed',  'Subject': 'Apples'} ```.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("document_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func UpdateDocument(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.UpdateDocumentOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(UpdateDocumentEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(UpdateDocumentCollectionID)
		}
		if flag.Name == "document_id" {
			optionsModel.SetDocumentID(UpdateDocumentDocumentID)
		}
		if flag.Name == "file" {
			file, fileErr := os.Open(UpdateDocumentFile)
			utils.HandleError(fileErr)

			optionsModel.SetFile(file)
		}
		if flag.Name == "filename" {
			optionsModel.SetFilename(UpdateDocumentFilename)
		}
		if flag.Name == "file_content_type" {
			optionsModel.SetFileContentType(UpdateDocumentFileContentType)
		}
		if flag.Name == "metadata" {
			optionsModel.SetMetadata(UpdateDocumentMetadata)
		}
	})

	result, _, responseErr := discovery.UpdateDocument(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteDocumentEnvironmentID string
var DeleteDocumentCollectionID string
var DeleteDocumentDocumentID string

func getDeleteDocumentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-document",
		Short: "Delete a document",
		Long: "If the given document ID is invalid, or if the document is not found, then the a success response is returned (HTTP status code `200`) with the status set to 'deleted'.",
		Run: DeleteDocument,
	}

	cmd.Flags().StringVarP(&DeleteDocumentEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&DeleteDocumentCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&DeleteDocumentDocumentID, "document_id", "", "", "The ID of the document.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("document_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteDocument(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.DeleteDocumentOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(DeleteDocumentEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(DeleteDocumentCollectionID)
		}
		if flag.Name == "document_id" {
			optionsModel.SetDocumentID(DeleteDocumentDocumentID)
		}
	})

	result, _, responseErr := discovery.DeleteDocument(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var QueryEnvironmentID string
var QueryCollectionID string
var QueryFilter string
var QueryQuery string
var QueryNaturalLanguageQuery string
var QueryPassages bool
var QueryAggregation string
var QueryCount int64
var QueryReturn string
var QueryOffset int64
var QuerySort string
var QueryHighlight bool
var QueryPassagesFields string
var QueryPassagesCount int64
var QueryPassagesCharacters int64
var QueryDeduplicate bool
var QueryDeduplicateField string
var QuerySimilar bool
var QuerySimilarDocumentIds string
var QuerySimilarFields string
var QueryBias string
var QuerySpellingSuggestions bool
var QueryXWatsonLoggingOptOut bool

func getQueryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "query",
		Short: "Query a collection",
		Long: "By using this method, you can construct long queries. For details, see the [Discovery documentation](https://cloud.ibm.com/docs/services/discovery?topic=discovery-query-concepts#query-concepts).",
		Run: Query,
	}

	cmd.Flags().StringVarP(&QueryEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&QueryCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&QueryFilter, "filter", "", "", "A cacheable query that excludes documents that don't mention the query content. Filter searches are better for metadata-type searches and for assessing the concepts in the data set.")
	cmd.Flags().StringVarP(&QueryQuery, "query", "", "", "A query search returns all documents in your data set with full enrichments and full text, but with the most relevant documents listed first. Use a query search when you want to find the most relevant search results.")
	cmd.Flags().StringVarP(&QueryNaturalLanguageQuery, "natural_language_query", "", "", "A natural language query that returns relevant documents by utilizing training data and natural language understanding.")
	cmd.Flags().BoolVarP(&QueryPassages, "passages", "", false, "A passages query that returns the most relevant passages from the results.")
	cmd.Flags().StringVarP(&QueryAggregation, "aggregation", "", "", "An aggregation search that returns an exact answer by combining query search with filters. Useful for applications to build lists, tables, and time series. For a full list of possible aggregations, see the Query reference.")
	cmd.Flags().Int64VarP(&QueryCount, "count", "", 0, "Number of results to return.")
	cmd.Flags().StringVarP(&QueryReturn, "return", "", "", "A comma-separated list of the portion of the document hierarchy to return.")
	cmd.Flags().Int64VarP(&QueryOffset, "offset", "", 0, "The number of query results to skip at the beginning. For example, if the total number of results that are returned is 10 and the offset is 8, it returns the last two results.")
	cmd.Flags().StringVarP(&QuerySort, "sort", "", "", "A comma-separated list of fields in the document to sort on. You can optionally specify a sort direction by prefixing the field with `-` for descending or `+` for ascending. Ascending is the default sort direction if no prefix is specified. This parameter cannot be used in the same query as the **bias** parameter.")
	cmd.Flags().BoolVarP(&QueryHighlight, "highlight", "", false, "When true, a highlight field is returned for each result which contains the fields which match the query with `<em></em>` tags around the matching query terms.")
	cmd.Flags().StringVarP(&QueryPassagesFields, "passages_fields", "", "", "A comma-separated list of fields that passages are drawn from. If this parameter not specified, then all top-level fields are included.")
	cmd.Flags().Int64VarP(&QueryPassagesCount, "passages_count", "", 0, "The maximum number of passages to return. The search returns fewer passages if the requested total is not found. The default is `10`. The maximum is `100`.")
	cmd.Flags().Int64VarP(&QueryPassagesCharacters, "passages_characters", "", 0, "The approximate number of characters that any one passage will have.")
	cmd.Flags().BoolVarP(&QueryDeduplicate, "deduplicate", "", false, "When `true`, and used with a Watson Discovery News collection, duplicate results (based on the contents of the **title** field) are removed. Duplicate comparison is limited to the current query only; **offset** is not considered. This parameter is currently Beta functionality.")
	cmd.Flags().StringVarP(&QueryDeduplicateField, "deduplicate_field", "", "", "When specified, duplicate results based on the field specified are removed from the returned results. Duplicate comparison is limited to the current query only, **offset** is not considered. This parameter is currently Beta functionality.")
	cmd.Flags().BoolVarP(&QuerySimilar, "similar", "", false, "When `true`, results are returned based on their similarity to the document IDs specified in the **similar.document_ids** parameter.")
	cmd.Flags().StringVarP(&QuerySimilarDocumentIds, "similar_document_ids", "", "", "A comma-separated list of document IDs to find similar documents.**Tip:** Include the **natural_language_query** parameter to expand the scope of the document similarity search with the natural language query. Other query parameters, such as **filter** and **query**, are subsequently applied and reduce the scope.")
	cmd.Flags().StringVarP(&QuerySimilarFields, "similar_fields", "", "", "A comma-separated list of field names that are used as a basis for comparison to identify similar documents. If not specified, the entire document is used for comparison.")
	cmd.Flags().StringVarP(&QueryBias, "bias", "", "", "Field which the returned results will be biased against. The specified field must be either a **date** or **number** format. When a **date** type field is specified returned results are biased towards field values closer to the current date. When a **number** type field is specified, returned results are biased towards higher field values. This parameter cannot be used in the same query as the **sort** parameter.")
	cmd.Flags().BoolVarP(&QuerySpellingSuggestions, "spelling_suggestions", "", false, "When `true` and the **natural_language_query** parameter is used, the **natural_languge_query** parameter is spell checked. The most likely correction is retunred in the **suggested_query** field of the response (if one exists). **Important:** this parameter is only valid when using the Cloud Pak version of Discovery.")
	cmd.Flags().BoolVarP(&QueryXWatsonLoggingOptOut, "x_watson_logging_opt_out", "", false, "If `true`, queries are not stored in the Discovery **Logs** endpoint.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func Query(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.QueryOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(QueryEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(QueryCollectionID)
		}
		if flag.Name == "filter" {
			optionsModel.SetFilter(QueryFilter)
		}
		if flag.Name == "query" {
			optionsModel.SetQuery(QueryQuery)
		}
		if flag.Name == "natural_language_query" {
			optionsModel.SetNaturalLanguageQuery(QueryNaturalLanguageQuery)
		}
		if flag.Name == "passages" {
			optionsModel.SetPassages(QueryPassages)
		}
		if flag.Name == "aggregation" {
			optionsModel.SetAggregation(QueryAggregation)
		}
		if flag.Name == "count" {
			optionsModel.SetCount(QueryCount)
		}
		if flag.Name == "return" {
			optionsModel.SetReturn(QueryReturn)
		}
		if flag.Name == "offset" {
			optionsModel.SetOffset(QueryOffset)
		}
		if flag.Name == "sort" {
			optionsModel.SetSort(QuerySort)
		}
		if flag.Name == "highlight" {
			optionsModel.SetHighlight(QueryHighlight)
		}
		if flag.Name == "passages_fields" {
			optionsModel.SetPassagesFields(QueryPassagesFields)
		}
		if flag.Name == "passages_count" {
			optionsModel.SetPassagesCount(QueryPassagesCount)
		}
		if flag.Name == "passages_characters" {
			optionsModel.SetPassagesCharacters(QueryPassagesCharacters)
		}
		if flag.Name == "deduplicate" {
			optionsModel.SetDeduplicate(QueryDeduplicate)
		}
		if flag.Name == "deduplicate_field" {
			optionsModel.SetDeduplicateField(QueryDeduplicateField)
		}
		if flag.Name == "similar" {
			optionsModel.SetSimilar(QuerySimilar)
		}
		if flag.Name == "similar_document_ids" {
			optionsModel.SetSimilarDocumentIds(QuerySimilarDocumentIds)
		}
		if flag.Name == "similar_fields" {
			optionsModel.SetSimilarFields(QuerySimilarFields)
		}
		if flag.Name == "bias" {
			optionsModel.SetBias(QueryBias)
		}
		if flag.Name == "spelling_suggestions" {
			optionsModel.SetSpellingSuggestions(QuerySpellingSuggestions)
		}
		if flag.Name == "x_watson_logging_opt_out" {
			optionsModel.SetXWatsonLoggingOptOut(QueryXWatsonLoggingOptOut)
		}
	})

	result, _, responseErr := discovery.Query(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var QueryNoticesEnvironmentID string
var QueryNoticesCollectionID string
var QueryNoticesFilter string
var QueryNoticesQuery string
var QueryNoticesNaturalLanguageQuery string
var QueryNoticesPassages bool
var QueryNoticesAggregation string
var QueryNoticesCount int64
var QueryNoticesReturn []string
var QueryNoticesOffset int64
var QueryNoticesSort []string
var QueryNoticesHighlight bool
var QueryNoticesPassagesFields []string
var QueryNoticesPassagesCount int64
var QueryNoticesPassagesCharacters int64
var QueryNoticesDeduplicateField string
var QueryNoticesSimilar bool
var QueryNoticesSimilarDocumentIds []string
var QueryNoticesSimilarFields []string

func getQueryNoticesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "query-notices",
		Short: "Query system notices",
		Long: "Queries for notices (errors or warnings) that might have been generated by the system. Notices are generated when ingesting documents and performing relevance training. See the [Discovery documentation](https://cloud.ibm.com/docs/services/discovery?topic=discovery-query-concepts#query-concepts) for more details on the query language.",
		Run: QueryNotices,
	}

	cmd.Flags().StringVarP(&QueryNoticesEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&QueryNoticesCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&QueryNoticesFilter, "filter", "", "", "A cacheable query that excludes documents that don't mention the query content. Filter searches are better for metadata-type searches and for assessing the concepts in the data set.")
	cmd.Flags().StringVarP(&QueryNoticesQuery, "query", "", "", "A query search returns all documents in your data set with full enrichments and full text, but with the most relevant documents listed first.")
	cmd.Flags().StringVarP(&QueryNoticesNaturalLanguageQuery, "natural_language_query", "", "", "A natural language query that returns relevant documents by utilizing training data and natural language understanding.")
	cmd.Flags().BoolVarP(&QueryNoticesPassages, "passages", "", false, "A passages query that returns the most relevant passages from the results.")
	cmd.Flags().StringVarP(&QueryNoticesAggregation, "aggregation", "", "", "An aggregation search that returns an exact answer by combining query search with filters. Useful for applications to build lists, tables, and time series. For a full list of possible aggregations, see the Query reference.")
	cmd.Flags().Int64VarP(&QueryNoticesCount, "count", "", 0, "Number of results to return. The maximum for the **count** and **offset** values together in any one query is **10000**.")
	cmd.Flags().StringSliceVarP(&QueryNoticesReturn, "return", "", nil, "A comma-separated list of the portion of the document hierarchy to return.")
	cmd.Flags().Int64VarP(&QueryNoticesOffset, "offset", "", 0, "The number of query results to skip at the beginning. For example, if the total number of results that are returned is 10 and the offset is 8, it returns the last two results. The maximum for the **count** and **offset** values together in any one query is **10000**.")
	cmd.Flags().StringSliceVarP(&QueryNoticesSort, "sort", "", nil, "A comma-separated list of fields in the document to sort on. You can optionally specify a sort direction by prefixing the field with `-` for descending or `+` for ascending. Ascending is the default sort direction if no prefix is specified.")
	cmd.Flags().BoolVarP(&QueryNoticesHighlight, "highlight", "", false, "When true, a highlight field is returned for each result which contains the fields which match the query with `<em></em>` tags around the matching query terms.")
	cmd.Flags().StringSliceVarP(&QueryNoticesPassagesFields, "passages_fields", "", nil, "A comma-separated list of fields that passages are drawn from. If this parameter not specified, then all top-level fields are included.")
	cmd.Flags().Int64VarP(&QueryNoticesPassagesCount, "passages_count", "", 0, "The maximum number of passages to return. The search returns fewer passages if the requested total is not found.")
	cmd.Flags().Int64VarP(&QueryNoticesPassagesCharacters, "passages_characters", "", 0, "The approximate number of characters that any one passage will have.")
	cmd.Flags().StringVarP(&QueryNoticesDeduplicateField, "deduplicate_field", "", "", "When specified, duplicate results based on the field specified are removed from the returned results. Duplicate comparison is limited to the current query only, **offset** is not considered. This parameter is currently Beta functionality.")
	cmd.Flags().BoolVarP(&QueryNoticesSimilar, "similar", "", false, "When `true`, results are returned based on their similarity to the document IDs specified in the **similar.document_ids** parameter.")
	cmd.Flags().StringSliceVarP(&QueryNoticesSimilarDocumentIds, "similar_document_ids", "", nil, "A comma-separated list of document IDs to find similar documents.**Tip:** Include the **natural_language_query** parameter to expand the scope of the document similarity search with the natural language query. Other query parameters, such as **filter** and **query**, are subsequently applied and reduce the scope.")
	cmd.Flags().StringSliceVarP(&QueryNoticesSimilarFields, "similar_fields", "", nil, "A comma-separated list of field names that are used as a basis for comparison to identify similar documents. If not specified, the entire document is used for comparison.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func QueryNotices(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.QueryNoticesOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(QueryNoticesEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(QueryNoticesCollectionID)
		}
		if flag.Name == "filter" {
			optionsModel.SetFilter(QueryNoticesFilter)
		}
		if flag.Name == "query" {
			optionsModel.SetQuery(QueryNoticesQuery)
		}
		if flag.Name == "natural_language_query" {
			optionsModel.SetNaturalLanguageQuery(QueryNoticesNaturalLanguageQuery)
		}
		if flag.Name == "passages" {
			optionsModel.SetPassages(QueryNoticesPassages)
		}
		if flag.Name == "aggregation" {
			optionsModel.SetAggregation(QueryNoticesAggregation)
		}
		if flag.Name == "count" {
			optionsModel.SetCount(QueryNoticesCount)
		}
		if flag.Name == "return" {
			optionsModel.SetReturn(QueryNoticesReturn)
		}
		if flag.Name == "offset" {
			optionsModel.SetOffset(QueryNoticesOffset)
		}
		if flag.Name == "sort" {
			optionsModel.SetSort(QueryNoticesSort)
		}
		if flag.Name == "highlight" {
			optionsModel.SetHighlight(QueryNoticesHighlight)
		}
		if flag.Name == "passages_fields" {
			optionsModel.SetPassagesFields(QueryNoticesPassagesFields)
		}
		if flag.Name == "passages_count" {
			optionsModel.SetPassagesCount(QueryNoticesPassagesCount)
		}
		if flag.Name == "passages_characters" {
			optionsModel.SetPassagesCharacters(QueryNoticesPassagesCharacters)
		}
		if flag.Name == "deduplicate_field" {
			optionsModel.SetDeduplicateField(QueryNoticesDeduplicateField)
		}
		if flag.Name == "similar" {
			optionsModel.SetSimilar(QueryNoticesSimilar)
		}
		if flag.Name == "similar_document_ids" {
			optionsModel.SetSimilarDocumentIds(QueryNoticesSimilarDocumentIds)
		}
		if flag.Name == "similar_fields" {
			optionsModel.SetSimilarFields(QueryNoticesSimilarFields)
		}
	})

	result, _, responseErr := discovery.QueryNotices(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var FederatedQueryEnvironmentID string
var FederatedQueryFilter string
var FederatedQueryQuery string
var FederatedQueryNaturalLanguageQuery string
var FederatedQueryPassages bool
var FederatedQueryAggregation string
var FederatedQueryCount int64
var FederatedQueryReturn string
var FederatedQueryOffset int64
var FederatedQuerySort string
var FederatedQueryHighlight bool
var FederatedQueryPassagesFields string
var FederatedQueryPassagesCount int64
var FederatedQueryPassagesCharacters int64
var FederatedQueryDeduplicate bool
var FederatedQueryDeduplicateField string
var FederatedQuerySimilar bool
var FederatedQuerySimilarDocumentIds string
var FederatedQuerySimilarFields string
var FederatedQueryBias string
var FederatedQueryCollectionIds string
var FederatedQueryXWatsonLoggingOptOut bool

func getFederatedQueryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "federated-query",
		Short: "Query multiple collections",
		Long: "By using this method, you can construct long queries that search multiple collection. For details, see the [Discovery documentation](https://cloud.ibm.com/docs/services/discovery?topic=discovery-query-concepts#query-concepts).",
		Run: FederatedQuery,
	}

	cmd.Flags().StringVarP(&FederatedQueryEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&FederatedQueryFilter, "filter", "", "", "A cacheable query that excludes documents that don't mention the query content. Filter searches are better for metadata-type searches and for assessing the concepts in the data set.")
	cmd.Flags().StringVarP(&FederatedQueryQuery, "query", "", "", "A query search returns all documents in your data set with full enrichments and full text, but with the most relevant documents listed first. Use a query search when you want to find the most relevant search results.")
	cmd.Flags().StringVarP(&FederatedQueryNaturalLanguageQuery, "natural_language_query", "", "", "A natural language query that returns relevant documents by utilizing training data and natural language understanding.")
	cmd.Flags().BoolVarP(&FederatedQueryPassages, "passages", "", false, "A passages query that returns the most relevant passages from the results.")
	cmd.Flags().StringVarP(&FederatedQueryAggregation, "aggregation", "", "", "An aggregation search that returns an exact answer by combining query search with filters. Useful for applications to build lists, tables, and time series. For a full list of possible aggregations, see the Query reference.")
	cmd.Flags().Int64VarP(&FederatedQueryCount, "count", "", 0, "Number of results to return.")
	cmd.Flags().StringVarP(&FederatedQueryReturn, "return", "", "", "A comma-separated list of the portion of the document hierarchy to return.")
	cmd.Flags().Int64VarP(&FederatedQueryOffset, "offset", "", 0, "The number of query results to skip at the beginning. For example, if the total number of results that are returned is 10 and the offset is 8, it returns the last two results.")
	cmd.Flags().StringVarP(&FederatedQuerySort, "sort", "", "", "A comma-separated list of fields in the document to sort on. You can optionally specify a sort direction by prefixing the field with `-` for descending or `+` for ascending. Ascending is the default sort direction if no prefix is specified. This parameter cannot be used in the same query as the **bias** parameter.")
	cmd.Flags().BoolVarP(&FederatedQueryHighlight, "highlight", "", false, "When true, a highlight field is returned for each result which contains the fields which match the query with `<em></em>` tags around the matching query terms.")
	cmd.Flags().StringVarP(&FederatedQueryPassagesFields, "passages_fields", "", "", "A comma-separated list of fields that passages are drawn from. If this parameter not specified, then all top-level fields are included.")
	cmd.Flags().Int64VarP(&FederatedQueryPassagesCount, "passages_count", "", 0, "The maximum number of passages to return. The search returns fewer passages if the requested total is not found. The default is `10`. The maximum is `100`.")
	cmd.Flags().Int64VarP(&FederatedQueryPassagesCharacters, "passages_characters", "", 0, "The approximate number of characters that any one passage will have.")
	cmd.Flags().BoolVarP(&FederatedQueryDeduplicate, "deduplicate", "", false, "When `true`, and used with a Watson Discovery News collection, duplicate results (based on the contents of the **title** field) are removed. Duplicate comparison is limited to the current query only; **offset** is not considered. This parameter is currently Beta functionality.")
	cmd.Flags().StringVarP(&FederatedQueryDeduplicateField, "deduplicate_field", "", "", "When specified, duplicate results based on the field specified are removed from the returned results. Duplicate comparison is limited to the current query only, **offset** is not considered. This parameter is currently Beta functionality.")
	cmd.Flags().BoolVarP(&FederatedQuerySimilar, "similar", "", false, "When `true`, results are returned based on their similarity to the document IDs specified in the **similar.document_ids** parameter.")
	cmd.Flags().StringVarP(&FederatedQuerySimilarDocumentIds, "similar_document_ids", "", "", "A comma-separated list of document IDs to find similar documents.**Tip:** Include the **natural_language_query** parameter to expand the scope of the document similarity search with the natural language query. Other query parameters, such as **filter** and **query**, are subsequently applied and reduce the scope.")
	cmd.Flags().StringVarP(&FederatedQuerySimilarFields, "similar_fields", "", "", "A comma-separated list of field names that are used as a basis for comparison to identify similar documents. If not specified, the entire document is used for comparison.")
	cmd.Flags().StringVarP(&FederatedQueryBias, "bias", "", "", "Field which the returned results will be biased against. The specified field must be either a **date** or **number** format. When a **date** type field is specified returned results are biased towards field values closer to the current date. When a **number** type field is specified, returned results are biased towards higher field values. This parameter cannot be used in the same query as the **sort** parameter.")
	cmd.Flags().StringVarP(&FederatedQueryCollectionIds, "collection_ids", "", "", "A comma-separated list of collection IDs to be queried against.")
	cmd.Flags().BoolVarP(&FederatedQueryXWatsonLoggingOptOut, "x_watson_logging_opt_out", "", false, "If `true`, queries are not stored in the Discovery **Logs** endpoint.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func FederatedQuery(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.FederatedQueryOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(FederatedQueryEnvironmentID)
		}
		if flag.Name == "filter" {
			optionsModel.SetFilter(FederatedQueryFilter)
		}
		if flag.Name == "query" {
			optionsModel.SetQuery(FederatedQueryQuery)
		}
		if flag.Name == "natural_language_query" {
			optionsModel.SetNaturalLanguageQuery(FederatedQueryNaturalLanguageQuery)
		}
		if flag.Name == "passages" {
			optionsModel.SetPassages(FederatedQueryPassages)
		}
		if flag.Name == "aggregation" {
			optionsModel.SetAggregation(FederatedQueryAggregation)
		}
		if flag.Name == "count" {
			optionsModel.SetCount(FederatedQueryCount)
		}
		if flag.Name == "return" {
			optionsModel.SetReturn(FederatedQueryReturn)
		}
		if flag.Name == "offset" {
			optionsModel.SetOffset(FederatedQueryOffset)
		}
		if flag.Name == "sort" {
			optionsModel.SetSort(FederatedQuerySort)
		}
		if flag.Name == "highlight" {
			optionsModel.SetHighlight(FederatedQueryHighlight)
		}
		if flag.Name == "passages_fields" {
			optionsModel.SetPassagesFields(FederatedQueryPassagesFields)
		}
		if flag.Name == "passages_count" {
			optionsModel.SetPassagesCount(FederatedQueryPassagesCount)
		}
		if flag.Name == "passages_characters" {
			optionsModel.SetPassagesCharacters(FederatedQueryPassagesCharacters)
		}
		if flag.Name == "deduplicate" {
			optionsModel.SetDeduplicate(FederatedQueryDeduplicate)
		}
		if flag.Name == "deduplicate_field" {
			optionsModel.SetDeduplicateField(FederatedQueryDeduplicateField)
		}
		if flag.Name == "similar" {
			optionsModel.SetSimilar(FederatedQuerySimilar)
		}
		if flag.Name == "similar_document_ids" {
			optionsModel.SetSimilarDocumentIds(FederatedQuerySimilarDocumentIds)
		}
		if flag.Name == "similar_fields" {
			optionsModel.SetSimilarFields(FederatedQuerySimilarFields)
		}
		if flag.Name == "bias" {
			optionsModel.SetBias(FederatedQueryBias)
		}
		if flag.Name == "collection_ids" {
			optionsModel.SetCollectionIds(FederatedQueryCollectionIds)
		}
		if flag.Name == "x_watson_logging_opt_out" {
			optionsModel.SetXWatsonLoggingOptOut(FederatedQueryXWatsonLoggingOptOut)
		}
	})

	result, _, responseErr := discovery.FederatedQuery(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var FederatedQueryNoticesEnvironmentID string
var FederatedQueryNoticesCollectionIds []string
var FederatedQueryNoticesFilter string
var FederatedQueryNoticesQuery string
var FederatedQueryNoticesNaturalLanguageQuery string
var FederatedQueryNoticesAggregation string
var FederatedQueryNoticesCount int64
var FederatedQueryNoticesReturn []string
var FederatedQueryNoticesOffset int64
var FederatedQueryNoticesSort []string
var FederatedQueryNoticesHighlight bool
var FederatedQueryNoticesDeduplicateField string
var FederatedQueryNoticesSimilar bool
var FederatedQueryNoticesSimilarDocumentIds []string
var FederatedQueryNoticesSimilarFields []string

func getFederatedQueryNoticesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "federated-query-notices",
		Short: "Query multiple collection system notices",
		Long: "Queries for notices (errors or warnings) that might have been generated by the system. Notices are generated when ingesting documents and performing relevance training. See the [Discovery documentation](https://cloud.ibm.com/docs/services/discovery?topic=discovery-query-concepts#query-concepts) for more details on the query language.",
		Run: FederatedQueryNotices,
	}

	cmd.Flags().StringVarP(&FederatedQueryNoticesEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringSliceVarP(&FederatedQueryNoticesCollectionIds, "collection_ids", "", nil, "A comma-separated list of collection IDs to be queried against.")
	cmd.Flags().StringVarP(&FederatedQueryNoticesFilter, "filter", "", "", "A cacheable query that excludes documents that don't mention the query content. Filter searches are better for metadata-type searches and for assessing the concepts in the data set.")
	cmd.Flags().StringVarP(&FederatedQueryNoticesQuery, "query", "", "", "A query search returns all documents in your data set with full enrichments and full text, but with the most relevant documents listed first.")
	cmd.Flags().StringVarP(&FederatedQueryNoticesNaturalLanguageQuery, "natural_language_query", "", "", "A natural language query that returns relevant documents by utilizing training data and natural language understanding.")
	cmd.Flags().StringVarP(&FederatedQueryNoticesAggregation, "aggregation", "", "", "An aggregation search that returns an exact answer by combining query search with filters. Useful for applications to build lists, tables, and time series. For a full list of possible aggregations, see the Query reference.")
	cmd.Flags().Int64VarP(&FederatedQueryNoticesCount, "count", "", 0, "Number of results to return. The maximum for the **count** and **offset** values together in any one query is **10000**.")
	cmd.Flags().StringSliceVarP(&FederatedQueryNoticesReturn, "return", "", nil, "A comma-separated list of the portion of the document hierarchy to return.")
	cmd.Flags().Int64VarP(&FederatedQueryNoticesOffset, "offset", "", 0, "The number of query results to skip at the beginning. For example, if the total number of results that are returned is 10 and the offset is 8, it returns the last two results. The maximum for the **count** and **offset** values together in any one query is **10000**.")
	cmd.Flags().StringSliceVarP(&FederatedQueryNoticesSort, "sort", "", nil, "A comma-separated list of fields in the document to sort on. You can optionally specify a sort direction by prefixing the field with `-` for descending or `+` for ascending. Ascending is the default sort direction if no prefix is specified.")
	cmd.Flags().BoolVarP(&FederatedQueryNoticesHighlight, "highlight", "", false, "When true, a highlight field is returned for each result which contains the fields which match the query with `<em></em>` tags around the matching query terms.")
	cmd.Flags().StringVarP(&FederatedQueryNoticesDeduplicateField, "deduplicate_field", "", "", "When specified, duplicate results based on the field specified are removed from the returned results. Duplicate comparison is limited to the current query only, **offset** is not considered. This parameter is currently Beta functionality.")
	cmd.Flags().BoolVarP(&FederatedQueryNoticesSimilar, "similar", "", false, "When `true`, results are returned based on their similarity to the document IDs specified in the **similar.document_ids** parameter.")
	cmd.Flags().StringSliceVarP(&FederatedQueryNoticesSimilarDocumentIds, "similar_document_ids", "", nil, "A comma-separated list of document IDs to find similar documents.**Tip:** Include the **natural_language_query** parameter to expand the scope of the document similarity search with the natural language query. Other query parameters, such as **filter** and **query**, are subsequently applied and reduce the scope.")
	cmd.Flags().StringSliceVarP(&FederatedQueryNoticesSimilarFields, "similar_fields", "", nil, "A comma-separated list of field names that are used as a basis for comparison to identify similar documents. If not specified, the entire document is used for comparison.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_ids")
	cmd.MarkFlagRequired("version")

	return cmd
}

func FederatedQueryNotices(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.FederatedQueryNoticesOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(FederatedQueryNoticesEnvironmentID)
		}
		if flag.Name == "collection_ids" {
			optionsModel.SetCollectionIds(FederatedQueryNoticesCollectionIds)
		}
		if flag.Name == "filter" {
			optionsModel.SetFilter(FederatedQueryNoticesFilter)
		}
		if flag.Name == "query" {
			optionsModel.SetQuery(FederatedQueryNoticesQuery)
		}
		if flag.Name == "natural_language_query" {
			optionsModel.SetNaturalLanguageQuery(FederatedQueryNoticesNaturalLanguageQuery)
		}
		if flag.Name == "aggregation" {
			optionsModel.SetAggregation(FederatedQueryNoticesAggregation)
		}
		if flag.Name == "count" {
			optionsModel.SetCount(FederatedQueryNoticesCount)
		}
		if flag.Name == "return" {
			optionsModel.SetReturn(FederatedQueryNoticesReturn)
		}
		if flag.Name == "offset" {
			optionsModel.SetOffset(FederatedQueryNoticesOffset)
		}
		if flag.Name == "sort" {
			optionsModel.SetSort(FederatedQueryNoticesSort)
		}
		if flag.Name == "highlight" {
			optionsModel.SetHighlight(FederatedQueryNoticesHighlight)
		}
		if flag.Name == "deduplicate_field" {
			optionsModel.SetDeduplicateField(FederatedQueryNoticesDeduplicateField)
		}
		if flag.Name == "similar" {
			optionsModel.SetSimilar(FederatedQueryNoticesSimilar)
		}
		if flag.Name == "similar_document_ids" {
			optionsModel.SetSimilarDocumentIds(FederatedQueryNoticesSimilarDocumentIds)
		}
		if flag.Name == "similar_fields" {
			optionsModel.SetSimilarFields(FederatedQueryNoticesSimilarFields)
		}
	})

	result, _, responseErr := discovery.FederatedQueryNotices(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetAutocompletionEnvironmentID string
var GetAutocompletionCollectionID string
var GetAutocompletionField string
var GetAutocompletionPrefix string
var GetAutocompletionCount int64

func getGetAutocompletionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-autocompletion",
		Short: "Get Autocomplete Suggestions",
		Long: "Returns completion query suggestions for the specified prefix.  /n/n **Important:** this method is only valid when using the Cloud Pak version of Discovery.",
		Run: GetAutocompletion,
	}

	cmd.Flags().StringVarP(&GetAutocompletionEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&GetAutocompletionCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&GetAutocompletionField, "field", "", "", "The field in the result documents that autocompletion suggestions are identified from.")
	cmd.Flags().StringVarP(&GetAutocompletionPrefix, "prefix", "", "", "The prefix to use for autocompletion. For example, the prefix `Ho` could autocomplete to `Hot`, `Housing`, or `How do I upgrade`. Possible completions are.")
	cmd.Flags().Int64VarP(&GetAutocompletionCount, "count", "", 0, "The number of autocompletion suggestions to return.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetAutocompletion(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.GetAutocompletionOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(GetAutocompletionEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(GetAutocompletionCollectionID)
		}
		if flag.Name == "field" {
			optionsModel.SetField(GetAutocompletionField)
		}
		if flag.Name == "prefix" {
			optionsModel.SetPrefix(GetAutocompletionPrefix)
		}
		if flag.Name == "count" {
			optionsModel.SetCount(GetAutocompletionCount)
		}
	})

	result, _, responseErr := discovery.GetAutocompletion(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ListTrainingDataEnvironmentID string
var ListTrainingDataCollectionID string

func getListTrainingDataCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-training-data",
		Short: "List training data",
		Long: "Lists the training data for the specified collection.",
		Run: ListTrainingData,
	}

	cmd.Flags().StringVarP(&ListTrainingDataEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&ListTrainingDataCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func ListTrainingData(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.ListTrainingDataOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(ListTrainingDataEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(ListTrainingDataCollectionID)
		}
	})

	result, _, responseErr := discovery.ListTrainingData(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var AddTrainingDataEnvironmentID string
var AddTrainingDataCollectionID string
var AddTrainingDataNaturalLanguageQuery string
var AddTrainingDataFilter string
var AddTrainingDataExamples string

func getAddTrainingDataCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "add-training-data",
		Short: "Add query to training data",
		Long: "Adds a query to the training data for this collection. The query can contain a filter and natural language query.",
		Run: AddTrainingData,
	}

	cmd.Flags().StringVarP(&AddTrainingDataEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&AddTrainingDataCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&AddTrainingDataNaturalLanguageQuery, "natural_language_query", "", "", "The natural text query for the new training query.")
	cmd.Flags().StringVarP(&AddTrainingDataFilter, "filter", "", "", "The filter used on the collection before the **natural_language_query** is applied.")
	cmd.Flags().StringVarP(&AddTrainingDataExamples, "examples", "", "", "Array of training examples.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func AddTrainingData(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.AddTrainingDataOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(AddTrainingDataEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(AddTrainingDataCollectionID)
		}
		if flag.Name == "natural_language_query" {
			optionsModel.SetNaturalLanguageQuery(AddTrainingDataNaturalLanguageQuery)
		}
		if flag.Name == "filter" {
			optionsModel.SetFilter(AddTrainingDataFilter)
		}
		if flag.Name == "examples" {
			var examples []discoveryv1.TrainingExample
			decodeErr := json.Unmarshal([]byte(AddTrainingDataExamples), &examples);
			utils.HandleError(decodeErr)

			optionsModel.SetExamples(examples)
		}
	})

	result, _, responseErr := discovery.AddTrainingData(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteAllTrainingDataEnvironmentID string
var DeleteAllTrainingDataCollectionID string

func getDeleteAllTrainingDataCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-all-training-data",
		Short: "Delete all training data",
		Long: "Deletes all training data from a collection.",
		Run: DeleteAllTrainingData,
	}

	cmd.Flags().StringVarP(&DeleteAllTrainingDataEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&DeleteAllTrainingDataCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteAllTrainingData(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.DeleteAllTrainingDataOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(DeleteAllTrainingDataEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(DeleteAllTrainingDataCollectionID)
		}
	})

	_, responseErr := discovery.DeleteAllTrainingData(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var GetTrainingDataEnvironmentID string
var GetTrainingDataCollectionID string
var GetTrainingDataQueryID string

func getGetTrainingDataCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-training-data",
		Short: "Get details about a query",
		Long: "Gets details for a specific training data query, including the query string and all examples.",
		Run: GetTrainingData,
	}

	cmd.Flags().StringVarP(&GetTrainingDataEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&GetTrainingDataCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&GetTrainingDataQueryID, "query_id", "", "", "The ID of the query used for training.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("query_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetTrainingData(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.GetTrainingDataOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(GetTrainingDataEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(GetTrainingDataCollectionID)
		}
		if flag.Name == "query_id" {
			optionsModel.SetQueryID(GetTrainingDataQueryID)
		}
	})

	result, _, responseErr := discovery.GetTrainingData(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteTrainingDataEnvironmentID string
var DeleteTrainingDataCollectionID string
var DeleteTrainingDataQueryID string

func getDeleteTrainingDataCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-training-data",
		Short: "Delete a training data query",
		Long: "Removes the training data query and all associated examples from the training data set.",
		Run: DeleteTrainingData,
	}

	cmd.Flags().StringVarP(&DeleteTrainingDataEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&DeleteTrainingDataCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&DeleteTrainingDataQueryID, "query_id", "", "", "The ID of the query used for training.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("query_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteTrainingData(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.DeleteTrainingDataOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(DeleteTrainingDataEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(DeleteTrainingDataCollectionID)
		}
		if flag.Name == "query_id" {
			optionsModel.SetQueryID(DeleteTrainingDataQueryID)
		}
	})

	_, responseErr := discovery.DeleteTrainingData(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var ListTrainingExamplesEnvironmentID string
var ListTrainingExamplesCollectionID string
var ListTrainingExamplesQueryID string

func getListTrainingExamplesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-training-examples",
		Short: "List examples for a training data query",
		Long: "List all examples for this training data query.",
		Run: ListTrainingExamples,
	}

	cmd.Flags().StringVarP(&ListTrainingExamplesEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&ListTrainingExamplesCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&ListTrainingExamplesQueryID, "query_id", "", "", "The ID of the query used for training.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("query_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func ListTrainingExamples(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.ListTrainingExamplesOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(ListTrainingExamplesEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(ListTrainingExamplesCollectionID)
		}
		if flag.Name == "query_id" {
			optionsModel.SetQueryID(ListTrainingExamplesQueryID)
		}
	})

	result, _, responseErr := discovery.ListTrainingExamples(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var CreateTrainingExampleEnvironmentID string
var CreateTrainingExampleCollectionID string
var CreateTrainingExampleQueryID string
var CreateTrainingExampleDocumentID string
var CreateTrainingExampleCrossReference string
var CreateTrainingExampleRelevance int64

func getCreateTrainingExampleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-training-example",
		Short: "Add example to training data query",
		Long: "Adds a example to this training data query.",
		Run: CreateTrainingExample,
	}

	cmd.Flags().StringVarP(&CreateTrainingExampleEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&CreateTrainingExampleCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&CreateTrainingExampleQueryID, "query_id", "", "", "The ID of the query used for training.")
	cmd.Flags().StringVarP(&CreateTrainingExampleDocumentID, "document_id", "", "", "The document ID associated with this training example.")
	cmd.Flags().StringVarP(&CreateTrainingExampleCrossReference, "cross_reference", "", "", "The cross reference associated with this training example.")
	cmd.Flags().Int64VarP(&CreateTrainingExampleRelevance, "relevance", "", 0, "The relevance of the training example.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("query_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func CreateTrainingExample(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.CreateTrainingExampleOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(CreateTrainingExampleEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(CreateTrainingExampleCollectionID)
		}
		if flag.Name == "query_id" {
			optionsModel.SetQueryID(CreateTrainingExampleQueryID)
		}
		if flag.Name == "document_id" {
			optionsModel.SetDocumentID(CreateTrainingExampleDocumentID)
		}
		if flag.Name == "cross_reference" {
			optionsModel.SetCrossReference(CreateTrainingExampleCrossReference)
		}
		if flag.Name == "relevance" {
			optionsModel.SetRelevance(CreateTrainingExampleRelevance)
		}
	})

	result, _, responseErr := discovery.CreateTrainingExample(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteTrainingExampleEnvironmentID string
var DeleteTrainingExampleCollectionID string
var DeleteTrainingExampleQueryID string
var DeleteTrainingExampleExampleID string

func getDeleteTrainingExampleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-training-example",
		Short: "Delete example for training data query",
		Long: "Deletes the example document with the given ID from the training data query.",
		Run: DeleteTrainingExample,
	}

	cmd.Flags().StringVarP(&DeleteTrainingExampleEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&DeleteTrainingExampleCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&DeleteTrainingExampleQueryID, "query_id", "", "", "The ID of the query used for training.")
	cmd.Flags().StringVarP(&DeleteTrainingExampleExampleID, "example_id", "", "", "The ID of the document as it is indexed.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("query_id")
	cmd.MarkFlagRequired("example_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteTrainingExample(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.DeleteTrainingExampleOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(DeleteTrainingExampleEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(DeleteTrainingExampleCollectionID)
		}
		if flag.Name == "query_id" {
			optionsModel.SetQueryID(DeleteTrainingExampleQueryID)
		}
		if flag.Name == "example_id" {
			optionsModel.SetExampleID(DeleteTrainingExampleExampleID)
		}
	})

	_, responseErr := discovery.DeleteTrainingExample(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var UpdateTrainingExampleEnvironmentID string
var UpdateTrainingExampleCollectionID string
var UpdateTrainingExampleQueryID string
var UpdateTrainingExampleExampleID string
var UpdateTrainingExampleCrossReference string
var UpdateTrainingExampleRelevance int64

func getUpdateTrainingExampleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-training-example",
		Short: "Change label or cross reference for example",
		Long: "Changes the label or cross reference query for this training data example.",
		Run: UpdateTrainingExample,
	}

	cmd.Flags().StringVarP(&UpdateTrainingExampleEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&UpdateTrainingExampleCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&UpdateTrainingExampleQueryID, "query_id", "", "", "The ID of the query used for training.")
	cmd.Flags().StringVarP(&UpdateTrainingExampleExampleID, "example_id", "", "", "The ID of the document as it is indexed.")
	cmd.Flags().StringVarP(&UpdateTrainingExampleCrossReference, "cross_reference", "", "", "The example to add.")
	cmd.Flags().Int64VarP(&UpdateTrainingExampleRelevance, "relevance", "", 0, "The relevance value for this example.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("query_id")
	cmd.MarkFlagRequired("example_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func UpdateTrainingExample(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.UpdateTrainingExampleOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(UpdateTrainingExampleEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(UpdateTrainingExampleCollectionID)
		}
		if flag.Name == "query_id" {
			optionsModel.SetQueryID(UpdateTrainingExampleQueryID)
		}
		if flag.Name == "example_id" {
			optionsModel.SetExampleID(UpdateTrainingExampleExampleID)
		}
		if flag.Name == "cross_reference" {
			optionsModel.SetCrossReference(UpdateTrainingExampleCrossReference)
		}
		if flag.Name == "relevance" {
			optionsModel.SetRelevance(UpdateTrainingExampleRelevance)
		}
	})

	result, _, responseErr := discovery.UpdateTrainingExample(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetTrainingExampleEnvironmentID string
var GetTrainingExampleCollectionID string
var GetTrainingExampleQueryID string
var GetTrainingExampleExampleID string

func getGetTrainingExampleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-training-example",
		Short: "Get details for training data example",
		Long: "Gets the details for this training example.",
		Run: GetTrainingExample,
	}

	cmd.Flags().StringVarP(&GetTrainingExampleEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&GetTrainingExampleCollectionID, "collection_id", "", "", "The ID of the collection.")
	cmd.Flags().StringVarP(&GetTrainingExampleQueryID, "query_id", "", "", "The ID of the query used for training.")
	cmd.Flags().StringVarP(&GetTrainingExampleExampleID, "example_id", "", "", "The ID of the document as it is indexed.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("collection_id")
	cmd.MarkFlagRequired("query_id")
	cmd.MarkFlagRequired("example_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetTrainingExample(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.GetTrainingExampleOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(GetTrainingExampleEnvironmentID)
		}
		if flag.Name == "collection_id" {
			optionsModel.SetCollectionID(GetTrainingExampleCollectionID)
		}
		if flag.Name == "query_id" {
			optionsModel.SetQueryID(GetTrainingExampleQueryID)
		}
		if flag.Name == "example_id" {
			optionsModel.SetExampleID(GetTrainingExampleExampleID)
		}
	})

	result, _, responseErr := discovery.GetTrainingExample(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteUserDataCustomerID string

func getDeleteUserDataCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-user-data",
		Short: "Delete labeled data",
		Long: "Deletes all data associated with a specified customer ID. The method has no effect if no data is associated with the customer ID. You associate a customer ID with data by passing the **X-Watson-Metadata** header with a request that passes data. For more information about personal data and customer IDs, see [Information security](https://cloud.ibm.com/docs/services/discovery?topic=discovery-information-security#information-security).",
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
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.DeleteUserDataOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "customer_id" {
			optionsModel.SetCustomerID(DeleteUserDataCustomerID)
		}
	})

	_, responseErr := discovery.DeleteUserData(&optionsModel)
	utils.HandleError(responseErr)

	ui.Ok()
}

var CreateEventType string
var CreateEventData string

func getCreateEventCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-event",
		Short: "Create event",
		Long: "The **Events** API can be used to create log entries that are associated with specific queries. For example, you can record which documents in the results set were 'clicked' by a user and when that click occured.",
		Run: CreateEvent,
	}

	cmd.Flags().StringVarP(&CreateEventType, "type", "", "", "The event type to be created.")
	cmd.Flags().StringVarP(&CreateEventData, "data", "", "", "Query event data object.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("type")
	cmd.MarkFlagRequired("data")
	cmd.MarkFlagRequired("version")

	return cmd
}

func CreateEvent(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.CreateEventOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "type" {
			optionsModel.SetType(CreateEventType)
		}
		if flag.Name == "data" {
			var data discoveryv1.EventData
			decodeErr := json.Unmarshal([]byte(CreateEventData), &data);
			utils.HandleError(decodeErr)

			optionsModel.SetData(&data)
		}
	})

	result, _, responseErr := discovery.CreateEvent(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var QueryLogFilter string
var QueryLogQuery string
var QueryLogCount int64
var QueryLogOffset int64
var QueryLogSort []string

func getQueryLogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "query-log",
		Short: "Search the query and event log",
		Long: "Searches the query and event log to find query sessions that match the specified criteria. Searching the **logs** endpoint uses the standard Discovery query syntax for the parameters that are supported.",
		Run: QueryLog,
	}

	cmd.Flags().StringVarP(&QueryLogFilter, "filter", "", "", "A cacheable query that excludes documents that don't mention the query content. Filter searches are better for metadata-type searches and for assessing the concepts in the data set.")
	cmd.Flags().StringVarP(&QueryLogQuery, "query", "", "", "A query search returns all documents in your data set with full enrichments and full text, but with the most relevant documents listed first.")
	cmd.Flags().Int64VarP(&QueryLogCount, "count", "", 0, "Number of results to return. The maximum for the **count** and **offset** values together in any one query is **10000**.")
	cmd.Flags().Int64VarP(&QueryLogOffset, "offset", "", 0, "The number of query results to skip at the beginning. For example, if the total number of results that are returned is 10 and the offset is 8, it returns the last two results. The maximum for the **count** and **offset** values together in any one query is **10000**.")
	cmd.Flags().StringSliceVarP(&QueryLogSort, "sort", "", nil, "A comma-separated list of fields in the document to sort on. You can optionally specify a sort direction by prefixing the field with `-` for descending or `+` for ascending. Ascending is the default sort direction if no prefix is specified.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("version")

	return cmd
}

func QueryLog(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.QueryLogOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "filter" {
			optionsModel.SetFilter(QueryLogFilter)
		}
		if flag.Name == "query" {
			optionsModel.SetQuery(QueryLogQuery)
		}
		if flag.Name == "count" {
			optionsModel.SetCount(QueryLogCount)
		}
		if flag.Name == "offset" {
			optionsModel.SetOffset(QueryLogOffset)
		}
		if flag.Name == "sort" {
			optionsModel.SetSort(QueryLogSort)
		}
	})

	result, _, responseErr := discovery.QueryLog(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetMetricsQueryStartTime string
var GetMetricsQueryEndTime string
var GetMetricsQueryResultType string

func getGetMetricsQueryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-metrics-query",
		Short: "Number of queries over time",
		Long: "Total number of queries using the **natural_language_query** parameter over a specific time window.",
		Run: GetMetricsQuery,
	}

	cmd.Flags().StringVarP(&GetMetricsQueryStartTime, "start_time", "", "", "Metric is computed from data recorded after this timestamp; must be in `YYYY-MM-DDThh:mm:ssZ` format.")
	cmd.Flags().StringVarP(&GetMetricsQueryEndTime, "end_time", "", "", "Metric is computed from data recorded before this timestamp; must be in `YYYY-MM-DDThh:mm:ssZ` format.")
	cmd.Flags().StringVarP(&GetMetricsQueryResultType, "result_type", "", "", "The type of result to consider when calculating the metric.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("version")

	return cmd
}

func GetMetricsQuery(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.GetMetricsQueryOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "start_time" {
			start_time, dateTimeParseErr := strfmt.ParseDateTime(GetMetricsQueryStartTime)
			utils.HandleError(dateTimeParseErr)

			optionsModel.SetStartTime(&start_time)
		}
		if flag.Name == "end_time" {
			end_time, dateTimeParseErr := strfmt.ParseDateTime(GetMetricsQueryEndTime)
			utils.HandleError(dateTimeParseErr)

			optionsModel.SetEndTime(&end_time)
		}
		if flag.Name == "result_type" {
			optionsModel.SetResultType(GetMetricsQueryResultType)
		}
	})

	result, _, responseErr := discovery.GetMetricsQuery(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetMetricsQueryEventStartTime string
var GetMetricsQueryEventEndTime string
var GetMetricsQueryEventResultType string

func getGetMetricsQueryEventCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-metrics-query-event",
		Short: "Number of queries with an event over time",
		Long: "Total number of queries using the **natural_language_query** parameter that have a corresponding 'click' event over a specified time window. This metric requires having integrated event tracking in your application using the **Events** API.",
		Run: GetMetricsQueryEvent,
	}

	cmd.Flags().StringVarP(&GetMetricsQueryEventStartTime, "start_time", "", "", "Metric is computed from data recorded after this timestamp; must be in `YYYY-MM-DDThh:mm:ssZ` format.")
	cmd.Flags().StringVarP(&GetMetricsQueryEventEndTime, "end_time", "", "", "Metric is computed from data recorded before this timestamp; must be in `YYYY-MM-DDThh:mm:ssZ` format.")
	cmd.Flags().StringVarP(&GetMetricsQueryEventResultType, "result_type", "", "", "The type of result to consider when calculating the metric.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("version")

	return cmd
}

func GetMetricsQueryEvent(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.GetMetricsQueryEventOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "start_time" {
			start_time, dateTimeParseErr := strfmt.ParseDateTime(GetMetricsQueryEventStartTime)
			utils.HandleError(dateTimeParseErr)

			optionsModel.SetStartTime(&start_time)
		}
		if flag.Name == "end_time" {
			end_time, dateTimeParseErr := strfmt.ParseDateTime(GetMetricsQueryEventEndTime)
			utils.HandleError(dateTimeParseErr)

			optionsModel.SetEndTime(&end_time)
		}
		if flag.Name == "result_type" {
			optionsModel.SetResultType(GetMetricsQueryEventResultType)
		}
	})

	result, _, responseErr := discovery.GetMetricsQueryEvent(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetMetricsQueryNoResultsStartTime string
var GetMetricsQueryNoResultsEndTime string
var GetMetricsQueryNoResultsResultType string

func getGetMetricsQueryNoResultsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-metrics-query-no-results",
		Short: "Number of queries with no search results over time",
		Long: "Total number of queries using the **natural_language_query** parameter that have no results returned over a specified time window.",
		Run: GetMetricsQueryNoResults,
	}

	cmd.Flags().StringVarP(&GetMetricsQueryNoResultsStartTime, "start_time", "", "", "Metric is computed from data recorded after this timestamp; must be in `YYYY-MM-DDThh:mm:ssZ` format.")
	cmd.Flags().StringVarP(&GetMetricsQueryNoResultsEndTime, "end_time", "", "", "Metric is computed from data recorded before this timestamp; must be in `YYYY-MM-DDThh:mm:ssZ` format.")
	cmd.Flags().StringVarP(&GetMetricsQueryNoResultsResultType, "result_type", "", "", "The type of result to consider when calculating the metric.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("version")

	return cmd
}

func GetMetricsQueryNoResults(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.GetMetricsQueryNoResultsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "start_time" {
			start_time, dateTimeParseErr := strfmt.ParseDateTime(GetMetricsQueryNoResultsStartTime)
			utils.HandleError(dateTimeParseErr)

			optionsModel.SetStartTime(&start_time)
		}
		if flag.Name == "end_time" {
			end_time, dateTimeParseErr := strfmt.ParseDateTime(GetMetricsQueryNoResultsEndTime)
			utils.HandleError(dateTimeParseErr)

			optionsModel.SetEndTime(&end_time)
		}
		if flag.Name == "result_type" {
			optionsModel.SetResultType(GetMetricsQueryNoResultsResultType)
		}
	})

	result, _, responseErr := discovery.GetMetricsQueryNoResults(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetMetricsEventRateStartTime string
var GetMetricsEventRateEndTime string
var GetMetricsEventRateResultType string

func getGetMetricsEventRateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-metrics-event-rate",
		Short: "Percentage of queries with an associated event",
		Long: "The percentage of queries using the **natural_language_query** parameter that have a corresponding 'click' event over a specified time window.  This metric requires having integrated event tracking in your application using the **Events** API.",
		Run: GetMetricsEventRate,
	}

	cmd.Flags().StringVarP(&GetMetricsEventRateStartTime, "start_time", "", "", "Metric is computed from data recorded after this timestamp; must be in `YYYY-MM-DDThh:mm:ssZ` format.")
	cmd.Flags().StringVarP(&GetMetricsEventRateEndTime, "end_time", "", "", "Metric is computed from data recorded before this timestamp; must be in `YYYY-MM-DDThh:mm:ssZ` format.")
	cmd.Flags().StringVarP(&GetMetricsEventRateResultType, "result_type", "", "", "The type of result to consider when calculating the metric.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("version")

	return cmd
}

func GetMetricsEventRate(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.GetMetricsEventRateOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "start_time" {
			start_time, dateTimeParseErr := strfmt.ParseDateTime(GetMetricsEventRateStartTime)
			utils.HandleError(dateTimeParseErr)

			optionsModel.SetStartTime(&start_time)
		}
		if flag.Name == "end_time" {
			end_time, dateTimeParseErr := strfmt.ParseDateTime(GetMetricsEventRateEndTime)
			utils.HandleError(dateTimeParseErr)

			optionsModel.SetEndTime(&end_time)
		}
		if flag.Name == "result_type" {
			optionsModel.SetResultType(GetMetricsEventRateResultType)
		}
	})

	result, _, responseErr := discovery.GetMetricsEventRate(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetMetricsQueryTokenEventCount int64

func getGetMetricsQueryTokenEventCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-metrics-query-token-event",
		Short: "Most frequent query tokens with an event",
		Long: "The most frequent query tokens parsed from the **natural_language_query** parameter and their corresponding 'click' event rate within the recording period (queries and events are stored for 30 days). A query token is an individual word or unigram within the query string.",
		Run: GetMetricsQueryTokenEvent,
	}

	cmd.Flags().Int64VarP(&GetMetricsQueryTokenEventCount, "count", "", 0, "Number of results to return. The maximum for the **count** and **offset** values together in any one query is **10000**.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("version")

	return cmd
}

func GetMetricsQueryTokenEvent(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.GetMetricsQueryTokenEventOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "count" {
			optionsModel.SetCount(GetMetricsQueryTokenEventCount)
		}
	})

	result, _, responseErr := discovery.GetMetricsQueryTokenEvent(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ListCredentialsEnvironmentID string

func getListCredentialsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-credentials",
		Short: "List credentials",
		Long: "List all the source credentials that have been created for this service instance. **Note:**  All credentials are sent over an encrypted connection and encrypted at rest.",
		Run: ListCredentials,
	}

	cmd.Flags().StringVarP(&ListCredentialsEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func ListCredentials(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.ListCredentialsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(ListCredentialsEnvironmentID)
		}
	})

	result, _, responseErr := discovery.ListCredentials(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var CreateCredentialsEnvironmentID string
var CreateCredentialsSourceType string
var CreateCredentialsCredentialDetails string
var CreateCredentialsStatus string

func getCreateCredentialsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-credentials",
		Short: "Create credentials",
		Long: "Creates a set of credentials to connect to a remote source. Created credentials are used in a configuration to associate a collection with the remote source.**Note:** All credentials are sent over an encrypted connection and encrypted at rest.",
		Run: CreateCredentials,
	}

	cmd.Flags().StringVarP(&CreateCredentialsEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&CreateCredentialsSourceType, "source_type", "", "", "The source that this credentials object connects to.-  `box` indicates the credentials are used to connect an instance of Enterprise Box.-  `salesforce` indicates the credentials are used to connect to Salesforce.-  `sharepoint` indicates the credentials are used to connect to Microsoft SharePoint Online.-  `web_crawl` indicates the credentials are used to perform a web crawl.=  `cloud_object_storage` indicates the credentials are used to connect to an IBM Cloud Object Store.")
	cmd.Flags().StringVarP(&CreateCredentialsCredentialDetails, "credential_details", "", "", "Object containing details of the stored credentials. Obtain credentials for your source from the administrator of the source.")
	cmd.Flags().StringVarP(&CreateCredentialsStatus, "status", "", "", "The current status of this set of credentials. `connected` indicates that the credentials are available to use with the source configuration of a collection. `invalid` refers to the credentials (for example, the password provided has expired) and must be corrected before they can be used with a collection.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func CreateCredentials(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.CreateCredentialsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(CreateCredentialsEnvironmentID)
		}
		if flag.Name == "source_type" {
			optionsModel.SetSourceType(CreateCredentialsSourceType)
		}
		if flag.Name == "credential_details" {
			var credential_details discoveryv1.CredentialDetails
			decodeErr := json.Unmarshal([]byte(CreateCredentialsCredentialDetails), &credential_details);
			utils.HandleError(decodeErr)

			optionsModel.SetCredentialDetails(&credential_details)
		}
		if flag.Name == "status" {
			optionsModel.SetStatus(CreateCredentialsStatus)
		}
	})

	result, _, responseErr := discovery.CreateCredentials(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetCredentialsEnvironmentID string
var GetCredentialsCredentialID string

func getGetCredentialsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-credentials",
		Short: "View Credentials",
		Long: "Returns details about the specified credentials. **Note:** Secure credential information such as a password or SSH key is never returned and must be obtained from the source system.",
		Run: GetCredentials,
	}

	cmd.Flags().StringVarP(&GetCredentialsEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&GetCredentialsCredentialID, "credential_id", "", "", "The unique identifier for a set of source credentials.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("credential_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetCredentials(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.GetCredentialsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(GetCredentialsEnvironmentID)
		}
		if flag.Name == "credential_id" {
			optionsModel.SetCredentialID(GetCredentialsCredentialID)
		}
	})

	result, _, responseErr := discovery.GetCredentials(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var UpdateCredentialsEnvironmentID string
var UpdateCredentialsCredentialID string
var UpdateCredentialsSourceType string
var UpdateCredentialsCredentialDetails string
var UpdateCredentialsStatus string

func getUpdateCredentialsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-credentials",
		Short: "Update credentials",
		Long: "Updates an existing set of source credentials.**Note:** All credentials are sent over an encrypted connection and encrypted at rest.",
		Run: UpdateCredentials,
	}

	cmd.Flags().StringVarP(&UpdateCredentialsEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&UpdateCredentialsCredentialID, "credential_id", "", "", "The unique identifier for a set of source credentials.")
	cmd.Flags().StringVarP(&UpdateCredentialsSourceType, "source_type", "", "", "The source that this credentials object connects to.-  `box` indicates the credentials are used to connect an instance of Enterprise Box.-  `salesforce` indicates the credentials are used to connect to Salesforce.-  `sharepoint` indicates the credentials are used to connect to Microsoft SharePoint Online.-  `web_crawl` indicates the credentials are used to perform a web crawl.=  `cloud_object_storage` indicates the credentials are used to connect to an IBM Cloud Object Store.")
	cmd.Flags().StringVarP(&UpdateCredentialsCredentialDetails, "credential_details", "", "", "Object containing details of the stored credentials. Obtain credentials for your source from the administrator of the source.")
	cmd.Flags().StringVarP(&UpdateCredentialsStatus, "status", "", "", "The current status of this set of credentials. `connected` indicates that the credentials are available to use with the source configuration of a collection. `invalid` refers to the credentials (for example, the password provided has expired) and must be corrected before they can be used with a collection.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("credential_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func UpdateCredentials(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.UpdateCredentialsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(UpdateCredentialsEnvironmentID)
		}
		if flag.Name == "credential_id" {
			optionsModel.SetCredentialID(UpdateCredentialsCredentialID)
		}
		if flag.Name == "source_type" {
			optionsModel.SetSourceType(UpdateCredentialsSourceType)
		}
		if flag.Name == "credential_details" {
			var credential_details discoveryv1.CredentialDetails
			decodeErr := json.Unmarshal([]byte(UpdateCredentialsCredentialDetails), &credential_details);
			utils.HandleError(decodeErr)

			optionsModel.SetCredentialDetails(&credential_details)
		}
		if flag.Name == "status" {
			optionsModel.SetStatus(UpdateCredentialsStatus)
		}
	})

	result, _, responseErr := discovery.UpdateCredentials(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteCredentialsEnvironmentID string
var DeleteCredentialsCredentialID string

func getDeleteCredentialsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-credentials",
		Short: "Delete credentials",
		Long: "Deletes a set of stored credentials from your Discovery instance.",
		Run: DeleteCredentials,
	}

	cmd.Flags().StringVarP(&DeleteCredentialsEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&DeleteCredentialsCredentialID, "credential_id", "", "", "The unique identifier for a set of source credentials.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("credential_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteCredentials(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.DeleteCredentialsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(DeleteCredentialsEnvironmentID)
		}
		if flag.Name == "credential_id" {
			optionsModel.SetCredentialID(DeleteCredentialsCredentialID)
		}
	})

	result, _, responseErr := discovery.DeleteCredentials(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var ListGatewaysEnvironmentID string

func getListGatewaysCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-gateways",
		Short: "List Gateways",
		Long: "List the currently configured gateways.",
		Run: ListGateways,
	}

	cmd.Flags().StringVarP(&ListGatewaysEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func ListGateways(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.ListGatewaysOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(ListGatewaysEnvironmentID)
		}
	})

	result, _, responseErr := discovery.ListGateways(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var CreateGatewayEnvironmentID string
var CreateGatewayName string

func getCreateGatewayCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-gateway",
		Short: "Create Gateway",
		Long: "Create a gateway configuration to use with a remotely installed gateway.",
		Run: CreateGateway,
	}

	cmd.Flags().StringVarP(&CreateGatewayEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&CreateGatewayName, "name", "", "", "User-defined name.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func CreateGateway(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.CreateGatewayOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(CreateGatewayEnvironmentID)
		}
		if flag.Name == "name" {
			optionsModel.SetName(CreateGatewayName)
		}
	})

	result, _, responseErr := discovery.CreateGateway(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var GetGatewayEnvironmentID string
var GetGatewayGatewayID string

func getGetGatewayCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-gateway",
		Short: "List Gateway Details",
		Long: "List information about the specified gateway.",
		Run: GetGateway,
	}

	cmd.Flags().StringVarP(&GetGatewayEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&GetGatewayGatewayID, "gateway_id", "", "", "The requested gateway ID.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("gateway_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func GetGateway(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.GetGatewayOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(GetGatewayEnvironmentID)
		}
		if flag.Name == "gateway_id" {
			optionsModel.SetGatewayID(GetGatewayGatewayID)
		}
	})

	result, _, responseErr := discovery.GetGateway(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}

var DeleteGatewayEnvironmentID string
var DeleteGatewayGatewayID string

func getDeleteGatewayCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete-gateway",
		Short: "Delete Gateway",
		Long: "Delete the specified gateway configuration.",
		Run: DeleteGateway,
	}

	cmd.Flags().StringVarP(&DeleteGatewayEnvironmentID, "environment_id", "", "", "The ID of the environment.")
	cmd.Flags().StringVarP(&DeleteGatewayGatewayID, "gateway_id", "", "", "The requested gateway ID.")
	cmd.Flags().StringVarP(&Version, "version", "v", "", "The API version date to use with the service, in \"YYYY-MM-DD\" format.")
	cmd.Flags().StringVarP(&OutputFormat, "output", "", "table", "Choose an output format - can be `json`, `yaml`, or `table`.")
	cmd.Flags().StringVarP(&JMESQuery, "jmes_query", "q", "", "Provide a JMESPath query to customize output.")

	cmd.MarkFlagRequired("environment_id")
	cmd.MarkFlagRequired("gateway_id")
	cmd.MarkFlagRequired("version")

	return cmd
}

func DeleteGateway(cmd *cobra.Command, args []string) {
	utils.ConfirmRunningCommand(OutputFormat)
	discovery, discoveryErr := discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			Authenticator: Authenticator,
			Version: Version,
		})
	utils.HandleError(discoveryErr)

	optionsModel := discoveryv1.DeleteGatewayOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	flagSet := cmd.Flags()
	flagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "environment_id" {
			optionsModel.SetEnvironmentID(DeleteGatewayEnvironmentID)
		}
		if flag.Name == "gateway_id" {
			optionsModel.SetGatewayID(DeleteGatewayGatewayID)
		}
	})

	result, _, responseErr := discovery.DeleteGateway(&optionsModel)
	utils.HandleError(responseErr)

	utils.PrintOutput(result, OutputFormat, JMESQuery)
}
