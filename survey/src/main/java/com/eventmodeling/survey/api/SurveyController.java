package com.eventmodeling.survey.api;

import com.eventmodeling.survey.EventStore;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
class SurveyController {

    @GetMapping("/")
    public String survey(){
        return "<!DOCTYPE html>\n" +
                "<html>\n" +
                "\n" +
                "<style>\n" +
                "  .login-page {\n" +
                "    width: 360px;\n" +
                "    padding: 8% 0 0;\n" +
                "    margin: auto;\n" +
                "  }\n" +
                "§\n" +
                "  .form {\n" +
                "    position: relative;\n" +
                "    z-index: 1;\n" +
                "    background: #FFFFFF;\n" +
                "    max-width: 360px;\n" +
                "    margin: 0 auto 100px;\n" +
                "    padding: 45px;\n" +
                "    text-align: center;\n" +
                "    box-shadow: 0 0 20px 0 rgba(0, 0, 0, 0.2), 0 5px 5px 0 rgba(0, 0, 0, 0.24);\n" +
                "  }\n" +
                "\n" +
                "  .form input {\n" +
                "    outline: 0;\n" +
                "    background: #f2f2f2;\n" +
                "    width: 100%;\n" +
                "    border: 0;\n" +
                "    margin: 0 0 15px;\n" +
                "    padding: 15px;\n" +
                "    box-sizing: border-box;\n" +
                "    font-size: 14px;\n" +
                "  }\n" +
                "\n" +
                "  .form button {\n" +
                "    text-transform: uppercase;\n" +
                "    outline: 0;\n" +
                "    background: #4CAF50;\n" +
                "    width: 100%;\n" +
                "    border: 0;\n" +
                "    padding: 15px;\n" +
                "    color: #FFFFFF;\n" +
                "    font-size: 14px;\n" +
                "    -webkit-transition: all 0.3 ease;\n" +
                "    transition: all 0.3 ease;\n" +
                "    cursor: pointer;\n" +
                "  }\n" +
                "  .form button:hover,.form button:active,.form button:focus {\n" +
                "    background: #43A047;\n" +
                "  }\n" +
                "  .form .message {\n" +
                "    margin: 15px 0 0;\n" +
                "    color: #b3b3b3;\n" +
                "    font-size: 12px;\n" +
                "  }\n" +
                "  .form .message a {\n" +
                "    color: #4CAF50;\n" +
                "    text-decoration: none;\n" +
                "  }\n" +
                "  .form .register-form {\n" +
                "    display: none;\n" +
                "  }\n" +
                "  .container {\n" +
                "    position: relative;\n" +
                "    z-index: 1;\n" +
                "    max-width: 300px;\n" +
                "    margin: 0 auto;\n" +
                "  }\n" +
                "  .container:before, .container:after {\n" +
                "    content: \"\";\n" +
                "    display: block;\n" +
                "    clear: both;\n" +
                "  }\n" +
                "  .container .info {\n" +
                "    margin: 50px auto;\n" +
                "    text-align: center;\n" +
                "  }\n" +
                "  .container .info h1 {\n" +
                "    margin: 0 0 15px;\n" +
                "    padding: 0;\n" +
                "    font-size: 36px;\n" +
                "    font-weight: 300;\n" +
                "    color: #1a1a1a;\n" +
                "  }\n" +
                "  .container .info span {\n" +
                "    color: #4d4d4d;\n" +
                "    font-size: 12px;\n" +
                "  }\n" +
                "  .container .info span a {\n" +
                "    color: #000000;\n" +
                "    text-decoration: none;\n" +
                "  }\n" +
                "  .container .info span .fa {\n" +
                "    color: #EF3B3A;\n" +
                "  }\n" +
                "\n" +
                "  body {\n" +
                "  }\n" +
                "\n" +
                "</style>\n" +
                "\n" +
                "<body>\n" +
                "\n" +
                "  <div class=\"survey-page\">\n" +
                "    <div class=\"form\">\n" +
                "      <form class=\"survey-form action=\"/survey\" method=\"post\">\n" +
                "        <p>Do you like apple?</p>\n" +
                "        <input name=\"answer\" type=\"text\" placeholder=\"name\" />\n" +
                "        <button>Complete</button>\n" +
                "      </form>\n" +
                "    </div>\n" +
                "  </div>\n" +
                "\n" +
                "</body>\n" +
                "\n" +
                "</html>\n";
    }

    @PostMapping("/survey")
    public void survey(@ModelAttribute CompleteSurvey completeSurvey){
        EventStore eventStore = new FileEventStore(new ObjectMapper());
        SurveyApplicationService surveyApplicationService = new SurveyApplicationService(eventStore);
        surveyApplicationService.handle(completeSurvey);

    }



}
