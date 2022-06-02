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
@CrossOrigin(origins = "*")
@RequestMapping("api/companies")
public class CompanyController {

    @Autowired
    CompanyService companyService;

    @PostMapping(consumes = "application/json", path = "/registerCompany")
    public ResponseEntity<?> add(@RequestBody CompanyDTO dto) {
        Company company  = companyService.add(dto);

        return new ResponseEntity<>(CompanyConverters.modelToDTO(company), HttpStatus.OK);
    }

    @GetMapping(path = "/allPendingCompanies")
    public ResponseEntity<?> getAll() {
        List<Company> companies = companyService.getAllCompaniesForApproving();
        return new ResponseEntity<>(CompanyConverters.modelsToDTOs(companies), HttpStatus.OK);
    }


    @PutMapping(path = "/approveCompanyRequest")
    public ResponseEntity<?> approveRequest(@RequestBody CompanyDTO dto){
       Company company =  companyService.approveCompanyRegistration(dto);
        return new ResponseEntity<>(company, HttpStatus.OK);
    }

    @PutMapping(path = "/declineCompanyRequest")
    public ResponseEntity<?> declineCompanyRequest( @RequestBody CompanyDTO companyDTO) {
        Company company = companyService.declineCompanyRegistration(companyDTO);
        return new ResponseEntity<>(company, HttpStatus.OK);
    }

    @GetMapping(path = "/myCompanies")
    public ResponseEntity<?> getMyCompanies() {
        List<Company> companies = companyService.getAllCompaniesForOwner();
        return new ResponseEntity<>(CompanyConverters.modelsToDTOs(companies), HttpStatus.OK);
    }

}
