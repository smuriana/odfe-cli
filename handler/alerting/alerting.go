/*
 * Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *     http://www.apache.org/licenses/LICENSE-2.0
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package alerting

import (
	"context"
	"odfe-cli/controller/alerting"
	entity "odfe-cli/entity/alerting"
)

//Handler is facalertine for controller
type Handler struct {
	alerting.Controller
}

// New returns new Handler instance
func New(controller alerting.Controller) *Handler {
	return &Handler{
		controller,
	}
}

// GetMonitorByID gets monitor based on monitor id
func GetMonitorByID(h *Handler, monitor string) ([]*entity.MonitorOutput, error) {
	return h.GetMonitorByID(monitor)
}

// GetMonitorByID gets monitor based on monitor id
func (h *Handler) GetMonitorByID(ID string) ([]*entity.MonitorOutput, error) {
	ctx := context.Background()
	monitor, err := h.GetMonitor(ctx, ID)
	if err != nil {
		return nil, err
	}
	output := []*entity.MonitorOutput{monitor}
	return output, nil
}
