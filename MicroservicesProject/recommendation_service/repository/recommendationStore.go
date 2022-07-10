package repository

import (
	"common/tracer"
	"context"
	"fmt"
	"recommendationS/model"
	"strings"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type RecommendationStore struct {
	recommendationDB *neo4j.Driver
}

func NewRecommendationStore(client *neo4j.Driver) RecommendationStoreI {
	return &RecommendationStore{
		recommendationDB: client,
	}
}

func (store *RecommendationStore) JobRecommendations(ctx context.Context, id string, experiences []*model.WorkExperience, skills []string, jobOffers []*model.JobOffer) ([]*model.JobsId, error) {

	span := tracer.StartSpanFromContext(ctx, "REPOSITORY JobRecommendations")
	defer span.Finish()

	session := (*store.recommendationDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	result, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {

		if !checkIfUserExist(ctx, id, transaction) {
			_, err := transaction.Run(
				"CREATE (new_user:USER{userID:$userID})",
				map[string]interface{}{"userID": id})

			if err != nil {
				return nil, err
			}

		}

		for _, s := range skills {
			if !checkIfSkillExist(ctx, s, transaction) {
				_, err := transaction.Run(
					"CREATE (new_skill:SKILL{name : $Name})",
					map[string]interface{}{"Name": s})

				if err != nil {
					return nil, err
				}

			}

			if !checkIfRelationshipExist(ctx, id, s, transaction) {
				fmt.Println("Veza ne postoji")
				result, err := transaction.Run(
					"MATCH (u:USER) WHERE u.userID=$uIDa "+
						"MATCH (s:SKILL) WHERE s.name=$name "+
						"CREATE (u)-[r:KNOWS]->(s) "+
						"RETURN u.userID",
					map[string]interface{}{"uIDa": id, "name": s})
				if err != nil {
					return nil, err
				}

				fmt.Println(result)
			}
		}

		for _, s := range experiences {
			if !checkIfExperienceExist(ctx, s.Description, transaction) {
				_, err := transaction.Run(
					"CREATE (new_exp:POSITION{description : $description}) ",
					map[string]interface{}{"description": s.Description})

				if err != nil {
					return nil, err
				}

			}

			if !checkIfExpRelationshipExist(ctx, id, s.Description, transaction) {
				result, err := transaction.Run(
					"MATCH (u:USER) WHERE u.userID=$uIDa "+
						"MATCH (s:POSITION) WHERE s.description=$description "+
						"CREATE (u)-[r:WORKED]->(s) "+
						"RETURN u.userID",
					map[string]interface{}{"uIDa": id, "description": s.Description})
				if err != nil {
					return nil, err
				}

				fmt.Println(result)
			}
		}

		for _, job := range jobOffers {

			if !jobOfferExist(ctx, job.Id.Hex(), transaction) {
				_, err := transaction.Run(
					"CREATE (new_job:JOB{position:$position, jobID:$jobID, description:$Description, preconditions: $preconditions})",
					map[string]interface{}{"jobID": job.Id.Hex(), "Description": job.Description, "preconditions": job.Precondition, "position": job.Position})

				if err != nil {
					return nil, err
				}
			}

			if !checkIfJobRelationshipExist(ctx, job.Id.Hex(), job.Precondition, transaction) {
				result, err := transaction.Run(
					"MATCH (j:JOB) WHERE j.jobID=$jobID "+
						"MATCH (s:SKILL) WHERE s.name=$name "+
						"CREATE (j)-[r:NEEDS]->(s) "+
						"RETURN j.jobID",
					map[string]interface{}{"jobID": job.Id.Hex(), "name": job.Precondition})
				if err != nil {
					return nil, err
				}

				fmt.Println(result)
			}

			if !checkIfJobPositionRelationshipExist(ctx, job.Id.Hex(), job.Position, transaction) {
				result, err := transaction.Run(
					"MATCH (j:JOB) WHERE j.jobID=$jobID "+
						"MATCH (s:POSITION) WHERE s.description=$position "+
						"CREATE (j)-[r:INCLUDES]->(s) "+
						"RETURN j.jobID",
					map[string]interface{}{"jobID": job.Id.Hex(), "position": job.Position})
				if err != nil {
					return nil, err
				}

				fmt.Println(result)
			}
		}

		var recommendation []*model.JobsId

		jobsRecommendations, err1 := getJobRecommendations(ctx, id, transaction)
		if err1 != nil {
			return recommendation, err1
		}

		for _, recommend := range jobsRecommendations {
			recommendation = append(recommendation, recommend)
		}

		return recommendation, err1
	})

	fmt.Println(err)

	return result.([]*model.JobsId), nil
}

func checkIfUserExist(ctx context.Context, userID string, transaction neo4j.Transaction) bool {
	result, _ := transaction.Run(
		"MATCH (existing_user:USER) WHERE existing_user.userID = $userID RETURN existing_user.userID",
		map[string]interface{}{"userID": userID})

	if result != nil && result.Next() && result.Record().Values[0] == userID {
		return true
	}
	return false
}

func checkIfSkillExist(ctx context.Context, skillName string, transaction neo4j.Transaction) bool {
	result, _ := transaction.Run(
		"MATCH (existing_skill:SKILL) WHERE toUpper(existing_skill.name) = $name RETURN toUpper(existing_skill.name)",
		map[string]interface{}{"name": strings.ToUpper(skillName)})

	if result != nil && result.Next() && result.Record().Values[0] == strings.ToUpper(skillName) {
		return true
	}
	return false
}

func checkIfExperienceExist(ctx context.Context, expName string, transaction neo4j.Transaction) bool {
	result, _ := transaction.Run(
		"MATCH (e:POSITION) WHERE toUpper(e.description) = $description RETURN toUpper(e.description)",
		map[string]interface{}{"description": strings.ToUpper(expName)})

	if result != nil && result.Next() && result.Record().Values[0] == strings.ToUpper(expName) {
		return true
	}
	return false
}

func checkIfRelationshipExist(ctx context.Context, userID, skillName string, transaction neo4j.Transaction) bool {
	result, _ := transaction.Run(
		"MATCH (u1:USER) WHERE u1.userID=$uIDa "+
			"MATCH (s:SKILL) WHERE s.name=$name "+
			"MATCH (u1)-[r:KNOWS]->(s) "+
			"RETURN r.date ",
		map[string]interface{}{"uIDa": userID, "name": skillName})

	if result != nil && result.Next() {
		fmt.Println("nadjena veza")
		return true

	}
	return false
}

func checkIfExpRelationshipExist(ctx context.Context, userID, expName string, transaction neo4j.Transaction) bool {
	result, _ := transaction.Run(
		"MATCH (u1:USER) WHERE u1.userID=$uIDa "+
			"MATCH (s:POSITION) WHERE s.description=$description "+
			"MATCH (u1)-[r:WORKED]->(s) "+
			"RETURN s.description",
		map[string]interface{}{"uIDa": userID, "description": expName})

	if result != nil && result.Next() {
		return true
	}
	return false
}

func jobOfferExist(ctx context.Context, jobId string, transaction neo4j.Transaction) bool {
	result, _ := transaction.Run(
		"MATCH (j:JOB) WHERE j.jobID = $id RETURN j.jobID",
		map[string]interface{}{"id": jobId})

	if result != nil && result.Next() && result.Record().Values[0] == jobId {
		return true
	}
	return false
}

func checkIfJobRelationshipExist(ctx context.Context, jobID, skillName string, transaction neo4j.Transaction) bool {
	result, _ := transaction.Run(
		"MATCH (j:JOB) WHERE j.jobID=$jobID "+
			"MATCH (s:SKILL) WHERE s.name=$name "+
			"MATCH (j)-[r:NEEDS]->(s) "+
			"RETURN r ",
		map[string]interface{}{"jobID": jobID, "name": skillName})

	if result != nil && result.Next() {
		return true
	}
	return false
}

func checkIfJobPositionRelationshipExist(ctx context.Context, jobID, position string, transaction neo4j.Transaction) bool {
	result, _ := transaction.Run(
		"MATCH (j:JOB) WHERE j.jobID=$jobID "+
			"MATCH (s:POSITION) WHERE s.description=$position "+
			"MATCH (j)-[r:INCLUDES]->(s) "+
			"RETURN r ",
		map[string]interface{}{"jobID": jobID, "position": position})

	if result != nil && result.Next() {
		return true
	}
	return false
}

func getJobRecommendations(ctx context.Context, userID string, transaction neo4j.Transaction) ([]*model.JobsId, error) {
	result, err := transaction.Run(
		"MATCH  (u1:USER)-[:KNOWS]->(u4:SKILL)<-[:NEEDS]-(u3:JOB) "+
			"WHERE u1.userID=$uID "+
			"RETURN distinct u3.jobID "+
			"LIMIT 20 "+
			"UNION "+
			"MATCH (u1:USER)-[:WORKED]->(u2:POSITION)<-[:INCLUDES]-(u3:JOB) "+
			"WHERE u1.userID=$uID "+
			"RETURN distinct u3.jobID "+
			" LIMIT 20",
		map[string]interface{}{"uID": userID})

	if err != nil {
		return nil, err
	}

	fmt.Println(result)

	var recommendation []*model.JobsId
	for result.Next() {
		recommendation = append(recommendation, &model.JobsId{Id: result.Record().Values[0].(string)})
	}
	return recommendation, nil
}
