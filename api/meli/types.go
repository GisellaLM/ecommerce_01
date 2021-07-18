package meli

import "time"

type TokenRequest struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	RedirectUri  string `json:"redirect_uri"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	UserId       int    `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
}

type TrendingProductsResponse struct {
	QueryData struct {
		Id            string `json:"id"`
		Criteria      string `json:"criteria"`
		HighlightType string `json:"highlight_type"`
	} `json:"query_data"`
	Content []struct {
		Id       string `json:"id"`
		Position int    `json:"position"`
		Type     string `json:"type"`
	} `json:"content"`
}

type GetItemResponse struct {
	ID                string      `json:"id"`
	SiteID            string      `json:"site_id"`
	Title             string      `json:"title"`
	Subtitle          interface{} `json:"subtitle"`
	SellerID          int         `json:"seller_id"`
	CategoryID        string      `json:"category_id"`
	OfficialStoreID   int         `json:"official_store_id"`
	Price             float64     `json:"price"`
	BasePrice         float64     `json:"base_price"`
	OriginalPrice     interface{} `json:"original_price"`
	CurrencyID        string      `json:"currency_id"`
	InitialQuantity   int         `json:"initial_quantity"`
	AvailableQuantity int         `json:"available_quantity"`
	SoldQuantity      int         `json:"sold_quantity"`
	SaleTerms         []struct {
		ID          string      `json:"id"`
		Name        string      `json:"name"`
		ValueID     string      `json:"value_id"`
		ValueName   string      `json:"value_name"`
		ValueStruct interface{} `json:"value_struct"`
		Values      []struct {
			ID     string      `json:"id"`
			Name   string      `json:"name"`
			Struct interface{} `json:"struct"`
		} `json:"values"`
	} `json:"sale_terms"`
	BuyingMode      string    `json:"buying_mode"`
	ListingTypeID   string    `json:"listing_type_id"`
	StartTime       time.Time `json:"start_time"`
	StopTime        time.Time `json:"stop_time"`
	Condition       string    `json:"condition"`
	Permalink       string    `json:"permalink"`
	ThumbnailID     string    `json:"thumbnail_id"`
	Thumbnail       string    `json:"thumbnail"`
	SecureThumbnail string    `json:"secure_thumbnail"`
	Pictures        []struct {
		ID        string `json:"id"`
		URL       string `json:"url"`
		SecureURL string `json:"secure_url"`
		Size      string `json:"size"`
		MaxSize   string `json:"max_size"`
		Quality   string `json:"quality"`
	} `json:"pictures"`
	VideoID      interface{} `json:"video_id"`
	Descriptions []struct {
		ID string `json:"id"`
	} `json:"descriptions"`
	AcceptsMercadopago           bool          `json:"accepts_mercadopago"`
	NonMercadoPagoPaymentMethods []interface{} `json:"non_mercado_pago_payment_methods"`
	Shipping                     struct {
		Mode        string `json:"mode"`
		FreeMethods []struct {
			ID   int `json:"id"`
			Rule struct {
				Default          bool        `json:"default"`
				FreeMode         string      `json:"free_mode"`
				FreeShippingFlag bool        `json:"free_shipping_flag"`
				Value            interface{} `json:"value"`
			} `json:"rule"`
		} `json:"free_methods"`
		Tags         []string    `json:"tags"`
		Dimensions   interface{} `json:"dimensions"`
		LocalPickUp  bool        `json:"local_pick_up"`
		FreeShipping bool        `json:"free_shipping"`
		LogisticType string      `json:"logistic_type"`
		StorePickUp  bool        `json:"store_pick_up"`
	} `json:"shipping"`
	InternationalDeliveryMode string `json:"international_delivery_mode"`
	SellerAddress             struct {
		City struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"city"`
		State struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"state"`
		Country struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"country"`
		SearchLocation struct {
			Neighborhood struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"neighborhood"`
			City struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"city"`
			State struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"state"`
		} `json:"search_location"`
		ID int `json:"id"`
	} `json:"seller_address"`
	SellerContact interface{} `json:"seller_contact"`
	Location      struct {
	} `json:"location"`
	CoverageAreas []interface{} `json:"coverage_areas"`
	Attributes    []struct {
		ID          string      `json:"id"`
		Name        string      `json:"name"`
		ValueID     string      `json:"value_id"`
		ValueName   string      `json:"value_name"`
		ValueStruct interface{} `json:"value_struct"`
		Values      []struct {
			ID     string      `json:"id"`
			Name   string      `json:"name"`
			Struct interface{} `json:"struct"`
		} `json:"values"`
		AttributeGroupID   string `json:"attribute_group_id"`
		AttributeGroupName string `json:"attribute_group_name"`
	} `json:"attributes"`
	Warnings      []interface{} `json:"warnings"`
	ListingSource string        `json:"listing_source"`
	Variations    []struct {
		ID                    int64   `json:"id"`
		Price                 float64 `json:"price"`
		AttributeCombinations []struct {
			ID          string      `json:"id"`
			Name        string      `json:"name"`
			ValueID     string      `json:"value_id"`
			ValueName   string      `json:"value_name"`
			ValueStruct interface{} `json:"value_struct"`
			Values      []struct {
				ID     string      `json:"id"`
				Name   string      `json:"name"`
				Struct interface{} `json:"struct"`
			} `json:"values"`
		} `json:"attribute_combinations"`
		AvailableQuantity int           `json:"available_quantity"`
		SoldQuantity      int           `json:"sold_quantity"`
		SaleTerms         []interface{} `json:"sale_terms"`
		PictureIds        []string      `json:"picture_ids"`
		CatalogProductID  interface{}   `json:"catalog_product_id"`
	} `json:"variations"`
	Status              string        `json:"status"`
	SubStatus           []interface{} `json:"sub_status"`
	Tags                []string      `json:"tags"`
	Warranty            string        `json:"warranty"`
	CatalogProductID    interface{}   `json:"catalog_product_id"`
	DomainID            string        `json:"domain_id"`
	ParentItemID        interface{}   `json:"parent_item_id"`
	DifferentialPricing interface{}   `json:"differential_pricing"`
	DealIds             []interface{} `json:"deal_ids"`
	AutomaticRelist     bool          `json:"automatic_relist"`
	DateCreated         time.Time     `json:"date_created"`
	LastUpdated         time.Time     `json:"last_updated"`
	Health              int           `json:"health"`
	CatalogListing      bool          `json:"catalog_listing"`
	Channels            []string      `json:"channels"`
}
