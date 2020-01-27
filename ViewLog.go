package main
import ( 
	"bufio"
	"io"
	"fmt"
	"os"
	"log"
	"time"
)

func main ()  {
	if len(os.Args) < 2 {
		os.Exit(0)
	}
	arg := os.Args[1]

	file, err := os.Open(arg)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	reader := bufio.NewReader(file)
	for true {
		line, err := reader.ReadString('\n')
		if (err != io.EOF) {
			fmt.Print(line) 
		} else {
			time.Sleep(time.Second / 10)
		}
	}
}