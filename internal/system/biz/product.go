package biz

import (
	"context"
	"errors"
	"github.com/liang21/terminator/pkg/pagination"
	"time"
)

type Product struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Deleted     int64     `json:"deleted"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
}

func (p Product) TableName() string {
	return "product"
}

type ProductRepo interface {
	// db
	ListProduct(ctx context.Context, meta pagination.ListMeta) (*ProductDTOList, error)
	GetProduct(ctx context.Context, id int64) (*Product, error)
	CreateProduct(ctx context.Context, product *Product) error
	UpdateProduct(ctx context.Context, id int64, product *Product) error
	DeleteProduct(ctx context.Context, id int64) error
}

type ProductUsecase struct {
	repo ProductRepo
}

func NewProductUsecase(repo ProductRepo) *ProductUsecase {
	return &ProductUsecase{repo: repo}
}

type ProductDTOList struct {
	TotalCount int64      `json:"total"` //总数
	Items      []*Product `json:"items"` //数据
}

func (pu *ProductUsecase) List(ctx context.Context, meta pagination.ListMeta) (productDtoList *ProductDTOList, err error) {
	productDtoList, err = pu.repo.ListProduct(ctx, meta)
	if err != nil {
		return
	}
	return productDtoList, nil
}

func (pu *ProductUsecase) Get(ctx context.Context, id int64) (product *Product, err error) {
	if id == 0 {
		return nil, errors.New("id is empty")
	}
	product, err = pu.repo.GetProduct(ctx, id)
	if err != nil {
		return
	}
	return product, nil
}

func (pu *ProductUsecase) Create(ctx context.Context, product *Product) error {
	return pu.repo.CreateProduct(ctx, product)
}

func (pu *ProductUsecase) Update(ctx context.Context, id int64, product *Product) error {
	if id == 0 {
		return errors.New("id is empty")
	}
	return pu.repo.UpdateProduct(ctx, id, product)
}

func (pu *ProductUsecase) Delete(ctx context.Context, id int64) error {
	if id == 0 {
		return errors.New("id is empty")
	}
	return pu.repo.DeleteProduct(ctx, id)
}
