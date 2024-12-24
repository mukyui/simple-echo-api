package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Struct สำหรับตอบข้อมูล JSON
type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// Handler สำหรับ Root Endpoint
func homeHandler(c echo.Context) error {
	response := Response{
		Message: "Welcome to Echo API!",
		Status:  200,
	}
	return c.JSON(http.StatusOK, response)
}
func helloHandler(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		name = "World"
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello, " + name + "!",
	})
}
func main() {
	// สร้าง Instance ของ Echo
	e := echo.New()

	// Route
	e.GET("/", homeHandler)
	e.GET("/hello", helloHandler)
	// Start Server
	port := ":8080"
	e.Logger.Fatal(e.Start(port))
}
