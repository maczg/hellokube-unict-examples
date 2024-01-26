## Kafka and Zookeeper

Kafka is a distributed publish-subscribe messaging system, designed to be fast, scalable and enable message persistence. 
All actors that interact with kafka can be considered as **clients** and they exchange **messages**.
A message is a piece of data that can potentially be everything. It is up to the client to define the message format in 
transmission and reception.

Each client assumes a specific role, based on the type of interaction with kafka.
If a client produces a message, it is called **producer**. 
If a client consumes a message, it is called **consumer**.

Kafka is typically deployed as a cluster of nodes. 
Each node is called *broker* and it is responsible for storing messages and delivering them to consumers.
A broker can be deployed on a dedicated machine (physical/virtual/container-based).

Kafka is a stateful system, and it requires a distributed coordination service to manage data consistency and replication.
Zookeeper is a distributed coordination service that kafka uses to manage its cluster. 
It is also used to manage the replication of topics, partitions and consumer offsets.
Immagine zookeeper as a **consensus** mechanism that allows kafka to work properly.

Kafka organize messages in **topics**. A topic is a category or feed name to which messages are published and consumed.
A topic is divided into **partitions**. Each partition is an ordered, immutable sequence of messages that is continually appended to—a commit log.
Each message in a partition is assigned and identified by its unique **offset** (think of the offset as an index).
A topic can be replicated across multiple brokers, so **replication** is the process of synchronizing data across multiple servers.

A consumer can subscribe to one or more topics and read messages from them.

If more than one consumer must read from the same topic, they can be organized in **consumer groups**.
A consumer group is a set of consumers that cooperate to consume messages from a topic.
Tipically, each consumer in a group reads from a different partition of the same topic.
<u> It is important to note that Kafka guarantees consistency only **within** a partition, not across partitions.</u>
It means that if a consumer group has more consumers than partitions, some consumers will be idle and will not receive messages.
On the other hand, if a consumer group has less consumers than partitions, some consumers will read from more than one partition.

## Deploying kafka and zookeeper on kubernetes

**StatefulSet or Deployment?** Technically, you can use either a StatefulSet or a Deployment to run Kafka on Kubernetes.
However, a StatefulSet is the recommended approach because it provides unique network identifiers and persistent storage for each pod. 

<u>Configuration of kafka replicas can be difficult so, for development purposes, we can use a single kafka node and deploy
it as a **deployment**.</u>

FYI, in production environment we usually deploy kafka as multi-node cluster. 
In this case, due to the huge amount of configuration parameters required, usually 
can be worth to consider the use of helm charts or operators to deploy kafka on kubernetes, 
but it is out of the scope of this tutorial.

**Exposing kafka service** Kafka uses TCP protocol to communicate with clients, so 
ingress is not a suitable solution. Two options are available: NodePort and LoadBalancer (cloud only).

Using a NodePort service in kind ([ref](https://kind.sigs.k8s.io/docs/user/configuration/#nodeport-with-port-mappings)) means that kind configuration must include the following lines:

```yaml
## 
# map container port 30092 to host port 9092
nodes:
  - role: worker
    extraPortMappings:
    - containerPort: 9092
      hostPort: 9092
```

**NB** Kafka pod must be run on the node that expose the proper port.

The service can be deployed using the following command:

```bash
apiVersion: v1
kind: Service
metadata:
  name: kafka
spec:
  type: NodePort
  ports:
  - name: kafka
    nodePort: 30092
    port: 9092
  selector:
    app: kafka
```

> [!NOTE]
> To expose kafka service on your host machine, you must add the following line to your /etc/hosts file:
> ```bash
> ## add the name of the service on the local machine, otherwise kafka clients will not be able to resolve
> ## the name of the listener
> 127.0.0.1 localhost kafka
> ```

See [confluent blog post](https://www.confluent.io/blog/kafka-client-cannot-connect-to-broker-on-aws-on-docker-etc/) for
more details about listener and configuration errors.