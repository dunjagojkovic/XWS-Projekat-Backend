package com.user.UserMicroservice.controller;

import com.user.UserMicroservice.config.CustomUserDetailsService;
import com.user.UserMicroservice.dto.*;
import com.user.UserMicroservice.model.User;
import com.user.UserMicroservice.security.TokenUtil;
import com.user.UserMicroservice.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.web.bind.annotation.*;

import java.util.List;

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
        return ResponseEntity.ok(responseDTO);
    }
    @GetMapping(path = "/current")
    public ResponseEntity<?> getCurrentUser() {

        return new ResponseEntity<>(userService.getCurrentUser(), HttpStatus.OK);
    }

    @PutMapping()
    @PreAuthorize("hasAuthority('User')")
    public ResponseEntity<?> edit(@RequestBody UserDTO dto) {
        User user = userService.edit(dto);

        return new ResponseEntity<>(user, HttpStatus.OK);
    }

    @PostMapping(path = "/changePassword")
    @PreAuthorize("hasAuthority('User') and hasPermission('hasAccess', 'WRITE')")
    public ResponseEntity<?> changePassword(@RequestBody ChangePasswordDTO changePasswordDTO) {
        User user = userService.changePassword(changePasswordDTO);

        if(user == null) {
            return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
        }

        return new ResponseEntity<>(HttpStatus.OK);
    }

    @PostMapping(path = "/filterUsers")
    public ResponseEntity<?> filterUsers(@RequestBody UserDTO dto){
        List<User> users = userService.filterUsers(dto);
        return new ResponseEntity<>(users, HttpStatus.OK);
    }
    
    @GetMapping(path = "/public")
    public ResponseEntity<?> getPublicProfile() {

    	List<String> users = userService.getPublicProfile();
    	
        if(users.isEmpty()) {
            return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
        }

        return new ResponseEntity<>(users, HttpStatus.OK);
    }
}
