package com.agent.model;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;
import java.time.LocalDateTime;

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

    public String getSurname() {
        return surname;
    }

    public void setSurname(String surname) {
        this.surname = surname;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public String getPhoneNumber() {
        return phoneNumber;
    }

    public void setPhoneNumber(String phoneNumber) {
        this.phoneNumber = phoneNumber;
    }

    public String getGender() {
        return gender;
    }

    public void setGender(String gender) {
        this.gender = gender;
    }

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getBirthDate() {
        return birthDate;
    }

    public void setBirthDate(String birthDate) {
        this.birthDate = birthDate;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

    public String getActivationCode() {
        return activationCode;
    }

    public void setActivationCode(String activationCode) {
        this.activationCode = activationCode;
    }

    public boolean isActivated() {
        return activated;
    }

    public void setActivated(boolean activated) {
        this.activated = activated;
    }

    public LocalDateTime getActivationCodeValidity() {
        return activationCodeValidity;
    }

    public void setActivationCodeValidity(LocalDateTime activationCodeValidity) {
        this.activationCodeValidity = activationCodeValidity;
    }

    public String getPasswordResetCode() {
        return passwordResetCode;
    }

    public void setPasswordResetCode(String passwordResetCode) {
        this.passwordResetCode = passwordResetCode;
    }

    public LocalDateTime getPasswordResetCodeValidity() {
        return passwordResetCodeValidity;
    }

    public void setPasswordResetCodeValidity(LocalDateTime passwordResetCodeValidity) {
        this.passwordResetCodeValidity = passwordResetCodeValidity;
    }

    public String getLoginCode() {
        return loginCode;
    }

    public void setLoginCode(String loginCode) {
        this.loginCode = loginCode;
    }

    public LocalDateTime getLoginCodeValidity() {
        return loginCodeValidity;
    }

    public void setLoginCodeValidity(LocalDateTime loginCodeValidity) {
        this.loginCodeValidity = loginCodeValidity;
    }
}
