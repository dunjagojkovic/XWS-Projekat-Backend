package com.agent.dto.converters;

import com.agent.dto.CompanyDTO;
import com.agent.model.Company;

import java.util.ArrayList;
import java.util.List;

public class CompanyConverters {

    public static CompanyDTO modelToDTO(Company company) {

        CompanyDTO dto = new CompanyDTO();
        dto.setAddress(company.getAddress());
        dto.setCity(company.getCity());
        dto.setName(company.getName());
        dto.setEmail(company.getEmail());
        dto.setContact(company.getContact());
        dto.setDescription(company.getDescription());
        dto.setState(company.getState());
        dto.setOwnerId(company.getOwnerId());
        dto.setStatus(company.getStatus());
        dto.setId(company.getId());

        return dto;
    }

    public static List<CompanyDTO> modelsToDTOs(List<Company> companies) {

        List<CompanyDTO> result = new ArrayList<>();

        for(Company company: companies) {
            result.add(modelToDTO(company));
        }
        return result;
    }

}
