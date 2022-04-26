package com.user.UserMicroservice.controller;

import com.user.UserMicroservice.config.CustomUserDetailsService;
import com.user.UserMicroservice.dto.LoginDTO;
import com.user.UserMicroservice.dto.LoginResponseDTO;
import com.user.UserMicroservice.dto.RegistrationDTO;
import com.user.UserMicroservice.model.User;
import com.user.UserMicroservice.security.TokenUtil;
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

    @Autowired
    private TokenUtil tokenUtils;

    @Autowired
    private PasswordEncoder passwordEncoder;

    @Autowired
    private CustomUserDetailsService customUserService;

    @PostMapping(consumes = "application/json", path = "/register")
    public ResponseEntity<?> registerClient(@RequestBody RegistrationDTO registrationDTO) {

        User user = userService.userRegistration(registrationDTO);

        if(user == null) {
            return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
        }

        return new ResponseEntity<>(HttpStatus.CREATED);
    }

    @PostMapping(path = "/login")
    public ResponseEntity<?> login(@RequestBody LoginDTO loginDTO) {

        User user = customUserService.findByUsername(loginDTO.getUsername());

        if (user == null || !passwordEncoder.matches(loginDTO.getPassword(), user.getPassword()) || !loginDTO.getUsername().equals(user.getUsername())) {
            return  ResponseEntity.ok(HttpStatus.UNAUTHORIZED);
        }

        String token = tokenUtils.generateToken(user.getUsername());
        LoginResponseDTO responseDTO = new LoginResponseDTO();
        responseDTO.setToken(token);
        return ResponseEntity.ok(responseDTO);
    }
}
