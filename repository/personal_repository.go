package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"go-api/model"
)

type PersonalRepository struct {
	connection *sql.DB
}

func NewPersonalRepository(connection *sql.DB) PersonalRepositoryInterface {
	return &PersonalRepository{
		connection: connection,
	}
}

func (pr *PersonalRepository) GetPersonalByID(id int) (model.Personal, error) {
	query := "SELECT id, address, neighborhood, state, city, cep, phone, email, website, linkedin, github FROM personal WHERE id=$1"
	row := pr.connection.QueryRow(query, id)

	var personalObj model.Personal
	err := row.Scan(
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
		if errors.Is(err, sql.ErrNoRows) {
			return model.Personal{}, errors.New("personal not found")
		}
		fmt.Println(err)
		return model.Personal{}, err
	}

	return personalObj, nil
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

func (pr *PersonalRepository) UpdatePersonal(personal model.Personal) (model.Personal, error) {
	query := "UPDATE personal SET " +
		"address=$1, city=$2, neighborhood=$3, state=$4, cep=$5, phone=$6, email=$7, website=$8, linkedin=$9, github=$10 " +
		"WHERE id=$11"

	stmt, err := pr.connection.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return model.Personal{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
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
		personal.ID,
	)
	if err != nil {
		fmt.Println(err)
		return model.Personal{}, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return model.Personal{}, errors.New("personal not found")
	}

	return personal, nil
}

func (pr *PersonalRepository) DeletePersonal(id int) error {
	query := "DELETE FROM personal WHERE id=$1"

	stmt, err := pr.connection.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("personal not found")
	}

	return nil
}
