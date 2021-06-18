package customer

type Facade interface {
	CountByIdEmployee (idEmployee int64) int
	Add(this *CustomerAggregate) int64
	Update(this *CustomerAggregate) int
}
