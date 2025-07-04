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

		fmt.Println("📔 CHUONG TRINH QUAN LY")
		fmt.Println("🎓1.QUAN LY SINH VIEN")
		fmt.Println("🧙2.QUAN LY GIANG VIEN")
		fmt.Println("◀️3.THOAT")

		chose := utils.ReadInput("CHON CHUC NANG:")

		switch chose {
		case "1":
			student.StudentMenu()
			continue
		case "2":
			teacher.TeacherMenu()
			continue
		case "3":
			return
		default:

			fmt.Println("⛔Ban da nhap sai! Vui long nhap lai!")
		}
		utils.ReadInput("Nhan Enter de tiep tuc ....")
	}
}

func main() {
	MenuManagement()

}
