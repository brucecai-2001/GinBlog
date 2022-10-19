package middleware

import (
	"GinBlog/utils"
	"GinBlog/utils/errmsg"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JwtKey = []byte(utils.Jwtkey)

type MyClaim struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

//生成Token
func SetToken(username, password string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)
	Claim := MyClaim{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "JoJo",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, Claim)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS
}

//验证Token
func CheckToken(token string) (*MyClaim, int) {
	settoken, _ := jwt.ParseWithClaims(token, &MyClaim{}, func(t *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, ok := settoken.Claims.(*MyClaim); ok && settoken.Valid {
		return key, errmsg.SUCCESS
	} else {
		return nil, errmsg.ERROR
	}
}

//jwt中间件

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		code := errmsg.SUCCESS

		if tokenHeader == "" {
			code = errmsg.ERROR_Token_Not_Exist
		}
		//验证Token格式
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_FORMAT
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.Get_Error_Msg(code),
			})
			c.Abort()
			return
		}
		//验证Token是否正确
		key, res := CheckToken(checkToken[1])
		if res == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.Get_Error_Msg(code),
			})
			c.Abort()
			return
		}
		//验证Token是否过期
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.Get_Error_Msg(code),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": errmsg.Get_Error_Msg(code),
		})
		c.Set("username", key.Username)
		c.Next()
	}
}
