# Demo project using 
* Gin
* PostgreSQL
* Gorm
* Viper
* Docker
* Docker-compose
* Ansible
* Minikube
* Postman

## Create the project
```bash
go mod init go-gin-gorm-pg
```

## Import dependencies
```bash
go mod tidy
```

## Build Project
```bash
go build .
```

## Run the project
```bash
docker-compose up
```

## Test
```bash
curl http://localhost:8080/books
```

## Test with Kubernetes
* Install ansible
* Install minikube

```bash
docker build -t marcelodsales/go-gin-gorm-pg:latest . # You can change the image name
minikube cache add marcelodsales/go-gin-gorm-pg:latest
ansible-playbook config-minikube-playbook.yaml -vv
ansible-playbook deploy-go-gin-gorm-pg-kubernetes-playbook.yaml -vv
minikube tunnel --cleanup
```

## Update kubernetes deployment
```bash
docker build -t marcelodsales/go-gin-gorm-pg:0.0.1 . # increase the version
minikube cache add marcelodsales/go-gin-gorm-pg:0.0.1
minikube cache reload # Always when the image is generated again
kubectl rollout pause deployment.v1.apps/go-gin-gorm-pg -n dev
kubectl --record deployment.apps/go-gin-gorm-pg set image deployment.v1.apps/go-gin-gorm-pg go-gin-gorm-pg=marcelodsales/go-gin-gorm-pg:0.0.1 -n dev
kubectl rollout resume deployment.v1.apps/go-gin-gorm-pg -n dev
```

### Access the app behind kubernetes
```bash
kubectl get ingress -n dev
curl 127.0.0.1:80/app/books
```