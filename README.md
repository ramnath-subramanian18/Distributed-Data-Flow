Question: Build a web server in Go to receive data external and route it through messaging systems

We need to start all the services for Kafka 
Run the commands after installing Kafka 
1.$ bin/zookeeper-server-start.sh config/zookeeper.properties
2.$ bin/kafka-server-start.sh config/server.properties
Install Redis and keep it running 

Use the command go run main. go and go run get_data_redis.go to execute the command 
