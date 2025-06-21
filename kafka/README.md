# Kafka cluster and implementations

- Purpose
   - demo low latency load balancing with elastic microservice via orchestration and automation tools 

- What does this repo do: 
Configure kafka container to be tested loacally, later shipped to k8s with a config, potentially with locally built Docker image 

- Tools
    - Orchestration: k8s, ArgoCD
    - Load Balancing: Horizontal Pod Autoscaler (HPA), Kubernetes Event-Driven Autoscaler (KEDA)
    - Kafka via Docker [4]
- The following are referenced from cited source [4]

    - Quick Run
    docker run -d --name=kafka -p 9092:9092 apache/kafka 
    - get cluster id
    docker exec -ti kafka /opt/kafka/bin/kafka-cluster.sh cluster-id --bootstrap-server :9092 

    - Create a sample topic and produce (or publish); type something after this; Ctrl+ C to terminate: 
        docker exec -ti kafka /opt/kafka/bin/kafka-console-producer.sh --bootstrap-server :9092 --topic demo
    - Confirm the messages were published: 
        docker exec -ti kafka /opt/kafka/bin/kafka-console-producer.sh --bootstrap-server :9092 --topic demo 

- Kafka build-in load balancing [1]
    - Round Robin
    - Key-based
    - Customized

- K. Hierarchy [2] - Message buffering 
    - K. cluster 
        - Brokers
            - Topics
                - Partitions 
                    - Messages

- Official doc on Load Balancing [3]
    - Locality 
        -  producer can control partition load-balancing by specifying key-based to allow data storing in one pratition 


## Reference 
 [1]https://www.confluent.io/learn/kafka-partition-strategy/?utm_source=chatgpt.com 
 [2] https://kafka.apache.org/21/javadoc/org/apache/kafka/streams/package-tree.html
 [3] https://kafka.apache.org/documentation/ 
 [4] https://docs.docker.com/guides/kafka/ 