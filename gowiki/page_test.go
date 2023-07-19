package main

import (
	"fmt"
	"testing"
)

func TestSave(t *testing.T) {
	page := &Page{Title: "welcome", Body: []byte("to the jungle")}

	err := page.save()
	if err != nil {
		t.Errorf("wrong behaveor")
	}
}

func TestLoadPage(t *testing.T) {
	page, err := LoadPage("welcome")

	if err != nil {
		t.Errorf("we got %v", err)
	}

	fmt.Printf("we got rigth title: %v and body: %v", page.Title, page.Body)
}

func TestFailLoadPage(t *testing.T) {
	page, err := LoadPage("evangelion")
	if err == nil {
		fmt.Printf("we shouldn't get %v or %v", page.Title, page.Body)
		t.Errorf("should throw an error")
	}

	fmt.Printf("the error ge got was %v", err)

}
