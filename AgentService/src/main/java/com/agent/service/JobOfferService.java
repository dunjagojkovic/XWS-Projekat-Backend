package com.agent.service;

import com.agent.dto.CommentDTO;
import com.agent.dto.CompanyDTO;
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
    CommentRepository commentRepository;
    
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
    
    public Comment comment(CommentDTO commentDTO){

        JobOffer jobOffer = jobOfferRepository.getById(commentDTO.getJobOfferId());

        Comment comment  = new Comment();
        comment.setUsername(commentDTO.getUsername());
        comment.setContent(commentDTO.getContent());
        comment.setOffer(jobOffer);

        return commentRepository.save(comment);
    }
    
    
    public List<Comment> getJobOfferComments(Long id){
    	
    	List<Comment> jobOfferComments = new ArrayList<>();
    	List<Comment> comments = commentRepository.findAll();
    
    	for(Comment comment: comments) {
    		if(comment.getOffer().getId() == id) {
    			jobOfferComments.add(comment);
    		}
    	}
    	
    	return jobOfferComments;
    	
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
    
    public Survey survey(SurveyDTO surveyDTO){

        JobOffer jobOffer = jobOfferRepository.getById(surveyDTO.getJobOfferId());

        Survey survey = new Survey();
        survey.setWorkEnvironment(surveyDTO.getWorkEnvironment());
        survey.setOpportunities(surveyDTO.getOpportunities());
        survey.setBenefits(surveyDTO.getBenefits());
        survey.setSalary(surveyDTO.getSalary());
        survey.setCommunication(surveyDTO.getCommunication());
        survey.setColleagues(surveyDTO.getColleagues());
        survey.setSupervision(surveyDTO.getSupervision());
        survey.setUsername(surveyDTO.getUsername());
        survey.setOffer(jobOffer);

        return surveyRepository.save(survey);
    }
    
    public List<Survey> getJobOfferSurveys(Long id){
    	
    	List<Survey> jobOfferSurveys = new ArrayList<>();
    	List<Survey> surveys = surveyRepository.findAll();
    
    	for(Survey survey: surveys) {
    		if(survey.getOffer().getId() == id) {
    			jobOfferSurveys.add(survey);
    		}
    	}
    	
    	return jobOfferSurveys;
    	
    }


}
