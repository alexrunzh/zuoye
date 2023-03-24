package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	os.Setenv("VERSION", "0.1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	fmt.Printf("OS VERSION: %v\n", version)

	for k, v := range r.Header {
		//fmt.Println(k, v)
		for _, vv := range v {
			fmt.Printf("Header key: %s, Header value: %s\n", k, v)
			w.Header().Set(k, vv)
		}
	}
	currentip := getCurrentIP(r)
	println(r.RemoteAddr)
	//fmt.Printf("clientip: %v\n", clientip)
	log.Printf("Success! clientip: %s", currentip)

	clientip := ClientIP(r)
	println(r.RemoteAddr)
	//fmt.Printf("clientip: %v\n", clientip)
	log.Printf("Success! clientip: %s", clientip)

}


func getCurrentIP(r *http.Request) string {

	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]

	}
	return ip
}


func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func healthz(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "200")
}

func main() {
	// 构架http服务器
	mux := http.NewServeMux()
	// 06. debug
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	// 指定根路由
	mux.HandleFunc("/", index)
	// healthz 路由
	mux.HandleFunc("/healthz", healthz)
	// 判断httpserver 是否启动成功
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Start httpserver failed, error: %v", err.Error())

	}
}
