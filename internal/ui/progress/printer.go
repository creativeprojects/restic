package progress

import "testing"

// A Printer can can return a new counter or print messages
// at different log levels.
// It must be safe to call its methods from concurrent goroutines.
type Printer interface {
	NewCounter(description string) *Counter

	// E prints to stderr
	E(msg string, args ...interface{})
	// S prints to stdout
	S(msg string, args ...interface{})
	// P prints to stdout unless quiet was passed
	P(msg string, args ...interface{})
	// V prints to stdout if verbose is set once
	V(msg string, args ...interface{})
	// VV prints to stdout if verbose is set twice
	VV(msg string, args ...interface{})
}

// NoopPrinter discards all messages
type NoopPrinter struct{}

var _ Printer = (*NoopPrinter)(nil)

func (*NoopPrinter) NewCounter(_ string) *Counter {
	return nil
}

func (*NoopPrinter) E(_ string, _ ...interface{}) {}

func (*NoopPrinter) S(_ string, _ ...interface{}) {}

func (*NoopPrinter) P(_ string, _ ...interface{}) {}

func (*NoopPrinter) V(_ string, _ ...interface{}) {}

func (*NoopPrinter) VV(_ string, _ ...interface{}) {}

// TestPrinter prints messages during testing
type TestPrinter struct {
	t testing.TB
}

func NewTestPrinter(t testing.TB) *TestPrinter {
	return &TestPrinter{
		t: t,
	}
}

var _ Printer = (*TestPrinter)(nil)

func (p *TestPrinter) NewCounter(_ string) *Counter {
	return nil
}

func (p *TestPrinter) E(msg string, args ...interface{}) {
	p.t.Logf("error: "+msg, args...)
}

func (p *TestPrinter) S(msg string, args ...interface{}) {
	p.t.Logf("stdout: "+msg, args...)
}

func (p *TestPrinter) P(msg string, args ...interface{}) {
	p.t.Logf("print: "+msg, args...)
}

func (p *TestPrinter) V(msg string, args ...interface{}) {
	p.t.Logf("verbose: "+msg, args...)
}

func (p *TestPrinter) VV(msg string, args ...interface{}) {
	p.t.Logf("verbose2: "+msg, args...)
}
