//go:build tools

package tools

import (
	_ "github.com/brumhard/alligotor"
	_ "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	_ "github.com/onsi/ginkgo/v2"
	_ "github.com/onsi/gomega"
	_ "github.com/pkg/errors"
	_ "go.uber.org/automaxprocs"
	_ "go.uber.org/zap"
	_ "k8s.io/api"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/client-go/kubernetes/scheme"
	_ "sigs.k8s.io/controller-runtime"
	_ "sigs.k8s.io/controller-runtime/tools/setup-envtest"
)
