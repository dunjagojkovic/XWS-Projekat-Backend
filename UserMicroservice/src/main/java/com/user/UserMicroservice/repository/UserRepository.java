package com.user.UserMicroservice.repository;

import com.user.UserMicroservice.model.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository
public interface UserRepository extends JpaRepository<User, Long> {

    Optional<User> findByUsername(String username);
    Optional<User> findById(Long id);
    Optional<User> findByActivationCodeAndActivatedTrue(String code);
    Optional<User> findByPasswordResetCode(String code);
    Optional<User> findByActivationCode(String code);
    List<User> findAllByIsPublic(Boolean isPublic);
    

}
