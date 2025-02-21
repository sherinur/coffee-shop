package god

import (
	"fmt"
	"net/http"
	"sync"
)

type Context struct {
	Request *http.Request
	Writer  http.ResponseWriter

	Params   map[string]string
	handlers HandlersChain
	index    int

	// Keys is a key/value pair for the context of each request.
	Keys map[string]any

	// This mutex protects Keys map.
	mu sync.RWMutex

	fullPath string
}

type (
	HandlersChain []HandlerFunc
	HandlerFunc   func(*Context)
)

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Request: r,
		Writer:  w,
		Params:  make(map[string]string),
		index:   -1,
		Keys:    make(map[string]any),
	}
}

// Next calls the next handler in the chain.
func (c *Context) Next() {
	c.index++
	for c.index < len(c.handlers) {
		c.handlers[c.index](c)
		c.index++
	}
}

// JSON sends a JSON response.
func (c *Context) JSON(code int, obj any) {
	json := &JSON{Data: obj}
	err := json.Render(code, c.Writer)
	if err != nil {
		fmt.Println("Error of rendering JSON response:", err)
	}
}

func (c *Context) Status(code int) {
	writeStatusCode(code, c.Writer)
}

// Get is used to store a new key/value pair for this context.
func (c *Context) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.Keys == nil {
		c.Keys = make(map[string]any)
	}

	c.Keys[key] = value
}

// Get returns the value and existence of the key.
// If the value does not exist it returns (nil, false)
func (c *Context) Get(key string) (value any, exists bool) {
	c.mu.RLock()
	defer c.mu.Unlock()

	value, exists = c.Keys[key]

	return value, exists
}

func (c *Context) FullPath() (fullPath string) {
	return c.fullPath
}
