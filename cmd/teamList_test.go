package cmd

import (
	"os"
	"regexp"
	"testing"
)

func TestTeamList(t *testing.T) {
	http := initFakeHTTP()

	jsonFile, _ := os.Open("../test/fixtures/teamList.json")
	defer jsonFile.Close()
	http.StubResponse(200, jsonFile)

	output, err := RunCommand(teamListCmd, "team list")
	if err != nil {
		t.Errorf("error running command `issue list`: %v", err)
	}
	eq(t, output.String(), `foo1
foo2
`)

	expectedIssues := []*regexp.Regexp{
		regexp.MustCompile(`(?m)^foo1$`),
		regexp.MustCompile(`(?m)^foo2$`),
	}

	for _, r := range expectedIssues {
		if !r.MatchString(output.String()) {
			t.Errorf("output did not match regexp /%s/\n> output\n%s\n", r, output)
			return
		}
	}
}
