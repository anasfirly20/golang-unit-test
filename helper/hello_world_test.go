package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldAnas(t *testing.T) {
	result := HelloWorld("Anas")

	if result != "Hello Anas" {
		// error
		t.Error("Result must be 'Hello Anas'")
	}

	fmt.Println("TestHelloWorldAnas Done")
}

func TestHelloWorldFirly(t *testing.T) {
	result := HelloWorld("Firly")

	if result != "Hello Firly" {
		// error
		t.Fatal("Result must be 'Hello Firly'")
	}
	
	fmt.Println("TestHelloWorldFirly Done")
}

