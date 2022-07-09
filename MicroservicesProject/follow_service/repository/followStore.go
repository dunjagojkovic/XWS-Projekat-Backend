package repository

import (
	"common/tracer"
	"context"
	"fmt"
	"followS/model"
	"io"
	"log"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
)

type FollowStore struct {
	driver       neo4j.Driver
	databaseName string
}

func NewFollowStore(driver *neo4j.Driver, dbName string) FollowStoreI {
	return &FollowStore{
		driver:       *driver,
		databaseName: dbName,
	}
}

func (store *FollowStore) Follows(ctx context.Context, id string) ([]*model.User, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	span := tracer.StartSpanFromContext(ctx, "REPOSITORY Follows")
	defer span.Finish()

	followers, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
			"MATCH (:User {userId:$id})-[follow:FOLLOWING]->(followed:User) RETURN followed, follow",
			map[string]interface{}{"id": id})
		if err != nil {
			return nil, err
		}
		results := []*model.User{}
		for records.Next() {
			record := records.Record()
			id, _ := record.Get("followed")
			relationship, _ := record.Get("follow")
			user := model.User{
				Id:           id.(dbtype.Node).Props["userId"].(string),
				TimeOfFollow: time.Time(relationship.(dbtype.Relationship).Props["timeStarted"].(dbtype.LocalDateTime)),
			}
			results = append(results, &user)
		}
		return results, nil
	})
	if err != nil {
		return nil, err
	}
	return followers.([]*model.User), nil
}

func (store *FollowStore) Followers(ctx context.Context, id string) ([]*model.User, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	span := tracer.StartSpanFromContext(ctx, "REPOSITORY Followers")
	defer span.Finish()

	followers, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
			"MATCH (:User {userId:$id})<-[follow:FOLLOWING]-(follower:User) RETURN follower, follow",
			map[string]interface{}{"id": id})
		if err != nil {
			return nil, err
		}
		results := []*model.User{}
		for records.Next() {
			record := records.Record()
			id, _ := record.Get("follower")
			relationship, _ := record.Get("follow")
			user := model.User{
				Id:           id.(dbtype.Node).Props["userId"].(string),
				TimeOfFollow: time.Time(relationship.(dbtype.Relationship).Props["timeStarted"].(dbtype.LocalDateTime)),
			}
			results = append(results, &user)
		}
		return results, nil
	})
	if err != nil {
		return nil, err
	}
	return followers.([]*model.User), nil
}

func (store *FollowStore) FollowRequests(ctx context.Context, id string) ([]*model.User, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	followers, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
			"MATCH (:User {userId:$id})-[follow:REQUESTING_FOLLOW]->(followed:User) RETURN followed, follow",
			map[string]interface{}{"id": id})
		if err != nil {
			return nil, err
		}
		results := []*model.User{}
		for records.Next() {
			record := records.Record()
			id, _ := record.Get("followed")
			relationship, _ := record.Get("follow")
			user := model.User{
				Id:           id.(dbtype.Node).Props["userId"].(string),
				TimeOfFollow: time.Time(relationship.(dbtype.Relationship).Props["timeSent"].(dbtype.LocalDateTime)),
			}
			results = append(results, &user)
		}
		return results, nil
	})
	if err != nil {
		return nil, err
	}
	return followers.([]*model.User), nil
}

func (store *FollowStore) FollowerRequests(ctx context.Context, id string) ([]*model.User, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	span := tracer.StartSpanFromContext(ctx, "REPOSITORY FollowerRequests")
	defer span.Finish()

	followers, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
			"MATCH (:User {userId:$id})<-[follow:REQUESTING_FOLLOW]-(follower:User) RETURN follower, follow",
			map[string]interface{}{"id": id})
		if err != nil {
			return nil, err
		}
		results := []*model.User{}
		for records.Next() {
			record := records.Record()
			id, _ := record.Get("follower")
			relationship, _ := record.Get("follow")
			user := model.User{
				Id:           id.(dbtype.Node).Props["userId"].(string),
				TimeOfFollow: time.Time(relationship.(dbtype.Relationship).Props["timeSent"].(dbtype.LocalDateTime)),
			}
			results = append(results, &user)
		}
		return results, nil
	})
	if err != nil {
		return nil, err
	}
	return followers.([]*model.User), nil
}
func (store *FollowStore) Relationship(ctx context.Context, followerId string, followedId string) (string, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	span := tracer.StartSpanFromContext(ctx, "REPOSITORY Relationship")
	defer span.Finish()

	relationship, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
			"MATCH (:User {userId:$followedId})<-[relationship]-(:User {userId:$followerId}) RETURN type(relationship)",
			map[string]interface{}{"followedId": followedId, "followerId": followerId})
		if err != nil {
			return nil, err
		}
		for records.Next() {
			record := records.Record()
			relationship := record.Values[0]
			results := relationship
			return results, nil
		}
		return "NO RELATIONSHIP", nil
	})
	if err != nil {
		return "", err
	}
	return relationship.(string), nil
}

