package product

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	dto "github.com/mohammaderm/krad/internal/dto/product"
	"github.com/mohammaderm/krad/internal/models"
	"github.com/mohammaderm/krad/log"
)

var (
	GLTProduct    = "SELECT product.id, product.productname, product.productnameen, product.imagepath, product.available, product.createdat, product.categoryid, category.categoryname, price.discountstatus, price.discountpercent, price.pricedis, price.price FROM product INNER JOIN category ON product.categoryid = category.id Inner join price ON product.priceid = price.id ORDER BY product.createdat DESC LIMIT 10;"
	GetByID       = "SELECT * FROM  product WHERE id = ?;"
	GetByCategory = "SELECT DISTINCT p.id , p.productname , p.imagepath , p.available, p.createdat, p2.discountstatus, p2.discountpercent , p2.price , p2.pricedis FROM product p INNER JOIN featuresvalueproduct f ON f.productid = p.id INNER JOIN featuresvalue f2 ON f.featuresvalueid = f2.id INNER JOIN price p2 ON p.priceid = p2.id WHERE p.categoryid = ? AND f2.id in (?) ORDER BY ? DESC LIMIT ? OFFSET ?;"
	limit         = 15
)

type (
	repository struct {
		logger log.Logger
		db     *sqlx.DB
	}
	ProductRepository interface {
		GLTProduct(ctx context.Context) ([]dto.GetProducts, error)
		GetByID(ctx context.Context, id int) (*models.Product, error)
		GetByCategory(ctx context.Context, offset, id int, filter []string, order string) ([]dto.GetProductsByCategory, error)
	}
)

func NewRepository(con *sqlx.DB, logger log.Logger) ProductRepository {
	return &repository{
		db:     con,
		logger: logger,
	}
}

// get last 10 row in product tables (GLTProduct)
func (r *repository) GLTProduct(ctx context.Context) ([]dto.GetProducts, error) {
	var result []dto.GetProducts
	err := r.db.SelectContext(ctx, &result, GLTProduct)
	if err != nil {
		r.logger.Error(&log.Field{
			Package:  "repository.product",
			Function: "GLTProduct",
			Params:   "_",
			Message:  err.Error(),
		})
		return nil, err
	}
	return result, nil
}

// get product by id
func (r *repository) GetByID(ctx context.Context, id int) (*models.Product, error) {
	var result models.Product
	err := r.db.GetContext(ctx, &result, GetByID, id)
	if errors.Is(err, sql.ErrNoRows) {
		r.logger.Error(&log.Field{
			Package:  "repository.product",
			Function: "GetByID",
			Params:   "_",
			Message:  err.Error(),
		})
		return nil, fmt.Errorf("%w:product not found", err)
	}
	if err != nil {
		r.logger.Error(&log.Field{
			Package:  "repository.product",
			Function: "GetByID",
			Params:   "_",
			Message:  err.Error(),
		})
		return nil, fmt.Errorf("failed to fetch product: %w", err)
	}
	return &result, nil

}

func (r *repository) GetByCategory(ctx context.Context, offset, id int, filter []string, order string) ([]dto.GetProductsByCategory, error) {
	var result []dto.GetProductsByCategory
	query, args, err := sqlx.In(GetByCategory, id, filter, order, limit, offset)
	if err != nil {
		return nil, err
	}
	query = r.db.Rebind(query)
	err = r.db.SelectContext(ctx, &result, query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}
