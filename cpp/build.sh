IMAGE_NAME=local/cpp
docker build -t ${IMAGE_NAME} .

CONTAINER_NAME=temp-cpp
docker create --name ${CONTAINER_NAME} ${IMAGE_NAME}:latest
docker cp ${CONTAINER_NAME}:/app/build/cpp.zip .
docker rm ${CONTAINER_NAME}
