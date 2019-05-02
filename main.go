package main

import (
	"github.com/labstack/echo"
	"time"

	"gopkg.in/mgo.v2"
	r "blog/user/repository"
	u "blog/user/usecase"
	h "blog/user/delivery/http"
)

const (
	hosts = "localhost:27017"
	// database   = "test_db"
	// username   = "user1"
	// password   = "example"
	collection = "blog"
)			

func main() {
	// Tạo phiên kết nối với MongDB
	info := &mgo.DialInfo{
		Addrs:   []string{hosts},
		Timeout: 60 * time.Second,
		// Database: database,
		// Username: username,
		// Password: password,
	}
	session, err := mgo.DialWithInfo(info)
	defer session.Close()
	if err != nil {
		panic(err)
	}
	timeoutContext := time.Duration(5) * time.Second
	userRepo := r.NewMongoUserRepository(session)
	uu := u.NewUserUsecase(userRepo, timeoutContext)
	e:= echo.New()
	h.NewUserHTTPHandler(e, uu)
	e.Logger.Fatal(e.Start(":1323"))
}
