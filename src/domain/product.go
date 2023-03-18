package domain

type Product struct {
	Id    int64
	Name  string
	Price int64
}

type ProductRepository interface {
	SaveProduct(product Product) error
	FindAll() ([]*Product, error)
	FindById(id int64) (*Product, error)
	UpdateProduct(id int64, product Product) error
	DeleteProduct(id int64) error
}
