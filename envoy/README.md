# what2do-envoy
The Envoy proxy for the Sports Admin service

The backend listens on port 12345, and the client should send to port 10000

`docker build -t what2do-envoy:latest .`

`docker run -d --name what2do-envoy -p 9901:9901 -p 10000:10000 what2do-envoy:latest`
