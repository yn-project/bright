package mixin

import (
	"entgo.io/ent"
	"yun.tea/block/bright/contract/pkg/db/ent/privacy"
	"yun.tea/block/bright/contract/pkg/db/rule"
)

func (TimeMixin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (TimeMixin) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rule.FilterTimeRule(),
		},
	}
}