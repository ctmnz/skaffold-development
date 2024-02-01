# Dev Cluster

## Dependencies
 - Install [Docker for Desktop](https://www.docker.com/products/docker-desktop/)
 - Install the follwoing CLI tools
```bash
brew install go-task
brew install skaffold
brew install istioctl
brew install kind
brew install helm
```

## Installation
```
task create-kind-cluster
```
| :exclamation:  You should wait for all the pods to be ready |
|-------------------------------------------------------------|

The local cluster installation has couple of tools installed for you:
- [Kiali](http://kiali.127.0.0.1.nip.io/)
- [Zipkin](http://jaeger.127.0.0.1.nip.io/zipkin/)
- [Graphana](http://grafana.127.0.0.1.nip.io/)

## How to start the demo

### For app development demo
```
cd Skaffold
skaffold dev
```

This will follow the instructions in ./Skaffold/skaffold.yaml and will
- Build and Upload to each Kind node a docker image using the ./Skaffold/Dockerifle manifest for instructions
- Deploy the application to the kubernetes cluster following the yaml manifests in the ./Skaffold/k8s folder

When the process is started, it keeps itself attached to the console and tailing the logs from the pod

You can open [the application in your browser](https://app01.127.0.0.1.nip.io/)
Open your favorite IDE and edit the application files (i.e. ./Skaffold/main.go)
Skaffold will constanly watching for changes in your files. When triggered it will redeploy everyting descrited in the ./Skaffold/skaffold.yaml manifest.


| :exclamation:  When you press Ctrl+C you are stopping and clearing the service instaled by Skaffold |
|-----------------------------------------------------------------------------------------------------|



### For helm chart development demo
```
cd Skaffold-helm
skaffold dev
```

This will follow the instructions in ./Skaffold-helm/skaffold.yaml and will
- Build and Upload to each Kind node a docker image using the ./Skaffold-helm/Dockerifle manifest for instructions
- Deploy the helm application to the kubernetes cluster following the helm manifests in the ./Skaffold-helm/app-helm/ folder

You can open [the application in your browser](https://app01.127.0.0.1.nip.io/)
Open your favorite IDE and edit the application or the helm chart files (i.e. ./Skaffold-helm/app-helm/templates/all.yaml)
Skaffold will constanly watching for changes in your files. When triggered it will redeploy everyting descrited in the ./Skaffold-helm/skaffold.yaml manifest.

## Clean up procedures

You can press Ctrl+C to stop the `scaffold dev` process. This will clean up the resources installed by skaffold
To remove the kubernetes cluster completely execute:
```bash
task destroy-kind-cluster
```


