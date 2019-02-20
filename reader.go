package avromaster

import (
	"io"
	"sync"

	"github.com/linkedin/goavro"
	"golang.org/x/xerrors"
)

type Reader interface {
	ReadSingle(data interface{}) error
	ReadMulti(max int, data interface{}) error
}

type reader struct {
	mux  *sync.Mutex
	r    io.Reader
	ocfr *goavro.OCFReader
}

func NewReader(r io.Reader) (Reader, error) {
	ocfr, err := goavro.NewOCFReader(r)
	if err != nil {
		return nil, xerrors.Errorf("failed to goavro.NewOCFReader : %w", err)
	}
	return &reader{mux: new(sync.Mutex), r: r, ocfr: ocfr}, nil
}

func (amr *reader) ReadSingle(data interface{}) error {
	amr.mux.Lock()
	if !amr.ocfr.Scan() {
		amr.mux.Unlock()
		return xerrors.New("failed to scan new rows")
	}
	in, err := amr.ocfr.Read()
	amr.mux.Unlock()
	if err != nil {
		return xerrors.Errorf("failed to goavro.Read : %w", err)
	}
	if err := mapToStruct(in, data); err != nil {
		return xerrors.Errorf("failed to mapToStruct : %w", err)
	}
	return nil
}

func (amr *reader) ReadMulti(max int, data interface{}) error {
	in := []interface{}{}
	for i := 0; i < max; i++ {
		amr.mux.Lock()
		if !amr.ocfr.Scan() {
			amr.mux.Unlock()
			return xerrors.New("failed to scan new rows")
		}
		tmp, err := amr.ocfr.Read()
		amr.mux.Unlock()
		if err != nil {
			return xerrors.Errorf("failed to goavro.Read : %w", err)
		}
		in = append(in, tmp)
	}
	if err := mapToStruct(in, data); err != nil {
		return xerrors.Errorf("failed to mapToStruct : %w", err)
	}
	return nil
}
