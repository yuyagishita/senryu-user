package db

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/yu-yagishita/nanpa-user/users"
)

// Database represents a simple interface so we can switch to a new system easily
// this is just basic and specific to this microservice
type Database interface {
	Init() error
	GetUserByName(string) (users.User, error)
	GetUser(string) (users.User, error)
	GetUsers() ([]users.User, error)
	CreateUser(*users.User) error
	GetUserAttributes(*users.User) error
	GetAddress(string) (users.Address, error)
	GetAddresses() ([]users.Address, error)
	CreateAddress(*users.Address, string) error
	GetCard(string) (users.Card, error)
	GetCards() ([]users.Card, error)
	Delete(string, string) error
	CreateCard(*users.Card, string) error
	Ping() error
}

var (
	database string
	//DefaultDb is the database set for the microservice
	DefaultDb Database
	//DBTypes is a map of DB interfaces that can be used for this service
	DBTypes = map[string]Database{}
	//ErrNoDatabaseFound error returnes when database interface does not exists in DBTypes
	ErrNoDatabaseFound = "No database with name %v registered"
	//ErrNoDatabaseSelected is returned when no database was designated in the flag or env
	ErrNoDatabaseSelected = errors.New("No DB selected")
)

func init() {
	flag.StringVar(&database, "database", os.Getenv("USER_DATABASE"), "Database to use, Mongodb or ...")

}

//Init inits the selected DB in DefaultDb
func Init() error {
	if database == "" {
		return ErrNoDatabaseSelected
	}
	err := Set()
	if err != nil {
		return err
	}
	return DefaultDb.Init()
}

//Set the DefaultDb
func Set() error {
	if v, ok := DBTypes[database]; ok {
		DefaultDb = v
		return nil
	}
	return fmt.Errorf(ErrNoDatabaseFound, database)
}

//Register registers the database interface in the DBTypes
func Register(name string, db Database) {
	DBTypes[name] = db
}

//CreateUser invokes DefaultDb method
func CreateUser(u *users.User) error {
	return DefaultDb.CreateUser(u)
}

//GetUserByName invokes DefaultDb method
func GetUserByName(n string) (users.User, error) {
	u, err := DefaultDb.GetUserByName(n)
	if err == nil {
		u.AddLinks()
	}
	return u, err
}

//GetUser invokes DefaultDb method
func GetUser(n string) (users.User, error) {
	u, err := DefaultDb.GetUser(n)
	if err == nil {
		u.AddLinks()
	}
	return u, err
}

//GetUsers invokes DefaultDb method
func GetUsers() ([]users.User, error) {
	us, err := DefaultDb.GetUsers()
	for k, _ := range us {
		us[k].AddLinks()
	}
	return us, err
}

//Delete invokes DefaultDb method
func Delete(entity, id string) error {
	return DefaultDb.Delete(entity, id)
}

//Ping invokes DefaultDB method
func Ping() error {
	return DefaultDb.Ping()
}
