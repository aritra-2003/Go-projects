package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)
const secretkey="supersecret"
func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"emails": email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretkey))
}
func VerifyToken(t string)(int64,error){
	pt,err:=jwt.Parse(t,func(t *jwt.Token)(interface{},error){
		_,ok:=t.Method.(*jwt.SigningMethodHMAC)
		if !ok{
			return nil,errors.New("unexpected signing")
		}
		return secretkey,nil
	})
	if err!=nil{
		return 0,errors.New("error in parsing")
	}
	tokenIsValid:=pt.Valid
	if !tokenIsValid{
		return 0,errors.New("token is invalid")
	}
cl,ok:=pt.Claims.(jwt.MapClaims)
   if !ok{
	return 0,errors.New("invalid token")
	
   }
//    emails:=cl["emails"].(string)
   userId:=int64(cl["userId"].(float64))
return userId, nil


}