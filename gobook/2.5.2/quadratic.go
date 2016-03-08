package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
)

type Solution struct {
	a  float64
	b  float64
	c  float64
	x1 complex128
	x2 complex128
}

func complexToString(c complex128) string {
	if imag(c) == 0.0 {
		return fmt.Sprintf("%v", real(c))
	}
	return fmt.Sprintf("%v", c)
}

func (s Solution) String() string {
	emptySolutions := Solution{}
	if s == emptySolutions {
		return ""
	}

	result := bytes.Buffer{}
	result.WriteString(fmt.Sprintf("%vx<sup>2</sup>", s.a))

	if s.b > 0 {
		result.WriteString(fmt.Sprintf("+%vx", s.b))
	} else if s.b < 0 {
		result.WriteString(fmt.Sprintf("%vx", s.b))
	}

	if s.c > 0 {
		result.WriteString(fmt.Sprintf("+%v", s.c))
	} else if s.c < 0 {
		result.WriteString(fmt.Sprintf("%v", s.c))
	}

	result.WriteString("=0 <br/>")

	result.WriteString(fmt.Sprintf("x1=%v", complexToString(s.x1)))
	if s.x1 != s.x2 {
		result.WriteString(fmt.Sprintf("<br/>x2=%v", complexToString(s.x2)))
	}

	return result.String()
}

type PageData struct {
	Result       Solution
	ErrorMessage template.HTML
}

func solveQuadratic(a, b, c float64) (x1, x2 complex128) {
	d := math.Pow(b, 2) - 4*a*c
	if d >= 0 {
		x1 = complex(-b/(2*a)+math.Sqrt(d)/(2*a), 0)
		x2 = complex(-b/(2*a)-math.Sqrt(d)/(2*a), 0)
	} else {
		d = math.Abs(d)
		x1 = complex(-b/(2*a), +math.Sqrt(d)/(2*a))
		x2 = complex(-b/(2*a), -math.Sqrt(d)/(2*a))
	}
	return
}

func parseFloatParam(paramName string, request *http.Request) (value float64, err error) {
	if values, found := request.Form[paramName]; found && len(values) > 0 {
		value, err = strconv.ParseFloat(values[0], 64)
	}
	return value, err
}

func parseForm(request *http.Request) (a, b, c float64, err error) {
	if a, err = parseFloatParam("a", request); err != nil {
		return
	}
	if b, err = parseFloatParam("b", request); err != nil {
		return
	}
	if c, err = parseFloatParam("c", request); err != nil {
		return
	}
	return
}

func processRequest(request *http.Request) (Solution, error) {
	a, b, c, err := parseForm(request)
	if err != nil || len(request.Form) == 0 {
		return Solution{}, err
	}

	if a == 0.0 {
		return Solution{}, fmt.Errorf("Parameter 'a' can't be 0")
	}

	x1, x2 := solveQuadratic(a, b, c)
	result := Solution{a, b, c, x1, x2}
	return result, nil
}

func homePageHandler(writer http.ResponseWriter, request *http.Request) {
	var errorMessage string
	var result Solution

	err := request.ParseForm()
	if err == nil {
		result, err = processRequest(request)
	}
	if err != nil {
		errorMessage = fmt.Sprintf("Error: %s", err)
	}

	funcMap := template.FuncMap{
		"unescape": func(s fmt.Stringer) template.HTML {
			return template.HTML(s.String())
		},
	}

	var t *template.Template
	if t, err = template.New("index.html").Funcs(funcMap).ParseFiles("templates/index.html"); err != nil {
		fmt.Fprint(writer, "Failed to load template: %s", err)
		return
	}

	t.Execute(writer, &PageData{Result: result, ErrorMessage: template.HTML(errorMessage)})
}

func main() {
	http.HandleFunc("/", homePageHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
