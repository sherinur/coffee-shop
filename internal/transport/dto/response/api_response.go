package response

type APIResponse struct {
	Status int `json:"status"`
	Body   any `json:"body"`
}
