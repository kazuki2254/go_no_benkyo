//

package main 

import (
		"net/http"
		"github.com/labstack/echo/v4"
		"fmt"
		"strconv"
		"encoding/json"
	)

type resj struct{
		Formula string
		Result	int
}
type fml struct{
	Val int
	Vall int
	Ope string
	
}
func sansu(a,b int,s string)int{
	switch s{
		case "-":
			return a-b
		case "+":
			return a+b
		default:
			return 0
	}
}

func encj(r resj)[]byte{
	m :=make(map[string]int)

	m=[r.Formula]r.Result
	j,err :=json.Marshal(m)
	if err != nil{
		}
	return  j
}
func getcalcu(c echo.Context)error{
	
		a,h:= strconv.Atoi(c.Param("id"))
		d,j:= strconv.Atoi(c.Param("ig"))
		b := c.Param("ie")
		rj :=resj{		
			Formula:strconv.Itoa(a) + b + strconv.Itoa(d),
		}
		fmt.Println("strconv OK")

		if( h!=nil || b=="" || j!=nil || (a*a<=0 && d*d<=0 ) ){
			//error prc
			goto err
		}
		
		

		rj.Result=sansu(a,d,b)

		fmt.Println("call calcu")
		fmt.Printf("param is\t%v\t%v\t%v \ntype is \t%T\t%T\t%T\nres=%d\n",a,b,d,a,b,d,rj.Result)
		//return c.String(http.StatusOK,fmt.Sprintf("param is\t%v\t%v\t%v \ntype is \t%T\t%T\t%T\nres=%d\n",a,b,d,a,b,d,rj.Result))
		return c.JSON(http.StatusOK,rj)
		
		//error prc
		err:
		fmt.Println("error")
		return c.String(http.StatusBadRequest,fmt.Sprintln(rj))
}

func postcalcu(c echo.Context)error{
	
	var (
		vv fml
		rj resj
	)
	if err:=c.Bind(&vv);err!=nil{
		//error prc
	}
		rj.Formula= strconv.Itoa(vv.Val) + vv.Ope + strconv.Itoa(vv.Vall)
		switch vv.Ope{
			case "-":
				rj.Result=vv.Val+vv.Vall
			case "+":
				rj.Result=vv.Val+vv.Vall
			default:
				//error prc
				goto err
			}
	return c.JSON(http.StatusOK,rj)
	err:
	fmt.Println("error")
	return c.String(http.StatusBadRequest,fmt.Sprintln(rj))
}

func main(){
		e:= echo.New()
		e.GET("/calcu/:id/:ie/:ig",getcalcu)
		e.POST("/calcu",postcalcu)
	e.Logger.Fatal(e.Start(":8000"))
}
