package repos

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/yosa12978/lizardpoint/internal/logging"
	"github.com/yosa12978/lizardpoint/internal/types"
)

type channelPostgres struct {
	db     *sql.DB
	logger logging.Logger
}

func NewChannelPostgres(db *sql.DB, logger logging.Logger) ChannelRepo {
	return &channelPostgres{
		db:     db,
		logger: logger,
	}
}

var getAllChannelsSQL = `
	SELECT c.id, c.name, 
	ARRAY(SELECT rp.role_name FROM read_permissions rp WHERE rp.channel_id=c.id) AS read_perms, 
	ARRAY(SELECT wp.role_name FROM write_permissions wp WHERE wp.channel_id=c.id) AS write_perms 
	FROM channels c;
`

func (c *channelPostgres) GetAll(ctx context.Context) ([]types.Channel, error) {
	channels := []types.Channel{}
	rows, err := c.db.QueryContext(ctx, getAllChannelsSQL)
	if err != nil {
		return channels, types.NewErrInternalFailure(err)
	}
	for rows.Next() {
		channel := types.Channel{}
		rp := []string{}
		wp := []string{}
		rows.Scan(
			&channel.Id,
			&channel.Name,
			(*pq.StringArray)(&rp),
			(*pq.StringArray)(&wp),
		)
		for _, v := range rp {
			channel.ReadPermissions =
				append(channel.ReadPermissions, types.Role{Name: v})
		}
		for _, v := range wp {
			channel.WritePermissions =
				append(channel.WritePermissions, types.Role{Name: v})
		}
		channels = append(channels, channel)
	}
	return channels, nil
}

var getChannelByIdSQL = `
	SELECT c.id, c.name, 
	ARRAY(SELECT rp.role_name FROM read_permissions rp WHERE rp.channel_id=$1) AS read_perms, 
	ARRAY(SELECT wp.role_name FROM write_permissions wp WHERE wp.channel_id=$1) AS write_perms 
	FROM channels c WHERE c.id=$1 ;
`

func (c *channelPostgres) GetById(ctx context.Context, id uuid.UUID) (*types.Channel, error) {
	channel := types.Channel{}
	rp := []string{}
	wp := []string{}
	err := c.db.QueryRowContext(ctx, getChannelByIdSQL).Scan(
		&channel.Id,
		&channel.Name,
		(*pq.StringArray)(&rp),
		(*pq.StringArray)(&wp),
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, types.NewErrNotFound(err)
		}
		return nil, types.NewErrInternalFailure(err)
	}
	for _, v := range rp {
		channel.ReadPermissions =
			append(channel.ReadPermissions, types.Role{Name: v})
	}
	for _, v := range wp {
		channel.WritePermissions =
			append(channel.WritePermissions, types.Role{Name: v})
	}
	return &channel, nil
}

var insertChannelSQL = `
	INSERT INTO channels (id, name) VALUES ($1, $2);
`

func (c *channelPostgres) Create(ctx context.Context, channel types.Channel) error {
	_, err := c.db.ExecContext(ctx, insertChannelSQL, channel.Id, channel.Name)
	if err != nil {
		return types.NewErrInternalFailure(err)
	}
	return nil
}

var updateChannelSQL = `
	UPDATE channels SET name=$1 WHERE id=$2;
`

func (c *channelPostgres) Update(ctx context.Context, channel types.Channel) error {
	_, err := c.db.ExecContext(ctx, updateChannelSQL, channel.Name, channel.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return types.NewErrNotFound(err)
		}
		return types.NewErrInternalFailure(err)
	}
	return nil
}

var deleteChannelSQL = `
	DELETE FROM channels WHERE id=$1;
`

func (c *channelPostgres) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := c.db.ExecContext(ctx, deleteChannelSQL, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return types.NewErrNotFound(err)
		}
		return types.NewErrInternalFailure(err)
	}
	return nil
}

var addWritePermSQL = `
	INSERT INTO write_permissions (channel_id, role_name) VALUES ($1, $2);
`

func (c *channelPostgres) AddWritePermission(ctx context.Context, channelId uuid.UUID, role string) error {
	_, err := c.db.ExecContext(ctx, addWritePermSQL, channelId, role)
	if err != nil {
		return types.NewErrInternalFailure(err)
	}
	return nil
}

var addReadPermSQL = `
	INSERT INTO read_permissions (channel_id, role_name) VALUES ($1, $2);
`

func (c *channelPostgres) AddReadPermission(ctx context.Context, channelId uuid.UUID, role string) error {
	_, err := c.db.ExecContext(ctx, addReadPermSQL, channelId, role)
	if err != nil {
		return types.NewErrInternalFailure(err)
	}
	return nil
}

var removeWritePermSQL = `
	DELETE FROM write_permissions WHERE channel_id=$1 AND role_name=$2;
`

func (c *channelPostgres) RemoveWritePermission(ctx context.Context, channelId uuid.UUID, role string) error {
	_, err := c.db.ExecContext(ctx, removeWritePermSQL, channelId, role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return types.NewErrNotFound(err)
		}
		return types.NewErrInternalFailure(err)
	}
	return nil
}

var removeReadPermSQL = `
	DELETE FROM read_permissions WHERE channel_id=$1 AND role_name=$2;
`

func (c *channelPostgres) RemoveReadPermission(ctx context.Context, channelId uuid.UUID, role string) error {
	_, err := c.db.ExecContext(ctx, removeReadPermSQL, channelId, role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return types.NewErrNotFound(err)
		}
		return types.NewErrInternalFailure(err)
	}
	return nil
}
