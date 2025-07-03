package main

import (
	"awesomeProject/teacher"
	"awesomeProject/utils"
	"fmt"
)

func main() {
	for {
		utils.ClearScreen()
		fmt.Println("📔 CHUONG TRINH QUAN LY")
		fmt.Println("🧑‍🎓1.QUAN LY SINH VIEN")
		fmt.Println("🧙2.QUAN LY GIANG VIEN")
		fmt.Println("◀️3.THOAT")

		chose := utils.ReadInput("CHON CHUC NANG:")

		switch chose {
		case "1":
			teacher.StudentMenu()
		case "2":
			fmt.Println("quan giang vien")
		case "3":
			return
		default:
			fmt.Println("⛔Ban da nhap sai! Vui long nhap lai!")
		}
		utils.ReadInput("Nhan Enter de tiep tuc ....")
	}

}
