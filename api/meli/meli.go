package meli

import (
	core2 "ecommerce/api/core"
)

type TokenRequest struct {
	grantType    string `json:"grant_type"`
	clientId     string `json:"client_id"`
	clientSecret string `json:"client_secret"`
	code         string `json:"code"`
	redirectUri  string `json:"redirect_uri"`
}

type TokenResponse struct {
	accessToken  string `json:"access_token"`
	tokenType    string `json:"token_type"`
	expiresIn    int    `json:"expires_in"`
	scope        string `json:"scope"`
	userId       int    `json:"user_id"`
	refreshToken string `json:"refresh_token"`
}

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
	/*
		Once working, add state param for security
		Para mayor seguridad, te recomendamos incluir el parámetro state en la URL de autorización para asegurarte que la respuesta pertenece a una solicitud iniciada por tu aplicación.
		En caso de no tener un identificador aleatorio seguro, puedes crearlo utilizando SecureRandom y deberá ser único por cada intento de llamada (request).
	*/

	//authenticar usuario get code
	//redirect to:
	//request "http://auth.mercadolibre.com.ar/authorization?response_type=code&client_id=$APP_ID&redirect_uri=$YOUR_URL"
	//redirect to: /authenticated with response (in my redirect url: "code=xxx")
	//this code can be used to get many access tokens so probably we should keep it somewhere
	//use the code to generate a token and use it to grant authorization to the user

	//if code == nil
	//redirect to meli to get code
	//on auth endpoint, get code passed as parameter in the request
	//save the code
	//post meli to get access token.
	//return auth token to the user probably... (header)
}

func (m *MeliService) GetLastSeenProducts() []*core2.Product {
	panic("implement me")
}

func (m *MeliService) GetTrendingProducts() []*core2.Product {
	panic("implement me")
}

func (m *MeliService) GetPopularCategories() []*core2.Category {
	panic("implement me")
}

func (m *MeliService) GetFilterCategories() []*core2.Category {
	panic("implement me")
}
