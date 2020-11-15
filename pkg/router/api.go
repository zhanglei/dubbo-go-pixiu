/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package router

import (
	"net/url"
)

import (
	"github.com/dubbogo/dubbo-go-proxy/pkg/config"
	"strings"
)

// API describes the minimum configuration of an RESTful api configure in gateway
type API struct {
	URLPattern    string `json:"urlPattern" yaml:"urlPattern"`
	config.Method `json:"method,inline" yaml:"method,inline"`
}

// GetURIParams returns the values retrieved from the rawURL
func (api *API) GetURIParams(rawURL url.URL) url.Values {
	sourceURL := strings.Split(rawURL.Path, "&")[0]
	return wildcardMatch(api.URLPattern, sourceURL)
}
