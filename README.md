# kubeinn

## Prerequisites
```
# Set default storage class
kubectl patch storageclass rook-cephfs  -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"true"}}}'
```

## Installation
```
# Create namespace
kubectl create namespace kubeinn
```

## Kustomize
```
kubectl kustomize ./configs > generated_config.yaml
kubectl apply -k ./configs
```

## Debugging
```
kubectl exec --stdin --tty -n acrp <POD> -- /bin/bash

kubectl run pgsql-postgresql-client --rm --tty -i --restart='Never' --namespace kubeinn --image postgres:13 --env="PGPASSWORD=pgpassword" --command -- psql postgres --host acrp-postgres-cluster-ip-service-v-1-0-0-a-1 -U postgres -d postgres -p 5432
```