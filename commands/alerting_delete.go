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
	deleteMonitorsCommandName = "delete"
	deleteMonitorIDFlagName   = "id"
)

//deleteMonitorsCmd deletes monitors based on id
var deleteMonitorsCmd = &cobra.Command{
	Use:   deleteMonitorsCommandName + " monitor_id ..." + " [flags] ",
	Short: "Delete monitor based on the ID",
	Long:  "Delete monitor based on the ID.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		err := deleteMonitor(args)
		DisplayError(err, deleteMonitorsCommandName)
	},
}

func init() {
	GetAlertingCommand().AddCommand(deleteMonitorsCmd)
	deleteMonitorsCmd.Flags().BoolP(deleteMonitorIDFlagName, "", false, "Input is monitor ID")
	deleteMonitorsCmd.Flags().BoolP("help", "h", false, "Help for "+deleteMonitorsCommandName)
}

//deleteMonitor deletes monitor with force by calling delete method provided
func deleteMonitor(monitors []string) error {
	commandHandler, err := GetAlertingHandler()
	if err != nil {
		return err
	}
	for _, monitor := range monitors {
		err = handler.DeleteMonitorByID(commandHandler, monitor)
		if err != nil {
			return err
		}
	}

	return nil
}
