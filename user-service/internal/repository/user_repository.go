package repository

import (
	"TurboTransit/user-service/internal/model"
	"database/sql"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *model.User) error {
    stmt, err := r.db.Prepare("INSERT INTO users (username, email, password_hash, salt, first_name, last_name) VALUES (?, ?, ?, ?, ?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(user.Username, user.Email, user.Password, user.FirstName, user.LastName)
    if err != nil {
        return err
    }
    return nil
}

func (r *UserRepository) GetUserByID(id int) (*model.User, error) {
    row := r.db.QueryRow("SELECT id, username, email, first_name, last_name, created_at, updated_at FROM users WHERE id = ?", id)
    user := &model.User{}
    err := row.Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return user, nil
}

func (r *UserRepository) UpdateUser(user *model.User) error {
    stmt, err := r.db.Prepare("UPDATE users SET username = ?, email = ?, first_name = ?, last_name = ?, updated_at = ? WHERE id = ?")
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(user.Username, user.Email, user.FirstName, user.LastName, user.UpdatedAt, user.ID)
    if err != nil {
        return err
    }
    return nil
}