package com.eventmodeling.survey.api;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
class SurveyController {


    @RequestMapping("/")
    public String hello(){
        return "hello";
    }

}