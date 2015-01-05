/*
 * Simple Site Redirector.
 * - this code is used for security reason (default site for HAProxy or so...)
 *
 * Code by Jioh L. Jung (ziozzang@gmail.com)
 *
 * Lisense under MIT License.
 *
 */
package main

import (
        "flag"
        "io"
        "net/http"
        "strconv"
        "fmt"
)

var url string

func main() {
        port := flag.Int("port", 80, "Listening port number")
        site := flag.String("site", "warning.or.kr", "Target Site")
        flag.Parse()

        url = *site

        fmt.Printf("Listening=> :%d Redrect to -> '%s'\n", *port, *site)

        server := http.Server{
                Addr:    ":" + strconv.Itoa(*port),
                Handler: &myHandler{},
        }

        server.ListenAndServe()
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        str := "<html>\n<head>\n<meta http-equiv=\"Refresh\" content=\"0; url=http://" + url + "/\" />\n</head>\n</html>"

        fmt.Println("Request:" + r.RemoteAddr)

        w.Header().Add("Location", "http://"+url+"/")
        w.WriteHeader(301) //Moved
        io.WriteString(w, str)
}
