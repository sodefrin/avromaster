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

func Test_mapToStruct(t *testing.T) {
	data := struct {
		Test string `json:"test"`
	}{}

	mapData := map[string]string{
		"test": "test",
	}

	err := mapToStruct(mapData, &data)
	if err != nil {
		t.Fatalf("failed to mapToStruct (%v)", err)
	}

	if data.Test != "test" {
		t.Fatalf("data.Test is not equal mapData.test (%v)", data.Test)
	}
}
