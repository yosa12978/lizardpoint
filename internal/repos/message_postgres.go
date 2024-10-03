package repos

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/yosa12978/lizardpoint/internal/logging"
	"github.com/yosa12978/lizardpoint/internal/types"
)

type messagePostgres struct {
	db     *sql.DB
	logger logging.Logger
}

func NewMessagePostgres(db *sql.DB, logger logging.Logger) Message {
	return &messagePostgres{db: db, logger: logger}
}

var getMessagesByChannelIdSQL = `
	SELECT m.id, m.content, m.edited,
	m.created_at, m.updated_at, m.parent_id,
	p.account_id AS parent_account_id, 
	p.username AS parent_account_username,
	a.id AS account_id,
	a.username AS account_username,
	c.id AS channel_id,
	c.name AS channel_name FROM messages m
	INNER JOIN channels c ON m.channel_id=c.id
	INNER JOIN accounts a ON m.account_id=a.id
	LEFT JOIN (
		SELECT pm.id, pm.account_id, pa.username 
		FROM messages pm INNER JOIN accounts pa ON pa.id=pm.account_id
	) p ON m.parent_id=p.id
	WHERE m.channel_id=$1 AND m.created_at <= $2
	ORDER BY created_at DESC OFFSET $3 LIMIT $4;
`

func (m *messagePostgres) GetByChannelId(
	ctx context.Context,
	channelId uuid.UUID,
	page int, limit int,
) ([]types.Message, error) {

	messages := []types.Message{}
	rows, err := m.db.QueryContext(ctx,
		getMessagesByChannelIdSQL,
		channelId,
		time.Now().UTC(),
		(page-1)*limit,
		limit,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return messages, nil
		}
		return messages, types.NewErrInternalFailure(err)
	}
	for rows.Next() {
		message := types.Message{}
		rows.Scan(
			&message.Id,
			&message.Content,
			&message.Edited,
			&message.CreatedAt,
			&message.UpdatedAt,
			&message.ParentId,
			&message.ParentAccountId,
			&message.ParentAccountUsername,
			&message.AccountId,
			&message.AccountUsername,
			&message.ChannelId,
			&message.ChannelName,
		)
		messages = append(messages, message)
	}
	return messages, nil
}

var getMessageRepliesSQL = `
	WITH RECURSIVE tree AS (
		SELECT m.id, m.content, m.edited,
		m.created_at, m.updated_at, m.parent_id,
		p.account_id AS parent_account_id, 
		p.username AS parent_account_username,
		a.id AS account_id,
		a.username AS account_username,
		c.id AS channel_id,
		c.name AS channel_name FROM messages m
		INNER JOIN channels c ON m.channel_id=c.id
		INNER JOIN accounts a ON m.account_id=a.id
		LEFT JOIN (
			SELECT pm.id, pm.account_id, pa.username 
			FROM messages pm INNER JOIN accounts pa ON pa.id=pm.account_id
		) p ON m.parent_id=p.id
		WHERE m.id=$1

		UNION ALL 

		SELECT r.id, r.content, r.edited,
		r.created_at, r.updated_at, r.parent_id,
		p.account_id AS parent_account_id, 
		p.username AS parent_account_username,
		a.id AS account_id,
		a.username AS account_username,
		c.id AS channel_id,
		c.name AS channel_name FROM messages r
		INNER JOIN channels c ON r.channel_id=c.id
		INNER JOIN accounts a ON r.account_id=a.id
		LEFT JOIN (
			SELECT pm.id, pm.account_id, pa.username 
			FROM messages pm INNER JOIN accounts pa ON pa.id=pm.account_id
		) p ON r.parent_id=p.id
		INNER JOIN tree ON r.parent_id = tree.id
	) SELECT * FROM tree WHERE created_at <=$2 OFFSET $3 LIMIT $4;
`

func (m *messagePostgres) GetReplies(
	ctx context.Context,
	parentId uuid.UUID,
	page int, limit int,
) ([]types.Message, error) {
	messages := []types.Message{}
	rows, err := m.db.QueryContext(ctx,
		getMessageRepliesSQL,
		parentId,
		time.Now().UTC(),
		(page-1)*limit,
		limit,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return messages, types.ErrNotFound
		}
		return messages, types.NewErrInternalFailure(err)
	}
	for rows.Next() {
		message := types.Message{}
		rows.Scan(
			&message.Id,
			&message.Content,
			&message.Edited,
			&message.CreatedAt,
			&message.UpdatedAt,
			&message.ParentId,
			&message.ParentAccountId,
			&message.ParentAccountUsername,
			&message.AccountId,
			&message.AccountUsername,
			&message.ChannelId,
			&message.ChannelName,
		)
		messages = append(messages, message)
	}
	return messages[1:], nil // also exclude parent message
}

var getMessageByIdSQL = `
	SELECT m.id, m.content, m.edited,
	m.created_at, m.updated_at, m.parent_id,
	p.account_id AS parent_account_id, 
	p.username AS parent_account_username,
	a.id AS account_id,
	a.username AS account_username,
	c.id AS channel_id,
	c.name AS channel_name FROM messages m
	INNER JOIN channels c ON m.channel_id=c.id
	INNER JOIN accounts a ON m.account_id=a.id
	LEFT JOIN (
		SELECT pm.id, pm.account_id, pa.username 
		FROM messages pm INNER JOIN accounts pa ON pa.id=pm.account_id
	) p ON m.parent_id=p.id
	WHERE m.id=$1;
`

func (m *messagePostgres) GetById(ctx context.Context, id uuid.UUID) (*types.Message, error) {
	message := types.Message{}
	err := m.db.QueryRowContext(ctx, getMessageByIdSQL, id).
		Scan(
			&message.Id,
			&message.Content,
			&message.Edited,
			&message.CreatedAt,
			&message.UpdatedAt,
			&message.ParentId,
			&message.ParentAccountId,
			&message.ParentAccountUsername,
			&message.AccountId,
			&message.AccountUsername,
			&message.ChannelId,
			&message.ChannelName,
		)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, types.ErrNotFound
		}
		return nil, types.NewErrInternalFailure(err)
	}
	return &message, nil
}

var insertMessageSQL = `
	INSERT INTO messages (
		id, content, edited, 
		created_at, updated_at, account_id, 
		channel_id, parent_id,
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
`

func (m *messagePostgres) Create(ctx context.Context, message types.Message) error {
	_, err := m.db.ExecContext(ctx,
		insertMessageSQL,
		message.Id,
		message.Content,
		message.Edited,
		message.CreatedAt,
		message.UpdatedAt,
		message.AccountId,
		message.ChannelId,
		message.ParentId,
	)
	if err != nil {
		return types.NewErrInternalFailure(err)
	}
	return nil
}

var updateMessageSQL = `
	UPDATE messages SET content=$1, edited=true, updated_at=$2 WHERE id=$3; 
`

func (m *messagePostgres) Update(ctx context.Context, message types.Message) error {
	_, err := m.db.ExecContext(ctx,
		updateMessageSQL,
		message.Content,
		message.UpdatedAt,
		message.Id,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return types.NewErrNotFound(err)
		}
		return types.NewErrInternalFailure(err)
	}
	return nil
}

var deleteMessageSQL = `
	DELETE FROM messages WHERE id=$1;
`

func (m *messagePostgres) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := m.db.ExecContext(ctx, deleteMessageSQL, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return types.NewErrNotFound(err)
		}
		return types.NewErrInternalFailure(err)
	}
	return nil
}
