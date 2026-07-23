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
