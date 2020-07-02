package logger

import (
	"encoding/json"
	"log"
)

type any interface{}

func JSON(data any) {
	if data, err := json.Marshal(data); err == nil {
		log.Print(string(data))
	} else {
		log.Fatal(err)
	}
}
