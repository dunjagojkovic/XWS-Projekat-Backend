package com.agent.model;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;

@Getter
@Setter
@Entity
@Table(name = "company_table")
public class Company {
    @Id
    @GeneratedValue(
            strategy = GenerationType.IDENTITY
    )
    private Long id;
    @Column(unique = true)
    private String name;
    private String state;
    private String city;
    private String address;
    private String contact;
    private String email;
    private String description;

    @ManyToOne
    private User owner;
    private String status;


}
