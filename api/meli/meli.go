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
	authUri  = "http://auth.mercadolibre.com.ar/authorization?response_type=code&client_id=%v&redirect_uri=%v"
	tokenUri = "https://api.mercadolibre.com/oauth/token"
)

type Service struct {
	appID    string
	secret   string
	redirect string
}

func NewService(appId string, secret string, redirect string) *Service {
	return &Service{
		appID:    appId,
		secret:   secret,
		redirect: redirect,
	}
}

func (m *Service) Authenticate(w http.ResponseWriter, r *http.Request) {
	uri := fmt.Sprintf(authUri, m.appID, m.redirect)

	http.Redirect(w, r, uri, http.StatusPermanentRedirect)
}

func (m *Service) Authorize(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	body, err := json.Marshal(TokenRequest{
		GrantType:    "authorization_code",
		ClientId:     m.appID,
		ClientSecret: m.secret,
		Code:         code,
		RedirectUri:  m.redirect,
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

	//category: ropa y accesorios
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
	for i := 0; i < len(rp.Content); i++ {
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

	return products
}

func (m *Service) GetPopularCategories(ctx context.Context) []*core.Category {
	panic("implement me")
}

func (m *Service) GetFilterCategories(ctx context.Context) []*core.Category {
	panic("implement me")
}
