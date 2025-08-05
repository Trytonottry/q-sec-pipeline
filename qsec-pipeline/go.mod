module github.com/qsec-pipeline/qsec-pipeline

go 1.23

require (
	github.com/go-logr/logr v1.4.2
	github.com/google/go-containerregistry v0.20.2
	github.com/onsi/ginkgo/v2 v2.20.0
	github.com/onsi/gomega v1.34.1
	github.com/sigstore/cosign/v2 v2.3.0
	github.com/spf13/cobra v1.8.1
	k8s.io/api v0.31.0
	k8s.io/apimachinery v0.31.0
	k8s.io/client-go v0.31.0
	sigs.k8s.io/controller-runtime v0.19.0
	sigs.k8s.io/kubebuilder-declarative-pattern v0.15.0
)