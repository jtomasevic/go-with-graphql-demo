// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/google/uuid"
)

type Crew struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Pirates []Pirate  `json:"pirates"`
}

type Mutation struct {
}

type Pirate struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Query struct {
}

type Ship struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Crew *Crew     `json:"crew"`
}

type UpsertCrew struct {
	ID     *uuid.UUID `json:"id,omitempty"`
	Name   string     `json:"name"`
	ShipID *string    `json:"shipId,omitempty"`
}

type UpsertPirate struct {
	ID     *uuid.UUID `json:"id,omitempty"`
	Name   string     `json:"name"`
	CrewID *string    `json:"crewId,omitempty"`
}

type UpsertShip struct {
	ID   *uuid.UUID `json:"id,omitempty"`
	Name string     `json:"name"`
}
