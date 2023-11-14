package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"text/template"
)

// 定义一个结构体，用于在模板中引用字段
type person struct {
	Body string `json:"name"`
	Age  int    `json:"age"`
}

func Test_main(t *testing.T) {
	// 从文件加载模板
	tmpl, err := template.ParseFiles("./template/main.tmpl")
	if err != nil {
		fmt.Println("Error loading template:", err)
		return
	}

	// 创建一个数据对象，将在模板中使用
	person := person{Body: "", Age: 30}

	// 将数据对象应用到模板，并输出结果
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, person)
	if err != nil {
		fmt.Println("Error applying template:", err)
		return
	}
	fmt.Println(buf.String())
}

func Test_parseFiles(t *testing.T) {
	fmt.Println(os.Getwd())
	s, err := parseFiles(body{})
	if err != nil {
		t.Error(err)
	}
	t.Log(s)
}

func Test_toJsonSchema(t *testing.T) {
	fmt.Println(os.Getwd())
	bd := body{
		Struct: "Req",
		Body: `
type Req struct {
	Data struct {
		Token string  // token
	} 
}
`,
	}
	bs, _ := json.Marshal(bd)
	res, err := toJsonSchema(bytes.NewBuffer(bs))
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
