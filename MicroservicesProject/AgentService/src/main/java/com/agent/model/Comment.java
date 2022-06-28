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
	 private Company company;
	

}
