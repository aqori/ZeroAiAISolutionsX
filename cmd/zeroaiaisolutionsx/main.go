// cmd/zeroaiaisolutionsx/main.go
package main

import (
"flag"
"log"
"os"

"zeroaiaisolutionsx/internal/zeroaiaisolutionsx"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := zeroaiaisolutionsx.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
