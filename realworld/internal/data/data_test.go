package data

import (
	"realworld/internal/conf"
	"testing"
)

func TestNewDB(t *testing.T) {
	//Mock or set up your conf.Data with valid database connection details
	mockConf := &conf.Data{
		Database: &conf.Data_Database{
			Dsn: "root:123456test@tcp(127.0.0.1:3309)/realworld?charset=utf8mb4&parseTime=True&loc=Local",
		},
	}

	db := NewDB(mockConf)
	if db == nil {
		t.Errorf("Expected non-nil DB")
	}
}
