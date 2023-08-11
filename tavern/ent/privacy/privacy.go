// Code generated by ent, DO NOT EDIT.

package privacy

import (
	"context"
	"fmt"

	"github.com/kcarretto/realm/tavern/ent"

	"entgo.io/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...any) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...any) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...any) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// Policy groups query and mutation policies.
	Policy = privacy.Policy

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy

	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

// MutationRuleFunc type is an adapter which allows the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// QueryMutationRule is an interface which groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, ent.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, ent.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ ent.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The BeaconQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type BeaconQueryRuleFunc func(context.Context, *ent.BeaconQuery) error

// EvalQuery return f(ctx, q).
func (f BeaconQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.BeaconQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.BeaconQuery", q)
}

// The BeaconMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type BeaconMutationRuleFunc func(context.Context, *ent.BeaconMutation) error

// EvalMutation calls f(ctx, m).
func (f BeaconMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.BeaconMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.BeaconMutation", m)
}

// The FileQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type FileQueryRuleFunc func(context.Context, *ent.FileQuery) error

// EvalQuery return f(ctx, q).
func (f FileQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.FileQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.FileQuery", q)
}

// The FileMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type FileMutationRuleFunc func(context.Context, *ent.FileMutation) error

// EvalMutation calls f(ctx, m).
func (f FileMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.FileMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.FileMutation", m)
}

// The QuestQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type QuestQueryRuleFunc func(context.Context, *ent.QuestQuery) error

// EvalQuery return f(ctx, q).
func (f QuestQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.QuestQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.QuestQuery", q)
}

// The QuestMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type QuestMutationRuleFunc func(context.Context, *ent.QuestMutation) error

// EvalMutation calls f(ctx, m).
func (f QuestMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.QuestMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.QuestMutation", m)
}

// The TagQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TagQueryRuleFunc func(context.Context, *ent.TagQuery) error

// EvalQuery return f(ctx, q).
func (f TagQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TagQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.TagQuery", q)
}

// The TagMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TagMutationRuleFunc func(context.Context, *ent.TagMutation) error

// EvalMutation calls f(ctx, m).
func (f TagMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.TagMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.TagMutation", m)
}

// The TaskQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TaskQueryRuleFunc func(context.Context, *ent.TaskQuery) error

// EvalQuery return f(ctx, q).
func (f TaskQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TaskQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.TaskQuery", q)
}

// The TaskMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TaskMutationRuleFunc func(context.Context, *ent.TaskMutation) error

// EvalMutation calls f(ctx, m).
func (f TaskMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.TaskMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.TaskMutation", m)
}

// The TomeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TomeQueryRuleFunc func(context.Context, *ent.TomeQuery) error

// EvalQuery return f(ctx, q).
func (f TomeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TomeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.TomeQuery", q)
}

// The TomeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TomeMutationRuleFunc func(context.Context, *ent.TomeMutation) error

// EvalMutation calls f(ctx, m).
func (f TomeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.TomeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.TomeMutation", m)
}

// The UserQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserQueryRuleFunc func(context.Context, *ent.UserQuery) error

// EvalQuery return f(ctx, q).
func (f UserQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UserQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.UserQuery", q)
}

// The UserMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserMutationRuleFunc func(context.Context, *ent.UserMutation) error

// EvalMutation calls f(ctx, m).
func (f UserMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.UserMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.UserMutation", m)
}
