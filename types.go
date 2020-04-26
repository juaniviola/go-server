package main

// Task type
type Task struct {
	ID          string `json:"id" xml:"id" form:"id" query:"id"`
	Name        string `json:"name" xml:"name" form:"name" query:"name"`
	TaskName    string `json:"taskname" xml:"taskname" form:"taskname" query:"taskname"`
	Description string `json:"description" xml:"description" form:"description" query:"description"`
}

// Tasks type
type Tasks struct {
	items []Task
}

// type methods
func (task *Tasks) addToSlice(t Task) {
	task.items = append(task.items, t)
}

func (task *Tasks) removeToSlice(id string) {
	for i, el := range task.items {
		if el.ID == id {
			task.items = append(task.items[:i], task.items[i+1:]...)
			break
		}
	}
}

func (task *Tasks) updateToSlice(id string, t Task) {
	for i, el := range task.items {
		if el.ID == id {
			task.items[i] = t
			break
		}
	}
}
