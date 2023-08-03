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
	Delete      int64     `json:"delete"`
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
	Items      []*Project `json:"items"` //数据
}

func (uc *ProductUsecase) List(ctx context.Context, meta pagination.ListMeta) (productDtoList *ProductDTOList, err error) {
	productDtoList, err = uc.repo.ListProduct(ctx, meta)
	if err != nil {
		return
	}
	return productDtoList, nil
}

func (uc *ProductUsecase) Get(ctx context.Context, id int64) (product *Product, err error) {
	if id == 0 {
		return nil, errors.New("id is empty")
	}
	product, err = uc.repo.GetProduct(ctx, id)
	if err != nil {
		return
	}
	return product, nil
}

func (uc *ProductUsecase) Create(ctx context.Context, product *Product) error {
	return uc.repo.CreateProduct(ctx, product)
}

func (uc *ProductUsecase) Update(ctx context.Context, id int64, product *Product) error {
	if id == 0 {
		return errors.New("id is empty")
	}
	return uc.repo.UpdateProduct(ctx, id, product)
}

func (uc *ProductUsecase) Delete(ctx context.Context, id int64) error {
	if id == 0 {
		return errors.New("id is empty")
	}
	return uc.repo.DeleteProduct(ctx, id)
}
