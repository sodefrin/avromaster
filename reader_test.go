package avromaster

import (
	"bytes"
	"testing"
)

func TestReadSingle(t *testing.T) {
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

	r, err := NewReader(b)
	if err != nil {
		t.Fatalf("failed to NewReader (%v)", err)
	}

	testOut := struct {
		Amount    int    `json:"amount"`
		Fee       int    `json:"fee"`
		CreatedAt string `json:"created_at"`
	}{}

	err = r.ReadSingle(&testOut)
	if err != nil {
		t.Fatalf("failed to ReadSingle (%v)", err)
	}
}

func TestReadMulti(t *testing.T) {
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

	r, err := NewReader(b)
	if err != nil {
		t.Fatalf("failed to NewReader (%v)", err)
	}

	testOut := []struct {
		Amount    int    `json:"amount"`
		Fee       int    `json:"fee"`
		CreatedAt string `json:"created_at"`
	}{}

	err = r.ReadMulti(2, &testOut)
	if err != nil {
		t.Fatalf("failed to ReadMulti (%v)", err)
	}
}
