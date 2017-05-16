To run the simulator:

    cd exec
    go run main.go -c config.xml

    Change the following values in config.xml:
      servers: pass the zookeepers host and port(for "Ensemble")
      numberofclients: number of clients to be used in simulation.
      messages: number of messages to be received by clients
      debug: set to true, to get debugging message, which will give idea about how the simulation is behaving,
             if set to false, will print only if lock acquired by a client, else nothing
