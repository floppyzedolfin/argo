package load

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/floppyzedolfin/argo/services/portdomain/client/portdomain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpserter_UpsertPorts(t *testing.T) {
	tt := map[string]struct {
		inputFile    string
		errMsg       string
		backendError string
		count        int
	}{
		"one port": {
			inputFile: "testdata/onePort.json",
			count:     1,
		},
		"two ports": {
			inputFile: "testdata/twoPorts.json",
			count:     2,
		},
		"no port": {
			inputFile: "testdata/noPort.json",
			count:     0,
		},
		"invalid input file": {
			inputFile: "testdata/invalid.json",
			errMsg:    "error while reading the beginning of the file",
		},
		"backend error": {
			inputFile:    "testdata/onePort.json",
			backendError: "a backend error",
			errMsg:       "error while calling backend",
		},
		"invalid port in input": {
			inputFile: "testdata/invalidPort.json",
			errMsg:    "unable to decode port",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			f, _ := os.Open(tc.inputFile)
			counter := 0
			u := upserter{backendUpserter: func(context.Context, *portdomain.UpsertRequest) error {
				if tc.backendError != "" {
					return fmt.Errorf(tc.backendError)
				}
				counter++
				return nil
			}}

			err := u.upsertPorts(context.Background(), f)

			if tc.errMsg != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.errMsg)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.count, counter)
			}
		})
	}
}
