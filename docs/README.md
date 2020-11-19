# Introduction
> Kubeinn is an open source manager for multi-tenant [Kubernetes](https://github.com/kubernetes/kubernetes) clusters. It provides cluster administrators with the basic tools to manage tenants of a shared Kubernetes cluster. Originally designed for [SingAREN](https://www.singaren.net.sg/).

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
KUBEINN_SCHUTTERIJ_URL=http://[YOUR-KUBERNETES-NODE-IP]:30002
KUBEINN_POSTGREST_URL=http://[YOUR-KUBERNETES-NODE-IP]:30001
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

## 4. Create namespace
```bash
kubectl create namespace kubeinn
```

## 5. Install using kustomize
```
kubectl apply -k ./config
```

# Concepts
## Users
There are three types of users in Kubeinn. 

### Innkeepers
Innkeepers (cluster administrators) are responsible for managing the shared cluster and have full cluster privileges.  

Innkeepers interact with Kubeinn via the Innkeeper Administration Portal.

### Reeves 
Reeves (team representatives) are responsible for managing projects and users registered under their team. Reeves act as the point of contact for their team; any issue which requires the intervention of an innkeeper will be brought up by the reeves. 

Reeves interact with Kubeinn via the Reeve Management Portal.

### Pilgrims 
Pilgrims (team members) are the users of the cluster. Pilgrims can create and delete projects. To request for special permissions (e.g. increase resource limits), pilgrims must first communicate the request to their respective reeves, who will then raise a ticket. 

Pilgrims interact with Kubeinn via the Pilgrim User Portal.

## Registration
### Account Registration
Kubeinn is designed to be a self-service resource provisioner, meaning that innkeepers should have as little involvement in the registration process as possible. Instead, the responsibility of user management is dedicated to the reeves.

However, as there is still a need to ensure that users on the platform are legitimate, reeve accounts must be approved by an innkeeper before the team is allowed to use the cluster.

The account registration process is as follows:
1. Reeve submits a registration request.
2. Innkeeper reviews the request, and either approves or rejects the request.
3. If approved, reeves will now be able to log into Kubeinn and create pilgrims.
4. When a pilgrim is created, a registration code is generated.
5. Reeves can forward this registration code to the corresponding team member.
6. The team member, or pilgrim, can register for an account by entering the registration code provided by the reeve.

### Project Registration
Following the creation of a pilgrim account, team members can now log into Kubeinn and create projects. To create a project, select the project tab, click create, enter the project's details and click submit. 

Once a project has been created, the team member may copy the kube configuration file to the clipboard. Instructions on how to access a Kubernetes cluster using a kube configuration file can be found [here](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/#).

## Tickets
We understand that there will be times when intervention by cluster administrators is necessary. For example, project members might wish to install Custom Resource Definitions (CRDs) which may potentially affect other cluster users.

That's why we have introduced a simplified ticket management service. When a project user wishes to raise a request, the respective reeve is required to raise a ticket. The innkeeper may then view these tickets and take the necessary actions.

# FAQ
> Who can see my projects?
Kubeinn is designed with tenant isolation in mind. Innkeepers will be able to view all projects in the cluster. Reeves will only be able to view projects by pilgrims registered under their team. Pilgrims will only be able to view their own projects.

> Is Kubeinn secure?
We never stores passwords in raw text - they are always hashed. Also, we use [JSON Web Tokens](https://jwt.io/) (JWTs) which require a signing key provided by you, for authentication.