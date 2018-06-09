package uolo_lens

import (
	"artemis/uolo_lens/model"
	"artemis/uolo_lens/utils"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type (
	ThingsHandler struct {
	}
)

func (handler *ThingsHandler) NewThings(ec echo.Context) error {
	type (
		NewThingsFrom struct {
			Content string `json:"content" form:"content"`
			Comment string `json:"comment" form:"comment"` //first comment
		}
		NewThingsResponse struct {
			Code    int32  `json:"code"`
			Message string `json:"message"`
		}
	)

	form := &NewThingsFrom{}
	err := ec.Bind(form)
	if err != nil {
		return ec.JSON(http.StatusOK, &NewThingsResponse{
			Code:    -1,
			Message: "表单错误",
		})
	}

	userId, login := utils.IsUserLogin(ec)
	if login == false {
		return ec.JSON(http.StatusOK, &NewThingsResponse{
			Code:    -2,
			Message: "登陆校验错误",
		})
	}

	thing := &model.UoloThing{
		Owner:       userId,
		Uuid:        utils.GenUUID(),
		Content:     form.Content,
		StartTime:   time.Now(),
		ArchiveGoal: false,
	}

	SaveErrorResponse := &NewThingsResponse{
		-3,
		"保存失败",
	}

	logrus.Debugln("New Things Comment:", form.Comment, " Content:", form.Content)

	session := Orm.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		logrus.Debugln("error:", err.Error())
		return ec.JSON(200, SaveErrorResponse)
	}

	if _, err := session.Insert(thing); err != nil {
		logrus.Debugln("error:", err.Error())
		return ec.JSON(200, SaveErrorResponse)
	}

	if len(form.Comment) != 0 { //如果有firs comment， 那么就新建一个comment 保存
		comment := &model.ThingsComment{
			ThingsId: thing.Uuid,
			Comment:  form.Comment,
		}
		if _, err := session.Insert(comment); err != nil {
			logrus.Debugln("error:", err.Error())
			return ec.JSON(200, SaveErrorResponse)
		}
	}

	if err := session.Commit(); err != nil {
		logrus.Debugln("error:", err.Error())
		return ec.JSON(200, SaveErrorResponse)
	}

	return ec.JSON(http.StatusOK, &NewThingsResponse{
		Code:    0,
		Message: "OK",
	})
}

func (handler *ThingsHandler) GetThingsList(ec echo.Context) error {

	type (
		Res struct {
			Code    int32              `json:"code"`
			Message string             `json:"message"`
			Things  []*model.UoloThing `json:"things,omitempty"`
		}
	)

	userId, login := utils.IsUserLogin(ec)
	if login == false {
		userId = "public_user"
	}

	things := make([]*model.UoloThing, 0)
	err := Orm.Where("owner = ?", userId).And("archive_goal=?", false).Find(&things)
	if err != nil {
		logrus.Debugln("err:", err.Error())
		return ec.JSON(200, &Res{
			Code:    -2,
			Message: "获取Things失败",
		})
	}

	return ec.JSON(200, &Res{
		Code:   0,
		Things: things,
	})
}

func (handler *ThingsHandler) GetDoneThingsList(ec echo.Context) error {
	type (
		Res struct {
			Code    int32              `json:"code"`
			Message string             `json:"message"`
			Things  []*model.UoloThing `json:"things,omitempty"`
		}
	)

	userId, login := utils.IsUserLogin(ec)
	if login == false {
		return ec.JSON(http.StatusOK, &Res{
			Code:    -1,
			Message: "登陆校验错误",
		})
	}

	things := make([]*model.UoloThing, 0)
	err := Orm.Where("owner = ?", userId).And("archive_goal=?", true).Find(&things)
	if err != nil {
		return ec.JSON(200, &Res{
			Code:    -2,
			Message: "获取Things失败",
		})
	}

	return ec.JSON(200, &Res{
		Code:   0,
		Things: things,
	})
}

func (handler *ThingsHandler) MarkUoloThingFinish(ec echo.Context) error {
	type (
		From struct {
			Uuid    string `json:"uuid" form:"uuid"`
			Comment string `json:"comment" form:"comment"`
		}
		DoneResponse struct {
			Code    int32  `json:"code"`
			Message string `json:"message"`
		}
	)
	Responser := func(errCode int32, msg string) error {
		return ec.JSON(200, &DoneResponse{
			Code:    errCode,
			Message: msg,
		})
	}

	form := &From{}
	err := ec.Bind(form)
	if err != nil {
		return Responser(-1, "参数错误")
	}

	userId, login := utils.IsUserLogin(ec)
	if login == false {
		return Responser(-2, "登陆信息错误")
	}

	result, err := Orm.Exec("update things set archive_goal = ? where things.owner = ? and things.uuid = ?", true, userId, form.Uuid)
	if err != nil {
		logrus.Errorf("Update UoloThing %s Failed in db", form.Uuid)
		return Responser(-3, "保存失败")
	}
	if lines, _ := result.RowsAffected(); lines == 0 {
		return Responser(-4, "数据不存在")
	}
	return Responser(0, "ok")
}

func (handler *ThingsHandler) DeleteUoloThing(ec echo.Context) error {
	type (
		From struct {
			Uuid string `json:"uuid" form:"uuid"`
		}
		DelRes struct {
			Code    int32  `json:"code"`
			Message string `json:"message"`
		}
	)

	form := &From{}
	err := ec.Bind(form)
	if err != nil {
		return ec.JSON(http.StatusOK, &DelRes{
			Code:    -1,
			Message: "表单错误",
		})
	}
	Responser := func(errCode int32, msg string) error {
		return ec.JSON(200, DelRes{
			Code:    errCode,
			Message: msg,
		})
	}

	userId, login := utils.IsUserLogin(ec)
	if login == false {
		return Responser(-2, "登陆校验错误")
	}

	session := Orm.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		logrus.Debugln("error:", err.Error())
		return Responser(-3, "删除失败，请重试")
	}

	thing := &model.UoloThing{
		Uuid:  form.Uuid,
		Owner: userId,
	}

	if lines, err := session.Where("uuid=?", form.Uuid).Delete(thing); err != nil {
		logrus.Debugln("error:", err.Error())
		return Responser(-4, "删除失败，请重试")
	} else if lines == 0 {
		return Responser(-4, "没有此记录")
	}
	comment := &model.ThingsComment{
		ThingsId: form.Uuid,
	}
	if _, err := session.Where("thing_id=?", form.Uuid).Delete(comment); err != nil {
		logrus.Debugln("error:", err.Error())
		return Responser(-5, "删除失败，请重试")
	}

	if err := session.Commit(); err != nil {
		logrus.Debugln("error:", err.Error())
		return Responser(-6, "删除失败，请重试")
	}

	return Responser(0, "ok")
}
