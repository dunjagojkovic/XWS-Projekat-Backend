package com.agent.model;


import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

import com.fasterxml.jackson.annotation.JsonBackReference;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
@Entity
@Table(name = "salary_table")
public class Salary {
	
	 @Id
	 @GeneratedValue(
			 strategy = GenerationType.IDENTITY
	 )
	 private Long id;
	 private String username;
	 private Long amount;
	 
	 @JsonBackReference
	 @ManyToOne
	 private JobOffer offer;
	 

}
