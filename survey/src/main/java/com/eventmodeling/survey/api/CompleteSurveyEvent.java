package com.eventmodeling.survey.api;

import com.eventmodeling.survey.DomainEvent;
import lombok.AllArgsConstructor;
import lombok.Getter;

@AllArgsConstructor
@Getter
public class CompleteSurveyEvent implements DomainEvent {
    private final String answer;
}
