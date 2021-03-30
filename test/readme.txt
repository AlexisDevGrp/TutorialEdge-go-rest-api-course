install kub cluster -> brew install kind
go get sigs.k8s.io/kind@v0.10.0
kind create cluster


do test -> go test ./... -tags=e2e -v

Creating cluster "kind" ...
 ✓ Ensuring node image (kindest/node:v1.20.2) 🖼 
 ✓ Preparing nodes 📦  
 ✓ Writing configuration 📜 
 ✓ Starting control-plane 🕹️ 
 ✓ Installing CNI 🔌 
 ✓ Installing StorageClass 💾 
Set kubectl context to "kind-kind"
You can now use your cluster with:

kubectl cluster-info --context kind-kind

Not sure what to do next? 😅  Check out https://kind.sigs.k8s.io/docs/user/quick-start/

1. install kube
https://v1-16.docs.kubernetes.io/docs/tasks/tools/install-kubectl/
Hi Franck (09:19 am)  ,  You are in -> test  (master *+)  $ curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 38.3M  100 38.3M    0     0  1321k      0  0:00:29  0:00:29 --:--:--  976k
Hi Franck (09:21 am)  ,  You are in -> test  (master *+)  $ chmod +x ./kubectl
Hi Franck (09:21 am)  ,  You are in -> test  (master *+)  $ sudo mv ./kubectl /usr/local/bin/kubectl
[sudo] password for admin: 
Hi Franck (09:21 am)  ,  You are in -> test  (master *+)  $ kubectl version
Client Version: version.Info{Major:"1", Minor:"20", GitVersion:"v1.20.5", GitCommit:"6b1d87acf3c8253c123756b9e61dac642678305f", GitTreeState:"clean", BuildDate:"2021-03-18T01:10:43Z", GoVersion:"go1.15.8", Compiler:"gc", Platform:"linux/amd64"}
Server Version: version.Info{Major:"1", Minor:"20", GitVersion:"v1.20.2", GitCommit:"faecb196815e248d3ecfb03c680a4507229c2a56", GitTreeState:"clean", BuildDate:"2021-01-21T01:11:42Z", GoVersion:"go1.15.5", Compiler:"gc", Platform:"linux/amd64"}
Hi Franck (09:22 am)  ,  You are in -> test  (master *+)  $ kubectl cluster-info --context kind-kind
Kubernetes control plane is running at https://127.0.0.1:33098
KubeDNS is running at https://127.0.0.1:33098/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.

then run:


kubectl cluster-info --context kind-kind

type of commans:

kubectl get pods
kubectl get services


envsubst < config/deployment.yaml > temp.yaml

docker build -t alexisdevgrp/comments-api .
docker push alexisdevgrp/comments-api:latest
envsubst < config/deployment.yml | kubectl apply -f -
kubectl gets pods
kubectl logs comments-api-6f57dd88db-6gblv
kubectl apply -f config/service.yml
kubectl get service
kubectl get endpoints
kubectl port-forward service/comments.api 9090:8080
go get github.com/dgrijalva/jwt-go
