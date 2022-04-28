package com.user.UserMicroservice.service;

import com.user.UserMicroservice.dto.RegistrationDTO;
import com.user.UserMicroservice.model.User;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import com.user.UserMicroservice.repository.UserRepository;
import com.user.UserMicroservice.config.SecurityUtils;

import java.util.Optional;

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
        user.setPublic(false);
        user.setBirthDate(registrationDTO.getBirthDate());

        return userRepository.save(user);

    }

    public User getCurrentUser() {

        String username = SecurityUtils.getCurrentUserLogin().get();
        return userRepository.findByUsername(username).get();
    }

}
