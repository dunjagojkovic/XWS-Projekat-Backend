package com.agent.service;

import com.agent.dto.CommentDTO;
import com.agent.dto.CompanyDTO;
import com.agent.dto.SurveyDTO;
import com.agent.model.Comment;
import com.agent.model.Company;
import com.agent.model.JobOffer;
import com.agent.model.Survey;
import com.agent.model.User;
import com.agent.repository.CommentRepository;
import com.agent.repository.CompanyRepository;
import com.agent.repository.SurveyRepository;
import com.agent.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

@Service
public class CompanyService {

    @Autowired
    CompanyRepository companyRepository;

    @Autowired
    UserRepository userRepository;

    @Autowired
    private UserService userService;
    
    @Autowired
    CommentRepository commentRepository;
    
    @Autowired
    SurveyRepository surveyRepository;

    public Company add(CompanyDTO companyDTO) {

        Company company = new Company();

        company.setName(companyDTO.getName());
        company.setEmail(companyDTO.getEmail());
        company.setAddress(companyDTO.getAddress());
        company.setCity(companyDTO.getCity());
        company.setState(companyDTO.getState());
        company.setContact(companyDTO.getContact());
        company.setDescription(companyDTO.getDescription());
        company.setOwner(userService.getCurrentUser());
        company.setStatus("Pending");

        return  companyRepository.save(company);
    }

    public List<Company> getAllCompaniesForApproving() {

        return companyRepository.findAllByStatus("Pending");
    }
    
    public List<Company> getApprovedCompanies() {
    	 return companyRepository.findAllByStatus("Approved");
    }

    public Company approveCompanyRegistration(CompanyDTO dto){

        Company companyToApprove = companyRepository.getById(dto.getId());

        companyToApprove.setStatus("Approved");
        companyToApprove.getOwner().setType("Company owner");
        userRepository.save(companyToApprove.getOwner());
        return  companyRepository.save(companyToApprove);
    }

    public Company declineCompanyRegistration(CompanyDTO companyDTO){

        Company companyToDecline = companyRepository.getById(companyDTO.getId());
        companyToDecline.setStatus("Declined");

        return companyRepository.save(companyToDecline);
    }


    public List<Company> getAllCompaniesForOwner(){
        return companyRepository.findAllByOwnerIdAndStatus(userService.getCurrentUser().getId(), "Approved");
    }

    public Company editCompanyInfo(CompanyDTO companyDTO) {
        Company company = companyRepository.getById(companyDTO.getId());

        if (companyDTO.getName() != null && !companyDTO.getName().equals("")){
            company.setName(companyDTO.getName());
        }
        if (companyDTO.getAddress() != null && !companyDTO.getAddress().equals("")){
            company.setAddress(companyDTO.getAddress());
        }
        if (companyDTO.getCity() != null && !companyDTO.getCity().equals("")){
            company.setCity(companyDTO.getCity());
        }
        if (companyDTO.getContact() != null && !companyDTO.getContact().equals("")){
            company.setContact(companyDTO.getContact());
        }
        if (companyDTO.getDescription() != null && !companyDTO.getDescription().equals("")){
            company.setDescription(companyDTO.getDescription());
        }
        if (companyDTO.getEmail() != null && !companyDTO.getEmail().equals("")){
            company.setEmail(companyDTO.getEmail());
        }
        if (companyDTO.getState() != null && !companyDTO.getState().equals("")){
            company.setState(companyDTO.getState());
        }

        return companyRepository.save(company);

    }

    public Company getCompanyInfo(Long id){
        return companyRepository.getById(id);
    }
    
    public Comment comment(CommentDTO commentDTO){

        Company company = companyRepository.getById(commentDTO.getCompanyId());

        Comment comment  = new Comment();
        comment.setUsername(commentDTO.getUsername());
        comment.setContent(commentDTO.getContent());
        comment.setCompany(company);

        return commentRepository.save(comment);
    }
    
    
    public List<Comment> getComments(Long id){
    	
    	List<Comment> companyComments = new ArrayList<>();
    	List<Comment> comments = commentRepository.findAll();
    
    	for(Comment comment: comments) {
    		if(comment.getCompany().getId() == id) {
    			companyComments.add(comment);
    		}
    	}
    	
    	return companyComments;
    	
    }
    
    public Survey survey(SurveyDTO surveyDTO){

        Company company = companyRepository.getById(surveyDTO.getCompanyId());

        Survey survey = new Survey();
        survey.setWorkEnvironment(surveyDTO.getWorkEnvironment());
        survey.setOpportunities(surveyDTO.getOpportunities());
        survey.setBenefits(surveyDTO.getBenefits());
        survey.setSalary(surveyDTO.getSalary());
        survey.setCommunication(surveyDTO.getCommunication());
        survey.setColleagues(surveyDTO.getColleagues());
        survey.setSupervision(surveyDTO.getSupervision());
        survey.setUsername(surveyDTO.getUsername());
        survey.setCompany(company);

        return surveyRepository.save(survey);
    }
    
    public List<Survey> getCompanySurveys(Long id){
    	
    	List<Survey> companySurveys = new ArrayList<>();
    	List<Survey> surveys = surveyRepository.findAll();
    
    	for(Survey survey: surveys) {
    		if(survey.getCompany().getId() == id) {
    			companySurveys.add(survey);
    		}
    	}
    	
    	return companySurveys;
    	
    }


}
