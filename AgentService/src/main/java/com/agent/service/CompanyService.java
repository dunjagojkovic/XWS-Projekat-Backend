package com.agent.service;

import com.agent.dto.CompanyDTO;
import com.agent.model.Company;
import com.agent.repository.CompanyRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class CompanyService {

    @Autowired
    CompanyRepository companyRepository;

    @Autowired
    private UserService userService;

    public Company add(CompanyDTO companyDTO) {

        Company company = new Company();

        company.setName(companyDTO.getName());
        company.setEmail(companyDTO.getEmail());
        company.setAddress(companyDTO.getAddress());
        company.setCity(companyDTO.getCity());
        company.setState(companyDTO.getState());
        company.setContact(companyDTO.getContact());
        company.setDescription(companyDTO.getDescription());
        company.setOwnerId(userService.getCurrentUser().getId());
        company.setStatus("Pending");

        return  companyRepository.save(company);
    }
}
