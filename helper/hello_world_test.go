package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldAnas(t *testing.T) {
	result := HelloWorld("Anas")

	if result != "Hello Anas" {
		// error
		t.Fail()
	}

	fmt.Println("TestHelloWorldAnas Done")
	
}

func TestHelloWorldFirly(t *testing.T) {
	result := HelloWorld("Firly")

	if result != "Hello Firly" {
		// error
		t.FailNow()
	}
	
	fmt.Println("TestHelloWorldFirly Done")
}

