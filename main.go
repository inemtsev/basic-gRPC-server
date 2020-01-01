package main

import (
	basketBallPlayer "basic-gRPC-proto"
	"basic-gRPC-server/models"
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"net"
	"time"

	"github.com/patrickmn/go-cache"
	"google.golang.org/grpc"
)

type server struct{}

var c *cache.Cache
var db *gorm.DB
const dbPath = "host=34.84.26.219 user=postgres password=JPHwcKkIGldw14wm"

func init() {
	var e error
	db, e = gorm.Open("postgres", dbPath)
	defer db.Close()

	if e != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&models.BasketballPlayer{})
}

func (*server) GetBasketballPlayer(ctx context.Context, r *basketBallPlayer.PlayerRequest) (*basketBallPlayer.PlayerResponse, error) {
	id := r.GetId()

	playerVal, found := c.Get(id)
	if found {
		fmt.Println("Checking cache...")
		player := playerVal.(*basketBallPlayer.Player)
		return &basketBallPlayer.PlayerResponse{
			Result: player,
		}, nil
	} else {
		fmt.Println("Checking database...")
		db, e := gorm.Open("postgres", dbPath)
		defer db.Close()
		if e != nil {
			panic(fmt.Sprintf("failed to connect to database: %v", e))
		}

		var p models.BasketballPlayer
		if e = db.First(&p, id).Error; e != nil {
			return nil, fmt.Errorf("could not find player with id: %v", id)
		} else {
			gRPCResult := p.GetgRPCModel()
			return &basketBallPlayer.PlayerResponse{
				Result:	&gRPCResult,
			}, nil
		}
	}
}

func main() {
	fmt.Println("Starting gRPC micro-service...")
	c = cache.New(60*time.Minute, 70*time.Minute)
	c.Set("1", &basketBallPlayer.Player{
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

	fmt.Println("Started micro-service successfully!")
}
