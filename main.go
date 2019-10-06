package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"os"
	"region-cn/cnf"
	"region-cn/web"
)

var config = &cnf.Config{}

var db *sql.DB

const (
	version   = "1.0.0"
	goVersion = "go1.13.1"
)

func main() {

	if config.Version {
		fmt.Println("Version   :", version)
		fmt.Println("Go version:", goVersion)
		return
	}

	db = NewDB()
	defer func() {
		if db != nil {
			_ = db.Close()
		}
	}()

	if err := db.Ping(); err != nil {
		logrus.Error("db connection is NOT reachable")
		os.Exit(1)
	}

	web.StartHttpServer(db, config)
}

func init() {
	cmd := flag.NewFlagSet("region-cn", flag.ExitOnError)
	cmd.StringVar(&config.Timezone, "timezone", "Asia/Shanghai", "timezone")
	cmd.StringVar(&config.HttpHost, "web-host", "0.0.0.0", "web host")
	cmd.IntVar(&config.HttpPort, "web-port", 8080, "web port")
	cmd.StringVar(&config.MySqlHost, "mysql-host", "localhost", "mysql host")
	cmd.IntVar(&config.MySqlPort, "mysql-port", 3306, "mysql port")
	cmd.StringVar(&config.MySqlUsername, "mysql-username", "regioncn", "mysql username")
	cmd.StringVar(&config.MySqlPassword, "mysql-password", "regioncn", "mysql password")
	cmd.StringVar(&config.MySqlDatabase, "mysql-database", "regioncn", "mysql password")
	cmd.StringVar(&config.ResponseType, "response-type", cnf.Protobuf, "mysql password")
	cmd.BoolVar(&config.IndentJson, "indent-json", false, "intent json")
	cmd.BoolVar(&config.Version, "version", false, "show version")
	_ = cmd.Parse(os.Args[1:])

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: false,
	})

	// 设置时区
	_ = os.Setenv("TZ", config.Timezone)

	logrus.Debugf("timezone       = %v", config.Timezone)
	logrus.Debugf("web host       = %v", config.HttpHost)
	logrus.Debugf("web port       = %v", config.HttpPort)
	logrus.Debugf("mysql host     = %v", config.MySqlHost)
	logrus.Debugf("mysql port     = %v", config.MySqlPort)
	logrus.Debugf("mysql username = %v", config.MySqlUsername)
	logrus.Debugf("mysql password = %v", config.MySqlPassword)
	logrus.Debugf("mysql database = %v", config.MySqlDatabase)
	logrus.Debugf("response type  = %v", config.ResponseType)

}

func NewDB() *sql.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
		config.MySqlUsername,
		config.MySqlPassword,
		config.MySqlHost,
		config.MySqlPort,
		config.MySqlDatabase,
	)

	logrus.Debugf("dsn = %v", dsn)

	if database, err := sql.Open("mysql", dsn); err != nil {
		panic(err)
	} else {
		return database
	}
}
