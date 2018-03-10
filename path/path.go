package path
import (
	"os"
)
func Exist(path string) bool{
	_,err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func IsDir(path string) bool{
	f, err := os.Stat(path)
	if err == nil{
		if f.IsDir(){
			return true
		}else{
			return false
		}
	}
	return false
}
