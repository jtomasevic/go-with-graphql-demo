package seas

import "github.com/google/uuid"

type Crew struct {
	ID      uuid.UUID
	Name    string
	Pirates []Pirate
}

type Pirate struct {
	ID     uuid.UUID
	Name   string
	CrewId uuid.UUID
}

type Ship struct {
	ID   uuid.UUID
	Name string
	Crew *Crew
}
