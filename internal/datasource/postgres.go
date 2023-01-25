package datasource

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v4/pgxpool"
)

type postgres struct {
	pool *pgxpool.Pool
}

func NewPostgres(pool *pgxpool.Pool) *postgres {
	p := postgres{pool}
	return &p
}

func (p *postgres) GetAllUsers() ([]User, error) {
	var users []User
	rows, err := p.pool.Query(context.Background(), `select u.id, u.name, u.email from "Users" as u`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (p *postgres) GetUserByID(id int) (*User, error) {
	var user User
	err := p.pool.QueryRow(context.Background(),
		`select u.id, u.name, u.email
		from "Users" as u
		where u.id = $1;`, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (p *postgres) GetPastebinByID(id int) (*Pastebin, error) {
	var pastebin Pastebin
	err := p.pool.QueryRow(context.Background(),
		`select u.id, u.content, u.user_id
		from "Pastebins" as u
		where u.id = $1;`, id).Scan(&pastebin.ID, &pastebin.Content, &pastebin.UserID)
	if err != nil {
		return nil, err
	}
	return &pastebin, nil
}

func (p *postgres) GetAllPastebins() ([]Pastebin, error) {
	var pastebins []Pastebin

	rows, err := p.pool.Query(context.Background(),
		`select c.id, c.content, c.user_id
		  	from "Pastebins" as c;`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {

		var pastebin Pastebin

		err = rows.Scan(&pastebin.ID, &pastebin.Content, &pastebin.UserID)
		if err != nil {
			return nil, err
		}

		pastebins = append(pastebins, pastebin)
	}

	return pastebins, nil
}

func (p *postgres) GetPastebinsForUser(userID int) ([]Pastebin, error) {
	var pastebins []Pastebin

	rows, err := p.pool.Query(context.Background(),
		`select c.id, c.content, c.user_id
		  	from  "Users" as u 
			left join "Pastebins" as c on c.user_id = u."id" where c.user_id = $1;`, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var pastebin Pastebin

		err = rows.Scan(&pastebin.ID, &pastebin.Content, &pastebin.UserID)
		if err != nil {
			return nil, err
		}

		pastebins = append(pastebins, pastebin)
	}

	return pastebins, nil
}

func (p *postgres) CreateNewUser(user *User) (int, error) {
	tx, err := p.pool.Begin(context.Background())
	if err != nil {
		return -1, err
	}
	defer tx.Rollback(context.Background())
	var id int
	err = tx.QueryRow(context.Background(), `INSERT into "Users" (name, email) VALUES ($1, $2) returning id`, user.Name, user.Email).Scan(&id)

	if err != nil {
		return -1, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return -1, err
	}
	return id, nil
}

func (p *postgres) AddUserPastebin(id int, interests []string) (int, error) {
	tx, err := p.pool.Begin(context.Background())
	if err != nil {
		return -1, err
	}

	defer tx.Rollback(context.Background())

	var ds *goqu.InsertDataset
	rows := make([]interface{}, len(interests))
	for i, interest := range interests {
		rows[i] = goqu.Record{"userId": id, "topic": interest}
	}
	ds = goqu.Insert("UserInterest").Rows(rows)

	sql, _, err := ds.ToSQL()
	if err != nil {
		return -1, err
	}
	_, err = tx.Exec(context.Background(), sql)
	if err != nil {
		return -1, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return -1, err
	}
	return len(interests), nil
}
