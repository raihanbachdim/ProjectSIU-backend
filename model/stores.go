package model

type Stores struct {
	Id          uint   `json:"id"`
	StoreName   string `json:store_name`
	Description string `json:description`
	Image       string `json:image`
	UserID      string `json:userid`
	User        User   `json:"user";gorm:"foreignkey:UserID"`
}
