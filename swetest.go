/*
Package swetest is the simple wrapper of swisseph library's swetest.

Swetest:
 test page: http://www.astro.com/swisseph/swetest.htm
 Help: http://www.astro.com/cgi/swetest.cgi?arg=-h&p=0
 ephemeris files: ftp://ftp.astro.com/pub/swisseph/ephe/

Example:
	package main

	import (
		"fmt"
		"log"

		"github.com/DestinyLab/go-swetest"
	)

	func main() {
		opt := []string{
			"-edir./resources/",
			"-b11.11.2017",
			"-ut00:00:00",
			"-p0",
			"-fZ",
			"-eswe",
			"-head",
		}
		s := swetest.New()
		res, err := s.Query(opt)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", res)
	}
*/
package swetest

import (
	"os/exec"
	"path/filepath"
	"runtime"
)

// A Swetest returns what you queryed to swisseph library's swetest.
type Swetest struct {
	binPath string
}

// New creates an instance of swetest.
func New() *Swetest {
	return &Swetest{binPath: getDefaultBinPath()}
}

// SetBinPath set the path to swetest(executable binary).
func (s *Swetest) SetBinPath(p string) {
	s.binPath = p
}

// Query query what you want.
func (s *Swetest) Query(q []string) ([]byte, error) {
	return exec.Command(s.binPath+"swetest", q[0:]...).Output()
}

func getDefaultBinPath() string {
	_, file, _, _ := runtime.Caller(0)

	return filepath.Dir(file)
}
