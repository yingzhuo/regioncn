package main

import (
	"database/sql"
	"fmt"
	"os"

	"region-cn/cnf"
	"region-cn/web"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/subchen/go-cli"
)

var config = &cnf.Config{}

var db *sql.DB

func main() {

	app := cli.NewApp()
	app.Name = "regioncn"
	app.Usage = "a http server of id-generator"
	app.UsageText = "[options]"
	app.Authors = "应卓 <yingzhor@gmail.com>"
	app.Version = "1.0.0"
	app.BuildInfo = &cli.BuildInfo{}
	app.Flags = []*cli.Flag{
		{
			Name:          "tz, timezone",
			Usage:         "timezone of application",
			DefValue:      "Asia/Shanghai",
			NoOptDefValue: "Asia/Shanghai",
			Value:         &config.Timezone,
		}, {
			Name:          "port",
			Usage:         "port of application",
			DefValue:      "8080",
			NoOptDefValue: "8080",
			Value:         &config.HttpPort,
		}, {
			Name:          "mysql-host",
			Usage:         "host of msql",
			DefValue:      "local",
			NoOptDefValue: "local",
			Value:         &config.MySqlHost,
		}, {
			Name:          "mysql-port",
			Usage:         "port of msql",
			DefValue:      "3306",
			NoOptDefValue: "3306",
			Value:         &config.MySqlPort,
		}, {
			Name:          "mysql-database",
			Usage:         "database of msql",
			DefValue:      "regioncn",
			NoOptDefValue: "regioncn",
			Value:         &config.MySqlDatabase,
		}, {
			Name:          "mysql-username",
			Usage:         "username of msql",
			DefValue:      "regioncn",
			NoOptDefValue: "regioncn",
			Value:         &config.MySqlUsername,
		}, {
			Name:          "mysql-password",
			Usage:         "password of msql",
			DefValue:      "regioncn",
			NoOptDefValue: "regioncn",
			Value:         &config.MySqlPassword,
		}, {
			Name:          "type",
			Usage:         "response type of http",
			DefValue:      "protobuf",
			NoOptDefValue: "protobuf",
			Value:         &config.ResponseType,
		}, {
			Name:          "i, indent",
			Usage:         "output indented json",
			DefValue:      "false",
			NoOptDefValue: "false",
			Hidden:        true,
			Value:         &config.ResponseType,
		},
	}

	app.Action = func(context *cli.Context) {

		logrus.SetOutput(os.Stdout)
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableColors: true,
			FullTimestamp: false,
		})

		// 设置时区
		_ = os.Setenv("TZ", config.Timezone)

		logrus.Debugf("pid            = %v", os.Getpid())
		logrus.Debugf("timezone       = %v", config.Timezone)
		logrus.Debugf("web port       = %v", config.HttpPort)
		logrus.Debugf("mysql host     = %v", config.MySqlHost)
		logrus.Debugf("mysql port     = %v", config.MySqlPort)
		logrus.Debugf("mysql username = %v", config.MySqlUsername)
		logrus.Debugf("mysql password = %v", config.MySqlPassword)
		logrus.Debugf("mysql database = %v", config.MySqlDatabase)
		logrus.Debugf("response type  = %v", config.ResponseType)

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

	app.Run(os.Args)
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
