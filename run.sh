docker build -t gin-demo:1.0 -f ./Dockerfile .

docker run -d -p 9002:8888 --name gin_demo_01 -t gin-demo:1.0
