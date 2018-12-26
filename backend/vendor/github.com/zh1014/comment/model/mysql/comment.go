/*
 * Revision History:
 *     Initial: 2018/09/12        Tong Yuehong
 */

package mysql

import (
	"database/sql"
	"errors"
	"time"

)

const (
	statusActive = 0
)

var (
	errInvalidInsert       = errors.New("insert comment: insert affected 0 rows")
	errInvalidChangeStatus = errors.New("change status: affected 0 rows")
)

type Comment struct {
	Id       uint64
	TargetID uint64
	Content  string
	UserID   string
	ParentID uint64
	Created  time.Time
}

type ListComment struct {
	Id       uint64
	TargetID uint64
	Content  string
	UserID   string
	ParentID uint64
	Created  time.Time
	Name string
	Avatar string
}

func Insert(db *sql.DB, insert string, targetID, parentID uint64, userID, content string) (uint64, error) {
	result, err := db.Exec(insert, targetID, content, userID, parentID, statusActive)

	if err != nil {
		return 0, err
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		return 0, errInvalidInsert
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

func ChangeStatus(db *sql.DB, changeStatus string, commentID uint64, status uint8) error {
	result, err := db.Exec(changeStatus, status, commentID)
	if err != nil {
		return err
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		return errInvalidChangeStatus
	}

	return nil
}

func ChangeContent(db *sql.DB, changeContent string, commentID uint64, content string) error {
	result, err := db.Exec(changeContent, content, commentID)
	if err != nil {
		return err
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		return errInvalidChangeStatus
	}

	return nil
}

func CommentsByUserID(db *sql.DB, commentsByUsers,userId string) ([]*Comment, error) {
	var (
		id       uint64
		targetID uint64
		content  string
		userID   string
		parentID uint64
		created  time.Time

		comments []*Comment
	)

	rows, err := db.Query(commentsByUsers, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&id, &targetID, &content, &userID, &parentID, &created); err != nil {
			return nil, err
		}

		comment := &Comment{
			Id:       id,
			TargetID: targetID,
			UserID:   userID,
			Content:  content,
			ParentID: parentID,
			Created:  created,
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func  CommentsByTargetID(db *sql.DB, commentsByTarget string, targetId uint64) ([]*ListComment, error) {
	var (
		id       uint64
		targetID uint64
		content  string
		userID   string
		parentID uint64
		created  time.Time
		userName     string
		userAvatar   string

		comments []*ListComment
	)

	rows, err := db.Query(commentsByTarget, targetId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&id, &targetID, &content, &userID, &parentID, &created, &userName, &userAvatar); err != nil {
			return nil, err
		}

		listComment := &ListComment{
			Id:       id,
			TargetID: targetID,
			UserID:   userID,
			Content:  content,
			ParentID: parentID,
			Created:  created,
			Name: userName,
			Avatar:userAvatar,
		}

		comments = append(comments, listComment)


	}

	return comments, nil
}
