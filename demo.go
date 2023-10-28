package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	ID    int
	Name  string
	Done  bool
}

var tasks []Task
var currentID = 1

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nTODO CLI")
		fmt.Println("1. Add task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Mark task as done")
		fmt.Println("4. Remove task")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter task name: ")
			scanner.Scan()
			taskName := scanner.Text()

			tasks = append(tasks, Task{ID: currentID, Name: taskName})
			currentID++
		case "2":
			for _, task := range tasks {
				status := "Pending"
				if task.Done {
					status = "Done"
				}
				fmt.Printf("%d. %s [%s]\n", task.ID, task.Name, status)
			}
		case "3":
			fmt.Print("Enter task ID to mark as done: ")
			scanner.Scan()
			var id int
			fmt.Sscan(scanner.Text(), &id)

			for i, task := range tasks {
				if task.ID == id {
					tasks[i].Done = true
					break
				}
			}
		case "4":
			fmt.Print("Enter task ID to remove: ")
			scanner.Scan()
			var id int
			fmt.Sscan(scanner.Text(), &id)

			for i, task := range tasks {
				if task.ID == id {
					tasks = append(tasks[:i], tasks[i+1:]...)
					break
				}
			}
		case "5":
			return
		default:
			fmt.Println("Invalid choice!")
		}
	}
}
