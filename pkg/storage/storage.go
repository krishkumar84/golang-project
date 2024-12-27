package storage
import "github.com/krishkumar84/golang-project/pkg/types"

type Storage interface {
	CreateUser(name string, email string, age int)(int64, error)
	GetUserById(id int64)(types.User, error) 
	GetAllUsers()([]types.User, error)
}
