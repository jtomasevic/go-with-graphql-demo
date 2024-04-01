package model

import "github.com/google/uuid"

// This id RDBS model.

type Pirate struct {
	Id     uuid.UUID
	Name   string
	CrewId *uuid.UUID
}

type Crew struct {
	Id     uuid.UUID
	Name   string
}

type Ship struct {
	Id     uuid.UUID
	Name   string
	CrewId *uuid.UUID
}