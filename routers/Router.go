package routers
import (
	"net/http"
	"proj-base/helper"
	"github.com/gorilla/mux"
	"github.com/astaxie/beego"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/auth0/go-jwt-middleware"
)
var app_version = beego.AppConfig.String("appversion")

var mySigningKey = []byte(helper.Secret)

var myHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("Auth")
	if cookie.Value == "none" {
		jwtmiddleware.OnError(w,r,"")
	}
	return
})

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter : func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
    },
    SigningMethod: jwt.SigningMethodHS256,
})

func InitRoutes() *mux.Router {
	router :=	mux.NewRouter()
	router	=	SampleRoutes(router)
	return router
}
