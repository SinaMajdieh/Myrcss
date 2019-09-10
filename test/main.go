package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nsf/sexp"
)

// type Init struct {
// 	Dir string
// }

func main() {

	// // const example_sexp = `
	// // 	(init l 1 b)
	// // `

	// if s, err := sexp.New(Init{Dir: "left"}); err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Printf("%#v", s)
	// }

	// // ast, err := sexp.Marshal(Init{Dir: "left"}, true)
	// // if err != nil {
	// // 	fmt.Println("2", err)
	// // 	return
	// // }
	// // fmt.Println(ast)
	// // i, err := sexp.Unmarshal(ast)
	// // fmt.Printf("%T", i)

	// // if err != nil {
	// // 	fmt.Println("1", err)
	// // 	return
	// // }

	// //fmt.Println(example.Pos[1])
	// //fmt.Println(example.Tgt)
	// const INIT = `
	// 	(init 2 2 2)
	// `
	// ast, err := sexp.Parse(strings.NewReader(INIT), nil)
	// if nil != err {
	// 	panic(err)
	// }

	// var p struct {
	// 	init int `sexp: "init" `
	// }
	// // p := Init{}
	// err = ast.Unmarshal(&p)
	// if nil != err {
	// 	panic(err)
	// }
	// fmt.Println(p.init)
	const example_sexp = `(init true) (init1 1.2)`

	var example struct {
		Bool  bool    `sexp:"init"`
		Float float64 `sexp:"init1"`
		//Tgt [3]float32 `sexp:"target,siblings"`
	}

	ast, err := sexp.Parse(strings.NewReader(example_sexp), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ast.Unmarshal(&example)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(example)
	// fmt.Println(example.Pos[2])
	// if "befor_kick_off" == example.Pos[2] {
	// 	fmt.Println("OK")
	// }
	// //fmt.Println(example.Tgt)
	// s, err := sexp.ParseString((example_sexp))
	// if nil != err {
	// 	panic(err)
	// }

	// for _, v := range s {
	// 	s := v.Head()
	// 	fmt.Print(s)
	// 	d := fmt.Sprint(s)
	// 	fmt.Printf("\n\n\n%s\n\n\n", d)
	// }
	fmt.Print(strconv.ParseBool("TRUE"))
	s := "sina"
	fmt.Println(s[0])
}
