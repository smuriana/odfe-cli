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
	"encoding/json"
	"odfe-cli/entity"
)

//Period represents time interval
type Period struct {
	Interval int32  `json:"interval"`
	Unit     string `json:"unit"`
}

//Bool type for must query
type Bool struct {
	Must []json.RawMessage `json:"must"`
}

//Query type to represent query
type Query struct {
	Bool Bool `json:"bool"`
}

//CreateFailedError structure if create failed
type CreateFailedError struct {
	Type   string `json:"type"`
	Reason string `json:"reason"`
}

//CreateError Error type in Create Response
type CreateError struct {
	Error  CreateFailedError `json:"error"`
	Status int32             `json:"status"`
}

//Configuration represents configuration in config file
type Configuration struct {
	Profiles []entity.Profile `mapstructure:"profiles"`
}

//Match specifies name
type Match struct {
	Name string `json:"name"`
}

//SearchQuery contains match names
type SearchQuery struct {
	Match Match `json:"match"`
}

//SearchRequest represents structure for search detectors
type SearchRequest struct {
	Query SearchQuery `json:"query"`
}

type Schedule struct {
	Period Period `json:"period"`
}

type Search struct {
	Indices []string    `json:"indices"`
	Query   SearchQuery `json:"query"`
}

type Input struct {
	Search json.RawMessage `json:"search"`
}

type Script struct {
	Source string `json:"source"`
	Lang   string `json:"lang,omitempty"`
}

type Condition struct {
	Script Script `json:"script"`
}
type Trigger struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Severity  string    `json:"severity"`
	Condition Condition `json:"condition"`
	Actions   []Action  `json:"actions"`
}

type Action struct {
	Name            string    `json:"name"`
	DestionationId  string    `json:"destination_id"`
	SubjectTemplate Script    `json:"subject_template"`
	MessageTemplate Script    `json:"message_template"`
	Severity        string    `json:"severity"`
	Condition       Condition `json:"condition"`
}

type Monitor struct {
	Type           string    `json:"type"`
	Name           string    `json:"name"`
	Enabled        bool      `json:"enabled"`
	Enabled_time   uint64    `json:"enabled_time"`
	Schedule       Schedule  `json:"schedule"`
	Inputs         []Input   `json:"inputs"`
	Triggers       []Trigger `json:"triggers"`
	LastUpdateTime uint64    `json:"last_update_time"`
}

//MonitorResponse represents monitor's setting
type MonitorResponse struct {
	ID      string  `json:"_id"`
	Version int32   `json:"_version"`
	Monitor Monitor `json:"monitor"`
}

//MonitorOutput represents detector's setting displayed to user
type MonitorOutput struct {
	ID            string
	Name          string `json:"name"`
	Version       int32  `json:"version"`
	LastUpdatedAt uint64 `json:"last_update_time"`
}

type Throttle struct {
	Value int32  `json:"value"`
	Unit  string `json:"unit"`
}

type ActionRequest struct {
	Name            string   `json:"name"`
	DestinationId   string   `json:"destination_id"`
	MessageTemplate Script   `json:"message_template"`
	ThrottleEnabled bool     `json:"throttle_enabled"`
	Throttle        Throttle `json:"throttle"`
	SubjectTemplate Script   `json:"subject_template"`
}

type TriggerRequest struct {
	Name      string          `json:"name"`
	Severity  string          `json:"severity"`
	Condition Condition       `json:"condition,omitempty"`
	Actions   []ActionRequest `json:"actions,omitempty"`
}

//CreateMonitorRequest represents request for alerting
type CreateMonitorRequest struct {
	Type     string           `json:"type"`
	Name     string           `json:"name"`
	Enabled  bool             `json:"enabled"`
	Schedule Schedule         `json:"schedule"`
	Inputs   []Input          `json:"inputs,omitempty"`
	Triggers []TriggerRequest `json:"triggers,omitempty"`
}

type UpdateMonitorUserInput struct {
	ID            string  `json:"_id"`
	Version       int32   `json:"_version"`
	Monitor       Monitor `json:"monitor"`
	LastUpdatedAt uint64  `json:"last_update_time"`
}
