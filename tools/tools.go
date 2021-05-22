/*********************************************************************

rss3go_hub: An alternative version of RSSHub for RSS3 written in go

Copyright (C) 2021 Nyawork, Candinya

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

 ********************************************************************/

package tools

import (
	"github.com/imdario/mergo"
	"github.com/nyawork/rss3go_lib/types"
)

func DeepMergePersona(old types.RSS3, patch interface{}) (types.RSS3, error) {

	var patchPersona types.RSS3

	if err := mergo.Map(&patchPersona, patch); err != nil {
		return old, err
	}
	err := mergo.Merge(&old, patchPersona, mergo.WithOverride)

	return old, err
}

func DeepMergeItems(old types.RSS3Items, patch interface{}) (types.RSS3Items, error) {

	var patchItems types.RSS3Items

	if err := mergo.Map(&patchItems, patch); err != nil {
		return old, err
	}
	err := mergo.Merge(&old, patchItems, mergo.WithOverride)

	return old, err
}

func DeepMergeList(old types.RSS3List, patch interface{}) (types.RSS3List, error) {

	var patchLink types.RSS3List

	if err := mergo.Map(&patchLink, patch); err != nil {
		return old, err
	}
	err := mergo.Merge(&old, patchLink, mergo.WithOverride)

	return old, err
}

func DeepMergeItem(old *types.RSS3Item, patch interface{}) error {

	var patchItem types.RSS3Item

	if err := mergo.Map(&patchItem, patch); err != nil {
		return err
	}
	err := mergo.Merge(old, patchItem, mergo.WithOverride)

	return err
}
