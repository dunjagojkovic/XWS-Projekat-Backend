package com.user.UserMicroservice.security;

import com.fasterxml.jackson.annotation.JsonBackReference;
import com.user.UserMicroservice.model.User;
import org.springframework.security.core.GrantedAuthority;

import javax.persistence.*;

@Entity
@Table(name = "permission")
public class Permission implements  GrantedAuthority{
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    Long id;
    String name;

    @JsonBackReference
    @ManyToOne
    private User user;


    public Permission() {
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }


    @Override
    public String getAuthority() {
        return name;
    }
}
