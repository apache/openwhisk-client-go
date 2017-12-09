//// +build unit

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
    "testing"
    "github.com/stretchr/testify/assert"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "strings"
    "net/url"
)

const (
    NODE_ACTION_NO_CODE = `{
        "name": "test",
        "publish": false,
        "annotations": [
            {
                "key": "exec",
                "value": "nodejs:6"
            }
        ],
        "version": "0.0.1",
        "exec": {
            "kind": "nodejs:6",
            "binary": false
        },
        "parameters": [],
        "limits": {
            "timeout": 60000,
            "memory": 256,
            "logs": 10
        },
        "namespace": "test@openwhisk"
    }`

    NODE_ACTION = `{
        "name": "test",
        "publish": false,
        "annotations": [
            {
                "key": "exec",
                "value": "nodejs:6"
            }
        ],
        "version": "0.0.1",
        "exec": {
            "kind": "nodejs:6",
            "code": "...",
            "binary": false
        },
        "parameters": [],
        "limits": {
            "timeout": 60000,
            "memory": 256,
            "logs": 10
        },
        "namespace": "test@openwhisk"
    }`
)

type ActionResponse struct {
    Body    string
}

type ActionRequest struct {
    Method  string
    URL     string
}

var actionResponse = &ActionResponse{}
var actionRequest = &ActionRequest{}

type MockClient struct {}

func (c *MockClient) NewRequestUrl(method string, urlRelResource *url.URL, body interface{}, includeNamespaceInUrl bool, appendOpenWhiskPath bool, encodeBodyAs string, useAuthentication bool) (*http.Request, error) {
    return &http.Request{}, nil
}

func (c *MockClient) NewRequest(method, urlStr string, body interface{}, includeNamespaceInUrl bool) (*http.Request, error) {
    actionRequest.Method = method
    actionRequest.URL = urlStr

    request, err := http.NewRequest(method, urlStr, nil)
    if (err != nil) {
        return &http.Request{}, err
    }

    return request, nil
}

func (c *MockClient) Do(req *http.Request, v interface{}, ExitWithErrorOnTimeout bool, secretToObfuscate ...ObfuscateSet) (*http.Response, error) {
    var reader = strings.NewReader(actionResponse.Body)

    dc := json.NewDecoder(reader)
    dc.UseNumber()
    err := dc.Decode(v)

    if err != nil {
        return nil, err
    }

    resp := &http.Response{
        StatusCode: 200,
        Body: ioutil.NopCloser(reader),
    }

    return resp, nil
}

func TestActionGet(t *testing.T) {
    assert := assert.New(t)
    mockClient := &MockClient{}
    actionService := &ActionService{client: mockClient}

    actionResponse.Body = NODE_ACTION_NO_CODE
    action, _, _ := actionService.Get("test", false)

    var exec Exec
    exec = *action.Exec
    var nilStr *string

    assert.Equal("GET", actionRequest.Method)
    assert.Equal("actions/test?code=false", actionRequest.URL)
    assert.Equal(nilStr, exec.Code)

    actionResponse.Body = NODE_ACTION
    action, _, _ = actionService.Get("test", true)
    assert.Equal("GET", actionRequest.Method)
    assert.Equal("actions/test?code=true", actionRequest.URL)
    assert.Equal("...", *action.Exec.Code)
}
