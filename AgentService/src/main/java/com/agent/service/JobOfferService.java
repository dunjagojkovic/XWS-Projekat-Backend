package com.agent.service;

import com.agent.dto.CommentDTO;
import com.agent.dto.CompanyDTO;
import com.agent.dto.CompanyOfferDTO;
import com.agent.dto.JobOfferDTO;
import com.agent.dto.SalaryDTO;
import com.agent.dto.SurveyDTO;
import com.agent.model.Comment;
import com.agent.model.Company;
import com.agent.model.JobOffer;
import com.agent.model.Salary;
import com.agent.model.Survey;
import com.agent.repository.CommentRepository;
import com.agent.repository.CompanyRepository;
import com.agent.repository.JobOfferRepository;
import com.agent.repository.SalaryRepository;
import com.agent.repository.SurveyRepository;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class JobOfferService {

    @Autowired
    JobOfferRepository jobOfferRepository;

    @Autowired
    CompanyRepository companyRepository;
    
    @Autowired
    SalaryRepository salaryRepository;
    
    @Autowired
    SurveyRepository surveyRepository;

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
    
 
    
    public Salary addSalary(SalaryDTO salaryDTO){

        JobOffer jobOffer = jobOfferRepository.getById(salaryDTO.getJobOfferId());

        Salary salary  = new Salary();
        salary.setUsername(salaryDTO.getUsername());
        salary.setAmount(salaryDTO.getAmount());
        salary.setOffer(jobOffer);

        return salaryRepository.save(salary);
    }
    
    public List<Salary> getJobOfferSalaries(Long id){
    	
    	List<Salary> jobOfferSalaries = new ArrayList<>();
    	List<Salary> salaries = salaryRepository.findAll();
    
    	for(Salary salary: salaries) {
    		if(salary.getOffer().getId() == id) {
    			jobOfferSalaries.add(salary);
    		}
    	}
    	
    	return jobOfferSalaries;
    	
    }
    
    
    public List<CompanyOfferDTO> getJobOffers(Long id){
    	
    	List<JobOffer> offers = new ArrayList<>();
    	offers = jobOfferRepository.findAllByCompanyId(id);
    	
    	List<CompanyOfferDTO> offersDTO = new ArrayList<>();
    	for(JobOffer offer: offers) {
    		CompanyOfferDTO dto = new CompanyOfferDTO();
    		dto.setId(offer.getId());
    		dto.setPosition(offer.getPosition());
    		dto.setSalary(offer.getSalary());
    		dto.setResponsibilities(offer.getResponsibilities());
    		dto.setRequirements(offer.getRequirements());
    		dto.setBenefit(offer.getBenefit());
    		dto.setCompany(offer.getCompany());
    		
    		offersDTO.add(dto);
    	}
    	return offersDTO;
    }


}
