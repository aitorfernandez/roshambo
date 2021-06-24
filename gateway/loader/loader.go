package loader

import (
	"context"

	"github.com/aitorfernandez/roshambo/gateway/client"
	"github.com/aitorfernandez/roshambo/pkg/ctxkey"
	"github.com/graph-gophers/dataloader"
)

var profileByAccountKey = ctxkey.New("profile_by_account")
var statsByAccountKey = ctxkey.New("stats_by_account")

// Map holds the keys to batch fucntions.
type Map map[ctxkey.Key]dataloader.BatchFunc

// New returns a lookup map of context keys to batch functions.
func New(c *client.Client) Map {
	return Map{
		profileByAccountKey: newProfileByAccountLoader(c),
		statsByAccountKey:   newStatsByAccountLoader(c),
	}
}

// Attach attaches dataloaders to the request's context.
func (m Map) Attach(ctx context.Context) context.Context {
	for k, batchFunc := range m {
		ldr := dataloader.NewBatchedLoader(batchFunc, dataloader.WithCache(dataloader.NewCache()))
		ctx = context.WithValue(ctx, k, ldr)
	}
	return ctx
}
