package parse

import (
	"fmt"
	"strconv"
	"strings"
)

type PareseIn interface {
	Parse()
	GetMin() string
	GetDayOfMonth() string
	GetMonth() string
	GetHour() string
	GetDay() string
	GetDayofWeek() string
	GetCommand() string
}

var delimit = "*/"

type Data struct {
	input  string
	values map[string]string
}

func NewData(input string) PareseIn {
	return &Data{input: input}
}

func trim(str string) string {

	rdel := []rune(delimit)
	for _, val := range rdel {
		str = strings.Trim(str, fmt.Sprintf("%c", val))
	}
	return str
}

func (d *Data) Parse() {

	a := strings.Fields(d.input)

	d.values = make(map[string]string)

	// extract minutes
	minuteFactor, _ := strconv.Atoi(trim(a[0]))
	val := 0
	min := ""
	for i := 0; val < 45; i++ {
		val = i * minuteFactor
		min = min + strconv.Itoa(val) + " "
	}

	d.values["minute"] = min

	//extract hours

	d.values["hour"] = trim(a[1])

	//extract month

	d.values["day of month"] = trim(strings.Replace(a[2], ",", " ", -1))

	var month = ""
	for i := 1; i <= 12; i++ {
		month = month + strconv.Itoa(i) + " "
	}

	d.values["month"] = month

	// extract day

	day := trim(a[3])
	if day == "" {
		day = a[4]
	}
	sday, _ := strconv.Atoi(day[0:1])
	endday, _ := strconv.Atoi(day[2:])
	fmt.Println(sday, endday)

	days := ""
	for i := sday; i <= endday; i++ {
		days = days + strconv.Itoa(i) + " "
	}
	d.values["day of week"] = days

	d.values["command"] = a[5]

}

func (d *Data) GetMonth() string {
	return d.values["month"]

}
func (d *Data) GetDayOfMonth() string {
	return d.values["day of month"]

}

func (d *Data) GetDay() string {
	return d.values["day of week"]

}

func (d *Data) GetMin() string {
	return d.values["minute"]

}

func (d *Data) GetHour() string {
	return d.values["hour"]

}

func (d *Data) GetDayofWeek() string {
	return d.values["day of week"]

}

func (d *Data) GetCommand() string {
	return d.values["command"]

}
