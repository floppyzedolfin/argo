package load

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/floppyzedolfin/argo/services/portdomain/client/portdomain"
)

// upserter is a utility that has been taught how to "upsert something on a backend"
type upserter struct {
	backendUpserter func(context.Context, *portdomain.UpsertRequest) error
}

func (u *upserter) upsertPorts(ctx context.Context, f io.Reader) error {
	decoder := json.NewDecoder(f)
	// remove delimiters
	_, err := decoder.Token()
	if err == io.EOF {
		// this file is empty
		return nil
	}
	if err != nil {
		return fmt.Errorf("error while reading the beginning of the file: %w", err)
	}

	for decoder.More() {
		if err = u.upsertPort(ctx, decoder); err != nil {
			return fmt.Errorf("unable to upsert port: %w", err)
		}
	}
	// we've finished parsing the entries of the file
	return nil
}

func (u *upserter) upsertPort(ctx context.Context, dec *json.Decoder) error {
	token, err := dec.Token()
	if err != nil {
		return fmt.Errorf("unable to get token from decoder: %w", err)
	}

	p := new(portdomain.Port)

	if err = dec.Decode(&p); err != nil {
		return fmt.Errorf("unable to decode port: %w", err)
	}
	p.Id = token.(string)

	if err = u.backendUpserter(ctx, &portdomain.UpsertRequest{Port: p}); err != nil {
		return fmt.Errorf("error while calling backend: %w", err)
	}

	return nil
}
