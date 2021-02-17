# test
```
kubectl exec -n <namespace> --stdin --tty <pod-name> -- /bin/bash
kubectl exec -n pilgrim-1-project-1 --stdin --tty alpine-same-ns -- /bin/sh
kubectl exec -n pilgrim-1-project-2 --stdin --tty alpine-different-ns -- /bin/sh
kubectl exec -n kubeinn --stdin --tty alpine-admin-ns -- /bin/sh

kubectl get pods -n pilgrim-1-project-1 -o wide
wget -qO- --timeout=2 <IP ADDRESS OF NGINX POD>
wget -qO- --timeout=2 192.168.141.180
```

# Add to instructions
kubectl label namespace/default kubeinn=admin
kubectl label namespace/kube-node-lease kubeinn=admin
kubectl label namespace/kube-public kubeinn=admin
kubectl label namespace/kube-system kubeinn=admin
kubectl label namespace/kubeinn kubeinn=admin
kubectl label namespace/rook-ceph kubeinn=admin

