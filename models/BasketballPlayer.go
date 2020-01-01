package models

import (
	basketBallPlayer "basic-gRPC-proto"
)

type BasePersistenceModel struct {
	ID	string `gorm:"type:string;primary_key"`
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
		Id:                   em.ID,
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
		em.ID = gRPCModel.Id
		em.FirstName = gRPCModel.FirstName
		em.LastName = gRPCModel.LastName
		em.Age = gRPCModel.Age
		em.PhotoUrl = gRPCModel.PhotoUrl
		em.PointsPerGame = gRPCModel.PointsPerGame
		em.AssistsPerGame = gRPCModel.AssistsPerGame
		em.ReboundsPerGame = gRPCModel.ReboundsPerGame
}