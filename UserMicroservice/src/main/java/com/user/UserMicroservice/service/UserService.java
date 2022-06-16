package com.user.UserMicroservice.service;

import com.user.UserMicroservice.config.SecurityUtils;
import com.user.UserMicroservice.dto.ChangePasswordDTO;
import com.user.UserMicroservice.dto.RegistrationDTO;
import com.user.UserMicroservice.dto.ResetPasswordDTO;
import com.user.UserMicroservice.dto.UserDTO;
import com.user.UserMicroservice.model.User;
import com.user.UserMicroservice.repository.UserRepository;

import ch.qos.logback.core.net.server.Client;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

import javax.servlet.http.HttpServletRequest;

@Service
public class UserService {

    @Autowired
    UserRepository userRepository;
    
    BCryptPasswordEncoder bCryptPasswordEncoder = new BCryptPasswordEncoder();
    
    @Autowired
    CodeService codeService;
    
    @Autowired
    MailService<String> mailService;
   


    public User userRegistration(RegistrationDTO registrationDTO, HttpServletRequest request) {

        Optional<User> optionalUser = userRepository.findByUsername(registrationDTO.getUsername());

        if(!optionalUser.isEmpty()) {
            return null;
        }

        User user = new User();
        user.setUsername(registrationDTO.getUsername());
        user.setPassword(bCryptPasswordEncoder.encode(registrationDTO.getPassword()));
        user.setEmail(registrationDTO.getEmail());
        user.setName(registrationDTO.getName());
        user.setSurname(registrationDTO.getSurname());
        user.setPhoneNumber(registrationDTO.getPhoneNumber());
        user.setGender(registrationDTO.getGender());
        user.setPublic(true);
        user.setBirthDate(registrationDTO.getBirthDate());
        String activationCode = codeService.generateActivationCodeForUSer(user);
        user.setActivationCode(bCryptPasswordEncoder.encode(activationCode));
        user.setActivated(false);
        user.setActivationCodeValidity(LocalDateTime.now().plusDays(5));
        
        mailService.sendUserRegistrationMail(user.getEmail(), activationCode, getSiteURL(request));
        user.setType("User");

        return userRepository.save(user);
    }
    
    public User getByUsername(String username) {
    	return userRepository.findByUsername(username).orElse(null);
    }

    public User getCurrentUser() {

        String username = SecurityUtils.getCurrentUserLogin().get();
        System.out.println(username);
        return userRepository.findByUsername(username).get();
    }

    public User edit(UserDTO userDTO) {

        //Optional<User> optionalUser = userRepository.findByUsername(userDTO.getUsername());
        Optional<User> optionalUser = userRepository.findById(getCurrentUser().getId());


        if (userDTO.getName() != null && !userDTO.getName().equals("")){
            optionalUser.get().setName(userDTO.getName());
        }
        if (userDTO.getEmail() != null && !userDTO.getEmail().equals("")) {
            optionalUser.get().setEmail(userDTO.getEmail());
        }
        if (userDTO.getSurname() != null && !userDTO.getSurname().equals("")) {
            optionalUser.get().setSurname(userDTO.getSurname());
        }
        if (userDTO.getPhoneNumber() != null && !userDTO.getPhoneNumber().equals("")) {
            optionalUser.get().setPhoneNumber(userDTO.getPhoneNumber());
        }
        if (userDTO.getBiography() != null && !userDTO.getBiography().equals("")) {
            optionalUser.get().setBiography(userDTO.getBiography());
        }
        if (userDTO.getBirthDate() != null && !userDTO.getBirthDate().equals("")) {
            optionalUser.get().setBirthDate(userDTO.getBirthDate());
        }
        if (userDTO.getGender() != null && !userDTO.getGender().equals("")) {
            optionalUser.get().setGender(userDTO.getGender());
        }
        if (userDTO.getEducation() != null && !userDTO.getEducation().equals("")) {
            optionalUser.get().setEducation(userDTO.getEducation());
        }
        if (userDTO.getHobby() != null && !userDTO.getHobby().equals("")) {
            optionalUser.get().setHobby(userDTO.getHobby());
        }
        if (userDTO.getUsername() != null && !userDTO.getUsername().equals("")) {
            optionalUser.get().setUsername(userDTO.getUsername());
        }
        if (userDTO.getWorkExperience() != null && !userDTO.getWorkExperience().equals("")) {
            optionalUser.get().setWorkExperience(userDTO.getWorkExperience());
        }
        if (userDTO.getInterest() != null && !userDTO.getInterest().equals("")) {
            optionalUser.get().setInterest(userDTO.getInterest());
        }
        if (userDTO.getPublic() != null && !userDTO.getPublic().equals("")) {
            optionalUser.get().setPublic(userDTO.getPublic());
        }
        
        return userRepository.save(optionalUser.get());
    }

