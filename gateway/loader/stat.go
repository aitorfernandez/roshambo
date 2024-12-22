package loader

import (
	"context"
	"fmt"
	"sync"

	"github.com/aitorfernandez/roshambo/gateway/client"
	pb "github.com/aitorfernandez/roshambo/proto"
	"github.com/graph-gophers/dataloader"
)

// LoadStatsByAccount loads stats via data loader for the given accountID.
func LoadStatsByAccount(ctx context.Context, accountID string) ([]*pb.Stat, error) {
	data, err := load(ctx, statsByAccountKey, accountID)
	if err != nil {
		return nil, err
	}
	ss, ok := data.([]*pb.Stat)
	if !ok {
		return nil, fmt.Errorf("gateway/loader.LoadStatsByAccount wrong type v.([]) %T", data)
	}
	return ss, nil
}

type statsByAccountLoader struct {
	client *client.Client
}

func newStatsByAccountLoader(c *client.Client) dataloader.BatchFunc {
	return statsByAccountLoader{c}.loadBatch
}

func (l statsByAccountLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(keys)
		results = make([]*dataloader.Result, n)
		wg      sync.WaitGroup
	)
	wg.Add(n)
	for i, k := range keys {
		go func(i int, k dataloader.Key) {
			defer wg.Done()
			data, err := l.client.ListStatsByAccount(ctx, &pb.ListStatsByAccountReq{
				AccountID: k.String(),
			})
			// results must be in the same order as keys.
			results[i] = &dataloader.Result{
				Data:  data,
				Error: err,
			}
		}(i, k)
	}
	wg.Wait()
	return results
}
