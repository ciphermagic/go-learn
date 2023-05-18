package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Proxy struct {
	Remark        string
	Prefix        string
	Upstream      string
	RewritePrefix string
}

var proxyMap = make(map[string]Proxy)

func main() {
	loadProxyList()
	http.HandleFunc("/", forward)
	http.ListenAndServe(":8000", nil)
}

func forward(w http.ResponseWriter, r *http.Request) {
	var upstream = ""
	if value, ok := proxyMap[r.RequestURI]; ok {
		upstream = value.Upstream + value.RewritePrefix
	}

	remote, _ := url.Parse(upstream)
	fmt.Printf("RequestURI: %s, upstream: %s", r.RequestURI, upstream)

	r.Host = remote.Host
	r.URL.Path = ""
	httputil.NewSingleHostReverseProxy(remote).ServeHTTP(w, r)
}

func loadProxyList() {
	dataList := []Proxy{
		{Remark: "集群列表", Prefix: "/cdc-admin-clusters", Upstream: "https://xxx", RewritePrefix: "/api/dev/config/taskGroup/clusters"},
		{Remark: "任务状态", Prefix: "/cdc-admin-status", Upstream: "https://xxx", RewritePrefix: "/api/dev/config/taskGroup/status"},
	}
	for _, d := range dataList {
		proxyMap[d.Prefix] = d
	}
}
