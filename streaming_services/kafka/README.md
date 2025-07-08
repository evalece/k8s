# Kafka cluster and implementations

- Purpose
   - demo low latency load balancing with elastic microservice via orchestration and automation tools 

- What does this repo do: 
Configure kafka container to be tested loacally, later shipped to k8s with a config, potentially with locally built Docker image 

- Quick Command 
```bash 
docker rm -f kafka
```
```bash 
docker compose -f compose.k.yml up 
```
- Tools
    - Orchestration: k8s, ArgoCD
    - Load Balancing: Horizontal Pod Autoscaler (HPA), Kubernetes Event-Driven Autoscaler (KEDA)
    - Kafka via Docker [4]
- The following are referenced from cited source [4]

    - Quick Run
    ```bash 
    docker run -d --name=kafka -p 9092:9092 apache/kafka 
    ```
    - get cluster id
    ```bash 
    docker exec -ti kafka /opt/kafka/bin/kafka-cluster.sh cluster-id --bootstrap-server :9092 
    ```
    - Create a sample topic and produce (or publish); type something after this; Ctrl+ C to terminate: 
    ```bash 
        docker exec -ti kafka /opt/kafka/bin/kafka-console-producer.sh --bootstrap-server :9092 --topic demo
    ```
    - Confirm the messages were published: 
    ```bash 
        docker exec -ti kafka /opt/kafka/bin/kafka-console-producer.sh --bootstrap-server :9092 --topic demo 
    ```
- Kafka build-in load balancing [1]
    - Round Robin
    - Key-based
    - Customized

- K. Hierarchy [2] - Message buffering 
    - K. cluster 
        - Brokers (Nodes, behaves a bit like Raft and ZK nodes, check Kraft ), can be +/- at runtime
            - Topics
                - Partitions- (File dir of nodes, cond. addressing. Offset applies here) defines how data is distributed or replicated over nodes, ex. pratition on keyword, topics 
                    - Messages (config. to have kaf. controlled offset or client offset; clients gets offset from the same partition)

- Official doc on Load Balancing [3]
    - Locality 
        -  producer can control partition load-balancing by specifying key-based to allow data storing in one pratition 

- K. with clients
    - K. removes messages after a set deadline, new message is appended, consumer keeps an offset value 
        - consumer record their offset
        - However, can be centralized and controlled by kafka

- Summary
    - Latency bottleneck potentially at: 
        - downstream consumer (at kafka -> consumer phase)
        - Kafka listening port (at producer-> kafka phase)
    - Our repo will try: 
        - Using Helm, Helmfile or AugoCD to scale down and up consumer microservice 

    - 
## Reference 
 [1]https://www.confluent.io/learn/kafka-partition-strategy/?utm_source=chatgpt.com 
 [2] https://kafka.apache.org/21/javadoc/org/apache/kafka/streams/package-tree.html
 [3] https://kafka.apache.org/documentation/ 
 [4] https://docs.docker.com/guides/kafka/ 