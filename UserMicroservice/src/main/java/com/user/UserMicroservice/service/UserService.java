package com.user.UserMicroservice.service;

import com.user.UserMicroservice.dto.RegistrationDTO;
import com.user.UserMicroservice.model.User;
import com.user.UserMicroservice.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.dao.PessimisticLockingFailureException;
import org.springframework.http.HttpStatus;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import org.springframework.web.server.ResponseStatusException;

import java.util.Optional;

@Service
public class UserService {

    @Autowired
    UserRepository userRepository;


    public User userRegistration(RegistrationDTO registrationDTO) {

        Optional<User> optionalUser = userRepository.findByUsername(registrationDTO.getUsername());

        if(!optionalUser.isEmpty()) {
            return null;
        }

        User user = new User();
        user.setUsername(registrationDTO.getUsername());
        //user.setPassword(encoder.encode(registrationDTO.getPassword()));
        user.setPassword(registrationDTO.getPassword());
        user.setEmail(registrationDTO.getEmail());
        user.setCity(registrationDTO.getCity());
        user.setCountry(registrationDTO.getCountry());
        user.setName(registrationDTO.getName());
        user.setSurname(registrationDTO.getSurname());
        user.setPhoneNumber(registrationDTO.getPhoneNumber());
        user.setGender(registrationDTO.getGender());
        user.setPublic(registrationDTO.getPublic());
        user.setBirthDate(registrationDTO.getBirthDate());

        return userRepository.save(user);

    }

}
