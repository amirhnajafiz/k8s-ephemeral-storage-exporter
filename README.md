# K8S Ephemeral Storage Exporter

Even though ephemeral storage is a resource you can set in Kubernetes manifests, there aren’t many tools (exporters) that show how much of it is actually being used. When this storage fills up, it often causes problems like running out of space, which makes the Kubelet remove (evict) pods. A simple exporter that reports how much ephemeral storage each pod is using would help with monitoring and setting up alerts before issues happen.
