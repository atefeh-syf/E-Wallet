package models

type Wallet struct {
	BaseModel
	Name        string  `gorm:"column:name;size:255" json:"name" form:"name"`
	Type        string  `gorm:"column:type;size:255" json:"type" form:"type"`
	Balance     float64 `sql:"type:decimal(64,0);"`
	Slug        string  `gorm:"column:slug;size:255" json:"slug" form:"slug"`
	Description string  `gorm:"column:description;size:255" json:"description" form:"description"`
	UserId      uint
	User        User `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
}
