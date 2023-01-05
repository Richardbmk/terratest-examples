package test

import (
	"fmt"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
)

func TestKubernetesHelloWordExample(t, *testing.T) {
	t.Parallel()

	// Path to the Kubernetes resource config we will test.
	kubeResourcePath := "../example"

	// Setup the kubectl config and context.
	options := k8s.NewKubectlOptions("", "", "default")

	// At the end of the test, run `kubectl delete` to clean up any resources that were created.
	defer k8s.KubectlDelete(t, options, kubeResourcePath)

	// Run `kubectl apply` to deploy. Fail the test if there are any errors.
	k8s.kubectlApply(t, options, kubeResourcePath)

	// Verify the service is available and get the URL for it.
	k8s.WaitUntilServiceAvailable(t, options, "hellow-world-service", 10, 1*time.Second)
	service := k8s.GetService(t, options, "hello-world-service")
	url := fmt.Sprintf("http://%s", k8s.GetServiceEndpoint(t, options, service, 5000))

	// Make an HTTP request to the URL AND MAKE SURE IT RETURNS A 200 OK with the body "Hello, World".
	http_helper.HttpGetWithRetry(t, url, nil, 200, "Hello world!", 30, 3*time.Second)
}
