package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {
	// printSomething(1)
	// printSomething(1.5)
	// printSomething("any or interface to accept any type of value")

	// result := add[int](1, 2)

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	// printSomething(todo)

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	outputData(userNote)
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title:")
	content := getUserInput("Note Content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the todo failed:", err)
		return err
	}

	fmt.Println("Saving the todo succeeded!")
	return nil
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

// func printSomething(value any)
// func printSomething(value interface{}) {
// 	intVal, ok := value.(int)
// 	if ok {
// 		fmt.Println("Integer:", intVal)
// 		return
// 	}
// 	floatVal, ok := value.(float64)
// 	if ok {
// 		fmt.Println("Float:", floatVal)
// 		return
// 	}
// 	stringVal, ok := value.(string)
// 	if ok {
// 		fmt.Println("String:", stringVal)
// 		return
// 	}
// 	// switch value.(type) {
// 	// case int:
// 	// 	fmt.Println("Integer:", value)
// 	// case float64:
// 	// 	fmt.Println("Float:", value)
// 	// case string:
// 	// 	fmt.Println("String:", value)
// 	// default:
// 	// 	fmt.Println("Any:", value)
// 	// }
// }

// func add[T int | float64 | string](a, b T) T {
// 	// aInt, aIsInt := a.(int)
// 	// bInt, bIsInt := b.(int)
// 	// if aIsInt && bIsInt {
// 	// 	return aInt + bInt
// 	// }
// 	// aFloat, aIsFloat := a.(float64)
// 	// bFloat, bIsFloat := b.(float64)
// 	// if aIsFloat && bIsFloat {
// 	// 	return aFloat + bFloat
// 	// }
// 	// aString, aIsString := a.(string)
// 	// bString, bIsString := b.(string)
// 	// if aIsString && bIsString {
// 	// 	return aString + bString
// 	// }
// 	return a + b
// }
