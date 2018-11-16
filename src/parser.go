package src

import "io"

type Parser interface {
	Parse(r io.ReadCloser) (body string, urls []string, err error)
}
