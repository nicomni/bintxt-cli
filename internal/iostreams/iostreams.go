package iostreams

import (
	"bytes"
	"io"
	"os"
)

type IOStreams struct {
	In     io.Reader
	Out    io.Writer
	ErrOut io.Writer
}

func SystemIO() *IOStreams {
	return &IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}
}

func InMemoryIO() (io *IOStreams, in, out, errOut *bytes.Buffer) {
	in = new(bytes.Buffer)
	out = new(bytes.Buffer)
	errOut = new(bytes.Buffer)
	io = &IOStreams{
		In:     in,
		Out:    out,
		ErrOut: errOut,
	}
	return
}
