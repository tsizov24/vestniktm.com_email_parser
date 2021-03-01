package helper

import (
	"github.com/360EntSecGroup-Skylar/excelize"
)

const (
	baseUrl = "https://m.vestniktm.com/"
)

var (
	excelFile = excelize.NewFile()
	urlCats = []int32{
		1,    // 8
		3,    // 1231
		11,   // 12
		145,  // 3
		8,    // 27
		360,  // 3
		6,    // 17
		12,   // 50
		9,    // 91
		352,  // 2
		4,    // 12
		10,   // 9
		7,    // 8
		147,  // 1
		344,  // 3
		155,  // 1
		2,    // 11
		121,  // 11
	}
)
