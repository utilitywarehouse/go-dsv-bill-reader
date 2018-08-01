# Bill DSV

Some files from Bill (`Comms.txt`) include multi-line strings that are not
quoted. No D/CSV reader seemed to deal with these files so a custom library was
necessary.

See the `data.csv` file for an example of this problem.

## Usage

`Reader` exports `Read` and `ReadAll` and it behaves like the standard CSV
reader with a few exceptions:

- No quoting, because Bill files don't include quotes (from what I can tell)
- No slice optimisation (yet)
- No comments
- Requires `FieldsPerRecord` to be set explicitly

For example:

```go
func main() {
    f, err := os.Open("Bill Data/Comms/18/07/31/Comms.txt")
    if err != nil {
        panic(err)
    }

    cr := billdsv.NewReader(f)
    cr.FieldsPerRecord = 29 // necessary

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
