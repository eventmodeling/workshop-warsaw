package com.eventmodeling.survey.api;

import com.eventmodeling.survey.EventStore;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
class SurveyController {




    @RequestMapping("/")
    public String hello(){
        return "hello";
    }

    @PostMapping("/")
    public void survey(@RequestBody CompleteSurvey completeSurvey){
        EventStore eventStore = new FileEventStore(new ObjectMapper());
        SurveyApplicationService surveyApplicationService = new SurveyApplicationService(eventStore);
        surveyApplicationService.handle(completeSurvey);

    }

}
