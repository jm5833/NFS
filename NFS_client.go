package main
import (
    "fmt"
    "strings"
    "bufio"
    "os"
)
func main() {
    reader := bufio.NewReader(os.Stdin)
    for{
        fmt.Print("Client->: ")
        command,_ := reader.ReadString('\n')
        command = strings.Replace(command, "\n", "", -1)
        fmt.Println(command)
    }
}
