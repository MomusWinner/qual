package integration

import (
	"app/internal/core"
	"app/tests/mocks"
	"fmt"
	"net/http"
	"testing"
)

func TestGetAllUsers(t *testing.T) {
	ctx := mocks.InitMockCtx(":4000")
	s := core.NewServer(ctx, false, false)

	go func() {
		s.Start()
	}()

	resp, _ := http.DefaultClient.Get("http://localhost:4000/healthcheck")
	fmt.Println(resp.Status)

	if resp.StatusCode != 200 {
		t.Error("Incorrect status code")
	}
}
