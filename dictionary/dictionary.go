package dictionary

import (
	"encoding/json"
	"errors"
	"os"
)

type Entry struct {
	Definition string
}

type Dictionary struct {
	filePath string
}

func New(filePath string) *Dictionary {
	return &Dictionary{filePath: filePath}
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

func (d *Dictionary) Add(word, definition string) error {
	entries, err := d.load()
	if err != nil {
		return err
	}

	if _, exists := entries[word]; exists {
		return errors.New("word already exists")
	}

	entries[word] = Entry{Definition: definition}
	return d.save(entries)
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

func (d *Dictionary) Remove(word string) error {
	entries, err := d.load()
	if err != nil {
		return err
	}

	if _, exists := entries[word]; !exists {
		return errors.New("word does not exist")
	}

	delete(entries, word)
	return d.save(entries)
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
