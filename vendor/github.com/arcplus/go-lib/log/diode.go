package log

import (
	"context"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"code.cloudfoundry.org/go-diodes"
	"github.com/rs/zerolog"
)

var bufPool = &sync.Pool{
	New: func() interface{} {
		return make([]byte, 0, 500)
	},
}

type Alerter func(missed int)

// Writer is a io.Writer wrapper that uses a diode to make Write lock-free,
// non-blocking and thread safe.
type Writer struct {
	lv   Level
	w    io.Writer
	d    *diodes.ManyToOne
	p    *diodes.Poller
	c    context.CancelFunc
	done chan struct{}
}

// NewAsyncWriter creates a writer wrapping w with a many-to-one diode in order to
// never block log producers and drop events if the writer can't keep up with
// the flow of data.
//
// Use a diode.Writer when
//
//     wr := diode.NewWriter(w, 1000, 10 * time.Millisecond, func(missed int) {
//         log.Printf("Dropped %d messages", missed)
//     })
//     log := zerolog.New(wr)
//
//
// See code.cloudfoundry.org/go-diodes for more info on diode.
func NewAsyncWriter(lv Level, w io.Writer, size int, poolInterval time.Duration, f Alerter) Writer {
	ctx, cancel := context.WithCancel(context.Background())
	d := diodes.NewManyToOne(size, diodes.AlertFunc(f))
	dw := Writer{
		lv: lv,
		w:  w,
		d:  d,
		p: diodes.NewPoller(d,
			diodes.WithPollingInterval(poolInterval),
			diodes.WithPollingContext(ctx)),
		c:    cancel,
		done: make(chan struct{}),
	}
	go dw.poll()
	return dw
}

func (dw Writer) Write(p []byte) (n int, err error) {
	// p is pooled in zerolog so we can't hold it passed this call, hence the
	// copy.
	p = append(bufPool.Get().([]byte), p...)
	dw.d.Set(diodes.GenericDataType(&p))
	return len(p), nil
}

func (dw Writer) WriteLevel(level Level, p []byte) (n int, err error) {
	if level < dw.lv {
		return len(p), nil
	}

	return dw.Write(p)
}

// Close releases the diode poller and call Close on the wrapped writer if
// io.Closer is implemented.
func (dw Writer) Close() error {
	dw.c()
	<-dw.done
	if w, ok := dw.w.(io.Closer); ok {
		return w.Close()
	}
	return nil
}

func (dw Writer) poll() {
	defer close(dw.done)
	for {
		d := dw.p.Next()
		if d == nil {
			return
		}
		p := *(*[]byte)(d)
		dw.w.Write(p)
		bufPool.Put(p[:0])
	}
}

// ConsoleConfig is conf for redis writer.
type ConsoleConfig struct {
	Async bool
}

func ConsoleWriter(conf ConsoleConfig) io.Writer {
	if conf.Async {
		wr := NewAsyncWriter(0, zerolog.ConsoleWriter{Out: os.Stdout}, 1000, 10*time.Millisecond, func(missed int) {
			log.Printf("Console Writer dropped %d messages", missed)
		})

		asyncWaitList = append(asyncWaitList, func() error {
			return wr.Close()
		})

		return wr
	}
	return zerolog.ConsoleWriter{Out: os.Stdout}
}