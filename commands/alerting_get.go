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
	"encoding/json"
	"fmt"
	"io"
	entity "odfe-cli/entity/alerting"
	"odfe-cli/handler/alerting"
	"os"

	"github.com/spf13/cobra"
)

const (
	getMonitorCommandName = "get"
	getMonitorIDFlagName  = "id"
)

//getMonitorsCmd prints Monitors configuration based on id, name or name regex pattern.
//default input is name pattern, one can change this format to be id by passing --id flag
var getMonitorsCmd = &cobra.Command{
	Use:   getMonitorCommandName + " monitor_id ..." + " [flags] ",
	Short: "Get Monitors based on a list of IDs, names, or name regex patterns",
	Long: "Get Monitors based on a list of IDs, names, or name regex patterns.\n" +
		"Wrap regex patterns in quotation marks to prevent the terminal from matching patterns against the files in the current directory.\nThe default input is detector name. Use the `--id` flag if input is detector ID instead of name",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := printMonitors(PrintlnMonitor, cmd, args)
		if err != nil {
			DisplayError(err, getMonitorCommandName)
		}
	},
}

type DisplayMonitor func(*cobra.Command, *entity.MonitorOutput) error

//fprint displays the list of detectors.
func fprintMonitor(cmd *cobra.Command, display DisplayMonitor, results []*entity.MonitorOutput) error {
	if results == nil {
		return nil
	}
	for _, d := range results {
		if err := display(cmd, d); err != nil {
			return err
		}
	}
	return nil
}

//printMonitors print Monitors
func printMonitors(display DisplayMonitor, cmd *cobra.Command, Monitors []string) error {
	idStatus, _ := cmd.Flags().GetBool(getMonitorIDFlagName)
	commandHandler, err := GetAlertingHandler()
	if err != nil {
		return err
	}
	// default is name
	action := alerting.GetMonitorByID
	if idStatus {
		action = alerting.GetMonitorByID
	}

	results, err := getMonitors(commandHandler, Monitors, action)
	if err != nil {
		return err
	}
	return fprintMonitor(cmd, display, results)
}

//getMonitors fetch detector from controller
func getMonitors(
	commandHandler *alerting.Handler, args []string, get func(*alerting.Handler, string) (
		[]*entity.MonitorOutput, error)) ([]*entity.MonitorOutput, error) {
	var results []*entity.MonitorOutput
	for _, monitor := range args {
		output, err := get(commandHandler, monitor)
		if err != nil {
			return nil, err
		}
		results = append(results, output...)
	}
	return results, nil
}

//getMonitorsByID gets detector output based on ID as argument
func GetMonitorByID(commandHandler *alerting.Handler, ID string) ([]*entity.MonitorOutput, error) {

	output, err := alerting.GetMonitorByID(commandHandler, ID)
	if err != nil {
		return nil, err
	}
	return output, nil
}

//FPrint prints detector configuration on writer
//Since this is json format, use indent function to pretty print before printing on writer
func FPrintMonitor(writer io.Writer, d *entity.MonitorOutput) error {
	formattedOutput, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(writer, string(formattedOutput))
	return err
}

//Println prints detector configuration on stdout
func PrintlnMonitor(cmd *cobra.Command, d *entity.MonitorOutput) error {
	return FPrintMonitor(os.Stdout, d)
}

func init() {
	GetAlertingCommand().AddCommand(getMonitorsCmd)
	getMonitorsCmd.Flags().BoolP(getMonitorIDFlagName, "", false, "Input is monitor ID")
	getMonitorsCmd.Flags().BoolP("help", "h", false, "Help for "+getMonitorCommandName)
}
