package com.agent.dto;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class JobOfferDTO {

    private Long id;
    private String position;
    private Long salary;
    private String responsibilities;
    private String requirements;
    private String benefit;
    private Long companyId;
}
