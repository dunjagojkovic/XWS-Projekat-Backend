package com.user.UserMicroservice.repository;

import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import com.user.UserMicroservice.model.Request;

@Repository
public interface RequestRepository extends JpaRepository<Request, Long> {

	Optional<Request> findByFollower(String username);
	Optional<Request> findByFollowing(String username);
    Optional<Request> findById(Long id);
}
