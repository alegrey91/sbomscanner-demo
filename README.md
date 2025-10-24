# SBOMscanner Demo

## Prerequisites

The following tools are required:

* k9s
* yq
* kubectl

Build the demo program:

```shell
go build .
```

## Demo

Run these commands before starting the demo:

```shell
docker pull ghcr.io/aquasecurity/trivy:0.52.0
docker tag localhost:5000/test-image:1.0
```

Run the demo:

```shell
./sbomscanner-demo --scan-demo
```
