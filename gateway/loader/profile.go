package loader

import (
	"context"
	"fmt"
	"sync"

	"github.com/aitorfernandez/roshambo/gateway/client"
	pb "github.com/aitorfernandez/roshambo/proto"
	"github.com/graph-gophers/dataloader"
)

// LoadProfileByAccount loads profile via data loader for the given accountID.
func LoadProfileByAccount(ctx context.Context, accountID string) (*pb.Profile, error) {
	data, err := load(ctx, profileByAccountKey, accountID)
	if err != nil {
		return nil, err
	}
	p, ok := data.(*pb.Profile)
	if !ok {
		return nil, fmt.Errorf("gateway/loader.LoaderProfileByAccount wrong type v.([]) %T", data)
	}
	return p, nil
}

type profileByAccountLoader struct {
	client *client.Client
}

func newProfileByAccountLoader(c *client.Client) dataloader.BatchFunc {
	return profileByAccountLoader{c}.loadBatch
}

func (l profileByAccountLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(keys)
		results = make([]*dataloader.Result, n)
		wg      sync.WaitGroup
	)
	wg.Add(n)
	for i, k := range keys {
		go func(i int, k dataloader.Key) {
			defer wg.Done()
			data, err := l.client.GetProfileByAccount(ctx, &pb.GetProfileByAccountReq{
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
