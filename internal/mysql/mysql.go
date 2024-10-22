package mysql

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var rdsUser string

func init() {
	rdsUser = os.Getenv("RDS_USER")
}

// IntializeDatabase 데이타베이스 초기생성
func IntializeDatabase(dbName string) (*gorm.DB, *gorm.DB, error) {
	masterDb, masterErr := initMasterDB(dbName)
	if masterErr != nil {
		return masterDb, nil, masterErr
	}

	replicaDb, repleErr := initReplicaDB(dbName)
	if repleErr != nil {
		return nil, nil, repleErr
	}

	return masterDb, replicaDb, nil
}

// Ping 디비연결 테스트
func Ping(dbName string) error {
	mysqlMasterConnStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("RDS_USER"), os.Getenv("RDS_PASSWORD"), os.Getenv("RDS_MASTER_HOST"), os.Getenv("RDS_PORT"), dbName)
	masterDb, _ := gorm.Open("mysql", mysqlMasterConnStr)
	masterDb.DB()
	defer masterDb.Close()
	err := masterDb.DB().Ping()
	if err != nil {
		return err
	}

	mysqlReplicaConnStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("RDS_USER"), os.Getenv("RDS_PASSWORD"), os.Getenv("RDS_REPLICA_HOST"), os.Getenv("RDS_PORT"), dbName)
	replicaDb, _ := gorm.Open("mysql", mysqlReplicaConnStr)

	replicaDb.DB()
	defer replicaDb.Close()
	err = replicaDb.DB().Ping()
	if err != nil {
		return err
	}
	return nil
}

func initMasterDB(dbName string) (*gorm.DB, error) {
	mysqlMasterConnStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("RDS_USER"), os.Getenv("RDS_PASSWORD"), os.Getenv("RDS_MASTER_HOST"), os.Getenv("RDS_PORT"), dbName)
	masterDb, err := gorm.Open("mysql", mysqlMasterConnStr)
	if err != nil {
		return nil, err
	}
	masterDb.DB()
	err = masterDb.DB().Ping()
	if err != nil {
		return nil, err
	}

	masterDb.LogMode(true)
	masterDb.DB().SetMaxIdleConns(5)
	masterDb.DB().SetMaxOpenConns(5)
	masterDb.DB().SetConnMaxLifetime(time.Minute * 5)
	masterDb.SingularTable(true)
	return masterDb, nil
}

func initReplicaDB(dbName string) (*gorm.DB, error) {
	mysqlReplicaConnStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("RDS_USER"), os.Getenv("RDS_PASSWORD"), os.Getenv("RDS_REPLICA_HOST"), os.Getenv("RDS_PORT"), dbName)
	replicaDb, _ := gorm.Open("mysql", mysqlReplicaConnStr)

	replicaDb.DB()
	err := replicaDb.DB().Ping()
	if err != nil {
		return nil, err
	}

	replicaDb.LogMode(true)
	replicaDb.DB().SetMaxIdleConns(5)
	replicaDb.DB().SetMaxOpenConns(5)
	replicaDb.DB().SetConnMaxLifetime(time.Minute * 5)
	replicaDb.SingularTable(true)
	return replicaDb, nil
}
