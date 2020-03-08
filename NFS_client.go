package main
import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "net"
    "log"
)

var server net.Conn
var err error
var port string = ":1337"

func main() {
    fmt.Println("Attempting to connect to NFS...")
    server,err = net.Dial("tcp", port)
    if err != nil{
        log.Fatal("Cannot connect to NFS!")
    }
    reader := bufio.NewReader(os.Stdin)
    for{
        fmt.Print("Client-> ")
        command,_ := reader.ReadString('\n')
        command = strings.Replace(command, "\n", "", -1)
        server.Write([]byte(command))
    }
}
