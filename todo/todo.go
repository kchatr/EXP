package todo

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Item struct {
	Text string
}

func SaveItems(filename string, items []Item) error {

	b, err := json.Marshal(items)
	err = ioutil.WriteFile(filename, b, 0644)

	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}