/*
 * Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package commands

import (
	"fmt"
	handler "odfe-cli/handler/alerting"

	"github.com/spf13/cobra"
)

const (
	createMonitorCommandName = "create"
	generateMonitor          = "generate-template"
)

//createMonitorCmd creates monitors with configuration from input file, if interactive mode is on,
//this command will prompt for confirmation on number of monitors will be created on executions.
var createMonitorCmd = &cobra.Command{
	Use:   createMonitorCommandName + " json-file-path ..." + " [flags] ",
	Short: "Create monitors based on JSON files",
	Long: "Create monitors based on a local JSON file\n" +
		"To begin, use `odfe-cli alerting create --generate-template` to generate a sample configuration. Save this template locally and update it for your use case. Then use `odfe-cli monitor create file-path` to create monitor.",
	Run: func(cmd *cobra.Command, args []string) {
		generateMonitor, _ := cmd.Flags().GetBool(generateMonitor)
		if generateMonitor {
			generateMonitorTemplate()
			return
		}
		//If no args, display usage
		if len(args) < 1 {
			fmt.Println(cmd.Usage())
			return
		}
		err := createMonitors(args)
		DisplayError(err, createMonitorCommandName)
	},
}

//generateMonitorTemplate prints sample monitor configuration
func generateMonitorTemplate() {
	monitor, _ := handler.GenerateMonitor()
	fmt.Println(string(monitor))
}

func init() {
	GetAlertingCommand().AddCommand(createMonitorCmd)
	createMonitorCmd.Flags().BoolP(generateMonitor, "g", false, "Output sample monitor configuration")
	createMonitorCmd.Flags().BoolP("help", "h", false, "Help for "+createMonitorCommandName)
}

//createMonitor create monitor based on configurations from fileNames
func createMonitors(fileNames []string) error {
	commandHandler, err := GetAlertingHandler()
	if err != nil {
		return err
	}
	for _, name := range fileNames {
		err = handler.CreateMonitor(commandHandler, name)
		if err != nil {
			return err
		}
	}
	return nil
}
