package com.user.UserMicroservice.controller;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.user.UserMicroservice.config.CustomUserDetailsService;
import com.user.UserMicroservice.dto.*;
import com.user.UserMicroservice.enums.LogEntryType;
import com.user.UserMicroservice.model.User;
import com.user.UserMicroservice.security.TokenUtil;
import com.user.UserMicroservice.service.LoggingService;
import com.user.UserMicroservice.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PostAuthorize;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.validation.BindingResult;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.servlet.view.RedirectView;

import java.io.IOException;
import java.time.LocalDateTime;
import java.util.List;

import javax.servlet.http.HttpServletRequest;
import javax.validation.Valid;

@RestController
@CrossOrigin(origins = "*")
@RequestMapping("api/users")
public class UserController {

    @Autowired
    private UserService userService;

    @Autowired
    private TokenUtil tokenUtils;
    
    private BCryptPasswordEncoder bCryptPasswordEncoder = new BCryptPasswordEncoder();

    @Autowired
    private CustomUserDetailsService customUserService;
    
    @Autowired
    LoggingService loggingService;

    @PostMapping(consumes = "application/json", path = "/register")
    public ResponseEntity<?> registerClient(HttpServletRequest request, @RequestBody RegistrationDTO registrationDTO, BindingResult result) { 	
    	if (result.hasErrors()){
    		System.out.println("Result has errors!");
    		return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
        }
    	User user = userService.userRegistration(registrationDTO, request);	
        if(user == null) {
        	System.out.println("Created user is null!");
            return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
        }
        try {
			loggingService.log(LogEntryType.NOTIFICATION, "DATA_NU", request.getRemoteAddr(), user.getEmail());
		} catch (IOException e) {
			e.printStackTrace();
		}
        return new ResponseEntity<>(HttpStatus.CREATED);
    }

    @PostMapping(path = "/login")
    public ResponseEntity<?> login(@RequestBody LoginDTO loginDTO, HttpServletRequest request) {

        User user = customUserService.findUserByUsername(loginDTO.getUsername());

        if(loggingService.containsPotentialSQLInjection(loginDTO.getUsername())) {
        	try {
    			loggingService.log(LogEntryType.ERROR, "DATA_SI", request.getRemoteAddr());
    		} catch (IOException e) {
    			e.printStackTrace();
    		}
        }
        if (user == null || !user.isActivated() 
        		|| !loginDTO.getUsername().equals(user.getUsername())){
            return  ResponseEntity.ok(HttpStatus.UNAUTHORIZED);
        }
        else if(loginDTO.getPassword()!=null && loginDTO.getPassword().length()>1
                && !bCryptPasswordEncoder.matches(loginDTO.getPassword(), user.getPassword()) ) {
        	try {
    			loggingService.log(LogEntryType.WARNING, "AUTH_FL", request.getRemoteAddr(), user.getEmail());
    		} catch (IOException e) {
    			e.printStackTrace();
    		}
        	return  ResponseEntity.ok(HttpStatus.UNAUTHORIZED);
        	
        }
        else if(loginDTO.getCode()!=null  && ((!bCryptPasswordEncoder.matches(loginDTO.getCode(), user.getLoginCode())) || LocalDateTime.now().isAfter(user.getLoginCodeValidity()))){
        	System.out.println("Login code is not valid!");
        	return  ResponseEntity.ok(HttpStatus.UNAUTHORIZED);
        }

        String token = tokenUtils.generateToken(user.getUsername(), user.getType());
        user.setLoginCode(null);
        user.setLoginCodeValidity(null);
        customUserService.saveUser(user);

        LoginResponseDTO responseDTO = new LoginResponseDTO();
        responseDTO.setToken(token);
        try {
			loggingService.log(LogEntryType.NOTIFICATION, "AUTH_SL", request.getRemoteAddr(), user.getEmail());
		} catch (IOException e) {
			e.printStackTrace();
		}
        return ResponseEntity.ok(responseDTO);
    }
    
    @GetMapping(path = "/current")
    public ResponseEntity<?> getCurrentUser() {

        return new ResponseEntity<>(userService.getCurrentUser(), HttpStatus.OK);
    }

