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
	"odfe-cli/client"
	alertingctrl "odfe-cli/controller/alerting"
	esctrl "odfe-cli/controller/es"
	alertinggateway "odfe-cli/gateway/alerting"
	esgateway "odfe-cli/gateway/es"
	handler "odfe-cli/handler/alerting"
	"os"

	"github.com/spf13/cobra"
)

const (
	alertingCommandName           = "alerting"
	alertingMonitorCommandName    = "monitor"
	alertingMonitorGetFlagName    = "get"
	alertingMonitorCreateFlagName = "create"
)

//alertingCommand is base command for Alerting plugin.
var alertingCommand = &cobra.Command{
	Use:   alertingCommandName,
	Short: "Manage the Alerting plugin",
	Long:  "Use the Alerting commands to create, configure, and manage monitors.",
}

func init() {
	alertingCommand.Flags().BoolP("help", "h", false, "Help for Alerting")
	GetRoot().AddCommand(alertingCommand)
}

//GetAlertingCommand returns Alerting base command, since this will be needed for subcommands
//to add as parent later
func GetAlertingCommand() *cobra.Command {
	return alertingCommand
}

//GetAlertingHandler returns handler by wiring the dependency manually
func GetAlertingHandler() (*handler.Handler, error) {
	c, err := client.New(nil)
	if err != nil {
		return nil, err
	}
	profile, err := GetProfile()
	if err != nil {
		return nil, err
	}
	g := alertinggateway.New(c, profile)
	esg := esgateway.New(c, profile)
	esc := esctrl.New(esg)
	ctr := alertingctrl.New(os.Stdin, esc, g)
	return handler.New(ctr), nil
}
