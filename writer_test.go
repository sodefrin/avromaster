package avromaster

import (
	"bytes"
	"testing"
)

func Test_WriteSingle(t *testing.T) {
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
	w, err := NewWriter(schema, b)
	if err != nil {
		t.Fatalf("failed to NewWriter (%v)", err)
	}

	test := struct {
		Amount    int    `json:"amount"`
		Fee       int    `json:"fee"`
		CreatedAt string `json:"created_at"`
	}{
		Amount:    100,
		Fee:       10,
		CreatedAt: "2019-01-29T18:36:46+09:00",
	}

	err = w.WriteSingle(test)
	if err != nil {
		t.Fatalf("failed to WriteSingle (%v)", err)
	}
}

func Test_WriteMulti(t *testing.T) {
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
	w, err := NewWriter(schema, b)
	if err != nil {
		t.Fatalf("failed to NewWriter (%v)", err)
	}

	test := []struct {
		Amount    int    `json:"amount"`
		Fee       int    `json:"fee"`
		CreatedAt string `json:"created_at"`
	}{
		{
			Amount:    100,
			Fee:       10,
			CreatedAt: "2019-01-29T18:36:46+09:00",
		},
		{
			Amount:    100,
			Fee:       10,
			CreatedAt: "2019-01-29T18:36:46+09:00",
		},
	}

	n, err := w.WriteMulti(test)
	if err != nil {
		t.Fatalf("failed to WriteMulti (%v)", err)
	}
	if n != len(test) {
		t.Fatalf("n should be equal to %d but has %d", len(test), n)
	}
}
