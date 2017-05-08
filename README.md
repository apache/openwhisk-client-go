# incubator-openwhisk-client-go

`go-whisk` is a Go client library for accessing the IBM Whisk API.


### Usage

```go
import "github.com/apache/incubator-openwhisk-client-go/whisk"
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

Whisk can be configured by passing in a `*whisk.Config` object as the second argument to `whisk.New( ... )`.  For example:

```go
u, _ := url.Parse("https://whisk.stage1.ng.bluemix.net:443/api/v1/")
config := &whisk.Config{
  Namespace: "_",
  AuthKey: "aaaaa-bbbbb-ccccc-ddddd-eeeee",
  BaseURL: u
}
client, err := whisk.Newclient(http.DefaultClient, config)
```


### Example
```go
import (
  "net/http"
  "net/url"

  "github.com/apache/incubator-openwhisk-client-go/whisk"
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
