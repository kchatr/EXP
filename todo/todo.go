/*
Copyright © 2020 Kaushik Chatterjee <kchatr1729@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package todo

import (
	"os"
	"strconv"
	"encoding/json"
	"io/ioutil"
)

/*
A struct is analagous to a class and represents a data structure containing a collection of properties associated with a To-Do item
*/
type Item struct {
	Text string
	Priority int
	Position int
	Done bool
}

type ByPri []Item

/*
Sets the priority of the desired To-Do (1 being the highest and 3 being the lowest; the default is 2).
*/
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

/*
An example of procedural abstraction: A helper method that provides the To-Do with a position label.
*/
func (i *Item) Label() string {
	return strconv.Itoa(i.Position) + "."
}

/*
An example of procedural abstraction: A helper method to make the position printing prettier.
*/
func (i *Item) PrettyPrint() string {
	if i.Priority == 1 {
		return "(1)-High"
	} else if i.Priority == 3 {
		return "(3)-Low"
	} else {
		return "(2)-Med" 
	}
}

/*
An example of procedural abstraction: A helper method to make the done printing prettier.
*/
func (i *Item) PrettyDone() string {
	if i.Done == true {
		return "✓"
	} else {
		return "X"
	}
}


/*
An example of procedural abstraction: A helper method to save To-Do items to the local file storing all user To-Do.
*/
func SaveItems(filename string, items []Item) error {

	b, err := json.Marshal(items)
	err = ioutil.WriteFile(filename, b, 0644)

	if err != nil {
		return err
	}


	return nil
}

/*
An example of procedural abstraction: A helper method to read & return To-Do items from the local file storing all user To-Dos.
*/
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


/*
A custom sorting algorithm that sorts the To-Do list based on the criteria of completion, priority, and finally position.
Uses iteration, sequencing, and selection in the algorithm; goes through To-Do list and swaps items if the first is 'lesser'
runs in a time complexity of O(n^2) and space complexity of O(n).
*/
func Sort(items []Item) {
	for i := 0; i < len(items) - 1; i++ {
		for j := i + 1; j < len(items); j++ {

		less := less(items, i, j)
		
			if !less {
				swap(items, i, j)
			}

		}

	}
}

/*
swap() is used as a helper methon in Sort(); an example of top down development and stepwise refinement.
Method is only available in todo.go i.e. not imported, so it is also an example of data hiding.
*/
func swap(s []Item, i, j int) {
	s[i], s[j] = s[j], s[i]
}

/*
less() is used as a helper methon in Sort(); an example of top down development and stepwise refinement.
Method is only available in todo.go i.e. not imported, so it is also an example of data hiding.
*/
func less(s []Item, i, j int) bool{
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