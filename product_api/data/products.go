package data

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"regexp"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeleteOn    string  `json:"-"`
}

type Products []*Product

var ProductNotFoundError = fmt.Errorf("product not found")

func (p *Product) Validate() error {
	validator := validator.New()
	validator.RegisterValidation("sku", validateSKU)
	return validator.Struct(p)
}

func validateSKU(field validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(field.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}

	return true
}

func (p *Products) ToJSON(writer io.Writer) error {
	e := json.NewEncoder(writer)
	return e.Encode(p)
}

func (p *Product) FromJSON(reader io.Reader) error {
	e := json.NewDecoder(reader)
	return e.Decode(p)
}

func findNextID() int {
	lastProduct := productList[len(productList)-1]
	return lastProduct.ID + 1
}

func findId(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ProductNotFoundError
}

func findIndexByProduct(id int) int {
	for i, p := range productList {

		if p.ID == id {
			return i
		}
	}

	return -1
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = findNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findId(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

func RemoveProduct(id int) error {
	pos := findIndexByProduct(id)
	if pos == -1 {
		return ProductNotFoundError
	}

	productList = append(productList[:pos], productList[pos+1])
	return nil
}

var productList = []*Product{

	&Product{
		ID:          1,
		Name:        "Capachino",
		Description: "Milk Coffee",
		Price:       2.45,
		SKU:         "dcnjd",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Capachino expresso",
		Description: "premium Coffee",
		Price:       6.45,
		SKU:         "dfbu",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
