package com.user.UserMicroservice.controller;

import com.user.UserMicroservice.dto.RegistrationDTO;
import com.user.UserMicroservice.model.User;
import com.user.UserMicroservice.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.web.bind.annotation.*;

@RestController
@CrossOrigin(origins = "http://localhost:4200")
@RequestMapping("api/users")
public class UserController {

    @Autowired
    private UserService userService;

    @PostMapping(consumes = "application/json", path = "/register")
    public ResponseEntity<?> registerClient(@RequestBody RegistrationDTO registrationDTO) {

        User user = userService.userRegistration(registrationDTO);

        if(user == null) {
            return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
        }

        return new ResponseEntity<>(HttpStatus.CREATED);
    }
}
