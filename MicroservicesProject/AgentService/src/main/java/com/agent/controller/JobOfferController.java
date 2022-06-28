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
    
    @GetMapping(value = "/offers/{id}")
    public ResponseEntity<?> getJobOffers(@PathVariable Long id) {
        List<CompanyOfferDTO> jobOffers = jobOfferService.getJobOffers(id);
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
    
}
