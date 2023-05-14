package model

type Post struct {
	Id        uint64 `json:"id,omitempty" gorm:"primary_key;auto_increment" `
	Post      string `json:"post,omitempty"`
	IsDeleted bool   `json:"isDeleted,omitempty" gorm:"default:false"`
	UserEmail string `json:"userEmail,omitempty"`
	CreatedAt int    `json:"createdAt" gorm:"autoCreateTime:nano"`
	UpdatedAt int    `json:"updatedAt" gorm:"autoUpdateTime:nano"`
}
