package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type test struct {
	name string
	request string
	expected string
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Anas")
	}
}

func BenchmarkHelloWorldFirly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Muhammad")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []test{
		{
			name: "Anas",
			request: "Anas",
			expected: "Hello Anas",
		},
		{
			name: "Firly",
			request: "Firly",
			expected: "Hello Firly",
		},
		{
			name: "Thomi",
			request: "Thomi",
			expected: "Hello Thomi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			expectedMessage := fmt.Sprintf("Result must be 'Hello %s'", test.name)
			require.Equal(t, test.expected, result, expectedMessage)
		})
	}
	
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestSubTest(t *testing.T) {
	t.Run("Anas", func (t *testing.T)  {
		result := HelloWorld("Anas")
		require.Equal(t, "Hello Anas", result, "Result must be 'Hello Anas'")
	})
	t.Run("Firly", func (t *testing.T)  {
		result := HelloWorld("Firly")
		require.Equal(t, "Hello Firly", result, "Result must be 'Hello Firly'")
	})
}

func TestSkip(t *testing.T){
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run on MAC OS")
	}

	result := HelloWorld("Firly")
	require.Equal(t, "Hello Firly", result, "Result must be 'Hello Firly'")
}

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

