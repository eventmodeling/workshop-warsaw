package com.eventmodeling.survey;

public interface EventStore {
    DomainEvent load(String aggregateid);
    void save(DomainEvent domainEvent);
}
