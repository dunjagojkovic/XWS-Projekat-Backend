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
	private Company company;
	

}
