package com.agent.controller;

import com.agent.dto.CompanyDTO;
import com.agent.dto.converters.CompanyConverters;
import com.agent.model.Company;
import com.agent.service.CompanyService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@CrossOrigin(origins = "*")
@RequestMapping("api/companies")
public class CompanyController {

    @Autowired
    CompanyService companyService;

    @PostMapping(consumes = "application/json", path = "/registerCompany")
    @PreAuthorize("hasAuthority('registerCompany') and hasAuthority('Potential company owner')")
    public ResponseEntity<?> add(@RequestBody CompanyDTO dto) {
        Company company  = companyService.add(dto);

        return new ResponseEntity<>(CompanyConverters.modelToDTO(company), HttpStatus.OK);
    }

    @GetMapping(path = "/allPendingCompanies")
    @PreAuthorize("hasAuthority('getAllCompanyRequests') and hasAuthority('Admin')")
    public ResponseEntity<?> getAll() {
        List<Company> companies = companyService.getAllCompaniesForApproving();
        return new ResponseEntity<>(CompanyConverters.modelsToDTOs(companies), HttpStatus.OK);
    }


    @PutMapping(path = "/approveCompanyRequest")
    @PreAuthorize("hasAuthority('approveCompanyRequest') and hasAuthority('Admin')")
    public ResponseEntity<?> approveRequest(@RequestBody CompanyDTO dto){
       Company company =  companyService.approveCompanyRegistration(dto);
        return new ResponseEntity<>(company, HttpStatus.OK);
    }

    @PutMapping(path = "/declineCompanyRequest")
    @PreAuthorize("hasAuthority('declineCompanyRequest') and hasAuthority('Admin')")
    public ResponseEntity<?> declineCompanyRequest( @RequestBody CompanyDTO companyDTO) {
        Company company = companyService.declineCompanyRegistration(companyDTO);
        return new ResponseEntity<>(company, HttpStatus.OK);
    }

    @GetMapping(path = "/myCompanies")
    @PreAuthorize("hasAuthority('getMyCompanies') and hasAuthority('Company owner')")
    public ResponseEntity<?> getMyCompanies() {
        List<Company> companies = companyService.getAllCompaniesForOwner();
        return new ResponseEntity<>(CompanyConverters.modelsToDTOs(companies), HttpStatus.OK);
    }

    @PutMapping(path = "/editCompanyInfo")
    @PreAuthorize("hasAuthority('editCompanyInfo') and hasAuthority('Company owner')")
    public ResponseEntity<?> editCompanyInfo(@RequestBody CompanyDTO dto) {
        Company company = companyService.editCompanyInfo(dto);

        return new ResponseEntity<>(company, HttpStatus.OK);
    }


    @GetMapping(path = "/{id}")
    @PreAuthorize("hasAuthority('getCompanyInfo') and hasAuthority('Company owner')")
    public ResponseEntity<?> getCompany(@PathVariable Long id) {

        Company company = companyService.getCompanyInfo(id);

        if(company == null) {
            return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
        }

        return new ResponseEntity<>(company, HttpStatus.OK);
    }

}
