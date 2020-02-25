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
func errorCheck(e error) bool {
    if e != nil{
        fmt.Println(e)
        return true
    }
    return false
}

func readFile(fname string, mode int){
    //manually setting the mode to make testing easier
    mode = os.O_RDWR
    file,err := os.OpenFile(fname, mode, 0755)
    if errorCheck(err){ return }

    //using stat here to get file length 
    fi,err := file.Stat()
    if errorCheck(err){ return }
    //create a slize the size of the file so that we grab the while file 
    ficon := make([]byte, fi.Size())

    //read the file, not interested in bytes read
    //since the size of the buffer = size of the file in bytes
    _,err = file.Read(ficon)
    if errorCheck(err){ return }
    fmt.Println(string(ficon))
}

func processCall(call string) {
    if len(call) <= 0{ return }
    args := strings.Split(call, " ")
    command := args[0]
    switch command {
        case "exit":
            os.Exit(0)
        case "read":
            fname := args[1]
            mode,err := strconv.Atoi(args[2])
            if err != nil{
                fmt.Println("Invalid mode")
                return
            }
            readFile(fname,mode)
    }
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
        processCall(call)
    }
}
