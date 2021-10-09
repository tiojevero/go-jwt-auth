package data

import (
	"fmt"

	"xorm.io/xorm"
)


func createDBEngine() (*xorm.Engine, error){
	connectionInfo := fmt.Sprintf("host=%s port=%d password=%s dbname=%s sslmode=disable", "localhost", 5432, "root", "password")
	engine, err := xorm.NewEngine("postgres", connectionInfo)
	
	if err != nil {
		return nil, err
	}
	
	if err := engine.Ping(); err != nil {
		return nil, err
	} 

	return engine, nil
}	