package avromaster

import "encoding/json"

func structToMap(data interface{}) (interface{}, error) {
	var ret interface{}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func mapToStruct(inData interface{}, outData interface{}) error {
	b, err := json.Marshal(inData)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, outData); err != nil {
		return err
	}

	return nil
}
