package queries

import (
	"egoist/internal/structs"
)



func (q *Queries) GetUsers() ([]structs.User, error) {
	var users []structs.User
	err := q.DB.Select(&users, "SELECT * FROM user")
	return users, err
}

func (q *Queries) GetUserByID(id string) (structs.User, error) {
	var user structs.User
	err := q.DB.Select(&user, "SELECT * FROM user WHERE id = ?", id)
	return user, err
}


func (q *Queries) InsertUser(user structs.User) (error) {
	_, err := q.DB.NamedExec(`INSERT INTO user (id, email, password, goal_weight)
        VALUES (:id, :email, :password, :goal_weight)`, user)
		
	return err
}