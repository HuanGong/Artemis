package uolo_center

import (
	"artemis/uolo_center/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

const (
	ErrInvalidArguments  = -1
	ErrBadDatabaseQuery  = -2
	ErrFailedEncryptPsd  = -3
	ErrUserAreadyExsist  = -4
	ErrIncorrectPassword = -5
	ErrEmailFormatError  = -6
	TokenMaxAge          = time.Hour * 2
)

type (
	Handler struct {
	}
	DefaultRes struct {
		Code    int                `json:"code"`
		Message string             `json:"message"`
		Token   *TokenInfo         `json:"token,omitempty"`
		User    *model.UserProfile `json:"user,omitempty"`
	}
	TokenInfo struct {
		Token  string `json:"token"`
		MaxAge int32  `json:"max_age"`
	}
)

func AbortWithError(c echo.Context, code int, message string) error {

	c.Response().Header().Set("WWW-Authenticate", "JWT realm = jwt realm")
	res := &DefaultRes{
		Code:    code,
		Message: message,
	}
	return c.JSON(http.StatusOK, res)
}

func (handler *Handler) SignUp(ec echo.Context) error {
	logrus.Infoln("SignUP enter")

	type SignUpForm struct {
		Name     string `form:"username" json:"username" binding:"required" query:"username"`
		Password string `form:"password" json:"password" binding:"required" query:"password"`
		Email    string `form:"email" json:"email" binding:"required" query:"email"`
	}

	form := &SignUpForm{}
	if err := ec.Bind(form); err != nil {
		logrus.Infof("Request Arguments Error When Bind")
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invailed Request Arguments"}
	}
	logrus.Infoln("SignUP data:", form)
	if !model.CheckEmail(form.Email) {
		return AbortWithError(ec, ErrEmailFormatError, "email格式错误")
	}
	if len(form.Password) < 6 {
		return AbortWithError(ec, ErrIncorrectPassword, "密码长度不够")
	}
	if len(form.Name) < 1 {
		return &echo.HTTPError{Code: ErrInvalidArguments, Message: "用户名非法"}
	}
	tNow := time.Now()

	user := &model.UserProfile{
		Name:     form.Name,
		Email:    form.Email,
		NickName: form.Name,
		AvastIco: form.Name, //TODO： gen avastinco path
		Sex:      1,
		Tags:     0,
		Status:   0,
		Birth:    tNow,
	}
	found, err := orm.Where("username = ?", form.Name).Get(&model.UserProfile{})
	if err != nil {
		return AbortWithError(ec, ErrBadDatabaseQuery, "数据查询失败")
	}
	if found {
		return AbortWithError(ec, ErrUserAreadyExsist, "用户已存在")
	}

	uuidGen, _ := uuid.NewV4()
	user.Id = uuidGen.String()

	if digest, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost); err != nil {
		return AbortWithError(ec, ErrFailedEncryptPsd, "签发密码失败")
	} else {
		user.Password = string(digest)
	}

	logrus.Infoln("Create New User:", user)

	_, err = orm.Insert(user)
	if err != nil {
		logrus.Errorln("Create UserProfile Failed With error:", err.Error())
		return AbortWithError(ec, ErrBadDatabaseQuery, "新建用户失败")
	}

	user.Password = ""

	return ec.JSON(http.StatusOK, DefaultRes{
		Code:    0,
		User:    user,
		Message: "注册成功",
	})
}

func (handler *Handler) Login(ec echo.Context) error {
	logrus.Debugln("Handler.Login Enter")
	type LoginForm struct {
		Name     string `form:"username" json:"username" binding:"required" query:"username"`
		Password string `form:"password" json:"password" binding:"required" query:"password"`
	}

	form := &LoginForm{}

	if err := ec.Bind(form); err != nil {
		return AbortWithError(ec, ErrInvalidArguments, "Missing usename or password")
	}

	u := &model.User{}
	found, err := orm.Where("username = ?", form.Name).Get(u)
	if err != nil {
		return AbortWithError(ec, ErrBadDatabaseQuery, "DB Query Error")
	}

	if found == false || bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(form.Password)) != nil {
		return AbortWithError(ec, ErrIncorrectPassword, "Incorrect Username / Password")
	}

	expire := time.Now().Add(TokenMaxAge)

	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.Id
	claims["exp"] = expire.Unix()

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(conf.JWTSecretkey))

	if err != nil {
		return AbortWithError(ec, ErrFailedEncryptPsd, "Create Token faild")
	}

	u.Password = "********"
	uidCookie := &http.Cookie{
		Name:     "UoloAU",
		Value:    tokenString,
		HttpOnly: false,
		Domain:   "",
		MaxAge:   int(time.Hour * 2),
	}
	ec.SetCookie(uidCookie)

	logrus.Debugln("Handler.Login Leave Success")
	return ec.JSON(http.StatusOK, DefaultRes{
		Code: 0,
		//User: u,
		Token: &TokenInfo{
			Token:  tokenString,
			MaxAge: 2,
		},
		Message: "Success",
	})
}

