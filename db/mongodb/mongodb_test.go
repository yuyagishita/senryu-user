package mongodb

import (
	"os"
	"testing"

	"github.com/yu-yagishita/senryu-user/users"
	"gopkg.in/mgo.v2/dbtest"
)

var (
	TestMongo  = Mongo{}
	TestServer = dbtest.DBServer{}
	TestUser   = users.User{
		Username: "username",
		Password: "blahblah",
		Email:    "username@yagi.com",
	}
)

func init() {
	TestServer.SetPath("/tmp")
}

func TestMain(m *testing.M) {
	TestMongo.Session = TestServer.Session()
	// TestMongo.EnsureIndexes()
	// TestMongo.Session.Close()
	exitTest(0)
}

func exitTest(i int) {
	// TestServer.Wipe()
	// TestServer.Stop()
	os.Exit(i)
}

// func TestInit(t *testing.T) {
// 	err := TestMongo.Init()
// 	if err.Error() != "no reachable servers" {
// 		t.Error("expecting no reachable servers error")
// 	}
// }
