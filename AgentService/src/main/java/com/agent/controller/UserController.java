package com.agent.controller;

import com.agent.config.CustomUserDetailsService;
import com.agent.dto.LoginDTO;
import com.agent.dto.LoginResponseDTO;
import com.agent.dto.RegistrationDTO;
import com.agent.model.User;
import com.agent.security.TokenUtil;
import com.agent.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.UUID;

@RestController
@CrossOrigin(origins = "http://localhost:4200")
@RequestMapping("api/users")
public class UserController {

    @Autowired
    private UserService userService;

    @Autowired
    private PasswordEncoder passwordEncoder;

    @Autowired
    private TokenUtil tokenUtils;

    @Autowired
    private CustomUserDetailsService customUserService;


    @PostMapping(consumes = "application/json", path = "/registerUser")
    public ResponseEntity<?> registerUser(@RequestBody RegistrationDTO registrationDTO) {

        User user = userService.userRegistration(registrationDTO);

        if(user == null) {
            return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
        }

        return new ResponseEntity<>(HttpStatus.CREATED);
    }

    @PostMapping(path = "/login")
    public ResponseEntity<?> login(@RequestBody LoginDTO loginDTO) {

        User user = customUserService.findUserByUsername(loginDTO.getUsername());

        if (user == null || !passwordEncoder.matches(loginDTO.getPassword(), user.getPassword()) || !loginDTO.getUsername().equals(user.getUsername())) {
            return  ResponseEntity.ok(HttpStatus.UNAUTHORIZED);
        }

        String token = tokenUtils.generateToken(user.getUsername());
        LoginResponseDTO responseDTO = new LoginResponseDTO();
        responseDTO.setToken(token);
        String key = UUID.randomUUID().toString();
        responseDTO.setKey(key);

        return ResponseEntity.ok(responseDTO);
    }
    @GetMapping(path = "/current")
    public ResponseEntity<?> getCurrentUser() {

        return new ResponseEntity<>(userService.getCurrentUser(), HttpStatus.OK);
    }
    
    @GetMapping(path = "/owners")
    public ResponseEntity<?> getCompanyOwners() {

        return new ResponseEntity<>(userService.companyOwners(), HttpStatus.OK);
    }

}
