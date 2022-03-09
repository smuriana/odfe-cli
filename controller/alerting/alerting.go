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
	"errors"
	"fmt"
	"io"
	"odfe-cli/controller/es"
	entity "odfe-cli/entity/alerting"
	"odfe-cli/gateway/alerting"
	"odfe-cli/mapper"
	alertingmapper "odfe-cli/mapper/alerting"

	"github.com/cheggaaa/pb/v3"
)

//go:generate go run -mod=mod github.com/golang/mock/mockgen -destination=mocks/mock_alerting.go -package=mocks . Controller

//Controller is an interface for the Alerting plugin controllers
type Controller interface {
	GetMonitor(context.Context, string) (*entity.MonitorOutput, error)
	CreateMonitors(context.Context, entity.CreateMonitorRequest) (*string, error)
	DeleteMonitor(context.Context, string, bool) error
	UpdateMonitor(context.Context, entity.UpdateMonitorUserInput, bool) ([]byte, error)
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

func validateCreateRequest(r entity.CreateMonitorRequest) error {
	if len(r.Type) < 1 {
		return fmt.Errorf("type field cannot be empty")
	}
	if len(r.Name) < 1 {
		return fmt.Errorf("name field cannot be empty")
	}
	if len(r.Inputs) < 1 {
		return fmt.Errorf("inputs cannot be empty")
	}
	return nil
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

//createProgressBar creates progress bar with suffix as counter and number of action completed, prefix as percentage
func createProgressBar(total int) *pb.ProgressBar {
	template := `{{string . "prefix"}}{{percent . }} {{bar . "[" "=" ">" "_" "]" }} {{counters . }}{{string . "suffix"}}`
	bar := pb.New(total)
	bar.SetTemplateString(template)
	bar.SetMaxWidth(65)
	bar.Start()
	return bar
}

func processEntityError(err error) error {
	var c entity.CreateError
	data := fmt.Sprintf("%v", err)
	responseErr := json.Unmarshal([]byte(data), &c)
	if responseErr != nil {
		return err
	}
	if len(c.Error.Reason) > 0 {
		return errors.New(c.Error.Reason)
	}
	return err
}

//CreateMonitor creates monitor based on user request
func (c controller) CreateMonitors(ctx context.Context, r entity.CreateMonitorRequest) (*string, error) {
	if err := validateCreateRequest(r); err != nil {
		return nil, err
	}
	response, err := c.gateway.CreateMonitor(ctx, r)

	if err != nil {
		return nil, processEntityError(err)
	}

	var data map[string]interface{}
	_ = json.Unmarshal(response, &data)
	monitorID := fmt.Sprintf("%s", data["_id"])

	return mapper.StringToStringPtr(monitorID), nil
}

//DeleteMonitor deletes monitor based on Id.
func (c controller) DeleteMonitor(ctx context.Context, id string, force bool) error {
	monitorOutput, err := c.GetMonitor(ctx, id)
	if err != nil {
		return err
	}
	if monitorOutput == nil {
		return nil
	}
	_, err = c.gateway.DeleteMonitor(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

//UpdateMonitor updates monitor based on MonitorID, if force is enabled, it overrides without checking whether
// user downloaded latest version before updating it, if enable is true, monitor will be enabled after update
func (c controller) UpdateMonitor(ctx context.Context, input entity.UpdateMonitorUserInput, force bool) ([]byte, error) {
	if len(input.ID) < 1 {
		return nil, fmt.Errorf("monitor Id cannot be empty")
	}
	if !force {
		latestDetector, err := c.GetMonitor(ctx, input.ID)
		if err != nil {
			return nil, err
		}
		if latestDetector.LastUpdatedAt > input.LastUpdatedAt {
			return nil, fmt.Errorf(
				"new version for monitor is available. Please fetch latest version and then merge your changes")
		}
	}
	payload, err := alertingmapper.MapToUpdateMonitor(input)
	if err != nil {
		return nil, err
	}
	monitorUpdated, err := c.gateway.UpdateMonitor(ctx, input.ID, payload)
	if err != nil {
		return nil, err
	}
	return monitorUpdated, nil
}
