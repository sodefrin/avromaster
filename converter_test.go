package avromaster

import (
	"fmt"
	"testing"
)

func Test_structToMap(t *testing.T) {
	data := struct {
		Test string `json:"test"`
	}{
		Test: "test",
	}

	ret, err := structToMap(data)
	if err != nil {
		t.Fatalf("failed to structToMap (%v)", err)
	}

	if fmt.Sprintf("%v", ret) != "map[test:test]" {
		t.Fatalf("ret is not equal map[test:test] (%v)", ret)
	}
}
