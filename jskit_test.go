package jskit

import (
	"testing"
)

type User struct {
	Name string `json:"name,omitempty"`
}

func TestTool_Stringify(t *testing.T) {

	valuesToStringifyTests := []struct {
		value interface{}
	}{
		{
			value: string([]byte(`{"name":"John Doe","email":"john.doe@example.com"}`)),
		},
		{
			value: struct {
				name  string
				email string
			}{name: "Patrick Kabwe", email: "patrick.kabwe@example.com"},
		},
		{
			value: map[string]string{
				"name":  "Ryan Kabwe",
				"email": "ryan.kabwe@example.com",
			},
		},
	}

	jskit := NewJSKit[User]()
	JSON := jskit.JSON()

	for _, valuesToStringify := range valuesToStringifyTests {
		_, err := JSON.Stringify(valuesToStringify.value)

		if err != nil {
			t.Errorf("Failed to Parse %v", valuesToStringify.value)
		}

	}
}

func TestTool_Parse(t *testing.T) {

	jsonToParseTests := []struct {
		json          string
		expectedValue interface{}
	}{
		{
			json:          string([]byte(`{"name":"John Doe","email":"john.doe@example.com"}`)),
			expectedValue: "John Doe",
		},
		{
			json:          string([]byte(`{"name":"Patrick Kabwe","email":"patrick.kabwe@example.com"}`)),
			expectedValue: "Patrick Kabwe",
		},
		{
			json:          string([]byte(`{"name":"Ryan Kabwe","email":"ryan.kabwe@example.com"}`)),
			expectedValue: "Ryan Kabwe",
		},
	}

	jskit := NewJSKit[User]()
	JSON := jskit.JSON()

	for _, jsonToParse := range jsonToParseTests {
		user2, err := JSON.Parse(jsonToParse.json, User{})
		if err != nil {
			t.Errorf("Failed to Parse %v", jsonToParse.json)
		}

		if user2.Name != jsonToParse.expectedValue {
			t.Errorf("Name is not %v", jsonToParse.expectedValue)
		}
	}
}

func TestTool_ToString(t *testing.T) {

	BytesToTest := []struct {
		value         []byte
		expectedValue interface{}
	}{
		{
			value:         []byte("How are you"),
			expectedValue: "How are you",
		},
	}

	jskit := NewJSKit[User]()

	for _, byte := range BytesToTest {
		str := jskit.ToString(byte.value)

		if str != byte.expectedValue {
			t.Errorf("Name is not %v", byte.expectedValue)
		}
	}
}
