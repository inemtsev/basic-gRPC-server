package models

import (
	basketBallPlayer "basic-gRPC-proto"
	"strconv"
)

type BasePersistenceModel struct {
	ID	uint64 `gorm:"type:int;primary_key"`
}

type BasketballPlayer struct {
	BasePersistenceModel //gorm.Model can be used if you don't mind ID being an integer
	FirstName string
	LastName string
	Age int32
	PhotoUrl string
	PointsPerGame int32
	AssistsPerGame int32
	ReboundsPerGame int32
}

func (em *BasketballPlayer) GetgRPCModel() basketBallPlayer.Player {
	return basketBallPlayer.Player{
		Id:                   strconv.FormatUint(em.ID, 10),
		FirstName:            em.FirstName,
		LastName:             em.LastName,
		Age:                  em.Age,
		PhotoUrl:             em.PhotoUrl,
		PointsPerGame:        em.PointsPerGame,
		AssistsPerGame:       em.AssistsPerGame,
		ReboundsPerGame:      em.ReboundsPerGame,
	}
}

func (em *BasketballPlayer) From(gRPCModel basketBallPlayer.Player) {
		u, e := strconv.ParseUint(gRPCModel.Id, 10, 64)
		if e != nil {
			panic("incorrect ID from gRPC")
		}

		em.ID = u
		em.FirstName = gRPCModel.FirstName
		em.LastName = gRPCModel.LastName
		em.Age = gRPCModel.Age
		em.PhotoUrl = gRPCModel.PhotoUrl
		em.PointsPerGame = gRPCModel.PointsPerGame
		em.AssistsPerGame = gRPCModel.AssistsPerGame
		em.ReboundsPerGame = gRPCModel.ReboundsPerGame
}