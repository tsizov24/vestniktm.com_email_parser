package helper

import (
	"bytes"
	"strconv"
)

func encode(protected string) string {
	var e bytes.Buffer
	r, _ := strconv.ParseInt(protected[0:2], 16, 0)
	for n := 4; n < len(protected)+2; n += 2 {
		i, _ := strconv.ParseInt(protected[n-2:n], 16, 0)
		e.WriteString(string(i ^ r))
	}
	return e.String()
}
