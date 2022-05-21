package com.user.UserMicroservice.controller;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.user.UserMicroservice.config.CustomUserDetailsService;
import com.user.UserMicroservice.dto.*;
import com.user.UserMicroservice.model.User;
import com.user.UserMicroservice.security.TokenUtil;
import com.user.UserMicroservice.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.web.bind.annotation.*;

import java.util.List;

import javax.servlet.http.HttpServletRequest;

@RestController
@CrossOrigin(origins = "http://localhost:4200")
@RequestMapping("api/users")
public class UserController {

    @Autowired
    private UserService userService;

    @Autowired
    private TokenUtil tokenUtils;
    
    @Autowired
    private BCryptPasswordEncoder bCryptPasswordEncoder;

    @Autowired
    private CustomUserDetailsService customUserService;

    @PostMapping(consumes = "application/json", path = "/register")
    public ResponseEntity<?> registerClient(HttpServletRequest request, @RequestBody RegistrationDTO registrationDTO) {

        User user = userService.userRegistration(registrationDTO, request);

        if(user == null) {
            return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
        }

        return new ResponseEntity<>(HttpStatus.CREATED);
    }

    @PostMapping(path = "/login")
    public ResponseEntity<?> login(@RequestBody LoginDTO loginDTO) {

        User user = customUserService.findUserByUsername(loginDTO.getUsername());

        if (user == null || !bCryptPasswordEncoder.matches(loginDTO.getPassword(), user.getPassword()) || !loginDTO.getUsername().equals(user.getUsername())) {
        	System.out.println(bCryptPasswordEncoder.matches(loginDTO.getPassword(), user.getPassword()));
        	System.out.println("Bcripted pass"+bCryptPasswordEncoder.encode(loginDTO.getPassword()));
        	System.out.println(user.getPassword());
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
    public ResponseEntity<?> edit(@RequestBody UserDTO dto) {
        User user = userService.edit(dto);

        return new ResponseEntity<>(HttpStatus.OK);
    }

    @PostMapping(path = "/password")
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
    
    @PostMapping(path = "/forgottenpassword")
    public ResponseEntity<?> forgottenPassword(HttpServletRequest request, @RequestBody ForgottenPasswordDTO dto){
    	User user = customUserService.findUserByUsername(dto.getUsername());
    	if(user == null) {
    		return ResponseEntity.ok(HttpStatus.BAD_REQUEST);
    	}
    	userService.forgottenPassword(user, request);
    	return ResponseEntity.ok(HttpStatus.OK);
    }
    
    @GetMapping(path = "/public")
    public ResponseEntity<?> getPublicProfile() {

    	List<String> users = userService.getPublicProfile();
    	
        if(users.isEmpty()) {
            return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
        }

        return new ResponseEntity<>(users, HttpStatus.OK);
    }
    
	@RequestMapping(method = RequestMethod.POST, value = "/checkActivationCode",consumes = MediaType.APPLICATION_JSON_VALUE,produces = MediaType.APPLICATION_JSON_VALUE)
	@CrossOrigin(origins = "*")
	public ResponseEntity<String> checkActivationCode(@RequestBody String c) throws JsonProcessingException{
		System.out.println("Password reset code "+c);
		String code = c.substring(1,c.length()-1);
		
		
		if(userService.userAlreadyActivated(code)) return ResponseEntity.ok("already validated");
				
		boolean valid = userService.checkActivationCode(code);
		
		if(valid) return ResponseEntity.ok("valid");
		else return ResponseEntity.ok("Acivation code expired!");
	}
	
	@RequestMapping(method = RequestMethod.POST, value = "/checkForgottenPassword",consumes = MediaType.APPLICATION_JSON_VALUE,produces = MediaType.APPLICATION_JSON_VALUE)
	@CrossOrigin(origins = "*")
	public ResponseEntity<String> resetPassword( @RequestBody String c) throws JsonProcessingException{
		System.out.println("Password reset code "+c);
		String code = c.substring(1,c.length()-1);
		
		
		if(!userService.resetCodeExists(code)) return ResponseEntity.ok("Code is not valid");
				
		boolean valid = userService.checkPasswordResetCode(code);
		
		if(valid) return ResponseEntity.ok("valid");
		else return ResponseEntity.ok("Reset code expired!");
	}
	
	
    
    
}
