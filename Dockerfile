FROM golang:alpine

# Name of maintiner
MAINTAINER DAVINDER PAL "dpsangwal@gmail.com"

# Creating Application folder in container
RUN mkdir -p /app

# Setting Home Directory for Application
WORKDIR /app

# Installing Glide Package Manager
#RUN apk update \
#    && apk install glide

# Copying Glide.yml
#ADD glide.yml /app

# Installing Golang Modules
#RUN glide install

# Exporting Ports for Application
EXPOSE 80

# Creating Volume for Persistent Data
#VOLUME ["/app-data"]

# Adding Source Code to Container
ADD . .

RUN go build
# Exporting Environment Varaiables
#ENV APPLICATION_ENVIRONMENT DEVELOPMENT

# Running Application
CMD ["/app/app]