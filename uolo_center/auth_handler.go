package uolo_center

import (
	"artemis/uolo_center/model"
	"artemis/uolo_center/utils"
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
	ErrUserAlreadyExist  = -4
	ErrIncorrectPassword = -5
	ErrEmailFormatError  = -6
	TokenMaxAge          = time.Hour * 24
)

type (
	AuthHandler struct {
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

func (handler *AuthHandler) SignUp(ec echo.Context) error {
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
		Sex:      1,
		Tags:     0,
		Status:   0,
		Birth:    tNow,
	}
	found, err := ormEngine.Where("username = ?", form.Name).Get(&model.UserProfile{})
	if err != nil {
		return AbortWithError(ec, ErrBadDatabaseQuery, "数据查询失败")
	}
	if found {
		return AbortWithError(ec, ErrUserAlreadyExist, "用户已存在")
	}

	uuidGen, _ := uuid.NewV4()
	user.Id = uuidGen.String()
	user.Avatar = user.Id

	if digest, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost); err != nil {
		return AbortWithError(ec, ErrFailedEncryptPsd, "签发密码失败")
	} else {
		user.Password = string(digest)
	}

	logrus.Infoln("Create New User:", user)

	_, err = ormEngine.Insert(user)
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

func (handler *AuthHandler) Login(ec echo.Context) error {
	logrus.Debugln("AuthHandler.Login Enter")

	type LoginForm struct {
		Name     string `form:"username" json:"username" binding:"required" query:"username"`
		Password string `form:"password" json:"password" binding:"required" query:"password"`
	}

	form := &LoginForm{}

	if err := ec.Bind(form); err != nil {
		return AbortWithError(ec, ErrInvalidArguments, "Missing username or password")
	}

	u := &model.UserProfile{}
	found, err := ormEngine.Where("username = ?", form.Name).Get(u)
	if err != nil {
		return AbortWithError(ec, ErrBadDatabaseQuery, "数据库错误")
	}

	if found == false || bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(form.Password)) != nil {
		return AbortWithError(ec, ErrIncorrectPassword, "用户名或密码错误")
	}

	expire := time.Now().Add(TokenMaxAge)

	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.Id
	claims["exp"] = expire.Unix()

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(appConf.JWTSecretkey))

	if err != nil {
		return AbortWithError(ec, ErrFailedEncryptPsd, "Create Token faild")
	}

	utils.SetJwtAuthCookie(ec, tokenString, ".echoface.cn")
	utils.SetUoloUserCookieIfNeeded(ec, u.Id, ".echoface.cn")

	logrus.Debugln("AuthHandler.Login Leave Success")
	return ec.JSON(http.StatusOK, DefaultRes{
		Code: 0,
		User: u,
		Token: &TokenInfo{
			Token:  tokenString,
			MaxAge: 12,
		},
		Message: "Success",
	})
}

func (handler *AuthHandler) AuthRefresh(ec echo.Context) error {

	oldToken := ec.Get("user").(*jwt.Token)
	if oldToken == nil {
		return AbortWithError(ec, ErrFailedEncryptPsd, "获取token失败")
	}
	oldClaims := oldToken.Claims.(jwt.MapClaims)
	oldId := oldClaims["id"]

	expire := time.Now().Add(TokenMaxAge)

	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = oldId
	claims["exp"] = expire.Unix()

	tokenString, err := token.SignedString([]byte(appConf.JWTSecretkey))
	if err != nil {
		return AbortWithError(ec, ErrFailedEncryptPsd, "签发token失败")
	}

	utils.SetJwtAuthCookie(ec, tokenString, ".echoface.cn")
	utils.SetUoloUserCookieIfNeeded(ec, oldId.(string), ".echoface.cn")

	return ec.JSON(http.StatusOK, DefaultRes{
		Code:    0,
		Message: "Success",
		Token: &TokenInfo{
			MaxAge: 12,
			Token:  tokenString,
		},
	})
}

func (handler *AuthHandler) AuthTest(ec echo.Context) error {
	logrus.Debugln("AuthHandler.AuthTest Enter")
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

	logrus.Debugf("AuthHandler.AutoTest Leave For UserUUID:%s", userUUID.(string))

	return ec.JSON(http.StatusOK, DefaultRes{
		Code:    0,
		Message: "Wellcome!, " + userUUID.(string),
	})
}

func (handler *AuthHandler) ResetPassword(ec echo.Context) error {
	type CPswdForm struct {
		Old     string `form:"old" json:"old" binding:"required"`
		New     string `form:"new" json:"new" binding:"required"`
		Confirm string `form:"confirm" json:"confirm" binding:"required"`
		Captcha string `form:"captcha" json:"captcha" binding:"required"`
	}

	token := ec.Get("user").(*jwt.Token)
	oldClaims := token.Claims.(jwt.MapClaims)
	userId := oldClaims["id"]

	form := &CPswdForm{}
	if err := ec.Bind(form); err != nil {
		logrus.Debugln("ResetPasswod Form Error, Request:", ec.Request())
		return AbortWithError(ec, ErrInvalidArguments, "Missing Required Arguments")
	}

	u := &model.UserProfile{}
	found, err := ormEngine.Where("id = ?", userId).Get(u)
	if err != nil {
		return AbortWithError(ec, ErrBadDatabaseQuery, "数据库错误")
	}

	if found == false || bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(form.Old)) != nil {
		return AbortWithError(ec, ErrIncorrectPassword, "密码错误")
	}

	if form.New == form.Old {
		return AbortWithError(ec, ErrIncorrectPassword, "新密码没有变动")
	}

	if form.New != form.Confirm || len(form.New) < 6 {
		return AbortWithError(ec, ErrInvalidArguments, "密码不一致")
	}

	if digest, err := bcrypt.GenerateFromPassword([]byte(form.New), bcrypt.DefaultCost); err != nil {
		return AbortWithError(ec, ErrFailedEncryptPsd, "签发密码失败")
	} else {
		u.Password = string(digest)
	}

	_, err = ormEngine.Update(u)
	if err != nil {
		return AbortWithError(ec, ErrBadDatabaseQuery, "保存失败")
	}

	return ec.JSON(http.StatusOK, DefaultRes{
		Code:    0,
		Message: "修改成功",
	})
}
