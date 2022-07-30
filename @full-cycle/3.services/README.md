### Services

How do we make it so someone can access our application, which pods will that person access?

Service is the gateway to our pods, it will act like a load balancer to redirect the traffic to the pods.

```bash
# create cluster
kind create cluster --config=k8s/kind.yaml --name=multiple

# apply deployments config
kubectl apply -f k8s/deployment.yaml

# apply service config
kubectl apply -f k8s/service.yaml
kubectl get services

# we are goint to have 2 service, one is created by default from k8s and the other is what we created from our service 
#goserver     ClusterIP   10.96.150.118   <none>        80/TCP    6s

# now we can call our applications by the service name and because k8s has a service discovery system we are gonna be able to communicate with our microservices/databases/webservers without knowing their ip address

kubectl port-forward service/goserver-service 8080:80
```