package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"text/template"
	"time"

	_ "github.com/invopop/jsonschema"
)

type body struct {
	// 1. 从请求体中读取数据
	Body   string `json:"body"`
	Struct string `json:"struct"`
}

func toJsonSchema(reader io.Reader) (string, error) {

	var bd body
	err := json.NewDecoder(reader).Decode(&bd)
	if err != nil {
		return "", err
	}
	bs, err := parseFiles(bd)
	if err != nil {
		return "", err
	}
	dir, path, err := write(bs)
	if err != nil {
		return "", err
	}
	return cmd(dir, path)
}

func parseFiles(bd body) ([]byte, error) {
	tpl, err := template.ParseFiles("./template/main.tmpl")
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	if err := tpl.Execute(buf, bd); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func write(bs []byte) (string, string, error) {
	dir := fmt.Sprintf("%d", time.Now().UnixNano())
	path := fmt.Sprintf("./%s/main.go", dir)
	err := os.Mkdir(dir, 0755)
	if err != nil {
		return "", "", err
	}
	f, err := os.Create(path)
	if err != nil {
		return "", "", err
	}
	_, err = f.Write(bs)
	if err != nil {
		return "", "", err
	}

	err = f.Close()
	if err != nil {
		return "", "", err
	}
	return dir, path, nil
}

func cmd(dir, path string) (s string, err error) {
	defer func() {
		os.RemoveAll(dir)
	}()
	cmd := exec.Command("go", "run", path)
	// 执行命令并等待完成
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s:%s", err.Error(), string(output))
	}

	// 输出命令执行结果
	return string(output), cmd.Err
}

func main() {
	cwd, _ := os.Getwd()
	fmt.Println("cwd:", cwd)
	htmlBs, err := os.ReadFile("./index.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(htmlBs)
	})

	http.HandleFunc("/action", func(w http.ResponseWriter, r *http.Request) {
		//读取请求体http.Request
		toJsonSchemaBs, err := toJsonSchema(r.Body)
		if err != nil {
			r.Body.Close()
			w.Write([]byte(err.Error()))
			return
		}
		r.Body.Close()
		w.Write([]byte(toJsonSchemaBs))
	})
	fmt.Println("server start at 8080")
	http.ListenAndServe(":8080", nil)
}
