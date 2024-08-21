package queries

import (
	"egoist/internal/structs"

	"github.com/google/uuid"
)



func (q *Queries) GetUsers() ([]structs.User, error) {
	var users []structs.User
	err := q.DB.Select(&users, "SELECT * FROM user")
	return users, err
}

func (q *Queries) GetUserByID(id string) (structs.User, error) {
	var user structs.User
	err := q.DB.Get(&user, "SELECT * FROM user WHERE id = ?", id)
	return user, err
}

func (q *Queries) GetUserByEmail(email string) (structs.User, error){
	var user structs.User
	err := q.DB.Get(&user, "SELECT * FROM user WHERE email = ?", email)
	return user, err
}


func (q *Queries) InsertUser(user structs.User) (error) {
	_, err := q.DB.NamedExec(`INSERT INTO user (id, email, password, goal_weight)
        VALUES (:id, :email, :password, :goal_weight)`, user)
		
	return err
}

func (q *Queries) CreateUser(email string, password *string) (string, error){
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	err = q.InsertUser(structs.User{
		ID:    id.String(),
		Email: email,
		Password: password,
	})

	if err != nil {
		return "", nil
	}

	return id.String(), err
}