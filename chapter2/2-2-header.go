package main

import (
	"net/http"
)

func main() {
	headers := http.Header{"token": {"kuysdfaeg6634fwr324brfh3urhf839hf349h"}}
	headers.Add("Accept-Charset","UTF-8")
	headers.Set("Host","www.shirdon.com")
	headers.Set("Location","www.baidu.com")
}
