package com.agent.model;


import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;

import com.fasterxml.jackson.annotation.JsonBackReference;


@Getter
@Setter
@Entity
@Table(name = "comment_table")
public class Comment {
	
	@Id
    @GeneratedValue(
            strategy = GenerationType.IDENTITY
    )
    private Long id;
	private String username;
	private String content;
	
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

	public String getContent() {
		return content;
	}

	public void setContent(String content) {
		this.content = content;
	}

	public JobOffer getOffer() {
		return offer;
	}

	public void setOffer(JobOffer offer) {
		this.offer = offer;
	}
}
