package com.agent.dto;


import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class CommentDTO {

	private Long companyId;
	private String username;
	private String content;
}
