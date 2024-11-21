package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

type Service interface {
	Health() map[string]string
	Close() error
}

type service struct {
}

type Time struct {
	T time.Time
}

type Quote struct {
	ID         int    `json:"id"`
	ShortID    string `json:"shortId"`
	Quote      string `json:"quote"`
	Speaker    string `json:"speaker"`
	Source     string `json:"source"`
	SourceLink string `json:"sourceLink"`
	Created    Time   `json:"date"`
}

type QuoteStore interface {
	GetQuoteByID(ctx context.Context, id string) (*Quote, error)
	GetRandomQuote(ctx context.Context) (*Quote, error)
}

type Database struct {
	DB                 *sql.DB
	url                string
	maxOpenConnections int
	maxIdleConnections int
	log                *log.Logger
}

type DatabaseOption func(*Database)

func WithBaseUrl(url string) DatabaseOption {
	return func(c *Database) {
		c.url = url
	}
}

func WithLogger(log *log.Logger) DatabaseOption {
	return func(c *Database) {
		c.log = log
	}
}

func NewDatabase(opt ...DatabaseOption) *Database {
	d := new(Database)
	for _, o := range opt {
		o(d)
	}

	if d.maxIdleConnections == 0 {
		d.maxIdleConnections = 10
	}

	if d.maxOpenConnections == 0 {
		d.maxOpenConnections = 10
	}
	return d
}

func (d *Database) Connect() error {

	var err error
	d.DB, err = sql.Open("sqlite3", d.url)
	if err != nil {
		return err
	}

	d.DB.SetMaxIdleConns(d.maxIdleConnections)
	d.DB.SetMaxOpenConns(d.maxOpenConnections)
	return nil
}

func (d *Database) GetQouteById(ctx context.Context, id string) (*Quote, error) {
	var q Quote
	query := "SELECT shortId, quote, speaker, source FROM quotes WHERE shortId = ?;"

	var args []any
	args = append(args, id)
	err := d.DB.QueryRowContext(ctx, query, args...).Scan(&q.ShortID, &q.Quote, &q.Speaker, &q.Source)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("Error occured while querying for quote id:%v, %v", id, err.Error())
		}
		return nil, fmt.Errorf("Unknown error occured querying database for id:%v, %v", id, err.Error())
	}
	return &q, nil
}

func (d *Database) GetRandomQoute(ctx context.Context) (*Quote, error) {
	var q Quote
	query := "SELECT id, quote, speaker, source FROM quotes ORDER BY RANDOM() LIMIT 1;"

	err := d.DB.QueryRowContext(ctx, query).Scan(&q.ShortID, &q.Quote, &q.Speaker, &q.Source)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("Error occured while querying for random quote: %v", err.Error())
		}
		return nil, fmt.Errorf("Unknown error occured querying database: %v", err.Error())
	}
	return &q, nil
}
