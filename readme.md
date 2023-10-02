GOLANG AND KAFKA API

1. `docker compose up`
2. `docker compose exec kafka bash`
3. `kafka-console-producer --bootstrap-server=localhost:9092 --topic=courses` (open queu console)
4. `docker compose exec app bash` (to see messages provided by kafka)
