package order_test

import (
	"bytes"
	"encoding/json"
	main "go/api/cmd"
	"go/api/internal/order"
	"net/http"
	"testing"
)

func TestOrder(t *testing.T) {

	//подготовака
	app := main.App()

	server := http.Server{
		Addr:    ":8081",
		Handler: app,
	}

	server.ListenAndServe()
	defer server.Close()

	//тест
	order := order.Order{
		UserID: 1,
		Products: []OrderProduct{
			{ProductID: 3, Quantity: 1},
			{ProductID: 4, Quantity: 2},
		},
	}

	// Преобразуем в JSON
	jsonData, err := json.Marshal(order)
	if err != nil {
		t.Fatal(err)
	}

	// Создаем запрос
	resp, err := http.Post(
		"http://localhost:8081/order",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Читаем ответ
	// body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 201 {
		t.Errorf("got %d, expected %d", resp.StatusCode, 201)
	}

}
