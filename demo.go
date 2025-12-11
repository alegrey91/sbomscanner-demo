package main

import (
	"github.com/saschagrunert/demo"
)

func main() {
	// Create a new demo CLI application
	d := demo.New()

	// Set demo properties
	d.Name = "Kubernetes Image Security Scanning Demo"
	d.Usage = "Demonstrates container image scanning with ScanJobs, SBOMs, and VEX integration"
	d.HideVersion = true

	// Register the demo run
	d.Add(scanDemo(), "scan-demo", "Complete security scanning workflow")

	// Run the application
	d.Run()
}

// scanDemo is the main demo run
func scanDemo() *demo.Run {
	r := demo.NewRun(
		"Kubernetes Security Scanning with VEX",
		"This demo shows how to:",
		"1. Push and scan container images",
		"2. Generate SBOMs and vulnerability reports",
		"3. Integrate VEX data to filter vulnerabilities",
	)

	// Step 1: Push image to local registry
	r.Step(demo.S(
		"Push the test image to the local registry",
	), demo.S(
		"docker push localhost:5000/test-image:1.0",
	))

	// Step 2: Verify image was pushed
	r.Step(demo.S(
		"Ensure the image was pushed properly",
	), demo.S(
		"curl -X GET http://localhost:5000/v2/_catalog | jq",
	))

	// Step 3: Check registry manifest
	r.Step(demo.S(
		"Check the registry manifest",
	), demo.S(
		"cat examples/registry.yaml",
	))

	// Step 4: Apply registry manifest
	r.Step(demo.S(
		"Apply the registry manifest to Kubernetes",
	), demo.S(
		"kubectl apply -f examples/registry.yaml",
	))

	// Step 5: Verify registry resource
	r.Step(demo.S(
		"Ensure the registry resource was created",
	), demo.S(
		"kubectl get registry",
	))

	// Step 6: Check scanjob manifest
	r.Step(demo.S(
		"Check the scanjob manifest",
	), demo.S(
		"cat examples/scanjob.yaml",
	))

	// Step 7: Apply scanjob manifest
	r.Step(demo.S(
		"Apply the scanjob manifest to initiate scanning",
	), demo.S(
		"kubectl apply -f examples/scanjob.yaml",
	))

	// Step 8: Verify scanjob resource
	r.Step(demo.S(
		"Ensure the scanjob resource was created",
	), demo.S(
		"kubectl get scanjob",
	))

	// Step 9: Manual check logs
	r.Step(demo.S(
		"[Manual] Open k9s and navigate to the scanjob logs",
		"Press Enter to continue after checking the logs",
	), nil)

	// Step 10: Verify SBOM creation
	r.Step(demo.S(
		"Ensure the SBOM (Software Bill of Materials) was created",
	), demo.S(
		"kubectl get sboms",
	))

	// Step 11: Show SBOM details
	r.Step(demo.S(
		"Show the SBOM resource details",
	), demo.S(
		"kubectl get sboms c8f387632342cf9b250ffe9db03a87d1b2cbbf4bbdb231e02fc10faabd537453 -o yaml | less",
	))

	// Step 12: Verify vulnerability report creation
	r.Step(demo.S(
		"Verify the vulnerability report was created",
	), demo.S(
		"kubectl get vulnerabilityreports",
	))

	// Step 13: Show vulnerability report details
	r.Step(demo.S(
		"Show the vulnerability report resource",
	), demo.S(
		"kubectl get vulnerabilityreports -o=jsonpath='{.items[0].report.summary}' | jq",
	))

	// Step 14: Show SBOM details
	r.Step(demo.S(
		"Show the VulnerabilityReport resource",
	), demo.S(
		"kubectl get vulnerabilityreport c8f387632342cf9b250ffe9db03a87d1b2cbbf4bbdb231e02fc10faabd537453 -o yaml | less",
	))

	// Step 15: Delete scanjob
	r.Step(demo.S(
		"Clean up the initial scanjob",
	), demo.S(
		"kubectl delete scanjob my-first-scanjob",
	))

	// Step 16: Clear screen
	r.Step(demo.S(
		"Clear the terminal for the VEX demonstration",
	), demo.S(
		"clear",
	))

	// Step 17: Check VEXHub manifest
	r.Step(demo.S(
		"Check the VEXHub manifest",
		"VEX (Vulnerability Exploitability eXchange) provides",
		"additional context about vulnerabilities",
	), demo.S(
		"cat examples/vexhub-aquasec.yaml",
	))

	// Step 18: Apply VEXHub manifest
	r.Step(demo.S(
		"Apply the VEXHub manifest to integrate VEX data",
	), demo.S(
		"kubectl apply -f examples/vexhub-aquasec.yaml",
	))

	// Step 19: Verify VEXHub resource
	r.Step(demo.S(
		"Verify the VEXHub resource was created",
	), demo.S(
		"kubectl get vexhub",
	))

	// Step 20: Re-apply scanjob
	r.Step(demo.S(
		"Re-apply the scanjob to run scan with VEX integration",
	), demo.S(
		"kubectl apply -f examples/scanjob.yaml",
	))

	// Step 21: Show SBOM details
	r.Step(demo.S(
		"Show the VulnerabilityReport resource",
	), demo.S(
		"kubectl get vulnerabilityreport c8f387632342cf9b250ffe9db03a87d1b2cbbf4bbdb231e02fc10faabd537453",
	))

	// Step 22: Show vulnerability report (with VEX)
	r.Step(demo.S(
		"Show the vulnerability report with VEX status",
	), demo.S(
		"kubectl get vulnerabilityreports -o=jsonpath='{.items[0].report.summary}' | jq",
	))

	return r
}
