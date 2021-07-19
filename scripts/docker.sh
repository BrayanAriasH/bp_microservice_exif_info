docker build -t bp-create-images .
docker run -p80:1503 -it bp-create-images