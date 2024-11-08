package handler

import "net/http"

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	// TODO: implement logic to Create a new order.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("There will be ORDER CREATING"))
}

func RetrieveOrders(w http.ResponseWriter, r *http.Request) {
	// TODO: implement logic to Retrieve all orders.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("There will be Retrieving all orders."))
}

func RetrieveOrder(w http.ResponseWriter, r *http.Request) {
	orderId := r.PathValue("id")

	// TODO: implement logic to Retrieve a specific order by ID.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("There will be Retrieve a specific order by ID: " + orderId))
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	orderId := r.PathValue("id")

	// TODO: implement logic to Update an existing order by ID.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("There will be Update an existing order by ID: " + orderId))
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	orderId := r.PathValue("id")

	// TODO: implement logic to Delete an order by ID.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("There will be Delete an order by ID: " + orderId))
}

func CloseOrder(w http.ResponseWriter, r *http.Request) {
	orderId := r.PathValue("id")

	// TODO: implement logic to Close an order by ID.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("There will be Close an order by ID: " + orderId))
}
