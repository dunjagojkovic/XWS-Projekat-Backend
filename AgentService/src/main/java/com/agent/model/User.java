package com.agent.model;

import com.agent.security.Permission;
import com.fasterxml.jackson.annotation.JsonManagedReference;
import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;
import java.time.LocalDateTime;
import java.util.HashSet;
import java.util.Set;

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
    @Column(unique = true)
    private String activationCode;
    private boolean activated;
    private LocalDateTime activationCodeValidity;
    @Column(unique = true)
    private String passwordResetCode;
    private LocalDateTime passwordResetCodeValidity;
    @Column(unique = true)
    private String loginCode;
    private LocalDateTime loginCodeValidity;
    public String getRole() { return type.toString(); }

    @JsonManagedReference
    @OneToMany(mappedBy = "user", fetch = FetchType.EAGER, cascade = CascadeType.ALL)
    private Set<Permission> permissions = new HashSet<>();

}
