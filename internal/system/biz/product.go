package biz

import (
	"context"
	"time"
)

type Product struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Delete      int64     `json:"delete"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
}

func (p Product) TableName() string {
	return "product"
}

type ProductRepo interface {
	// db
	ListProduct(ctx context.Context) ([]*Product, error)
	GetProduct(ctx context.Context, id int64) (*Product, error)
	CreateProduct(ctx context.Context, product *Product) error
	UpdateProduct(ctx context.Context, id int64, product *Product) error
	DeleteProduct(ctx context.Context, id int64) error
	// redis
	GetProductLike(ctx context.Context, id int64) (rv int64, err error)
	IncProductLike(ctx context.Context, id int64) error
}
