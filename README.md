`docker-compose up ` for run containers

`curl localhost:8080/worker` list workers

`curl -X POST -H 'Content-Type: application/json' -d '{"name":"Misha"}' localhost:8080/worker` create worker

`curl localhost:8080/worker/1` get worker

`curl -X PUT -H 'Content-Type: application/json' -d '{"name":"Migel First"}' localhost:8080/worker/1` update worker

`curl -X DELETE localhost:8080/worker/1` delete worker

`curl localhost:8080/schedule` list schedules

`curl -X POST -H 'Content-Type: application/json' -d '{"worker_id":"1","start":"2019-10-29T09:00:00.0Z","end":"2019-10-29T18:00:00.0Z"}' localhost:8080/schedule` create schedule

stop all containers `docker container stop $(docker container ls -aq)`

remove all containers `docker container rm $(docker container ls -aq)`
