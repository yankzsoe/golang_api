package tools

import (
	"strconv"
)

type Conversion struct{}

func (c *Conversion) StrToInt(p string) (r int) {
	if p == "" {
		p = "1"
	}

	val, err := strconv.Atoi(p)
	if err != nil {
		panic(err.Error())
	}
	r = val
	return r
}
