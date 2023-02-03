### Test workflow submission



Hi, the task has been completed, the default temporal server port has been changed to 7233 andthe main.go with the temporal client has been updated too

a few extensive files have been added, the .env file stating some universal variables, the docker-compose.yml which has all the configurations required to create and launch the container. the temporal-deployment.yaml is a kubernetes configfile stating everything required for deployment. a few security considerations were implemented. the CICD.md states my recommendations for the app's continuous integration and continuous deployment .

## How would i improve the app

To improve a Temporal app, i can focus on optimizing performance, improving scalability, enhancing reliability, adding security, and enhancing the user experience. This can include using profiling tools, scaling with container orchestration, implementing retries and circuit breakers, adding authentication and authorization, using encryption, and improving the user interface and feedback.


STEPS TO START

```bash
docker compose up
```

```bash
go run ./worker/main.go
```

Execute the helloworkflow

```bash
go run ./starter/main.go
```

Hoping i made it to the next challenge/step 👋🏻.

