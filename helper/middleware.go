package helper

/*
    NOTE:
        The general purpose of this package is for different
        second hand open source code. Examples are;
        1. MiddleWares
        2. JWT Authentication
*/
import (
	"time"
	"net/http"
	jwt "github.com/dgrijalva/jwt-go"
)

var Secret  =   "!@#$%^&*()@<3GodLoveYouAndMe<3@)((*&#$#$@#)@@)"

/*
    This is a public method
    @param    string          required
    @param    ResponseWriter  required
    @return   int, string
*/
func SignToken(crypt string,w http.ResponseWriter)(int,string){
    token, err := generateToken(crypt,w)
    if err != nil {
        return http.StatusInternalServerError, ""
    } else {
        return http.StatusOK, token
    }
}
/**
    This is a private method, package specific access
    This method will generate sign token SigningMethodHS256 signing method with signkey provided


    @param    string          required
    @param    ResponseWriter  required

    @return   string, interface
*/
func generateToken(crypt string,w http.ResponseWriter)(string,interface{}){
    var signkey  =  []byte(Secret)
	claims      :=  generateClaims(crypt)
    jwtoken     :=  jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
    token, err  :=  jwtoken.SignedString(signkey)
    if err != nil {
        return "", err
    }
    return token, nil
}
/**
    This is a private method, package specific access
    This method will generate claims that includes expiration


    @param    string      required

    @return   claims
*/
func generateClaims(crypt string)jwt.MapClaims{
    claims := make(jwt.MapClaims)
    /* this will expire in 2 hour */
    claims["exp"]   =   time.Now().Add(time.Hour + time.Duration(2)).Unix()
    claims["sub"] 	= 	crypt
    return claims
}
