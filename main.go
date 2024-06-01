package main

import "fmt"

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos []Todo

func main() {
	tempTodo := Add("Make a breakfast")
	tempTodo.Complete()
	tempTodo.Edit("Make a coffee!")

	testTodo := Add("Make a breakfast")
	testTodo.Complete()
	testTodo.Complete()
	Remove(2)
	newTodo := Add("Make a breakfast")
	newTodo.Complete()
	fmt.Println(todos)
}

func (t *Todo) Edit(title string) {
	for i, todo := range todos {
		if todo.Id == t.Id {
			todos[i].Title = title
			break
		}
	}
}

func (t *Todo) Complete() {
	for i, todo := range todos {
		if todo.Id == t.Id {
			todos[i].Completed = !todos[i].Completed
			break
		}
	}
}

func Add(title string) Todo {
	todo := Todo{len(todos) + 1, title, false}
	todos = append(todos, todo)
	return todo
}

func Remove(id int) Todo {
	for i, todo := range todos {
		if todo.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return todo
		}
	}
	return Todo{}
}
