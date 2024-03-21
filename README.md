# gorabbitmq


Simple project for rabbitMQ written in Go

On linux run
```
$sudo chmod +x a.sh
$./a.sh
$go get -u -v -f all
```
This will install rabbitMq on you Linux machine

To start RabbitMQ run

```
$sudo systemctl status rabbitmq-server
```
For starting the UI part of rabbitMQ

```
$sudo rabbitmq-plugins enable rabbitmq_management
```
Open the browser and go to

```
http://localhost:15672
```
Username: guest
Password: guest

Now run both command in different terminal
```
$go run rabbitmq/producer/producer.go
$go run rabbitmq/consumer/consumer.go
```