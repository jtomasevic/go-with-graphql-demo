package repo

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	"github.com/google/uuid"
	"github.com/jtomasevic/go-with-graphql-demo/src/services/7seas/data_store/model"
	_ "github.com/proullon/ramsql/driver"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("ramsql", "Pirates")
	if err != nil {
		log.Fatalf("sql.Open : Error : %s\n", err)
		panic(err)
	}

	return db
}

func InitDb() error {
	batch := []string{
		`CREATE TABLE pirate (id UUID PRIMARY KEY, name TEXT, crew_id UUID);`,
		`CREATE TABLE crew (id UUID PRIMARY KEY, name TEXT);`,
		`CREATE TABLE ship (id UUID PRIMARY KEY, name TEXT, crew_id UUID);`,
	}
	db := GetConnection()

	defer db.Close()

	for _, b := range batch {
		// fmt.Println(b)
		_, err := db.Exec(b)
		if err != nil {
			log.Fatalf("sql.Exec: Error: %s\n", err)
			return err
		}
	}
	return nil
}

func PopulateDb(hugeDataSet bool) error {
	var ships []model.Ship
	var crews []model.Crew
	var pirates []model.Pirate
	if hugeDataSet {
		ships, crews, pirates = createTestObjectsForStresTest()
	} else {
		ships, crews, pirates = createTestObjects()
	}
	batch := []string{}
	for _, ship := range ships {
		batch = append(batch, fmt.Sprintf("INSERT INTO ship (id, name, crew_id) VALUES ('%s', '%s', '%s');", ship.Id.String(), ship.Name, ship.CrewId.String()))
	}
	for _, crew := range crews {
		batch = append(batch, fmt.Sprintf("INSERT INTO crew (id, name) VALUES ('%s', '%s');", crew.Id.String(), crew.Name))
	}
	for _, pirate := range pirates {
		batch = append(batch, fmt.Sprintf("INSERT INTO pirate (id, name, crew_id) VALUES ('%s', '%s', '%s');", pirate.Id.String(), pirate.Name, pirate.CrewId.String()))
	}

	// }
	db := GetConnection()

	defer db.Close()

	for _, b := range batch {
		// fmt.Println(b)
		_, err := db.Exec(b)
		if err != nil {
			log.Fatalf("sql.Exec: Error: %s\n", err)
			return err
		}
	}
	return nil
}

var (
	cursedCrewId         = uuid.New()
	flyingDutchmanCrewId = uuid.New()
	blackPearlId         = uuid.New()
	flyingDutchmanId     = uuid.New()
)

func createTestObjects() ([]model.Ship, []model.Crew, []model.Pirate) {
	ships := []model.Ship{
		{
			Id:     blackPearlId,
			Name:   "Black Pearl",
			CrewId: &cursedCrewId,
		}, {
			Id:     flyingDutchmanId,
			Name:   "Flying Dutchman",
			CrewId: &flyingDutchmanCrewId,
		},
	}
	crews := []model.Crew{
		{
			Id:   cursedCrewId,
			Name: "Cursed Crew",
		}, {
			Id:   flyingDutchmanCrewId,
			Name: "Crew of the Flying Dutchman",
		},
	}
	pirates := []model.Pirate{
		{
			Id:     uuid.New(),
			Name:   "Jack Sparrow",
			CrewId: &cursedCrewId,
		},
		{
			Id:     uuid.New(),
			Name:   "Hector Barbossa",
			CrewId: &cursedCrewId,
		},
		{
			Id:     uuid.New(),
			Name:   "Pintel",
			CrewId: &cursedCrewId,
		},
		{
			Id:     uuid.New(),
			Name:   "Jack the Monkey",
			CrewId: &cursedCrewId,
		},
		{
			Id:     uuid.New(),
			Name:   "Koehler",
			CrewId: &cursedCrewId,
		},
		{
			Id:     uuid.New(),
			Name:   "Twigg",
			CrewId: &cursedCrewId,
		},
		{
			Id:     uuid.New(),
			Name:   "Jacoby",
			CrewId: &cursedCrewId,
		},
		{
			Id:     uuid.New(),
			Name:   "Mallot",
			CrewId: &cursedCrewId,
		},
		{
			Id:     uuid.New(),
			Name:   "Scurvy",
			CrewId: &cursedCrewId,
		},
		{
			Id:     uuid.New(),
			Name:   "Davy Jones",
			CrewId: &flyingDutchmanCrewId,
		},
		{
			Id:     uuid.New(),
			Name:   "Bootstrap Bill Turner",
			CrewId: &flyingDutchmanCrewId,
		},
		{
			Id:     uuid.New(),
			Name:   "Maccus",
			CrewId: &flyingDutchmanCrewId,
		},
		{
			Id:     uuid.New(),
			Name:   "Koleniko",
			CrewId: &flyingDutchmanCrewId,
		},
		{
			Id:     uuid.New(),
			Name:   "Palifico",
			CrewId: &flyingDutchmanCrewId,
		},
		{
			Id:     uuid.New(),
			Name:   "Jimmy Legs",
			CrewId: &flyingDutchmanCrewId,
		},
		{
			Id:     uuid.New(),
			Name:   "Clanker",
			CrewId: &flyingDutchmanCrewId,
		},
		{
			Id:     uuid.New(),
			Name:   "Ogilvey",
			CrewId: &flyingDutchmanCrewId,
		},
		{
			Id:     uuid.New(),
			Name:   "Hadras",
			CrewId: &flyingDutchmanCrewId,
		},
	}

	return ships, crews, pirates
}

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func createTestObjectsForStresTest() ([]model.Ship, []model.Crew, []model.Pirate) {
	ships := []model.Ship{}
	crews := []model.Crew{}
	pirates := []model.Pirate{}

	for i := 0; i < 200; i++ {
		crew := model.Crew{
			Id:   uuid.New(),
			Name: randString(12),
		}
		for i := 0; i < 200; i++ {
			pirate := model.Pirate{
				Id:     uuid.New(),
				CrewId: &crew.Id,
				Name:   randString(12),
			}
			pirates = append(pirates, pirate)
		}
		crews = append(crews, crew)
	}

	for i := 0; i < 200; i++ {
		ship := model.Ship{
			Id:     uuid.New(),
			Name:   randString(8),
			CrewId: &crews[i].Id,
		}
		ships = append(ships, ship)
	}
	return ships, crews, pirates
}
