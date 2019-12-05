package main

import (
	basketBallPlayer "basic-gRPC-proto"
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/patrickmn/go-cache"
	"google.golang.org/grpc"
)

type server struct{}

var c *cache.Cache

func (*server) GetBasketballPlayer(ctx context.Context, r *basketBallPlayer.PlayerRequest) (*basketBallPlayer.PlayerResponse, error) {
	id := r.GetId()

	playerVal, found := c.Get(id)
	if found {
		player := playerVal.(basketBallPlayer.Player)
		return &basketBallPlayer.PlayerResponse{
			Result: &player,
		}, nil
	}

	return nil, fmt.Errorf("could not find player with id: %v", id)
}

func main() {
	fmt.Println("Starting gRPC micro-service...")
	c = cache.New(60*time.Minute, 70*time.Minute)
	c.Set("1", basketBallPlayer.Player{
		Id:              "1",
		FirstName:       "James",
		LastName:        "LeBron",
		Age:             33,
		PhotoUrl:        "https://ak-static.cms.nba.com/wp-content/uploads/headshots/nba/1610612747/2019/260x190/2544.png",
		PointsPerGame:   26,
		AssistsPerGame:  11,
		ReboundsPerGame: 8,
	}, cache.DefaultExpiration)

	l, e := net.Listen("tcp", ":50051")
	if e != nil {
		log.Fatalf("Failed to start listener %v", e)
	}

	s := grpc.NewServer()
	basketBallPlayer.RegisterPlayerServiceServer(s, &server{})

	if e := s.Serve(l); e != nil {
		log.Fatalf("failed to serve %v", e)
	}
}
