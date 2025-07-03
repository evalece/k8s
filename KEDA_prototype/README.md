# Kick Start on KEDA locally with Minikube

- Note that KEDA is a K8s built-in feature and should co-ordinate with provisioning in CI/CD pipelies 
- Side note on tools differentiations
    - Provisioning - configuring software 
        - Tools example : Terraform, Ansible, ArgoCD
    - Auto Scaling - Scaling up and down either horizontally or vertically, in KEDA's case, horizotally. 
        - K8s
            - Scale on pre-configured logics based on thresold 
            - Event-driven; using metrics to drive auto scaling 
            - scale on pods 
            - note that booting and shooting off mircroservice takes time

- Our Tech stack -Prometheus, Prometheus client, Graphana 
    - Prometheus
        - DB on cluster
        - handles requests on scape frequency 
        - outputs metrics 
        - our case: posting logs via client API 
            - designated where a watched out log is
    - KEDA compatibility (see https://prometheus.io/docs/introduction/overview/ for shematic)
        - Note in the schematic that data can be pulled or pushed to prom. 
        - KEDA compatible on prom pulling path (pull metrics) via Pushgateway

## Steps

- 1. Setup controllers and install if needed 
    - 1. Deploy Controller 


- 2. Deploy mircroservice
    - assign logs (customized)
    - prototype image: prometheus_client

- 3. CRDs setup-configuring driving events, auto-scaling target and thresold
    - ScaledObjects and ScaledJobs
    - Configure ScaledObject

    
    
### Definitions     
    - ScaledObject
        - Defining monitoring targets for KEDA 
        - ...and scaling targets for deployment with assigned thresold
    

# Implemntation Steps

- main.go

- KEDA
    - CRD (Custom Resource Definition), uses existing API with additional controller logics. 
     - make k8s running (helm communicates to K8s API Server)
    - Pull Helm Charts
        - Need to pull the K8s manifest format for an image (i.e, such as accessing defined env var)
        - need to do helm repo add then helm install ... because no centralized hub for helm chart
   
- Quick check 
    - minikube dashboard
- Log tracking 
    - Work around for not having free quota for Dockerhub images
    - compile main.go into binary and ship Slim Docker image only 
        - https://docs.docker.com/build/building/export/ 
  

- Manifests  https://kubernetes.io/docs/concepts/services-networking/service/  
    - recall k8s Arch:  https://kubernetes.io/docs/concepts/architecture/ 
    - Development: Pod create or delete
        - Pods
            - Owns their own IP
    - Services: IP management of Pods, using grouping service grouping logics (i.e, Backend vs Frontend)
        - ex: pods listening on one endpoint
            - define endpoint in service
            - k8s tracks
        - Endpont definition: usually pods
        - pod access policy
        - Use Selector(optional) to manage a set of pods
        - Ingress if entry point for HTTP traffic access




### Readings:
- Docker Builder
- Multi- stage builder