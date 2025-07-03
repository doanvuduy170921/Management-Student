package teacher

import (
	"awesomeProject/utils"
	"fmt"
)

func StudentMenu() {
	for {
		utils.ClearScreen()
		fmt.Println("==========QUAN LY SINH VIEN===========")
		fmt.Println("1. THEM SINH VIEN")
		fmt.Println("2. SUA THONG TIN SINH VIEN")
		fmt.Println("3. XOA SINH VIEN")
		fmt.Println("4. TIM KIEM SINH VIEN")
		fmt.Println("5. DANH SACH TAT CA SINH VIEN")
		fmt.Println("6. QUAY LAI")
		fmt.Println("CHON CHUC NANG :")
	}
}
