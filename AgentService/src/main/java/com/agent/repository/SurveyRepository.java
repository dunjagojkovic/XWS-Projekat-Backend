package com.agent.repository;

import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;

import com.agent.model.Survey;


public interface SurveyRepository extends JpaRepository<Survey, Long> {

	Optional<Survey> findById(Long id);
}
