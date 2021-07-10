package src

import (
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	_ "github.com/lib/pq"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

//TODO use the database/sql package to create the tables, insert the data and retrieve the results. double iterate for nested objects
//TODO does a query builder solves this nested structs problem?
type InMemoryService struct {
	//mu protects products and categories
	mu sync.Mutex

	Products   []*Product
	Categories []*Category
}

func (r *InMemoryService) Init() {
	//Dummy products
	for i := 1; i <= 10; i++ {
		pi := r.createProduct(uint(i), 1, time.Now())

		r.Products = append(r.Products, pi)
	}

	//Dummy categories
	for i := 1; i <= 10; i++ {
		ci := r.createCategory(uint(i))

		r.Categories = append(r.Categories, ci)
	}
}

func (r *InMemoryService) createProduct(p uint, c uint, lastSeen time.Time) *Product {
	strp := strconv.FormatUint(uint64(p), 10)

	name, desc, image := "p: "+strp, "desc: "+strp, "image: "+strp
	quantity, timesSeen := p*2, p*10

	return &Product{
		Id:                &p,
		Name:              &name,
		Description:       &desc,
		ImageLink:         &image,
		CategoryId:        &c,
		AvailableQuantity: &quantity,
		Variations:        nil,
		Shipping:          nil,
		Price:             nil,
		Installments:      nil,
		TimesSeen:         &timesSeen,
		LastSeen:          lastSeen,
	}
}

func (r *InMemoryService) createCategory(c uint) *Category {
	strc := strconv.FormatUint(uint64(c), 10)
	n := "cate: " + strc
	t := c * 2
	return &Category{Id: &c, Name: &n, TimesSeen: &t}
}

func (r *InMemoryService) GetLastSeenProducts() []*Product {
	r.mu.Lock()
	//sort by last seen
	sort.SliceStable(r.Products, func(i, j int) bool {
		return r.Products[i].LastSeen.After(r.Products[j].LastSeen)
	})
	r.mu.Unlock()

	//return the first 6 elements
	return r.Products[:6]
}

func (r *InMemoryService) GetTrendingProducts() []*Product {
	r.mu.Lock()
	//sort by last seen
	sort.SliceStable(r.Products, func(i, j int) bool {
		return *r.Products[i].TimesSeen > *r.Products[j].TimesSeen
	})
	r.mu.Unlock()

	//return the first 8 elements
	return r.Products[:8]
}

func (r *InMemoryService) GetPopularCategories() []*Category {
	r.mu.Lock()
	//order categories by times seen
	sort.SliceStable(r.Categories, func(i, j int) bool {
		return *r.Categories[i].TimesSeen > *r.Categories[j].TimesSeen
	})
	r.mu.Unlock()

	//return first 6 categories
	return r.Categories[:6]
}

func (r *InMemoryService) GetFilterCategories() []*Category {
	r.mu.Lock()
	//order categories by name
	sort.SliceStable(r.Categories, func(i, j int) bool {
		return *r.Categories[i].Name < *r.Categories[j].Name
	})
	r.mu.Unlock()

	//return all
	return r.Categories
}

type PostgresService struct {
	db *sqlx.DB
}

func NewPostgresService(db *sqlx.DB) *PostgresService {
	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
	return &PostgresService{
		db: db,
	}
}

func (p PostgresService) GetLastSeenProducts() []*Product {
	var ps []*Product

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

func (p PostgresService) GetTrendingProducts() []*Product {
	panic("implement me")
}

func (p PostgresService) GetPopularCategories() []*Category {
	panic("implement me")
}

func (p PostgresService) GetFilterCategories() []*Category {
	panic("implement me")
}
