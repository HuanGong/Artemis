package model

type (
	LensList struct {
		Articles []*Article `json:"articles" xorm:"articles"`
	}
)
