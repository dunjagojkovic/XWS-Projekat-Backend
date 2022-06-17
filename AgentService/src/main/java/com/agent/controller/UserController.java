package com.agent.controller;

import com.agent.config.CustomUserDetailsService;
import com.agent.dto.*;
import com.agent.model.User;
import com.agent.security.TokenUtil;
import com.agent.service.UserService;
import com.fasterxml.jackson.core.JsonProcessingException;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import java.time.LocalDateTime;
import java.util.List;
import java.util.UUID;

@RestController
@CrossOrigin(origins = "*")
@RequestMapping("api/users")
public class UserController {

    @Autowired
    private UserService userService;

    @Autowired
    private PasswordEncoder passwordEncoder;

    private BCryptPasswordEncoder bCryptPasswordEncoder = new BCryptPasswordEncoder();

    @Autowired
    private TokenUtil tokenUtils;

    @Autowired
    private CustomUserDetailsService customUserService;


    @PostMapping(consumes = "application/json", path = "/registerUser")
    public ResponseEntity<?> registerUser(HttpServletRequest request, @RequestBody RegistrationDTO registrationDTO) {

        User user = userService.userRegistration(registrationDTO, request);

        if(user == null) {
            return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
        }

        return new ResponseEntity<>(HttpStatus.CREATED);
    }

    @PostMapping(path = "/login")
    public ResponseEntity<?> login(@RequestBody LoginDTO loginDTO) {

        User user = customUserService.findUserByUsername(loginDTO.getUsername());
        if (user == null || !loginDTO.getUsername().equals(user.getUsername())
        || !user.isActivated()) {

            return  new ResponseEntity<>("User does not exist or not activated!", HttpStatus.UNAUTHORIZED);
        }
        else if(loginDTO.getPassword()!=null && loginDTO.getPassword().length()>1
                && !bCryptPasswordEncoder.matches(loginDTO.getPassword(), user.getPassword()) ) {

            System.out.println("Wrong password!");
            return new ResponseEntity<>("Wrong password!", HttpStatus.UNAUTHORIZED);
        }
        else if(loginDTO.getCode()!=null
                && ((!bCryptPasswordEncoder.matches(loginDTO.getCode(), user.getLoginCode()))
                || LocalDateTime.now().isAfter(user.getLoginCodeValidity()))){

            System.out.println("Login code is not valid!");
            return  ResponseEntity.ok(HttpStatus.UNAUTHORIZED);
        }

        user.setLoginCode(null);
        user.setLoginCodeValidity(null);
        customUserService.saveUser(user);

        String token = tokenUtils.generateToken(user.getUsername());
        LoginResponseDTO responseDTO = new LoginResponseDTO();
        responseDTO.setToken(token);
        String key = UUID.randomUUID().toString();
        responseDTO.setKey(key);


        return new ResponseEntity<>(responseDTO, HttpStatus.OK);
    }
    @GetMapping(path = "/current")
    public ResponseEntity<?> getCurrentUser() {
        return new ResponseEntity<>(userService.getCurrentUser(), HttpStatus.OK);
    }
    
    @GetMapping(path = "/owners")
    public ResponseEntity<?> getCompanyOwners() {
        return new ResponseEntity<>(userService.companyOwners(), HttpStatus.OK);
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
    public ResponseEntity<?> resetPassword( @RequestBody ResetPasswordDTO dto) throws JsonProcessingException {
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
