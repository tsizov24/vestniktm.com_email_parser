package helper

import "fmt"

var (
	cnt = 0
)

func setCellValue(email string) {
	cnt++
	notFatalNotification(excelFile.SetCellValue("Sheet1", fmt.Sprintf("A%d", cnt), listUrl))
	notFatalNotification(excelFile.SetCellValue("Sheet1", fmt.Sprintf("B%d", cnt), adUrl))
	notFatalNotification(excelFile.SetCellValue("Sheet1", fmt.Sprintf("C%d", cnt), email))
}

func saveFile() {
	notFatalNotification(excelFile.SaveAs("emails.xlsx"))
}
