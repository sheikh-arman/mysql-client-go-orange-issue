package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"k8s.io/klog/v2"
	_ "k8s.io/klog/v2"
	"log"
	"os"
	"time"
)

func main() {
	for true {
		klog.Infoln("Querying..")
		dbHost := os.Getenv("DB_HOST")
		if dbHost == "" {
			dbHost = "mmr-lab.mysql-mmr-lab-usjzu.svc:3306"
		}
		dbName := os.Getenv("DB_NAME")
		if dbName == "" {
			dbName = "orange"
		}
		dbPass := os.Getenv("MYSQL_ROOT_PASSWORD")
		if dbPass == "" {
			dbPass = "fM97uAv)nY)TNqTU"
		}
		in := Appscode{
			dbHost:     dbHost,
			dbName:     dbName,
			dbPassword: dbPass,
		}
		db := in.openConnection()
		_, err := db.Query("select * from orange.sample_table")
		if err != nil {
			klog.Errorf("Error querying orange.sample_table %v", err)
		}
		fmt.Println("")
		time.Sleep(time.Second * 30)
	}
}

type Appscode struct {
	dbHost     string `json:"dbhost"`
	dbName     string `json:"dbname"`
	dbPassword string `json:"dbpassword"`
}

func (in Appscode) openConnection() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s", "root", in.dbPassword, in.dbHost, in.dbName))
	if err != nil {
		klog.Errorf("Error opening database %s\n", err.Error())
	}
	return db
}

func (in Appscode) closeConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalf("closing connection %s\n", err.Error())
	}
}
