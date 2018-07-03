package uolo_center

import (
	"artemis/component/avatar_keeper"
	"artemis/uolo_center/model"
	"artemis/uolo_center/utils"
	"bytes"
	"encoding/base64"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type (
	OptionConfig struct {
		Path          string
		DefaultAvatar string
	}

	ProfileHandler struct {
		config OptionConfig
		keeper *avatar_keeper.AvatarKeeper
	}

	ProfileResponse struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	}
)

func NewProfileHandler(conf OptionConfig) *ProfileHandler {
	keeper := avatar_keeper.NewAvatarKeeperWithFileStorage(conf.Path)
	return &ProfileHandler{
		config: conf,
		keeper: keeper,
	}
}

func (handler *ProfileHandler) RegisterRouter(ec *echo.Echo) {
	// https://api.echoface.cn/user/profile/save
	ec.POST("/user/profile/save", handler.SaveUserProfile)
	ec.POST("/user/profile/avatar/upload", handler.UploadProfilePhoto)
	ec.GET("/user/profile/avatar", handler.GetProfilePhoto)

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

func (handler *ProfileHandler) UploadProfilePhoto(ec echo.Context) error {
	logrus.Debugln("UploadProfilePhoto Enter")
	type (
		In struct {
			Id   string `json:"id" form:"id"`
			Data string `json:"img" form:"img"`
		}
		Res struct {
			Code    int32  `json:"code"`
			Message string `json:"message"`
		}
	)
	in := &In{}
	res := &Res{}
	if err := ec.Bind(in); err != nil {
		logrus.Debugln("UploadProfilePhoto Bind Form Error")
		res.Code = -1
		res.Message = "参数错误"
		return ec.JSON(http.StatusOK, res)
	}

	in.Data = strings.Replace(in.Data, " ", "+", -1)
	//logrus.Debugln("data:", in.Data)

	rawData, err := base64.StdEncoding.DecodeString(in.Data)
	if err != nil {
		res.Code = -2
		res.Message = "解码失败"
		return ec.JSON(http.StatusOK, res)
	}
	buf := bytes.NewBuffer(rawData)

	if err := handler.keeper.SaveNewProfilePhoto(in.Id, buf); err != nil {
		res.Code = -3
		res.Message = "存储失败"
		return ec.JSON(http.StatusOK, res)
	}

	res.Code = 0
	res.Message = "OK"
	return ec.JSON(http.StatusOK, res)
}

func (handler *ProfileHandler) GetProfilePhoto(ec echo.Context) error {
	type (
		In struct {
			Standard string `query:"standard" form:"standard" binding:"required"`
		}
	)
	userId, login := utils.IsUserLogin(ec)
	if !login || len(userId) == 0 {
		if defaultAvatar, err := ioutil.ReadFile(handler.config.DefaultAvatar); err == nil {
			return ec.Blob(http.StatusOK, "image/png", defaultAvatar)
		} else {
			return echo.ErrNotFound
		}
	}
	in := &In{}
	if err := ec.Bind(in); err != nil {
		return echo.ErrUnsupportedMediaType
	}
	level := avatar_keeper.KAvatarMiddle
	switch in.Standard {
	case "small":
		level = avatar_keeper.KAvatarSmall
	case "middle":
		level = avatar_keeper.KAvatarMiddle
	case "normal":
		level = avatar_keeper.KAvatarNormal
	default:
		break
	}
	data, err := handler.keeper.GetSingleAvatar(userId, level)
	if err != nil {
		return echo.ErrNotFound
	}

	return ec.Blob(http.StatusOK, "image/png", data)
}
