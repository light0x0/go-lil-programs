package main
import "fmt"
import "os"

func main(){
	args := os.Args
	if len(args) > 2{
		fmt.Println("Too many arguments")
		return
	}
	content , e := os.ReadFile(args[1])
	if e != nil {
		fmt.Println("File name missing")
	}
	fmt.Print(string(content))

}