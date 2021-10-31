package learngojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName: "Raisa",
		LastName:  "Supriatna",
		Age:       21,
		Married:   false,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