func (store *FollowStore) Follow(ctx context.Context, followerId string, followedId string) (string, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	span := tracer.StartSpanFromContext(ctx, "REPOSITORY Follow")
	defer span.Finish()

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		currentTime := neo4j.LocalDateTime(time.Now())
		result, err := tx.Run(
			"MERGE (followed:User {userId: $followedId}) "+
				"ON CREATE SET followed.userId = $followedId "+
				"MERGE (follower:User {userId: $followerId}) "+
				"ON CREATE SET follower.userId = $followerId "+
				"MERGE (followed) <- [fol:FOLLOWING] - (follower) "+
				"ON CREATE SET fol.timeStarted = $timeStarted",
			map[string]interface{}{"followedId": followedId, "followerId": followerId,
				"timeStarted": currentTime})
		if err != nil {
			return nil, err
		}
		return result.Consume()
	})
	if err != nil {
		return "Failed to follow: " + followerId + " -> " + followedId, err
	}
	return session.LastBookmark(), nil
}

func (store *FollowStore) FollowRequest(ctx context.Context, followerId string, followedId string) (string, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	span := tracer.StartSpanFromContext(ctx, "REPOSITORY FollowRequest")
	defer span.Finish()

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		currentTime := neo4j.LocalDateTime(time.Now())
		result, err := tx.Run(
			"MERGE (followed:User {userId: $followedId}) "+
				"ON CREATE SET followed.userId = $followedId "+
				"MERGE (follower:User {userId: $followerId}) "+
				"ON CREATE SET follower.userId = $followerId "+
				"MERGE (followed) <- [req:REQUESTING_FOLLOW] - (follower) "+
				"ON CREATE SET req.timeSent = $timeSent",
			map[string]interface{}{"followedId": followedId, "followerId": followerId,
				"timeSent": currentTime})
		if err != nil {
			return nil, err
		}
		return result.Consume()
	})
	if err != nil {
		return "Failed to create follow request: " + followerId + " -> " + followedId, err
	}
	return session.LastBookmark(), nil
}

func (store *FollowStore) AcceptFollow(ctx context.Context, followerId string, followedId string) (string, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	span := tracer.StartSpanFromContext(ctx, "REPOSITORY AcceptFollow")
	defer span.Finish()

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		currentTime := neo4j.LocalDateTime(time.Now())
		result, err := tx.Run(
			"MATCH (followed:User {userId: $followedId}) "+
				"MATCH (follower:User {userId: $followerId}) "+
				"MATCH (followed) <- [followRequest:REQUESTING_FOLLOW] - (follower) "+
				"MERGE (followed) <- [fol:FOLLOWING] - (follower) "+
				"ON CREATE SET fol.timeStarted = $timeStarted "+
				"DELETE followRequest",
			map[string]interface{}{"followedId": followedId, "followerId": followerId,
				"timeStarted": currentTime})
		if err != nil {
			return nil, err
		}
		return result.Consume()
	})
	if err != nil {
		return "Failed to create follow request: " + followerId + " -> " + followedId, err
	}
	return session.LastBookmark(), nil
}

func (store *FollowStore) Unfollow(ctx context.Context, followerId string, followedId string) (string, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	span := tracer.StartSpanFromContext(ctx, "REPOSITORY Unfollow")
	defer span.Finish()

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(
			"MATCH (followed:User {userId: $followedId}) "+
				"MATCH (follower:User {userId: $followerId}) "+
				"MATCH (followed) <- [follow:FOLLOWING] - (follower) "+
				"DELETE follow",
			map[string]interface{}{"followedId": followedId, "followerId": followerId})
		if err != nil {
			return nil, err
		}
		return result.Consume()
	})
	if err != nil {
		return "Failed to follow: " + followerId + " -> " + followedId, err
	}
	return session.LastBookmark(), nil
}

func (store *FollowStore) FollowRequestRemove(ctx context.Context, followerId string, followedId string) (string, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	span := tracer.StartSpanFromContext(ctx, "REPOSITORY FollowRequestRemove")
	defer span.Finish()

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(
			"MATCH (followed:User {userId: $followedId}) "+
				"MATCH (follower:User {userId: $followerId}) "+
				"MATCH (followed) <- [followRequest:REQUESTING_FOLLOW] - (follower) "+
				"DELETE followRequest",
			map[string]interface{}{"followedId": followedId, "followerId": followerId})
		if err != nil {
			return nil, err
		}
		return result.Consume()
	})
	if err != nil {
		return "Failed to create follow request: " + followerId + " -> " + followedId, err
	}
	return session.LastBookmark(), nil
}

func (store *FollowStore) Recommended(ctx context.Context, id string) ([]*model.User, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	span := tracer.StartSpanFromContext(ctx, "REPOSITORY Recommended")
	defer span.Finish()

	followers, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
			"MATCH r = (u1:User{userId:$id}) - [f1:FOLLOWING] -> (u2:User) - [f2:FOLLOWING] -> (u3:User) "+
				"WHERE NOT exists((u1) - [:FOLLOWING] -> (u3))"+
				"RETURN u3, count(f2) as fol_num "+
				"ORDER BY fol_num DESC "+
				"LIMIT 15",
			map[string]interface{}{"id": id})
		if err != nil {
			return nil, err
		}
		results := []*model.User{}
		for records.Next() {
			record := records.Record()
			id, _ := record.Get("u3")
			user := model.User{
				Id: id.(dbtype.Node).Props["userId"].(string),
			}
			results = append(results, &user)
		}
		return results, nil
	})
	if err != nil {
		return nil, err
	}
	return followers.([]*model.User), nil
}

func unsafeClose(closeable io.Closer) {
	if err := closeable.Close(); err != nil {
		log.Fatal(fmt.Errorf("could not close resource: %w", err))
	}
}
