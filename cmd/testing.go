package cmd

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/google/shlex"
	"github.com/krakowitzerm/ghadmin/api"
	"github.com/krakowitzerm/ghadmin/context"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type QuestionStub struct {
	Name    string
	Value   interface{}
	Default bool
}

func initFakeHTTP() *api.FakeHTTP {
	http := &api.FakeHTTP{}
	apiClientForContext = func(context.Context) (*api.Client, error) {
		return api.NewClient(api.ReplaceTripper(http)), nil
	}
	return http
}

func eq(t *testing.T, got interface{}, expected interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}

type cmdOut struct {
	outBuf, errBuf *bytes.Buffer
}

func (c cmdOut) String() string {
	return c.outBuf.String()
}

func (c cmdOut) Stderr() string {
	return c.errBuf.String()
}

func RunCommand(cmd *cobra.Command, args string) (*cmdOut, error) {
	rootCmd := cmd.Root()
	argv, err := shlex.Split(args)
	if err != nil {
		return nil, err
	}
	rootCmd.SetArgs(argv)

	outBuf := bytes.Buffer{}
	cmd.SetOut(&outBuf)
	errBuf := bytes.Buffer{}
	cmd.SetErr(&errBuf)

	// Reset flag values so they don't leak between tests
	// FIXME: change how we initialize Cobra commands to render this hack unnecessary
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		switch v := f.Value.(type) {
		case pflag.SliceValue:
			_ = v.Replace([]string{})
		default:
			switch v.Type() {
			case "bool", "string", "int":
				_ = v.Set(f.DefValue)
			}
		}
	})

	_, err = rootCmd.ExecuteC()
	cmd.SetOut(nil)
	cmd.SetErr(nil)

	return &cmdOut{&outBuf, &errBuf}, err
}
