package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Firly")
	require.Equal(t, "Hello Firly", result, "Result must be 'Hello Firly'")
	fmt.Println("TestHelloWorldRequire executed")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Firly")
	assert.Equal(t, "Hello Firly", result, "Result must be 'Hello Firly'")
	fmt.Println("TestHelloWorldAssert executed")
}

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

