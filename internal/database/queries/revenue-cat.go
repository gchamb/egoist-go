package queries

import "egoist/internal/structs"

func (q *Queries) CreateSubscriber(subscriber structs.RevenueCatSubscriber) (error){
	query := `INSERT INTO revenue_cat_subscriber (id, transaction_id, product_id, purchased_at_ms, expiration_at_ms, user_id)
        VALUES (:id, :transaction_id, :product_id, :purchased_at_ms, :expiration_at_ms, :user_id)`

	_, err := q.DB.NamedExec(query, subscriber)
	return err
}

func (q *Queries) UpdateSubscriber(expiration int64, uid string) (error){
	query := `UPDATE revenue_cat_subscriber SET expiration_at_ms = ? where user_id = ?"`

	_, err := q.DB.Exec(query, expiration, uid)

	return err
}
