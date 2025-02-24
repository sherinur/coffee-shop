package binding

import "net/http"

const (
	MIMEJSON  = "application/json"
	MIMEHTML  = "text/html"
	MIMEXML   = "application/xml"
	MIMEPlain = "text/plain"
	MIMEYAML  = "application/x-yaml"
	MIMETOML  = "application/toml"
)

type Binding interface {
	Name() string
	Bind(*http.Request, any) error
}

type BindingBody interface {
	Binding
	BindBody([]byte, any) error
}

var JSON BindingBody = jsonBinding{}

// TODO: Implement validator struct and interface
