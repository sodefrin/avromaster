package avromaster

import (
	"encoding/json"

	"golang.org/x/xerrors"
)

func structToMap(data interface{}) (interface{}, error) {
	var ret interface{}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, xerrors.Errorf("failed to marshal json: %w", err)
	}
	if err := json.Unmarshal(b, &ret); err != nil {
		return nil, xerrors.Errorf("failed to unmarshal json: %w", err)
	}

	return ret, nil
}

func mapToStruct(inData interface{}, outData interface{}) error {
	b, err := json.Marshal(inData)
	if err != nil {
		return xerrors.Errorf("failed to marshal json: %w", err)
	}
	if err := json.Unmarshal(b, outData); err != nil {
		return xerrors.Errorf("failed to unmarshal json: %w", err)
	}

	return nil
}
