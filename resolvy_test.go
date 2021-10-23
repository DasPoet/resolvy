package resolvy

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func Example() {
	type Message struct {
		CreatedAt time.Time   `resolvy:"createdAt,omitempty"`
		Data      interface{} `resolvy:"data,omitempty"`
	}
	msg := Message{CreatedAt: time.Now()}

	data, err := MarshalJSON(msg, MarshalConfig{
		Marshalers: map[string]FieldMarshaler{
			"createdAt": func() (interface{}, error) {
				return msg.CreatedAt.String(), nil
			},
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(data))
}

func Test(t *testing.T) {
	Example()
}
