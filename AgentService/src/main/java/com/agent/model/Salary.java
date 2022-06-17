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

	public Long getId() {
		return id;
	}

	public void setId(Long id) {
		this.id = id;
	}

	public String getUsername() {
		return username;
	}

	public void setUsername(String username) {
		this.username = username;
	}

	public Long getAmount() {
		return amount;
	}

	public void setAmount(Long amount) {
		this.amount = amount;
	}

	public JobOffer getOffer() {
		return offer;
	}

	public void setOffer(JobOffer offer) {
		this.offer = offer;
	}
}
