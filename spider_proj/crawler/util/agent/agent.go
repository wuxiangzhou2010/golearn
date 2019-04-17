package agent

import "io"

type Agent interface {
	Get(url string) (resp io.Reader, err error)
}