    @PutMapping()
    @PreAuthorize("hasAuthority('editInfo')")
    public ResponseEntity<?> edit(@RequestBody UserDTO dto, HttpServletRequest request) {
        User user = userService.edit(dto);
        try {
			loggingService.log(LogEntryType.NOTIFICATION, "DATA_EU", request.getRemoteAddr(), user.getEmail());
		} catch (IOException e) {
			e.printStackTrace();
		}
        return new ResponseEntity<>(user, HttpStatus.OK);
    }

    @PostMapping(path = "/changePassword")
    @PreAuthorize("hasAuthority('changePassword')")
    public ResponseEntity<?> changePassword(@RequestBody ChangePasswordDTO changePasswordDTO, HttpServletRequest request) {
        User user = userService.changePassword(changePasswordDTO);

        if(user == null) {
        	try {
    			loggingService.log(LogEntryType.ERROR, "DATA_XU", request.getRemoteAddr());
    		} catch (IOException e) {
    			e.printStackTrace();
    		}
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
    		try {
    			loggingService.log(LogEntryType.ERROR, "DATA_XU", request.getRemoteAddr());
    		} catch (IOException e) {
    			e.printStackTrace();
    		}
    		return ResponseEntity.ok(HttpStatus.BAD_REQUEST);
    	}
    	userService.forgottenPassword(user, request);
    	try {
			loggingService.log(LogEntryType.ERROR, "DATA_PC", request.getRemoteAddr(), user.getEmail());
		} catch (IOException e) {
			e.printStackTrace();
		}
    	return ResponseEntity.ok(HttpStatus.OK);
    }
    
    @PostMapping(path = "/loginCode")
    public ResponseEntity<?> getLoginCode(HttpServletRequest request, @RequestBody ForgottenPasswordDTO dto  )
    {
    	User user = customUserService.findUserByUsername(dto.getUsername());
    	if(user == null) {
    		return ResponseEntity.ok(HttpStatus.BAD_REQUEST);
    	}
    	userService.getLoginCode(user, request);
    	return new ResponseEntity<String>("Code is sent to your email address!", HttpStatus.OK);
    }
    
    @GetMapping(path = "/users")
    public ResponseEntity<?> users(){
        List<User> users = userService.users();
        return new ResponseEntity<>(users, HttpStatus.OK);
    }
    
    @GetMapping(path = "/public")
    public ResponseEntity<?> getPublicProfile() {

    	List<User> users = userService.getPublicProfile();
    	
        if(users.isEmpty()) {
            return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
        }

        return new ResponseEntity<>(users, HttpStatus.OK);
    }
    
	@RequestMapping(method = RequestMethod.POST, value = "/checkActivationCode",consumes = "text/plain",produces = MediaType.APPLICATION_JSON_VALUE)
	@CrossOrigin(origins = "*")
	public ResponseEntity<String> checkActivationCode(@RequestBody String c) throws JsonProcessingException{
		System.out.println("Activation code "+c);
		String code = c;	
		if(userService.userAlreadyActivated(code)) {
			System.out.println("Already validated!");
			return ResponseEntity.ok("already validated");
		} 
		
		boolean valid = userService.checkActivationCode(code);
		
		if(valid) {
			System.out.println("Valid!");
			return new ResponseEntity<String>("Valid", HttpStatus.OK);
		}
		
		else {
			System.out.println("Activation code expired!");
			return ResponseEntity.ok("Activation code expired!");
		}
	}
	
	@RequestMapping(method = RequestMethod.POST, value = "/checkForgottenPassword",consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
	@CrossOrigin(origins = "*")
	public ResponseEntity<?> resetPassword( @RequestBody ResetPasswordDTO dto) throws JsonProcessingException{
		System.out.println("Password reset code "+dto.getCode());
		String code = dto.getCode();
		
		
		if(!userService.resetCodeExists(code))
			return new ResponseEntity<String>(HttpStatus.NOT_FOUND);
		User user = userService.findByPasswordResetCode(code);

		
		boolean valid = userService.checkPasswordResetCode(dto);
		
		if(valid) return new ResponseEntity<String>(HttpStatus.OK);
		else {
			System.out.println("Expired!");
			return new ResponseEntity<String>(HttpStatus.GATEWAY_TIMEOUT);
		}
	}
    
    
    @GetMapping(path = "/user/{username}")
    @PreAuthorize("hasAuthority('getUserInfo')")
    public ResponseEntity<?> getUser(@PathVariable String username) {

    	User user = userService.getUser(username);
    	
        if(user == null) {
            return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
        }

        return new ResponseEntity<>(user, HttpStatus.OK);
    }
}
