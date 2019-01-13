package main

import (
  "fmt"
  "log"
  "os"

  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()

  app.Name = "release"
  app.Usage = "Display GhostBSD release information"
  app.Version = "18.12"
  app.Action = func(c *cli.Context) error {
    fmt.Println("GhostBSD 18.12")
    return nil
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
