package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// use for before/after unit test
// only executed once on each package
func TestMain(m *testing.M) {
	fmt.Println("Start unit test")

	m.Run()

	fmt.Println("End unit test")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Wawan")
	if result != "Hello Wawan" {
		// t.Fail()

		t.Error("Result:", result)
	}

	// this still executed
	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldWindah(t *testing.T) {
	result := HelloWorld("Windah")
	if result != "Hello Windah" {
		// t.FailNow()

		t.Fatal("Result:", result)
	}

	// this not executed
	fmt.Println("TestHelloWorldWindah Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Wawan")
	assert.Equal(t, "Hello Wawan", result, "result must be 'Hello Wawan")
	fmt.Println("TestHelloWorldAssert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Wawan")
	require.Equal(t, "Hello Wawan", result, "result must be 'Hello Wawan")
	fmt.Println("TestHelloWorldRequire Done")
}

func TestHelloWorldSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Unit test can't run on Mac OS")
	}

	result := HelloWorld("Wawan")
	require.Equal(t, "Hello Wawan", result, "result must be 'Hello Wawan")
	fmt.Println("TestHelloWorldSkip Done")
}

func TestSubTest(t *testing.T) {
	t.Run("Wawan", func(t *testing.T) {
		result := HelloWorld("Wawan")
		require.Equal(t, "Hello Wawan", result)
	})
	t.Run("Windah", func(t *testing.T) {
		result := HelloWorld("Windah")
		require.Equal(t, "Hello Windah", result)
	})
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Wawan)",
			request:  "Wawan",
			expected: "Hello Wawan",
		},
		{
			name:     "HelloWorld(Windah)",
			request:  "Windah",
			expected: "Hello Windah",
		},
		{
			name:     "HelloWorld(Brando)",
			request:  "Brando",
			expected: "Hello Brando",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Wawan")
	}
}

func BenchmarkHelloWorldWindahBasudara(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Windah Basudara")
	}
}

func BenchmarkHelloWorlSub(b *testing.B) {
	b.Run("Wawan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Wawan")
		}
	})

	b.Run("Windah Basudara", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Windah Basudara")
		}
	})
}

func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "HelloWorld(Wawan)",
			request: "Wawan",
		},
		{
			name:    "HelloWorld(Windah)",
			request: "Windah",
		},
		{
			name:    "HelloWorld(Brando)",
			request: "Brando",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}
