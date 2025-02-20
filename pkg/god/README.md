# God - A Minimalist Web Framework for Go

**God** is a lightweight and minimalist web framework for Go, designed to simplify the process of building web applications and APIs. It provides essential features like routing, middleware support, and context management, making it easy to create robust and scalable web services.

## Installation

To install the **God** framework, use the following command:

```bash
go get github.com/sherinur/coffee-shop/pkg/god
```

## Getting Started

### Initializing the Router

To start using **God**, you need to initialize a new router instance. The router is the core of the framework, handling all incoming HTTP requests and routing them to the appropriate handlers.

```go
package main

import (
	"github.com/sherinur/coffee-shop/pkg/god"
)

func main() {
	// Create a new router instance
	router := god.Default()

	// Define a route
	router.GET("/hello", func(c *god.Context) {
		c.JSON(http.StatusOK, god.H{"message": "Hello, World!"})
	})

	// Start the server
	router.Run(":8080")
}
```

### Using Context

The `Context` type in **God** is used to pass request-specific information between handlers. It provides methods to access the request, response, and other metadata.

```go
func handler(c *god.Context) {
	// Access request parameters
	name := c.Params["name"]

	// Set a response status code
	c.Status(http.StatusOK)

	// Send a JSON response
	c.JSON(http.StatusOK, god.H{"name": name})
}
```

## Framework Structure

### Types and Instances

1. **Context (`god.Context`)**
   - **Purpose**: The `Context` type encapsulates the HTTP request and response, providing methods to interact with them.
   - **Key Methods**:
     - `Next()`: Calls the next handler in the chain.
     - `JSON(code int, obj any)`: Sends a JSON response.
     - `Status(code int)`: Sets the HTTP status code.
     - `Set(key string, value any)`: Stores a key/value pair in the context.
     - `Get(key string) (value any, exists bool)`: Retrieves a value from the context.

2. **Router (`god.Router`)**
   - **Purpose**: The `Router` type is responsible for routing incoming HTTP requests to the appropriate handlers.
   - **Key Methods**:
     - `Handle(method, path string, handlers ...HandlerFunc)`: Registers a new route with a method and path.
     - `ServeHTTP(w http.ResponseWriter, req *http.Request)`: Implements the `http.Handler` interface.
     - `GET(path string, handler HandlerFunc)`: Registers a GET route.
     - `POST(path string, handler HandlerFunc)`: Registers a POST route.
     - `PUT(path string, handler HandlerFunc)`: Registers a PUT route.
     - `DELETE(path string, handler HandlerFunc)`: Registers a DELETE route.
     - `Run(addr string) error`: Starts the HTTP server.

3. **JSON (`god.JSON`)**
   - **Purpose**: The `JSON` type is used to render JSON responses.
   - **Key Methods**:
     - `Render(code int, w http.ResponseWriter) error`: Renders the JSON response.
     - `WriteJSONResponse(code int, w http.ResponseWriter) error`: Writes the JSON response to the HTTP response writer.

4. **HandlersChain (`god.HandlersChain`)**
   - **Purpose**: A slice of `HandlerFunc` that represents a chain of handlers to be executed in sequence.

5. **HandlerFunc (`god.HandlerFunc`)**
   - **Purpose**: A function type that handles HTTP requests. It takes a `Context` as its only argument.

### Directory Structure

The **God** framework is organized as follows:

```
god/
├── context.go       # Contains the Context type and related methods
├── json.go          # Contains the JSON type and related methods
├── README.md        # Documentation for the framework
├── router.go        # Contains the Router type and related methods
└── utils.go         # Utility functions
```

## Example Usage

Here’s a complete example demonstrating how to use the **God** framework to create a simple web server:

```go
package main

import (
	"net/http"

	"github.com/sherinur/coffee-shop/pkg/god"
)

func main() {
	router := god.Default()

	router.GET("/hello", func(c *god.Context) {
		c.JSON(http.StatusOK, god.H{"message": "Hello, World!"})
	})

	router.POST("/greet", func(c *god.Context) {
		name := c.Params["name"]
		c.JSON(http.StatusOK, god.H{"message": "Hello, " + name})
	})

	router.Run(":8080")
}
```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on the [GitHub repository](https://github.com/sherinur/coffee-shop/tree/main/pkg).

## Author

- **Nurislam Sheri**
  - GitHub: [sherinur](https://github.com/sherinur/)
  - Repository: [coffee-shop](https://github.com/sherinur/coffee-shop/tree/main/pkg)

---

**God** is designed to be simple yet powerful, providing the essential tools needed to build web applications in Go. Whether you're building a small API or a larger web service, **God** has you covered. Happy coding! 🚀