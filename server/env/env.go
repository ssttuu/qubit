package env

import "github.com/gocql/gocql"

type Env struct {
	CqlSession *gocql.Session
}
