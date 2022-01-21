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
	"context"
	"fmt"
	"net/http"
	"net/url"
	"odfe-cli/client"
	"odfe-cli/entity"
	gw "odfe-cli/gateway"
)

const (
	baseURL           = "_opendistro/_alerting/monitors"
	searchURLTemplate = baseURL + "/_search"
	deleteURLTemplate = baseURL + "/%s"
	getURLTemplate    = baseURL + "/%s"
	updateURLTemplate = baseURL + "/%s"
)

//go:generate go run -mod=mod github.com/golang/mock/mockgen  -destination=mocks/mock_ad.go -package=mocks . Gateway

// Gateway interface to AD Plugin
type Gateway interface {
	GetMonitor(context.Context, string) ([]byte, error)
}

type gateway struct {
	gw.HTTPGateway
}

// New creates new Gateway instance
func New(c *client.Client, p *entity.Profile) Gateway {
	return &gateway{*gw.NewHTTPGateway(c, p)}
}

func (g *gateway) buildGetURL(ID string) (*url.URL, error) {
	endpoint, err := gw.GetValidEndpoint(g.Profile)
	if err != nil {
		return nil, err
	}
	endpoint.Path = fmt.Sprintf(getURLTemplate, ID)
	return endpoint, nil
}

// GetMonitor Returns all information about a detector based on the monitor_id.
// It calls http request: GET _opendistro/_alerting/monitors/<monitorId>
func (g *gateway) GetMonitor(ctx context.Context, ID string) ([]byte, error) {
	getURL, err := g.buildGetURL(ID)
	if err != nil {
		return nil, err
	}
	monitorRequest, err := g.BuildRequest(ctx, http.MethodGet, "", getURL.String(), gw.GetDefaultHeaders())
	if err != nil {
		return nil, err
	}
	response, err := g.Call(monitorRequest, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return response, nil
}