## Go Whisk Client



##### Simple
```go
import (
  "net/http"
  "net/url"

  whisk "github.ibm.com/Bluemix/go-whisk"
)

func main() {
  httpClient := http.DefaultClient
  whisk, err := whisk.New(httpClient, nil)
  if err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }

  actions, resp, err := whisk.Actions.Get("actionName")
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
