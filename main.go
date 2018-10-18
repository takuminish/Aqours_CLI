package main

import (
    "fmt"
    "log"
    "os"

    "github.com/urfave/cli"
)

func main() {
    app := cli.NewApp()
    app.Name = "Aqours"
    app.Usage = "Aqours information"
    app.Flags = []cli.Flag {
        cli.BoolFlag {
	    Name: "profile,p",
	    Usage: "Aqours Character's profile",
	},
    }
    app.Commands = []cli.Command{
        {
	    Name: "chika",
	    Usage: "chika",
	    Action: chika,
	},
    }

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}

func chika(c *cli.Context) {
    var isProfile = c.GlobalBool("profile")
    if isProfile {
        fmt.Println("Profile:\n")
    }
    fmt.Println("chika")
}

