package main

import (
	"errors"
	"fmt"
	"strconv"
)

type MyBigNumber struct {
}

func (m MyBigNumber) sum(strNum1 string, strNum2 string) error {
	// kiểm tra strNum1 và strNum2 có phải kiểu số không
	_, err := strconv.ParseUint(strNum1, 10, 64)
	if err != nil {
		err = errors.New("chuỗi thứ nhất không phải kiểu số")
		return err
	}
	_, err = strconv.ParseUint(strNum2, 10, 64)
	if err != nil {
		err = errors.New("chuỗi thứ hai không phải kiểu số")
		return err
	}

	length1 := len(strNum1)
	length2 := len(strNum2)
	max := length1
	min := length2
	if length1 < length2 {
		max = length2
		min = length1
	}
	stringZero := createString(max - min)

	if length1 == min {
		strNum1 = stringZero + strNum1
	} else {
		strNum2 = stringZero + strNum2
	}
	result := ""
	remember := 0
	for i := max - 1; i >= 0; i-- {
		ki1 := strNum1[i] - 48
		ki2 := strNum2[i] - 48
		total := ki1 + ki2
		totalKi := int(total) % 10
		if total >= 10 {
			if remember == 1 {
				fmt.Printf("Bước %d, lấy %d cộng %d được %d cộng 1 được %d, ghi %d, nhớ 1", max-i, ki1, ki2, total, total+1, total%10)
				totalKi = totalKi + 1
			} else {
				fmt.Printf("Bước %d, lấy %d cộng %d được %d, ghi %d, nhớ 1", max-i, ki1, ki2, total, total%10)
			}
			fmt.Println()
		} else {
			if remember == 1 {
				if total+1 == 10 {
					totalKi = 0
					fmt.Printf("Bước %d, lấy %d cộng %d được %d cộng 1 được %d, ghi %d nhớ 1", max-i, ki1, ki2, total, total+1, (total+1)%10)
					total = 10
				} else {
					fmt.Printf("Bước %d, lấy %d cộng %d được %d cộng 1 được %d, ghi %d", max-i, ki1, ki2, total, total+1, total+1)
					totalKi = totalKi + 1
				}
			} else {
				fmt.Printf("Bước %d, lấy %d cộng %d được %d, ghi %d", max-i, ki1, ki2, total, total)
			}
			fmt.Println()
		}
		remember = int(total) / 10
		s := strconv.Itoa(totalKi)
		result = s + result
	}
	if remember == 1 {
		result = "1" + result
	}
	fmt.Printf("Tổng hai số là: %s", result)
	fmt.Println()
	return nil
}

func createString(length int) (str string) {
	for i := 0; i < length; i++ {
		str += "0"
	}
	return str
}

func main() {
	fmt.Println("=================Start================")
	fmt.Print("Nhập số thứ nhất: ")
	var strNum1, strNum2 string
	_, errScanln := fmt.Scanln(&strNum1)
	if errScanln != nil {
		return
	}
	fmt.Print("Nhập số thứ hai: ")
	_, errScanln = fmt.Scanln(&strNum2)
	if errScanln != nil {
		return
	}
	c := MyBigNumber{}
	err := c.sum(strNum1, strNum2)
	if err != nil {
		fmt.Println("Lỗi: ", err.Error())
	}
	fmt.Println("=================END==================")
}
