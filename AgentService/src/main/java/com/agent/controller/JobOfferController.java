package com.agent.controller;

import com.agent.dto.CommentDTO;
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
import org.springframework.web.bind.annotation.*;

@RestController
@CrossOrigin(origins = "http://localhost:4200")
@RequestMapping("api/jobs")
public class JobOfferController {

    @Autowired
    private JobOfferService jobOfferService;

    @PostMapping(path = "/addOffer")
    ResponseEntity<?> add(@RequestBody JobOfferDTO jobOfferDTO){
        JobOffer jobOffer = jobOfferService.addJobOffer(jobOfferDTO);
        return new ResponseEntity<>(jobOffer, HttpStatus.OK);
    }
    
    @PostMapping(path = "/comment")
    ResponseEntity<?> comment(@RequestBody CommentDTO commentDTO){
        Comment comment = jobOfferService.comment(commentDTO);
        return new ResponseEntity<>(comment, HttpStatus.OK);
    }
    
    @GetMapping(value = "/comments/{id}")
    public ResponseEntity<?> getJobOfferComments(@PathVariable Long id) {
        List<Comment> comments = jobOfferService.getJobOfferComments(id);
        return new ResponseEntity<>(comments, HttpStatus.OK);
    }
    
    @GetMapping(value = "/offers")
    public ResponseEntity<?> getJobOffers() {
        List<JobOffer> jobOffers = jobOfferService.getJobOffers();
        return new ResponseEntity<>(jobOffers, HttpStatus.OK);
    }
    
    @PostMapping(path = "/salary")
    ResponseEntity<?> addSalary(@RequestBody SalaryDTO salaryDTO){
        Salary salary = jobOfferService.addSalary(salaryDTO);
        return new ResponseEntity<>(salary, HttpStatus.OK);
    }
    
    @GetMapping(value = "/salaries/{id}")
    public ResponseEntity<?> getJobOfferSalaries(@PathVariable Long id) {
        List<Salary> salaries = jobOfferService.getJobOfferSalaries(id);
        return new ResponseEntity<>(salaries, HttpStatus.OK);
    }
    
    @PostMapping(path = "/survey")
    ResponseEntity<?> survey(@RequestBody SurveyDTO surveyDTO){
        Survey survey = jobOfferService.survey(surveyDTO);
        return new ResponseEntity<>(survey, HttpStatus.OK);
    }
    
    @GetMapping(value = "/surveys/{id}")
    public ResponseEntity<?> getJobOfferSurveys(@PathVariable Long id) {
        List<Survey> surveys = jobOfferService.getJobOfferSurveys(id);
        return new ResponseEntity<>(surveys, HttpStatus.OK);
    }
}
