package com.agent.security;

import org.springframework.security.core.GrantedAuthority;

public class Authority implements GrantedAuthority {
    String name;

    public String getName(){
        return name;
    }
    public void setName(String name){
        this.name = name;
    }

    @Override
    public String getAuthority() { return name; }
}