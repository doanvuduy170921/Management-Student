package teacher

import (
	"awesomeProject/utils"
	"fmt"
)

var teachers []Teacher

func addTeacher() {
	fmt.Println("=======THÊM GIẢNG VIÊN==========")
	var id int
	for {
		id = utils.GetPositiveInt("Nhập id : ")
		if utils.IsContainId(id, teachers) {
			break
		}
		fmt.Println("⛔Id đã tồn tại! Vui lòng nhập Id khác!")
	}

	name := utils.GetNoneEmptyStr("Nhập tên : ")
	subject := utils.GetNoneEmptyStr("Nhập môn dạy : ")
	baseSalary := utils.GetPositiveFl("Nhập lương cơ bản : ")
	bonus := utils.GetPositiveFl("Nhập tiền thưởng : ")

	teacher := Teacher{
		Id:         id,
		Name:       name,
		Subject:    subject,
		BaseSalary: baseSalary,
		Bonus:      bonus,
	}
	teachers = append(teachers, teacher)
	fmt.Println("✅Thêm giảng viên thành công!")

}
func editTeacher() {
	fmt.Println("=======SỬA GIẢNG VIÊN==========")
	id := utils.GetPositiveInt("Nhập id giảng viên cần sửa :")
	for i, teacher := range teachers {
		if teacher.Id == id {
			name := utils.GetOptionalString(fmt.Sprintf("Nhập tên giảng viên cần sửa (%s) : ", teacher.Name), teacher.Name)
			subject := utils.GetOptionalString(fmt.Sprintf("Nhập môn học cần sửa (%s) : ", teacher.Subject), teacher.Subject)
			baseSalary := utils.GetOptionalFloat(fmt.Sprintf("Nhập lương cơ bản cần sửa (%.2f) : ", teacher.BaseSalary), teacher.BaseSalary)
			bonus := utils.GetOptionalFloat(fmt.Sprintf("Nhập tiền thưởng cần sửa (%.2f) : ", teacher.Bonus), teacher.Bonus)

			teachers[i] = Teacher{
				Id:         id,
				Name:       name,
				Subject:    subject,
				BaseSalary: baseSalary,
				Bonus:      bonus,
			}
			fmt.Println("✅Cập nhật giảng viên thành công!")
			return
		}

	}
	fmt.Println("⛔Không tìm thấy ID của giảng viên cần sửa!")

}
func deleteTeacher() {
	fmt.Println("=======XÓA GIẢNG VIÊN==========")
	id := utils.GetPositiveInt("Nhập ID của giảng viên cần xóa : ")
	for i, teacher := range teachers {
		if teacher.Id == id {
			teachers = append(teachers[:i], teachers[i+1:]...)
			return
		}
	}
	fmt.Println("⛔ID giảng viên không có trong danh sách!")
}
func findTeacher() {
	fmt.Println("=======TÌM GIẢNG VIÊN==========")
	id := utils.GetPositiveInt("Nhập ID của giảng viên cần tìm : ")
	for _, teacher := range teachers {
		if teacher.Id == id {
			teacher.GetInfo()
			return
		}
	}
	fmt.Println("⛔ID giảng viên không có trong danh sách!")
}
func listTeacher() {
	fmt.Println("=======DANH SÁCH GIẢNG VIÊN==========")
	if teachers == nil {
		fmt.Println("DANH SÁCH GIẢNG VIÊN RỖNG!")
		return
	}
	for _, val := range teachers {
		val.GetInfo()
	}
}

func MenuTeacher() {
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
			fmt.Println("⛔Bạn đã nhập sai! Vui lòng nhập lại!")
		}
		utils.ReadInput("Nhấn Enter để tiếp tục.....")
	}
}
