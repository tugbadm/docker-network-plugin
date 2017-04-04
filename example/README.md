1- Create a bridge, __mydriver__ plugin uses a bridge.
```bash
brctl addbr br1
ifconfig br1 10.0.1.1
ifconfig br1 up
```

2- run main.go
```bash
go run main.go
```
3- Create the json file on /etc/docker/plugins
mynetwork.json
```json
{
    "Name": "mynetwork",
    "Addr": "tcp://127.0.0.1:8010"
}
```

4- To create network
```bash
docker network create -d mynetwork mynet -o bridge=br1 --ip-range=10.0.1.1/24 --subnet=10.0.1.1/24
```
You might not set ip-range and subnet, but if you do you will have an isolated network.

5- To run a container
```bash
docker run -ti -network=mynet ubuntu:14.04 bash
```

more details checkout: https://medium.com/@tugba/docker-networking-with-plugin-8fc3ce97444

