package com.agent.controller;

import com.agent.dto.CommentDTO;
import com.agent.dto.CompanyOfferDTO;
import com.agent.dto.JobOfferDTO;
import com.agent.dto.SalaryDTO;
import com.agent.dto.SurveyDTO;
import com.agent.dto.converters.CompanyConverters;
import com.agent.model.JobOffer;
import com.agent.model.Salary;
import com.agent.model.Survey;
import com.agent.model.Comment;
import com.agent.model.Company;
import com.agent.service.JobOfferService;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.*;

@RestController
@CrossOrigin(origins = "*")
@RequestMapping("api/jobs")
public class JobOfferController {

    @Autowired
    private JobOfferService jobOfferService;

    @PostMapping(path = "/addOffer")
    @PreAuthorize("hasAuthority('addJobOffer') and hasAuthority('Company owner')")
    ResponseEntity<?> add(@RequestBody JobOfferDTO jobOfferDTO){
        JobOffer jobOffer = jobOfferService.addJobOffer(jobOfferDTO);
        return new ResponseEntity<>(jobOffer, HttpStatus.OK);
    }
    
    @PostMapping(path = "/comment")
    @PreAuthorize("hasAuthority('comment') and hasAuthority('User')")
    ResponseEntity<?> comment(@RequestBody CommentDTO commentDTO){
        Comment comment = jobOfferService.comment(commentDTO);
        return new ResponseEntity<>(comment, HttpStatus.OK);
    }
    
    @GetMapping(value = "/comments/{id}")
    @PreAuthorize("hasAuthority('getJobOfferComments') and hasAuthority('User')")
    public ResponseEntity<?> getJobOfferComments(@PathVariable Long id) {
        List<Comment> comments = jobOfferService.getJobOfferComments(id);
        return new ResponseEntity<>(comments, HttpStatus.OK);
    }
    
    @GetMapping(value = "/offers")
    @PreAuthorize("hasAuthority('getJobOffers') and hasAuthority('User')")
    public ResponseEntity<?> getJobOffers() {
        List<CompanyOfferDTO> jobOffers = jobOfferService.getJobOffers();
        return new ResponseEntity<>(jobOffers, HttpStatus.OK);
    }
    
    @PostMapping(path = "/salary")
    @PreAuthorize("hasAuthority('addSalary') and hasAuthority('Company owner')")
    ResponseEntity<?> addSalary(@RequestBody SalaryDTO salaryDTO){
        Salary salary = jobOfferService.addSalary(salaryDTO);
        return new ResponseEntity<>(salary, HttpStatus.OK);
    }
    
    @GetMapping(value = "/salaries/{id}")
    @PreAuthorize("hasAuthority('getJobOfferSalaries') and hasAuthority('User')")
    public ResponseEntity<?> getJobOfferSalaries(@PathVariable Long id) {
        List<Salary> salaries = jobOfferService.getJobOfferSalaries(id);
        return new ResponseEntity<>(salaries, HttpStatus.OK);
    }
    
    @PostMapping(path = "/survey")
    @PreAuthorize("hasAuthority('addSurvey') and hasAuthority('User')")
    ResponseEntity<?> survey(@RequestBody SurveyDTO surveyDTO){
        Survey survey = jobOfferService.survey(surveyDTO);
        return new ResponseEntity<>(survey, HttpStatus.OK);
    }
    
    @GetMapping(value = "/surveys/{id}")
    @PreAuthorize("hasAuthority('getJobOfferSurveys') and hasAuthority('User')")
    public ResponseEntity<?> getJobOfferSurveys(@PathVariable Long id) {
        List<Survey> surveys = jobOfferService.getJobOfferSurveys(id);
        return new ResponseEntity<>(surveys, HttpStatus.OK);
    }
}
