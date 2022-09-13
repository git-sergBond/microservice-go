Собрать образ:
docker  build . -t time-service:0.1.0

Запустить контейнер:
docker run -p 8080:8080 time-service:0.1.0

КубКонтрол команды:
kubectl get nodes
kubectl config current-context