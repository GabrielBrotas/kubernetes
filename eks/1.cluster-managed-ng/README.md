
1. create cluster using eksctl
```bash
# create cluster with 2 nodes t3.micro
eksctl create cluster --name cluster-test --nodegroup-name nodegroup-default --node-type t3.micro --nodes 2

# list k8s resources
kubectl get all
```

2. using yaml file
```bash
eksctl create cluster --config-file=eks/1.cluster-managed-ng/cluster.yaml

eksctl create nodegroup --config-file=eks/1.cluster-managed-ng/node-group.yaml

eksctl get nodegroup --cluster=cluster-yaml 

kubectl apply -f eks/1.cluster-managed-ng/nginx.yaml

# access pod port from localhost
kubectl port-forward pod/<pod name> 9090:80

eksctl delete cluster --name cluster-yaml
```


## EKS Nodes Regular Update
Each node is actually running in a ec2 ami;
but if k8s version needs to be updated, or we need to deploy new security patches,..
the ec2 ami needs to be updated,

in order to do that we need to create the AMIs to be used by the pods
and if want to update something in the k8s version we need to update our AMIs by ourselves
and we need to destroy the old images to update to the new ones so the application may
have some downtime


## EKS Managed Nodegroups
- Create and manage ec2 workers for you
instead of we have to create the ami with security patches and k8s version, aws releases and
manage in our behalf, and they release new version with bug fix, security patches for eks worker nodes
the ami updata has no app downtime, no overhead of user managed orchestration, auto scaling group is used behind the scenes


```bash
eksctl create cluster --name <name> --version 1.15 --nodegroup-name <nodegrpname> --node-tpe t3.micro --nodes 2 --managed
```