package avromaster

import (
	"io"

	"github.com/linkedin/goavro"
	"golang.org/x/xerrors"
)

type Writer interface {
	WriteSingle(data interface{}) error
	WriteMulti(data interface{}) (int, error)
}

type writer struct {
	schema string
	w      io.Writer
	codec  *goavro.Codec
	ocfw   *goavro.OCFWriter
}

func NewWriter(schema string, w io.Writer) (Writer, error) {
	config := goavro.OCFConfig{W: w, Schema: schema}
	ocfw, err := goavro.NewOCFWriter(config)
	if err != nil {
		return nil, xerrors.Errorf("failed to goavro.NewOCFWriter : %w", err)
	}
	codec, err := goavro.NewCodec(schema)
	if err != nil {
		return nil, xerrors.Errorf("failed to goavro.NewCodec : %w", err)
	}

	return &writer{schema: schema, w: w, codec: codec, ocfw: ocfw}, nil
}

func (amw *writer) WriteSingle(data interface{}) error {
	d, err := structToMap(data)
	if err != nil {
		return xerrors.Errorf("failed to structToMap : %w", err)
	}
	datum := []interface{}{d}
	if err := amw.ocfw.Append(datum); err != nil {
		return xerrors.Errorf("failed to goavro.Append : %w", err)
	}
	return nil
}

func (amw *writer) WriteMulti(data interface{}) (int, error) {
	datum, err := structToMap(data)
	if err != nil {
		return 0, xerrors.Errorf("failed to structToMap : %w", err)
	}
	if err := amw.ocfw.Append(datum); err != nil {
		return 0, xerrors.Errorf("failed to goavro.Append : %w", err)
	}
	return len(datum.([]interface{})), nil
}
