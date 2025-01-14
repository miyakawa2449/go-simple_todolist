package main

import (
	"bufio"
	"fmt"
	"os"
)

type ToDo struct {
	ID          int
	Description string
	Done        bool
}

var todos []ToDo
var nextID int = 1

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- ToDo List ---")
		fmt.Println("1. Add ToDo")
		fmt.Println("2. List ToDos")
		fmt.Println("3. Mark ToDo as Done")
		fmt.Println("4. Delete ToDo")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			addToDo(scanner)
		case "2":
			listToDos()
		case "3":
			markToDoAsDone(scanner)
		case "4":
			deleteToDo(scanner)
		case "5":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

func addToDo(scanner *bufio.Scanner) {
	fmt.Print("Enter the ToDo description: ")
	scanner.Scan()
	description := scanner.Text()
	todo := ToDo{
		ID:          nextID,
		Description: description,
		Done:        false,
	}
	todos = append(todos, todo)
	nextID++
	fmt.Println("ToDo added successfully!")
}

func listToDos() {
	if len(todos) == 0 {
		fmt.Println("No ToDos available.")
		return
	}
	fmt.Println("\nCurrent ToDos:")
	for _, todo := range todos {
		status := "Not Done"
		if todo.Done {
			status = "Done"
		}
		fmt.Printf("%d. [%s] %s\n", todo.ID, status, todo.Description)
	}
}

func markToDoAsDone(scanner *bufio.Scanner) {
	fmt.Print("Enter the ID of the ToDo to mark as done: ")
	scanner.Scan()
	var id int
	fmt.Sscanf(scanner.Text(), "%d", &id)

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Done = true
			fmt.Println("ToDo marked as done!")
			return
		}
	}
	fmt.Println("ToDo not found.")
}

func deleteToDo(scanner *bufio.Scanner) {
	fmt.Print("Enter the ID of the ToDo to delete: ")
	scanner.Scan()
	var id int
	fmt.Sscanf(scanner.Text(), "%d", &id)

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Println("ToDo deleted successfully!")
			return
		}
	}
	fmt.Println("ToDo not found.")
}
