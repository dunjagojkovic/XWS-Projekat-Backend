package com.user.UserMicroservice.repository;

import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import com.user.UserMicroservice.model.Follow;

@Repository
public interface FollowRepository extends JpaRepository<Follow, Long> {

	Optional<Follow> findByFollower(String username);
	Optional<Follow> findByFollowing(String username);
    Optional<Follow> findById(Long id);
	
}
