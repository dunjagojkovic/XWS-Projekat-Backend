package com.agent.controller;

import com.agent.dto.JobOfferDTO;
import com.agent.model.JobOffer;
import com.agent.service.JobOfferService;
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
}
