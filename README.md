# EMQX

In this example, first we setup a ```EMQX``` cluster with 3 nodes, using ```docker-compose```. After that we are going to use ```Nginx``` as a
load balancer to manage the input traffic that is being transmitted to cluster nodes. Finally, we are using ```mqtt.go``` library to test the
cluster by publishing events and subscribing over a topic.
