package dictionary

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

type Entry struct {
	Definition string
}

type Dictionary struct {
	filePath     string
	addChan      chan addRequest
	removeChan   chan string
	wg           sync.WaitGroup
}

type addRequest struct {
	word       string
	definition string
	result     chan error
}

func New(filePath string) *Dictionary {
	d := &Dictionary{
		filePath:   filePath,
		addChan:    make(chan addRequest),
		removeChan: make(chan string),
	}

	go d.processRequests()

	return d
}

func (d *Dictionary) processRequests() {
	entries, err := d.load()
	if err != nil {
	}

	for {
		select {
		case addReq := <-d.addChan:
			if _, exists := entries[addReq.word]; exists {
				addReq.result <- errors.New("word already exists")
			} else {
				entries[addReq.word] = Entry{Definition: addReq.definition}
				addReq.result <- d.save(entries)
			}

		case word := <-d.removeChan:
			if _, exists := entries[word]; !exists {
			} else {
				delete(entries, word)
				d.save(entries)
			}
		}
	}
}

func (d *Dictionary) Add(word, definition string) error {
	resultChan := make(chan error)
	d.addChan <- addRequest{word, definition, resultChan}
	return <-resultChan
}

func (d *Dictionary) Remove(word string) error {
	d.removeChan <- word
	return nil
}

func (d *Dictionary) Get(word string) (Entry, error) {
	entries, err := d.load()
	if err != nil {
		return Entry{}, err
	}

	entry, exists := entries[word]
	if !exists {
		return Entry{}, errors.New("word does not exist")
	}

	return entry, nil
}

func (d *Dictionary) List() ([]string, error) {
	entries, err := d.load()
	if err != nil {
		return nil, err
	}

	var words []string
	for word := range entries {
		words = append(words, word)
	}
	return words, nil
}

func (d *Dictionary) load() (map[string]Entry, error) {
	file, err := os.Open(d.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]Entry), nil
		}
		return nil, err
	}
	defer file.Close()

	var entries map[string]Entry
	err = json.NewDecoder(file).Decode(&entries)
	if err != nil {
		return make(map[string]Entry), nil
	}
	return entries, nil
}

func (d *Dictionary) save(entries map[string]Entry) error {
	file, err := os.Create(d.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(entries)
}
