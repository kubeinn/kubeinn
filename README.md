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
kubectl kustomize ./config > generated_config.yaml
kubectl apply -k ./config
```

## Uninstall
```
kubectl delete -k ./configs
```

## Debugging
```
kubectl exec --stdin --tty -n kubeinn kubeinn-postgres-deployment-v-1-0-0-a-1-86d95cdf8f-gxksm -- /bin/bash
psql -U postgres
kubectl logs kubeinn-postgres-deployment-v-1-0-0-a-1-6bcb5f84c-jvkrt -n kubeinn

May need to wait for pvc to be bound before creating Postgres deployment

kubeinn-postgres-deployment-v-1-0-0-a-1-6bcb5f84c-jx77r
```