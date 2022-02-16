package middlewares

import (
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	opsProcessed *prometheus.CounterVec
}

func NewMetrics() *Metrics {
	opsProcessed := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "accounts_ops_processed_total",
		Help: "Total of operations processed by accounts-core-api ",
	}, []string{"method", "path", "statuscode"})
	return &Metrics{opsProcessed: opsProcessed}
}
func (m *Metrics) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wi := &responseWriterInterceptor{
			statusCode:     http.StatusOK,
			ResponseWriter: w,
		}
		next.ServeHTTP(wi, r)

		m.opsProcessed.With(prometheus.Labels{"method": r.Method, "path": r.RequestURI, "statuscode": strconv.Itoa(wi.statusCode)}).Inc()
	})
}

type responseWriterInterceptor struct {
	http.ResponseWriter
	statusCode int
}
