// mutexes
package goroutines

import "sync"
import "time"
import "fmt"

type SafeStringCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func NewSafeStringCounter() *SafeStringCounter {
	return &SafeStringCounter{v: make(map[string]int)}
}

// Inc increments the counter for the given key.
func (c *SafeStringCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeStringCounter) Dec(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]--
	c.mux.Unlock()
}

func (c *SafeStringCounter) Count(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func Mutexes() {

	c := NewSafeStringCounter()

	for i := 0; i < 100; i++ {

		go c.Inc("somekey")

	}

	time.Sleep(time.Second)
	fmt.Println("safe count", c.Count("somekey"))
}
