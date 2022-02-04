package repository

import (
	"api/auth"
	"api/models"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) SaveUser(user models.User) (uint64, error) {
	statement, erro := r.db.Prepare("INSERT INTO users(name, email, password, admin) VALUES($1,$2,$3,$4)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()
	hashedPassword, erro := auth.Hash(user.Password)
	if erro != nil {
		return 0, erro
	}
	user.Password = string(hashedPassword)
	resultado, erro := statement.Exec(user.Name, user.Email, user.Password, false)
	if erro != nil {
		return 0, erro
	}
	lastInsertedId, erro := resultado.RowsAffected()
	if erro != nil {
		return 0, erro
	}
	return uint64(lastInsertedId), nil
}

func (r UserRepository) GetAllUsers() ([]models.User, error) {
	rows, err := r.db.Query("select id, name, email, admin from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Admin); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
func (r UserRepository) GetUserById(id uint64) ([]models.User, error) {
	rows, err := r.db.Query("select id, name, email, admin from users where id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	if rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Admin); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r UserRepository) GetUserByEmail(email string) (models.User, error) {
	rows, err := r.db.Query("select id, name, email, password from users where email = $1", email)
	if err != nil {
		return models.User{}, err
	}
	var user models.User
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}
