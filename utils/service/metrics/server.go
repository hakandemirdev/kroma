package metrics

import (
	"context"
	"net"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/wemixkanvas/kanvas/utils/service/httputil"
)

func ListenAndServe(ctx context.Context, r *prometheus.Registry, hostname string, port int) error {
	addr := net.JoinHostPort(hostname, strconv.Itoa(port))
	server := &http.Server{
		Addr: addr,
		Handler: promhttp.InstrumentMetricHandler(
			r, promhttp.HandlerFor(r, promhttp.HandlerOpts{}),
		),
	}
	return httputil.ListenAndServeContext(ctx, server)
}
