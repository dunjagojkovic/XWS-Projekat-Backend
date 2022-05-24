package com.user.UserMicroservice.service;

import org.springframework.stereotype.Service;

import com.user.UserMicroservice.model.User;

import net.bytebuddy.utility.RandomString;

@Service
public class CodeService {
	public String generateActivationCodeForUSer(User user) {
		return RandomString.make(64);
	}
	
	public String generatePasswordResetCode(User user) {
		return RandomString.make(64);
	}

	public String generateLoginCode(User user) {
		// TODO Auto-generated method stub
		return RandomString.make(8);
	}
	

}
