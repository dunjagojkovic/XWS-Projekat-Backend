package com.agent.service;

import com.agent.dto.JobOfferDTO;
import com.agent.model.Company;
import com.agent.model.JobOffer;
import com.agent.repository.CompanyRepository;
import com.agent.repository.JobOfferRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class JobOfferService {

    @Autowired
    JobOfferRepository jobOfferRepository;

    @Autowired
    CompanyRepository companyRepository;

    public JobOffer addJobOffer(JobOfferDTO jobOfferDTO){

        Company company = companyRepository.getById(jobOfferDTO.getCompanyId());

        JobOffer jobOffer = new JobOffer();
        jobOffer.setBenefit(jobOfferDTO.getBenefit());
        jobOffer.setCompany(company);
        jobOffer.setPosition(jobOfferDTO.getPosition());
        jobOffer.setRequirements(jobOfferDTO.getRequirements());
        jobOffer.setResponsibilities(jobOfferDTO.getResponsibilities());
        jobOffer.setSalary(jobOfferDTO.getSalary());

        return jobOfferRepository.save(jobOffer);
    }
}
