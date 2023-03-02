package entity

import "mime/multipart"

type Customer struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type Product struct {
	//ID        int     `json:"id"`
	OfferID   int     `json:"offerID"`
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
	Quantity  int     `json:"quantity"`
	Available bool    `json:"available"`
}

type FormData struct {
	ID   int                   `form:"id"`
	File *multipart.FileHeader `form:"file"`
}
