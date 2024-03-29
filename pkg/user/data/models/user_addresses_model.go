package models

type UserAddress struct {
	BaseModel
	Address		string `gorm:"type:string;size:1000;not null"`
	Phone		string `gorm:"type:string;size:1000;not null"`
	User   User `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	UserId int
}

