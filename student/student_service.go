package student

import (
	"awesomeProject/utils"
	"fmt"
)

var students []Student

func addStudent() {
	var scores []float64
	var id int
	fmt.Println("=========THÊM SINH VIÊN==========")

	for {
		id = utils.GetPositiveInt("Nhập Id: ")
		if utils.IsContainId(id, students) {
			break
		}
		fmt.Println("ID đã tồn tại! Vui lòng nhập ID khác!")
	}

	name := utils.GetNoneEmptyStr("Nhập Name: ")
	class := utils.GetNoneEmptyStr("Nhập lớp: ")
	countScore := utils.GetPositiveInt("Nhập số lượng điểm: ")

	for i := 1; i <= countScore; i++ {
		score := utils.GetPositiveFl(fmt.Sprintf("Nhập điểm thứ %d :", i))
		scores = append(scores, score)
	}
	student := Student{
		Id:     id,
		Name:   name,
		Class:  class,
		Scores: scores,
	}
	students = append(students, student)
	fmt.Println("✅Thêm sinh viên thành công!")

}
func editStudent() {
	fmt.Println("=======SỬA SINH VIÊN==========")
	id := utils.GetPositiveInt("Nhập id cần sửa : ")
	for i, s := range students {
		if s.Id == id {
			name := utils.GetOptionalString(fmt.Sprintf("Nhập tên sinh viên cần sửa (%s)", s.Name), s.Name)
			class := utils.GetOptionalString(fmt.Sprintf("Nhập lớp cần sửa (%s)", s.Class), s.Class)

			newScores := make([]float64, len(s.Scores))
			for idx, val := range s.Scores {
				temp := fmt.Sprintf("Nhập điểm thứ %d (%.2f)", idx+1, val)
				newScores[idx] = utils.GetOptionalFloat(temp, val)
			}
			students[i] = Student{
				Id:     id,
				Name:   name,
				Class:  class,
				Scores: newScores,
			}
			fmt.Println("✅Cập nhật sinh viên thành công!")
			return
		}
	}
	fmt.Println("⛔Không tìm thấy ID của sinh viên cần sửa!")

}
func deleteStudent() {
	fmt.Println("=======XÓA SINH VIÊN==========")
	id := utils.GetPositiveInt("Nhập ID sinh viên cần xóa :")
	for i, s := range students {
		if s.Id == id {
			students = append(students[:i], students[i+1:]...)
			fmt.Println("✅Xóa sinh viên thành công!")
			return
		}
	}
	fmt.Println("⛔Không tìm thấy ID của sinh viên cần xóa!")

}
func findStudent() {
	fmt.Println("=======TÌM SINH VIÊN==========")
	id := utils.GetPositiveInt("Nhập ID cần tìm :")
	for _, s := range students {
		if s.Id == id {
			s.GetInfo()
			return
		}
	}
	fmt.Println("⛔Không tìm thấy ID của sinh viên cần tìm!")

}
func listStudent() {
	fmt.Println("=======DANH SÁCH SINH VIÊN==========")
	if students == nil {
		fmt.Println("DANH SÁCH RỖNG!")
		return
	}
	for _, student := range students {
		student.GetInfo()
	}
}

func Menu() {
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
			fmt.Println("⛔Bạn đã nhập sai! Vui lòng nhập lai!")
		}
		utils.ReadInput("\nNhấn Enter de tiếp tục ....")
	}
}
