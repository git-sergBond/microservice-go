Окружение:
podman
minikube

Собрать образ:
podman  build . -t go-hello

Запустить контейнер:
podman run -p 8080:8080 go-hello

Миникуб:
minikube start --vm-driver=virtualbox
minikube stop
https://kubernetes.io/ru/docs/tasks/tools/install-minikube/
https://kubernetes.io/ru/docs/setup/learning-environment/minikube/

КубКонтрол:
kubectl config current-context