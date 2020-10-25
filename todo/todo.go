package todo

import (
	"strconv"
	"encoding/json"
	"io/ioutil"
)

type Item struct {
	Text string
	Priority int
	position int
}

type ByPri []Item

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:	
		i.Priority = 2
	}
}

func (i *Item) Label() string {
	return strconv.Itoa(i.position) + "."
}

func (i *Item) PPrinting() string {
	if i.Priority == 1 {
		return "(1)"
	} else if i.Priority == 3{
		return "(3)"
	} else {
		return " "
	}
}

func SaveItems(filename string, items []Item) error {

	b, err := json.Marshal(items)
	err = ioutil.WriteFile(filename, b, 0644)

	if err != nil {
		return err
	}


	return nil
}

func ReadItems(filename string) ([]Item, error) {

	b, err := ioutil.ReadFile(filename)

	if err != nil {
		return []Item{}, err
	}

	var items []Item

	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}

	for i, _ := range items {
		items[i].position = i + 1
	}

	return items, nil
}

func (s ByPri) Len() int {
	return len(s)
}

func (s ByPri) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByPri) Less(i, j int) bool{
	if s[i].Priority == s[j].Priority {
		return s[i].position < s[j].position
	} else {
		return s[i].Priority < s[j].Priority
	}
}