    public User changePassword(ChangePasswordDTO changePasswordDTO) {

        User user = getCurrentUser();

        if(!bCryptPasswordEncoder.matches(changePasswordDTO.getOldPassword(), user.getPassword())) {
            return null;
        }

        user.setPassword(bCryptPasswordEncoder.encode(changePasswordDTO.getNewPassword()));

        return userRepository.save(user);
    }

    public List<User> filterUsers(UserDTO dto) {
        List<User> users = userRepository.findAllByIsPublic(true);
        List<User> results = new ArrayList<>();

        for(User user: users){
            if(user.getUsername().toLowerCase().contains(dto.getSearchTerm().toLowerCase()) || user.getName().toLowerCase().contains(dto.getSearchTerm().toLowerCase()) || user.getSurname().toLowerCase().contains(dto.getSearchTerm().toLowerCase())){
                if(!userExists(user, results)){
                    results.add(user);
                }
            }
        }
        return results;
    }
    
    public List<User> users() {
        List<User> users = userRepository.findAll();
        List<User> result = new ArrayList<>();
        
        User user = getCurrentUser();
        
        for(User u: users) {
        	
        	if(!u.getUsername().equals(user.getUsername())) {
        		result.add(u);
        	}
        	
        }
        
        return result;
    }
    
    

    public boolean userExists(User user, List<User> users) {

        for(User u: users){
            if(u.getId().equals(user.getId())){
                return true;
            }
        }
        return false;
    }
    
    public List<User> getPublicProfile() {
    	
    	List<User> users = userRepository.findAllByIsPublic(true);
        List<User> results = new ArrayList<>();
//    	List<String> usernames = new ArrayList<>();
    	
    	for(User user: users) {
//            usernames.add(user.getUsername());
    		results.add(user);
    	}

//        return usernames;
    	return results;
    	
    }
    
    public User getUser(String username) {
    	
    	User user = userRepository.findByUsername(username).get();
    	System.out.println(username);
        
    	return user;
    	
    }

	public void forgottenPassword(User user, HttpServletRequest request) {
		String resetCode = codeService.generatePasswordResetCode(user);
		mailService.sendLinkToResetPassword(user.getEmail(), resetCode, getSiteURL(request));
		user.setPasswordResetCode(bCryptPasswordEncoder.encode(resetCode));
		user.setPasswordResetCodeValidity(LocalDateTime.now().plusMinutes(5));
		userRepository.save(user);
		
	}

	public boolean userAlreadyActivated(String code) {
		User user = findByActivation(code);
		System.out.println("Found user= "+user);
		return user!=null && user.isActivated();
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

	public boolean resetCodeExists(String code) {
		User u = findByPasswordResetCode(code);
		return u!=null;
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
	
	private String getSiteURL(HttpServletRequest request) {
		return request.getHeader("origin");
	}

	public void getLoginCode(User user, HttpServletRequest request) {
		String loginCode = codeService.generateLoginCode(user);
		mailService.sendCodetToEmail(user.getEmail(), loginCode, getSiteURL(request));
		user.setLoginCode(bCryptPasswordEncoder.encode(loginCode));
		user.setLoginCodeValidity(LocalDateTime.now().plusMinutes(5));
		userRepository.save(user);
		
	}
}
