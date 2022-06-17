package com.agent.dto;


import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class SalaryDTO {
	
	private Long jobOfferId;
	private String username;
	private Long amount;
	public Long getJobOfferId() {
		return jobOfferId;
	}
	public void setJobOfferId(Long jobOfferId) {
		this.jobOfferId = jobOfferId;
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
}
