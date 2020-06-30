/*
   Nging is a toolbox for webmasters
   Copyright (C) 2018-present  Wenhui Shen <swh@admpub.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package notice

var (
	defaultUserNotices *userNotices
	debug              bool //= true
)

func SetDebug(on bool) {
	debug = on
}

func Default() *userNotices {
	if defaultUserNotices == nil {
		defaultUserNotices = NewUserNotices(debug)
	}
	return defaultUserNotices
}

func OnClose(fn ...func(user string)) *userNotices {
	return Default().OnClose(fn...)
}

func OnOpen(fn ...func(user string)) *userNotices {
	return Default().OnOpen(fn...)
}

func Send(user string, message *Message) error {
	return Default().Send(user, message)
}

func Recv(user string, clientID string) <-chan *Message {
	return Default().Recv(user, clientID)
}

func RecvJSON(user string, clientID string) ([]byte, error) {
	return Default().RecvJSON(user, clientID)
}

func RecvXML(user string, clientID string) ([]byte, error) {
	return Default().RecvXML(user, clientID)
}

func CloseClient(user string, clientID string) bool {
	return Default().CloseClient(user, clientID)
}

func OpenClient(user string) string {
	return Default().OpenClient(user)
}

func CloseMessage(user string, types ...string) {
	Default().CloseMessage(user, types...)
}

func OpenMessage(user string, types ...string) {
	Default().OpenMessage(user, types...)
}

func Clear() {
	Default().Clear()
}
