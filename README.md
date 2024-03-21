# gorabbitmq


Simple project for rabbitMQ written in Go

On linux run
```
$sudo chmod +x a.sh
$./a.sh
$go mod download
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
$go run producer/producer.go
$go run consumer/consumer.go
```