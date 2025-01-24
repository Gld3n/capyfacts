package models

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"

	_ "github.com/jackc/pgx/v5/pgxpool"
)

type Category string

const (
	Diet     Category = "diet"
	Habitat  Category = "habitat"
	Behavior Category = "behavior"
)

type Fact struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Category  Category  `json:"category"`
	CreatedAt time.Time `json:"created_at" db:"created"`
}

type FactsModelInterface interface {
	Random() (Fact, error)
	Create(*Fact) error
	GetAll(category Category, limit, offset int) ([]Fact, error)
	Edit(*Fact) error
	Delete(id int) error
}

type FactsModel struct {
	DB *pgxpool.Pool
}

func (f FactsModel) Random() (Fact, error) {
	var fact Fact

	stmt := `SELECT id, title, content, category, created FROM facts ORDER BY RANDOM() LIMIT 1`

	err := f.DB.QueryRow(context.Background(), stmt).Scan(&fact.ID, &fact.Title, &fact.Content, &fact.Category, &fact.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Fact{}, errors.New("there's not a single fact yet")
		}
		return Fact{}, err
	}

	return fact, nil
}

func (f FactsModel) Create(fact *Fact) error {
	stmt := `INSERT INTO facts (title, content, category, created) VALUES (@title, @content, @category, @created)`

	tx, err := f.DB.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("error starting transaction: %s", err.Error())
	}
	defer tx.Rollback(context.Background())

	args := pgx.NamedArgs{
		"title":    fact.Title,
		"content":  fact.Content,
		"category": fact.Category,
		"created":  fact.CreatedAt,
	}

	fmt.Printf("==> Named args: %+v", args)

	_, err = tx.Exec(context.Background(), stmt, args)
	if err != nil {
		return fmt.Errorf("error inserting new fact: %s", err.Error())
	}

	if err = tx.Commit(context.Background()); err != nil {
		return fmt.Errorf("error commiting fact creation: %s", err.Error())
	}

	return nil
}

func (f FactsModel) GetAll(category Category, limit, offset int) ([]Fact, error) {
	// TODO: filter by category
	stmt := `SELECT id, title, content, category, created FROM facts ORDER BY created LIMIT $1 OFFSET $2`

	rows, err := f.DB.Query(context.Background(), stmt, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("query error: %s", err.Error())
	}

	facts, err := pgx.CollectRows(rows, pgx.RowToStructByName[Fact])
	if err != nil {
		return []Fact{}, fmt.Errorf("CollectRows error: %s", err.Error())
	}

	return facts, nil
}

func (f FactsModel) Edit(fact *Fact) error {
	//TODO implement me
	panic("implement me")
}

func (f FactsModel) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
