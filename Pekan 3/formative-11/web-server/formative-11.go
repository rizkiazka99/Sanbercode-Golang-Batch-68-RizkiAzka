package webserver

import "math"

func VolumeTabung(r int, t int) float64 {
	return math.Pi * math.Pow(float64(r), 2) * float64(t)
}

func LuasAlasTabung(r int) float64 {
	return math.Pi * math.Pow(float64(r), 2)
}

func KelilingLuasAlasTabung(r int) float64 {
	return 2 * math.Pi * float64(r)
}
