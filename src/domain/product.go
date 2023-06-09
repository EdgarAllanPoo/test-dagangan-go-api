package domain

type Product struct {
	Id       int64
	Name     string
	Price    int64
	Category string
}

type ProductRepository interface {
	SaveProduct(product Product) error
	FindAll(category string, limit, offset int) ([]*Product, int64, error)
	FindById(id int64) (*Product, error)
	UpdateProduct(id int64, product Product) error
	DeleteProduct(id int64) error
}
