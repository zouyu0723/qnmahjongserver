// Package dao contains the types for schema 'mj'.
package dao

// GENERATED BY XO. DO NOT EDIT.

import "errors"

// Game represents a row from 'mj.game'.
type Game struct {
	IndexID     int32  `json:"index_id"`     // index_id
	Channel     int32  `json:"channel"`      // channel
	Version     int32  `json:"version"`      // version
	Size        int32  `json:"size"`         // size
	Module      string `json:"module"`       // module
	MjTypes     string `json:"mj_types"`     // mj_types
	Enabled     int32  `json:"enabled"`      // enabled
	UpdateType  int32  `json:"update_type"`  // update_type
	DownloadURL string `json:"download_url"` // download_url
	SvnVersion  int32  `json:"svn_version"`  // svn_version

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Game exists in the database.
func (g *Game) Exists() bool {
	return g._exists
}

// Deleted provides information if the Game has been deleted from the database.
func (g *Game) Deleted() bool {
	return g._deleted
}

// Insert inserts the Game to the database.
func (g *Game) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if g._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO mj.game (` +
		`channel, version, size, module, mj_types, enabled, update_type, download_url, svn_version` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, g.Channel, g.Version, g.Size, g.Module, g.MjTypes, g.Enabled, g.UpdateType, g.DownloadURL, g.SvnVersion)
	res, err := db.Exec(sqlstr, g.Channel, g.Version, g.Size, g.Module, g.MjTypes, g.Enabled, g.UpdateType, g.DownloadURL, g.SvnVersion)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	g.IndexID = int32(id)
	g._exists = true

	return nil
}

// Update updates the Game in the database.
func (g *Game) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !g._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if g._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE mj.game SET ` +
		`channel = ?, version = ?, size = ?, module = ?, mj_types = ?, enabled = ?, update_type = ?, download_url = ?, svn_version = ?` +
		` WHERE index_id = ?`

	// run query
	XOLog(sqlstr, g.Channel, g.Version, g.Size, g.Module, g.MjTypes, g.Enabled, g.UpdateType, g.DownloadURL, g.SvnVersion, g.IndexID)
	_, err = db.Exec(sqlstr, g.Channel, g.Version, g.Size, g.Module, g.MjTypes, g.Enabled, g.UpdateType, g.DownloadURL, g.SvnVersion, g.IndexID)
	return err
}

// Save saves the Game to the database.
func (g *Game) Save(db XODB) error {
	if g.Exists() {
		return g.Update(db)
	}

	return g.Insert(db)
}

// Delete deletes the Game from the database.
func (g *Game) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !g._exists {
		return nil
	}

	// if deleted, bail
	if g._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM mj.game WHERE index_id = ?`

	// run query
	XOLog(sqlstr, g.IndexID)
	_, err = db.Exec(sqlstr, g.IndexID)
	if err != nil {
		return err
	}

	// set deleted
	g._deleted = true

	return nil
}

// GameByIndexID retrieves a row from 'mj.game' as a Game.
//
// Generated from index 'game_index_id_pkey'.
func GameByIndexID(db XODB, indexID int32) (*Game, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`index_id, channel, version, size, module, mj_types, enabled, update_type, download_url, svn_version ` +
		`FROM mj.game ` +
		`WHERE index_id = ?`

	// run query
	XOLog(sqlstr, indexID)
	g := Game{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, indexID).Scan(&g.IndexID, &g.Channel, &g.Version, &g.Size, &g.Module, &g.MjTypes, &g.Enabled, &g.UpdateType, &g.DownloadURL, &g.SvnVersion)
	if err != nil {
		return nil, err
	}

	return &g, nil
}

// GamesByChannel retrieves a row from 'mj.game' as a Game.
//
// Generated from index 'idx_channel'.
func GamesByChannel(db XODB, channel int32) ([]*Game, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`index_id, channel, version, size, module, mj_types, enabled, update_type, download_url, svn_version ` +
		`FROM mj.game ` +
		`WHERE channel = ?`

	// run query
	XOLog(sqlstr, channel)
	q, err := db.Query(sqlstr, channel)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Game{}
	for q.Next() {
		g := Game{
			_exists: true,
		}

		// scan
		err = q.Scan(&g.IndexID, &g.Channel, &g.Version, &g.Size, &g.Module, &g.MjTypes, &g.Enabled, &g.UpdateType, &g.DownloadURL, &g.SvnVersion)
		if err != nil {
			return nil, err
		}

		res = append(res, &g)
	}

	return res, nil
}

// GameByChannelVersion retrieves a row from 'mj.game' as a Game.
//
// Generated from index 'uidx_channel_version'.
func GameByChannelVersion(db XODB, channel int32, version int32) (*Game, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`index_id, channel, version, size, module, mj_types, enabled, update_type, download_url, svn_version ` +
		`FROM mj.game ` +
		`WHERE channel = ? AND version = ?`

	// run query
	XOLog(sqlstr, channel, version)
	g := Game{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, channel, version).Scan(&g.IndexID, &g.Channel, &g.Version, &g.Size, &g.Module, &g.MjTypes, &g.Enabled, &g.UpdateType, &g.DownloadURL, &g.SvnVersion)
	if err != nil {
		return nil, err
	}

	return &g, nil
}