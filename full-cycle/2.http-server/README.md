
## First step -----------------
```bash
# create app image
docker build -t gbrotas/go-hello .

# push to docker hub
docker push gbrotas/go-hello 

# create k8s cluster
kind create cluster --config=k8s/kind.yaml --name=multiple

# apply pod config on k8s
kubectl apply -f k8s/pod.yaml

# verify if the pod is running
kubectl get pods

# if we want to test our pod we can do the following:
kubectl port-forward pod/goserver 8080:80
# now we can access our localhost 8080 that will be attached to the k8s pod server
```

## Second step -----------------
```bash
# Working with replica sets to manage our pods
# remove the pod running
kubectl delete -f k8s/pod.yaml

# Apply replica set configuration
kubectl apply -f k8s/replica-set.yaml
kubectl get pods
kubectl get replicasets # verify replica sets configuration 

#if we remove a pod it wil be created another one
kubectl delede pod <podname>
```

## Third step -----------------
Deployment 

**Problem:** <br >
If we push another image to docker hub
```bash
docker build -t gbrotas/go-hello:v2 .
docker push gbrotas/go-hello:v2
```

update our replica-set.yaml
```yaml
...  
spec:
  ...
  template:
    metadata:
      labels:
        app: "goserver"
    spec:
      containers:
        - name: "goserver"
          image: gbrotas/go-hello:v2      
```
and apply the changes

```
kubectl apply -f k8s/replica-set.yaml
```

Our pods running will not be removed to update to the new version; <br/>
We can verify the pods image running with the command
```bash
kubectl describe pod <pod name>
```

we need to delete the pods manually in order to generate a new pod with the new replica-set version
```bash
kubectl delede pod <podname>
```

**Solution** <br >

**Deployment** -> **ReplicaSet** -> **Pod** <br >
The deployment will create a replica set and the replicaset will create the pods.
If we change our deployment version it will create another replicaset and delete the older version

```bash
# delete the replica set
kubectl get replicasets
kubectl delete replicasets goserver

kubectl get pods # verify if the pods has been deleted

# deploy our deployment
kubectl apply -f k8s/deployment.yaml
kubectl get deployments
kubectl get replicasets
kubectl get pods

# update the deployment.yaml version
kubectl apply -f k8s/deployment.yaml
# it will terminate older pods and generate new ones with the new versions

# our pods will be updated with 0 downtime but the replicasets will not be removed, it will just not be used

```

## Fourth step -----------------
Rollout / Rollback

if our image is broken and we applied the change on k8s deployment we can rollout to a previous history
```bash
# get the history of our deployment service
kubectl rollout history deployment goserver

# go back to the previous version
kubectl rollout undo deployment goserver

# go back to a specific version
kubectl rollout undo deployment goserver --to-revision=<revision number>

# check deployment info
kubectl describe deployment goserver
```