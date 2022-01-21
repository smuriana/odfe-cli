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
	"encoding/json"
	"fmt"
	"io"
	"odfe-cli/controller/es"
	entity "odfe-cli/entity/alerting"
	"odfe-cli/gateway/alerting"
	alertingmapper "odfe-cli/mapper/alerting"
)

//go:generate go run -mod=mod github.com/golang/mock/mockgen -destination=mocks/mock_alerting.go -package=mocks . Controller

//Controller is an interface for the Alerting plugin controllers
type Controller interface {
	GetMonitor(context.Context, string) (*entity.MonitorOutput, error)
}

type controller struct {
	reader  io.Reader
	gateway alerting.Gateway
	esCtrl  es.Controller
}

//New returns new Controller instance
func New(reader io.Reader, esCtrl es.Controller, gateway alerting.Gateway) Controller {
	return &controller{
		reader,
		gateway,
		esCtrl,
	}
}

//GetMonitor fetch monitor based on MonitorID
func (c controller) GetMonitor(ctx context.Context, ID string) (*entity.MonitorOutput, error) {
	if len(ID) < 1 {
		return nil, fmt.Errorf("monitor Id: %s cannot be empty", ID)
	}
	response, err := c.gateway.GetMonitor(ctx, ID)
	if err != nil {
		return nil, err
	}
	var data entity.MonitorResponse
	err = json.Unmarshal(response, &data)
	if err != nil {
		return nil, err
	}
	return alertingmapper.MapToMonitorOutput(data)
}
