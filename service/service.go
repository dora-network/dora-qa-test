package service

import (
	"context"
	"log"

	"dora-dev-test/datastore"

	api "dora-dev-test/api/v1"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	api.UnimplementedDoraDevTestServiceServer
	ds datastore.DataStore
}

func (s Service) HealthCheck(ctx context.Context, empty *emptypb.Empty) (*api.HealthCheckResponse, error) {
	return &api.HealthCheckResponse{
		LastHeartbeat: timestamppb.Now(),
	}, nil
}

func (s Service) GetTicks(ctx context.Context, request *api.GetTicksRequest) (*api.GetTicksResponse, error) {
	var (
		from *int64
		to   *int64
	)

	if request.GetStart() != nil {
		startTime := request.GetStart().AsTime().UTC()
		start := startTime.UnixNano()
		from = &start
	}

	if request.GetEnd() != nil {
		endTime := request.GetEnd().AsTime().UTC()
		end := endTime.UnixNano()
		to = &end
	}

	log.Printf("Getting ticks for assetID: %s, from: %d, to: %d",
		request.Symbol, *from, *to)

	ticks, err := s.ds.GetTicks(
		ctx,
		request.Symbol,
		from,
		to,
		request.Offset,
		request.Limit,
	)
	if err != nil {
		log.Printf("Error getting ticks: %v", err)
		return nil, err
	}

	log.Printf("Received %d ticks", len(ticks))
	ticksProto := make([]*api.Tick, len(ticks))

	for i, tick := range ticks {
		ticksProto[i] = &api.Tick{
			AssetId:   tick.AssetID,
			Timestamp: timestamppb.New(tick.Timestamp),
			LastPrice: tick.LastPrice,
			LastSize:  tick.LastSize,
			BestBid:   tick.BestAsk,
			BestAsk:   tick.BestBid,
		}
	}

	return &api.GetTicksResponse{
		Ticks: ticksProto,
	}, nil
}

func NewService(ds datastore.DataStore) Service {
	return Service{
		ds: ds,
	}
}
