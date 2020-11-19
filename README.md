# PoC of a service and infrastructure around it

* The simplest server written in Go
* Dockerfile to create the image:
    * `docker build -t poc .`

## Deploy from a local image
* Deploy service from a local image (without an online repository):
    * `minikube start`
    * `eval $(minikube -p minikube docker-env)`
    * `kubectl create -f poc.yml`
    * [source](https://medium.com/swlh/how-to-run-locally-built-docker-images-in-kubernetes-b28fbc32cc1d)
    