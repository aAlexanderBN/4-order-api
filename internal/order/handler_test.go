package order_test

import (
	"bytes"
	"encoding/json"
	"go/api/configs"
	"go/api/internal/myuser"
	"go/api/internal/order"
	"go/api/internal/product"
	"go/api/pkg/db"
	"go/api/pkg/jwt"
	"go/api/pkg/middleware"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb() *gorm.DB {
	err := godotenv.Load("internal/testing/.env")

	if err != nil {
		panic("Error loading .env file, using dafault config")
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})

	return db
}

func setupTestDB(t *testing.T) (*gorm.DB, *configs.Config, func()) {

	conf := configs.LoadTestConfig()
	db := db.NewDb(conf)

	// Миграции для тестовой БД
	db.AutoMigrate(&myuser.Users{}, &product.Product{}, &order.Order{}, &order.OrderProduct{})

	cleanup := func() {
		// Очистка после теста
		db.Exec("DELETE FROM order_products")
		db.Exec("DELETE FROM orders")
		db.Exec("DELETE FROM products")
		db.Exec("DELETE FROM users")
	}

	return db.DB, conf, cleanup
}

func createTestUser(t *testing.T, db *gorm.DB) *myuser.Users {
	user := &myuser.Users{
		Phone: "+79991111111",
		Name:  "Test User",
	}
	db.Create(user)
	return user
}

func createTestProducts(t *testing.T, db *gorm.DB) []product.Product {
	products := []product.Product{
		{Name: "Product 1", Description: "Test"},
		{Name: "Product 2", Description: "Test"},
	}
	db.Create(&products)
	return products
}

func TestOrder(t *testing.T) {
	db, conf, cleanup := setupTestDB(t)
	defer cleanup() // Удаление данных после теста

	user := createTestUser(t, db)
	products := createTestProducts(t, db)

	// 4. Создание тестового роутера с реальными handlers
	router := http.NewServeMux()
	userRepo := myuser.NewUserRepository(db)
	orderRepo := order.NewOrderRepository(db)

	order.NewOrderHandler(router, order.OrderHandlerDeps{
		OrderRepository: orderRepo,
		Config:          conf,
		UserRepository:  userRepo,
	})

	// 5. Создание тестового сервера
	ts := httptest.NewServer(middleware.Logging(router))
	defer ts.Close()

	// Создать JWT токен для пользователя
	token, _ := jwt.NewJWT("/2+XnmJGz1j3ehIVI/5P9kl+CghrE3DcS7rnT+qar5w=").Create(user.Phone)

	orderData := order.Order{
		UserID: 1,
		Products: []order.OrderProduct{
			{ProductID: products[0].ID, Quantity: 1},
			{ProductID: products[1].ID, Quantity: 3},
		},
	}

	jsonData, err := json.Marshal(orderData)
	if err != nil {
		t.Fatal(err)
	}

	// Формирование запроса с токеном
	req, err := http.NewRequest("POST", ts.URL+"/order", bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 201 {
		t.Errorf("got %d, expected %d", resp.StatusCode, 201)
	}

	var orderread order.Order

	err = json.Unmarshal(body, &orderread)
	if err != nil {
		t.Fatal(err)
	}

	if orderread.Products[0].ProductID != orderData.Products[0].ProductID {
		t.Errorf("Product got %d, expected %d", orderread.Products[0].ProductID, orderData.Products[0].ProductID)
	}

}
