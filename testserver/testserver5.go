//testserver5

package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type keisan struct {
	Fo string
	V1 string
	V2 string
	Op string
	Re string
}

func sansu(k keisan) keisan {
	v1, er1 := strconv.ParseFloat(k.V1, 64)
	v2, er2 := strconv.ParseFloat(k.V2, 64)
	k.Fo = k.V1 + k.Op + k.V2

	switch {
	case er1 != nil || er2 != nil:
		k.Re = "bad value"
	case k.Op == "+":
		k.Re = fmt.Sprintf("%v", v1+v2)
	case k.Op == "-":
		k.Re = fmt.Sprintf("%v", v1-v2)
	default:
		k.Re = "unknown operand"
	}
	fmt.Printf("%v\t%v\t%v\t%v\n",v1,v2,k.Fo,k)
	return k
}

func gethandle(e echo.Context) error {
	kekka := keisan{
		V1: e.Param("v1"),
		V2: e.Param("v2"),
		Op: e.Param("op"),
	}
	
	kekka = sansu(kekka)
	m := make(map[string]string)
	m[kekka.Fo] = kekka.Re

	if jf, err := json.Marshal(m); err == nil {
		return e.Blob(http.StatusOK, "application/json", jf)
		//return e.Blob  (http.StatusOK,"application/json",
		//[]byte(fmt.Sprintf("{\"%v\":\"%v\"}\n",kekka.Fo,kekka.Re)))
	} else {
		return e.String(http.StatusBadRequest, kekka.Re)
	}

}
func poshandle(e echo.Context) error {
	var kekka keisan
	if err := e.Bind(&kekka); err != nil {
		return e.String(http.StatusBadRequest, "Binding error")
	}

	kekka = sansu(kekka)
	m := make(map[string]string)
	m[kekka.Fo] = kekka.Re

	if jf, err := json.Marshal(m); err == nil {
		return e.Blob(http.StatusOK, "application/json", jf)
	} else {
		return e.String(http.StatusBadRequest, kekka.Re)
	}
}

func main() {
	e := echo.New()
	e.GET("/cal/:v1/:op/:v2", gethandle)
	e.POST("/cal", poshandle)
	e.Logger.Fatal(e.Start(":8000"))
}
