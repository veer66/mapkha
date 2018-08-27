# Mapkha
Thai word segmentation (wordcut; word boundary identification; ตัดคำ) program in Go (golang)

## Example 1: Word segmentation

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
    dict, e := m.LoadDefaultDict()
    check(e)
    wordcut := m.NewWordcut(dict)
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        fmt.Println(strings.Join(wordcut.Segment(scanner.Text()), "|"))
    }
}

```


## Example 2: Word wrap

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
    dict, e := m.LoadDefaultDict()
    check(e)
    wordcut := m.NewWordcut(dict)
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        fmt.Println(strings.Join(wordcut.WordWrap(scanner.Text(), 30), "|"))
    }
}

```

# Contributors
* Vee Satayamas
* Weerasak Chongnguluam (@iporsut)
* Phongphan Phienphanich (@kapong)
