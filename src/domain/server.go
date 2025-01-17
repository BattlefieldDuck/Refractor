/*
 * This file is part of Refractor.
 *
 * Refractor is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package domain

import (
	"Refractor/pkg/broadcast"
	"context"
	"database/sql"
	"time"
)

type Server struct {
	ID           int64     `json:"id"`
	Game         string    `json:"game"`
	Name         string    `json:"name"`
	Address      string    `json:"address"`
	RCONPort     string    `json:"rcon_port"`
	RCONPassword string    `json:"-"`
	Deactivated  bool      `json:"deactivated"`
	CreatedAt    time.Time `json:"created_at"`
	ModifiedAt   time.Time `json:"modified_at"`
	IsFragment   bool      `json:"-"` // not a database field
}

type DBServer struct {
	ID           int64
	Game         string
	Name         string
	Address      string
	RCONPort     string
	RCONPassword string
	Deactivated  bool
	CreatedAt    sql.NullTime
	ModifiedAt   sql.NullTime
}

func (dbs DBServer) Server() *Server {
	s := &Server{
		ID:           dbs.ID,
		Game:         dbs.Game,
		Name:         dbs.Name,
		Address:      dbs.Address,
		RCONPort:     dbs.RCONPort,
		RCONPassword: dbs.RCONPassword,
		Deactivated:  dbs.Deactivated,
	}

	if dbs.CreatedAt.Valid {
		s.CreatedAt = dbs.CreatedAt.Time
	}

	if dbs.ModifiedAt.Valid {
		s.ModifiedAt = dbs.ModifiedAt.Time
	}

	return s
}

type ServerData struct {
	NeedsUpdate         bool
	ServerID            int64
	Status              string
	PlayerCount         int
	OnlinePlayers       map[string]*Player
	ReconnectInProgress bool
}

type ServerRepo interface {
	Store(ctx context.Context, server *Server) error
	GetByID(ctx context.Context, id int64) (*Server, error)
	GetAll(ctx context.Context) ([]*Server, error)
	Deactivate(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, args UpdateArgs) (*Server, error)
	Exists(ctx context.Context, args FindArgs) (bool, error)
}

type ServerUpdateSubscriber func(server *Server)

type ServerService interface {
	Store(c context.Context, server *Server) error
	GetByID(c context.Context, id int64) (*Server, error)
	GetAll(c context.Context) ([]*Server, error)
	// GetAllAccessible returns all servers on which the requesting user has the permission flag ViewServers set.
	GetAllAccessible(c context.Context) ([]*Server, error)
	Deactivate(c context.Context, id int64) error
	CreateServerData(id int64) error
	GetAllServerData() ([]*ServerData, error)
	GetServerData(id int64) (*ServerData, error)
	Update(c context.Context, id int64, args UpdateArgs) (*Server, error)
	HandlePlayerJoin(fields broadcast.Fields, serverID int64, game Game)
	HandlePlayerQuit(fields broadcast.Fields, serverID int64, game Game)
	HandleServerStatusChange(serverID int64, status string)
	HandlePlayerListUpdate(serverID int64, players []*OnlinePlayer, game Game)
	SubscribeServerUpdate(sub ServerUpdateSubscriber)
}
