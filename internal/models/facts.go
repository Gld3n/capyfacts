package models

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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

func ValidateCategory(category string) (*Category, error) {
	castedCategory := Category(category)
	if castedCategory != Diet && castedCategory != Habitat && castedCategory != Behavior {
		return nil, fmt.Errorf("invalid category: '%s'", category)
	}
	return &castedCategory, nil
}

func NewFact(title, content, category string) (*Fact, error) {
	validatedCategory, err := ValidateCategory(category)
	if err != nil {
		return nil, err
	}

	return &Fact{
		Title:     title,
		Content:   content,
		Category:  *validatedCategory,
		CreatedAt: time.Now(),
	}, nil
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
	stmt := `INSERT INTO facts (title, content, category, created) VALUES (@title, @content, @category, NOW())`

	tx, err := f.DB.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't begin transaction: %s", err.Error())
	}
	defer tx.Rollback(context.Background())

	args := pgx.NamedArgs{
		"title":    fact.Title,
		"content":  fact.Content,
		"category": fact.Category,
	}

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
	var stmt strings.Builder

	stmt.WriteString(`SELECT id, title, content, category, created FROM facts WHERE isDeleted = FALSE`)
	if category != "" {
		stmt.WriteString(` AND category = @category`)
	}
	stmt.WriteString(` ORDER BY created LIMIT @limit OFFSET @offset`)

	args := pgx.NamedArgs{
		"category": category,
		"limit":    limit,
		"offset":   offset,
	}

	rows, err := f.DB.Query(context.Background(), stmt.String(), args)
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
	var setClauses []string
	args := pgx.NamedArgs{"id": fact.ID}

	if len(fact.Title) > 0 {
		setClauses = append(setClauses, "title = @title")
		args["title"] = fact.Title
	}
	if len(fact.Content) > 0 {
		setClauses = append(setClauses, "content = @content")
		args["content"] = fact.Content
	}
	if len(fact.Category) > 0 {
		setClauses = append(setClauses, "category = @category")
		args["category"] = fact.Category
	}

	if len(setClauses) == 0 {
		return fmt.Errorf("no fields to update")
	}

	stmt := fmt.Sprintf(`UPDATE facts SET %s WHERE id = @id`, strings.Join(setClauses, ", "))

	tx, err := f.DB.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't begin transaction: %s", err.Error())
	}
	defer tx.Rollback(context.Background())

	cmdTag, err := tx.Exec(context.Background(), stmt, args)
	if err != nil {
		return fmt.Errorf("couldn't update fact %d: %s", fact.ID, err.Error())
	}

	if cmdTag.RowsAffected() != 1 {
		return fmt.Errorf("couldn't update fact: no fact with id %d", fact.ID)
	}

	if err = tx.Commit(context.Background()); err != nil {
		return fmt.Errorf("couldn't commit transaction: %s", err.Error())
	}

	return nil
}

func (f FactsModel) Delete(id int) error {
	stmt := `UPDATE facts SET isDeleted = TRUE WHERE id = $1`

	tx, err := f.DB.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't begin transaction: %s", err.Error())
	}
	defer tx.Rollback(context.Background())

	cmdTag, err := tx.Exec(context.Background(), stmt, id)
	if err != nil {
		return fmt.Errorf("couldn't delete fact %d: %s", id, err.Error())
	}

	if cmdTag.RowsAffected() != 1 {
		return fmt.Errorf("couldn't delete fact: no fact with id %d", id)
	}

	if err = tx.Commit(context.Background()); err != nil {
		return fmt.Errorf("couldn't commit transaction: %s", err.Error())
	}

	return nil
}
