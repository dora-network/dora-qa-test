package redis

import (
	"context"
	"dora-dev-test/data"
	"encoding/json"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type DataStore struct {
	rdb *redis.Client
}

func (d DataStore) SaveTick(ctx context.Context, tick data.Tick) error {
	score := tick.Timestamp.UTC().UnixNano()
	log.Printf("Saving tick: %v, score: %d", tick, score)
	cmd := d.rdb.ZAddNX(ctx, tick.AssetID, redis.Z{Score: float64(score), Member: tick})
	return cmd.Err()
}

func (d DataStore) GetTicks(ctx context.Context, assetID string, from, to, offset, limit *int64) ([]data.Tick, error) {
	log.Printf("Getting ticks for assetID: %s, from: %d, to: %d, limit: %d", assetID, *from, *to, limit)
	rangeBy := &redis.ZRangeBy{
		Min:    fmt.Sprintf("%d", *from),
		Max:    fmt.Sprintf("%d", *to),
		Offset: 0,
		Count:  0,
	}

	if limit != nil {
		rangeBy.Offset = *offset
		rangeBy.Count = *limit
	}

	cmd := d.rdb.ZRangeByScore(ctx, assetID, rangeBy)

	if cmd.Err() != nil {
		return nil, cmd.Err()
	}

	var ticks []data.Tick

	for _, res := range cmd.Val() {
		tick := data.Tick{}
		if err := json.Unmarshal([]byte(res), &tick); err != nil {
			continue
		}
		ticks = append(ticks, tick)
	}

	return ticks, nil
}

func NewDataStore(rdb *redis.Client) *DataStore {
	return &DataStore{
		rdb: rdb,
	}
}
