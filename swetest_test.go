package swetest

import "testing"

func TestNew(t *testing.T) {
	New()
}

func TestSetPath(t *testing.T) {
	s := New()
	s.SetPath("path/to")

	if s.path != "path/to" {
		t.Error("SetPath error")
	}
}

func TestQuery(t *testing.T) {
	opt := []string{`-b11.11.2016 -ut00:00:00 -p0 -fZ -eswe`}
	s := New()
	s.SetPath("./resources/")
	_, err := s.Query(opt)
	if err != nil {
		t.Error(err)
	}
}

func Benchmark_New(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New()
	}
}

func Benchmark_SetPath(b *testing.B) {
	s := New()
	for i := 0; i < b.N; i++ {
		s.SetPath("path/to")
	}
}

func Benchmark_Query(b *testing.B) {
	s := New()
	opt := []string{
		"-b11.11.2016",
		"-ut00:00:00",
		"-p0",
		"-fZ",
		"-eswe",
		"-head",
	}
	for i := 0; i < b.N; i++ {
		s.Query(opt)
	}
}

func Benchmark_getDefaultPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getDefaultPath()
	}
}
