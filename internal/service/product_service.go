package service

import (
	"fmt"
	"merchant/internal/entity"
	"merchant/internal/repository"
)

type ProductService struct {
	productRepo repository.IProductRepository
}

type IProductService interface {
	EditStore(map[int][]entity.Product) (string, error)
}

func NewProductService(productRepo repository.IProductRepository) *ProductService {
	var productService = ProductService{}
	productService.productRepo = productRepo
	return &productService
}

func (s *ProductService) EditStore(productMap map[int][]entity.Product) (string, error) {
	var idFromDB int
	id := 0
	for k, _ := range productMap {
		id = k
	}
	rows, err := s.productRepo.GetProductsByCustomerID(id)
	if err != nil {
		return "", fmt.Errorf("Error in getting data from DB: %s", err.Error())
	}

	productsExisted := make([]entity.Product, 0)
	for rows.Next() {
		product := entity.Product{}
		if err := rows.Scan(&idFromDB, &product.OfferID, &product.Name, &product.Price, &product.Quantity); err != nil {
			return "", fmt.Errorf("Error in scanning rows: %s", err.Error())
		}
		productsExisted = append(productsExisted, product)
	}

	inserted := 0
	updated := 0
	deleted := 0
	for _, valIn := range productMap[id] {
		check := false
		for _, valExist := range productsExisted {
			if valIn.OfferID == valExist.OfferID {
				check = true
				if valIn.Available == true {
					valIn.Quantity = valExist.Quantity + valIn.Quantity
				} else {
					valIn.Quantity = valExist.Quantity - valIn.Quantity
				}
				if valIn.Quantity <= 0 {
					err = s.productRepo.Delete(id, valIn)
					if err != nil {
						return "", fmt.Errorf("Error in deleting row: %s", err.Error())
					}
					deleted++
				} else {
					err = s.productRepo.Update(id, valIn)
					if err != nil {
						return "", fmt.Errorf("Error in updating row: %s", err.Error())
					}
					updated++
				}
			}
		}
		if check == false {
			err = s.productRepo.Insert(id, valIn)
			if err != nil {
				return "", fmt.Errorf("Error in inserting row: %s", err.Error())
			}
			inserted++
		}
	}

	return fmt.Sprintf("Создано новых товаров: %v \n Обновлено товаров: %v \n Удалено товаров: %v", inserted, updated, deleted), nil

}
