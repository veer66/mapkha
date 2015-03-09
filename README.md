# mapkha
Thai word segmentation program in Go

## Example

    ```go
    package main

    import ("fmt"
        "strings"
        "bufio"
        "os"
        m "github.com/veer66/mapkha"
    )

    func check(e error) {
        if e != nil {
            panic(e)
        }
    }

    func main() {
        dict, e := m.LoadDict("tdict-std.txt")
        check(e)
        scanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {
            fmt.Println(strings.Join(m.Segment(scanner.Text(), dict), "|"))
        }
    }

    ```
