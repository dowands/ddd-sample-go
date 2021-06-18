package query

import (
	_const "go-ddd/config/const"
	. "go-ddd/internal/client/client_object"
)
import "go-ddd/internal/infrastructure"

func ListCustomer(page int) (result []*CusomterCO) {
	list := infrastructure.ListCustomer((page-1) * _const.PERPAGE, _const.PERPAGE)

	for _, item := range list {
		result = append(result, &CusomterCO{
			item.Id, item.IdEmployee, item.Name,
		})
	}

	return
}

func TotalCustomer() int64 {
	return infrastructure.TotalCustomer()
}
