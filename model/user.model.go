package model

type User struct {
	Id        uint64 `json:"id,omitempty" gorm:"primary_key;auto_increment" `
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	IsActive  bool   `json:"isActive,omitempty" gorm:"default:true"`
	Password  string `json:"password,omitempty"`
	Salt      string `json:"salt,omitempty"`
	Role      string `json:"role,omitempty"`
	IsDeleted bool   `json:"isDeleted,omitempty" gorm:"default:false"`
	CreatedAt int    `json:"createdAt" gorm:"autoCreateTime:nano"`
	UpdatedAt int    `json:"updatedAt" gorm:"autoUpdateTime:nano"`
}
