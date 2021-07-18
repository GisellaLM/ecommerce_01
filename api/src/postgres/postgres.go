package postgres

import (
	"ecommerce/api/src/core"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	"strings"
)

//TODO use the database/sql package to create the tables, insert the data and retrieve the results. double iterate for nested objects
//TODO does a query builder solves this nested structs problem?
type PostgresService struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) *PostgresService {
	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
	return &PostgresService{
		db: db,
	}
}

func (p *PostgresService) GetLastSeenProducts() []*core.Product {
	var ps []*core.Product

	//check nested properties
	//probably I should iterate and set each one my self
	p.db.Select(&ps, `
		SELECT 
		id,
		name,
		description,
		image_link,
		category_id,
		available_quantity,
		stars,
		variations,
		shipping,
		price,
		installments,
		times_seen, 
		last_seen,  
		FROM products
	`)

	return ps
}

func (p *PostgresService) GetTrendingProducts() []*core.Product {
	panic("implement me")
}

func (p *PostgresService) GetPopularCategories() []*core.Category {
	panic("implement me")
}

func (p *PostgresService) GetFilterCategories() []*core.Category {
	panic("implement me")
}
