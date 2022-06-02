package com.agent.model;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;

@Getter
@Setter
@Entity
@Table(name = "user_table")
public class User {

    @Id
    @GeneratedValue(
            strategy = GenerationType.IDENTITY
    )
    private Long id;
    private String name;
    private String surname;
    @Column(unique = true)
    private String email;
    private String password;
    private String phoneNumber;
    private String gender;
    @Column(unique = true)
    private String username;
    private String birthDate;
    private String type;

    public String getRole() { return type.toString(); }


}
