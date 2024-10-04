// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package db

import (
	"context"
	"database/sql"
)

const createChatRoom = `-- name: CreateChatRoom :one
INSERT INTO chat_rooms (id, name, created_at, updated_at)
VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, name, created_at, updated_at, last_message_sent_at
`

type CreateChatRoomParams struct {
	ID   string
	Name string
}

type CreateChatRoomRow struct {
	ID                string
	Name              string
	CreatedAt         string
	UpdatedAt         string
	LastMessageSentAt sql.NullString
}

func (q *Queries) CreateChatRoom(ctx context.Context, arg CreateChatRoomParams) (CreateChatRoomRow, error) {
	row := q.db.QueryRowContext(ctx, createChatRoom, arg.ID, arg.Name)
	var i CreateChatRoomRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastMessageSentAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (name, session_id, created_at, updated_at)
VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, name, session_id, created_at, updated_at
`

type CreateUserParams struct {
	Name      string
	SessionID string
}

type CreateUserRow struct {
	ID        int64
	Name      string
	SessionID string
	CreatedAt string
	UpdatedAt string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Name, arg.SessionID)
	var i CreateUserRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.SessionID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getChatRoom = `-- name: GetChatRoom :one
SELECT
    id,
    name,
    last_message_sent_at,
    created_at,
    updated_at
FROM chat_rooms
WHERE chat_rooms.id = ?
`

func (q *Queries) GetChatRoom(ctx context.Context, id string) (ChatRoom, error) {
	row := q.db.QueryRowContext(ctx, getChatRoom, id)
	var i ChatRoom
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LastMessageSentAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getLastMessages = `-- name: GetLastMessages :many
SELECT
    messages.id,
    messages.chat_room_id,
    messages.user_id,
    users.name AS user_name,
    messages.message,
    messages.created_at,
    messages.updated_at
FROM messages
         JOIN users ON messages.user_id = users.id
WHERE messages.chat_room_id = ?
ORDER BY messages.created_at
LIMIT ?
`

type GetLastMessagesParams struct {
	ChatRoomID string
	Limit      int64
}

type GetLastMessagesRow struct {
	ID         int64
	ChatRoomID string
	UserID     int64
	UserName   string
	Message    string
	CreatedAt  string
	UpdatedAt  string
}

func (q *Queries) GetLastMessages(ctx context.Context, arg GetLastMessagesParams) ([]GetLastMessagesRow, error) {
	rows, err := q.db.QueryContext(ctx, getLastMessages, arg.ChatRoomID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetLastMessagesRow
	for rows.Next() {
		var i GetLastMessagesRow
		if err := rows.Scan(
			&i.ID,
			&i.ChatRoomID,
			&i.UserID,
			&i.UserName,
			&i.Message,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserBySessionId = `-- name: GetUserBySessionId :one
SELECT id, name, created_at, updated_at, session_id FROM users WHERE session_id = ?
`

func (q *Queries) GetUserBySessionId(ctx context.Context, sessionID string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserBySessionId, sessionID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.SessionID,
	)
	return i, err
}

const insertMessage = `-- name: InsertMessage :exec
INSERT INTO messages (chat_room_id, user_id, username, message, created_at, updated_at)
VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, chat_room_id, user_id, username, message, created_at, updated_at
`

type InsertMessageParams struct {
	ChatRoomID string
	UserID     int64
	Username   string
	Message    string
}

func (q *Queries) InsertMessage(ctx context.Context, arg InsertMessageParams) error {
	_, err := q.db.ExecContext(ctx, insertMessage,
		arg.ChatRoomID,
		arg.UserID,
		arg.Username,
		arg.Message,
	)
	return err
}

const updateChatRoomLastMessageSentAt = `-- name: UpdateChatRoomLastMessageSentAt :exec
UPDATE chat_rooms
SET last_message_sent_at = CURRENT_TIMESTAMP, updated_at = CURRENT_TIMESTAMP
WHERE id = ?
`

func (q *Queries) UpdateChatRoomLastMessageSentAt(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, updateChatRoomLastMessageSentAt, id)
	return err
}