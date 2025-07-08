package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func ReadInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func GetPositiveInt(prompt string) int {
	for {
		val, err := strconv.Atoi(ReadInput(prompt))
		if err == nil && val > 0 {
			return val
		}
		fmt.Println("❌ Gia tri khong hop le! Hay nhap so nguyen duong.")
	}
}

func GetOptionalString(prompt string, oldValue string) string {
	input := ReadInput(prompt)
	if input == "" {
		return oldValue
	}
	return input
}

func GetOptionalFloat(prompt string, oldValue float64) float64 {
	input := ReadInput(prompt)
	if input == "" {
		return oldValue
	}
	val, err := strconv.ParseFloat(input, 64)
	if err != nil && val < 0 {
		fmt.Println("Giá trị không hợp lệ! Giữ nguyên giá trị cũ")
		return oldValue
	}
	return val
}

func GetPositiveFl(prompt string) float64 {
	for {
		val, err := strconv.ParseFloat(ReadInput(prompt), 64)
		if err == nil && val > 0 {
			return val
		}
		fmt.Println("❌ Giá trị không hợp lệ! Hãy nhập số thực dương!.")
	}
}

func GetNoneEmptyStr(prompt string) string {
	for {
		input := ReadInput(prompt)
		if input != "" {
			return input
		}
		fmt.Println("❌ Giá trị không hợp lệ! Vui lòng không để trống!")
	}
}

type HasId interface {
	GetId() int
}

func IsContainId[T HasId](id int, list []T) bool {
	for _, st := range list {
		if st.GetId() == id {
			return false
		}
	}
	return true
}

func ClearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("⛔Error clearing screen :", err)
	}
}
