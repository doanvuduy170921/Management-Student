package teacher

import (
	"awesomeProject/utils"
	"fmt"
)

func addTeacher() {
	fmt.Println("Them sinh vien")
}
func editTeacher() {
	fmt.Println("Sua sinh vien")
}
func deleteTeacher() {
	fmt.Println("Xoa sinh vien")
}
func findTeacher() {
	fmt.Println("Tim sinh vien")
}
func listTeacher() {
	fmt.Println("Danh sach sinh vien")
}

func TeacherMenu() {
	for {
		utils.ClearScreen()
		fmt.Println("==========QUẢN LÝ GIẢNG VIÊN===========")
		fmt.Println("1. THÊM GIẢNG VIÊN")
		fmt.Println("2. SỬA THÔNG TIN GIẢNG VIÊN")
		fmt.Println("3. XÓA GIẢNG VIÊN")
		fmt.Println("4. TÌM KIẾM GIẢNG VIÊN")
		fmt.Println("5. DANH SÁCH GIẢNG VIÊN")
		fmt.Println("6. QUAY LẠI")

		chose := utils.ReadInput("CHỌN CHỨC NĂNG:")
		switch chose {
		case "1":
			addTeacher()
		case "2":
			editTeacher()
		case "3":
			deleteTeacher()
		case "4":
			findTeacher()
		case "5":
			listTeacher()
		case "6":
			return
		default:
			fmt.Println("⛔Ban da nhap sai! Vui long nhap lai!")
		}
		utils.ReadInput("Nhan Enter de tiep tuc ....")
	}
}
