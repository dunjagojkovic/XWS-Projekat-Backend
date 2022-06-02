package com.agent.model;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;
import java.util.HashSet;
import java.util.Set;

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

    @OneToMany(mappedBy = "company", fetch = FetchType.LAZY, cascade = CascadeType.ALL)
    private Set<JobOffer> positions = new HashSet<>();


}
