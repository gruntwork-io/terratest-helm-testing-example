package test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/random"
)

func TestPodDeploysContainerImage(t *testing.T) {
	// Path to the helm chart we will test
	helmChartPath := "../charts/minimal-pod"

	// Setup the kubectl config and context. Here we choose to use the defaults, which is:
	// - HOME/.kube/config for the kubectl config file
	// - Current context of the kubectl config file
	// We also specify that we are working in the default namespace (required to get the Pod)
	kubectlOptions := k8s.NewKubectlOptions("", "")
	kubectlOptions.Namespace = "default"

	// Setup the args. For this test, we will set the following input values:
	// - image=nginx:1.15.8
	options := &helm.Options{
		SetValues: map[string]string{"image": "nginx:1.15.8"},
	}

	// We generate a unique release name so that we can refer to after deployment.
	// By doing so, we can schedule the delete call here so that at the end of the test, we run
	// `helm delete RELEASE_NAME` to clean up any resources that were created.
	releaseName := fmt.Sprintf("nginx-%s", strings.ToLower(random.UniqueId()))
	defer helm.Delete(t, options, releaseName, true)

	// Deploy the chart using `helm install`. Note that we use the version without `E`, since we want to assert the
	// install succeeds without any errors.
	helm.Install(t, options, helmChartPath, releaseName)

	// Now that the chart is deployed, verify the deployment. This function will open a tunnel to the Pod and hit the
	// nginx container endpoint.
	podName := fmt.Sprintf("%s-minimal-pod", releaseName)
	verifyNginxPod(t, kubectlOptions, podName)
}

// verifyNginxPod will open a tunnel to the Pod and hit the endpoint to verify the nginx welcome page is shown.
func verifyNginxPod(t *testing.T, kubectlOptions *k8s.KubectlOptions, podName string) {
	// Wait for the pod to come up. It takes some time for the Pod to start, so retry a few times.
	retries := 15
	sleep := 5 * time.Second
	k8s.WaitUntilPodAvailable(t, kubectlOptions, podName, retries, sleep)

	// We will first open a tunnel to the pod, making sure to close it at the end of the test.
	tunnel := k8s.NewTunnel(kubectlOptions, k8s.ResourceTypePod, podName, 0, 80)
	defer tunnel.Close()
	tunnel.ForwardPort(t)

	// ... and now that we have the tunnel, we will verify that we get back a 200 OK with the nginx welcome page.
	// It takes some time for the Pod to start, so retry a few times.
	endpoint := fmt.Sprintf("http://%s", tunnel.Endpoint())
	http_helper.HttpGetWithRetryWithCustomValidation(
		t,
		endpoint,
		retries,
		sleep,
		func(statusCode int, body string) bool {
			return statusCode == 200 && strings.Contains(body, "Welcome to nginx")
		},
	)
}
