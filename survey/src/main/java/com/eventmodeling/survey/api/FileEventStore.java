package com.eventmodeling.survey.api;

import com.eventmodeling.survey.DomainEvent;
import com.eventmodeling.survey.EventStore;

public class FileEventStore implements EventStore {
    @Override
    public DomainEvent load(String aggregateid) {
        throw new UnsupportedOperationException();
    }

    @Override
    public void save(DomainEvent domainEvent) {
        throw new UnsupportedOperationException();
    }
}
