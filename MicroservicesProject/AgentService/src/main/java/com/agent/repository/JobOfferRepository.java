package com.agent.repository;

import com.agent.model.JobOffer;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository
public interface JobOfferRepository extends JpaRepository<JobOffer, Long> {

    Optional<JobOffer> findById(Long id);
    List<JobOffer> findAllByCompanyId(Long id);
}
