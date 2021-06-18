package conn

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type conn struct {
	Type string
	DataSource string
}

var ConnConfig = conn{
	Type:       "sqlite3",
	DataSource: "file::memory:?mode=memory&cache=shared&loc=auto",
}

var SqlConn, err = sql.Open(ConnConfig.Type, ConnConfig.DataSource)

func init()  {
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
}

