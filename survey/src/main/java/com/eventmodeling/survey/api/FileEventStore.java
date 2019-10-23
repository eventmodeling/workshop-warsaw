package com.eventmodeling.survey.api;

import com.eventmodeling.survey.DomainEvent;
import com.eventmodeling.survey.EventStore;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.UUID;

public class FileEventStore implements EventStore {

    public static final String EVENTS_DIRECTORY = "/events";



    @Override
    public DomainEvent load(String aggregateid) {
        throw new UnsupportedOperationException();
    }

    @Override
    public void save(DomainEvent domainEvent) {
        try {
            final Path path = Files.createFile(Paths.get(EVENTS_DIRECTORY, String.format("%s-%s", domainEvent.getClass().getSimpleName(),
                    UUID.randomUUID().toString())));
            final File file = path.toFile();
            ObjectMapper objectMapper = new ObjectMapper();
            objectMapper.writeValue(file, domainEvent);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
