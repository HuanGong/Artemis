package avatar_data

import (
	"fmt"
	"testing"
)

func TestGetWeather(t *testing.T) {
	w, err := GetWeather("南京", "afdb952ba7aa4a419acd05e754587b17")
	fmt.Println(w, err)
}
