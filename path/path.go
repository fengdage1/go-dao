package path
import (
	"os"
)
func Exist(path string) (bool,error){
	_,err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func IsDir(path string) (bool,error){
	f, err := os.Stat(path)
	if err == nil{
		if f.IsDir(){
			return true,nil
		}else{
			return false,nil
		}
	}
	return false,err
}
