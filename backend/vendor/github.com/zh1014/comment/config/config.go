/*
 * Revision History:
 *     Initial: 2018/09/17        Tong Yuehong
 */

package config

import (
	"database/sql"
)

type Config struct {
	UserDB     string
	UserTable  string
	UserID     string
	UserName   string
	UserAvatar string
	CommentDB  string
	CommentTable string
	RequestDomain string

	PreCommentCheck func(db *sql.DB, userID string, targetID int64) error
}
