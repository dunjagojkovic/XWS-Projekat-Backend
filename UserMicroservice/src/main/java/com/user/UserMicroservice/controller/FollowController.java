package com.user.UserMicroservice.controller;


import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.user.UserMicroservice.dto.FollowDTO;
import com.user.UserMicroservice.service.FollowService;

@RestController
@RequestMapping("api/follow")
public class FollowController {

	@Autowired
    private FollowService followService;

	@PostMapping(path = "/follower") 
    public ResponseEntity<?> follow(@RequestBody FollowDTO dto) {
		followService.follow(dto);
    	return new ResponseEntity<>(HttpStatus.OK);
    }
	
	@PostMapping(path = "/accept") 
    public ResponseEntity<?> accept(@RequestBody FollowDTO dto) {
		followService.accept(dto);
    	return new ResponseEntity<>(HttpStatus.OK);
    }
	
	@GetMapping(value = "/following/{username}")
    public ResponseEntity<?> getFollowingUsers(@PathVariable String username) {

    	List<String> users = followService.getFollowing(username);
    	
        if(users.isEmpty()) {
            return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
        }

        return new ResponseEntity<>(users, HttpStatus.OK);
    }
	
}
