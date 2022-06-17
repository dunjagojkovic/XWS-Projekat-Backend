package com.agent.dto;


import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class SurveyDTO {
	
	private Long jobOfferId;
	private String workEnvironment;
	private String opportunities;
	private String benefits;
	private String salary;
	private String supervision;
	private String communication;
	private String colleagues;
	private String username;
	public Long getJobOfferId() {
		return jobOfferId;
	}
	public void setJobOfferId(Long jobOfferId) {
		this.jobOfferId = jobOfferId;
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
}
