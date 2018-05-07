package apmmysql

import (
	"github.com/go-sql-driver/mysql"

	"github.com/elastic/apm-agent-go/contrib/apmsql"
)

func init() {
	apmsql.Register("mysql", &mysql.MySQLDriver{}, apmsql.WithDSNParser(ParseDSN))
}
