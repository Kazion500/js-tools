## JSKit Go Module

JSKit is a Go module that provides a set of utility functions commonly used in JavaScript programming. It offers an interface to handle JSON data, allowing users to stringify and parse JSON strings. Additionally, JSKit provides other useful functions like sorting arrays, searching for elements in arrays, and manipulating strings.

### Installation

To use this module in your Go project, you can run the following command:

```golang
go get github.com/Kazion500/jskit
```

After installing the package, you can import it in your Go code:

```golang
import "github.com/Kazion500/jskit"
```

### Usage

The JSKit module provides an `interface` that allows users to work with JSON data. The JSON function returns an `interface` that provides methods to stringify and parse JSON data. The ToString function returns a `string` representation of a `byte` array.

```golang
type AnyType interface{}

type JSKit[T AnyType] interface {
	JSON() JSON[T]
	ToString(value []byte) string
}

type JSON[T AnyType] interface {
	Stringify(any) (string, error)
	Parse(string, T) (T, error)
}
```

The `jsont` and `jskit` structs implement the JSON and JSKit interfaces respectively. The NewJSKit function returns a new JSKit interface.

### Examples
Here's an example of how to use this module to stringify and parse JSON data:

```golang
package main

import (
    "fmt"
    "github.com/username/repo/jskit"
)

type Person struct {
    Name string `json:"name"`
    Age int `json:"age"`
}

func main() {
    person := Person{"John Doe", 30}
    jskit := jskit.NewJSKit[Person]()
    jsonString, _ := jskit.JSON.Stringify(person)
    fmt.Println(jsonString)

    // Output: {"name":"John Doe","age":30}

    var p Person
    jskit.JSON.Parse(jsonString, &p)
    fmt.Println(p.Name)

    // Output: John Doe

}
```

### Contributing

We welcome contributions to this project. To contribute, please fork this repository and create a pull request.

### Author

#### Email: [Patrick](mailto:patrickckabwe@gmail.com)

#### Full Name: [Patrick Kabwe](github.com/Kazion500)
