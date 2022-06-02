package main

import (
  "encoding/json"
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
  "net"
  "strings"
  "flag"
  "os"
  "log"
)

var (
  Log      *log.Logger
)


func init() {
    // set location of log file
   var logpath = "log/cert.log"
   log.Println(logpath)

   flag.Parse()
   var file, err1 = os.OpenFile(logpath, os.O_APPEND|os.O_WRONLY, 0644) 

   if err1 != nil {
      panic(err1)
   }
      Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
     
}
// Oncall Stuct
type Oncall struct {
  Details string `json:"details"`
}


var listner = make(chan *Oncall)

func writer(stats *Oncall) {
	listner <- stats
}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/", rootHandler).Methods("GET")
  router.HandleFunc("/oncall", oncallHandler).Methods("POST")
  router.HandleFunc("/app1", oncallHandler).Methods("POST")
  router.HandleFunc("/app2", oncallHandler).Methods("POST")
  router.HandleFunc("/app3", oncallHandler).Methods("POST")
  router.HandleFunc("/app4", oncallHandler).Methods("POST")
  router.HandleFunc("/app5", oncallHandler).Methods("POST")


  go oncallfunc()
  go app1()
  go app2()
  go app3()
  go app4()
  go app5()

  //log.Print("Starting server on port 8888!")
  Log.Printf("Starting server on port 8884!")
  Log.Fatal(http.ListenAndServe(":8884", router))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Home")
}

func oncallHandler(w http.ResponseWriter, r *http.Request) {

  ip, err := getIP(r)
  if err != nil {
    w.WriteHeader(400)
    w.Write([]byte("No valid ip"))
  }
  w.WriteHeader(200)
  Log.Println(ip +" - " , r.Header)

  //Log.Printf("on-call handler was called")

  var statistics Oncall
  //Log.Print(statistics)
  if err := json.NewDecoder(r.Body).Decode(&statistics); err != nil {
    Log.Printf("ERROR: %s", err)
    http.Error(w, "Bad request", http.StatusTeapot)
    return
  }
  defer r.Body.Close()
  go writer(&statistics)
}

func getIP(r *http.Request) (string, error) {
    //Get IP from the X-REAL-IP header
    ip := r.Header.Get("X-REAL-IP")
    netIP := net.ParseIP(ip)
    if netIP != nil {

        return ip, nil
    }

    //Get IP from X-FORWARDED-FOR header
    ips := r.Header.Get("X-FORWARDED-FOR")
    splitIps := strings.Split(ips, ",")
    for _, ip := range splitIps {
        netIP := net.ParseIP(ip)
        if netIP != nil {

            return ip, nil
        }
    }

    //Get IP from RemoteAddr
    ip, _, err := net.SplitHostPort(r.RemoteAddr)
    if err != nil {
        return "", err
    }
    netIP = net.ParseIP(ip)
    if netIP != nil {
        //log.Println(ip)
        return ip, nil
    }
    return "", fmt.Errorf("No valid ip found")
}

func oncallfunc() {
  for {
    val := <-listner
    in_out := fmt.Sprintf("%v ", val.Details)
    Log.Printf(in_out)
  }
}

func app1() {
  for {
    val := <-listner
    in_out := fmt.Sprintf("%v ", val.Details)
    Log.Printf(in_out)
  }
}

func app2() {
  for {
    val := <-listner
    in_out := fmt.Sprintf("%v ", val.Details)
    Log.Printf(in_out)
  }
}

func app3() {
  for {
    val := <-listner
    in_out := fmt.Sprintf("%v ", val.Details)
    Log.Printf(in_out)
  }
}

func app4() {
  for {
    val := <-listner
    in_out := fmt.Sprintf("%v ", val.Details)
    Log.Printf(in_out)
  }
}

func app5() {
  for {
    val := <-listner
    in_out := fmt.Sprintf("%v ", val.Details)
    Log.Printf(in_out)
  }
}
