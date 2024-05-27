package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerHello(t *testing.T) {
	testcases := []struct {
		Name   string
		Expect string
	}{
		{Name: "Eduardo", Expect: "Hello Eduardo"},
		{Name: "Jocko", Expect: "Hello Jocko"},
		{Name: "", Expect: "Hello World"},
	}

	server := httptest.NewServer(http.HandlerFunc(HandlerHello))
	defer server.Close()

	for _, tc := range testcases {
		path := fmt.Sprintf("%s/?name=%s", server.URL, tc.Name)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", resp.StatusCode, http.StatusOK)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		var resBody map[string]string
		err = json.Unmarshal(body, &resBody)
		if err != nil {
			t.Fatal(err)
		}

		if resBody["message"] != tc.Expect {
			t.Fatalf("expected %s, got %s", tc.Expect, resBody["message"])
		}
	}
}
