1. Start minikube at a terminal
```bash 
minikube start
```

2. Add KEDA Helm chart  (Helm charts contains K8s resources definitions to deploy K8s manifests)

```bash
helm repo add kedacore https://kedacore.github.io/charts
helm repo update
```

3. Install KEDA on MiniKube (helm install will install the chart on to the k8s cluster)
```bash
helm install keda kedacore/keda --namespace keda --create-namespace # KEDA can be namespaced, thought the CRDs it controll remain unaffected.
```
optional, if want to change namespace, remove and install again.
```bash
helm uninstall keda --namespace default
```
In case manually deleting .sh leftover requires (suggested by o4):
```bash
kubectl delete crd cloudeventsources.eventing.keda.sh \
    scaledjobs.keda.sh \
    scaledobjects.keda.sh \
    triggerauthentications.keda.sh
```

###### Creating a log generator Docker Image and Container ####

4. Create dockerfile on metric generator. 
Side note on port routings:
- Configure listening ports in dockerfiles 
- outbound port can be specified at manifest level 
see KEDA_prototype/Dockerfile 

```bash 
cd KEDA_prototype/
```
- Build Docker image inside minikube's Docker Deamon

```bash
eval $(minikube docker-env)  
docker build -t datapull:latest .
```

- To check if inside k8s Docker Daemon:
```bash
echo $DOCKER_HOST
```

- To rest to Host machine Docker env:
```bash
eval $(minikube docker-env -u)
```

5.