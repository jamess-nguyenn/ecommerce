package connection

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type MysqlDatabase struct {
	Master *gorm.DB
	Slave  *gorm.DB
}

func buildDsnMaster() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_MYSQL_WRITE_USERNAME"),
		os.Getenv("DB_MYSQL_WRITE_PASSWORD"),
		os.Getenv("DB_MYSQL_WRITE_HOST"),
		os.Getenv("DB_MYSQL_WRITE_PORT"),
		os.Getenv("DB_MYSQL_DATABASE"),
	)

	return dsn
}

func buildDsnSlave() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_MYSQL_READ_USERNAME"),
		os.Getenv("DB_MYSQL_READ_PASSWORD"),
		os.Getenv("DB_MYSQL_READ_HOST"),
		os.Getenv("DB_MYSQL_READ_PORT"),
		os.Getenv("DB_MYSQL_DATABASE"),
	)

	return dsn
}

func openConnectionMysql(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	return db, nil
}

func getInstanceMysql(db *gorm.DB) (*sql.DB, error) {
	sqlDb, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get the database instance: %w", err)
	}

	return sqlDb, nil
}

func pingConnectionMysql(db *gorm.DB) error {
	sqlDb, err := getInstanceMysql(db)
	if err != nil {
		return err
	}

	if err = sqlDb.Ping(); err != nil {
		return fmt.Errorf("failed to ping the database: %w", err)
	}

	return nil
}

func closeConnectionMysql(db *gorm.DB) error {
	sqlDb, err := getInstanceMysql(db)
	if err != nil {
		return err
	}

	if err = sqlDb.Close(); err != nil {
		return fmt.Errorf("failed to close the database: %w", err)
	}

	return nil
}

func ConnectMysql() (*MysqlDatabase, error) {
	// open a connection to the master database
	masterDb, err := openConnectionMysql(buildDsnMaster())
	if err != nil {
		return nil, err
	}

	// ping the master database to check the connection
	if err = pingConnectionMysql(masterDb); err != nil {
		return nil, err
	}

	// open a connection to the slave database
	slaveDb, err := openConnectionMysql(buildDsnSlave())
	if err != nil {
		return nil, err
	}

	// ping the slave database to check the connection
	if err = pingConnectionMysql(slaveDb); err != nil {
		return nil, err
	}

	return &MysqlDatabase{
		Master: masterDb,
		Slave:  slaveDb,
	}, nil
}

func (db *MysqlDatabase) Close() error {
	// close a connection to the master database
	errMaster := closeConnectionMysql(db.Master)

	// close a connection to the slave database
	errSlave := closeConnectionMysql(db.Slave)

	// check errors
	if errMaster != nil || errSlave != nil {
		return fmt.Errorf("errors occurred while closing the databases. Master: %v, Slave: %v", errMaster, errSlave)
	}

	return nil
}
