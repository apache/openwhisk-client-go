<!--
#
# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
-->

# Openwhisk Client Go
[![License](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0)
[![Build Status](https://travis-ci.com/apache/openwhisk-client-go.svg?branch=master)](https://travis-ci.com/apache/openwhisk-client-go)

This project `openwhisk-client-go` is a Go client library to access the Openwhisk API.

---

## Building the project

### Prerequisites

The Openwhisk Go Client library requires you to [Download and install GoLang](https://golang.org/dl/) onto your local machine.

> **Note** Go version 1.15 or higher is recommended

Make sure you select the package that fits your local environment, and [set the GOPATH environment variable](https://github.com/golang/go/wiki/SettingGOPATH).

### Download the source code from GitHub

As the code is managed using GitHub, it is easiest to retrieve the code using the `git clone` command.

If you just want to build the code and do not intend to be a Contributor, you can clone the latest code from the Apache repository:

```sh
git clone git@github.com:apache/openwhisk-client-go
```

You can also specify a release (tag), if you do not want the latest code, by using the `--branch <tag>` flag. For example, you can clone the source code for the tagged 1.1.0 [release](https://github.com/apache/openwhisk-client-go/releases)

```sh
git clone --branch 1.1.0 git@github.com:apache/openwhisk-client-go
```

You can also pull the code from a fork of the repository. If you intend to become a Contributor to the project, read the section [Contributing to the project](#contributing-to-the-project) below on how to setup a fork.

### Building using `go build`

Change into the cloned project directory and use the following command to build all packages:

```sh
$ go build -v ./...
```

or simply build just the whisk commands:

```sh
$ go build -v ./whisk
```

> **Note**: There is no `main` function in this project as the `./whish` packages are treated together as a client library.

### Testing using `go test`

Open a terminal, change into the project directory and use the following command to run the unit tests:

```
$ go test -v ./... -tags=unit
```

You should see all the unit tests passed; if not, please [log an issue](https://github.com/apache/openwhisk-client-go/issues) for us.

---

## Configuration

This Go client library is used to access the OpenWhisk API, so please make sure you have an OpenWhisk service running somewhere
available for you to run this library.

We use a configuration file called _wskprop_ to specify all the parameters necessary for this Go client library to access the OpenWhisk services. Make sure you create or edit the file _~/.wskprops_, and add the mandatory parameters APIHOST, APIVERSION, NAMESPACE and AUTH.

- The parameter `APIHOST` is the OpenWhisk API hostname (for example, openwhisk.ng.bluemix.net, 172.17.0.1, and so on).
- The parameter `APIVERSION` is the version of OpenWhisk API to be used to access the OpenWhisk resources.
- The parameter `NAMESPACE` is the OpenWhisk namespace used to specify the OpenWhisk resources about to be accessed.
- The parameter `AUTH` is the authentication key used to authenticate the incoming requests to the OpenWhisk services.

For more information regarding the REST API of OpenWhisk, please refer to [OpenWhisk REST API](https://github.com/apache/openwhisk/blob/master/docs/rest_api.md).

## Usage

```go
import "github.com/apache/openwhisk-client-go/whisk"
```

Construct a new whisk client, then use various services to access different parts of the whisk api.  For example to get the `hello` action:

```go
client, _ := whisk.NewClient(http.DefaultClient, nil)
action, resp, err := client.Actions.List("hello")
```

Some API methods have optional parameters that can be passed. For example, to list the first 30 actions, after the 30th action:
```go
client, _ := whisk.NewClient(http.DefaultClient, nil)

options := &whisk.ActionListOptions{
  Limit: 30,
  Skip: 30,
}

actions, resp, err := client.Actions.List(options)
```

By default, this Go client library is automatically configured by the configuration file _wskprop_. The parameters of APIHOST, APIVERSION,
NAMESPACE and AUTH will be used to access the OpenWhisk services.

In addition, it can also be configured by passing in a `*whisk.Config` object as the second argument to `whisk.New( ... )`.  For example:

```go
config := &whisk.Config{
  Host: "openwhisk.ng.bluemix.net",
  Version: "v1"
  Namespace: "_",
  AuthKey: "aaaaa-bbbbb-ccccc-ddddd-eeeee"
}
client, err := whisk.Newclient(http.DefaultClient, config)
```

### Example

You need to have an OpenWhisk service accessible, to run the following example.

```go
import (
  "net/http"
  "net/url"

  "github.com/apache/openwhisk-client-go/whisk"
)

func main() {
  client, err := whisk.NewClient(http.DefaultClient, nil)
  if err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }

  options := &whisk.ActionListOptions{
    Limit: 30,
    Skip: 30,
  }

  actions, resp, err := client.Actions.List(options)
  if err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }

  fmt.Println("Returned with status: ", resp.Status)
  fmt.Println("Returned actions: \n%+v", actions)

}
```

Then build it with the go tool:

```
$ go build
```

If the openWhisk service is available and your configuration is correct, you should receive the status and the actions with the above example.
