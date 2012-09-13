package main

import (
    "log"
    "net/http"
    "net/url"
    "os"
    "github.com/alphazero/Go-Redis"
    "fmt"
    "strings"
    "strconv"
    "math/rand"
)
var(    
		redis_url, _ = url.Parse(os.Getenv("REDISTOGO_URL"))
    password, _ = redis_url.User.Password()
    url_parts = strings.Split(redis_url.Host, ":")
    port,_ = strconv.Atoi(url_parts[1])
    host = url_parts[0]
    spec = redis.DefaultSpec().Host(host).Port(port).Password(password)
    client, e = redis.NewSynchClientWithSpec (spec) // TODO handle err, no pw etc
)

const charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {
    http.HandleFunc("/", processLink)
    http.HandleFunc("/create", addLink)
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

func processLink(w http.ResponseWriter, req *http.Request) {
    link := req.URL.Path[1:]
    value, _ := client.Get(link)
    if value == nil {
			http.Error(w, "Not Found", http.StatusNotFound)
		} else {
      log.Printf("Url found: %s", value)
			http.Redirect(w, req, fmt.Sprintf("%s", value), http.StatusFound)
		}
}

func addLink(w http.ResponseWriter, req *http.Request) {
	key := makeKey()
	form_val := req.FormValue("link")
	val := form_val[0:len(form_val)]
	byte_val := []byte(val)
  log.Printf("Added: %s", byte_val)
  client.Set(key, byte_val) //TODO handle error (eg key exists)
  fmt.Fprintf(w, "%s", key)
}

func makeKey() string {
  //TODO generate same key, for same url
  s := ""
  for i := 0; i < 5; i++ {
    index := rand.Intn(len(charSet))
    s += charSet[index:1+index]
  }
  return s
}
