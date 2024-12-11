FROM openjdk:19-ea-1-jdk-oracle

WORKDIR /app

COPY ./target/worker-0.0.1-SNAPSHOT.jar /app/worker.jar

ENTRYPOINT ["java", "-jar", "worker.jar"]
