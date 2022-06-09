package com.agent.service;

import com.agent.config.SecurityUtils;
import com.agent.dto.RegistrationDTO;
import com.agent.model.User;
import com.agent.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

@Service
public class UserService {

    @Autowired
    UserRepository userRepository;

    @Autowired
    PasswordEncoder encoder;


    public User userRegistration(RegistrationDTO registrationDTO) {

        Optional<User> optionalUser = userRepository.findByUsername(registrationDTO.getUsername());

        if(!optionalUser.isEmpty()) {
            return null;
        }

        User user = new User();
        user.setUsername(registrationDTO.getUsername());
        user.setPassword(encoder.encode(registrationDTO.getPassword()));
        user.setEmail(registrationDTO.getEmail());
        user.setName(registrationDTO.getName());
        user.setSurname(registrationDTO.getSurname());
        user.setPhoneNumber(registrationDTO.getPhoneNumber());
        user.setGender(registrationDTO.getGender());
        user.setBirthDate(registrationDTO.getBirthDate());
        user.setType(registrationDTO.getType());

        return userRepository.save(user);
    }

    public User getCurrentUser() {

        String username = SecurityUtils.getCurrentUserLogin().get();
        return userRepository.findByUsername(username).get();
    }

    
    public List<String> companyOwners() {
    	List<String> owners = new ArrayList<>();
    	List<User> users = new ArrayList<>(); 
    	
    	users = userRepository.findByType("Company owner");
    	for(User user : users) {
    		owners.add(user.getUsername());
    	}
    	
    	return owners;
    }



}
