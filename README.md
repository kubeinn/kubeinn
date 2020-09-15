# kubeinn

## Prerequisites
```
# Set default storage class
kubectl patch storageclass rook-cephfs  -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"true"}}}'
# Ingress controller installed
## e.g. Using traefik ingress controller
helm repo add traefik https://containous.github.io/traefik-helm-chart
helm repo update
helm install traefik traefik/traefik
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
kubectl exec --stdin --tty -n kubeinn <POD> -- /bin/bash
psql -U postgres

kubectl run pgsql-postgresql-client --rm --tty -i --restart='Never' --namespace kubeinn --image postgres:13 --env="PGPASSWORD=pgpassword" --command -- psql postgres --host acrp-postgres-cluster-ip-service-v-1-0-0-a-1 -U postgres -d postgres -p 5432
```