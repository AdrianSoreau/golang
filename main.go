package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"estiam-main/dictionary"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	d := dictionary.New("dictionary.json")

	for {
		fmt.Println("\nChoose an action [add, define, remove, list, exit]:")
		action, _ := reader.ReadString('\n')
		action = strings.TrimSpace(action)

		switch action {
		case "add":
			actionAddConcurrent(d, reader)
		case "define":
			actionDefineConcurrent(d, reader)
		case "remove":
			actionRemoveConcurrent(d, reader)
		case "list":
			actionList(d)
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Action not recognized.")
		}
	}
}

func actionAddConcurrent(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter word: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)

	fmt.Print("Enter definition: ")
	definition, _ := reader.ReadString('\n')
	definition = strings.TrimSpace(definition)

	go func() {
		err := d.Add(word, definition)
		if err != nil {
			fmt.Println("Failed to add word:", err)
		} else {
			fmt.Println("Word added.")
		}
	}()
}

func actionDefineConcurrent(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter word: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)

	go func() {
		entry, err := d.Get(word)
		if err != nil {
			fmt.Println("Failed to find word:", err)
		} else {
			fmt.Println("Definition:", entry)
		}
	}()
}

func actionRemoveConcurrent(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter word to remove: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)

	go func() {
		err := d.Remove(word)
		if err != nil {
			fmt.Println("Failed to remove word:", err)
		} else {
			fmt.Println("Word removed.")
		}
	}()
}

func actionList(d *dictionary.Dictionary) {
	words, err := d.List()
	if err != nil {
		fmt.Println("Error listing words:", err)
		return
	}

	for _, word := range words {
		entry, err := d.Get(word)
		if err != nil {
			fmt.Println("Error getting definition for word:", word, "Error:", err)
			continue
		}
		fmt.Println(word, ":", entry)
	}
}
