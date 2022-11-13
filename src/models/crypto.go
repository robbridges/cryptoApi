package models

type Crypto struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Amount_Owned float64 `json:"amount_owned"`
	Image_Src    string  `json:"image_src"`
}
