package util

import "testing"

func TestStringToTime(t *testing.T) {
	var convertedTime = StringToTime("2019-02-14T18:43:23")
	if convertedTime.Year() != 2019 {
		t.Errorf("Espera que ano seja igual à 2019")
	}
}
