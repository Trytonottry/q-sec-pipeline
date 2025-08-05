package main

import (
	"flag"
	"os"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/qsec-pipeline/qsec-pipeline/internal/controller"
)

func main() {
	var metricsAddr string
	flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "metrics addr")
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseDevMode(true)))
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{Metrics: metricsAddr})
	if err != nil { os.Exit(1) }

	if err = (&controller.QSecPolicyReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil { os.Exit(1) }

	if err = mgr.Start(ctrl.SetupSignalHandler()); err != nil { os.Exit(1) }
}