package student

import (
	"awesomeProject/utils"
	"fmt"
)

func addStudent() {
	fmt.Println("Them sinh vien")
}
func editStudent() {
	fmt.Println("Sua sinh vien")
}
func deleteStudent() {
	fmt.Println("Xoa sinh vien")
}
func findStudent() {
	fmt.Println("Tim sinh vien")
}
func listStudent() {
	fmt.Println("Danh sach sinh vien")
}

func StudentMenu() {
	for {
		utils.ClearScreen()
		fmt.Println("==========QUẢN LÝ SINH VIÊN===========")
		fmt.Println("1. THÊM SINH VIÊN")
		fmt.Println("2. SỬA THÔNG TIN SINH VIÊN")
		fmt.Println("3. XÓA SINH VIÊN")
		fmt.Println("4. TÌM KIẾM SINH VIÊN")
		fmt.Println("5. DANH SÁCH SINH VIÊN")
		fmt.Println("6. QUAY LẠI")

		chose := utils.ReadInput("CHỌN CHỨC NĂNG:")
		switch chose {
		case "1":
			addStudent()
		case "2":
			editStudent()
		case "3":
			deleteStudent()
		case "4":
			findStudent()
		case "5":
			listStudent()
		case "6":
			return
		default:
			fmt.Println("⛔Ban da nhap sai! Vui long nhap lai!")
		}
		utils.ReadInput("Nhan Enter de tiep tuc ....")
	}
}
