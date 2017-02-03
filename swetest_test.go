package swetest

import "testing"

var opt = []string{
	"-edir./resources/",
	"-b11.11.2016",
	"-ut00:00:00",
	"-p0",
	"-fZ",
	"-eswe",
	"-head",
}

func TestNew(t *testing.T) {
	New()
}

func TestSetBinPath(t *testing.T) {
	s := New()
	s.SetBinPath("path/to")

	if s.binPath != "path/to" {
		t.Error("SetBinPath error")
	}
}

func TestQuery(t *testing.T) {
	s := New()
	s.SetBinPath("./")
	res, err := s.Query(opt)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(res))
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New()
	}
}

func BenchmarkSetPath(b *testing.B) {
	s := New()
	for i := 0; i < b.N; i++ {
		s.SetBinPath("path/to")
	}
}

func BenchmarkQuery(b *testing.B) {
	s := New()
	for i := 0; i < b.N; i++ {
		s.Query(opt)
	}
}

func BenchmarkGetDefaultBinPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getDefaultBinPath()
	}
}
