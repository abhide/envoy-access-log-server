IMAGE=envoy-als-server
IMAGE_TAG=latest
NAMESPACE=als-server
CLUSTER=cluster01

build:
	docker build -t ${IMAGE}:${IMAGE_TAG} ./

kindly-push:
	kind load docker-image ${IMAGE}:${IMAGE_TAG} --name=${CLUSTER}

kindly-deploy:
	kubectl create namespace ${NAMESPACE} || true
	kubectl apply -f k8s/deploy.yaml -n ${NAMESPACE}

clean-ns:
	kubectl delete namespace ${NAMESPACE}

all: build kindly-push kindly-deploy
