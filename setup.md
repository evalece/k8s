#### If at anytime local cluster becomes too heavy... 

```bash
minikube pause
```


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

- Check
```bash
helm list -A
#or 
kubectl get pods -n keda

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

5. Create k8s manifest for the Docker image 
- for example:
KEDA_prototype/log-generator.yml 
- Mannually applying the manifest:
```bash
kubectl apply -f log-generator.yaml
```

If already exist:
```bash
kubectl delete -f log-generator.yaml
```
- Consider Automation: 
```bash
cd KEDA_prototype/helm_log-gen

helm install loggen .   # "loggen" is the release name
helm upgrade loggen . # if changes
helm uninstall loggen # if do not want 

```

- To check if the service is up 
```bash 
kubectl get svc
```



#### Deploy Prometheus (Prom) to catch loggen output (check pod, cluster and cloud networking https://medium.com/google-cloud/understanding-kubernetes-networking-pods-7117dd28727 )
# Summary: in K8s cluster, each node has a bridge network holding a router gateway for overlaying logics. 
# The following implementation assumes no network restriction and Prom is able to find Loggen's IP by consulting K8s's DNS (the Core DNS).
# Need to check later: Not sure how namespace will impact DNS efficiency 
- install
```bash 
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

helm install prometheus prometheus-community/prometheus \
  --namespace monitoring --create-namespace
```
- check 
```bash
kubectl get pods -n monitoring
```
- Hook loggen to output to Prom as upstream 










#### Configuring Auto scaling for KEDA

- see  https://keda.sh/docs/2.15/reference/scaledobject-spec/ , note Horizontal Pod Autoscaler= (HPA)
```bash
KEDA_prototype/KEDA_scale/scaledobject.yaml 
kubectl apply -f scaledobject.yaml
```

- Observe auto scaling

```bash
kubectl get hpa
kubectl get pods -w

```


### Question for later

1. Will namespace impact CoreDNS efficiency? 
