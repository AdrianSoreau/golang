package main

import (
	"fmt"
	"sort"
	"estiam-main/dictionary"
)

func main() {
	d := dictionary.New()

	d.Add("Gopher", "A person who uses the Go programming language.")
	d.Add("Variable", "A data storage element in programming.")
	d.Add("Function", "A block of code that performs a specific task.")

	word := "Gopher"
	entry, err := d.Get(word)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Definition of '%s': %s\n", word, entry)
	}

	err = d.Remove("Variable")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Word 'Variable' removed.")
	}

	words, entries := d.List()
	sort.Strings(words)
	fmt.Println("\nList of words and their definitions:")
	for _, w := range words {
		fmt.Printf("%s : %s\n", w, entries[w])
	}
}
