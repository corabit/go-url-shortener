package models

type UrlCreateRequest struct {
	Url string `json:"url"`
}

type UrlCreateResponse struct {
	Id string `json:"id"`
	ShortUrl string `json:"short_url"`
}
