/*
 * Revision History:
 *     Initial: 2018/09/18        Tong Yuehong
 */

package services

import (
	"database/sql"

	"github.com/TechCatsLab/comment/config"
	"github.com/TechCatsLab/comment/model/mysql"
)

type CommentService struct {
	db   *sql.DB
	SQLs []string
}

const (
	commentDB = iota
	commentTable
	commentInsert
	commentChangeStatus
	commentChangeContent
	commentListByTarget
	commentListByUser
)

func NewService(c *config.Config, db *sql.DB) *CommentService {
	cs := &CommentService{
		db: db,
		SQLs: []string{
			`CREATE DATABASE IF NOT EXISTS ` + c.CommentDB,
			`CREATE TABLE IF NOT EXISTS ` + c.CommentDB + `.` + c.CommentTable + ` (
				id 			INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
				target_id 	BIGINT UNSIGNED NOT NULL,
				user_id 	BIGINT UNSIGNED NOT NULL,
				content     TEXT NOT NULL,
				parent_id   BIGINT UNSIGNED NOT NULL,
				created_at 	DATETIME NOT NULL DEFAULT current_timestamp,
				status      TINYINT UNSIGNED NOT NULL DEFAULT 0,
				PRIMARY KEY (id),
				INDEX(target_id),
                INDEX(user_id)
			) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;`,
			`INSERT INTO ` + c.CommentDB  + `.` + c.CommentTable + `(target_id, content, user_id, parent_id, status) VALUES (?,?,?,?,?)`,
			`UPDATE ` + c.CommentDB  + `.` + c.CommentTable + ` SET status = ? WHERE id = ? LIMIT 1`,
			`UPDATE ` + c.CommentDB  + `.` + c.CommentTable + ` SET content = ? WHERE id = ? LIMIT 1`,
			`SELECT A.id, A.target_id, A.content, A.user_id, A.parent_id, A.created_at,` +
				`B.` + c.UserName + `,` + `B.` + c.UserAvatar +
				` FROM ` + c.CommentDB  + `.` + c.CommentTable + ` AS A,` + c.UserDB + `.` + c.UserTable + ` AS B ` +
				` WHERE A.user_id = B.` + c.UserID +
				` AND A.status = 0 AND A.target_id = ? LOCK IN SHARE MODE`,
			`SELECT id,target_id,content,user_id,parent_id,created_at FROM ` + c.CommentDB  + `.` + c.CommentTable + ` WHERE user_id = ? AND status = 0 LOCK IN SHARE MODE`,
		},
	}

	return cs
}

func (cs *CommentService) CreateDB() error {
	_, err := cs.db.Exec(cs.SQLs[commentDB])
	return err
}

func (cs *CommentService) CreateTable() error {
	_, err := cs.db.Exec(cs.SQLs[commentTable])
	return err
}

func (cs *CommentService) Insert(targetID, parentID, userID uint64, content string) (uint64, error) {
	return mysql.Insert(cs.db, cs.SQLs[commentInsert], targetID, parentID, userID, content)
}

func (cs *CommentService) ChangeStatus(commentID uint64, status uint8) error {
	return mysql.ChangeStatus(cs.db, cs.SQLs[commentChangeStatus], commentID, status)
}

func (cs *CommentService) ChangeContent(commentID uint64, content string) error {
	return mysql.ChangeContent(cs.db, cs.SQLs[commentChangeContent], commentID, content)
}

func (cs *CommentService) CommentsByTarget(targetId uint64) ([]*mysql.ListComment, error) {
	comments, err := mysql.CommentsByTargetID(cs.db, cs.SQLs[commentListByTarget], targetId)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (cs *CommentService) CommentsByUser(userId uint64) ([]*mysql.Comment, error) {
	comments, err := mysql.CommentsByUserID(cs.db, cs.SQLs[commentListByUser], userId)
	if err != nil {
		return nil, err
	}

	return comments, nil
}
