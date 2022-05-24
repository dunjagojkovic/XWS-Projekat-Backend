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
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.validation.BindingResult;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.servlet.view.RedirectView;

import java.time.LocalDateTime;
import java.util.List;

import javax.servlet.http.HttpServletRequest;
import javax.validation.Valid;

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
        System.out.println("User created!");
        return new ResponseEntity<>(HttpStatus.CREATED);
    }

    @PostMapping(path = "/login")
    public ResponseEntity<?> login(@RequestBody LoginDTO loginDTO) {
    	System.out.println("Login dto"+loginDTO);

        User user = customUserService.findUserByUsername(loginDTO.getUsername());

        if (user == null || !user.isActivated() 
        		|| !loginDTO.getUsername().equals(user.getUsername())
        		){
        	System.out.println(bCryptPasswordEncoder.matches(loginDTO.getPassword(), user.getPassword()));
        	System.out.println("Bcripted pass"+bCryptPasswordEncoder.encode(loginDTO.getPassword()));
        	System.out.println(user.getPassword());
            return  ResponseEntity.ok(HttpStatus.UNAUTHORIZED);
        }
        else if(loginDTO.getPassword()!=null && loginDTO.getPassword().length()>1 && !bCryptPasswordEncoder.matches(loginDTO.getPassword(), user.getPassword()) ) {
        	System.out.println("Wrong password");
        	return  ResponseEntity.ok(HttpStatus.UNAUTHORIZED);
        	
        }
        else if(loginDTO.getCode()!=null  && ((!bCryptPasswordEncoder.matches(loginDTO.getCode(), user.getLoginCode())) || LocalDateTime.now().isAfter(user.getLoginCodeValidity()))){
        	System.out.println("Login code is not valid!");
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
    @PreAuthorize("hasAuthority('User')")
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
    
    @GetMapping(path = "/public")
    public ResponseEntity<?> getPublicProfile() {

    	List<String> users = userService.getPublicProfile();
    	
        if(users.isEmpty()) {
            return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
        }

        return new ResponseEntity<>(users, HttpStatus.OK);
    }
    
	@RequestMapping(method = RequestMethod.POST, value = "/checkActivationCode",consumes = "application/json",produces = MediaType.APPLICATION_JSON_VALUE)
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
			System.out.println("Acivation code expired!");
			return ResponseEntity.ok("Acivation code expired!");
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
	
	
    
    
}
