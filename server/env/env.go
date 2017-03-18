package env

import (
	"github.com/gocql/gocql"
	"github.com/bitly/go-nsq"
	"database/sql"
)

type Env struct {
	CqlSession *gocql.Session
	NsqProducer *nsq.Producer
	Postgres *sql.DB
}
