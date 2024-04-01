package models

import (
	"database/sql"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	Id int `gorm:"primarykey"`

	CreatedAt  time.Time    `gorm:"type:TIMESTAMP with time zone;not null"`
	ModifiedAt sql.NullTime `gorm:"type:TIMESTAMP with time zone;null"`
	DeletedAt  sql.NullTime `gorm:"type:TIMESTAMP with time zone;null"`

	CreatedBy  int            `gorm:"not null"`
	ModifiedBy *sql.NullInt64 `gorm:"null"`
	DeletedBy  *sql.NullInt64 `gorm:"null"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")

	var userId = -1
	
	if value != nil {
		UserIdInt, _ := strconv.Atoi(value.(string))
		userId = UserIdInt
	}
	m.CreatedAt = time.Now().UTC()
	m.CreatedBy = userId
	return
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var userId = &sql.NullInt64{Valid: false}
	
	if value != nil {
		UserIdInt, _ := strconv.ParseInt(value.(string), 10, 64)
		userId = &sql.NullInt64{Valid: true, Int64: UserIdInt}
	}
	m.ModifiedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	m.ModifiedBy = userId
	return
}

func (m *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var userId = &sql.NullInt64{Valid: false}
	
	if value != nil {
		userId = &sql.NullInt64{Valid: true, Int64: int64(value.(float64))}
	}
	m.DeletedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	m.DeletedBy = userId
	return
}