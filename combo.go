package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
//    "net"  // use for unix socket
    "net/http"
    "net/url"
    "os"
)

var js_root string  // path to js files

func viewHandler(w http.ResponseWriter, r *http.Request) {
    files := make([]string, 0)
    q := r.URL.RawQuery
    params, _ := url.ParseQuery(q)

    // populate files from url query
    for k, v := range params {
        if v[0] == "" {
            // just a key, e.g. ?foo&bar&baz
            // assume filename and append path
            file := js_root + k
            files = append(files, file)
        }
        // else this has a value, e.g. ?foo=bar
        // maybe add support for parameters?
    }
    
    // create slice of file contents
    output := make([][]byte, 0)
    for _, file := range files {
        c, e := ioutil.ReadFile(file)
        if e == nil {
            output = append(output, c)
        } else {
            fmt.Println(e)
        }
    }

    // TODO add support for css or js content type
    w.Header().Set("Content-Type", "text/javascript")
    fmt.Fprintf(w, "%s", bytes.Join(output, []byte{}))
}

func main() {
    js_root = os.Getenv("GOOEY_JS_ROOT")
    if js_root == "" {
        fmt.Println("Missing environment variable GOOEY_JS_ROOT")
        os.Exit(1)
    }
    fmt.Println("Gooey combo server running")
    // net socket
    http.HandleFunc("/combo/", viewHandler)
    http.ListenAndServe("127.0.0.1:9000", nil)
    // unix socket
    /*l, err := net.Listen("unix", "/tmp/foo.bar.sock")
    if err != nil {
        fmt.Printf("%s\n", err)
    } else {
        http.Serve(l, nil)
    }*/
}
