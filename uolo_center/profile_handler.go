package uolo_center

import (
	"artemis/uolo_center/model"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type (
	ProfileHandler struct {
	}

	ProfileResponse struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	}
)

func NewProfileHandler() *ProfileHandler {
	return &ProfileHandler{}
}

func (handler *ProfileHandler) RegisterRouter(ec *echo.Echo) {
	// https://api.echoface.cn/user/profile/save
	ec.POST("/user/profile/save", handler.SaveUserProfile)
}

func (handler *ProfileHandler) SaveUserProfile(e echo.Context) error {

	type (
		Profile struct {
			Id       string `json:"id" form:"id" binding:"required"`
			Sex      int32  `json:"sex" form:"sex" binding:"required"`
			NickName string `json:"nickName" form:"nickName" binding:"required"`
			BirthDay string `json:"birthday" form:"birthday" binding:"required"`
			//LivePlace  string `json:"livePlace" form:"livePlace" binding:"required"`
			//Profession string `json:"profession" form:"profession" binding:"required"`
		}
	)
	in := &Profile{}
	res := ProfileResponse{}

	responser := func(code int32, msg string) error {
		res.Code = code
		res.Message = msg
		return e.JSON(http.StatusOK, res)
	}

	if err := e.Bind(in); err != nil {
		return responser(-1, "参数错误")
	}

	logrus.Debugln("SaveUserProfile Got Submited Data:", in)

	profile := &model.UserProfile{}

	found, err := ormEngine.Where("id = ?", in.Id).Get(profile)
	if err != nil {
		logrus.Errorln("Select * From user_profile where id = ", in.Id, " Failed", err.Error())
		return responser(-2, "数据库错误")
	}
	if !found {
		logrus.Errorln("Select * From user_profile where id = ", in.Id, " Failed", err.Error())
		return responser(-3, "没有此用户")
	}

	profile.Sex = in.Sex
	profile.NickName = in.NickName
	if birth, err := time.Parse(time.RFC3339Nano, in.BirthDay); err == nil {
		profile.Birth = birth
	} else {
		logrus.Errorln("Birth Day Parse Failed, ", in.BirthDay)
	}

	if _, err := ormEngine.Update(profile); err != nil {
		logrus.Errorln("xorm update error ", err.Error())
		return responser(-4, "保存失败")
	}

	return responser(0, "success")
}

func (handler *ProfileHandler) UpdateUserAvastImage(e echo.Context) error {
	return e.JSON(http.StatusOK, &ProfileResponse{
		Code:    0,
		Message: "success",
	})
}
