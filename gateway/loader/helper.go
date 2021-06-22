package loader

import (
	"context"
	"fmt"

	"github.com/aitorfernandez/roshambo/pkg/ctxkey"
	"github.com/graph-gophers/dataloader"
)

// extract is a helper function to make common get-value, assert-type, return-error-or-value
// operations easier.
func extract(ctx context.Context, k ctxkey.Key) (*dataloader.Loader, error) {
	ldr, ok := ctx.Value(k).(*dataloader.Loader)
	if !ok {
		return nil, fmt.Errorf("unable to find %s loader on the request context", k)
	}
	return ldr, nil
}

// load resolves the given key.
func load(ctx context.Context, loaderKey ctxkey.Key, k string) (interface{}, error) {
	ldr, err := extract(ctx, loaderKey)
	if err != nil {
		return nil, err
	}
	// Load returns a thunk, is a function returned from a function that is a closure over a value, an interface value and error.
	v, err := ldr.Load(ctx, dataloader.StringKey(k))()
	if err != nil {
		return nil, err
	}
	return v, nil
}
