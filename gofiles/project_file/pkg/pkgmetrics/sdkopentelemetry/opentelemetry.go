package sdkopentelemetry

import (
	"context"
	"errors"
	_log "log"
	"time"

	"connectrpc.com/otelconnect"
	prom "github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"

	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"

	// "go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"

	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"

	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Telemetry
var (
	otelInterceptor        *otelconnect.Interceptor
	Register               *prom.Registry = prom.NewRegistry()
	PROMETHEUS_CONFIG_PATH string
	GRAFANA_CONFIG_PATH    string
)

func init() {
	prepare_interceptor()
}

func GetOtelInterceptor() *otelconnect.Interceptor {
	return otelInterceptor
}

func prepare_interceptor() {
	var err error

	if otelInterceptor == nil {
		setupOTelSDK(context.Background())
		otelInterceptor, err = otelconnect.NewInterceptor()
		if err != nil {
			_log.Fatal(err)
		}
	}
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func newTracerProvider() (*trace.TracerProvider, error) {
	traceExporter, err := stdouttrace.New(
		stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, err
	}

	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter,
			// O valor padrão é 5s. Definimos em 1s para propósito de demonstração.
			trace.WithBatchTimeout(time.Second)),
	)
	return tracerProvider, nil
}

func newMeterProvider() (*metric.MeterProvider, error) {
	// metricExporter, err := stdoutmetric.New()
	// if err != nil {
	// 	return nil, err
	// }
	// Incluye registro de vars de golang
	// metricExporter, err := prometheus.New()
	// if err != nil {
	// 	return nil, err
	// }
	// para no registrar vars de golang
	metricExporter, err := prometheus.New(prometheus.WithRegisterer(Register))
	if err != nil {
		return nil, err
	}

	// 1. Definir la identidad de tu aplicación
	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("product-api"), // Nombre del servicio
			semconv.ServiceVersionKey.String("v1.0.0"),   // Versión
			semconv.DeploymentName("dev"),                // Entorno
		),
	)
	if err != nil {
		return nil, err
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithReader(metricExporter),
		metric.WithResource(res),
		//metric.WithReader(metric.NewPeriodicReader(metricExporter,
		// O valor padrão é 1m. Definimos em 3s para propósito de demonstração.
		//metric.WithInterval(10*time.Second))),
	)
	return meterProvider, nil
}

func newLoggerProvider() (*log.LoggerProvider, error) {
	logExporter, err := stdoutlog.New()
	if err != nil {
		return nil, err
	}

	loggerProvider := log.NewLoggerProvider(
		log.WithProcessor(log.NewBatchProcessor(logExporter)),
	)
	return loggerProvider, nil
}

// setupOTelSDK inicializa o pipeline do OpenTelemetry.
// Caso não retorne um erro, certifique-se de executar o método shutdown para realizar a finalização adequada.
func setupOTelSDK(ctx context.Context) (func(context.Context) error, error) {
	var shutdownFuncs []func(context.Context) error
	var err error

	// shutdown chama as funções de finalização registradas via shutdownFuncs.
	// Os erros das chamadas são concatenados.
	// Cada função de finalização registrada será invocada uma única vez.
	shutdown := func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	// handleErr chama a função shutdown para finalizar corretamente e garante que todos os erros serão retornados.
	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}

	// Inicializa o Propagator.
	prop := newPropagator()
	otel.SetTextMapPropagator(prop)

	// Inicializa o Trace Provider.
	// tracerProvider, err := newTracerProvider()
	// if err != nil {
	// 	handleErr(err)
	// 	return shutdown, err
	// }
	// shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
	// otel.SetTracerProvider(tracerProvider)

	// Inicializa o Meter Provider.
	meterProvider, err := newMeterProvider()
	if err != nil {
		handleErr(err)
		return shutdown, err
	}
	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)
	otel.SetMeterProvider(meterProvider)

	// Inicializa o Logger Provider.
	// loggerProvider, err := newLoggerProvider()
	// if err != nil {
	// 	handleErr(err)
	// 	return shutdown, err
	// }
	// shutdownFuncs = append(shutdownFuncs, loggerProvider.Shutdown)
	// global.SetLoggerProvider(loggerProvider)

	return shutdown, err
}
