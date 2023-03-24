docker build -f Dockerfile . -t httpserver:v0.1

docker run -itd httpserver:v0.1

docker ps

lsns -t net # get pid

nsenter -t {{pid}} -n ip a
