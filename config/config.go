//  Copyright (c) 2015 Couchbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	// token maps
	_ "github.com/qri-io/bleve/analysis/tokenmap"

	// fragment formatters
	_ "github.com/qri-io/bleve/search/highlight/format/ansi"
	_ "github.com/qri-io/bleve/search/highlight/format/html"

	// fragmenters
	_ "github.com/qri-io/bleve/search/highlight/fragmenter/simple"

	// highlighters
	_ "github.com/qri-io/bleve/search/highlight/highlighter/ansi"
	_ "github.com/qri-io/bleve/search/highlight/highlighter/html"
	_ "github.com/qri-io/bleve/search/highlight/highlighter/simple"

	// char filters
	_ "github.com/qri-io/bleve/analysis/char/html"
	_ "github.com/qri-io/bleve/analysis/char/regexp"
	_ "github.com/qri-io/bleve/analysis/char/zerowidthnonjoiner"

	// analyzers
	_ "github.com/qri-io/bleve/analysis/analyzer/custom"
	_ "github.com/qri-io/bleve/analysis/analyzer/keyword"
	_ "github.com/qri-io/bleve/analysis/analyzer/simple"
	_ "github.com/qri-io/bleve/analysis/analyzer/standard"
	_ "github.com/qri-io/bleve/analysis/analyzer/web"

	// token filters
	_ "github.com/qri-io/bleve/analysis/token/apostrophe"
	_ "github.com/qri-io/bleve/analysis/token/compound"
	_ "github.com/qri-io/bleve/analysis/token/edgengram"
	_ "github.com/qri-io/bleve/analysis/token/elision"
	_ "github.com/qri-io/bleve/analysis/token/keyword"
	_ "github.com/qri-io/bleve/analysis/token/length"
	_ "github.com/qri-io/bleve/analysis/token/lowercase"
	_ "github.com/qri-io/bleve/analysis/token/ngram"
	_ "github.com/qri-io/bleve/analysis/token/shingle"
	_ "github.com/qri-io/bleve/analysis/token/stop"
	_ "github.com/qri-io/bleve/analysis/token/truncate"
	_ "github.com/qri-io/bleve/analysis/token/unicodenorm"

	// tokenizers
	_ "github.com/qri-io/bleve/analysis/tokenizer/exception"
	_ "github.com/qri-io/bleve/analysis/tokenizer/regexp"
	_ "github.com/qri-io/bleve/analysis/tokenizer/single"
	_ "github.com/qri-io/bleve/analysis/tokenizer/unicode"
	_ "github.com/qri-io/bleve/analysis/tokenizer/web"
	_ "github.com/qri-io/bleve/analysis/tokenizer/whitespace"

	// date time parsers
	_ "github.com/qri-io/bleve/analysis/datetime/flexible"
	_ "github.com/qri-io/bleve/analysis/datetime/optional"

	// languages
	_ "github.com/qri-io/bleve/analysis/lang/ar"
	_ "github.com/qri-io/bleve/analysis/lang/bg"
	_ "github.com/qri-io/bleve/analysis/lang/ca"
	_ "github.com/qri-io/bleve/analysis/lang/cjk"
	_ "github.com/qri-io/bleve/analysis/lang/ckb"
	_ "github.com/qri-io/bleve/analysis/lang/cs"
	_ "github.com/qri-io/bleve/analysis/lang/el"
	_ "github.com/qri-io/bleve/analysis/lang/en"
	_ "github.com/qri-io/bleve/analysis/lang/eu"
	_ "github.com/qri-io/bleve/analysis/lang/fa"
	_ "github.com/qri-io/bleve/analysis/lang/fr"
	_ "github.com/qri-io/bleve/analysis/lang/ga"
	_ "github.com/qri-io/bleve/analysis/lang/gl"
	_ "github.com/qri-io/bleve/analysis/lang/hi"
	_ "github.com/qri-io/bleve/analysis/lang/hy"
	_ "github.com/qri-io/bleve/analysis/lang/id"
	_ "github.com/qri-io/bleve/analysis/lang/in"
	_ "github.com/qri-io/bleve/analysis/lang/it"
	_ "github.com/qri-io/bleve/analysis/lang/pt"

	// kv stores
	_ "github.com/qri-io/bleve/index/store/boltdb"
	_ "github.com/qri-io/bleve/index/store/goleveldb"
	_ "github.com/qri-io/bleve/index/store/gtreap"
	_ "github.com/qri-io/bleve/index/store/moss"

	// index types
	_ "github.com/qri-io/bleve/index/upsidedown"
)
