package repo

import (
	"context"
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
	//TODO implement me
	panic("implement me")
}

func (p *productRepo) GetProduct(ctx context.Context, id int64) (*biz.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p *productRepo) CreateProduct(ctx context.Context, product *biz.Product) error {
	//TODO implement me
	panic("implement me")
}

func (p *productRepo) UpdateProduct(ctx context.Context, id int64, product *biz.Product) error {
	//TODO implement me
	panic("implement me")
}

func (p *productRepo) DeleteProduct(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
