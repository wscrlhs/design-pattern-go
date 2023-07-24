package singleton_test

import (
	"design-pattern/singleton"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetInstance(t *testing.T) {
	single := singleton.GetInstance()
	assert.NotEqual(t, nil, single)

	single = singleton.GetInstance()
	assert.NotEqual(t, nil, single)

	singleton.Clear()
}

func TestGetInstance2(t *testing.T) {
	single2 := singleton.GetInstance2()
	assert.NotEqual(t, nil, single2)

	single2 = singleton.GetInstance2()
	assert.NotEqual(t, nil, single2)

	singleton.Clear2()
}

func ExampleGetInstance() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			singleton.GetInstance()
		}()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()

	wg.Wait()

	// Output:
	// Creating single instance now.
	// Single instance already created.
	// Single instance already created.
	// Single instance already created.
	// Single instance already created.
	// Single instance already created.
	// Single instance already created.
	// Single instance already created.
	// Single instance already created.
	// Single instance already created.

}

func ExampleGetInstance2() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
			singleton.GetInstance2()
		}()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()

	wg.Wait()

	// Output:
	// Creating single2 instance now.
	// Single2 instance already created.
	// Single2 instance already created.
	// Single2 instance already created.
	// Single2 instance already created.
	// Single2 instance already created.
	// Single2 instance already created.
	// Single2 instance already created.
	// Single2 instance already created.
	// Single2 instance already created.
}
