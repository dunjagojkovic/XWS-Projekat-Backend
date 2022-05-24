package com.user.UserMicroservice.dto;

public class LoginDTO {

    private String username;
    private String password;
    private String code;
    
    

    public String getCode() {
		return code;
	}

	public void setCode(String code) {
		this.code = code;
	}

	public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getPassword(){
        return password;
    }
    public void setPassword(String password){
        this.password = password;
    }

	@Override
	public String toString() {
		return "LoginDTO [username=" + username + ", password=" + password + ", code=" + code + "]";
	}
}