package customer

type EmployeeVO struct {
	Id int64
	AllowingAssign bool
	Facade Facade
}

func (this *EmployeeVO) IsAllowingAssign() bool {
	if !this.AllowingAssign {
		return false
	}

	count := this.Facade.CountByIdEmployee(this.Id)
	return count <= 3
}

