package main

import (
        "io"
        "log"
        "os"
        "time"
        "net/http"
        "strconv"
        "fmt"
        // 引用私有包
        //"gitee.com/jxwuqingyan/go-mod-test"

)

var loger *log.Logger
func init() {
        os.Mkdir("./logs", os.ModePerm)
        file := "./logs/" + time.Now().Format("2006-01-02") + ".log"
        logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
        if err != nil {
                panic(err)
        }
        loger = log.New(logFile, "[log write to file]",log.LstdFlags | log.Lshortfile | log.LUTC) // 将文件设置为loger作为输出
        return
}


var iCnt int = 0;
func helloHandler(w http.ResponseWriter, r * http.Request) {
    iCnt++;
    str := "Hello world ! (" + strconv.Itoa(iCnt) + ")"
    io.WriteString(w, str)
    fmt.Println(str)
    loger.Println(str + "this is write to file!") // 使用的时候，需要采用loger作为输出的前缀
    //  引用私有包go-mod-test
    //go-mod-test.Hello()
}

func main() {
         ht := http.HandlerFunc(helloHandler)
         if ht != nil {
             http.Handle("/hello", ht)
         }
         err := http.ListenAndServe(":8090", nil)
         if err != nil {
             log.Fatal("ListenAndServe: ", err.Error())
         }
}
