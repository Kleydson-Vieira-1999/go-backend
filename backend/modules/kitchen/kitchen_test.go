package kitchen

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/store"
	"github.com/gin-gonic/gin"
)

func TestStreamKitchenEvents_MissingEstablishmentID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	_ = store.NewBroker()

	r := gin.New()
	r.GET("/stream", StreamKitchenEvents)

	server := httptest.NewServer(r)
	defer server.Close()

	resp, err := http.Get(server.URL + "/stream")
	if err != nil {
		t.Fatalf("failed to GET: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code 500, got %d", resp.StatusCode)
	}
}

func TestStreamKitchenEvents_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	broker := store.NewBroker()
	go broker.Start()

	r := gin.New()
	r.GET("/stream", StreamKitchenEvents)

	server := httptest.NewServer(r)
	defer server.Close()

	payload := `{"order_id": "123", "status": "pending"}`

	// Broadcast asynchronously to avoid blocking http.Get
	go func() {
		time.Sleep(100 * time.Millisecond)
		broker.Broadcast("test_est_id", payload)
	}()

	// Connect as client (blocks until headers are flushed by the async broadcast)
	resp, err := http.Get(server.URL + "/stream?establishment_id=test_est_id")
	if err != nil {
		t.Fatalf("failed to GET: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", resp.StatusCode)
	}

	// Read stream response
	reader := bufio.NewReader(resp.Body)
	line, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("failed to read line: %v", err)
	}

	if !strings.Contains(line, "message") && !strings.Contains(line, "data") {
		t.Logf("Line read: %q", line)
	}
}
