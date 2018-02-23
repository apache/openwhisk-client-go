// +build unit

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

package whisk

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

const (
	TRIGGER_GET_NO_RULES = `{
                "namespace": "test@openwhisk_dev",
                "name": "testTrigger",
                "publish": false,
                "version": "0.0.1",
                "limits": {}
        }`

	TRIGGER_GET_WITH_RULES = `{
                "namespace": "test@openwhisk_dev",
                "name": "testTrigger",
                "publish": false,
                "version": "0.0.1",
                "limits": {},
                "rules": {
                        "guest/inactiverule": {
                                "action": {
                                        "name": "web-echo-env",
                                        "path": "guest"
                                },
                        "status": "inactive"
                        }
                }
        }`
)

type TriggerResponse struct {
	Body string
}

type TriggerRequest struct {
	Method string
	URL    string
}

var triggerResponse = &TriggerResponse{}
var triggerRequest = &TriggerRequest{}

type MockTriggerClient struct{}

func (c *MockTriggerClient) NewRequestUrl(method string, urlRelResource *url.URL, body interface{}, includeNamespaceInUrl bool, appendOpenWhiskPath bool, encodeBodyAs string, useAuthentication bool) (*http.Request, error) {
	return &http.Request{}, nil
}

func (c *MockTriggerClient) NewRequest(method, urlStr string, body interface{}, includeNamespaceInUrl bool) (*http.Request, error) {
	triggerRequest.Method = method
	triggerRequest.URL = urlStr

	request, err := http.NewRequest(method, urlStr, nil)
	if err != nil {
		fmt.Printf("http.NewRequest() failure: %s\n", err)
		return &http.Request{}, err
	}

	return request, nil
}

func (c *MockTriggerClient) Do(req *http.Request, v interface{}, ExitWithErrorOnTimeout bool, secretToObfuscate ...ObfuscateSet) (*http.Response, error) {
	var reader = strings.NewReader(triggerResponse.Body)

	dc := json.NewDecoder(reader)
	dc.UseNumber()
	err := dc.Decode(v)

	if err != nil {
		fmt.Printf("json decode failure: %s\n", err)
		return nil, err
	}

	resp := &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(reader),
	}

	return resp, nil
}

func TestTriggerGet(t *testing.T) {
	assert := assert.New(t)
	mockClient := &MockTriggerClient{}
	triggerService := &TriggerService{client: mockClient}
	var nilMap map[string]interface{}

	triggerResponse.Body = TRIGGER_GET_NO_RULES
	trigger, _, _ := triggerService.Get("testTrigger")
	assert.Equal("GET", triggerRequest.Method)
	assert.Equal("triggers/testTrigger", triggerRequest.URL)
	assert.Equal(nilMap, trigger.Rules)

	triggerResponse.Body = TRIGGER_GET_WITH_RULES
	var expectedTrigger map[string]interface{}
	json.Unmarshal([]byte(triggerResponse.Body), &expectedTrigger)
	expectedRules, _ := expectedTrigger["rules"]
	trigger, _, _ = triggerService.Get("testTrigger")
	assert.Equal("GET", triggerRequest.Method)
	assert.Equal("triggers/testTrigger", triggerRequest.URL)
	assert.Equal(expectedRules, trigger.Rules)
}
