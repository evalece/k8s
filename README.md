# The K8s Repo

- Purpose:
    - A quick demo for K8s and Cloud service deployment with industry DevOps CD/CI operations 

- What It Does:
    - Lints and builds container images (Docker + GitHub Actions) 
        - K8s orchestration 
        - Cloud service deployment (GCP, AWS)
        - Validate builds and pushes 

- Author: 
    - lj2liu@uwaterloo.ca

- ChatGPT 4o assistance in: 
    -  Readme proofreading. 

## Development from scratch (without cloning)

### Minikube 

brew install kubectl   


### k8s 

0. K8s Quick Guide:
    - Namespace- [1]
     Kubernetes namespace 
    - Configmap are can be mounted like data volumes[4]

1. Apply Helm to enable automation on k8s compose

    - 1. Create k8s work directory with Helm     
    - 2. helm create k8s
    - 3. helm template myapp ./my-chart -f values.yaml > rendered.yaml
    - 4. kubectl apply -f rendered.yaml



2. Prometheus

helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install prometheus prometheus-community/prometheus




3. Helmfile- Centralize helm chart for release-name and chart-name (warning: use with caution)
brew install helmfile

helmfile apply

4. ArgoCD [3]- Alternative to Helmfile  but update my cluster with git push helm chart (using ArgoCD Application)

    -  Create argocd namespace
[3]
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

kubectl get secret argocd-initial-admin-secret -n argocd -o jsonpath="{.data.password}" | base64 -d

UI
kubectl port-forward svc/argocd-server -n argocd 8080:443

## Reference

Kube cluster terms
[1]https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ 
https://kubernetes.io/docs/concepts/architecture/ 

[3]https://argo-cd.readthedocs.io/en/stable/ 
[4]https://kubernetes.io/docs/concepts/configuration/configmap/ 