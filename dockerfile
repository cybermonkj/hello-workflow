FROM temporalio/server
COPY /main.go server
WORKDIR /
ENTRYPOINT ["temporal", "worker", "--address", "0.0.0.0:7933", "--tasklist", "sample-tasklist", "--namespace", "sample-namespace"]
