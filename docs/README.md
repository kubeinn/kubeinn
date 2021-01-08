# Introduction
> Kubeinn is an open source manager for multi-tenant [Kubernetes](https://github.com/kubernetes/kubernetes) clusters. It provides cluster administrators with the basic tools to manage tenants of a shared Kubernetes cluster.

![](./img/kubeinn-demo.gif)

# Installation
You can deploy Kubeinn on any Kubernetes 1.9+ cluster in a matter of minutes, if not seconds. 

## Prerequisites
- Kubernetes version 1.9 or higher
- Persistent Volume Claims

## 1. Clone this repository
```bash
git clone https://github.com/kubeinn/kubeinn.git
cd kubeinn
```

## 2. Set environment variables
```
# ./config/configmaps/frontend/.env
KUBEINN_SCHUTTERIJ_URL=http://[YOUR-KUBERNETES-NODE-IP]:[YOUR-KUBERNETES-NODE-PORT]
```

## 3. Set kustomization secrets
```
# ./config/kustomization.yaml
secretGenerator:
  - name: pgpassword
    literals:
      - POSTGRES_PASSWORD=[YOUR-POSTGRES-PASSWORD]
  - name: jwt-signing-key
    literals:
      - JWT_SIGNING_KEY=[YOUR-256-BITS-JWT-SIGNING-KEY]
```

## 4. Copy kube config
```bash
# kube config is usually located at /root/.kube/config
# May differ according to your cloud provider
cp /root/.kube/config config/configmaps/backend/admin-config
```

## 5. Create namespace
```bash
kubectl create namespace kubeinn
```

## 6. Install using kustomize
```bash
kubectl apply -k ./config
```

# Concepts
## Users
There are two types of users in Kubeinn. 

### Innkeepers
Innkeepers (cluster administrators) are responsible for managing the shared cluster and have full cluster privileges.  

Innkeepers interact with Kubeinn via the Innkeeper Administration Portal.

### Pilgrims 
Pilgrims (project supervisors) are the users of the cluster. While pilgrims can share access to the cluster with other project members, they must be responsible for managing users registered under their projects. Pilgrims act as the point of contact for their projects. Any issue which requires the intervention of a cluster administrator must be brought up by the pilgrims.

Pilgrims can create and delete projects. To request for special permissions (e.g. increase resource limits), pilgrims raise a ticket which an innkeeper will attend to subsequently. 

Pilgrims interact with Kubeinn via the Pilgrim User Portal.

## Registration
### Account Registration
Kubeinn is designed to be a self-service resource provisioner, meaning that innkeepers should have as little involvement in the registration process as possible. However, as there is still a need to ensure that users on the platform are legitimate, pilgrim accounts must be approved by an innkeeper before the project supervisor is allowed to use the cluster.

The responsibility of managing users under a project is dedicated to the pilgrims, as they can choose who they wish to share their access configuration with.

The account registration process is as follows:
1. Pilgrim submits a registration request.
2. Innkeeper reviews the request, and either approves or rejects the request.
3. If approved, pilgrims will now be able to log into Kubeinn.

### Project Registration
Following the creation of a pilgrim account, pilgrims can create projects. To create a project, select the project tab, click create, enter the project's details and click submit. Once a project has been created, the team member may copy the kube configuration file to the clipboard. 

Instructions on how to access a Kubernetes cluster using a kube configuration file can be found [here](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/#).

The kube configuration file may be shared with other members of the project.

## Tickets
We understand that there will be times when intervention by cluster administrators is necessary. For example, project members might wish to install Custom Resource Definitions (CRDs) which may potentially affect other cluster users.

That's why we have introduced a simplified ticket management service. A project supervisor may raise a ticket, while innkeepers may view these tickets and take the necessary actions.

# FAQ
> Who can see my projects?
Kubeinn is designed with tenant isolation in mind. Innkeepers will be able to view all projects in the cluster. Pilgrims will only be able to view projects registered under their account.

> Is Kubeinn secure?
We never stores passwords in raw text - they are always hashed. Also, we use [JSON Web Tokens](https://jwt.io/) (JWTs) which require a signing key provided by you, for authentication.