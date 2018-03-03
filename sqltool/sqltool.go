package sqltool
import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	)

func CreateConnPool(host string,user string,pwd string,db string ,port string,charset string) (*sql.DB,error){
	return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",user,pwd,host,port,db,charset))
}
