package com.eventmodeling.survey;

import java.util.List;

public interface EventStore {
    List<DomainEvent> load(String aggregateid);
    void save(DomainEvent domainEvent);
}
