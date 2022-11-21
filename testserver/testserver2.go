package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func sansu(v1, v2, op string) (string, bool) {
	iv1, ve1 := strconv.Atoi(v1)
	iv2, ve2 := strconv.Atoi(v2)
	switch {
	case ve1 != nil:
		return "error bad value 1\n", false
	case ve2 != nil:
		return "error bad value 2\n", false
	case op == "":
		return "error no operand\n", false
	case !(iv1*iv1 > 0):
		return "error incrrect value 1\n", false
	case !(iv2*iv2 > 0):
		return "error increect value 2\n", false
	}
	switch op {
	case "+":
		return strconv.Itoa(iv1 + iv2), true
	case "-":
		return strconv.Itoa(iv1 - iv2), true
	default:
		return "error unknown operand\n", false
	}
}

func handle(c echo.Context) error {

	kekka, se := sansu(c.Param("v1"), c.Param("v2"), c.Param("op"))
	if se == true {
		fmt.Println(kekka)
		return c.Blob(http.StatusOK, "application/json",
			[]byte(fmt.Sprintf("{\"%v%v%v\":\"%v\"}\n", c.Param("v1"), c.Param("op"), c.Param("v2"), kekka)))
	} else {
		fmt.Println("error", kekka)
		return c.String(http.StatusBadRequest, kekka)
	}

	//	return c.Blob(http.StatusOK,"application/json",[]byte(fmt.Sprintf("{\"%v%v%v\":\"%v\"}",val1,ope,val2,res)))

}
func main() {
	e := echo.New()
	e.GET("/cal/:v1/:op/:v2", handle)
	e.Logger.Fatal(e.Start(":8000"))
}
