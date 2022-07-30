
## Ingress
single entry point for the application.
Works as a target group where we can define the rules, if the user is acessing XPTO /users url we are going to forward the request to users service, if is accessing /products ...

It makes easy to manage the services IP's.

It's similar to a API Gateway or Reverse Proxy.


https://kubernetes.github.io/ingress-nginx/