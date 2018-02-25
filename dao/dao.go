package dao
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type ConnInfo struct{
	Driver string
	DB string
	User string
	Pwd string
	Host string
	Port string
	Charset string
	Table string
}

type Dao struct{
	_conn *sql.DB
	_table string
	_where string
}


func Create(info ConnInfo) (*Dao,error){
	driver := "mysql"
	if info.Driver != ""{
		driver = info.Driver
	}
	charset := "utf8"
	if info.Charset != ""{
		charset = info.Charset
	}
	port := "3306"
	if info.Port != ""{
		port = info.Port
	}
	host := "localhost"
	if info.Host != ""{
		host = info.Host
	}
	user := info.User
	pwd := info.Pwd
	db := info.DB

	sentence := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",user,pwd,host,port,db,charset)//"root:333221@tcp(localhost:3306)/book?charset=utf8"
	conn,err := sql.Open(driver, sentence)
	dao := &Dao{_conn:conn,_table:info.Table}
	return dao,err
}

func (dao Dao) Table(table string){
	dao._table=table
}

func (dao Dao) Where(where string){
	dao._where = where
}

func query(){

}

func exec(){

}