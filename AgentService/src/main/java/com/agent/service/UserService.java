package com.agent.service;

import com.agent.config.SecurityUtils;
import com.agent.dto.RegistrationDTO;
import com.agent.dto.ResetPasswordDTO;
import com.agent.model.User;
import com.agent.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import javax.servlet.http.HttpServletRequest;
import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

@Service
public class UserService {

    @Autowired
    UserRepository userRepository;

    @Autowired
    PasswordEncoder encoder;

    @Autowired
    CodeService codeService;

    @Autowired
    MailService<String> mailService;

    BCryptPasswordEncoder bCryptPasswordEncoder = new BCryptPasswordEncoder();


    public User userRegistration(RegistrationDTO registrationDTO, HttpServletRequest request) {

        Optional<User> optionalUser = userRepository.findByUsername(registrationDTO.getUsername());

        if(!optionalUser.isEmpty()) {
            return null;
        }

        User user = new User();
        user.setUsername(registrationDTO.getUsername());
        user.setEmail(registrationDTO.getEmail());
        user.setName(registrationDTO.getName());
        user.setSurname(registrationDTO.getSurname());
        user.setPhoneNumber(registrationDTO.getPhoneNumber());
        user.setGender(registrationDTO.getGender());
        user.setBirthDate(registrationDTO.getBirthDate());
        user.setType(registrationDTO.getType());
        user.setActivated(false);
        user.setPassword(bCryptPasswordEncoder.encode(registrationDTO.getPassword()));

        String activationCode = codeService.generateActivationCodeForUSer(user);
        user.setActivationCode(bCryptPasswordEncoder.encode(activationCode));
        user.setActivationCodeValidity(LocalDateTime.now().plusDays(5));

        mailService.sendUserRegistrationMail(user.getEmail(), activationCode, getSiteURL(request));

        return userRepository.save(user);
    }

    public User getCurrentUser() {

        String username = SecurityUtils.getCurrentUserLogin().get();
        return userRepository.findByUsername(username).get();
    }

    
    public List<String> companyOwners() {
    	List<String> owners = new ArrayList<>();
    	List<User> users = new ArrayList<>(); 
    	
    	users = userRepository.findByType("Company owner");
    	for(User user : users) {
    		owners.add(user.getUsername());
    	}
    	
    	return owners;
    }

    private String getSiteURL(HttpServletRequest request) {
        return request.getHeader("origin");
    }


    public boolean resetCodeExists(String code) {
        User u = findByPasswordResetCode(code);
        return u!=null;
    }

    public User findByPasswordResetCode(String code) {
        List<User> allUsers = userRepository.findAll();
        User foundUser = null;
        for(User u : allUsers) {
            if(bCryptPasswordEncoder.matches(code, u.getPasswordResetCode())) {
                System.out.println("Maching!");
                foundUser = u;
            }
        }
        return foundUser;
    }

    public boolean checkPasswordResetCode(ResetPasswordDTO dto) {
        System.out.println("akt kod koji se trazi "+ dto.getCode());
        User u = findByPasswordResetCode(dto.getCode());
        if(u!=null && LocalDateTime.now().isBefore(u.getPasswordResetCodeValidity())) {
            u.setPasswordResetCode(null);
            u.setPasswordResetCodeValidity(null);
            u.setPassword(bCryptPasswordEncoder.encode(dto.getNewPassword()));
            userRepository.save(u);
            return true;
        }
        return false;
    }

    public boolean userAlreadyActivated(String code) {
        User user = findByActivation(code);
        System.out.println("Found user= "+user);
        return user!=null && user.isActivated();
    }

    private User findByActivation(String code) {
        List<User> allUsers = userRepository.findAll();
        User foundUser = null;
        for(User u : allUsers) {
            if(bCryptPasswordEncoder.matches(code, u.getActivationCode())) {
                System.out.println("Maching!");
                foundUser = u;
            }
        }
        return foundUser;
    }

    public boolean checkActivationCode(String code) {
        System.out.println("akt kod koji se trazi "+ code);
        User u = findByActivation(code);
        if(u!=null && LocalDateTime.now().isBefore(u.getActivationCodeValidity())) {
            u.setActivated(true);
            userRepository.save(u);
            return true;
        }
        return false;
    }

    public void getLoginCode(User user, HttpServletRequest request) {
        String loginCode = codeService.generateLoginCode(user);
        mailService.sendCodetToEmail(user.getEmail(), loginCode, getSiteURL(request));
        user.setLoginCode(bCryptPasswordEncoder.encode(loginCode));
        user.setLoginCodeValidity(LocalDateTime.now().plusMinutes(5));
        userRepository.save(user);
    }

    public void forgottenPassword(User user, HttpServletRequest request) {
        String resetCode = codeService.generatePasswordResetCode(user);
        mailService.sendLinkToResetPassword(user.getEmail(), resetCode, getSiteURL(request));
        user.setPasswordResetCode(bCryptPasswordEncoder.encode(resetCode));
        user.setPasswordResetCodeValidity(LocalDateTime.now().plusMinutes(5));
        userRepository.save(user);
    }
}
