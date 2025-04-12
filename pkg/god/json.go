package god

import (
	"encoding/json"
	"net/http"
)

type JSON struct {
	Data any
}

const (
	jsonContentType = "application/json"
)

// god.H is a alias for hashmap (map[string]interface{}).
// It simplifies the creation of JSON responses.
type H map[string]interface{}

func (r *JSON) Render(code int, w http.ResponseWriter) error {
	return r.WriteJSONResponse(code, w)
}

func (r *JSON) WriteJSONResponse(code int, w http.ResponseWriter) error {
	r.WriteContentType(w)

	data, err := json.Marshal(r.Data)
	if err != nil {
		return err
	}

	w.WriteHeader(code)
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}
func (r *JSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(jsonContentType, w)
}
