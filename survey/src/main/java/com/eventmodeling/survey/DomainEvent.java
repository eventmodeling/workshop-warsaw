package com.eventmodeling.survey;

public interface DomainEvent {
    String getAggreggateId();
    String getAggregateName();
}
