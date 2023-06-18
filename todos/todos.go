package main

import "fmt"

type TodoList struct {
	tasks []*Todo
}

func (t *TodoList) AddTask(task *Todo) {
	t.tasks = append(t.tasks, task)
}

func (t *TodoList) RemoveTask(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *TodoList) printTasks() {
	fmt.Println("✅ TO DO's")

	for i, t := range t.tasks {
		var check string

		if t.done {
			check = "[x]"
		} else {
			check = "[ ]"
		}

		fmt.Printf("%d. %s %s\n       %s\n", i+1, check, t.name, t.description)
	}
}

type Todo struct {
	description string
	done        bool
	name        string
}

func (t *Todo) ToggleDone() {
	t.done = !t.done
}

func (t *Todo) UpdateDescription(description string) {
	t.description = description
}

func (t *Todo) UpdateName(name string) {
	t.name = name
}

func main() {
	t1 := &Todo{
		description: "Es el curso practico: creacion de un web server",
		done:        false,
		name:        "Completar mi curso de Go",
	}

	t2 := &Todo{
		description: "Es el curso practico: creacion de un web server",
		done:        false,
		name:        "Completar mi curso de Python",
	}

	fmt.Printf("%+v\n", t1)

	t1.ToggleDone()
	t1.UpdateName("Finalizar mi curso de GO")
	t1.UpdateDescription("Completar mi curso de cuanto antes")

	fmt.Printf("%+v\n", t1)

	tl := &TodoList{
		tasks: []*Todo{
			t1,
			t2,
		},
	}

	t3 := &Todo{
		description: "Es el curso practico: creacion de un web server",
		done:        false,
		name:        "Completar mi curso de React",
	}

	tl.AddTask(t3)
	tl.RemoveTask(1)
	tl.printTasks()
}
