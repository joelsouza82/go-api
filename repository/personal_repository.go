package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type PersonalRepository struct {
	connection *sql.DB
}

func NewPersonalRepository(connection *sql.DB) PersonalRepository {
	return PersonalRepository{
		connection: connection,
	}
}

func (pr *PersonalRepository) GetPersonals() ([]model.Personal, error) {
	query := "SELECT id, address, neighborhood, state, city, cep, phone, email, website, linkedin, github FROM personal"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Personal{}, err
	}

	var personalList []model.Personal
	var personalObj model.Personal

	for rows.Next() {
		err = rows.Scan(
			&personalObj.ID,
			&personalObj.Address,
			&personalObj.Neighborhood,
			&personalObj.State,
			&personalObj.City,
			&personalObj.Cep,
			&personalObj.Phone,
			&personalObj.Email,
			&personalObj.Website,
			&personalObj.Linkedin,
			&personalObj.Github,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Personal{}, err
		}

		personalList = append(personalList, personalObj)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return []model.Personal{}, err
	}

	rows.Close()

	return personalList, nil
}

func (pr *PersonalRepository) CreatePersonal(personal model.Personal) (int, error) {
	var id int
	query, err := pr.connection.Prepare("INSERT INTO personal " +
		"(address, city, neighborhood, state, cep, phone, email, website, linkedin, github) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(
		personal.Address,
		personal.City,
		personal.Neighborhood,
		personal.State,
		personal.Cep,
		personal.Phone,
		personal.Email,
		personal.Website,
		personal.Linkedin,
		personal.Github,
	).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}
