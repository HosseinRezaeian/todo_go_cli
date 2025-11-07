package memorystore

import (
	"todo_cli/entity"
)

type Category struct {
	categories []entity.Category
}

func (c Category) DoesThisUserHaveThisCategoryID(userId, categoryId int) bool {
	iffound := false
	for _, c := range c.categories {
		if c.Id == categoryId && c.UserId == userId {
			iffound = true
			break

		}
	}
	return iffound
}
