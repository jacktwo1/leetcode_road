package test

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InputOutput() {
	inputs := bufio.NewScanner(os.Stdin)
	for inputs.Scan() { //每次读入一行
		data := strings.Split(inputs.Text(), " ") //通过空格将他们分割，并存入一个字符串切片
		var sum int
		for i := range data {
			val, _ := strconv.Atoi(data[i]) //将字符串转换为int
			sum += val
		}
		fmt.Println(sum)
	}
}
