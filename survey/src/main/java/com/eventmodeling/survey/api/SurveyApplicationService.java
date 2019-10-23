package com.eventmodeling.survey.api;

import com.eventmodeling.survey.EventStore;

public class SurveyApplicationService {

    private EventStore eventStore;

    public SurveyApplicationService(EventStore eventStore) {

        this.eventStore = eventStore;
    }

    public void handle(CompleteSurvey completeSurvey) {
        eventStore.save(new CompleteSurveyEvent("dupadupa"));
    }
}
