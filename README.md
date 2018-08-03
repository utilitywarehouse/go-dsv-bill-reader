# Bill DSV

Some files from Bill (`Comms.txt`) include multi-line strings that are not
quoted. No D/CSV reader seemed to deal with these files so a custom library was
necessary.

For example:

```csv
1000|first string|final string
1001|second string
that is multi-line|final string
1002|third string|final string
```

Record `1001` contains a value that spans multiple lines but there are no quotes
which is usually the strategy to handle multi-line values in DSV formats.

## Usage

`Reader` exports `Read` and `ReadAll` and it behaves like the standard CSV
reader with a few exceptions:

- No quoting, because Bill files don't include quotes (from what I can tell)
- No slice optimisation (yet)
- No comments

For example:

```go
func main() {
    f, err := os.Open("Bill Data/Comms/18/07/31/Comms.txt")
    if err != nil {
        panic(err)
    }

    cr := billdsv.NewReader(f)

    for {
        row, err := cr.Read()
        if err != nil {
            fmt.Println(err)
            break
        }
        fmt.Println(row)
    }
}
```
