package models

import "math"

var DB []Vars

type Vars struct {
	A      int
	B      int
	C      int
	Nroots int
}

func (vars *Vars) Check() {
	var d int = int(math.Pow(float64(vars.B), 2)) - (4 * vars.A * vars.C)
	if d > 0 {
		vars.Nroots = 2
	} else if d == 0 {
		vars.Nroots = 1
	} else {
		vars.Nroots = 0
	}
	DB = append(DB, *vars)
}
