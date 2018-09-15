package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
	"net/http"
	"strconv"

    "github.com/julienschmidt/httprouter"
)

// /Hello/:langにハンドルされているHello関数
func Hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    lang := p.ByName("lang") // langパラメーターを取得する
    fmt.Fprintf(w, lang)     // レスポンスに値を書き込む
}

func FizzBuzz(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	num_temp := p.ByName("num") // langパラメーターを取得する
	var num int
	var err = ""
	num, _ = strconv.Atoi(num_temp)
	//num, nil = strconv.Atoi(num_temp)
	/*
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}
	*/

	switch {
	//case  type != "int"
		//404 Error
	case num%15 == 0:
		fmt.Fprintf(w, "FIZZ BUZZ!")
	case num%3 == 0:
		fmt.Fprintf(w, "FIZZ!")
	case num%5 == 0:
		fmt.Fprintf(w, "BUZZ!")
	default:
		//fmt.Println(num)
		var res string
		res = strconv.Itoa(num)
		fmt.Fprintf(w, res)
	}

	//fmt.Fprintf(w, num)     // レスポンスに値を書き込む

}


// /ExampleにハンドルされているExample関数
func Example(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    defer r.Body.Close() // Example関数が終了する時に実行されるdeferステートメント

    // リクエストボディを読み取る
    bodyBytes, err := ioutil.ReadAll(r.Body)
    if err != nil {
        // リクエストボディの読み取りに失敗した => 400 Bad Requestエラー
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // JSONパラメーターを構造体にする為の定義
    type ExampleParameter struct {
        ID   int    `json:"id"`
        Name string `json:"name"`
    }
    var param ExampleParameter

    // ExampleParameter構造体に変換
    err = json.Unmarshal(bodyBytes, &param)
    if err != nil {
        // JSONパラメーターを構造体への変換に失敗した => 400 Bad Requestエラー
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // 構造体に変換したExampleParameterを文字列にしてレスポンスに書き込む
	fmt.Fprintf(w, fmt.Sprintf("%+v\n", param))
}

func main() {
    router := httprouter.New() // HTTPルーターを初期化

    // /HelloにGETリクエストがあったらHello関数にハンドルする
    // :langはパラメーターとして扱われる
    router.GET("/Hello/:lang", Hello)

    // /ExampleにPOSTリクエストがあったらExample関数にハンドルする
	router.POST("/Example", Example)

	router.GET("/FizzBuzz/:num", FizzBuzz)

    // Webサーバーを8080ポートで立ち上げる
    err := http.ListenAndServe(":8080", router)
    if err != nil {
        log.Fatal(err)
    }
}