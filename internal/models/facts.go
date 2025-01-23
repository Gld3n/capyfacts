package models

import (
	"context"
	"errors"
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
	CreatedAt time.Time `json:"created_at"`
}

type FactsModelInterface interface {
	Random() (Fact, error)
	Create(Fact) error
	GetAll(category, limit, offset int) ([]Fact, error)
	Edit(Fact) error
	Delete(id int) error
}

type FactsModel struct {
	DB *pgxpool.Pool
}

func (f FactsModel) Random() (Fact, error) {
	var fact Fact

	stmt := `SELECT id, title, content, category, created FROM facts ORDER BY RANDOM() LIMIT 1`

	tx, err := f.DB.Begin(context.Background())
	if err != nil {
		return Fact{}, err
	}
	defer tx.Rollback(context.Background())

	err = tx.QueryRow(context.Background(), stmt).Scan(&fact.ID, &fact.Title, &fact.Content, &fact.Category, &fact.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Fact{}, errors.New("there's not a single fact yet")
		}
		return Fact{}, err
	}

	return fact, nil
}

func (f FactsModel) Create(fact Fact) error {
	//TODO implement me
	panic("implement me")
}

func (f FactsModel) GetAll(category, limit, offset int) ([]Fact, error) {
	//TODO implement me
	panic("implement me")
}

func (f FactsModel) Edit(fact Fact) error {
	//TODO implement me
	panic("implement me")
}

func (f FactsModel) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
