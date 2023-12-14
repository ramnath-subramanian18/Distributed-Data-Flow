Question: Build a webserver in Go to receive data external and route through messaging systems

We need to start the all the services for kafka 
Run the commands after installing kafka 
1.$ bin/zookeeper-server-start.sh config/zookeeper.properties
2.$ bin/kafka-server-start.sh config/server.properties
Install Redis and keep it running 

Use the command go run main.go and go run get_data_redis.go to execute the command 
