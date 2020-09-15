# kubeinn

## Prerequisites
```
# Create namespace
kubectl create namespace kubeinn

# Set default storage class
kubectl patch storageclass rook-cephfs  -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"true"}}}'

# Ingress controller installed
## e.g. Using traefik ingress controller
helm repo add traefik https://containous.github.io/traefik-helm-chart
helm repo update
helm install --namespace kubeinn traefik traefik/traefik --values ./configs/traefik/values.yaml
```

## Installation
```
# Create namespace
kubectl create namespace kubeinn
kubectl kustomize ./configs/kustomize > generated_config.yaml
kubectl apply -k ./configs/kustomize
```

## Debugging
```
kubectl exec --stdin --tty -n kubeinn kubeinn-postgres-deployment-v-1-0-0-a-1-6bcb5f84c-mwxhk -- /bin/bash
psql -U postgres
kubectl logs kubeinn-postgrest-deployment-v-1-0-0-a-1-655c5b8ff8-gg6qv -n kubeinn

May need to wait for pvc to be bound before creating Postgres deployment

kubeinn-postgres-deployment-v-1-0-0-a-1-6bcb5f84c-jx77r
```