package main

import (
  "fmt"
  "log"
  "os"

  "github.com/urfave/cli/v2"
)

func main() {

        var (
                protocolHTTP bool
                protocolGRPC bool
        )


        // HTTP
        // RPC

        app := &cli.App{
                Name: "My Database",
                Usage: "Strongly-Typed Data Access Layer",
                Commands: []*cli.Command{
                        {
                                Name:    "run",
                                Aliases: nil,
                                Usage:   "Run the server.",
                                Action:  func(c *cli.Context) error {
                                        fmt.Println("added task: ", c.Args().First())
                                        return nil
                                },
                                Flags: []cli.Flag {
                                        &cli.BoolFlag{
                                                Name: "http",
                                                Usage: "Serve data over HTTP",
                                                Destination: &protocolHTTP,
                                        },
                                        &cli.BoolFlag{
                                                Name: "grpc",
                                                Usage: "Serve data over GRPC",
                                                Destination: &protocolGRPC,
                                        },
                                },
                        },
                        {
                                Name:    "docs",
                                Aliases: nil,
                                Usage:   "Run documentation server.",
                                Action:  func(c *cli.Context) error {
                                        fmt.Println("added task: ", c.Args().First())
                                        return nil
                                },
                        },
                },
        }

        err := app.Run(os.Args)
        if err != nil {
                log.Fatal(err)
        }
}

