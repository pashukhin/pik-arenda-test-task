
Implementation of test task for PIK-Arenda.
See https://docs.google.com/document/d/12s8dlqyHxM2JJ98Kkovfwzp6TOS3GN2w6Ts08kN2ev4/edit?usp=sharing for details.

### How to run
`docker-compose up ` for run containers.
Edit `docker-compose.yml` to change host, port and so on if you need it.
By defaults REST API server runs at localhost:8080.

### API examples
- `curl localhost:8080/worker` list workers
- `curl -X POST -H 'Content-Type: application/json' -d '{"name":"Misha"}' localhost:8080/worker` create worker
- `curl localhost:8080/worker/1` get worker
- `curl -X PUT -H 'Content-Type: application/json' -d '{"name":"Migel First"}' localhost:8080/worker/1` update worker
- `curl -X DELETE localhost:8080/worker/1` delete worker
- `curl localhost:8080/schedule` list schedules. 
Optional query params, `from` and `to`, are json datetime string, for example, `2019-10-30T09:00:00.0Z`
- `curl -X POST -H 'Content-Type: application/json' -d '{"worker_id":1,"start":"2019-10-30T09:00:00.0Z","end":"2019-10-30T18:00:00.0Z"}' localhost:8080/schedule` create schedule
-`curl localhost:8080/schedule/1` get schedule
- `curl -X PUT -H 'Content-Type: application/json' -d '{"start":"2019-10-30T10:00:00.0Z"}' localhost:8080/schedule/1` update schedule
- `curl -X DELETE localhost:8080/schedule/1` delete schedule
- `curl localhost:8080/task` list tasks. 
Optional query params, `from` and `to`, are json datetime string, for example, `2019-10-30T09:00:00.0Z`
- `curl -X POST -H 'Content-Type: application/json' -d '{start":"2019-10-30T09:00:00.0Z","end":"2019-10-30T18:00:00.0Z"}' localhost:8080/task` create task
- `curl localhost:8080/task/1` get task
- `curl -X PUT -H 'Content-Type: application/json' -d '{"start":"2019-10-30T10:00:00.0Z"}' localhost:8080/task/1` update task
- `curl -X PUT -H 'Content-Type: application/json' -d '{"cancelled":true}' localhost:8080/task/1` cancel task
- `curl -X DELETE localhost:8080/task/1` delete task
- `curl localhost:8080/free_schedule` list free schedule function intervals.
Optional query params, `from` and `to`, are json datetime string, for example, `2019-10-30T09:00:00.0Z`.
Result looks like
`[{"id":2,"start":"2019-10-30T09:00:00Z","end":"2019-10-30T15:00:00Z","value":1},{"id":1,"start":"2019-10-30T15:00:00Z","end":"2019-10-30T18:00:00Z","value":2}]`
- `curl localhost:8080/free_slot` list free slots. 
Optional query params, `from` and `to`, are json datetime string, for example, `2019-10-30T09:00:00.0Z`.
Result looks like `[{"start":"2019-10-30T09:00:00Z","end":"2019-10-30T18:00:00Z"}]`.

#####Useful commands
- stop all containers `docker container stop $(docker container ls -aq)`
- remove all containers `docker container rm $(docker container ls -aq)`
