package numconvert

import (
	"strconv"
	"strings"
	"math"
)

/**
	EUnit 英文金额基本单位
 */
type EUnit struct {
	thousand 	string		//千位
	hundred		string		//百位
	ten 		string		//十位
	one 		string		//个位
	unit		string		//单位
}


type EnlishAmount struct {
	trillion 	*EUnit
	billion 	*EUnit
	million 	*EUnit
	thousand	*EUnit
	one 		*EUnit
}



func (eunit *EUnit) toString() string {
	var oneMapping = map[string]string{"1":"one", "2":"two", "3":"three", "4":"four", "5":"five", "6":"six", "7":"seven", "8":"eight", "9":"nine", "0":"zero"}
	var tenMapping = map[string]string{"2":"twenty", "3":"thirty", "4":"forty", "5":"fifty", "6":"sixty", "7":"seventy", "8":"eighty", "9":"ninety", "0":""}
	var tenOneMapping = map[string]string{"10":"ten", "11":"eleven", "12":"twelve", "13":"thirteen", "14":"fourteen", "15":"fifteen", "16":"sixteen", "17":"seventeen", "18":"eighteen", "19":"nineteen"}
	var thousandName, hundredName, tenName, oneName string

	if eunit != nil {
		if eunit.thousand != "" {
			thousandName = oneMapping[eunit.thousand] + " thousand"
		}
		if eunit.hundred != "" && eunit.hundred != "0" {
			hundredName = oneMapping[eunit.hundred] + " hundred"
		}
		if eunit.ten != "" && eunit.ten != "0" {
			if eunit.ten == "1" {
				data := eunit.ten + eunit.one
				name := tenOneMapping[data]
				if hundredName != "" {
					return thousandName + " " + hundredName+" and " + name +" " +eunit.unit
				} else {
					return thousandName + " " + name +" " +eunit.unit
				}
			} else {
				tenName = tenMapping[eunit.ten]
				oneName = oneMapping[eunit.one]
				if hundredName != "" {
					if oneName == "zero" {
						return thousandName +" " + hundredName+" and " + tenName+" " + eunit.unit
					} else {
						return thousandName+" " + hundredName+" and " + tenName+"-"+oneName+" " + eunit.unit
					}
				} else {
					if oneName == "zero" {
						return thousandName +" "+ tenName+" " + eunit.unit
					} else {
						return thousandName+" " + tenName+"-"+oneName+" " + eunit.unit
					}
				}

			}

		}
		if eunit.one != "" && eunit.one != "0" {
			oneName = " " + oneMapping[eunit.one] + " "
		}

		return thousandName + hundredName + tenName + oneName + eunit.unit
	}

	return ""
}






func (ea *EnlishAmount) toString() string {
	return ea.trillion.toString() + ea.billion.toString() + ea.million.toString() + ea.thousand.toString() + ea.one.toString()
}




func formUnit(data []byte, unitName string) *EUnit {
	unit := &EUnit{}
	if len(data) == 1 {
		unit.one = string(data[0])
	} else if len(data) == 2 {
		unit.ten = string(data[0])
		unit.one = string(data[1])
	} else if len(data) == 3 {
		unit.hundred = string(data[0])
		unit.ten = string(data[1])
		unit.one = string(data[2])
	} else if len(data) == 4 {
		unit.thousand = string(data[0])
		unit.hundred = string(data[1])
		unit.ten = string(data[2])
		unit.one = string(data[3])
	}

	unit.unit = unitName
	return unit
}





/**
	Transfer 数字转为中文金额
	number 	金额
	prec 	精度 小数点右边传负数，小数点左边传正数
 */
func Convert2Enlish(number float64,prec int) string{

	var enlishUnit = []string{"dime", "cent"}

	var symbol string
	var strNum []byte

	if number == 0 {
		return "zero"
	} else if number < 0 {
		symbol = "negative "
	}

	number = float64(int(number/math.Pow(10, float64(prec)))) * float64(math.Pow(10, float64(prec)))
	strNum = []byte(strconv.FormatFloat(number, 'f', 2,64))

	pointIndex := strings.Split(string(strNum),".")
	changePart := pointIndex[1]
	bigPart    := []byte(pointIndex[0])

	bigLen     := len(bigPart)
	var Result = &EnlishAmount{}

	if bigLen < 4 {
		unit := formUnit(bigPart,"")
		Result.one = unit
	} else if bigLen < 7 {
		Result.one 		= formUnit(bigPart[bigLen-3:], "")
		Result.thousand = formUnit(bigPart[:bigLen-3], "thousand")

	} else if bigLen < 10 {
		Result.one 		= formUnit(bigPart[bigLen-3:], "")
		Result.thousand = formUnit(bigPart[bigLen-6:bigLen-3], "thousand")
		Result.million  = formUnit(bigPart[:bigLen-6], "million")
	} else if bigLen < 14 {
		Result.one 		= formUnit(bigPart[bigLen-3:], "")
		Result.thousand = formUnit(bigPart[bigLen-6:bigLen-3], "thousand")
		Result.million 	= formUnit(bigPart[bigLen-9:bigLen-6], "million")
		Result.billion 	= formUnit(bigPart[:bigLen-9], "billion")
	} else {
		Result.one 		= formUnit(bigPart[bigLen-3:], "")
		Result.thousand = formUnit(bigPart[bigLen-6:bigLen-3], "thousand")
		Result.million  = formUnit(bigPart[bigLen-9:bigLen-6], "million")
		Result.billion  = formUnit(bigPart[bigLen-13:bigLen-9], "billion")
		Result.trillion = formUnit(bigPart[:bigLen-13], "trillion")
	}

	var changeName string
	var oneMapping = map[string]string{"1":"one", "2":"two", "3":"three", "4":"four", "5":"five", "6":"six", "7":"seven", "8":"eight", "9":"nine", "0":"zero"}
	for i,v := range changePart {
		data := oneMapping[string(v)]
		dataName := enlishUnit[i]
		changeName += data + " " + dataName + " "
	}

	return symbol + " " + Result.toString() + "dollar " + changeName
}
