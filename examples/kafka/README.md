### Development Kafka Deploy

[kafka-all-in-one](./kafka-all-in-one.yaml) will deploy a single kafka broker and zookeeper.
It is inspired by [this](https://dev.to/thegroo/running-kafka-on-kubernetes-for-local-development-2a54) blog post.
It includes also kafka exporter. In addition, to reach kafka from host machine, it is required to add the name of kafka service in
`/etc/hosts` file.


Kafka is exposed as a NodePort service on port 30092. This is the port of the container that represents
the kind node. In order to reach kafka from the host machine, adjust configuration of kind is neede.

See [kind config](../../kind-config.yaml) at extraPortMappings section.

> [!IMPORTANT]
> This configuration is not suitable for production environment.
> 
> It will not use persistent storage and it will not scale. 
> 
> /etc/hosts file must be updated with the following line:
> ```bash
> ## add the name of the service on the local machine, otherwise kafka clients will not be able to resolve
> ## the name of the listener
> 127.0.0.1 localhost kafka
> ```


> [!WARNING]
> This configuration is tested only on linux host machine but it should work on windows and macos too.
> 
> After deploying kafka, you can interact with the kafka cluster through the kafka cli.
> 
> Download the cli from [here](https://kafka.apache.org/downloads).
> Extract the archive and use the scripts in the bin folder to interact with kafka.
> 
> ```bash
> ## create a topic
> bin/kafka-topics.sh --create --topic test --bootstrap-server localhost:9092
> ```