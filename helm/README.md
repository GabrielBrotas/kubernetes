## Problems
- Consistency Issues
- Revision History (Revision 1, 2, 3...)
- Manage files separately

## Requirements
1. k8s cluster
```bash
kind create cluster
```


## Helm
is a package management, equivalent to npm, apt, pip,...
helm is a package management in the k8s world

a package is a chart in Helm
we can install, deploy, pull charts to our package manager
ex:
```bash
helm install apache bitnami/apache --namespace=web
helm upgrade apache bitnami/apache --namespace=web
helm rollback apache 1 --namespace=web
helm uninstall apache 
```

We can also pass params to the template files using `values.yaml`

## Basic Commands
```bash
helm repo list # available repo on your machine
helm repo add bitnami https://charts.bitnami.com/bitnami # add repo

helm repo remove bitnami # remove repo from your machine

#Search the repository:

helm search repo mysql # seach package within your available repos
helm search repo database --versions

#Install a package:

helm install mydb bitnami/mysql # install on pod

helm status mydb


ROOT_PASSWORD=$(kubectl get secret --namespace default mydb-mysql -o jsonpath="{.data.mysql-root-password}" | base64 --decode)
helm upgrade --namespace default mysql-release bitnami/mysql --set auth.rootPassword=$ROOT_PASSWORD

helm uninstall mysql-release

# List all helm packages in this cluster
helm list
helm list --namespace <ns name>
helm uninstall mydb
```
