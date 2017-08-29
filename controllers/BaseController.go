package controllers

import(
    "net/http"
    "io/ioutil"
    "crypto/md5"
    "encoding/hex"
    "encoding/json"
    "proj-base/helper"
)

func HttpReq(w http.ResponseWriter, req *http.Request) (map[string]interface{}, string) {
    response,err := ioutil.ReadAll(req.Body)
    if err == nil {
        json_data := make(map[string]interface{})
        json.Unmarshal(response, &json_data)
        if len(json_data) != 0 {
            return json_data, ""
        } else {
            return nil, "Request parameters not found or empty."
        }
    } else {
        return nil, "Invalid request."
    }
}

func HttpRes(obj interface{},w http.ResponseWriter) (string){
	json_data,json_error := json.Marshal(obj)
    if json_error == nil {
        w.Header().Set("Content-Type", "application/json")
        w.Write(json_data)
        return ""
    }
    return "Invalid response data."
}

func HttpMethod(req *http.Request) (string) {
    return req.Method
}

func Encrypt(key string) string {
	token := ")*&^%$#@!13423423Proverbs"+key+"NineteenFourteen!@#$%^&*()_03746253"
    o := md5.New()
    o.Write([]byte(token))
    return hex.EncodeToString(o.Sum(nil))
}

func Decrypt(key string, crypt string) bool {
	token :=  ")*&^%$#@!13423423Proverbs" + key + "NineteenFourteen!@#$%^&*()_03746253"
    m := md5.New()
    m.Write([]byte(token))
    encrypt  := hex.EncodeToString(m.Sum(nil))
	if encrypt == crypt {
		return true;
	}
	return false;
}

func JWTokenizer(w http.ResponseWriter, r *http.Request){
	_, token := helper.SignToken("thisHasToBeARandomString",w)
	HttpRes(map[string]interface{}{
		"statusCode": 200,
		"devMessage": token,
	} , w)
}

func ProtectedAction(w http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	HttpRes(map[string]interface{}{
		"statusCode": 200,
		"devMessage": "Nice have successfully conquered Westeros",
	} , w)
}
