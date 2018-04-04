package lunar

import (
	"testing"
	"fmt"
	"time"
)

func TestLunarToSolar(t *testing.T) {
	ln := Lunar{
		IsLeap: false,
		LunarYear: 1992,
		LunarMonth: 1,
		LunarDay: 15,
	}
	solar := LunarToSolar(ln)
	fmt.Printf("Solar Callendar Year:%d, Month:%d, Day:%d", solar.SolarYear, solar.SolarMonth, solar.SolarDay)
}

func TestSolarToLunar(t *testing.T) {
	solar := Solar{
		SolarDay:	18,
		SolarMonth: 2,
		SolarYear: 1992,
	}
	ln := SolarToLunar(solar)
	fmt.Printf("BirthDay:%d-%d-%d, Leap:%t\n",ln.LunarYear, ln.LunarMonth,ln.LunarDay, ln.IsLeap)

	now := time.Now()
	solar = Solar{
		SolarDay: now.Day(),
		SolarMonth: int(now.Month()),
		SolarYear: now.Year(),
	}
	ln = SolarToLunar(solar)
	fmt.Printf("Tody Lunar:%d-%d-%d, Leap:%t \n",ln.LunarYear, ln.LunarMonth,ln.LunarDay, ln.IsLeap)
}