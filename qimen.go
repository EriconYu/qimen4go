package qimen

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
)

//go:embed engine.cjs
var engine []byte
var enginePath string
var engineErr error
var once sync.Once

func path() (string, error) {
	once.Do(func() {
		f, e := os.CreateTemp("", "qimen-engine-*.cjs")
		if e != nil {
			engineErr = e
			return
		}
		if _, e = f.Write(engine); e != nil {
			engineErr = e
			return
		}
		f.Close()
		enginePath = f.Name()
	})
	return enginePath, engineErr
}

type engineEnvelope struct {
	Ok        bool
	Data      map[string]any
	Canonical map[string]any
	Error     string
}

func run(input map[string]any) (*engineEnvelope, error) {
	p, e := path()
	if e != nil {
		return nil, e
	}
	b, _ := json.Marshal(input)
	cmd := exec.Command("node", p)
	cmd.Stdin = bytesReader(b)
	out, e := cmd.CombinedOutput()
	if e != nil {
		return nil, fmt.Errorf("engine: %w: %s", e, out)
	}
	var envelope engineEnvelope
	if e = json.Unmarshal(out, &envelope); e != nil {
		return nil, e
	}
	if !envelope.Ok {
		return nil, fmt.Errorf("%s", envelope.Error)
	}
	return &envelope, nil
}

func Calculate(input map[string]any) (map[string]any, error) {
	envelope, err := run(input)
	if err != nil {
		return nil, err
	}
	return envelope.Data, nil
}

func Canonical(input map[string]any) (map[string]any, error) {
	envelope, err := run(input)
	if err != nil {
		return nil, err
	}
	return envelope.Canonical, nil
}

type reader struct {
	b []byte
	i int
}

func bytesReader(b []byte) *reader { return &reader{b: b} }
func (r *reader) Read(p []byte) (int, error) {
	if r.i >= len(r.b) {
		return 0, io.EOF
	}
	n := copy(p, r.b[r.i:])
	r.i += n
	return n, nil
}
