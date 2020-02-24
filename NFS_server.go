package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)


/*
    starter commands
        read
        write
        create
    command format: command(arg1, arg2, ...)
    
    currently printing output, once the NFS client is created 
    then it'll be converted to a send command to the client 

    read    1
    write   2
    execute 4
*/
func ReadFile(fname string, mode int){
    file,err := os.OpenFile(fname, mode, 0755)
    if err != nil{
        fmt.Println(err)
        return
    }
    fmt.Println(file.Read())
}
func ProcessCall(call string) {
    cindex := strings.Index(call, "(")
    if cindex == -1{
        fmt.Println("Invalid syntax")
        return
    }

    command := call[:cindex]
    args := strings.Split(call[cindex+1:len(call)-1], ",")
    fmt.Println(args)

    if command == "read"{
        fname := args[0]
        mode,err := strconv.Atoi(args[1])
        if err != nil{
            fmt.Println("Invalid mode")
            return
        }
        ReadFile(fname,mode)
    }else if command == "exit"{
        os.Exit(0)
    }
    /*else if command == "write"{

    }else if command == "create"{

    }else{
        fmt.Println("invalid command")
    }*/
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
