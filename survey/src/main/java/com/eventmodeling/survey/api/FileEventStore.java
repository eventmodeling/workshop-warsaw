package com.eventmodeling.survey.api;

import com.eventmodeling.survey.DomainEvent;
import com.eventmodeling.survey.EventStore;
import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.AllArgsConstructor;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.Collections;
import java.util.List;
import java.util.UUID;
import java.util.stream.Collectors;
import java.util.stream.Stream;

@AllArgsConstructor
public class FileEventStore implements EventStore {

    public static final String EVENTS_DIRECTORY = "/events";

    private ObjectMapper objectMapper;



    @Override
    public List<DomainEvent> load(String aggregateid) {
        try (Stream<Path> paths = Files.walk(Paths.get("/home/you/Desktop"))) {
            return paths
                    .filter(Files::isRegularFile)
                    .map(file -> mapEvent(file))
                    .collect(Collectors.toList());
        } catch (IOException e) {
            e.printStackTrace();
        }
        return Collections.EMPTY_LIST;
    }

    private DomainEvent mapEvent(Path file) {
        try {
            return objectMapper.readValue(file.toFile(), resolveClassName(file));
        } catch (IOException e) {
            e.printStackTrace();
        }
        return null;
    }

    private Class<? extends DomainEvent> resolveClassName(Path file) {
        return CompleteSurveyEvent.class;
    }

    @Override
    public void save(DomainEvent domainEvent) {
        try {
            final Path path = Files.createFile(Paths.get(EVENTS_DIRECTORY, String.format("%s-%s", domainEvent.getClass().getSimpleName(),
                    UUID.randomUUID().toString())));
            final File file = path.toFile();
            objectMapper.writeValue(file, domainEvent);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
