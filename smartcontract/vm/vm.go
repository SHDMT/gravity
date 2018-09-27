// This file is part of the Dazzle Gravity library.
//
// The Dazzle Gravity library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Dazzle Gravity library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Dazzle Gravity library. If not, see <e <http://www.gnu.org/licenses/>./>.
package vm

import "github.com/SHDMT/gravity/platform/consensus/structure"

const (
	//LispScriptCode is lisp code number
	LispScriptCode = 1
)

//VM is virtual machine interface
type VM interface {
	ID() uint32
	Version() uint16
	ScriptCode() byte

	Config() Config
	Context() Context

	Exec(contract structure.Contract) bool
	SetEnv(context Context, config Config)
}
