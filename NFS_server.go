package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    "net"
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
//port to use for the NFS server, chosen at random
var port string = ":1337"

func errorCheck(e error) bool {
    if e != nil{
        fmt.Println(e)
        return true
    }
    return false
}

func readFile(fname string, client net.Conn){
    //manually setting the mode to make testing easier
    mode := os.O_RDWR
    file,err := os.OpenFile(fname, mode, 0755)
    if errorCheck(err){ return }

    //defer close to make sure file is closed 
    defer file.Close()

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
    //after reading the file, send the read file over to the client 
    client.Write(ficon)
}

//function to check that the offset isn't larger than the file size
func checkOffset(offset int64, file *os.File) int64{
    fs,err := file.Stat()
    if errorCheck(err){ return int64(0) }
    size := fs.Size()
    if offset > size{
        return size
    }
    return offset
}

//mode == "append" - adds at the ends of the list
//mode == "replace" - replaces everything after offset with content 
func writeToFile(fname string, offset int64, mode string, content []byte){
    fileMode := os.O_RDWR | os.O_CREATE
    //add in the append flag if we're appending to the file
    if mode == "append"{ fileMode = fileMode | os.O_APPEND }

    file,err := os.OpenFile(fname, fileMode, 0755)
    if errorCheck(err){ return }
    defer file.Close()

    //convert the content and offset to []byte and int64 
    //to work with the file functions 
    switch mode{
        case "append":
            _,err = file.Write(content)
        case "replace":
            offset = checkOffset(offset,file)
            _,err = file.WriteAt(content,offset)
    }
    if errorCheck(err){ return }
}

func processCall(call string, client net.Conn) {
    if len(call) <= 0{ return }
    args := strings.Split(call, " ")
    command := args[0]
    switch command {
        case "exit":
            os.Exit(0)
        case "read":
            fname := args[1]
            readFile(fname, client)
        case "write":
            fname := args[1]
            off,err := strconv.Atoi(args[2])
            if err != nil{
                fmt.Println("Invalid offset")
                return
            }
            offset := int64(off)
            mode := args[3]
            content := []byte(args[4])
            if err != nil{
                fmt.Println("Invalid offset")
            }
            writeToFile(fname,offset,mode,content)
    }
}
//function that accepts incoming connections
func acceptClients(server net.Listener){
    fmt.Println("Shh, listening...")
    for{
        client,err := server.Accept()
        if errorCheck(err){ return }
        fmt.Println("A client has connected.")
        go handleClient(client)
    }
}

func handleClient(client net.Conn){
    //create the buffer to read messages from NfS clients
    readbuf := bufio.NewScanner(client)
    //buffer to respond to NFS clients
    for readbuf.Scan(){
        call := readbuf.Text()
        fmt.Println(call)
        processCall(call,client)
    }
}

//starter function
func main() {
    server,err := net.Listen("tcp",port)
    if errorCheck(err){ return }
    acceptClients(server)

    /*
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter command:")
    for {
        fmt.Print("-> ")
        //read everything up to the newline(user hitting enter)
        call,_ := reader.ReadString('\n')
        call = strings.Replace(call, "\n", "", -1)
        processCall(call)
    */
}
