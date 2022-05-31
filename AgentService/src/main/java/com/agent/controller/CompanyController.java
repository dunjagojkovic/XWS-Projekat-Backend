package com.agent.controller;

import com.agent.dto.CompanyDTO;
import com.agent.dto.converters.CompanyConverters;
import com.agent.model.Company;
import com.agent.service.CompanyService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@CrossOrigin(origins = "http://localhost:4200")
@RequestMapping("api/companies")
public class CompanyController {

    @Autowired
    CompanyService companyService;

    @PostMapping()
    public ResponseEntity<?> add(@RequestBody CompanyDTO dto) {
        Company company  = companyService.add(dto);

        return new ResponseEntity<>(CompanyConverters.modelToDTO(company), HttpStatus.OK);
    }

}
