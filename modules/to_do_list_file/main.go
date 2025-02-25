package to_do_list_file

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type todo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (t todo) String() string {
	return "Task - ID:" + strconv.Itoa(t.Id) + " Name:" + t.Name
}

var todos = []todo{}

var todoFile = "todos.txt"

func Run() {
	loadTodosFromFile()
	for {
		fmt.Println("Welcome to this todo list program")
		fmt.Println("1. List todos")
		fmt.Println("2. Add todo")
		fmt.Println("3. Delete todo")
		fmt.Println("4. Exit")
		fmt.Println("Seleccione una opción")

		//Defining reader from operative system
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		option, _ := strconv.Atoi(input[:len(input)-1])

		switch option {
		case 1:
			fmt.Println("LIST TODOS")
			listTodos()
		case 2:
			fmt.Println("ADD TODO")
			addTodo()
			saveTodosToFile()
		case 3:
			fmt.Println("DELETE TODO")
			deleteTodo()
			saveTodosToFile()
		case 4:
			fmt.Println("EXIT")
			fmt.Println("SAVING TODOS")
			return
		default:
			fmt.Println("Insert one valid value")
		}
	}

}

func loadTodosFromFile() {
	file, err := os.Open(todoFile)
	if err != nil {
		fmt.Println("No tasks in file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		part := strings.Split(line, ",")
		id, _ := strconv.Atoi(part[0])
		name := part[1]
		todo := todo{
			Id:   id,
			Name: name,
		}
		todos = append(todos, todo)
	}
}

func saveTodosToFile() {
	file, err := os.Create(todoFile)
	if err != nil {
		fmt.Println("Error creating file")
		return
	}
	defer file.Close()

	for _, todo := range todos {
		_, _ = file.WriteString(fmt.Sprintf("%d,%s\n", todo.Id, todo.Name))
	}
}

func listTodos() []todo {
	if len(todos) == 0 {
		fmt.Println("No tasks")
	}
	for _, todo := range todos {
		fmt.Println(todo)
	}
	return todos
}

func addTodo() {
	fmt.Println("Insert todo name")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	todo := todo{
		Id:   len(todos) + 1,
		Name: name,
	}
	todos = append(todos, todo)
	fmt.Println("Task added")
}

func deleteTodo() {
	fmt.Println("Insert todo id to delete")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	id, _ := strconv.Atoi(input)

	if _, err := findTodoById(id); err != nil {
		fmt.Println("Todo not deleted: Todo not found")
		return
	}

	todos = append(todos[:id], todos[id+1:]...)
}

func findTodoById(id int) (todo, error) {
	for _, todo := range todos {
		if todo.Id == id {
			return todo, nil
		}
	}
	return todo{}, fmt.Errorf("task not found")
}
