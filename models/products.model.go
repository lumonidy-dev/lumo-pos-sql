package models

import "time"

// Table Products
// ID: int, primary key, auto increment
// Name: string, unique index, not null
// Price: float64, not null
// Quantity: int, omitempty
type Products struct {
	ID        int       `gorm:"type:int;primary_key;auto_increment" json:"id,omitempty"`
	Name      string    `gorm:"uniqueIndex;not null" json:"name,omitempty"`
	Price     float64   `gorm:"not null" json:"price,omitempty"`
	Quantity  int       `gorm:"omitempty" json:"quantity,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}

type CreateProductRequest struct {
	Name      string    `json:"name" binding:"required"`
	Price     float64   `json:"price" binding:"required"`
	Quantity  int       `json:"quantity,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UpdateProductRequest struct {
	Name      string    `json:"name,omitempty"`
	Price     float64   `json:"price,omitempty"`
	Quantity  int       `json:"quantity,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
