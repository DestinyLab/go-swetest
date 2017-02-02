/*
Package swetest is the simple wrapper of swisseph library's swetest.

Swetest test page: http://www.astro.com/swisseph/swetest.htm

Help: http://www.astro.com/cgi/swetest.cgi?arg=-h&p=0
*/
package swetest

import (
	"os/exec"
	"path/filepath"
	"runtime"
)

// A Swetest returns what you queryed to swisseph library's swetest.
type Swetest struct {
	path string
}

// New creates an instance of swetest.
func New() *Swetest {
	return &Swetest{path: getDefaultPath()}
}

// SetPath set the path of resources.
func (s *Swetest) SetPath(p string) {
	s.path = p
}

// Query query what you want.
func (s *Swetest) Query(q []string) ([]byte, error) {
	return exec.Command(s.path+"swetest", q[0:]...).Output()
}

func getDefaultPath() string {
	_, file, _, _ := runtime.Caller(0)

	return filepath.Dir(file) + "/resources/"
}
