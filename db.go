package main

import (
	"fmt"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type User struct {
	Id int64
	Name string
	Email string
	Password string
}


func createDBEngine() (*xorm.Engine, error){
	connectionInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "root", "go-jwt-auth")
	engine, err := xorm.NewEngine("postgres", connectionInfo)
	
	if err != nil {
		return nil, err
	}
	
	if err := engine.Ping(); err != nil {
		return nil, err
	} 

	if err := engine.Sync(new(User)); err != nil {
		return nil, err
	}

	return engine, nil
}	