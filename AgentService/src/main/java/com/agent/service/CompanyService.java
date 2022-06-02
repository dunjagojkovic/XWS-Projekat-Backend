package com.agent.service;

import com.agent.dto.CompanyDTO;
import com.agent.model.Company;
import com.agent.model.User;
import com.agent.repository.CompanyRepository;
import com.agent.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

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

    public List<Company> getAllCompaniesForApproving() {

            return companyRepository.findAllByStatus("Pending");
    }

    public Company approveCompanyRegistration(CompanyDTO dto){

        Company companyToApprove = companyRepository.getById(dto.getId());
        Optional<User> potentialOwner = userRepository.findById(dto.getOwnerId());

        companyToApprove.setStatus("Approved");
        potentialOwner.get().setType("Company owner");
        userRepository.save(potentialOwner.get());


        return  companyRepository.save(companyToApprove);
    }

    public Company declineCompanyRegistration(CompanyDTO companyDTO){

        Company companyToDecline = companyRepository.getById(companyDTO.getId());
        companyToDecline.setStatus("Declined");

        return companyRepository.save(companyToDecline);
    }



}
