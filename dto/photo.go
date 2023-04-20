package dto

import "time"

type NewPhotoRequest struct {
	Title     string `json:"title" valid:"required~title cannot be empty" example:"Fotoku"`
	Caption   string `json:"caption`
	photo_url string `json:"photo_url" valid:"required~photo url cannot be empty" example:"http://imageurl.com"`
}
type NewPhotoResponse struct {
	Result     string `json:"result"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
type PhotoResponse struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:caption`
	photo_url string    `json:"imageUrl"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type GetPhotoResponse struct {
	Result     string          `json:"result"`
	Message    string          `json:"message"`
	StatusCode int             `json:"statusCode"`
	Data       []PhotoResponse `json:"data"`
}
