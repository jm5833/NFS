package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
)
/*
    starter commands
        read
        write
        create
    command format: command(arg1, arg2, ...)

*/

func ProcessCall(call string) {
    cindex := strings.Index(call, "(")
    if cindex == -1{
        fmt.Println("Invalid syntax")
        return
    }
    //command := call[:cindex]
    args := strings.Split(call[cindex+1:len(call)-1], ",")
    fmt.Println(args)
}

//starter function
func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter command:")
    for {
        fmt.Print("->")
        //read everything up to the newline(user hitting enter)
        call,_ := reader.ReadString('\n')
        call = strings.Replace(call, "\n", "", -1)
        ProcessCall(call)
    }
}
