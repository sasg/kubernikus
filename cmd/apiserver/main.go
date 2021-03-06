package main

import (
	goflag "flag"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/pflag"

	apipkg "github.com/sapcc/kubernikus/pkg/api"
	"github.com/sapcc/kubernikus/pkg/api/rest"
	"github.com/sapcc/kubernikus/pkg/api/rest/operations"
	"github.com/sapcc/kubernikus/pkg/api/spec"
	logutil "github.com/sapcc/kubernikus/pkg/util/log"
	"github.com/sapcc/kubernikus/pkg/version"
)

var (
	namespace   string
	metricsPort int
)

func init() {
	pflag.StringVar(&namespace, "namespace", "kubernikus", "Namespace the apiserver should work in")
	pflag.IntVar(&metricsPort, "metrics-port", 9100, "Lister port for metric exposition")
}

func main() {
	if f := goflag.Lookup("logtostderr"); f != nil {
		f.Value.Set("true") // log to stderr by default
	}

	swaggerSpec, err := spec.Spec()
	if err != nil {
		fmt.Printf(`failed to parse swagger spec: %s`, err)
		os.Exit(1)
	}

	var server *rest.Server // make sure init is called

	pflag.Usage = func() {
		fmt.Fprint(os.Stderr, "Usage:\n")
		fmt.Fprint(os.Stderr, "  kubernikus-apiserver [OPTIONS]\n\n")

		title := "Kubernikus"
		fmt.Fprint(os.Stderr, title+"\n\n")
		desc := swaggerSpec.Spec().Info.Description
		if desc != "" {
			fmt.Fprintf(os.Stderr, desc+"\n\n")
		}
		fmt.Fprintln(os.Stderr, pflag.CommandLine.FlagUsages())
	}
	// parse the CLI flags
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine) //slurp in glog flags
	pflag.Parse()
	//goflag.CommandLine.Parse([]string{}) //https://github.com/kubernetes/kubernetes/issues/17162
	logger := logutil.NewLogger(pflag.CommandLine)

	api := operations.NewKubernikusAPI(swaggerSpec)

	rt := &apipkg.Runtime{
		Namespace: namespace,
		Logger:    logger,
	}
	rt.Kubernikus, rt.Kubernetes, err = rest.NewKubeClients(logger)
	if err != nil {
		logger.Log(
			"msg", "failed to create kubernetes clients",
			"err", err)
		os.Exit(1)
	}

	if err := rest.Configure(api, rt); err != nil {
		logger.Log(
			"msg", "failed to configure API server",
			"err", err)
		os.Exit(1)
	}
	logger.Log(
		"msg", "starting Kubernikus API",
		"namespace", namespace,
		"version", version.GitCommit)

	// get server with flag values filled out
	server = rest.NewServer(api)
	defer server.Shutdown()

	//Setup metrics listener
	metricsHost := "0.0.0.0"
	metricsListener, err := net.Listen("tcp", fmt.Sprintf("%v:%v", metricsHost, metricsPort))
	logger.Log(
		"msg", "Exposing metrics",
		"host", metricsHost,
		"port", metricsPort,
		"err", err)
	if err == nil {
		go http.Serve(metricsListener, promhttp.Handler())
		api.ServerShutdown = func() {
			metricsListener.Close()
		}
	}

	server.ConfigureAPI()
	if err := server.Serve(); err != nil {
		logger.Log(
			"msg", "failed to start API server",
			"err", err)
		os.Exit(1)
	}

}
