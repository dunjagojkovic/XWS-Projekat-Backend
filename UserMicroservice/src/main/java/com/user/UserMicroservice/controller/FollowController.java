package com.user.UserMicroservice.controller;


import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
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
    @PreAuthorize("hasAuthority('User')")
    public ResponseEntity<?> follow(@RequestBody FollowDTO dto) {
		followService.follow(dto);
    	return new ResponseEntity<>(HttpStatus.OK);
    }
	
	@PostMapping(path = "/accept")
    @PreAuthorize("hasAuthority('User')")
    public ResponseEntity<?> accept(@RequestBody FollowDTO dto) {
		followService.accept(dto);
    	return new ResponseEntity<>(HttpStatus.OK);
    }
	
	@PostMapping(path = "/deny")
    @PreAuthorize("hasAuthority('User')")
    public ResponseEntity<?> deny(@RequestBody FollowDTO dto) {
		followService.deny(dto);
    	return new ResponseEntity<>(HttpStatus.OK);
    }
	
	@GetMapping(value = "/following/{username}")
    @PreAuthorize("hasAuthority('User')")
    public ResponseEntity<?> getFollowingUsers(@PathVariable String username) {

    	List<String> users = followService.getFollowing(username);
    	
        return new ResponseEntity<>(users, HttpStatus.OK);
    }
	
	@GetMapping(value = "/requested/{username}")
    @PreAuthorize("hasAuthority('User')")
    public ResponseEntity<?> getRequestedUsers(@PathVariable String username) {

    	List<String> users = followService.getRequests(username);
    	
        return new ResponseEntity<>(users, HttpStatus.OK);
    }
	
	@GetMapping(value = "/requests/{username}")
    @PreAuthorize("hasAuthority('User')")
    public ResponseEntity<?> getRequests(@PathVariable String username) {

    	List<String> users = followService.getSentRequests(username);
    
        return new ResponseEntity<>(users, HttpStatus.OK);
    }
}
