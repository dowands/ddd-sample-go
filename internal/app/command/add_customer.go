package command

import (
	. "go-ddd/internal/client/request_object"
	"go-ddd/internal/domain"
	"go-ddd/pkg/expection"
)

func AddCustomer(cusomterRO CusomterRO) expection.Exception {
	return domain.CustomerDomainService{}.Create(cusomterRO)
}
