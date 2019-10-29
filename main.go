package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"region-cn/cnf"
	"region-cn/web"

	_ "github.com/go-sql-driver/mysql"
	"github.com/subchen/go-cli"
)

var config = &cnf.Config{}

var db *sql.DB

// build info
var (
	BuildVersion   string
	BuildGitBranch string
	BuildGitRev    string
	BuildGitCommit string
	BuildDate      string
)

func main() {

	app := cli.NewApp()
	app.Name = "regioncn"
	app.Usage = "a http server of region info"
	app.UsageText = "[options]"
	app.Authors = "应卓 <yingzhor@gmail.com>"
	app.Version = BuildVersion
	app.BuildInfo = &cli.BuildInfo{
		GitBranch:   BuildGitBranch,
		GitCommit:   BuildGitCommit,
		GitRevCount: BuildGitRev,
		Timestamp:   BuildDate,
	}

	app.SeeAlso = `https://github.com/yingzhuo/regioncn
https://github.com/yingzhuo/regioncn-mysql
https://github.com/yingzhuo/regioncn-golang-client
https://github.com/yingzhuo/regioncn-java-client`

	app.Flags = []*cli.Flag{
		{
			Name:     "tz, timezone",
			Usage:    "timezone of application",
			DefValue: "Asia/Shanghai",
			Value:    &config.Timezone,
		}, {
			Name:     "port",
			Usage:    "port of application",
			DefValue: "8080",
			Value:    &config.HttpPort,
		}, {
			Name:     "mysql-host",
			Usage:    "host of msql",
			DefValue: "local",
			Value:    &config.MySqlHost,
		}, {
			Name:     "mysql-port",
			Usage:    "port of msql",
			DefValue: "3306",
			Value:    &config.MySqlPort,
		}, {
			Name:     "mysql-database",
			Usage:    "database of msql",
			DefValue: "regioncn",
			Value:    &config.MySqlDatabase,
		}, {
			Name:     "mysql-username",
			Usage:    "username of msql",
			DefValue: "regioncn",
			Value:    &config.MySqlUsername,
		}, {
			Name:     "mysql-password",
			Usage:    "password of msql",
			DefValue: "regioncn",
			Value:    &config.MySqlPassword,
		}, {
			Name:     "type",
			Usage:    "response type of http",
			DefValue: "protobuf",
			Value:    &config.ResponseType,
		}, {
			Name:          "i, indent",
			Usage:         "output indented json",
			DefValue:      "false",
			NoOptDefValue: "false",
			IsBool:        true,
			Value:         &config.ResponseType,
		},
	}

	app.OnActionPanic = func(context *cli.Context, e error) {
		msg := e.Error()
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}
		_, _ = os.Stderr.WriteString(msg)
		os.Exit(1)
	}

	app.Action = func(context *cli.Context) {
		// 设置时区
		_ = os.Setenv("TZ", config.Timezone)

		fmt.Printf("pid            = %v\n", os.Getpid())
		fmt.Printf("timezone       = %v\n", config.Timezone)
		fmt.Printf("web port       = %v\n", config.HttpPort)
		fmt.Printf("mysql host     = %v\n", config.MySqlHost)
		fmt.Printf("mysql port     = %v\n", config.MySqlPort)
		fmt.Printf("mysql username = %v\n", config.MySqlUsername)
		fmt.Printf("mysql password = %v\n", config.MySqlPassword)
		fmt.Printf("mysql database = %v\n", config.MySqlDatabase)
		fmt.Printf("response type  = %v\n", config.ResponseType)
		if strings.EqualFold(config.ResponseType, "json") {
			fmt.Printf("indent json    = %v", config.Indent)
		}

		db = NewDB()
		defer func() {
			if db != nil {
				_ = db.Close()
			}
		}()

		if err := db.Ping(); err != nil {
			panic(err)
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

	if database, err := sql.Open("mysql", dsn); err != nil {
		panic(err)
	} else {
		return database
	}
}
