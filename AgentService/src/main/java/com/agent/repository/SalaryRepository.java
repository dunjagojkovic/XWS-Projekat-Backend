package com.agent.repository;

import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;


import com.agent.model.Salary;

public interface SalaryRepository extends JpaRepository<Salary, Long> {

	Optional<Salary> findById(Long id);
}
