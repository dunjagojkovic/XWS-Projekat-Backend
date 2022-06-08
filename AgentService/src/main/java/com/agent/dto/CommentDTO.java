package com.agent.dto;


import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class CommentDTO {

	private Long jobOfferId;
	private String username;
	private String content;
}
