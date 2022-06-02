package com.agent.dto;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class CompanyDTO {

    private Long id;
    private String name;
    private String state;
    private String city;
    private String address;
    private String contact;
    private String email;
    private String description;
    private Long ownerId;
    private String status;



}
