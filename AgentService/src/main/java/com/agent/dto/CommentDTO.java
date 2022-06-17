package com.agent.dto;


import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class CommentDTO {

	private Long jobOfferId;
	private String username;
	private String content;
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
	public String getContent() {
		return content;
	}
	public void setContent(String content) {
		this.content = content;
	}
}
