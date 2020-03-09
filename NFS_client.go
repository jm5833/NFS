package main
import (
    "fmt"
    //"strings"
    "bufio"
    "os"
    "net"
    "log"
)

var server net.Conn
var err error
var port string = ":1337"

//function to display messages from server
func handleServerMessage(server net.Conn){
    readscan := bufio.NewScanner(server)
    for readscan.Scan(){
        fmt.Println(readscan.Text())
    }
}

//function to accept input from stdin
//to send over to server
func acceptInput(server net.Conn){
    reader := bufio.NewReader(os.Stdin)
    go handleServerMessage(server)
    for{
        fmt.Print("Client-> ")
        command,_ := reader.ReadString('\n')
        //command = strings.Replace(command, "\n", "", -1)
        server.Write([]byte(command))
    }

}

//function to connect to the server
//returns a net.Conn object that's for the server
func serverConnect() net.Conn{
    fmt.Println("Attempting to connect to NFS...")
    server,err = net.Dial("tcp", port)
    if err != nil{
        log.Fatal("Cannot connect to NFS!")
    }
    fmt.Println("Connected!")
    return server
}

func main() {
    server = serverConnect()
    acceptInput(server)
}
