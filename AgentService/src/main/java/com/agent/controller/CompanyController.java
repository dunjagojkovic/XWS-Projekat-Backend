package com.agent.controller;

import com.agent.dto.CompanyDTO;
import com.agent.dto.converters.CompanyConverters;
import com.agent.model.Company;
import com.agent.service.CompanyService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@CrossOrigin(origins = "http://localhost:4200")
@RequestMapping("api/companies")
public class CompanyController {

    @Autowired
    CompanyService companyService;

    @PostMapping(consumes = "application/json", path = "/registerCompany")
    public ResponseEntity<?> add(@RequestBody CompanyDTO dto) {
        Company company  = companyService.add(dto);

        return new ResponseEntity<>(CompanyConverters.modelToDTO(company), HttpStatus.OK);
    }

    @GetMapping(path = "/allCompanies")
    public ResponseEntity<?> getAll() {
        List<Company> companies = companyService.getAll();
        return new ResponseEntity<>(CompanyConverters.modelsToDTOs(companies), HttpStatus.OK);
    }

}
