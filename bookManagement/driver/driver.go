package driver

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	vipers "github.com/mayurkhairnar2525/bookManagement/viper"
	log "github.com/sirupsen/logrus"
)

type DataStore struct {
	Db *sqlx.DB
}

func ConnectDB() (*DataStore, error) {
	dbConf, err := vipers.GetDbconfigs()
	if err != nil {
		log.Fatal(err)
	}
	dbInstance := &DataStore{}
	dbInstance.Db, err = sqlx.Connect(dbConf.Drivername, dbConf.Username+":"+dbConf.Password+"@tcp("+dbConf.Host+":"+dbConf.Port+")/"+dbConf.DbName)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Println("Database : connected successfully")
	return dbInstance, nil
}