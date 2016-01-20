## Go Whisk Client



##### Simple
```go
import (
  "net/http"
  "net/url"

  "github.ibm.com/Bluemix/go-whisk/whisk"
)

func main() {
  client, err := whisk.New(http.DefaultClient, nil)
  if err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }

  options := &whisk.ActionListOptions{
    Limit: 30,
    Skip: 0,
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

##### Configuration

Whisk can be configured by passing in a `*whisk.Config` object as the second argument to `whisk.New( ... )`.  Its declaration is:

```go
package whisk

type Config struct {
	Namespace string // NOTE :: Default is "_"
	AuthToken string
	BaseURL   *url.URL // NOTE :: Default is "whisk.stage1.ng.bluemix.net"
	Version   string
	Verbose   bool
}
```
