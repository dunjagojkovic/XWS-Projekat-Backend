package com.agent.dto;

import com.agent.model.Company;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class CompanyOfferDTO {
	
	private Long id;
    private String position;
    private Long salary;
    private String responsibilities;
    private String requirements;
    private String benefit;
    private Company company;

}
