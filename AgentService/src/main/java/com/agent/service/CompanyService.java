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
        company.setOwner(userService.getCurrentUser());
        company.setStatus("Pending");

        return  companyRepository.save(company);
    }

    public List<Company> getAllCompaniesForApproving() {

        return companyRepository.findAllByStatus("Pending");
    }

    public List<Company> getAllApprovedCompanies() {

        return companyRepository.findAllByStatus("Approved");
    }

    public Company approveCompanyRegistration(CompanyDTO dto){

        Company companyToApprove = companyRepository.getById(dto.getId());

        companyToApprove.setStatus("Approved");
        companyToApprove.getOwner().setType("Company owner");
        userRepository.save(companyToApprove.getOwner());
        return  companyRepository.save(companyToApprove);
    }

    public Company declineCompanyRegistration(CompanyDTO companyDTO){

        Company companyToDecline = companyRepository.getById(companyDTO.getId());
        companyToDecline.setStatus("Declined");

        return companyRepository.save(companyToDecline);
    }


    public List<Company> getAllCompaniesForOwner(){
        return companyRepository.findAllByOwnerId(userService.getCurrentUser().getId());
    }

    public Company editCompanyInfo(CompanyDTO companyDTO) {
        Company company = companyRepository.getById(companyDTO.getId());

        if (companyDTO.getName() != null && !companyDTO.getName().equals("")){
            company.setName(companyDTO.getName());
        }
        if (companyDTO.getAddress() != null && !companyDTO.getAddress().equals("")){
            company.setAddress(companyDTO.getAddress());
        }
        if (companyDTO.getCity() != null && !companyDTO.getCity().equals("")){
            company.setCity(companyDTO.getCity());
        }
        if (companyDTO.getContact() != null && !companyDTO.getContact().equals("")){
            company.setContact(companyDTO.getContact());
        }
        if (companyDTO.getDescription() != null && !companyDTO.getDescription().equals("")){
            company.setDescription(companyDTO.getDescription());
        }
        if (companyDTO.getEmail() != null && !companyDTO.getEmail().equals("")){
            company.setEmail(companyDTO.getEmail());
        }
        if (companyDTO.getState() != null && !companyDTO.getState().equals("")){
            company.setState(companyDTO.getState());
        }

        return companyRepository.save(company);

    }

    public Company getCompanyInfo(Long id){
        return companyRepository.getById(id);
    }


}
