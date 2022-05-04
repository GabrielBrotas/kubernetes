### Configurations
How can we store environment variables no k8s? how can we store password and secret values?

```bash
# create app image
docker build -t gbrotas/go-hello:v5 .

# push to docker hub
docker push gbrotas/go-hello:v5 

# create cluster
kind create cluster --config=k8s/kind.yaml --name=multiple

# apply deployments config
kubectl apply -f k8s/deployment.yaml

# apply service config
kubectl apply -f k8s/service.yaml

# apply configmap to second method
kubectl apply -f k8s/configmap-env.yaml

# access the pods from the service
kubectl port-forward service/goserver-service 8080:80
```