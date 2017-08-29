/*
	NOTE
		This sample is intended for guide and examples only
		some of the function does not have some data validation
		or even tokenizers.
*/

package controllers

import(
	"strconv"
	"net/http"
	_"database/sql"
	"proj-base/models"
	"github.com/gorilla/mux"
	"github.com/astaxie/beego/orm"
    _"github.com/go-sql-driver/mysql"
)
/*
	NOTE
	This is a sample REQUEST READ only function
	{
		"first_name"	:	"Conrado",
		"last_name"		:	"Consolacion"
	}
*/
func SampleController(w http.ResponseWriter, r *http.Request) {
	if HttpMethod(r) == "POST" {
		result, resultError := HttpReq(w,r)
		if resultError == "" {
			if result["first_name"].(string) == "" {
				HttpRes("First Name is required.", w)
				return
			} else if result["last_name"].(string) == "" {
				HttpRes("Last Name is required.", w)
				return
			} else {
				newSample                 :=  new(models.SampleData)
				newSample.SampleFirstName =   result["first_name"].(string)
				newSample.SampleLastName  =   result["last_name"].(string)
				HttpRes(newSample, w)
				return
			}
		}else {
			HttpRes(resultError, w)
			return
		}
	} else {
		HttpRes("Invalid request method", w)
		return
	}
}
/*
	NOTE
	This is a sample insert function
	{
		"first_name"	:	"Conrado",
		"last_name"		:	"Consolacion"
	}
*/
func SampleInsertController(w http.ResponseWriter, r *http.Request) {
	if HttpMethod(r) == "POST" {
		result, resultError := HttpReq(w,r)
		if resultError == "" {
			if result["first_name"].(string) == "" {
				HttpRes("First Name is required.", w)
				return
			} else if result["last_name"].(string) == "" {
				HttpRes("Last Name is required.", w)
				return
			} else {
				var newSample models.SampleData
				newSample.SampleFirstName =   result["first_name"].(string)
				newSample.SampleLastName  =   result["last_name"].(string)
				create := orm.NewOrm()
			    key, err := create.Insert(&newSample)
				if err == nil {
					id := strconv.Itoa(int(key))
					devMessage := map[string]interface{}{
						"key"		:	id,
						"token" 	: 	Encrypt(id),
						"first_name":	newSample.SampleFirstName,
						"last_name"	:	newSample.SampleLastName,
					}
					HttpRes(devMessage, w)
					return
				} else {
					HttpRes(err, w)
					return
				}
			}
		} else {
			HttpRes(resultError, w)
			return
		}
	} else {
		HttpRes("Invalid request method", w)
		return
	}
}
/*
	NOTE
	This is a sample insert function
	{
		"key"			:	1,
		"first_name"	:	"Conrado",
		"last_name"		:	"Consolacion"
	}
*/
func SampleReadController(w http.ResponseWriter, r *http.Request) {
	if HttpMethod(r) == "GET" {
		if r.FormValue("first_name") == "" {
			HttpRes("First name is required", w)
			return
		}
		if r.FormValue("last_name") == "" {
			HttpRes("Last Name is required", w)
			return
		}
		if r.FormValue("key") == "" {
			HttpRes("Key is required", w)
			return
		}
		var readSample models.SampleData
		read := orm.NewOrm()
	    err := read.
				QueryTable("SampleData").
				Filter("Autokey",r.FormValue("key")).
				One(&readSample)
		if err == nil {
			HttpRes(readSample, w)
			return
		} else {
			HttpRes(err, w)
			return
		}
	} else {
		HttpRes("Invalid request method", w)
		return
	}
}
/*
	NOTE
	This is a sample insert function
	{
		"key"			: 	1,
		"first_name"	:	"Conrado",
		"last_name"		:	"Consolacion"
	}
*/
func SampleUpdateController(w http.ResponseWriter, r *http.Request) {
	if HttpMethod(r) == "POST" {
		result, resultError := HttpReq(w,r)
		if resultError == "" {
			if result["first_name"].(string) == "" {
				HttpRes("First name is required",w)
				return
			} else if result["last_name"].(string) == "" {
				HttpRes("Last name is required.",w)
				return
			} else if result["key"].(float64) == 0 {
				HttpRes("key is requred",w)
				return
			} else {
				data_update := orm.Params{
			        "sample_first_name" :   result["first_name"].(string),
			        "sample_last_name"  :   result["last_name"].(string),
			    }
			    update := orm.NewOrm()
			    _, err := update.QueryTable("SampleData").Filter("Autokey",result["key"]).Update(data_update)
				if err == nil {
					HttpRes(data_update,w)
					return
				} else {
					HttpRes(err, w)
					return
				}
			}
		} else {
			HttpRes(resultError,w)
			return
		}
	} else {
		HttpRes("Invalid request method", w)
		return
	}
}
/*
	NOTE
	This is a sample insert function
	{
		"key"	:	1
	}
*/
func SampleDeleteController(w http.ResponseWriter, r *http.Request) {
	if HttpMethod(r) == "POST" {
		result, resultError := HttpReq(w,r)
		if resultError == "" {
			if result["key"].(float64) == 0 {
				HttpRes("key is requred",w)
				return
			} else {
				delete := orm.NewOrm()
			    _,err := delete.QueryTable("SampleData").Filter("Autokey",result["key"]).Delete()
				if err == nil {
					HttpRes("Data deleted",w)
					return
				} else {
					HttpRes(err, w)
					return
				}
			}
		} else {
			HttpRes(resultError,w)
			return
		}
	} else {
		HttpRes("Invalid request method", w)
		return
	}
}

/*
	NOTE
	This is for the second approach of crud operation only
	{
		"first_name"	:	"Conrado",
		"last_name"		:	"Consolacion"
	}
*/
func AllInOneCrudController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	link := mux.Vars(r)
	if HttpMethod(r) == "POST" {
		if link["action"] == "add" {
			result, resultError := HttpReq(w,r)
			if resultError == "" {
				if result["first_name"].(string) == "" {
					HttpRes("First Name is required.", w)
					return
				} else if result["last_name"].(string) == "" {
					HttpRes("Last Name is required.", w)
					return
				} else {
					var newSample models.SampleData
					newSample.SampleFirstName =   result["first_name"].(string)
					newSample.SampleLastName  =   result["last_name"].(string)
					create := orm.NewOrm()
				    key, err := create.Insert(&newSample)
					if err == nil {
						id := strconv.Itoa(int(key))
						devMessage := map[string]interface{}{
							"key"		:	id,
							"token" 	: 	Encrypt(id),
							"first_name":	newSample.SampleFirstName,
							"last_name"	:	newSample.SampleLastName,
						}
						HttpRes(devMessage, w)
						return
					} else {
						HttpRes(err, w)
						return
					}
				}
			} else {
				HttpRes(resultError, w)
				return
			}
		} else {
			HttpRes("Invalid action type", w)
			return
		}
	} else {
		HttpRes("Invalid request method", w)
		return
	}
}
