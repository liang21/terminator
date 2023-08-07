package repo

import (
	"context"
	"errors"

	"github.com/liang21/terminator/internal/system/biz"
	"github.com/liang21/terminator/pkg/pagination"
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"
)

type productRepo struct {
	db  *xorm.Engine
	rdb *redis.Client
}

func NewProductRepo(db *xorm.Engine, rdb *redis.Client) biz.ProductRepo {
	return &productRepo{
		db:  db,
		rdb: rdb,
	}
}

func (p *productRepo) ListProduct(ctx context.Context, meta pagination.ListMeta) (*biz.ProductDTOList, error) {
	var products []*biz.Product
	offset := pagination.GetPageOffset(meta.PageSize, meta.PageToken)
	count, err := p.db.Limit(int(meta.PageSize), int(offset)).Desc("name").And("deleted = ?", 0).FindAndCount(&products)
	if err != nil {
		return nil, err
	}
	productDtoList := &biz.ProductDTOList{}
	productDtoList.TotalCount = count
	productDtoList.Items = products
	return productDtoList, nil
}

func (p *productRepo) GetProduct(ctx context.Context, id int64) (*biz.Product, error) {
	product := &biz.Product{Deleted: 0}
	err := p.db.Where("id = ?", id).Find(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *productRepo) CreateProduct(ctx context.Context, product *biz.Product) error {
	result, err := p.db.Insert(product)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("product create failed")
	}
	return nil
}

func (p *productRepo) UpdateProduct(ctx context.Context, id int64, product *biz.Product) error {
	result, err := p.db.Where("id = ?", id).Update(product)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("product update failed")
	}
	return nil
}

func (p *productRepo) DeleteProduct(ctx context.Context, id int64) error {
	result, err := p.db.Where("id = ?", id).Update(&biz.Product{Deleted: 1})
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("product delete failed")
	}
	return nil
}
