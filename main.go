package main

import (
  "os"
  "github.com/codegangsta/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "sc-listen"
  app.Usage = "stream soundcloud from the terminal"
  app.Action = func(c *cli.Context) {
    println("Not yet implemented")
  }

  app.Run(os.Args)
}

