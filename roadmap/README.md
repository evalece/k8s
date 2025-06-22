# Dev Roadmp and Plan

- High Level Planning
    - Low latency streaming 
        - Elastic consumer microservice
        - Partition 
            - Kafka build-ins:
                - by keyword to enable locality
                - Round Robin 
                - customized methods
    - High throughput?
        - (Less seem) Kafka broker ++ to handle throught put

- Implemntation 
    - Local 
        - Microservice
            - Kafka - downstream: kafka producer, upstream: kafka consumer
            - Kafka clients - downstream: Kafka, upstream: data sink
                - Consumer 
                - Producer
            - Data sink - downstream: kafka consumer, upstream (asynch) : end user

    - Cloud 