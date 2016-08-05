// mauIRC-server - The IRC bouncer/backend system for mauIRC clients.
// Copyright (C) 2016 Tulir Asokan

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Package messages contains mauIRC client <-> server messages
package messages

// Message types
const (
	MsgRaw         = "raw"
	MsgInvite      = "invite"
	MsgNickChange  = "nickchange"
	MsgNetData     = "netdata"
	MsgChanData    = "chandata"
	MsgWhois       = "whois"
	MsgClear       = "clear"
	MsgDelete      = "delete"
	MsgChanList    = "chanlist"
	MsgCmdResponse = "cmdresponse"
	MsgMessage     = "message"
	MsgKick        = "kick"
	MsgMode        = "mode"
	MsgClose       = "close"
	MsgOpen        = "open"
)

// Container is a basic wrapper for a type string and the actual message object
type Container struct {
	Type   string      `json:"type"`
	Object interface{} `json:"object"`
}

// Message wraps an IRC message
type Message struct {
	ID        int64    `json:"id"`
	Network   string   `json:"network"`
	Channel   string   `json:"channel"`
	Timestamp int64    `json:"timestamp"`
	Sender    string   `json:"sender"`
	Command   string   `json:"command"`
	Message   string   `json:"message"`
	OwnMsg    bool     `json:"ownmsg"`
	Preview   *Preview `json:"preview,omitempty"`
}

// ParseMessage parses a Message object from a generic object
func ParseMessage(obj interface{}) (msg Message) {
	mp, ok := obj.(map[string]interface{})
	if !ok {
		return
	}

	msg.ID, _ = mp["id"].(int64)
	msg.Timestamp, _ = mp["timestamp"].(int64)
	msg.Network, _ = mp["network"].(string)
	msg.Channel, _ = mp["channel"].(string)
	msg.Sender, _ = mp["sender"].(string)
	msg.Command, _ = mp["command"].(string)
	msg.Message, _ = mp["message"].(string)
	msg.OwnMsg, _ = mp["ownmsg"].(bool)
	pw, ok := mp["preview"]
	if ok {
		msg.Preview = ParsePreview(pw)
	}
	return
}

// RawMessage is a raw IRC message
type RawMessage struct {
	Network string `json:"network"`
	Message string `json:"message"`
}

// ParseRawMessage parses a RawMessage object from a generic object
func ParseRawMessage(obj interface{}) (msg RawMessage) {
	mp, ok := obj.(map[string]interface{})
	if !ok {
		return
	}

	msg.Network, _ = mp["network"].(string)
	msg.Message, _ = mp["message"].(string)
	return
}

// NickChange is the message for IRC nick changes
type NickChange struct {
	Network string `json:"network"`
	Nick    string `json:"nick"`
}

// ParseNickChange parses a NickChange object from a generic object
func ParseNickChange(obj interface{}) (msg NickChange) {
	mp, ok := obj.(map[string]interface{})
	if !ok {
		return
	}

	msg.Network, _ = mp["network"].(string)
	msg.Nick, _ = mp["nick"].(string)
	return
}

// NetData contains basic network data
type NetData struct {
	Name      string `json:"name"`
	User      string `json:"user"`
	Realname  string `json:"realname"`
	Nick      string `json:"nick"`
	IP        string `json:"ip"`
	Port      uint16 `json:"port"`
	SSL       bool   `json:"ssl"`
	Connected bool   `json:"connected"`
}

// ParseNetData parses a NetData object from a generic object
func ParseNetData(obj interface{}) (msg NetData) {
	mp, ok := obj.(map[string]interface{})
	if !ok {
		return
	}

	msg.Name, _ = mp["name"].(string)
	msg.User, _ = mp["user"].(string)
	msg.Realname, _ = mp["realname"].(string)
	msg.Nick, _ = mp["nick"].(string)
	msg.IP, _ = mp["ip"].(string)
	msg.Port, _ = mp["port"].(uint16)
	msg.SSL, _ = mp["ssl"].(bool)
	msg.Connected, _ = mp["connected"].(bool)
	return
}

// ChanList contains a channel list and network name
type ChanList struct {
	Network string   `json:"network"`
	List    []string `json:"list"`
}

// ParseChanList parses a ChanList object from a generic object
func ParseChanList(obj interface{}) (msg ChanList) {
	mp, ok := obj.(map[string]interface{})
	if !ok {
		return
	}

	msg.Network, _ = mp["network"].(string)
	msg.List, _ = mp["nick"].([]string)
	return
}

// Invite an user to a channel
type Invite struct {
	Network string `json:"network"`
	Channel string `json:"channel"`
	Sender  string `json:"sender"`
}

// ParseInvite parses a Invite object from a generic object
func ParseInvite(obj interface{}) (msg Invite) {
	mp, ok := obj.(map[string]interface{})
	if !ok {
		return
	}

	msg.Network, _ = mp["network"].(string)
	msg.Channel, _ = mp["channel"].(string)
	msg.Sender, _ = mp["sender"].(string)
	return
}

// CommandResponse is a response to an user-sent internal command
type CommandResponse struct {
	Success       bool   `json:"success"`
	SimpleMessage string `json:"simple-message"`
	Message       string `json:"message"`
}

// ParseCommandResponse parses a CommandResponse object from a generic object
func ParseCommandResponse(obj interface{}) (msg CommandResponse) {
	mp, ok := obj.(map[string]interface{})
	if !ok {
		return
	}

	msg.Success, _ = mp["success"].(bool)
	msg.SimpleMessage, _ = mp["simple-message"].(string)
	msg.Message, _ = mp["message"].(string)
	return
}

// ClearHistory tells the client to clear the specific channel
type ClearHistory struct {
	Network string `json:"network"`
	Channel string `json:"channel"`
}

// ParseClearHistory parses a ClearHistory object from a generic object
func ParseClearHistory(obj interface{}) (msg ClearHistory) {
	mp, ok := obj.(map[string]interface{})
	if !ok {
		return
	}

	msg.Network, _ = mp["network"].(string)
	msg.Channel, _ = mp["channel"].(string)
	return
}

// WhoisData contains WHOIS information
type WhoisData struct {
	Channels   map[string]string `json:"channels"`
	Nick       string            `json:"nick"`
	User       string            `json:"user"`
	Host       string            `json:"host"`
	RealName   string            `json:"realname"`
	Away       string            `json:"away"`
	Server     string            `json:"server"`
	ServerInfo string            `json:"server-info"`
	IdleTime   int64             `json:"idle"`
	SecureConn bool              `json:"secure-connection"`
	Operator   bool              `json:"operator"`
}

// ParseWhoisData parses a WhoisData object from a generic object
func ParseWhoisData(obj interface{}) (msg WhoisData) {
	mp, ok := obj.(map[string]interface{})
	if !ok {
		return
	}

	chlist, _ := mp["channels"].(map[string]interface{})
	msg.Channels = make(map[string]string)
	for ch, role := range chlist {
		msg.Channels[ch], _ = role.(string)
	}

	msg.Nick, _ = mp["nick"].(string)
	msg.User, _ = mp["user"].(string)
	msg.Host, _ = mp["host"].(string)
	msg.RealName, _ = mp["realname"].(string)
	msg.Away, _ = mp["away"].(string)
	msg.Server, _ = mp["server"].(string)
	msg.ServerInfo, _ = mp["server-info"].(string)
	msg.IdleTime, _ = mp["idle"].(int64)
	msg.SecureConn, _ = mp["secure-connection"].(bool)
	msg.Operator, _ = mp["operator"].(bool)
	return
}
