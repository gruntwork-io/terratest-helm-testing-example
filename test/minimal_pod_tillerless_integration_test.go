package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/random"
)

func TestPodDeploysContainerImageTillerless(t *testing.T) {
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
	// - fullnameOverride=minimal-pod-RANDOM_STRING
	// We use a fullnameOverride so we can find the Pod later during verification
	podName := fmt.Sprintf("minimal-pod-%s", strings.ToLower(random.UniqueId()))
	options := &helm.Options{
		SetValues: map[string]string{"image": "nginx:1.15.8", "fullnameOverride": podName},
	}

	// Run RenderTemplate to render the template and capture the output.
	output := helm.RenderTemplate(t, options, helmChartPath, []string{})

	// Make sure to delete the resources at the end of the test
	defer k8s.KubectlDeleteFromString(t, kubectlOptions, output)

	// Now use kubectl to apply the rendered template
	k8s.KubectlApplyFromString(t, kubectlOptions, output)

	// Now that the chart is deployed, verify the deployment. This function will open a tunnel to the Pod and hit the
	// nginx container endpoint.
	verifyNginxPod(t, kubectlOptions, podName)
}
