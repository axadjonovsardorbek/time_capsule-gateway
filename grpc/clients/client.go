package clients

import (
	"gateway/config"
	cp "gateway/genproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClients struct {
	Memory          cp.MemoriesServiceClient
	Media           cp.MediasServiceClient
	SharedMemory    cp.SharedMemoriesServiceClient
	Comment         cp.CommentsServiceClient
	Milestone       cp.MilestonesServiceClient
	CustomEvent     cp.CustomEventsServiceClient
	PersonalEvent   cp.PersonalEventsServiceClient
	HistoricalEvent cp.HistoricalEventsServiceClient
}

func NewGrpcClients(cfg *config.Config) (*GrpcClients, error) {
	connM, err := grpc.NewClient(cfg.MEMORY_HOST+cfg.MEMORY_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connT, err := grpc.NewClient(cfg.TIMELINE_HOST+cfg.TIMELINE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &GrpcClients{
		Memory:          cp.NewMemoriesServiceClient(connM),
		Media:           cp.NewMediasServiceClient(connM),
		SharedMemory:    cp.NewSharedMemoriesServiceClient(connM),
		Comment:         cp.NewCommentsServiceClient(connM),
		Milestone:       cp.NewMilestonesServiceClient(connT),
		CustomEvent:     cp.NewCustomEventsServiceClient(connT),
		PersonalEvent:   cp.NewPersonalEventsServiceClient(connT),
		HistoricalEvent: cp.NewHistoricalEventsServiceClient(connT),
	}, nil
}
