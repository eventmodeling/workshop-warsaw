package com.eventmodeling.survey.api;

import com.eventmodeling.survey.EventStore;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.core.io.InputStreamResource;
import org.springframework.core.io.Resource;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.RestController;

import java.io.IOException;

@RestController
class SurveyController {

    @Value("${:classpath:survey.html}")
    private Resource index;

    @GetMapping("/")
    @ResponseBody
    public ResponseEntity actions() throws IOException {
        return ResponseEntity.ok(new InputStreamResource(index.getInputStream()));
    }

    @PostMapping("/")
    public void survey(@ModelAttribute CompleteSurvey completeSurvey){
        EventStore eventStore = new FileEventStore(new ObjectMapper());
        SurveyApplicationService surveyApplicationService = new SurveyApplicationService(eventStore);
        surveyApplicationService.handle(completeSurvey);

    }



}
