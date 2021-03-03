# envoy-access-log-server
Sample implementation for Envoy AccessLog Server.

Starts a gRPC Envoy AccessLog Server using v3 xDS and emits access logs entries to stdout.
Both HTTP and TCP access logs are supported

# Pre-Requisites:
- Docker
- [kind-cluster](https://github.com/abhide/kind-clusters)

# Deploy Access Log Server
```bash
➜  make all
docker build -t envoy-als-server:latest ./
Sending build context to Docker daemon  15.11MB
Step 1/10 : FROM golang:alpine3.12
alpine3.12: Pulling from library/golang
f84cab65f19f: Pull complete 
882e2a9d04d9: Pull complete 
3fd3821a34fb: Pull complete 
c41b5234fec1: Pull complete 
3b5ab1f9fb69: Pull complete 
Digest: sha256:9f18292b374d40c6547422c8979c694fa9654e609d2d6694722a94098a403b2c
Status: Downloaded newer image for golang:alpine3.12
 ---> 8d4cbc6fcb0f
Step 2/10 : WORKDIR /go/src/github.com/abhide/envoy-access-log-server/
 ---> Running in 43c29fc7dd10
Removing intermediate container 43c29fc7dd10
 ---> 605e839295f3
Step 3/10 : COPY main.go .
 ---> 2db1346e5f00
Step 4/10 : COPY go.mod .
 ---> e3c22f7f944d
Step 5/10 : COPY go.sum .
 ---> f7354d419944
Step 6/10 : RUN go build -o envoy-access-log-server ./main.go
 ---> Running in 6b89fdde093f
go: downloading github.com/envoyproxy/go-control-plane v0.9.9-0.20201210154907-fd9021fe5dad
go: downloading google.golang.org/grpc v1.35.0
go: downloading golang.org/x/net v0.0.0-20190311183353-d8887717615a
go: downloading github.com/golang/protobuf v1.4.2
go: downloading golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a
go: downloading google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
go: downloading google.golang.org/protobuf v1.25.0
go: downloading golang.org/x/text v0.3.0
go: downloading github.com/envoyproxy/protoc-gen-validate v0.1.0
go: downloading github.com/cncf/udpa/go v0.0.0-20201120205902-5459f2c99403
Removing intermediate container 6b89fdde093f
 ---> a6d46ca58e49
Step 7/10 : FROM alpine:3.12
3.12: Pulling from library/alpine
f84cab65f19f: Already exists 
Digest: sha256:a295107679b0d92cb70145fc18fb53c76e79fceed7e1cf10ed763c7c102c5ebe
Status: Downloaded newer image for alpine:3.12
 ---> 88dd2752d2ea
Step 8/10 : WORKDIR /root/
 ---> Running in 8847590c986b
Removing intermediate container 8847590c986b
 ---> 53679f3ecd74
Step 9/10 : COPY --from=0 /go/src/github.com/abhide/envoy-access-log-server/envoy-access-log-server .
 ---> fe46a83b5c66
Step 10/10 : CMD ["./envoy-access-log-server"]
 ---> Running in 02845dd35cea
Removing intermediate container 02845dd35cea
 ---> c145a00ce3cc
Successfully built c145a00ce3cc
Successfully tagged envoy-als-server:latest
kind load docker-image envoy-als-server:latest --name=cluster01
Image: "envoy-als-server:latest" with ID "sha256:c145a00ce3ccac0f1ec343c99d3940c4ec93ab098123a4db00d81431956b9b37" not yet present on node "cluster01-control-plane", loading...
kubectl create namespace als-server || true
namespace/als-server created
kubectl apply -f k8s/deploy.yaml -n als-server
deployment.apps/envoy-als-server created
service/envoy-als-server-svc created

➜  kubectl get pods -n als-server
NAME                               READY   STATUS    RESTARTS   AGE
envoy-als-server-7967d969d-wjq4c   1/1     Running   0          10s

➜  kubectl get svc -n als-server 
NAME                   TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
envoy-als-server-svc   ClusterIP   10.96.168.117   <none>        8080/TCP   14s

➜  kubectl logs envoy-als-server-7967d969d-wjq4c -n als-server
2021/03/03 06:22:59 Starting ALS Server
```
# Delete Access Log Server
```bash
➜  make clean-ns
kubectl delete namespace als-server
namespace "als-server" deleted

```