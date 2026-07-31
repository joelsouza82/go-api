package repository

import (
	"database/sql"
	"errors"
	"go-api/model"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetPersonalByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewPersonalRepository(db)

	personalMock := model.Personal{
		ID:      1,
		Address: "Test Address",
		Email:   "test@email.com",
	}

	expectedSQL := regexp.QuoteMeta("SELECT id, address, neighborhood, state, city, cep, phone, email, website, linkedin, github FROM personal WHERE id=$1")

	t.Run("Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "address", "neighborhood", "state", "city", "cep", "phone", "email", "website", "linkedin", "github"}).
			AddRow(personalMock.ID, personalMock.Address, personalMock.Neighborhood, personalMock.State, personalMock.City, personalMock.Cep, personalMock.Phone, personalMock.Email, personalMock.Website, personalMock.Linkedin, personalMock.Github)

		mock.ExpectQuery(expectedSQL).WithArgs(personalMock.ID).WillReturnRows(rows)

		personal, err := repo.GetPersonalByID(personalMock.ID)

		assert.NoError(t, err)
		assert.Equal(t, personalMock, personal)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Not Found", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).WithArgs(personalMock.ID).WillReturnError(sql.ErrNoRows)

		_, err := repo.GetPersonalByID(personalMock.ID)

		assert.Error(t, err)
		assert.Equal(t, "personal not found", err.Error())
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Query Error", func(t *testing.T) {
		dbError := errors.New("database error")
		mock.ExpectQuery(expectedSQL).WithArgs(personalMock.ID).WillReturnError(dbError)

		_, err := repo.GetPersonalByID(personalMock.ID)

		assert.Error(t, err)
		assert.Equal(t, dbError, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUpdatePersonal(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewPersonalRepository(db)

	personalToUpdate := model.Personal{
		ID:           1,
		Address:      "New Address",
		City:         "New City",
		Neighborhood: "New Neighborhood",
		State:        "NS",
		Cep:          "00000-000",
		Phone:        "999999999",
		Email:        "new@email.com",
		Website:      "new.com",
		Linkedin:     "new-linkedin",
		Github:       "new-github",
	}

	// Define a query esperada (usando expressão regular para flexibilidade)
	expectedSQL := regexp.QuoteMeta("UPDATE personal SET " +
		"address=$1, city=$2, neighborhood=$3, state=$4, cep=$5, phone=$6, email=$7, website=$8, linkedin=$9, github=$10 " +
		"WHERE id=$11")

	t.Run("Success", func(t *testing.T) {
		// Configura o mock para esperar a query de preparação
		mock.ExpectPrepare(expectedSQL).
			ExpectExec().
			WithArgs(
				personalToUpdate.Address, personalToUpdate.City,
				personalToUpdate.Neighborhood, personalToUpdate.State,
				personalToUpdate.Cep, personalToUpdate.Phone, personalToUpdate.Email,
				personalToUpdate.Website, personalToUpdate.Linkedin,
				personalToUpdate.Github, personalToUpdate.ID,
			).
			WillReturnResult(sqlmock.NewResult(1, 1)) // 1 linha afetada

		updatedPersonal, err := repo.UpdatePersonal(personalToUpdate)

		assert.NoError(t, err)
		assert.Equal(t, personalToUpdate, updatedPersonal)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Not Found", func(t *testing.T) {
		mock.ExpectPrepare(expectedSQL).
			ExpectExec().
			WithArgs(
				personalToUpdate.Address, personalToUpdate.City,
				personalToUpdate.Neighborhood, personalToUpdate.State,
				personalToUpdate.Cep, personalToUpdate.Phone, personalToUpdate.Email,
				personalToUpdate.Website, personalToUpdate.Linkedin,
				personalToUpdate.Github, personalToUpdate.ID,
			).
			WillReturnResult(sqlmock.NewResult(0, 0)) // 0 linhas afetadas

		_, err := repo.UpdatePersonal(personalToUpdate)

		assert.Error(t, err)
		assert.Equal(t, "personal not found", err.Error())
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Exec Error", func(t *testing.T) {
		dbError := errors.New("database error")

		mock.ExpectPrepare(expectedSQL).ExpectExec().WillReturnError(dbError)

		_, err := repo.UpdatePersonal(personalToUpdate)

		assert.Error(t, err)
		assert.Equal(t, dbError, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestDeletePersonal(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewPersonalRepository(db)
	personalID := 1

	expectedSQL := regexp.QuoteMeta("DELETE FROM personal WHERE id=$1")

	t.Run("Success", func(t *testing.T) {
		mock.ExpectPrepare(expectedSQL).
			ExpectExec().
			WithArgs(personalID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.DeletePersonal(personalID)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Not Found", func(t *testing.T) {
		mock.ExpectPrepare(expectedSQL).
			ExpectExec().
			WithArgs(personalID).
			WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.DeletePersonal(personalID)

		assert.Error(t, err)
		assert.Equal(t, "personal not found", err.Error())
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Exec Error", func(t *testing.T) {
		mock.ExpectPrepare(expectedSQL).ExpectExec().WillReturnError(errors.New("db error"))
		err := repo.DeletePersonal(personalID)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
