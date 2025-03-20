package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/AgustinCardozo/labo-so/crud/models"
	"github.com/AgustinCardozo/labo-so/utils/database"
	"time"
)

// IUserRepository define la interfaz para manejar usuarios.
type IUserRepository interface {
	Create(user models.User) error
	Delete(id int) error
	GetAll() ([]models.User, error)
	GetById(id int) (*models.User, error)
	Update(id int, user models.User) error
}

// UserRepository implementa IUserRepository.
type UserRepository struct{}

func (ur *UserRepository) Create(user models.User) error {
	db, err := database.GetSQLConnection()

	if err != nil {
		return err
	}

	query := "INSERT INTO GD2015C1.dbo.Usuarios (usuario, password, mail, nombre, createdAt) VALUES (@p1, @p2, @p3, @p4, @p5)"
	date := time.Now()
	_, err = db.ExecContext(context.Background(), query, user.User, "password", user.Mail, user.Name, date)

	return err
}

func (ur *UserRepository) Delete(id int) error {
	db, err := database.GetSQLConnection()

	if err != nil {
		return err
	}

	query := "DELETE FROM GD2015C1.dbo.Usuarios WHERE id = @id"
	_, err = db.ExecContext(context.Background(), query, sql.Named("id", id))

	return err
}

// GetAll obtiene todos los usuarios.
func (ur *UserRepository) GetAll() ([]models.User, error) {
	db, err := database.GetSQLConnection()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, usuario, mail, nombre FROM GD2015C1.dbo.Usuarios")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanUsers(rows)
}

// GetById obtiene un usuario por ID.
func (ur *UserRepository) GetById(id int) (*models.User, error) {
	db, err := database.GetSQLConnection()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	row := db.QueryRow("SELECT id, usuario, mail, nombre FROM dbo.Usuarios WHERE id = @id", sql.Named("id", id))
	return scanUser(row)
}

func (ur *UserRepository) Update(id int, user models.User) error {
	db, err := database.GetSQLConnection()

	if err != nil {
		return err
	}

	query := "UPDATE GD2015C1.dbo.Usuarios SET usuario = @usuario, mail = @mail WHERE id = @id"
	_, err = db.ExecContext(context.Background(), query, sql.Named("usuario", user.User), sql.Named("mail", user.Mail), sql.Named("id", id))

	return err
}

// scanUsers mapea múltiples filas a una lista de usuarios.
func scanUsers(rows *sql.Rows) ([]models.User, error) {
	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.User, &user.Mail, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// scanUser mapea una fila a un usuario.
func scanUser(row *sql.Row) (*models.User, error) {
	var user models.User
	err := row.Scan(&user.Id, &user.User, &user.Mail, &user.Name)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("usuario no encontrado")
		}
		return nil, err
	}

	return &user, nil
}
