# Build configuration instruction
docker build -t configuration \
--build-arg APP="server_configuration" \
--progress=plain --no-cache .
# Docker run instruction
docker run -tid --name configuration --net config -p 8080:8080 configuration

# Build authentication instruction
docker build -t authentication \
--build-arg APP="authentication_service" \
--progress=plain --no-cache .
# Docker run instruction
docker run -tid --name authentication --net config -p 8081:8081 authentication

# Build authorization instruction
docker build -t authorization \
--build-arg APP="authorization_service" \
--progress=plain --no-cache .
# Docker run instruction
docker run -tid --name authorization --net config -p 8082:8082 authorization

# Build jwt instruction
docker build -t jwt \
--build-arg APP="jwt_service" \
--progress=plain --no-cache .
# Docker run instruction
docker run -tid --name jwt --net config -p 8083:8083 jwt