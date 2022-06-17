package com.user.UserMicroservice.model;

import com.fasterxml.jackson.annotation.JsonManagedReference;
import com.user.UserMicroservice.security.Permission;

import java.time.LocalDateTime;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

import javax.persistence.*;

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
    private String email;
    private String password;
    private String phoneNumber;
    private String gender;
    private String username;
    private Boolean isPublic;
    private String biography;
    private String birthDate;
    private String workExperience;  //entitet umesto stringa ime, id
    private String education;
    private String hobby;
    private String interest;
    private String activationCode;
    private boolean activated;
    private LocalDateTime activationCodeValidity;
    private String passwordResetCode;
    private LocalDateTime passwordResetCodeValidity;
    private String loginCode;

    public Set<Permission> getPermissions() {
        return permissions;
    }

    public void setPermissions(Set<Permission> permissions) {
        this.permissions = permissions;
    }

    private LocalDateTime loginCodeValidity;

    @JsonManagedReference
    @OneToMany(mappedBy = "user", fetch = FetchType.EAGER, cascade = CascadeType.ALL)
    private Set<Permission> permissions = new HashSet<>();


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

	public LocalDateTime getActivationCodeValidity() {
		return activationCodeValidity;
	}
    private String type;


    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

    public String getRole() { return type.toString(); }


	public void setActivationCodeValidity(LocalDateTime activationCodeValidity) {
		this.activationCodeValidity = activationCodeValidity;
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

    public Boolean getPublic() {
        return isPublic;
    }

    public void setPublic(Boolean aPublic) {
        isPublic = aPublic;
    }

    public String getBiography() {
        return biography;
    }

    public void setBiography(String biography) {
        this.biography = biography;
    }

    public String getBirthDate() {
        return birthDate;
    }

    public void setBirthDate(String birthDate) {
        this.birthDate = birthDate;
    }

    public String getWorkExperience() {
        return workExperience;
    }

    public void setWorkExperience(String workExperience) {
        this.workExperience = workExperience;
    }

    public String getEducation() {
        return education;
    }

    public void setEducation(String education) {
        this.education = education;
    }

    public String getHobby() {
        return hobby;
    }

    public void setHobby(String hobby) {
        this.hobby = hobby;
    }

    public String getInterest() {
        return interest;
    }

    public void setInterest(String interest) {
        this.interest = interest;
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

	public Boolean getIsPublic() {
		return isPublic;
	}

	public void setIsPublic(Boolean isPublic) {
		this.isPublic = isPublic;
	}


    
    
}
