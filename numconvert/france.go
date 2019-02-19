package numconvert

import (
	"math"
	"strconv"
	"strings"
)

func Convert2France(number float64, prec int) string {
	strNum := strconv.FormatFloat(number, 'f', 2, 64)
	arrNum := strings.Split(strNum, ".")
	change,_ := strconv.ParseFloat(arrNum[1],  64)
	bigPart,_:= strconv.ParseFloat(arrNum[0],64)
	if  change > 0 {
		return Convert2France(bigPart, prec) + "Dinars" + " et " + Convert2France(change, prec) + " Centimes"
	}

	if number < 0 {
		return "moins" + Convert2France(math.Abs(number),prec)
	}


	number = float64(int(number/math.Pow(10, float64(prec)))) * float64(math.Pow(10, float64(prec)))
	prec = 0

	if number < 17 {
		switch number {
		case 0 :
		return "zero"
		case 1 :
		return "un"
		case 2 :
		return "deux"
		case 3 :
		return "trois"
		case 4 :
		return "quatre"
		case 5 :
		return "cinq"
		case 6 :
		return "six"
		case 7 :
		return "sept"
		case 8 :
		return "huit"
		case 9 :
		return "neuf"
		case 10 :
		return "dix"
		case 11 :
		return "onze"
		case 12 :
		return "douze"
		case 13 :
		return "treize"
		case 14 :
		return "quatorze"
		case 15 :
		return "quinze"
		case 16 :
		return "seize"
		}
	} else if number < 20 {
		return "dix-" + Convert2France(number - 10, prec)
	} else if number < 100 {
		if int(number) % 10 == 0 {
			switch number {
			case 20 :
			return "vingt"
			case 30 :
			return "trente"
			case 40 :
			return "quarante"
			case 50 :
			return "cinquante"
			case 60 :
			return "soixante"
			case 70 :
			return "soixante-dix"
			case 80 :
			return "quatre-vingt"
			case 90 :
			return "quatre-vingt-dix"
			}
		} else if  number== 21 || number== 31 ||number== 41 ||number== 51 ||number== 61 ||number== 71 ||number== 81 ||number== 91{
			if int(number / 10) * 10 < 70 {
				return Convert2France(float64(int(number / 10) * 10), prec) + "-et-un"
			} else if number == 71 {
				return "soixante-et-onze"
			} else if number == 81 {
				return "quatre-vingt-un"
			} else if number == 91 {
				return "quatre-vingt-onze"
			}
		} else if number < 70 {
			return Convert2France(number - float64(int(number) % 10), prec) + "-" + Convert2France(float64(int(number) % 10), prec)
		} else if number < 80 {
			return Convert2France(60, prec) + "-" + Convert2France(float64(int(number)%20), prec)
		} else {
			return Convert2France(80, prec) + "-" + Convert2France(float64(int(number) % 20), prec)
		}
	} else if number == 100 {
		return "cent"
	} else if number < 200 {
		return Convert2France(100, prec) + " " + Convert2France(float64(int(number)%100), prec)
	}else if number < 1000 {
		if int(number)%100 != 0 {
			return Convert2France(float64(int(number/100)), prec) + " cent " + Convert2France(float64(int(number)%100), prec)
		} else {
			return Convert2France(float64(int(number/100)), prec) + " cent "
		}
	} else if number == 1000 {
		return "mille"
	} else if number < 2000 {
		return Convert2France(1000, prec) + " " + Convert2France(float64(int(number) % 1000), prec)
	} else if number < 1000000 {
		if int(number) % 1000 != 0 {
			return Convert2France(float64(int(number/1000)), prec) + " mille " + Convert2France(float64(int(number) % 1000), prec)
		} else {
			return Convert2France(float64(int(number/1000)), prec) + " mille "
		}
	} else if number == 1000000 {
		return "millions"
	} else if number < 2000000 {
		return Convert2France(1000000, prec) + " " + Convert2France(float64(int(number) % 1000000), prec)
	} else if number < 1000000000 {
		if int(number) % 1000000 != 0 {
			return Convert2France(float64(int(number/1000000)), prec) + " millions " + Convert2France(float64(int(number) % 1000000), prec)
		} else {
			return Convert2France(float64(int(number/1000000)), prec) + " millions "
		}
	} else if number == 1000000000{
		return "milliard"
	} else if number < 2000000000{
		return Convert2France(1000000000, prec) + " " + Convert2France(float64(int(number) % 1000000000), prec)
	} else {
		if int(number) % 1000000000 != 0 {
			return Convert2France(float64(int(number/1000000)), prec) + " milliard " + Convert2France(float64(int(number) % 1000000000), prec)
		} else {
			return Convert2France(float64(int(number/1000000)), prec) + " milliard "
		}

	}

	return ""
}
