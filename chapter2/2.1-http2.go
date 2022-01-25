package main

import (
	"fmt"
	"log"
	"net/http"
)

/*func main() {
	headers := http.Header{"token": {"kuysdfaeg6634fwr324brfh3urhf839hf349h"}}
	headers.Add("Accept-Charset","UTF-8")
	headers.Set("Host","www.shirdon.com")
	headers.Set("Location","www.baidu.com")
}*/

/*func main() {
	// 启动服务器
	srv := &http.Server{Addr: ":8088", Handler: http.HandlerFunc(handle)}
	// 用TLS启动服务器，因为我们运行的是http/2，它必须是与TLS一起运行。
	log.Printf("Serving on https://0.0.0.0:8088")
	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
}

//处理器方法
func handle(w http.ResponseWriter, r *http.Request) {
	// 记录请求协议
	log.Printf("Got connection: %s", r.Proto)
	// 向客户发送一条消息
	w.Write([]byte("Hello this is a HTTP/2 message!"))
}*/

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is HTTPS Server!")
}

func main() {
	http.HandleFunc("/hi", handler)
	err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal("ListenerAndServe:", err)
	}
}
