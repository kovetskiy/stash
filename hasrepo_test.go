package stash

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHasRepository(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := *r.URL
		if url.Path != "/rest/api/1.0/repos" {
			t.Fatalf("GetRepositories() URL path expected to be /rest/api/1.0/repos but found %s\n", url.Path)
		}
		if r.Header.Get("Authorization") != "Basic dTpw" {
			t.Fatalf("Want  Basic dTpw but found %s\n", r.Header.Get("Authorization"))
		}
		fmt.Fprintln(w, repos)
	}))
	defer testServer.Close()

	url, _ := url.Parse(testServer.URL)
	stashClient := NewClient("u", "p", url)
	repositories, err := stashClient.GetRepositories()
	if err != nil {
		t.Fatalf("GetRepositories() not expecting an error, but received: %v\n", err)
	}

	if _, ok := HasRepository(repositories, "ssh://git@example.com:9999/teami/rabbit.git"); !ok {
		t.Fatalf("HasRepositories() expecting to contain %s, but did not\n", "ssh://git@example.com:9999/teami/rabbit.git")
	}
}
