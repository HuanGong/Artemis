package uolo_center

import (
	"time"
	"net/http"
	"github.com/labstack/echo"
	"artemis/uolo_center/model"
	"github.com/sirupsen/logrus"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"github.com/satori/go.uuid"
)

const (
	ExpireTime   = time.Hour * 2
)

type (
	Handler struct {
	}
)

func AbortWithError(c echo.Context, code int, message string) error {
	type Res struct {
		Code int
		Message string
	}

	c.Response().Header().Set("WWW-Authenticate", "JWT realm = jwt realm")
	res := &Res {
		Code: code,
		Message: message,
	}
	return c.JSON(code, res)
}

func (handler *Handler) SignUp(ec echo.Context) error {
	logrus.Infoln("SignUP enter")


	user := &model.User{}
	if err := ec.Bind(user); err != nil {
		logrus.Infof("Request Arguments Error When Bind")
		return &echo.HTTPError{Code:http.StatusBadRequest, Message: "Invailed Request Arguments",}
	}

	logrus.Infoln("user:", user)
	if user.Email == "" || user.Password == "" {
		return &echo.HTTPError{Code:http.StatusBadRequest, Message: "Email or Password Error",}
	}

	found, err := orm.Where("username = ?", user.Name).Get(user)
	if err != nil {
		logrus.Errorln("db error:", err.Error())
		return AbortWithError(ec, http.StatusInternalServerError, "数据查询失败")
	}
	if found {
		return AbortWithError(ec, -1, "用户已存在")
	}

	uuidGen, _ := uuid.NewV4()
	user.Id = uuidGen.String()

	if digest, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost); err != nil {
		return AbortWithError(ec, -1, err.Error())
	} else {
		user.Password = string(digest)
	}

	logrus.Infoln("Create New User:", user)

	_, err = orm.Insert(user)
	if err != nil {
		return AbortWithError(ec, -1, "新建用户失败")
	}

	return ec.JSON(http.StatusOK, user)
}

func (handler *Handler) Login(ec echo.Context) error {

	type LoginForm struct {
		Name 	 string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}

	form := &LoginForm{}

	if err := ec.Bind(form); err != nil {
		return AbortWithError(ec, http.StatusBadRequest, "Missing usename or password")
	}

	u := &model.User{}
	found, err := orm.Where("username = ?", form.Name).Get(u)
	if err != nil {
		return AbortWithError(ec, http.StatusInternalServerError, "DB Query Error")
	}

	if found == false || bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(form.Password)) != nil {
		return AbortWithError(ec, http.StatusUnauthorized, "Incorrect Username / Password")
	}

	expire := time.Now().Add(ExpireTime)

	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.Id
	claims["exp"] = expire.Unix()

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(conf.JWTSecretkey))

	if err != nil {
		return AbortWithError(ec, http.StatusUnauthorized, "Create JWT Token faild")
	}

	json := map[string]string {
		"token": tokenString,
		"expire": expire.Format(time.RFC3339),
	}
	return ec.JSON(http.StatusOK, json)
}

func (handler *Handler) AuthRefresh(ec echo.Context) error {

	user := ec.Get("user").(*jwt.Token)

	oldClaims := user.Claims.(jwt.MapClaims)
	oldId := oldClaims["id"]


	expire := time.Now().Add(ExpireTime)

	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = expire.Unix()
	claims["id"] = oldId

	tokenString, err := token.SignedString([]byte(conf.JWTSecretkey))

	if err != nil {
		return AbortWithError(ec, http.StatusUnauthorized, "Create JWT Token faild")
	}

	json := map[string]string {
		"token": tokenString,
		"expire": expire.Format(time.RFC3339),
	}
	return ec.JSON(http.StatusOK, json)
}

func (handler *Handler) AuthTest(ec echo.Context) error {
	user := ec.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["id"]
	return ec.String(http.StatusOK, "Welcome "+ name.(string) +"!")
}