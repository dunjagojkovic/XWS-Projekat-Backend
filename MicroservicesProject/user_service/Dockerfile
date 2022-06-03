FROM openjdk:8-jdk-alpine 
EXPOSE 8000
WORKDIR /user/src
COPY /common ../common
COPY /*.jar app.jar
ENTRYPOINT ["java","-jar","app.jar"]