func (handler *Handler) AuthRefresh(ec echo.Context) error {

	oldToken := ec.Get("user").(*jwt.Token)

	oldClaims := oldToken.Claims.(jwt.MapClaims)
	oldId := oldClaims["id"]

	expire := time.Now().Add(TokenMaxAge)

	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = oldId
	claims["exp"] = expire.Unix()

	tokenString, err := token.SignedString([]byte(conf.JWTSecretkey))

	if err != nil {
		return AbortWithError(ec, ErrFailedEncryptPsd, "签发token失败")
	}

	uidCookie := &http.Cookie{
		Name:     "UoloAU",
		Value:    tokenString,
		HttpOnly: true,
		MaxAge:   int(TokenMaxAge),
	}
	ec.SetCookie(uidCookie)

	return ec.JSON(http.StatusOK, DefaultRes{
		Code:    0,
		Message: "Success",
		Token: &TokenInfo{
			MaxAge: 2,
			Token:  tokenString,
		},
	})
}

func (handler *Handler) AuthTest(ec echo.Context) error {
	logrus.Debugln("Handler.AuthTest Enter")
	for _, cookie := range ec.Cookies() {
		logrus.Debugln("cookie: ", cookie)
	}

	ctxData := ec.Get("user")
	if ctxData == nil {
		return echo.ErrForbidden
	}
	jwtToken := ctxData.(*jwt.Token)
	claims := jwtToken.Claims.(jwt.MapClaims)
	userUUID := claims["id"]

	logrus.Debugf("Handler.AutoTest Leave For UserUUID:%s", userUUID.(string))

	return ec.JSON(http.StatusOK, DefaultRes{
		Code:    0,
		Message: "Wellcome!, " + userUUID.(string),
	})
}

func (handler *Handler) ResetPassword(ec echo.Context) error {
	type CPswdForm struct {
		Name    string `form:"name" json:"name" binding:"required"`
		Old     string `form:"old" json:"old" binding:"required"`
		New     string `form:"new" json:"new" binding:"required"`
		Confirm string `form:"confirm" json:"confirm" binding:"required"`
	}

	token := ec.Get("user").(*jwt.Token)
	oldClaims := token.Claims.(jwt.MapClaims)
	userId := oldClaims["id"]

	form := &CPswdForm{}
	if err := ec.Bind(form); err != nil {
		logrus.Debugln("ResetPasswod Form Error, Request:", ec.Request())
		return AbortWithError(ec, ErrInvalidArguments, "Missing Required Arguments")
	}

	u := &model.User{}
	found, err := orm.Where("username = ?", form.Name).Get(u)
	//select user from xxxx where username=xxxx
	if err != nil {
		return AbortWithError(ec, ErrBadDatabaseQuery, "数据库错误")
	}

	if found == false || bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(form.Old)) != nil {
		return AbortWithError(ec, ErrIncorrectPassword, "密码错误")
	}

	if userId != u.Id { //防止获取到别人用户名和密码的情况下，冒用别人的token来修改别人的密码
		return AbortWithError(ec, ErrIncorrectPassword, "ID不匹配")
	}

	if form.New != form.Confirm {
		return AbortWithError(ec, ErrInvalidArguments, "密码不一致")
	}

	if digest, err := bcrypt.GenerateFromPassword([]byte(form.New), bcrypt.DefaultCost); err != nil {
		return AbortWithError(ec, ErrFailedEncryptPsd, "签发密码失败")
	} else {
		u.Password = string(digest)
	}

	_, err = orm.Insert(u)
	if err != nil {
		return AbortWithError(ec, ErrBadDatabaseQuery, "保存失败")
	}

	return ec.JSON(http.StatusOK, DefaultRes{
		Code:    0,
		Message: "修改成功",
	})
}
