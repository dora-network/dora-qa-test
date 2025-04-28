package consumer

import (
	"context"
	"dora-dev-test/data"
	"dora-dev-test/datastore"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Consumer interface {
	Start(ctx context.Context)
	Save(ctx context.Context, tick data.Tick) error
	Stop()
}

type consumer struct {
	client *kgo.Client
	ds     datastore.DataStore
	cancel context.CancelFunc
}

func (c *consumer) Save(ctx context.Context, tick data.Tick) error {
	return c.ds.SaveTick(ctx, tick)
}

func (c *consumer) Start(parent context.Context) {
	ctx, cancel := context.WithCancel(parent)
	c.cancel = cancel

	go c.start(ctx)
}

func (c *consumer) start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			pollCtx, cancel := context.WithTimeout(ctx, time.Second)
			fetches := c.client.PollRecords(pollCtx, 0)
			cancel()
			fetches.EachError(func(msg string, code int32, err error) {
				if !errors.Is(err, context.Canceled) && !errors.Is(err, context.DeadlineExceeded) {
					log.Printf("Error while fetching records: %s, code: %d, err: %v", msg, code, err)
				}
			})
			fetches.EachRecord(func(record *kgo.Record) {
				log.Printf("Received record: %s", string(record.Value))
				var tick data.Tick
				err := json.Unmarshal(record.Value, &tick)
				if err != nil {
					log.Printf("Error unmarshalling record: %v", err)
				}
				if err := c.Save(ctx, tick); err != nil {
					log.Printf("Error saving tick: %v", err)
				} else {
					log.Printf("Saved tick: %v", tick)
				}
			})
		}
	}
}

func (c *consumer) Stop() {
	if c.cancel != nil {
		c.cancel()
	}
}

func NewConsumer(client *kgo.Client, ds datastore.DataStore) Consumer {
	return &consumer{client: client, ds: ds}
}
