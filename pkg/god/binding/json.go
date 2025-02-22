package binding

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type jsonBinding struct{}

func (jsonBinding) Name() string {
	return "json"
}

func (jsonBinding) Bind(req *http.Request, obj any) error {
	return decodeJSON(req.Body, obj)
}

func (jsonBinding) BindBody(body []byte, obj any) error {
	return decodeJSON(bytes.NewReader(body), obj)
}

// TODO: Write the appropriate error messages
func decodeJSON(r io.Reader, obj any) error {
	dec := json.NewDecoder(r)
	return dec.Decode(obj)
}
