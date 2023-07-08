package main

import (
	"fmt"
	"github.com/pkg/errors"
	"go-gorm-gen-example/model"
	"go-gorm-gen-example/query"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	gormdb, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(errors.Wrap(err, "failed to connect database"))
	}
	query.SetDefault(gormdb)
	adminUser := model.User{Username: "modi", Age: 18, Role: "admin"}
	newUser := model.User{Username: "modi2", Age: 18, Role: "user"}
	if err := query.User.Create(&adminUser, &newUser); err != nil {
		panic(errors.Wrap(err, "failed to create user"))
	}

	u := query.User
	user, err := query.User.Where(u.Username.Eq("modi")).First()
	if err != nil {
		panic(errors.Wrap(err, "failed to query user"))
	}
	fmt.Println(user)

	users, err := query.User.FilterWithUsernameAndRole("modi", "admin")
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		fmt.Println(user)
	}

}
