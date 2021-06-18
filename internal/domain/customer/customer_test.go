package customer

import (
	"go-ddd/internal/client/expection"
	"testing"
)

type TestCustomerGateway struct {

}

func (g TestCustomerGateway) CountByIdEmployee(idEmployee int64) int {
	return 4
}
func (g TestCustomerGateway) Add(aggr *CustomerAggregate) int64 {
	return 123
}
func (g TestCustomerGateway) Update(aggr *CustomerAggregate) int {
	return 1
}

var facadeTest = TestCustomerGateway{}

func TestCustomer(t *testing.T)  {
	aggregate := CustomerAggregate{
		EmployeeVO: EmployeeVO{
			Id:     6,
			Facade: facadeTest,
		},
		Name:   "test name",
		Facade: facadeTest,
	}
	
	if _, err := aggregate.Update(); err != expection.Expections.CAN_NOT_ASSIGN {
		t.Error("测试失败")
	}
}
