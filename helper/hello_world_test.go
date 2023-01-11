package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmark := []struct {
		name    string
		request string
	}{
		{
			name:    "Aldi",
			request: "Aldi",
		},
		{
			name:    "Tegar",
			request: "Tegar",
		},
	}

	for _, bencbenchmark := range benchmark {
		b.Run(bencbenchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				helloWorld(bencbenchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Aldi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			helloWorld("Aldi")
		}
	})
	b.Run("Tegar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			helloWorld("Tegar")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helloWorld("Aldi")
	}
}
func BenchmarkHelloWorldTegar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helloWorld("Tegar")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Aldi",
			request:  "Aldi",
			expected: "Hello Aldi",
		},
		{
			name:     "Tegar",
			request:  "Tegar",
			expected: "Hello Tegar",
		},
		{
			name:     "Prakoso",
			request:  "Prakoso",
			expected: "Hello Prakoso",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := helloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Aldi", func(t *testing.T) {
		result := helloWorld("Aldi")
		require.Equal(t, "Hello Aldi", result, "Result must be 'Hello Aldi'")
	})
	t.Run("Tegar", func(t *testing.T) {
		result := helloWorld("Tegar")
		require.Equal(t, "Hi Tegar", result, "Result must be 'Hi Tegar'")
	})
}

func TestMain(m *testing.M) {
	// Before Test
	fmt.Println("Before Unit Test")

	m.Run()

	// After Test
	fmt.Println("Before Unit Test")
}

func TestHelloWorld(t *testing.T) {
	result := helloWorld("Aldi")
	if result != "Hello Aldi" {
		panic("Result is not 'Hello Aldi'")
	}

	fmt.Println("TestHelloWorldAldi Done")
}
func TestHelloWorldTegar(t *testing.T) {
	result := helloWorld("Tegar")
	if result != "Hello Tegar" {
		panic("Result is not 'Hello Tegar'")
	}

	fmt.Println("TestHelloWorld Tegar Done")
}

func TestHiWorldFail(t *testing.T) {
	result := helloWorld("Prakoso")
	if result != "Hi Prakoso" {
		t.Fail()
	}

	fmt.Println("TestHiWorld Prakoso Done")
}
func TestHiWorldFailNow(t *testing.T) {
	result := helloWorld("Prakoso")
	if result != "Hi Prakoso" {
		t.FailNow()
	}

	fmt.Println("TestHiWorld Prakoso Done")
}
func TestHiWorldError(t *testing.T) {
	result := helloWorld("Prakoso")
	if result != "Hi Prakoso" {
		t.Error("Result must be 'Hi Prakoso'")
	}

	fmt.Println("TestHiWorldError Prakoso Done")
}
func TestHiWorldFatal(t *testing.T) {
	result := helloWorld("Prakoso")
	if result != "Hi Prakoso" {
		t.Fatal("Result must be 'Hi Prakoso'")
	}

	fmt.Println("TestHiWorldFatal Prakoso Done")
}

// Assetion
func TestHelloWorldAssertion(t *testing.T) {
	result := helloWorld("Prakoso")
	assert.Equal(t, "Hi Prakoso", result, "Result must be 'Hi Prakoso'") // Assert memanggil fail()
	fmt.Println("TestHelloWorld with Assertion Done")

}

// Require
func TestHelloWorldRequire(t *testing.T) {
	result := helloWorld("Prakoso")
	require.Equal(t, "Hi Prakoso", result, "Result must be 'Hi Prakoso'") // Require memanggil failNow()
	fmt.Println("TestHelloWorld with Require Done")

}

// Skip atau membatalkan test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Windows OS")
	}

	result := helloWorld("Prakoso")
	require.Equal(t, "Hi Prakoso", result, "Result must be 'Hi Prakoso'")
}
