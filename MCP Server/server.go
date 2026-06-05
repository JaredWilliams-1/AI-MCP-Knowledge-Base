package main

import (
    "bufio";
    "os";
    "fmt"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()

        fmt.Println(line)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("There was an error scanning.")
        os.Exit(1)
    }
}