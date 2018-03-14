package whisk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func hello(event json.RawMessage) (json.RawMessage, error) {
	var obj map[string]interface{}
	json.Unmarshal(event, &obj)
	name, ok := obj["name"].(string)
	if !ok {
		name = "Stranger"
	}
	fmt.Printf("name=%s\n", name)
	msg := map[string]string{"message": ("Hello, " + name + "!")}
	return json.Marshal(msg)
}

func Example_repl() {
	in := bytes.NewBufferString("{\"name\":\"Mike\"}\nerr\n")
	repl(hello, in, os.Stdout)
	// Output:
	// name=Mike
	// {"message":"Hello, Mike!"}
	// name=Stranger
	// {"message":"Hello, Stranger!"}
}

func ExampleStart() {
	StartWithArgs(hello, []string{"{\"name\":\"Mike\"}", "err"})
	// Output:
	// name=Mike
	// {"message":"Hello, Mike!"}
	// name=Stranger
	// {"message":"Hello, Stranger!"}
}
