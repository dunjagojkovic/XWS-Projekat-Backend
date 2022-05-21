package com.user.UserMicroservice.service;

import com.user.UserMicroservice.config.SecurityUtils;
import com.user.UserMicroservice.dto.ChangePasswordDTO;
import com.user.UserMicroservice.dto.RegistrationDTO;
import com.user.UserMicroservice.dto.UserDTO;
import com.user.UserMicroservice.model.User;
import com.user.UserMicroservice.repository.UserRepository;

import ch.qos.logback.core.net.server.Client;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

import javax.servlet.http.HttpServletRequest;

@Service
public class UserService {

    @Autowired
    UserRepository userRepository;
    
    @Autowired
    BCryptPasswordEncoder bCryptPasswordEncoder;
    
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

        Optional<User> optionalUser = userRepository.findByUsername(userDTO.getUsername());

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
            if (userDTO.getPassword() != null && !userDTO.getPassword().equals("")) {
                optionalUser.get().setPassword(userDTO.getPassword());
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

    public boolean userExists(User user, List<User> users) {

        for(User u: users){
            if(u.getId().equals(user.getId())){
                return true;
            }
        }
        return false;
    }
    
    public List<String> getPublicProfile() {
    	
    	List<User> users = userRepository.findAllByIsPublic(true);
    	
    	List<String> usernames = new ArrayList<>();
    	
    	for(User user: users) {
    		usernames.add(user.getUsername());
    	}
    	
    	return usernames;
    	
    }

	public void forgottenPassword(User user, HttpServletRequest request) {
		String resetCode = codeService.generatePasswordResetCode(user);
		mailService.sendLinkToResetPassword(user.getEmail(), resetCode, getSiteURL(request));
		user.setPasswordResetCode(bCryptPasswordEncoder.encode(resetCode));
		user.setPasswordResetCodeValidity(LocalDateTime.now().plusMinutes(5));
		userRepository.save(user);
		
	}

	public boolean userAlreadyActivated(String code) {
		return userRepository.findByActivationCodeAndActivatedTrue(bCryptPasswordEncoder.encode(code))!=null;
	}

	public boolean checkActivationCode(String code) {
		System.out.println("akt kod koji se trazi "+ code);
    	User u = userRepository.findByActivationCode(bCryptPasswordEncoder.encode(code)).get();
    	 if(u!=null && LocalDateTime.now().isBefore(u.getActivationCodeValidity())) {
    		 u.setActivated(true);
    		 userRepository.save(u);
    		 return true;
    	 }
		return false;
	}

	public boolean resetCodeExists(String code) {
		return userRepository.findByPasswordResetCode(bCryptPasswordEncoder.encode(code))!=null;
	}

	public boolean checkPasswordResetCode(String code) {
		System.out.println("akt kod koji se trazi "+ code);
    	User u = userRepository.findByPasswordResetCode(bCryptPasswordEncoder.encode(code)).get();
    	 if(u!=null && LocalDateTime.now().isBefore(u.getPasswordResetCodeValidity())) {
    		 u.setPasswordResetCode(null);
    		 u.setPasswordResetCodeValidity(null);
    		 userRepository.save(u);
    		 return true;
    	 }
		return false;
	}
	
	private String getSiteURL(HttpServletRequest request) {
		return request.getHeader("origin");
	}
}
