package Scanner_ports
import (
	"fmt"
	"net"
)

func main() {
	var link string
	fmt.Println("Enter the IP address or link to scan: ")
	fmt.Scanln(&link)

	for i := 1; i <= 1024; i++ {
		go func(j int) {
			address := fmt.Sprintf("%s:%d", link, j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}

	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
