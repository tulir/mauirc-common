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

// Preview is an URL preview
type Preview struct {
	Text  *Text  `json:"text,omitempty"`
	Image *Image `json:"image,omitempty"`
}

// Text is some text to preview
type Text struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	SiteName    string `json:"sitename,omitempty"`
}

// Image is an image to preview
type Image struct {
	URL    string `json:"url"`
	Type   string `json:"type"`
	Width  uint64 `json:"width"`
	Height uint64 `json:"height"`
}
