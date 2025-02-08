# Golang Microservice integration with Prometheus Metrics

### Golang microservice expose metrics via /metrics API
![img.png](img.png)

### Create some APIs(/ping) to let Prometheus collect
![img_4.png](img_4.png)


### Configure Prometheus to link to Golang microservice
![img_5.png](img_5.png)


#### Then run Prometheus with this config

`
docker run -d --name prometheus -p 9090:9090 -v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus
`

#### Prometheus run in :9090

#### It can see the Golang service 
![img_1.png](img_1.png)


![img_2.png](img_2.png)


#### We can show the statistic using Graph in Prometheus
![img_3.png](img_3.png)