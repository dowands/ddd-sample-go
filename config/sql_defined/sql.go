package sql_defined

var Sqls = map[string]string{
	"listCusomter": "select * from customer limit ?, ?",
	"totalCusomter": "select count(*) from customer",
	"countByIdEmployee": "select count(*) from customer where id_employee = ?",
	"addCustomer": "insert into customer (id_employee, name) values (?, ?)",
	"updateCustomer": "update customer set id_employee = ?, name = ? where id = ?",
}
