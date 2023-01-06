package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var (
		i    interface{}
		m    map[string]interface{}
		ok   bool
		err  error
		data []byte
	)

	data = []byte(`{"name":"example", "foo":123, "bar":true, "baz":"something"}`)
	if err = json.Unmarshal(data, &i); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("i: %+v\n", i)

	if m, ok = i.(map[string]interface{}); !ok {
		log.Fatal("failed to type assert data")
	}

	fmt.Printf("m: %+v\n", m)

	if _, ok := m["name"]; ok {
		m["name"] = "dynamic"
	}

	fmt.Printf("m: %+v\n", m)

	if data, err = json.Marshal(m); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("data: %+v\n", string(data))
}
