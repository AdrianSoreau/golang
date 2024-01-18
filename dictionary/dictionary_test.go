package dictionary

import (
    "testing"
)

func setupTest() *Dictionary {
    return New("test_dictionary.json")
}

func TestAdd(t *testing.T) {
    d := setupTest()

    err := d.Add("test", "A test word")
    if err != nil {
        t.Errorf("Add failed when it should succeed: %v", err)
    }

    err = d.Add("test", "A test word")
    if err == nil {
        t.Errorf("Add succeeded when it should fail")
    }
}

func TestGet(t *testing.T) {
    d := setupTest()
    d.Add("test", "A test word")

    _, err := d.Get("test")
    if err != nil {
        t.Errorf("Get failed when it should succeed: %v", err)
    }

    _, err = d.Get("nonexistent")
    if err == nil {
        t.Errorf("Get succeeded when it should fail")
    }
}

func TestRemove(t *testing.T) {
    d := setupTest()
    d.Add("test", "A test word")

    err := d.Remove("test")
    if err != nil {
        t.Errorf("Remove failed when it should succeed: %v", err)
    }

    err = d.Remove("nonexistent")
    if err == nil {
        t.Errorf("Remove succeeded when it should fail")
    }
}

func TestList(t *testing.T) {
    d := setupTest()
    d.Add("test", "A test word")

    words, err := d.List()
    if err != nil || len(words) == 0 {
        t.Errorf("List failed: %v", err)
    }
}
