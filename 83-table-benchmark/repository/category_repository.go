package repository

import "79-mock/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
