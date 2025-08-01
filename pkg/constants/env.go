// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package constants

const (
	EnvOTELServiceName      = "OTEL_SERVICE_NAME"
	EnvOTELResourceAttrs    = "OTEL_RESOURCE_ATTRIBUTES"
	EnvOTELPropagators      = "OTEL_PROPAGATORS"
	EnvOTELTracesSampler    = "OTEL_TRACES_SAMPLER"
	EnvOTELTracesSamplerArg = "OTEL_TRACES_SAMPLER_ARG"

	EnvOTELExporterOTLPEndpoint      = "OTEL_EXPORTER_OTLP_ENDPOINT"
	EnvOTELExporterCertificate       = "OTEL_EXPORTER_OTLP_CERTIFICATE"
	EnvOTELExporterClientCertificate = "OTEL_EXPORTER_OTLP_CLIENT_CERTIFICATE"
	EnvOTELExporterClientKey         = "OTEL_EXPORTER_OTLP_CLIENT_KEY"

	InstrumentationPrefix                           = "instrumentation.opentelemetry.io/"
	AnnotationDefaultAutoInstrumentationJava        = InstrumentationPrefix + "default-auto-instrumentation-java-image"
	AnnotationDefaultAutoInstrumentationNodeJS      = InstrumentationPrefix + "default-auto-instrumentation-nodejs-image"
	AnnotationDefaultAutoInstrumentationPython      = InstrumentationPrefix + "default-auto-instrumentation-python-image"
	AnnotationDefaultAutoInstrumentationDotNet      = InstrumentationPrefix + "default-auto-instrumentation-dotnet-image"
	AnnotationDefaultAutoInstrumentationGo          = InstrumentationPrefix + "default-auto-instrumentation-go-image"
	AnnotationDefaultAutoInstrumentationApacheHttpd = InstrumentationPrefix + "default-auto-instrumentation-apache-httpd-image"
	AnnotationDefaultAutoInstrumentationNginx       = InstrumentationPrefix + "default-auto-instrumentation-nginx-image"

	LabelTargetAllocator                         = "opentelemetry.io/target-allocator"
	ResourceAttributeAnnotationPrefix            = "resource.opentelemetry.io/"
	KubernetesLastAppliedConfigurationAnnotation = "^kubectl\\.kubernetes\\.io/last-applied-configuration$"

	EnvPodName  = "OTEL_RESOURCE_ATTRIBUTES_POD_NAME"
	EnvPodUID   = "OTEL_RESOURCE_ATTRIBUTES_POD_UID"
	EnvPodIP    = "OTEL_POD_IP"
	EnvNodeName = "OTEL_RESOURCE_ATTRIBUTES_NODE_NAME"
	EnvNodeIP   = "OTEL_NODE_IP"

	TACollectorTLSDirPath      = "/tls"
	TACollectorCAFileName      = "ca.crt"
	TACollectorTLSKeyFileName  = "tls.key"
	TACollectorTLSCertFileName = "tls.crt"
)

var (
	LabelAppName = []string{
		"app.kubernetes.io/instance",
		"app.kubernetes.io/name",
	}
	LabelAppVersion = []string{"app.kubernetes.io/version"}
)
