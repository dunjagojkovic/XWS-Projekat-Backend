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
        user.setPublic(false);
        user.setBirthDate(registrationDTO.getBirthDate());

        return userRepository.save(user);
    }

    public User getCurrentUser() {

        String username = SecurityUtils.getCurrentUserLogin().get();
        return userRepository.findByUsername(username).get();
    }

    public User edit(UserDTO userDTO) {

        Optional<User> optionalUser = userRepository.findById(getCurrentUser().getId());

            optionalUser.get().setName(userDTO.getName());
            optionalUser.get().setEmail(userDTO.getEmail());
            optionalUser.get().setPhoneNumber(userDTO.getPhoneNumber());
            optionalUser.get().setSurname(userDTO.getSurname());
            optionalUser.get().setPassword(userDTO.getPassword());
            optionalUser.get().setBiography(userDTO.getBiography());
            optionalUser.get().setBirthDate(userDTO.getBirthDate());
            optionalUser.get().setGender(userDTO.getGender());
            optionalUser.get().setEducation(userDTO.getEducation());
            optionalUser.get().setHobby(userDTO.getHobby());
            optionalUser.get().setUsername(userDTO.getUsername());
            optionalUser.get().setWorkExperience(userDTO.getWorkExperience());
            optionalUser.get().setInterest(userDTO.getInterest());
            optionalUser.get().setPublic(userDTO.getPublic());

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
}
