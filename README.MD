Окружение:
podman
minikube

Собрать образ:
docker  build . -t time-service:0.1.0

Запустить контейнер:
docker run -p 8080:8080 time-service:0.1.0

Миникуб:
minikube start --vm-driver=virtualbox
minikube stop
https://kubernetes.io/ru/docs/tasks/tools/install-minikube/
https://kubernetes.io/ru/docs/setup/learning-environment/minikube/

КубКонтрол:
kubectl config current-context