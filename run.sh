
 kill -9 $(lsof -ti:8080)

#Remove all containers:
docker rm -f $(docker ps -aq --filter "name=api-container")
docker rm -f $(docker ps -aq --filter "name=sql-container")

#Remove all images:
docker rmi -f $(docker images -aq --filter "reference=api")
docker rmi -f $(docker images -aq --filter "reference=mysql")

#Run all container:
 docker-compose up --build