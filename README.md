# EMQX

This is an example for using distributed ```EMQX```.
In this example, we are building our cluster using **docker**. At first, we setup a ```EMQX``` cluster with 3 nodes. When the whole cluster is up and running, we are going to use ```Nginx``` as a
load balancer to manage the input traffic that is being transmitted to cluster nodes.

In order to test the cluster validation, we setup a BlackBox exporter. We are using ```mqtt.go``` library to test the cluster by publishing events and subscribing over a topic.

## Setup

Use the following docker-compose command in order to bring up the whole example.

```sh
docker-compose -f emqx/docker-compose.yml up -d
```

After that you should have the three following containers:

- emqx-cluster-1-container, emqx-cluster-2-container, emqx-cluster-3-container
- nginx-container
- emqx-client-container (blackbox)

> NOTE: In order to stop the containers use the following command:
> ```docker-compose -f emqx/docker-compose.yml down```

## Logs

If you inspect the client-container logs, you can see that the whole cluster is responding:

```shell
emqx[client]   enter Publish
emqx[client]   sending publish message, topic: rides
emqx[net]      obound msg to write 0
emqx[net]      obound wrote msg, id: 0
emqx[net]      outgoing waiting for an outbound message
2024/03/11 08:46:44 topic:rides
new-ride
emqx[net]      startIncoming Received Message
emqx[net]      startIncomingComms: got msg on ibound
emqx[net]      startIncomingComms: received publish, msgId: 0
emqx[net]      logic waiting for msg on ibound
emqx[client]   enter Publish
emqx[client]   sending publish message, topic: rides
emqx[net]      obound msg to write 0
emqx[net]      obound wrote msg, id: 0
emqx[net]      outgoing waiting for an outbound message
2024/03/11 08:46:44 topic:rides
new-ride
```
