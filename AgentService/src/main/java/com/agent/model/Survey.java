package com.agent.model;

import lombok.Getter;
import lombok.Setter;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

import com.fasterxml.jackson.annotation.JsonBackReference;

@Getter
@Setter
@Entity
@Table(name = "survey_table")
public class Survey {
	@Id
	@GeneratedValue(
			strategy = GenerationType.IDENTITY
	)
	private Long id;
	private String workEnvironment;
	private String opportunities;
	private String benefits;
	private String salary;
	private String supervision;
	private String communication;
	private String colleagues;
	private String username;
	
	@JsonBackReference
	@ManyToOne
	private JobOffer offer;

	public Long getId() {
		return id;
	}

	public void setId(Long id) {
		this.id = id;
	}

	public String getWorkEnvironment() {
		return workEnvironment;
	}

	public void setWorkEnvironment(String workEnvironment) {
		this.workEnvironment = workEnvironment;
	}

	public String getOpportunities() {
		return opportunities;
	}

	public void setOpportunities(String opportunities) {
		this.opportunities = opportunities;
	}

	public String getBenefits() {
		return benefits;
	}

	public void setBenefits(String benefits) {
		this.benefits = benefits;
	}

	public String getSalary() {
		return salary;
	}

	public void setSalary(String salary) {
		this.salary = salary;
	}

	public String getSupervision() {
		return supervision;
	}

	public void setSupervision(String supervision) {
		this.supervision = supervision;
	}

	public String getCommunication() {
		return communication;
	}

	public void setCommunication(String communication) {
		this.communication = communication;
	}

	public String getColleagues() {
		return colleagues;
	}

	public void setColleagues(String colleagues) {
		this.colleagues = colleagues;
	}

	public String getUsername() {
		return username;
	}

	public void setUsername(String username) {
		this.username = username;
	}

	public JobOffer getOffer() {
		return offer;
	}

	public void setOffer(JobOffer offer) {
		this.offer = offer;
	}
}
