package api

// Logger defines the contract for logging.
type Logger interface {
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Panicf(template string, args ...interface{})
}

// A Counter represents a monotonically increasing value.
type Counter interface {
	// With is used to provide label values when updating a Counter. This must be
	// used to provide values for all LabelNames provided to CounterOpts.
	With(labelValues ...string) Counter

	// Add increments a counter value.
	Add(delta float64)
}

// A Gauge is a meter that expresses the current value of some metric.
type Gauge interface {
	// With is used to provide label values when recording a Gauge value. This
	// must be used to provide values for all LabelNames provided to GaugeOpts.
	With(labelValues ...string) Gauge

	// Add increments a Gauge value.
	Add(delta float64)

	// Set is used to update the current value associated with a Gauge.
	Set(value float64)
}

// A Histogram is a meter that records an observed value into quantized
// buckets.
type Histogram interface {
	// With is used to provide label values when recording a Histogram
	// observation. This must be used to provide values for all LabelNames
	// provided to HistogramOpts.
	With(labelValues ...string) Histogram
	Observe(value float64)
}

// Diagnostics contains set of various metrics and logging interfaces
type Diagnostics struct {
	Logger    Logger
	Counter   Counter
	Gauge     Gauge
	Histogram Histogram
}

// WrapNils returns Diagnostics with mocked empty implemetations
func (d Diagnostics) WrapNils() (out Diagnostics) {
	out = d

	if d.Logger == nil {
		out.Logger = EmptyLogger{}
	}

	if d.Counter == nil {
		out.Counter = EmptyCounter{}
	}

	if d.Gauge == nil {
		out.Gauge = EmptyGauge{}
	}

	if d.Histogram == nil {
		out.Histogram = EmptyHistogram{}
	}

	return
}

// --------------------------------------

// EmptyLogger is used to prevent nil interface errors
type EmptyLogger struct{}

func (EmptyLogger) Debugf(string, ...interface{}) {}
func (EmptyLogger) Infof(string, ...interface{})  {}
func (EmptyLogger) Errorf(string, ...interface{}) {}
func (EmptyLogger) Warnf(string, ...interface{})  {}
func (EmptyLogger) Panicf(string, ...interface{}) {}

// EmptyCounter is used to prevent nil interface errors
type EmptyCounter struct{}

func (EmptyCounter) With(...string) Counter { return EmptyCounter{} }
func (EmptyCounter) Add(float64)            {}

// EmptyGauge is used to prevent nil interface errors
type EmptyGauge struct{}

func (EmptyGauge) With(...string) Gauge { return EmptyGauge{} }
func (EmptyGauge) Add(float64)          {}
func (EmptyGauge) Set(float64)          {}

// EmptyHistogram is used to prevent nil interface errors
type EmptyHistogram struct{}

func (EmptyHistogram) With(...string) Histogram { return &EmptyHistogram{} }
func (EmptyHistogram) Observe(float64)          {}

// --------------------------------------
