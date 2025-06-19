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
Assume all images have been built, configuration for extra "customization" is done via configuration (or from image's perspective, driven by container's env var).Thus, K8s focuses on deining distributed containerizations, enabling image, container and load balancing at cluster level.
Thus, in a k8s manifest, we have the following catagory: Pod (i.e., container level), Deployment + ReplicaSet (i.e., replications of pod), StatefulSet (Pod ID management), Running Jobs (i.e.,at pod level, purpose is to track if desired state is achieved, if not, define kill or rety logic with the follwoing)- Job, CronJob, DaemonSet 
    - Namespace- [1]
     Kubernetes namespace 
    - Configmap are can be mounted like data volumes[4]:
        - Basically apply a configMap pod to accept mounted files 
        - shortcut kubectl create configmap my-config --from-file=config.json
        kubectl create configmap my-config --from-file=config.json -n namespace_tag


1. Apply Helm to enable automation on k8s compose
Helm chart orchestrates k8s manifests (i.e., multiple pods specs), enabling image and k8s resources modifications
    - 1. Create k8s work directory with Helm     
    - 2. helm create k8s
    - 3. helm template myapp ./my-chart -f values.yaml > rendered.yaml
    - 4. kubectl apply -f rendered.yaml


2. How-to
    - 0. start with helm gen. dir (1) or a desinated dir with (2)
        - 1. helm install my-app ./my-chart
        - 2. helm install my-app ./my-chart

    - 1. template to create k8s manifests for declarative assignment (i.e. define desired state of k8s cluster) 
    - 2. Chart.yaml to keep track on all release-names and chart-names 
    - 3. values.yaml for k8s resources spec

    helm install myapp ./my-chart -f values.yaml




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


2. Prometheus

helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install prometheus prometheus-community/prometheus

## Reference

Kube cluster terms
[1]https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ 
https://kubernetes.io/docs/concepts/architecture/ 

[3]https://argo-cd.readthedocs.io/en/stable/ 
[4]https://kubernetes.io/docs/concepts/configuration/configmap/ 