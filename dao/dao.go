package dao
type Dao struct{
	_db string
	_user string
	_pwd string
	table string
}

func Create() *Dao{
	return new(Dao)
}

func (dao Dao) db(dbname string){
	dao._db = dbname
}

func (dao Dao) user(username string){
	dao._user = username
}

func (dao Dao) pwd(password string){
	dao._pwd = password
}