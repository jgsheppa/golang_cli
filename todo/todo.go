package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Item struct {
	Text string
	Priority int
	position int
	Done bool
}


func (i *Item) Label() string {
	return strconv.Itoa(i.position) + "."
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("This error will need to be fixed. CLI trying to open file before it exists.")
		return []Item{}, err
	}

	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}

	for i := range items {
		items[i].position = i + 1
	}
	return items, nil
}

func (i *Item) SetPriority(priority int) {
	switch priority {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

func (item *Item) PrettyPrint() string {
	if item.Priority == 1 {
		return "(1)"
	}
	if item.Priority == 3 {
		return "(3)"
	}

	return " "
}

func (item *Item) PrettyDone() string {
	if item.Done {
		return "X"
	}
	return ""
}

type ByPri []Item

func (s ByPri) Len() int      { return len(s) }
func (s ByPri) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPri) Less(i, j int) bool {
	if s[i].Done != s[j].Done {
		return s[i].Done
	}
	
		if s[i].Priority == s[j].Priority {
			return s[i].position < s[j].position
		}
		return s[i].Priority < s[j].Priority
	}