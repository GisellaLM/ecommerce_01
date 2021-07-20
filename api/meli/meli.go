package meli

import (
	"bytes"
	"context"
	"ecommerce/api/core"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	trendingCacheKey = "meli-trending"
	authUri          = "http://auth.mercadolibre.com.ar/authorization?response_type=code&client_id=%v&redirect_uri=%v"
	tokenUri         = "https://api.mercadolibre.com/oauth/token"
)

type Config struct {
	AppID    string
	Secret   string
	Redirect string
	Cache    core.Cache
}

type Service struct {
	Config
}

func NewService(c Config) *Service {
	return &Service{
		c,
	}
}

func (m *Service) Authenticate(w http.ResponseWriter, r *http.Request) {
	uri := fmt.Sprintf(authUri, m.AppID, m.Redirect)

	http.Redirect(w, r, uri, http.StatusPermanentRedirect)
}

func (m *Service) Authorize(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	body, err := json.Marshal(TokenRequest{
		GrantType:    "authorization_code",
		ClientId:     m.AppID,
		ClientSecret: m.Secret,
		Code:         code,
		RedirectUri:  m.Redirect,
	})

	if err != nil {
		fmt.Println(err)
	}

	req, _ := http.NewRequest("POST", tokenUri, bytes.NewBuffer(body))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	c := http.Client{}
	res, err := c.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	var tr TokenResponse
	err = json.NewDecoder(res.Body).Decode(&tr)

	if err != nil {
		fmt.Println(err)
	}

	//save all and redirect to home
	fmt.Println(tr)
}

func (m *Service) GetLastSeenProducts(ctx context.Context) []*core.Product {
	panic("implement me")
}

func (m *Service) GetTrendingProducts(ctx context.Context) []*core.Product {
	tkn := ctx.Value("token").(string)

	if prods := m.Cache.Get(trendingCacheKey); prods != nil {
		return prods.([]*core.Product)
	}

	//category: clothes and accessories - ropa y accesorios
	uri := "https://api.mercadolibre.com/highlights/MLA/category/MLA1430"
	req, _ := http.NewRequest("GET", uri, nil)

	req.Header.Set("Authorization", "Bearer "+tkn)
	req.Header.Set("Content-Type", "application/json")

	c := &http.Client{}

	res, err := c.Do(req)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var rp TrendingProductsResponse
	json.NewDecoder(res.Body).Decode(&rp)

	//type product: https://api.mercadolibre.com/products/
	var products []*core.Product
	for i := 0; i < 3; i++ { //len(rp.Content)
		if rp.Content[i].Type == "ITEM" {
			url := "https://api.mercadolibre.com/items/" + rp.Content[i].Id + "?attributes=id,title,subtitle,thumbnail,available_quantity,price"
			req, _ := http.NewRequest("GET", url, nil)

			req.Header.Set("Authorization", "Bearer "+tkn)
			req.Header.Set("Content-Type", "application/json")

			res, err := c.Do(req)

			if err != nil {
				fmt.Println(err)
				return nil
			}

			var item GetItemResponse
			err = json.NewDecoder(res.Body).Decode(&item)

			if err != nil {
				fmt.Println(err)
				continue
			}

			id := uint(i)
			q := uint(item.AvailableQuantity)

			products = append(products, &core.Product{
				Id:                &id,
				Name:              &item.Title,
				Description:       &item.Title,
				ImageLink:         &item.Thumbnail,
				CategoryId:        nil,
				AvailableQuantity: &q,
				Link:              &item.Permalink,
				Stars:             [5]*uint{},
				Variations:        nil,
				Shipping:          nil,
				Price:             &item.Price,
				Installments:      nil,
				TimesSeen:         nil,
				LastSeen:          time.Time{},
			})
		}
	}

	//cache
	//add auth user token or something like that to the cache
	//maybe we should query meli to get some of the user data like its id o name
	//something we can use as a unique key.
	m.Cache.Set(trendingCacheKey, products)

	return products
}

func (m *Service) GetPopularCategories(ctx context.Context) []*core.Category {
	panic("implement me")
}

func (m *Service) GetFilterCategories(ctx context.Context) []*core.Category {
	panic("implement me")
}
