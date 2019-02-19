package numconvert

import (
	"strconv"
	"regexp"
	"math"
)

/**
	Transfer 数字转为中文金额
	number 	金额
	prec 	精度 小数点右边传负数，小数点左边传正数
 */
func Convert2China(number float64,prec int) string{
	var chinaMapping = map[string]string{"1":"壹", "2":"贰", "3":"叁", "4":"肆", "5":"伍", "6":"陆", "7":"柒", "8":"捌", "9":"玖", "0":"零"}
	var chinaUnit = []string{"分", "角", "圆", "拾", "佰", "仟", "万", "拾", "佰", "仟", "亿", "拾", "佰", "仟"}
	var chinaRegex = [][]string{
		{"零拾", "零"}, {"零佰", "零"}, {"零仟", "零"}, {"零零零", "零"}, {"零零", "零"},
		{"零角零分", "整"}, {"零分", "整"}, {"零角", "零"}, {"零亿零万零元", "亿元"},
		{"亿零万零元", "亿元"}, {"零亿零万", "亿"}, {"零万零元", "万元"}, {"万零元", "万元"},
		{"零亿", "亿"}, {"零万", "万"}, {"拾零圆", "拾元"}, {"零圆", "元"}, {"零零", "零"}}

	var numUpper,unitName,china string
	var strNum []byte


	if number == 0 {
		return "零"
	} else if number < 0 {
		china = "负"
	}


	number = float64(int(number/math.Pow(10, float64(prec)))) * float64(math.Pow(10, float64(prec)))
	strNum = []byte(strconv.FormatFloat(number, 'f', 3,64))
	unitIndex := len(strNum) - 2

	for _,v := range strNum {
		value := string(v)
		if unitIndex >= 1 && value != "." {
			if value == "0" {
				numUpper = "零"
			} else {
				numUpper = chinaMapping[value]
			}
			unitName = chinaUnit[unitIndex-1]
			china = china + numUpper + unitName
			unitIndex = unitIndex - 1
		}
	}

	for i := range chinaRegex {
		reg := regexp.MustCompile(chinaRegex[i][0])
		china = reg.ReplaceAllString(china, chinaRegex[i][1])
	}

	return china
}
