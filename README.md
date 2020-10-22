## How to use
  
### Example
main.go
```go
import (
	"fmt"

	"github.com/EricChiou/config"
)

// Config struct
type Config struct {
	Acc  string `key:"ACC"` // you can add key in tag for loading config by it.
	Pwd  string // if have no key tag, it will load config by name directly.
	IP   string
	ENV  string
	Int  int
	Bool bool
}

func main() {
	var cfg Config
	config.Load("./config.ini", &cfg)

	fmt.Println("Acc = " + cfg.Acc)
	fmt.Println("Pwd = " + cfg.Pwd)
	fmt.Println("IP = " + cfg.IP)
	fmt.Println("ENV = " + cfg.ENV)
	fmt.Println("Int =", cfg.Int)
	fmt.Println("Bool =", cfg.Bool)
	// Acc = your_account
	// Pwd = your_password
	// IP = server_ip
	// ENV = envirment
	// Int = 123456
	// Bool = true
}
```
config.ini
```ini
ACC=your_account
Pwd=your_password
IP =server_ip
ENV= envirment
Int=123456
Bool=true
```