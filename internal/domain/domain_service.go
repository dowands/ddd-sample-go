package domain

import (
	"go-ddd/internal/client/request_object"
	"go-ddd/internal/domain/customer"
	"go-ddd/internal/infrastructure"
	"go-ddd/pkg/expection"
)

type CustomerDomainService struct {

}

var facade customer.Facade = infrastructure.CustomerGateway{}

func (d CustomerDomainService) Create(ro request_object.CusomterRO) (err expection.Exception) {
	aggrgate := customer.CustomerAggregate{
		EmployeeVO: customer.EmployeeVO{
			Id:     ro.IdEmployee,
			Facade: facade,
		},
		Name:   ro.Name,
		Facade: facade,
	}
	_, err = aggrgate.Update()
	return
}
