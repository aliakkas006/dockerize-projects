# OpenJDK 11 as base image
FROM openjdk:11-jre-slim

# Define the maintainer
LABEL maintainer="Java App <101>"

# Create app directory
WORKDIR /usr/app

# Copy the built JAR file into the container
COPY target/spring-boot-docker-app.jar .

# Expose port that Spring Boot listens on
EXPOSE 8080

# Run the JAR
ENTRYPOINT ["java", "-jar", "spring-boot-docker-app.jar"]
