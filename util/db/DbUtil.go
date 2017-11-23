package db

import (
	"sync"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/koding/multiconfig"
	"github.com/bragfoo/saman/util/config"
	"bytes"
	"github.com/siddontang/go/log"
)

type Db struct {
	Handle *sql.DB
}

var instance *Db
var once sync.Once

func GetInstance() *Db {
	once.Do(func() {
		d := Db{}
		dbConf := initDatabaseConfig()
		db, err := sql.Open("mysql", getDataSource(dbConf))
		if nil != err {
			log.Error(err)
		} else {
			db.SetMaxOpenConns(dbConf.MaxConnections)
			db.SetMaxIdleConns(dbConf.MaxConnections)
			db.SetMaxIdleConns(dbConf.MaxIdleConns)
			d.Handle = db
		}
		instance = &d
	})
	return instance
}

func getPasswd(p string) string {
	return p
}

func getDataSource(dbConf *config.DbType) string {
	var buffer bytes.Buffer
	//  username:password@protocol(address:port)/dbname
	buffer.WriteString(dbConf.Username)
	buffer.WriteString(":")
	buffer.WriteString(getPasswd(dbConf.Password))
	buffer.WriteString("@")
	buffer.WriteString(dbConf.Protocol)
	buffer.WriteString("(")
	buffer.WriteString(dbConf.Address)
	buffer.WriteString(":")
	buffer.WriteString(dbConf.Port)
	buffer.WriteString(")/")
	buffer.WriteString(dbConf.Dbname)
	return buffer.String()
}

func initDatabaseConfig() *config.DbType {
	confPath := "conf/config.toml"
	m := multiconfig.NewWithPath(confPath)
	conf := &config.ConfType{}
	m.MustLoad(conf)
	return conf.Db
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	pool := GetInstance()
	errPing := pool.Handle.Ping()
	if nil != errPing {
		log.Error(errPing)
		return nil, errPing
	} else {
		return pool.Handle.Query(query, args...)
	}
}

func QueryRow(query string, args interface{}) (*sql.Row, error) {
	pool := GetInstance()
	errPing := pool.Handle.Ping()
	if nil != errPing {
		log.Error(errPing)
		return nil, errPing
	} else {
		return pool.Handle.QueryRow(query, args), nil
	}
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	pool := GetInstance()
	errPing := pool.Handle.Ping()
	if nil != errPing {
		log.Error(errPing)
		return nil, errPing
	} else {
		return pool.Handle.Exec(query, args)
	}
}

func Prepare(query string) (*sql.Stmt, error) {
	pool := GetInstance()
	errPing := pool.Handle.Ping()
	if nil != errPing {
		log.Error("Database connection has gone dead!please check your database")
		log.Error(errPing)
		return nil, errPing
	} else {
		return pool.Handle.Prepare(query)
	}
}
