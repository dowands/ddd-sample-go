package infrastructure

import (
	"go-ddd/internal/domain/customer"
	"go-ddd/internal/infrastructure/database"
	. "go-ddd/internal/infrastructure/database/dataobject"
)

type CustomerGateway struct {
}

func (g CustomerGateway) CountByIdEmployee(idEmployee int64) int {
	return database.CountByIdEmployee(idEmployee)
}
func (g CustomerGateway) Add(aggr *customer.CustomerAggregate) int64 {
	customerDO := &CusomterDO{
		IdEmployee: aggr.EmployeeVO.Id,
		Name:       aggr.Name,
	}
	return database.AddCustomer(customerDO)
}
func (g CustomerGateway) Update(aggr *customer.CustomerAggregate) int {
	customerDO := &CusomterDO{
		Id: aggr.Id,
		IdEmployee: aggr.EmployeeVO.Id,
		Name:       aggr.Name,
	}
	return int(database.UpdateCustomer(customerDO))
}

func ListCustomer(start int, limit int) []*CusomterDO {
	return database.ListCustomer(start, limit)
}

func TotalCustomer() int64 {
	return database.TotalCustomer()
}
