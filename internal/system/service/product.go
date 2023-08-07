package service

import (
	"context"
	v1 "github.com/liang21/terminator/api/system/v1"
	"github.com/liang21/terminator/internal/system/biz"
	"github.com/liang21/terminator/pkg/pagination"
	"strconv"
)

type ProductService struct {
	v1.UnimplementedProductServiceServer
	product *biz.ProductUsecase
}

func NewProductService(product *biz.ProductUsecase) *ProductService {
	return &ProductService{product: product}
}

func (p *ProductService) ListProduct(ctx context.Context, req *v1.ListProductRequest) (*v1.ListProductReply, error) {
	parseIndex, err := strconv.ParseInt(req.GetPageToken(), 10, 64)
	if err != nil {
		return nil, err
	}
	meta := pagination.ListMeta{
		PageSize:  int64(req.GetPageSize()),
		PageToken: parseIndex,
	}
	products, err := p.product.List(ctx, meta)
	product := make([]*v1.Product, 0, len(products.Items))
	for _, item := range products.Items {
		product = append(product, &v1.Product{Id: item.Id, Name: item.Name, Description: item.Description})
	}
	if err != nil {
		return nil, err
	}
	return &v1.ListProductReply{Total: products.TotalCount, Results: product}, nil
}

func (p *ProductService) GetProduct(ctx context.Context, req *v1.GetProductRequest) (*v1.GetProductReply, error) {
	product, err := p.product.Get(ctx, req.GetId())
	return &v1.GetProductReply{Product: &v1.Product{Name: product.Name, Description: product.Description}}, err
}

func (p *ProductService) CreateProduct(ctx context.Context, req *v1.CreateProductRequest) (*v1.CreateProductReply, error) {
	err := p.product.Create(ctx, &biz.Product{Name: req.GetName(), Description: req.GetDescription()})
	return &v1.CreateProductReply{Product: &v1.Product{Name: req.Name, Description: req.Description}}, err
}

func (p *ProductService) UpdateProduct(ctx context.Context, req *v1.UpdateProductRequest) (*v1.UpdateProductReply, error) {
	err := p.product.Update(ctx, req.GetId(), &biz.Product{Id: req.GetId(), Name: req.GetName(), Description: req.GetDescription()})
	return &v1.UpdateProductReply{}, err
}

func (p *ProductService) DeleteProduct(ctx context.Context, req *v1.DeleteProductRequest) (*v1.DeleteProductReply, error) {
	err := p.product.Delete(ctx, req.GetId())
	return &v1.DeleteProductReply{}, err
}
