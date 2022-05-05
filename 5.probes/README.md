
Just send traffic when the application is ready and working properly


## Healthz

deployment.yaml
```yaml
...
    spec:
      containers:
        - name: "goserver"
          image: gbrotas/go-hello:v7.0
```

```bash
docker build -t gbrotas/go-hello:v7.0 .
docker push gbrotas/go-hello:v7.0

kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml

kubectl port-forward service/goserver-service 8080:80
```

Even after 25 seconds that we placed to break the application the pod will still be running because we don't have a solution to verify if the application is working properly

**solution**
## Liviness Probe
We have 3 ways of check the liveness of our application
- http: call a endpoint
- cmd: execute a commando inside our container
- tcp: try to establish a tcp connection

deployment.yaml
```yaml
...
    spec:
      containers:
        - name: "goserver"
          image: gbrotas/go-hello:v7.0
          livenessProbe:
            httpGet:
              path: /healthz
              port: 4000 # container port
            periodSeconds: 5 # every 5 seconds
            # how many times it needs to throw error to restart the container
            failureThreshold: 1 # limit of failure times
            # how long this request should take to we decide that is with failure
            timeoutSeconds: 1 # if it takes 5 seconds will throw a failure
            # how many times it needs to get success to verify if the application is health
            successThreshold: 1       
```

Verify health check
```bash
## this way we can watch the pods being restarted each 30 seconds
kubectl apply -f k8s/deployment.yaml && watch -n1 kubectl get pods

kubectl describe pod <podname> # we can verify the pod history to check if was unhealth
```

## Readiness
Verify if the application is ready to receive traffic.
Sometimes the container is up and the application is not ready to receive traffic because can be creating migrations, connecting to the database or any other reason.
With readiness we can verify if the application can receive traffic.
If the application cannot receive traffic the readiness will redirect the traffic to another pod

```bash
docker build -t gbrotas/go-hello:v7.1 .
docker push gbrotas/go-hello:v7.1
```

```yaml
...
    spec:
      containers:
        ...
          readinessProbe:
            httpGet:
              path: /healthz2
              port: 4000
            periodSeconds: 5
            failureThreshold: 1 # if got one failure the application is not ready
            timeoutSeconds: 1 
            successThreshold: 1
```
Test it
```bash
kubectl apply -f k8s/deployment.yaml && watch -n1 kubectl get pods
```
The container will just be ready after the 10seconds that we set in our healthz2 to our application start

## Combining Liveness and Readiness
The way it is configured now will break our application because the liveness will keep checking the container even before the readiness say that it's ready to receive traffic, so the liveness will keep restarting our container many times.

```yaml
...
    spec:
      containers:
        - name: "goserver"
          image: gbrotas/go-hello:v7.1
          livenessProbe:
            httpGet:
              path: /healthz
              port: 4000 # container port
            periodSeconds: 5 # every 5 seconds
            # how many times it needs to throw error to restart the container
            failureThreshold: 1 # limit of failure times
            # how long this request should take to we decide that is with failure
            timeoutSeconds: 1 # if it takes 5 seconds will throw a failure
            # how many times it needs to get success to verify if the application is health
            successThreshold: 1
          
          # only access the application when ready
          readinessProbe:
            httpGet:
              path: /healthz2
              port: 4000
            periodSeconds: 5
            failureThreshold: 1 # if got one failure the application is not ready
            timeoutSeconds: 1 
            successThreshold: 1
            initialDelaySeconds: 5 # wait 5 seconds until start check if the application can receive traffic
            
```

We need to make sure readiness will execute first and then call the liveness
```yaml
...
    spec:
      containers:
        - name: "goserver"
          image: gbrotas/go-hello:v7.1
          livenessProbe:
            httpGet:
              path: /healthz
              port: 4000 # container port
            periodSeconds: 5
            # how many times it needs to throw error to restart the container
            failureThreshold: 1 # limit of failure times
            # how long this request should take to we decide that is with failure
            timeoutSeconds: 1 # if it takes 5 seconds will throw a failure
            # how many times it needs to get success to verify if the application is health
            successThreshold: 1
            initialDelaySeconds: 15 # more time so we have time to readiness verify the application

          # only access the application when ready
          readinessProbe:
            httpGet:
              path: /healthz2
              port: 4000
            periodSeconds: 3
            failureThreshold: 1 # if got one failure the application is not ready
            timeoutSeconds: 1 
            successThreshold: 1
            initialDelaySeconds: 5 # wait 5 seconds until start check if the application can receive traffic

```

Readiness will keep listening even before our application start, to make sure our container can receive traffic, if for some reason our pod broke the readiness will forward the traffic to other pods while the liveness try to create another instance

## startup Probe
The same thing as ReadinessProbe but it will just execute in the first run of the pod, and then while it is executing the liveness and readniss will keep waiting and will start to listen only after startup probe finish to verify if the application is ready

```yaml
...
    spec:
      containers:
        - name: "goserver"
          image: gbrotas/go-hello:v7.1

          startupProbe:
            httpGet:
              path: /healthz2
              port: 4000
            periodSeconds: 3
            failureThreshold: 10 # 10 * 3 = 30s, it will be listening the application for at least 30 second, if the application doesn't start after this time it will never be up            
            timeoutSeconds: 1 
            successThreshold: 1
            initialDelaySeconds: 5 # wait 5 seconds until start check if the application can receive traffic

          livenessProbe:
            httpGet:
              path: /health
              port: 4000 # container port
            periodSeconds: 5
            # how many times it needs to throw error to restart the container
            failureThreshold: 1 # limit of failure times
            # how long this request should take to we decide that is with failure
            timeoutSeconds: 1 # if it takes 5 seconds will throw a failure
            # how many times it needs to get success to verify if the application is health
            successThreshold: 1

# only access the application when ready
          readinessProbe:
            httpGet:
              path: /healthz
              port: 4000
            periodSeconds: 3
            failureThreshold: 1 # if got one failure the application is not ready
            timeoutSeconds: 1 
            successThreshold: 1
```
liviness and readiness will just start after startup say that the application is ready.
If startup say the application

```bash
docker build -t gbrotas/go-hello:v7.2 .
docker push gbrotas/go-hello:v7.2
```
```bash
kubectl apply -f k8s/deployment.yaml && watch -n1 kubectl get pods
```