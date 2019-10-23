package com.eventmodeling.survey.api;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import org.springframework.web.bind.annotation.GetMapping;

@NoArgsConstructor
@Getter
@Setter
public class CompleteSurvey {
    String answer;
}
