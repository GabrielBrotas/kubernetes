
Namespace creates a logical virtual separation in our environment.
Everytime we install something in our environment it will install on the default namespace.

Let's suppose we have 2 projects, we can install and manage them by the namespace.
Ex: project(namespace) 1 will have x of resource and project(namespace) 2 will have less resources.
Also we can manage the security of each one.

```bash
kubectl get namespace

kubecetl get pods -n=<namespace>

# create
kubectl create namespace <name> # kubectl create namespace dev

# apply changes
kubectl apply -f <path> -n=dev # it will only apply the changes to dev namespace
```

## Create context to use differente namespace
```bash
kubectl config view # get cluster and context config
kubect config current-context # get current context

kubectl create namespace dev

# create a context to work with different workspace
# kubectl config set-context <name> --namespace=<namespace name> --cluster=<get the cluter name from config view> --user=<get user name from config view>
kubectl config set-context dev --namespace=dev --cluster=kind-multiple --user=kind-multiple

# kubectl config use-context <context-name>
kubectl config use-context dev

# now every change that we apply it's going to apply on dev context instead of default.
kubectl apply -f k8s/deployment.yaml


# it shouldn't be deployed at default namespace
kubectl apply -f k8s/deployment.yaml -n=default
kubectl apply -f k8s/deployment.yaml # dev namespace
```

## Service Accounts
Credentials to communicate with the kubectl api.
by default k8s has one default account to deploy every resource

```bash
kubectl get serviceaccounts
```

the problem is that the service account has a master access, this account can do anything it wants.
the recommendations is to have one account for project and define just the necessary permission this account will need.

```bash
    kubectl apply -f k8s/security.yaml
```