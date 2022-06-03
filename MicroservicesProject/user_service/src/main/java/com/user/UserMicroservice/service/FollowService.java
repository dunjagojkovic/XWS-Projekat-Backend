package com.user.UserMicroservice.service;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.user.UserMicroservice.dto.FollowDTO;
import com.user.UserMicroservice.model.Follow;
import com.user.UserMicroservice.model.Request;
import com.user.UserMicroservice.model.User;
import com.user.UserMicroservice.repository.FollowRepository;
import com.user.UserMicroservice.repository.RequestRepository;
import com.user.UserMicroservice.repository.UserRepository;

@Service
public class FollowService {

	@Autowired
	FollowRepository followRepository;
	
	@Autowired
	RequestRepository requestRepository;
	
	@Autowired
	UserRepository userRepository;
	
	public void follow(FollowDTO followDTO) {
		
		Optional<User> user = userRepository.findByUsername(followDTO.getFollowing());
		if(user.get().getIsPublic() == true) {
			Follow f = new Follow();
			f.setFollower(followDTO.getFollower());
			f.setFollowing(followDTO.getFollowing());
			
			followRepository.save(f);
		} else {
			Request r = new Request();
			r.setFollower(followDTO.getFollower());
			r.setFollowing(followDTO.getFollowing());
			
			requestRepository.save(r);
		}
		
	}
	
	public void accept(FollowDTO followDTO) {
		
		Follow f = new Follow();
		f.setFollower(followDTO.getFollower());
		f.setFollowing(followDTO.getFollowing());
		
		followRepository.save(f);

		List<Request> requests = requestRepository.findAll();
		
		for(Request r : requests) {
			if(r.getFollower().equals(followDTO.getFollower()) && r.getFollowing().equals(followDTO.getFollowing())) {
				requestRepository.deleteById(r.getId());
			}	
		}	
	}
	
	public void deny(FollowDTO followDTO) {
		
		List<Request> requests = requestRepository.findAll();
		
		for(Request r : requests) {
			if(r.getFollower().equals(followDTO.getFollower()) && r.getFollowing().equals(followDTO.getFollowing())) {
				requestRepository.deleteById(r.getId());
			}	
		}	
	}
	
	public List<String> getRequests(String username) {
		
		List<Request> requests = requestRepository.findAll();
		List<String> requested = new ArrayList<>();
		
		for(Request request: requests) {
			
			if(request.getFollower().equals(username)) {
				requested.add(request.getFollowing());
			}
			
		}
		
		return requested;
	}
	
	public List<String> getSentRequests(String username) {
		
		List<Request> requests = requestRepository.findAll();
		List<String> sentRequests = new ArrayList<>();
		
		for(Request request: requests) {
			
			if(request.getFollowing().equals(username)) {
				sentRequests.add(request.getFollower());
			}
			
		}
		
		return sentRequests;
	}
	
	public List<String> getFollowing(String username) {
		
		List<Follow> follows = followRepository.findAll();
		List<String> following = new ArrayList<>();
		
		for(Follow follow: follows) {
			
			if(follow.getFollower().equals(username)) {
				following.add(follow.getFollowing());
			}
			
		}
		
		return following;
	}
	
}
