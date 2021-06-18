package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	conn2 "go-ddd/config/conn"
	"go-ddd/internal/adapter/controller"
	"net/http"
)

func main() {
	dbinit()

	gin.SetMode(gin.ReleaseMode)
	if err := http.ListenAndServe(":9090", controller.Route()); err != nil {
		fmt.Printf("%v", err)
	} else {
		fmt.Print("stop server success")
	}
}

func dbinit(){
	_, _ = conn2.SqlConn.Exec("CREATE TABLE customer (id bigint auto_increment NULL,id_employee bigint DEFAULT NULL,name varchar(255) DEFAULT NULL,PRIMARY KEY (id));")
	_, _ = conn2.SqlConn.Exec("CREATE TABLE employee (id bigint auto_increment NULL,allowingAssign tinyint DEFAULT 1,PRIMARY KEY (id));")
	_, _ = conn2.SqlConn.Exec("insert into customer (id_employee, name) values (6, 'test 1');")
	_, _ = conn2.SqlConn.Exec("insert into customer (id_employee, name) values (6, 'test 2');")
	_, _ = conn2.SqlConn.Exec("insert into customer (id_employee, name) values (7, 'test 3');")
}