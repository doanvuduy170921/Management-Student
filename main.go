package main

import (
	"awesomeProject/student"
	"awesomeProject/teacher"
	"awesomeProject/utils"
	"fmt"
)

func MenuManagement() {
	for {
		utils.ClearScreen()

		fmt.Println("📔 CHƯƠNG TRÌNH QUẢN LÝ")
		fmt.Println("🎓1.QUẢN LÝ SINH VIÊN")
		fmt.Println("🧙2.QUẢN LÝ GIẢNG VIÊN")
		fmt.Println("◀️3.THOÁT")

		chose := utils.ReadInput("CHỌN CHỨC NĂNG:")

		switch chose {
		case "1":
			student.Menu()
			continue
		case "2":
			teacher.MenuTeacher()
			continue
		case "3":
			return
		default:

			fmt.Println("⛔Bạn đã nhập sai! Vui lòng nhập lại!")
		}
		utils.ReadInput("Nhấn Enter để tiếp tục.....")
	}
}

func main() {
	MenuManagement()

}
