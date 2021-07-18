package meli

import "ecommerce/api/src/core"

type MeliService struct {
	AppID  string
	Secret string
}

func NewService(appId string, secret string) *MeliService {
	return &MeliService{
		AppID:  appId,
		Secret: secret,
	}
}

func (m *MeliService) Authenticated() {
	"http://auth.mercadolibre.com.ar/authorization?response_type=code&client_id=$APP_ID&redirect_uri=$YOUR_URL"
}

func (m *MeliService) GetLastSeenProducts() []*core.Product {
	panic("implement me")
}

func (m *MeliService) GetTrendingProducts() []*core.Product {
	panic("implement me")
}

func (m *MeliService) GetPopularCategories() []*core.Category {
	panic("implement me")
}

func (m *MeliService) GetFilterCategories() []*core.Category {
	panic("implement me")
}
