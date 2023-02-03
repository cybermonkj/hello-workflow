package main

import (
	"log"

	"github.com/wafi/hello-workflow/helloworkflow"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.NewClient(client.Options{
		HostPort: "localhost:7233",
	})
	
	if err != nil {
		log.Fatalln("Unable to make client", err)
	}

	defer c.Close()

	w := worker.New(c, "hello-world", worker.Options{})
	w.RegisterWorkflow(helloworkflow.Workflow)
	w.RegisterActivity(helloworkflow.Activity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start workflow", err)
	}
}
