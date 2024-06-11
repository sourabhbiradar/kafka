# kafka

Event driven programming ; Go + Kafka + Docker

Apache Kafka is a `Distributed Event Streaming System`.

Producer - Consumer Architecture

`Architecture Components` :

- `Producer` : Producers publish data to topics. Producers can specify which partition to send messages to, or let Kafka decide based on the message key.

- `Consumer` : Consumers subscribe to topics and process messages. They can belong to a consumer group.

- `Broker` : Its a server , responsible for `storing data / message & serving clients` (Receives messages from Producer & sends them to consumer). Each broker has a `Unique ID`.

- `Cluster` : Group of Brokers.

- `Topic` ** : A `unique name` given to a data stream or message stream.

`Producer/Producers >>msg>> Broker::Topic <<subscribe ; msg>> Consumer`

- `Partition` : Incase of large data ,Topic is further divided into Partitions. Partitions are distributed on different machines. Number of partitions are decided during topic creation.

- `Offset` : `Sequence number` assigned to each message in each partition of topic is called Offset. For a given topic, different partitions have different offsets & are always local to topic partitions.

`unique identity for any message = topic name + partition number + offset number`

- `ZooKeeper` : Manages and coordinating brokers, and maintaining track about topics, partitions, and configurations. Also determines leader elections.


- `Single Broker Cluster` : The Kafka cluster having only one broker is called Single Broker Cluster.
- `Multi-Broker Cluster` : The Kafka cluster having two or more brokers is called Multi-Broker Cluster.
- `Consumer Group` : Group of consumers that share the workload.
Multiple consumer groups subscribe to the same or different topics.

NOTE : Two or more consumers of same consumer group do not receive the common message. They always receive a different message because the offset pointer moves to the next number once the message is consumed by any of the consumers in that consumer group.

`cmd` :

go get -u github.com/confluentinc/confluent-kafka-go/kafka

brew install kafka

brew services start kafka

in working dir :

docker-compose up -d

go run producer.go

ggo run consumer.go
