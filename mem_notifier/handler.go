package mem_notifier

import (
	"gopkg.in/mgo.v2"
	"github.com/labstack/echo"
	"artemis/mem_notifier/model"
	"gopkg.in/mgo.v2/bson"
	"github.com/sirupsen/logrus"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	key = "secret"
)

type (
	Handler struct {
		DB *mgo.Session
	}
)

func (handler *Handler) SignUp(ec echo.Context) error {
	logrus.Infoln("SignUP enter")

	user := &model.User{
		ID: bson.NewObjectId(),
	}
	if err := ec.Bind(user); err != nil {
		logrus.Infof("Request Arguments Error When Bind")
		return &echo.HTTPError{Code:http.StatusBadRequest, Message: "Invaild Request Arguments",}
	}

	if user.Email == "" || user.Password == "" {
		return &echo.HTTPError{Code:http.StatusBadRequest, Message: "Email or Password Error",}
	}

	/*
	db := handler.DB.Clone()
	defer db.Close()

	if err := db.DB("MemNotifier").C("users").Insert(user); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "User Not Created",}
	}
	*/
	logrus.Infoln("Create New User:", user)
	return ec.JSON(http.StatusCreated, user)
}

func (handler *Handler) Login(ec echo.Context) error {
	u := new(model.User)
	if err := ec.Bind(u); err != nil {
		return err
	}

	// Find user
	/*
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("twitter").C("users").
		Find(bson.M{"email": u.Email, "password": u.Password}).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email or password"}
		}
		return
	}
	*/
	//-----
	// JWT
	//-----

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response
	var err error
	u.Token, err = token.SignedString([]byte(key))
	if err != nil {
		return err
	}

	u.Password = "" // Don't send password
	return ec.JSON(http.StatusOK, u)
}