# EMQX

In this example, first we setup a ```EMQX``` cluster with 3 nodes, using ```docker-compose```. After that we are going to use ```Nginx``` as a
load balancer to manage the input traffic that is being transmitted to cluster nodes. Finally, we are using ```mqtt.go``` library to test the
cluster by publishing events and subscribing over a topic.

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
