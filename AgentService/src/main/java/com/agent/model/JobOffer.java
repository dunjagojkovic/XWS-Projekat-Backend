package com.agent.model;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;

@Getter
@Setter
@Entity
@Table(name = "jobs_table")
public class JobOffer {

    @Id
    @GeneratedValue(
            strategy = GenerationType.IDENTITY
    )
    private Long id;
    private String position;
    private Long salary;
    private String responsibilities;
    private String requirements;
    private String benefit;

    @ManyToOne
    private Company company;



}
