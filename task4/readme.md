# Task - Microservice

Assume we want to expose the API from `task_1` as a microservice in production with more complete endpoints.
What would it take to get it to production? Explain your process.

We need to dockerize the project with creating settings for docker-compose.yml This api is already dockerized.

We can add an endpoint to api for health check of pods.

I added k8s.yaml file to create k8s cluster. 

Also we can use a monitoring tool like logdna to understand problems. There is no guarantee that our pods will be live so the problems can be occur from our cluster and also from our code. We need to analyze it.

## Bonus points
Create a microservice demo based on the greeting-API on minikube.