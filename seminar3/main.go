package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Message struct {
	Title string `json:"title_name"`
	Text string `json:"text_message"`
}

func handleHello(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./hello.html.tpl"))

	m := &Message{
        Title: "Hello, world",
        Text:  "hogehoge",
	}
	
	if err := t.ExecuteTemplate(writer, "hello.html.tpl", m); err != nil {
		log.Fatal("error in handleHello")
	}
}

func handleHelloApi(w http.ResponseWriter, r *http.Request) {
	// 構造体Messageを使って、渡す値を作成(値は何でも良い)
	response := Message {
		Title: "Hello API",
		Text: "hogehogehoge!",
	}

    // json.Marshal()メソッドに作成した構造体を渡し、JSONボディを取得
	// (encoding/jsonのimportが必要です)
	bodyJson, err := json.Marshal(response)

    // err検知、ログ出し
    if err != nil {
        log.Fatal(err)
    }

	// w.Write()に作成したJSONを渡して、レスポンスする
	w.Write(bodyJson)
}

func main_old() {
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/hello_api", handleHelloApi)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

type Template struct {
	templates *template.Template
}

func (template *Template) Render(w io.Writer, name string, data interface{}, echo echo.Context) error {
	return template.templates.ExecuteTemplate(w,name,data)
}

func Hello(c echo.Context) error {
	m := &Message{
		Title: "NEW!",
		Text: "TEXT!!",
	}
	return c.Render(http.StatusOK, "hello", m)
}

func HellOAPI(c echo.Context) error {
	m := &Message{
		Title: "This is new hogehoge!",
		Text: "This is new fugafuga!!",
	}
	return c.JSON(http.StatusOK, m)
}

type Result struct {
	Answer int `json:"result"`
}

func plus(num1 int, num2 int) int {
	return num1 + num2
}

func Plus(c echo.Context) error {
	// requestからfirst, secondを取得する
	num1,_ := strconv.Atoi(c.QueryParam("first"))
	num2,_ := strconv.Atoi(c.QueryParam("second"))

	// 実際に足し合わせた値を取得する
	ans := plus(num1,num2)

	// result構造体に変換する
	result := &Result{Answer: ans,}

	// JSONを利用して返す
	return c.JSON(http.StatusOK, result)
}

type Numbers struct {
	first 	int `json: "first" query: "first"`
	second 	int `json: "second" query: "second"`
}

func main() {
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("./*.tpl")),
	}
	e.Renderer = t
	e.GET("/hello", Hello)
	e.GET("/hello_api", HellOAPI)

	e.GET("plus", Plus)

	e.Logger.Fatal(e.Start(":8080"))
}
