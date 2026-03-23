package main
import (
  "net/http"; "net/http/httputil"; "net/url"; "sync/atomic"
)
func main(){ backends:=[]string{"http://localhost:9001","http://localhost:9002"}; var idx uint64; http.HandleFunc("/", func(w http.ResponseWriter,r *http.Request){ i:=atomic.AddUint64(&idx,1); target,_:=url.Parse(backends[(i-1)%uint64(len(backends))]); httputil.NewSingleHostReverseProxy(target).ServeHTTP(w,r) }); http.ListenAndServe(":8083", nil) }
