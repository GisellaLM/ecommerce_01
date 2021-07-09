package src

import (
	"sort"
	"strconv"
	"sync"
	"time"
)

//TODO use the database/sql package to create the tables, insert the data and retrieve the results. double iterate for nested objects
//TODO does a query builder solves this nested structs problem?
type ServiceImp struct {
	//mu protects products and categories
	mu sync.Mutex

	Products   []*Product
	Categories []*Category
}

func (r *ServiceImp) Init() {
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

func (r *ServiceImp) createProduct(p uint, c uint, lastSeen time.Time) *Product {
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

func (r *ServiceImp) createCategory(c uint) *Category {
	strc := strconv.FormatUint(uint64(c), 10)
	n := "cate: " + strc
	t := c * 2
	return &Category{Id: &c, Name: &n, TimesSeen: &t}
}

func (r *ServiceImp) GetLastSeenProducts() []*Product {
	r.mu.Lock()
	//sort by last seen
	sort.SliceStable(r.Products, func(i, j int) bool {
		return r.Products[i].LastSeen.After(r.Products[j].LastSeen)
	})
	r.mu.Unlock()

	//return the first 6 elements
	return r.Products[:6]
}

func (r *ServiceImp) GetTrendingProducts() []*Product {
	r.mu.Lock()
	//sort by last seen
	sort.SliceStable(r.Products, func(i, j int) bool {
		return *r.Products[i].TimesSeen > *r.Products[j].TimesSeen
	})
	r.mu.Unlock()

	//return the first 8 elements
	return r.Products[:8]
}

func (r *ServiceImp) GetPopularCategories() []*Category {
	r.mu.Lock()
	//order categories by times seen
	sort.SliceStable(r.Categories, func(i, j int) bool {
		return *r.Categories[i].TimesSeen > *r.Categories[j].TimesSeen
	})
	r.mu.Unlock()

	//return first 6 categories
	return r.Categories[:6]
}

func (r *ServiceImp) GetFilterCategories() []*Category {
	r.mu.Lock()
	//order categories by name
	sort.SliceStable(r.Categories, func(i, j int) bool {
		return *r.Categories[i].Name < *r.Categories[j].Name
	})
	r.mu.Unlock()

	//return all
	return r.Categories
}
