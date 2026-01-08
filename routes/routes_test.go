package routes

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/xOphiuchus/OnlyDrive/middlewares"
	"github.com/xOphiuchus/OnlyDrive/services"
)

func TestHealthRoute(t *testing.T) {
	app := fiber.New()

	app.Use(middlewares.Recovery())
	app.Use(middlewares.Logger())
	app.Use(middlewares.CORS())

	app.Get("/healthz", func(c *fiber.Ctx) error {
		return services.HealthCheck(c)
	})

	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Failed to execute request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	if !strings.Contains(string(body), "healthy") {
		t.Errorf("Expected 'healthy' in response, got %q", string(body))
	}
}
