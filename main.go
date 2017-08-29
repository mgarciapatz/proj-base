package main

import (
	"fmt"
	"net/http"
	"proj-base/routers"
	"github.com/rs/cors"
	_"proj-base/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
    _"github.com/go-sql-driver/mysql"
	"github.com/codegangsta/negroni"
)
var (
	app_port 	= ":"+beego.AppConfig.String("httpport")

	db_host		= beego.AppConfig.String("db_host")
	db_port     = beego.AppConfig.String("db_port")
    db_user     = beego.AppConfig.String("db_user")
    db_pass     = beego.AppConfig.String("db_pass")
    db_name     = beego.AppConfig.String("db_name")
    runmode     = beego.AppConfig.String("runmode")
)

func init() {
	err := orm.RegisterDataBase("default", "mysql", db_user + ":" + db_pass + "@tcp(" + db_host + ":" + db_port + ")/" + db_name + "?charset=utf8")
	if err != nil {
        panic(fmt.Sprintf("Database connection error: %+v", err))
    }
	if runmode == "dev" {
		orm.Debug = true
	}
}

func main() {
	cross_origin := cors.New(cors.Options{
		AllowCredentials	:	true,
		OptionsPassthrough	:	false,
		AllowedHeaders		:	[]string{"Origin", "Authorization", "Access-Control-Allow-Origin","Content-Type"},
		AllowedMethods		:	[]string{"GET", "POST", "DELETE","PATCH"},
	})
	router	:=	routers.InitRoutes()
	negro	:=	negroni.Classic()
	negro.Use(cross_origin)
	negro.UseHandler(router)
	http.ListenAndServe(app_port, negro)
}
