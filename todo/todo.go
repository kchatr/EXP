package todo

import (
	"os"
	"strconv"
	"encoding/json"
	"io/ioutil"
)

type Item struct {
	Text string
	Priority int
	Position int
	Done bool
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
	return strconv.Itoa(i.Position) + "."
}

func (i *Item) PrettyPrint() string {
	if i.Priority == 1 {
		return "(1)-High"
	} else if i.Priority == 3{
		return "(3)-Low"
	} else {
		return "(2)-Med" 
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

	if _, file_err := os.Stat(filename); os.IsNotExist(file_err) {
		return []Item{}, nil
	} 

	b, err := ioutil.ReadFile(filename)

	if err != nil {
		return []Item{}, err
	}

	var items []Item

	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}

	for i, _ := range items {
		items[i].Position = i + 1
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
	i1 := s[i]
	i2 := s[j]

	if i1.Done != i2.Done {
		return i1.Done
	} else if i1.Priority != i2.Priority{
		return i1.Priority < i2.Priority
	} else {
		return i1.Position < i2.Position
	}
}

func (i *Item) PrettyDone() string {
	if i.Done == true {
		return "✓"
	} else {
		return "X"
	}
}