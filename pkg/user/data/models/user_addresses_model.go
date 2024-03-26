package models

type Address struct {
	BaseModel
	Address		string `gorm:"type:string;size:1000;not null"`
	Phone		string `gorm:"type:string;size:1000;not null"`
	UserAddresses *[]UserAddress

}

type UserAddress struct {
	BaseModel
	User   User `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	Address   Address `gorm:"foreignKey:AddressId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	UserId int
	AddressId int
}

