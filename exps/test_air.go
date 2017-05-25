package exps

import (
	"github.com/sheng/air"
	"net/http"
	"io"
	"log"
)

//22316 requests done.
//All requests done.
//
//Summary:
//Total:        0.6562 secs
//Slowest:      0.0202 secs
//Fastest:      0.0001 secs
//Average:      0.0011 secs
//Requests/sec: 45717.5441
//Total data:   330000 bytes
//Size/request: 11 bytes
//
//Status code distribution:
//[200] 30000 responses
//
//Response time histogram:
//0.000 [1]     |
//0.002 [27603] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
//0.004 [2096]  |∎∎∎
//0.006 [177]   |
//0.008 [27]    |
//0.010 [27]    |
//0.012 [13]    |
//0.014 [2]     |
//0.016 [14]    |
//0.018 [38]    |
//0.020 [2]     |
//
//Latency distribution:
//10% in 0.0005 secs
//25% in 0.0006 secs
//50% in 0.0008 secs
//75% in 0.0012 secs
//90% in 0.0019 secs
//95% in 0.0025 secs
//99% in 0.0041 secs
//

//
func RunAir(){
	a := air.New()
	a.GET("/", homeHandler)
	a.Serve()
}

func homeHandler(c *air.Context) error{
	return c.String("hello world")
}

//22506 requests done.
//All requests done.
//
//Summary:
//Total:        0.7456 secs
//Slowest:      0.0705 secs
//Fastest:      0.0001 secs
//Average:      0.0012 secs
//Requests/sec: 40234.7401
//Total data:   330000 bytes
//Size/request: 11 bytes
//
//Status code distribution:
//[200] 30000 responses
//
//Response time histogram:
//0.000 [1]     |
//0.007 [29890] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
//0.014 [52]    |
//0.021 [45]    |
//0.028 [4]     |
//0.035 [0]     |
//0.042 [4]     |
//0.049 [0]     |
//0.056 [0]     |
//0.063 [2]     |
//0.070 [2]     |
//
//Latency distribution:
//10% in 0.0006 secs
//25% in 0.0009 secs
//50% in 0.0010 secs
//75% in 0.0012 secs
//90% in 0.0020 secs
//95% in 0.0026 secs
//99% in 0.0046 secs

// vanilla http server
func RunSimple(){
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request){
		rw.Header().Add("Content-Type", "text/plain")
		io.WriteString(rw, "hello world")
	})
	log.Fatal(http.ListenAndServe(":2333", nil))
}
