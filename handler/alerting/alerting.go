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
	"io/ioutil"
	"odfe-cli/controller/alerting"
	entity "odfe-cli/entity/alerting"
	"os"
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

//CreateMonitor creates monitor based on file configurations
func CreateMonitor(h *Handler, fileName string) error {
	return h.CreateMonitor(fileName)
}

/*CreateMonitor Creates a monitor.
It calls http request: POST _opendistro/_alerting/monitors
Sample Input:
{
  "type": "monitor",
  "name": "test-monitor",
  "enabled": true,
  "schedule": {
    "period": {
      "interval": 1,
      "unit": "MINUTES"
    }
  },
  "inputs": [{
    "search": {
      "indices": ["movies"],
      "query": {
        "size": 0,
        "aggregations": {},
        "query": {
          "bool": {
            "filter": {
              "range": {
                "@timestamp": {
                  "gte": "||-1h",
                  "lte": "",
                  "format": "epoch_millis"
                }
              }
            }
          }
        }
      }
    }
  }],
  "triggers": [{
    "name": "test-trigger",
    "severity": "1",
    "condition": {
      "script": {
        "source": "ctx.results[0].hits.total.value > 0",
        "lang": "painless"
      }
    },
    "actions": [{
      "name": "test-action",
      "destination_id": "ld7912sBlQ5JUWWFThoW",
      "message_template": {
        "source": "This is my message body."
      },
      "throttle_enabled": true,
      "throttle": {
        "value": 27,
        "unit": "MINUTES"
      },
      "subject_template": {
        "source": "TheSubject"
      }
    }]
  }]
}*/

//GenerateMonitor generate sample monitor to provide skeleton for users
func GenerateMonitor() ([]byte, error) {

	return json.MarshalIndent(entity.CreateMonitorRequest{
		Type:     "monitor",
		Name:     "Monitor Name",
		Enabled:  true,
		Schedule: "",
		Triggers: []entity.TriggerRequest{
			{
				Name:     "",
				Severity: "",
				Condition: entity.Condition{
					Script: entity.Script{
						Source: "ctx.results[0].hits.total.value > 0",
						Lang:   "painless",
					},
				},
				Actions: []entity.ActionRequest{
					{
						Name:          "test-action",
						DestinationId: "ld7912sBlQ5JUWWFThoW",
						MessageTemplate: entity.Script{
							Source: "This is my message body.",
						},
						ThrottleEnabled: true,
						Throttle: entity.Throttle{
							Value: 5,
							Unit:  "MINUTES",
						},
						SubjectTemplate: entity.Script{
							Source: "TheSubject",
						},
					},
				},
			},
		},
	}, "", "  ")
}

//CreateMonitor creates monitor based on file configurations
func (h *Handler) CreateMonitor(fileName string) error {
	if len(fileName) < 1 {
		return fmt.Errorf("file name cannot be empty")
	}

	jsonFile, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("failed to open file %s due to %v", fileName, err)
	}
	defer func() {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println("failed to close json:", err)
		}
	}()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var request entity.CreateMonitorRequest
	err = json.Unmarshal(byteValue, &request)
	// typecasting byte array to string
	if err != nil {
		return fmt.Errorf("file %s cannot be accepted due to %v", fileName, err)
	}
	ctx := context.Background()
	monitorId, err := h.CreateMonitors(ctx, request)
	if err != nil {
		fmt.Printf("Error %v", err)
		fmt.Println()
		return err
	}
	if monitorId != nil {
		fmt.Printf("Successfully created monitor with ID %v", *monitorId)
		fmt.Println()
		return nil
	}
	return err
}

//UpdateMonitor updates monitor based on file configurations
func (h *Handler) UpdateMonitor(fileName string, force bool) error {
	if len(fileName) < 1 {
		return fmt.Errorf("file name cannot be empty")
	}

	jsonFile, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("failed to open file %s due to %v", fileName, err)
	}
	defer func() {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println("failed close json file due to ", err)
		}
	}()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var request entity.UpdateMonitorUserInput
	err = json.Unmarshal(byteValue, &request)
	if err != nil {
		return fmt.Errorf("file %s cannot be accepted due to %v", fileName, err)
	}
	ctx := context.Background()
	err = h.Controller.UpdateMonitor(ctx, request, force)
	if err != nil {
		return err
	}
	fmt.Println("Successfully updated monitor.")
	return nil
}

// UpdateMonitor updates monitor based on file configurations
func UpdateMonitor(h *Handler, fileName string, force bool) error {
	return h.UpdateMonitor(fileName, force)
}

func DeleteMonitorByID(h *Handler, monitorId string) error {
	return h.DeleteMonitorByID(monitorId)
}

//DeleteMonitorByID delete monitor based on monitorId
func (h *Handler) DeleteMonitorByID(monitorID string) error {

	ctx := context.Background()
	err := h.DeleteMonitor(ctx, monitorID)
	if err != nil {
		return err
	}
	return err
}
