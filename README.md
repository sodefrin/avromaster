## avromaster

simple avrofile reader & writer. this repository is wrap of https://github.com/linkedin/goavro

## Write & Read

```
schema := `
{
  "type": "record",
  "name": "transactions",
  "fields" : [
    {"name": "amount", "type": "int"},
    {"name": "fee", "type": "int"},
    {"name": "created_at", "type": "string"}
  ]
}`

b := &bytes.Buffer{}

// pass srting & io.Writer 
w, _ := NewWriter(schema, b)

test := struct {
  Amount    int    `json:"amount"`
  Fee       int    `json:"fee"`
  CreatedAt string `json:"created_at"`
}{
  Amount:    100,
  Fee:       10,
  CreatedAt: "2019-01-29T18:36:46+09:00",
}

w.WriteSingle(test)

testOut := struct {
  Amount    int    `json:"amount"`
  Fee       int    `json:"fee"`
  CreatedAt string `json:"created_at"`
}{}

// pass io.Reader
r, _ := NewReader(b)

r.ReadSingle(&testOut)

```
