package model

type (
	MemCalendar struct {
		Id       int64  `json:"id" xorm:"id"`           // unique id
		Type     int32  `json:"type" xorm:"type"`       // 0: solar type, 1: lunar
		FireTime int64  `json:"feature" xorm:"feature"` // trick time by day
		Msg      string `json:"msg" xorm:"msg"`         // content
	}
)

func (posts *MemCalendar) GetName() string {
	return "calendar"
}
