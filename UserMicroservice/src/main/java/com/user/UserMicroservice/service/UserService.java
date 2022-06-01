package com.user.UserMicroservice.service;

import com.user.UserMicroservice.config.SecurityUtils;
import com.user.UserMicroservice.dto.ChangePasswordDTO;
import com.user.UserMicroservice.dto.RegistrationDTO;
import com.user.UserMicroservice.dto.UserDTO;
import com.user.UserMicroservice.model.User;
import com.user.UserMicroservice.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

@Service
public class UserService {

    @Autowired
    UserRepository userRepository;

    @Autowired
    PasswordEncoder encoder;


    public User userRegistration(RegistrationDTO registrationDTO) {

        Optional<User> optionalUser = userRepository.findByUsername(registrationDTO.getUsername());

        if(!optionalUser.isEmpty()) {
            return null;
        }

        User user = new User();
        user.setUsername(registrationDTO.getUsername());
        user.setPassword(encoder.encode(registrationDTO.getPassword()));
        user.setEmail(registrationDTO.getEmail());
        user.setName(registrationDTO.getName());
        user.setSurname(registrationDTO.getSurname());
        user.setPhoneNumber(registrationDTO.getPhoneNumber());
        user.setGender(registrationDTO.getGender());
        user.setPublic(true);
        user.setBirthDate(registrationDTO.getBirthDate());
        user.setType("User");

        return userRepository.save(user);
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

        if(!encoder.matches(changePasswordDTO.getOldPassword(), user.getPassword())) {
            return null;
        }

        user.setPassword(encoder.encode(changePasswordDTO.getNewPassword()));

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
}
