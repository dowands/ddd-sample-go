package database

import (
	conn2 "go-ddd/config/conn"
	"go-ddd/config/sql_defined"
	. "go-ddd/internal/infrastructure/database/dataobject"
)

func ListCustomer(start int, limit int) (result []*CusomterDO) {
	//使用stmt降低sql注入风险
	stmt, _ := conn2.SqlConn.Prepare(sql_defined.Sqls["listCusomter"])
	defer stmt.Close()

	rows, _ := stmt.Query(start, limit)
	if rows == nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id, idEmployee int64
		var name string
		_ = rows.Scan(&id, &idEmployee, &name)
		result = append(result, &CusomterDO{
			id, idEmployee, name,
		})
	}

	return
}

func TotalCustomer() (result int64) {
	stmt, _ := conn2.SqlConn.Prepare(sql_defined.Sqls["totalCusomter"])
	defer stmt.Close()

	rows, _ := stmt.Query()
	defer rows.Close()

	_ = rows.Scan(&result)

	return
}

func CountByIdEmployee(idEmployee int64) (result int) {
	stmt, _ := conn2.SqlConn.Prepare(sql_defined.Sqls["countByIdEmployee"])
	defer stmt.Close()

	rows, _ := stmt.Query(idEmployee)
	defer rows.Close()

	_ = rows.Scan(&result)

	return
}

func AddCustomer(do *CusomterDO) (result int64) {
	stmt, _ := conn2.SqlConn.Prepare(sql_defined.Sqls["addCustomer"])
	defer stmt.Close()

	rows, _ := stmt.Exec(do.IdEmployee, do.Name)
	id, _ := rows.LastInsertId()
	return id
}

func UpdateCustomer(do *CusomterDO) int64 {
	stmt, _ := conn2.SqlConn.Prepare(sql_defined.Sqls["updateCustomer"])
	defer stmt.Close()

	rows, _ := stmt.Exec(do.IdEmployee, do.Name, do.Id)
	numbers, _ := rows.RowsAffected()
	return numbers
}