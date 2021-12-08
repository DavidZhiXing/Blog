package tcp_limit

type LimitReader struct {
	R io.Reader
	Limiter *rate.Limiter
}

func (l *LimitReader) Read(p []byte) (n int, err error) {
	if l.Limiter != nil {
		l.Limiter.Wait(context.Background())
	}
	return l.R.Read(p)
}

func NewLimitReader(r io.Reader, limiter *rate.Limiter) *LimitReader {
	return &LimitReader{r, limiter}
}

// test
func main() {
	// Create a new limiter with a rate of 1 request per second.
	limiter := rate.NewLimiter(rate.Limit(1), 1)

	// Create a new LimitReader with the limiter.
	reader := NewLimitReader(os.Stdin, limiter)

	// Read from stdin, and print to stdout.
	io.Copy(os.Stdout, reader)
}