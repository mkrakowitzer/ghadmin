package api

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
)

func TestTeamList(t *testing.T) {

	http := &FakeHTTP{}
	client := NewClient(ReplaceTripper(http))

	http.StubResponse(200, bytes.NewBufferString(`
	{
		"data": {
		  "organization": {
			"teams": {
			  "edges": []
			}
		  }
		}
	  }`))

	cmd := &cobra.Command{}

	_, err := TeamList(client, cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(http.Requests) != 1 {
		t.Fatalf("expected 1 HTTP requests, seen %d", len(http.Requests))
	}

}
