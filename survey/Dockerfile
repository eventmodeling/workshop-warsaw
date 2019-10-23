FROM adoptopenjdk/openjdk11
VOLUME /tmp

COPY mvnw .
COPY .mvn .mvn
COPY pom.xml .
COPY src src

RUN ./mvnw install -DskipTests
RUN mkdir -p target/dependency && (cd target/dependency; jar -xf ../*.jar)

ENTRYPOINT ["java","-Djava.security.egd=file:/dev/./urandom","-jar","target/survey-0.0.1-SNAPSHOT.jar"]
