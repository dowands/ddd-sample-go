package customer

import (
	. "go-ddd/internal/client/expection"
	"go-ddd/pkg/expection"
)

type CustomerAggregate struct {
	Id int64
	EmployeeVO EmployeeVO
	Name string
	Facade Facade
}

func (this *CustomerAggregate) Update() (bool, expection.Exception) {
	if !this.EmployeeVO.IsAllowingAssign() {
		return false, Expections.CAN_NOT_ASSIGN
	}

	if this.Id == 0 {
		this.Id = this.Facade.Add(this)
	}else{
		return this.Facade.Update(this) > 0, Expections.EMPTY
	}

	return this.Id > 0, Expections.EMPTY
}
