package product

import (
	"database/sql"
	"productGofr/models"
	"productGofr/stores"
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type product struct {
}

func New() stores.Istore {
	return product{}
}

func (p product) UserById(ctx *gofr.Context, id int) (*models.Product, error) {
	var prd models.Product
	rows := ctx.DB().QueryRow("select * from Product where id = ?", id)
	if rows.Err() != nil {
		return &prd, rows.Err()
	}
	err := rows.Scan(&prd.Id, &prd.Name, &prd.Type)
	if err == sql.ErrNoRows {
		return &prd, errors.EntityNotFound{Entity: "products", ID: strconv.Itoa(id) /*fmt.Sprint(id)*/}
	}
	return &prd, nil
}