How to check the metrics of our pods? how can we know if we need more pods or how to know if our pods are holding the load?

HPA - Horizontal Pod Autoscaling

### Metrics Server
It collects real time data of how much resource each pod from k8s is consuming at that moment.
If we want we can store this data in anothers tools such as promethous with grafana that generates dashboard to us.

In a cloud provider (EKS, AKS,...) this metric server is configured by default but using with kind we need to setup everything.

## Setup with kind
```bash
cd k8s
wget https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml

mv components.yaml metrics-server.yaml

# add the flag at line 140
# --kubelet-insecure-tls 

kubectl apply -f k8s/metrics-server.yaml

kubectl get apiservices # make sure it created kube-system/metrics-server
```

## Resources
In order to setup and monitor our services we need to define the resources our pods will need.
How much our pods needs to run and what is the limit we will allow it to take

vCPU => defines the processing power of the machine, does the computing work
    1 vCPU -> 1000 milicores
    500m, the half of 1 vCPU
    0.5, percentage, half of our vCPU

Memory => holds the data

deployment.yaml
```yaml
...
    spec:
      containers:
        - name: "goserver"
          image: gbrotas/go-hello:v7.2

          # we are reserving this resource from our cluster for our pod
          resources:
            # the minimum our container(pod) needs to run
            requests: 
              cpu: 128m # unit => vCPU
              memory: 100
        # how far this pod can use a resource
            limits: 
              cpu: 256Mi
              memory: 200
...
```

When we say a resource needs some cpu and memory it will reserve from the cluster and no other pod will be able to take that slice we defined.
which node has 128m cpu and 100m memory, if is the Node1 the k8s scheduler will get this space from the Node1 to create our pod.

If none of our Nodes has cpu and memory to meet the pod requirements the pod will be at pending stage waiting some machine to be added to the cluster so it will be able to be deployed or another pod die to release resource.

**observation**
Our cluster has a limited amount of resource, ex: 3000m CPU
So you probably wondering that we can have 30 pods with 100m CPU to use, 30 * 100 = 3000mCPU
but, sometimes our pods will receive access spikes and because our pods limits are 256Mi all of them will try to get more space from the cluster and that will not be possible because 30 * 256 = 7680mCPU
we need to avoid that.
**tip** => the sum of the pods limits should not exceed the machine limit.

```bash
kubectl apply -f k8s/deployment.yaml

# verify the pod usage (cpu and memory)
kubectl top pod <podname>  
```

How to scale ?

## HPA - Horizontal Pod Autoscaling
Responsible to get the metrics and analyse the cpu, memory, traffic and based on this data it will provision more replicas (scale in and out)

most of the case we are going to use CPU to verify if we need more replicas or not.

```yaml
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler

metadata:
  name: goserver-hpa

spec:
  # listen the deployment
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: goserver

  # in production we need at least 3
  minReplicas: 1 

  maxReplicas: 5

  # in production we can use somethin arout 70
  targetCPUUtilizationPercentage: 20
```

when the pods cpu reach the specified percentage the hpa will start to create more replicas
Normally we are going to use 70% because sometimes our application last longer to start running so if we set 90 it will not have enough time to start the others pod

```bash
kubectl apply -f k8s/hpa.yaml

# here we can verify our deployment target cpu usage
kubectl get hpa

#NAME           REFERENCE             TARGETS   MINPODS   MAXPODS   REPLICAS   AGE
#goserver-hpa   Deployment/goserver   0%/20%    1         5         1          22s
```

## Stress test with Fortio
Keep listening the memory and cpu usage
```bash
watch -n1 kubectl get hpa
```

Open another terminal to make stress test

Generate a pod with the fortio image that will be removed after execute the command
-qps 800 => 800 queries per second
-t 120s  => keep sending requests for 120 secons
-c 70    => concurrencie/parallelism of 70

```bash
kubectl run -it fortio --rm --image=fortio/fortio -- load -qps 800 -t 180s -c 70 "http://goserver-service/healthz"
```

