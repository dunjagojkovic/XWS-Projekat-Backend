package com.agent.repository;

import com.agent.model.Company;
import com.agent.model.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository
public interface CompanyRepository extends JpaRepository<Company, Long> {

    List<Company> findAll();
    List<Company> findAllByStatus(String status);
    List<Company> findAllByOwnerId(Long id);

}
