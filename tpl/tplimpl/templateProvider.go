// Copyright 2017-present The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tplimpl

import (
	"github.com/gohugoio/hugo/deps"
)

// TemplateProvider manages templates.
type TemplateProvider struct{}

// DefaultTemplateProvider is a globally available TemplateProvider.
var DefaultTemplateProvider *TemplateProvider

// Update updates the Hugo Template System in the provided Deps
// with all the additional features, templates & functions.
func (*TemplateProvider) Update(deps *deps.Deps) error {
	newTmpl := newTemplateAdapter(deps)
	deps.Tmpl = newTmpl
	deps.TextTmpl = newTmpl.wrapTextTemplate(newTmpl.text.standalone)
	// These needs to be there at parse time.
	newTmpl.initTemplateExecuter()

	if err := newTmpl.loadEmbedded(); err != nil {
		return err
	}

	if deps.WithTemplate != nil {
		err := deps.WithTemplate(newTmpl)
		if err != nil {
			return err
		}

	}

	return newTmpl.markReady()

}

// Clone clones.
func (*TemplateProvider) Clone(d *deps.Deps) error {
	t := d.Tmpl.(*templateHandler)
	t.clone(d)
	return nil
}
