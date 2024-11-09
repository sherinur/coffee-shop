package handler

// type OrderHandler interface {
// 	CreateOrder(w http.ResponseWriter, r *http.Request)
// 	RetrieveOrders(w http.ResponseWriter, r *http.Request)
// 	RetrieveOrder(w http.ResponseWriter, r *http.Request)
// 	UpdateOrder(w http.ResponseWriter, r *http.Request)
// 	DeleteOrder(w http.ResponseWriter, r *http.Request)
// 	CloseOrder(w http.ResponseWriter, r *http.Request)
// }

// type orderHandler struct {
// 	OrderService *service.OrderService
// }

// func NewOrderHandler(s *service.OrderService) OrderHandler {
// 	return &orderHandler{OrderService: s}
// }

// func (h *orderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
// 	// TODO: implement logic to Create a new order.
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("There will be ORDER CREATING"))
// }

// func (h *orderHandler) RetrieveOrders(w http.ResponseWriter, r *http.Request) {
// 	// TODO: implement logic to Retrieve all orders.
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("There will be Retrieving all orders."))
// }

// func (h *orderHandler) RetrieveOrder(w http.ResponseWriter, r *http.Request) {
// 	orderId := r.PathValue("id")

// 	// TODO: implement logic to Retrieve a specific order by ID.
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("There will be Retrieve a specific order by ID: " + orderId))
// }

// func (h *orderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
// 	orderId := r.PathValue("id")

// 	// TODO: implement logic to Update an existing order by ID.
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("There will be Update an existing order by ID: " + orderId))
// }

// func (h *orderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
// 	orderId := r.PathValue("id")

// 	// TODO: implement logic to Delete an order by ID.
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("There will be Delete an order by ID: " + orderId))
// }

// func (h *orderHandler) CloseOrder(w http.ResponseWriter, r *http.Request) {
// 	orderId := r.PathValue("id")

// 	// TODO: implement logic to Close an order by ID.
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("There will be Close an order by ID: " + orderId))
// }
