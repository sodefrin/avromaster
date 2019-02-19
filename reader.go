package avromaster

import (
	"io"

	"github.com/linkedin/goavro"
	"golang.org/x/xerrors"
)

type Reader interface {
	ReadSingle(data interface{}) error
	ReadMulti(max int, data interface{}) error
}

type reader struct {
	r    io.Reader
	ocfr *goavro.OCFReader
}

func NewReader(r io.Reader) (Reader, error) {
	ocfr, err := goavro.NewOCFReader(r)
	if err != nil {
		return nil, err
	}
	return &reader{r: r, ocfr: ocfr}, nil
}

func (amr *reader) ReadSingle(data interface{}) error {
	if !amr.ocfr.Scan() {
		return xerrors.New("no rows")
	}
	_, err := amr.ocfr.Read()
	if err != nil {
		return err
	}
	return nil
}

func (amr *reader) ReadMulti(max int, data interface{}) error {
	return nil
}
