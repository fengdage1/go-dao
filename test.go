package main
import (
	"fmt"
	"./dao"
)
type User struct {
	     id   int
	     name string
	     age  int
	 }
func main(){
	var t struct{name string;user string}
	t.name="1"
	t.user="l"

	u:=User{ name:"ok", age:12}
	fmt.Println(u)
	d,err := dao.Create(dao.ConnInfo{User:"root",Pwd:"333221abc",DB:"book",Host:"127.0.0.1",Table:"list"})
	if err!=nil{
		fmt.Println("err")
	}
	fmt.Println(d)

}
