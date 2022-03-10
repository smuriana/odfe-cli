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
	handler "odfe-cli/handler/alerting"

	"github.com/spf13/cobra"
)

const (
	updateMonitorCommandName = "update"
	forceMonitorFlagName     = "force"
)

//updateMonitorCmd creates monitors with configuration from input file, if interactive mode is on,
//this command will prompt for confirmation on number of monitors will be created on executions.
var updateMonitorCmd = &cobra.Command{
	Use:   updateMonitorCommandName + " json-file-path ..." + " [flags] ",
	Short: "Update monitors based on JSON files",
	Long: "Update monitors based on a local JSON file\n" +
		"To begin, use `odfe-cli get monitor_id to retrieve the monitor configuration in JSON format. Save this JSON locally and update it for your use case. Then use `odfe-cli monitor update file-path` to update monitor.",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool(forceMonitorFlagName)
		err := updateMonitors(args, force)
		if err != nil {
			DisplayError(err, updateDetectorsCommandName)
		}
	},
}

func init() {
	GetAlertingCommand().AddCommand(updateMonitorCmd)
	updateMonitorCmd.Flags().BoolP("help", "h", false, "Help for "+updateMonitorCommandName)
}

//updateMonitor update monitor based on configurations from fileNames
func updateMonitors(fileNames []string, force bool) error {
	commandHandler, err := GetAlertingHandler()
	if err != nil {
		return err
	}
	for _, name := range fileNames {
		err = handler.UpdateMonitor(commandHandler, name, force)
		if err != nil {
			return err
		}
	}
	return nil
}
