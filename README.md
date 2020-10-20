# verint-swt-go
Go module to generate Service to Service (SWT) Tokens used with Verint's Web Services

## Disclaimer

I'm not a Go developer.  I managed to fumble my way through this implementation and it seems to work.

If you have any suggestions for enhancements please open an issue and PR.


#### Usage

```go
package main

import (
	"fmt"
	"verint-swt-go/verint-swt"

)

func main() {

	token := verint_swt.GenerateToken("VUXWPSXS", 
                            "iV2hwht_9spjpqb7UbbS-YHyuWzoRgo50j0MQ2s7Mls", 
                            "POST", 
                            "http://wfo-server/DASWebApi/Query/ExecuteDynamicQuery")
	fmt.Println("Token: ", token)
}
```

## Release Notes

### 1.0.0
Initial release