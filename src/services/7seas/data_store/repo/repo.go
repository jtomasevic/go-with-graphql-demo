package repo

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	datasource "github.com/jtomasevic/go-with-graphql-demo/src/services/7seas/data_store/data_source"
	"github.com/jtomasevic/go-with-graphql-demo/src/services/7seas/data_store/model"
)

type PiratesRepo struct {
	DataSource datasource.DataSource
}

func (r *PiratesRepo) GetPirates(ctx context.Context, crew_id *uuid.UUID) ([]model.Pirate, error) {

	query := "select id, name, crew_id from pirate"
	// for the purpose of demo, remove latter:
	fmt.Println(query)
	if crew_id != nil {
		query = fmt.Sprintf("%s where crew_id = '%s'", query, crew_id.String())
	}
	results := []model.Pirate{}
	rows, err := r.DataSource.Db().Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id uuid.UUID
		var name string
		var crewId uuid.UUID
		err := rows.Scan(&id, &name, &crewId)
		if err != nil {
			return nil, err
		}
		results = append(results, model.Pirate{
			Id:     id,
			Name:   name,
			CrewId: &crewId,
		})
	}

	return results, nil
}

func (r *PiratesRepo) GetPiratesByIds(ctx context.Context, ids []uuid.UUID) ([]model.Pirate, error) {
	var queryBuilder strings.Builder
	queryBuilder.WriteString("SELECT id, name, crew_id FROM pirate WHERE id IN (")
	for _, id := range ids {
		queryBuilder.WriteString(fmt.Sprintf("'%s',", id.String()))
	}
	query, _ := strings.CutSuffix(queryBuilder.String(), ",")
	query = fmt.Sprintf("%s)", query)
	// for the purpose of demo, remove latter:
	fmt.Println(query)
	results := []model.Pirate{}
	rows, err := r.DataSource.Db().Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id uuid.UUID
		var name string
		var crewId uuid.UUID
		err := rows.Scan(&id, &name, &crewId)
		if err != nil {
			return nil, err
		}
		results = append(results, model.Pirate{
			Id:     id,
			Name:   name,
			CrewId: &crewId,
		})
	}

	return results, nil
}

func (r *PiratesRepo) GetCrews(ctx context.Context) ([]model.Crew, error) {

	query := "select id, name from crew"
	// for the purpose of demo, remove latter:
	fmt.Println(query)

	results := []model.Crew{}
	rows, err := r.DataSource.Db().Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id uuid.UUID
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		results = append(results, model.Crew{
			Id:   id,
			Name: name,
		})
	}

	return results, nil
}

func (repo *PiratesRepo) GetCrew(ctx context.Context, id uuid.UUID) (model.Crew, error) {
	query := fmt.Sprintf("select id, name from crew where id = '%s", id.String())
	// for the purpose of demo, remove latter:
	fmt.Println(query)

	row := repo.DataSource.Db().QueryRow(query)
	var dbid uuid.UUID
	var name string
	err := row.Scan(&dbid, &name)
	if err != nil {
		return model.Crew{}, err
	}
	result := model.Crew{
		Id:   dbid,
		Name: name,
	}
	return result, nil
}

func (r *PiratesRepo) GetShips(ctx context.Context) ([]model.Ship, error) {

	query := "SELECT id, name, crew_id FROM ship"
	// fmt.Println(query)
	results := []model.Ship{}
	rows, err := r.DataSource.Db().Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id uuid.UUID
		var name string
		var crewId uuid.UUID
		err := rows.Scan(&id, &name, &crewId)
		if err != nil {
			return nil, err
		}
		results = append(results, model.Ship{
			Id:     id,
			Name:   name,
			CrewId: &crewId,
		})
	}

	return results, nil
}
func (repo *PiratesRepo) GetShip(ctx context.Context, id uuid.UUID) (model.Ship, error) {
	query := fmt.Sprintf("SELECT id, name, crew_id FROM ship where id = '%s", id.String())
	// for the purpose of demo, remove latter:
	fmt.Println(query)

	row := repo.DataSource.Db().QueryRow(query)
	var dbid uuid.UUID
	var name string
	var crewId uuid.UUID
	err := row.Scan(&dbid, &name, &crewId)
	if err != nil {
		return model.Ship{}, err
	}
	result := model.Ship{
		Id:     dbid,
		Name:   name,
		CrewId: &crewId,
	}
	return result, nil
}

func (r *PiratesRepo) GetCrewsByIds(ctx context.Context, ids []uuid.UUID) ([]model.Crew, error) {
	var queryBuilder strings.Builder
	queryBuilder.WriteString("SELECT id, name FROM crew WHERE id IN (")
	for _, id := range ids {
		queryBuilder.WriteString(fmt.Sprintf("'%s',", id.String()))
	}
	query, _ := strings.CutSuffix(queryBuilder.String(), ",")
	query = fmt.Sprintf("%s)", query)
	// for the purpose of demo, remove latter:
	fmt.Println(query)

	results := []model.Crew{}
	rows, err := r.DataSource.Db().Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id uuid.UUID
		var name string

		err := rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		results = append(results, model.Crew{
			Id:   id,
			Name: name,
		})
	}

	return results, nil
}

func (repo *PiratesRepo) GetPiratesInCrews(ctx context.Context, crew_ids []uuid.UUID) ([]model.Pirate, error) {
	var queryBuilder strings.Builder
	queryBuilder.WriteString("SELECT id, name, crew_id FROM pirate WHERE crew_id IN (")
	for _, id := range crew_ids {
		queryBuilder.WriteString(fmt.Sprintf("'%s',", id.String()))
	}
	query, _ := strings.CutSuffix(queryBuilder.String(), ",")
	query = fmt.Sprintf("%s)", query)
	// for the purpose of demo, remove latter:
	fmt.Println(query)

	// fmt.Println(query)
	results := []model.Pirate{}
	rows, err := repo.DataSource.Db().Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id uuid.UUID
		var name string
		var crewId uuid.UUID
		err := rows.Scan(&id, &name, &crewId)
		if err != nil {
			return nil, err
		}
		results = append(results, model.Pirate{
			Id:     id,
			Name:   name,
			CrewId: &crewId,
		})
	}

	return results, nil
}

func (repo *PiratesRepo) GetCrewForShip(ctx context.Context, shipId uuid.UUID) (model.Crew, error) {
	query := fmt.Sprintf("SELECT crew.id, crew.name FROM crew JOIN ship ON ship.crew_id = crew.id where ship.id = '%s'", shipId)
	// for the purpose of demo, remove latter:
	fmt.Println(query)

	results := []model.Crew{}
	rows, err := repo.DataSource.Db().Query(query)
	if err != nil {
		return model.Crew{}, err
	}

	for rows.Next() {
		var id uuid.UUID
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			return model.Crew{}, err
		}
		results = append(results, model.Crew{
			Id:   id,
			Name: name,
		})
	}
	if len(results) != 1 {
		return model.Crew{}, errors.New("Ship must have only one crew")
	}
	return results[0], nil
}
