# Measures

## Commandments of Cost Optimization

Sometimes we have a tendency to just spin up a ec2 and install our application
we don't check how much memory it's using in average, how much cpu it's using in average
and a lot of times there are a lot of wastage in there
and we're just paying for extra stuff that you don't need.

### Right Sizing
    - utilize pod requests, limits, resources quotas
    - use open source, third party tools to tune pod requests, limits

tools that tell you how much cpu/memory is being wasted, how much money you can save:
Open Source: 
    - RSG - Right Size Guide
    - Kubecost - has a limit on the open source
    - Kubernetes Resource Report - shows dollar amount you can save
    - Goldilocks

Third Party:
    - Kubecost
    - New Relic
    - CloudHealth by vmware


###  Auto Scaling
Once we know how much cpu and memory our pod will need we can scale up
and we only pay more when the traffic increases.
    - Once pods are optimized, enable auto scaling
    - utilize HPA, Cluster Autoscaling, Proportional Autoscaling

### Down Scaling
    - Terminate pods unnecessary during nights, weekends
    - Utilize DevOps

### EC2 Purchase Options
    - Use RI Instances, Spot, Savings Plan