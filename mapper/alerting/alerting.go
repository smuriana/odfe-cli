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

package alerting

import (
	"odfe-cli/entity/alerting"
)

const (
	minutesKey = "m"
	minutes    = "Minutes"
)

func MapToMonitorOutput(response alerting.MonitorResponse) (*alerting.MonitorOutput, error) {
	return &alerting.MonitorOutput{
		ID:            response.ID,
		Name:          response.Monitor.Name,
		Version:       response.Version,
		LastUpdatedAt: response.Monitor.LastUpdateTime,
	}, nil
}

func MapToMonitor(monitor alerting.UpdateMonitorUserInput) (*alerting.UpdateMonitor, error) {
	return &monitor.Monitor, nil
}

//MapToCreateMonitor maps to CreateMonitor
/*func MapToCreateMonitor(request alerting.CreateMonitorRequest) (*alerting.CreateMonitorRequest, error) {
	var inputs []alerting.Input

	interval, err := mapToInterval(request.Interval)
	if err != nil {
		return nil, err
	}
	delay, err := mapToInterval(request.Delay)
	if err != nil {
		return nil, err
	}
	return &ad.CreateDetector{
		Name:        request.Name,
		Description: request.Description,
		TimeField:   request.TimeField,
		Index:       request.Index,
		Features:    features,
		Filter:      request.Filter,
		Interval:    *interval,
		Delay:       *delay,
	}, nil
}*/
