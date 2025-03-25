# this is dependant on the language the app is build on, and to know which version of the base image to use
# go to the docker site and find the appropriate image, use alpine for lightweight image
# Lec 22 notes, the problem with running this code before the staging code is that the image size is large
# and the output image is 50 times the alpine image and it is due to the fact that, the large image 
# contains golang and all the packages required by the image when all that is required for the container 
# is the output binary file that will be executed. 

# So to solve the problem, we want to produce an iamge that contains only the output binary file
# to do so,

# Specify the build stage, where we build only the binary file
# build stage is specified using the AS keyword
FROM golang:1.23.4-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
# when we spin the container during the build phase in docker compose
# the postgres instance will be without a database hence to run our mgrate command we do as follows
RUN apk --no-cache add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz


# Run stage to ensure that the output image size will be small, we will use the alpine image
FROM alpine:3.21
WORKDIR /app
# this copy command specifies that we will copy the file from the build stage to the working directory mian folder
COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
# copy the env file to the docker container to be used
# however this is not best practice and we will learn how to replace this with the actual production config
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration

EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]