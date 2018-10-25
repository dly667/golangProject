package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"runtime/debug"
	"time"
)

//错误码常量
const (
	INVALID_BODY = 1
	INVALID_PARA = 2
	INNER_ERROR  = 3
)

//请求包体
type Req struct {
	Msg string `json:"msg"`
}

//回复包体
type Rsp struct {
	Status int64       `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data,omitempty"`
}

//路由器
var Router *mux.Router
type Process func(query map[string][]string, body []byte, rsp *Rsp)

func GenHandler(pro Process) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//url param
		query := r.URL.Query()
		//post data
		sbody, e := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if e != nil {
			w.WriteHeader(500)
			return
		}
		//process
		var rsp Rsp
		fmt.Printf("Req : %v\n", string(sbody))
		pro(query, sbody, &rsp)
		fmt.Printf("Rsp : %d, %v\n", rsp.Status, rsp.Data)

		//rsp
		buf, e := json.Marshal(&rsp)
		if e != nil {
			w.WriteHeader(500)
		}
		w.Write([]byte(buf))
	}
}

//服务
func ProcessEcho(query map[string][]string, body []byte, rsp *Rsp) {
	//解析url参数
	var isEcho bool
	echo := query["echo"]
	if echo != nil {
		//整型值使用strconv转换
		//state, e := strconv.ParseInt(echo[0], 10, 32)
		if echo[0] == "true" {
			isEcho = true
		} else {
			isEcho = false
		}
	}

	//解析post数据
	var req Req
	if e := json.Unmarshal(body, &req); e != nil {
		fmt.Printf("unmarshal fail %v [%v]\n", e, string(body))
		rsp.Status = INVALID_BODY
		rsp.Msg = "Invalid Req Body"
		return
	}

	//其他服务操作

	//回复数据
	if isEcho {
		rsp.Data = req.Msg
	}

	return
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
		}
	}()

	Router = mux.NewRouter()
	//cgi router
	Router.HandleFunc("/echo_svr", GenHandler(ProcessEcho))
	svr := http.Server{
		Addr:         ":38080",
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		Handler:      Router,
	}
	svr.ListenAndServe()
}

