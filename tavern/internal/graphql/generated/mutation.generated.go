// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/kcarretto/realm/tavern/internal/ent"
	"github.com/kcarretto/realm/tavern/internal/graphql/models"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type MutationResolver interface {
	CreateQuest(ctx context.Context, beaconIDs []int, input ent.CreateQuestInput) (*ent.Quest, error)
	UpdateBeacon(ctx context.Context, beaconID int, input ent.UpdateBeaconInput) (*ent.Beacon, error)
	UpdateHost(ctx context.Context, hostID int, input ent.UpdateHostInput) (*ent.Host, error)
	CreateTag(ctx context.Context, input ent.CreateTagInput) (*ent.Tag, error)
	UpdateTag(ctx context.Context, tagID int, input ent.UpdateTagInput) (*ent.Tag, error)
	ClaimTasks(ctx context.Context, input models.ClaimTasksInput) ([]*ent.Task, error)
	SubmitTaskResult(ctx context.Context, input models.SubmitTaskResultInput) (*ent.Task, error)
	CreateTome(ctx context.Context, input ent.CreateTomeInput) (*ent.Tome, error)
	UpdateUser(ctx context.Context, userID int, input ent.UpdateUserInput) (*ent.User, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) field_Mutation_claimTasks_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 models.ClaimTasksInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNClaimTasksInput2githubᚗcomᚋkcarrettoᚋrealmᚋtavernᚋinternalᚋgraphqlᚋmodelsᚐClaimTasksInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_createQuest_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 []int
	if tmp, ok := rawArgs["beaconIDs"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("beaconIDs"))
		arg0, err = ec.unmarshalNID2ᚕintᚄ(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["beaconIDs"] = arg0
	var arg1 ent.CreateQuestInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg1, err = ec.unmarshalNCreateQuestInput2githubᚗcomᚋkcarrettoᚋrealmᚋtavernᚋinternalᚋentᚐCreateQuestInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg1
	return args, nil
}

func (ec *executionContext) field_Mutation_createTag_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 ent.CreateTagInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNCreateTagInput2githubᚗcomᚋkcarrettoᚋrealmᚋtavernᚋinternalᚋentᚐCreateTagInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_createTome_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 ent.CreateTomeInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNCreateTomeInput2githubᚗcomᚋkcarrettoᚋrealmᚋtavernᚋinternalᚋentᚐCreateTomeInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_submitTaskResult_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 models.SubmitTaskResultInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNSubmitTaskResultInput2githubᚗcomᚋkcarrettoᚋrealmᚋtavernᚋinternalᚋgraphqlᚋmodelsᚐSubmitTaskResultInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_updateBeacon_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 int
	if tmp, ok := rawArgs["beaconID"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("beaconID"))
		arg0, err = ec.unmarshalNID2int(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["beaconID"] = arg0
	var arg1 ent.UpdateBeaconInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg1, err = ec.unmarshalNUpdateBeaconInput2githubᚗcomᚋkcarrettoᚋrealmᚋtavernᚋinternalᚋentᚐUpdateBeaconInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg1
	return args, nil
}

func (ec *executionContext) field_Mutation_updateHost_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 int
	if tmp, ok := rawArgs["hostID"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("hostID"))
		arg0, err = ec.unmarshalNID2int(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["hostID"] = arg0
	var arg1 ent.UpdateHostInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg1, err = ec.unmarshalNUpdateHostInput2githubᚗcomᚋkcarrettoᚋrealmᚋtavernᚋinternalᚋentᚐUpdateHostInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg1
	return args, nil
}

func (ec *executionContext) field_Mutation_updateTag_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 int
	if tmp, ok := rawArgs["tagID"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("tagID"))
		arg0, err = ec.unmarshalNID2int(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["tagID"] = arg0
	var arg1 ent.UpdateTagInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg1, err = ec.unmarshalNUpdateTagInput2githubᚗcomᚋkcarrettoᚋrealmᚋtavernᚋinternalᚋentᚐUpdateTagInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg1
	return args, nil
}

func (ec *executionContext) field_Mutation_updateUser_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 int
	if tmp, ok := rawArgs["userID"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("userID"))
		arg0, err = ec.unmarshalNID2int(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["userID"] = arg0
	var arg1 ent.UpdateUserInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg1, err = ec.unmarshalNUpdateUserInput2githubᚗcomᚋkcarrettoᚋrealmᚋtavernᚋinternalᚋentᚐUpdateUserInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg1
	return args, nil
}

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Mutation_createQuest(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_createQuest(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		directive0 := func(rctx context.Context) (interface{}, error) {
			ctx = rctx // use context from middleware stack in children
			return ec.resolvers.Mutation().CreateQuest(rctx, fc.Args["beaconIDs"].([]int), fc.Args["input"].(ent.CreateQuestInput))
		}
		directive1 := func(ctx context.Context) (interface{}, error) {
			role, err := ec.unmarshalNRole2githubᚗcomᚋkcarrettoᚋrealmᚋtavernᚋinternalᚋgraphqlᚋmodelsᚐRole(ctx, "USER")
			if err != nil {
				return nil, err
			}
			if ec.directives.RequireRole == nil {
				return nil, errors.New("directive requireRole is not implemented")
			}
			return ec.directives.RequireRole(ctx, nil, directive0, role)
		}

		tmp, err := directive1(rctx)
		if err != nil {
			return nil, graphql.ErrorOnPath(ctx, err)
		}
		if tmp == nil {
			return nil, nil
		}
		if data, ok := tmp.(*ent.Quest); ok {
			return data, nil
		}
		return nil, fmt.Errorf(`unexpected type %T from directive, should be *github.com/kcarretto/realm/tavern/internal/ent.Quest`, tmp)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*ent.Quest)
	fc.Result = res
	return ec.marshalOQuest2ᚖgithubᚗcomᚋkcarrettoᚋrealmᚋtavernᚋinternalᚋentᚐQuest(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_createQuest(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Quest_id(ctx, field)
			case "createdAt":
				return ec.fieldContext_Quest_createdAt(ctx, field)
			case "lastModifiedAt":
				return ec.fieldContext_Quest_lastModifiedAt(ctx, field)
			case "name":
				return ec.fieldContext_Quest_name(ctx, field)
			case "parameters":
				return ec.fieldContext_Quest_parameters(ctx, field)
			case "tome":
				return ec.fieldContext_Quest_tome(ctx, field)
			case "bundle":
				return ec.fieldContext_Quest_bundle(ctx, field)
			case "tasks":
				return ec.fieldContext_Quest_tasks(ctx, field)
			case "creator":
				return ec.fieldContext_Quest_creator(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Quest", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_createQuest_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_updateBeacon(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_updateBeacon(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		directive0 := func(rctx context.Context) (interface{}, error) {
			ctx = rctx // use context from middleware stack in children
			return ec.resolvers.Mutation().UpdateBeacon(rctx, fc.Args["beaconID"].(int), fc.Args["input"].(ent.UpdateBeaconInput))
		}
		directive1 := func(ctx context.Context) (interface{}, error) {
			role, err := ec.unmarshalNRole2githubᚗcomᚋkcarrettoᚋrealmᚋtavernᚋinternalᚋgraphqlᚋmodelsᚐRole(ctx, "USER")
			if err != nil {
				return nil, err
			}
			if ec.directives.RequireRole == nil {
				return nil, errors.New("directive requireRole is not implemented")
			}
			return ec.directives.RequireRole(ctx, nil, directive0, role)
		}

		tmp, err := directive1(rctx)
		if err != nil {
			return nil, graphql.ErrorOnPath(ctx, err)
		}
		if tmp == nil {
			return nil, nil
		}
		if data, ok := tmp.(*ent.Beacon); ok {
			return data, nil
		}
		return nil, fmt.Errorf(`unexpected type %T from directive, should be *github.com/kcarretto/realm/tavern/internal/ent.Beacon`, tmp)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*ent.Beacon)
	fc.Result = res
	return ec.marshalNBeacon2ᚖgithubᚗcomᚋkcarrettoᚋrealmᚋtavernᚋinternalᚋentᚐBeacon(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_updateBeacon(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Beacon_id(ctx, field)
			case "name":
				return ec.fieldContext_Beacon_name(ctx, field)
			case "principal":
				return ec.fieldContext_Beacon_principal(ctx, field)
			case "identifier":
				return ec.fieldContext_Beacon_identifier(ctx, field)
			case "agentIdentifier":
				return ec.fieldContext_Beacon_agentIdentifier(ctx, field)
			case "lastSeenAt":
				return ec.fieldContext_Beacon_lastSeenAt(ctx, field)
			case "interval":
				return ec.fieldContext_Beacon_interval(ctx, field)
			case "host":
				return ec.fieldContext_Beacon_host(ctx, field)
			case "tasks":
				return ec.fieldContext_Beacon_tasks(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Beacon", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_updateBeacon_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_updateHost(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_updateHost(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		directive0 := func(rctx context.Context) (interface{}, error) {
			ctx = rctx // use context from middleware stack in children
			return ec.resolvers.Mutation().UpdateHost(rctx, fc.Args["hostID"].(int), fc.Args["input"].(ent.UpdateHostInput))
		}
		directive1 := func(ctx context.Context) (interface{}, error) {
			role, err := ec.unmarshalNRole2githubᚗcomᚋkcarrettoᚋrealmᚋtavernᚋinternalᚋgraphqlᚋmodelsᚐRole(ctx, "USER")
			if err != nil {
				return nil, err
			}
			if ec.directives.RequireRole == nil {
				return nil, errors.New("directive requireRole is not implemented")
			}
			return ec.directives.RequireRole(ctx, nil, directive0, role)
		}

		tmp, err := directive1(rctx)
		if err != nil {
			return nil, graphql.ErrorOnPath(ctx, err)
		}
		if tmp == nil {
			return nil, nil
		}
		if data, ok := tmp.(*ent.Host); ok {
			return data, nil
		}
		return nil, fmt.Errorf(`unexpected type %T from directive, should be *github.com/kcarretto/realm/tavern/internal/ent.Host`, tmp)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*ent.Host)
	fc.Result = res
	return ec.marshalNHost2ᚖgithubᚗcomᚋkcarrettoᚋrealmᚋtavernᚋinternalᚋentᚐHost(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_updateHost(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Host_id(ctx, field)
			case "identifier":
				return ec.fieldContext_Host_identifier(ctx, field)
			case "name":
				return ec.fieldContext_Host_name(ctx, field)
			case "primaryIP":
				return ec.fieldContext_Host_primaryIP(ctx, field)
			case "platform":
				return ec.fieldContext_Host_platform(ctx, field)
			case "lastSeenAt":
				return ec.fieldContext_Host_lastSeenAt(ctx, field)
			case "tags":
				return ec.fieldContext_Host_tags(ctx, field)
			case "beacons":
				return ec.fieldContext_Host_beacons(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Host", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_updateHost_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_createTag(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_createTag(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		directive0 := func(rctx context.Context) (interface{}, error) {
			ctx = rctx // use context from middleware stack in children
			return ec.resolvers.Mutation().CreateTag(rctx, fc.Args["input"].(ent.CreateTagInput))
		}
		directive1 := func(ctx context.Context) (interface{}, error) {
			role, err := ec.unmarshalNRole2githubᚗcomᚋkcarrettoᚋrealmᚋtavernᚋinternalᚋgraphqlᚋmodelsᚐRole(ctx, "ADMIN")
			if err != nil {
				return nil, err
			}
			if ec.directives.RequireRole == nil {
				return nil, errors.New("directive requireRole is not implemented")
			}
			return ec.directives.RequireRole(ctx, nil, directive0, role)
		}

		tmp, err := directive1(rctx)
		if err != nil {
			return nil, graphql.ErrorOnPath(ctx, err)
		}
		if tmp == nil {
			return nil, nil
		}
		if data, ok := tmp.(*ent.Tag); ok {
			return data, nil
		}
		return nil, fmt.Errorf(`unexpected type %T from directive, should be *github.com/kcarretto/realm/tavern/internal/ent.Tag`, tmp)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*ent.Tag)
	fc.Result = res
	return ec.marshalNTag2ᚖgithubᚗcomᚋkcarrettoᚋrealmᚋtavernᚋinternalᚋentᚐTag(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_createTag(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Tag_id(ctx, field)
			case "name":
				return ec.fieldContext_Tag_name(ctx, field)
			case "kind":
				return ec.fieldContext_Tag_kind(ctx, field)
			case "hosts":
				return ec.fieldContext_Tag_hosts(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Tag", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_createTag_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_updateTag(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_updateTag(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		directive0 := func(rctx context.Context) (interface{}, error) {
			ctx = rctx // use context from middleware stack in children
			return ec.resolvers.Mutation().UpdateTag(rctx, fc.Args["tagID"].(int), fc.Args["input"].(ent.UpdateTagInput))
		}
		directive1 := func(ctx context.Context) (interface{}, error) {
			role, err := ec.unmarshalNRole2githubᚗcomᚋkcarrettoᚋrealmᚋtavernᚋinternalᚋgraphqlᚋmodelsᚐRole(ctx, "USER")
			if err != nil {
				return nil, err
			}
			if ec.directives.RequireRole == nil {
				return nil, errors.New("directive requireRole is not implemented")
			}
			return ec.directives.RequireRole(ctx, nil, directive0, role)
		}

		tmp, err := directive1(rctx)
		if err != nil {
			return nil, graphql.ErrorOnPath(ctx, err)
		}
		if tmp == nil {
			return nil, nil
		}
		if data, ok := tmp.(*ent.Tag); ok {
			return data, nil
		}
		return nil, fmt.Errorf(`unexpected type %T from directive, should be *github.com/kcarretto/realm/tavern/internal/ent.Tag`, tmp)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*ent.Tag)
	fc.Result = res
	return ec.marshalNTag2ᚖgithubᚗcomᚋkcarrettoᚋrealmᚋtavernᚋinternalᚋentᚐTag(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_updateTag(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Tag_id(ctx, field)
			case "name":
				return ec.fieldContext_Tag_name(ctx, field)
			case "kind":
				return ec.fieldContext_Tag_kind(ctx, field)
			case "hosts":
				return ec.fieldContext_Tag_hosts(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Tag", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_updateTag_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_claimTasks(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_claimTasks(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().ClaimTasks(rctx, fc.Args["input"].(models.ClaimTasksInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]*ent.Task)
	fc.Result = res
	return ec.marshalNTask2ᚕᚖgithubᚗcomᚋkcarrettoᚋrealmᚋtavernᚋinternalᚋentᚐTaskᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_claimTasks(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Task_id(ctx, field)
			case "createdAt":
				return ec.fieldContext_Task_createdAt(ctx, field)
			case "lastModifiedAt":
				return ec.fieldContext_Task_lastModifiedAt(ctx, field)
			case "claimedAt":
				return ec.fieldContext_Task_claimedAt(ctx, field)
			case "execStartedAt":
				return ec.fieldContext_Task_execStartedAt(ctx, field)
			case "execFinishedAt":
				return ec.fieldContext_Task_execFinishedAt(ctx, field)
			case "output":
				return ec.fieldContext_Task_output(ctx, field)
			case "error":
				return ec.fieldContext_Task_error(ctx, field)
			case "quest":
				return ec.fieldContext_Task_quest(ctx, field)
			case "beacon":
				return ec.fieldContext_Task_beacon(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Task", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_claimTasks_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_submitTaskResult(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_submitTaskResult(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().SubmitTaskResult(rctx, fc.Args["input"].(models.SubmitTaskResultInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*ent.Task)
	fc.Result = res
	return ec.marshalOTask2ᚖgithubᚗcomᚋkcarrettoᚋrealmᚋtavernᚋinternalᚋentᚐTask(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_submitTaskResult(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Task_id(ctx, field)
			case "createdAt":
				return ec.fieldContext_Task_createdAt(ctx, field)
			case "lastModifiedAt":
				return ec.fieldContext_Task_lastModifiedAt(ctx, field)
			case "claimedAt":
				return ec.fieldContext_Task_claimedAt(ctx, field)
			case "execStartedAt":
				return ec.fieldContext_Task_execStartedAt(ctx, field)
			case "execFinishedAt":
				return ec.fieldContext_Task_execFinishedAt(ctx, field)
			case "output":
				return ec.fieldContext_Task_output(ctx, field)
			case "error":
				return ec.fieldContext_Task_error(ctx, field)
			case "quest":
				return ec.fieldContext_Task_quest(ctx, field)
			case "beacon":
				return ec.fieldContext_Task_beacon(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Task", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_submitTaskResult_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_createTome(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_createTome(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		directive0 := func(rctx context.Context) (interface{}, error) {
			ctx = rctx // use context from middleware stack in children
			return ec.resolvers.Mutation().CreateTome(rctx, fc.Args["input"].(ent.CreateTomeInput))
		}
		directive1 := func(ctx context.Context) (interface{}, error) {
			role, err := ec.unmarshalNRole2githubᚗcomᚋkcarrettoᚋrealmᚋtavernᚋinternalᚋgraphqlᚋmodelsᚐRole(ctx, "USER")
			if err != nil {
				return nil, err
			}
			if ec.directives.RequireRole == nil {
				return nil, errors.New("directive requireRole is not implemented")
			}
			return ec.directives.RequireRole(ctx, nil, directive0, role)
		}

		tmp, err := directive1(rctx)
		if err != nil {
			return nil, graphql.ErrorOnPath(ctx, err)
		}
		if tmp == nil {
			return nil, nil
		}
		if data, ok := tmp.(*ent.Tome); ok {
			return data, nil
		}
		return nil, fmt.Errorf(`unexpected type %T from directive, should be *github.com/kcarretto/realm/tavern/internal/ent.Tome`, tmp)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*ent.Tome)
	fc.Result = res
	return ec.marshalNTome2ᚖgithubᚗcomᚋkcarrettoᚋrealmᚋtavernᚋinternalᚋentᚐTome(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_createTome(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Tome_id(ctx, field)
			case "createdAt":
				return ec.fieldContext_Tome_createdAt(ctx, field)
			case "lastModifiedAt":
				return ec.fieldContext_Tome_lastModifiedAt(ctx, field)
			case "name":
				return ec.fieldContext_Tome_name(ctx, field)
			case "description":
				return ec.fieldContext_Tome_description(ctx, field)
			case "paramDefs":
				return ec.fieldContext_Tome_paramDefs(ctx, field)
			case "eldritch":
				return ec.fieldContext_Tome_eldritch(ctx, field)
			case "files":
				return ec.fieldContext_Tome_files(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Tome", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_createTome_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_updateUser(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_updateUser(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		directive0 := func(rctx context.Context) (interface{}, error) {
			ctx = rctx // use context from middleware stack in children
			return ec.resolvers.Mutation().UpdateUser(rctx, fc.Args["userID"].(int), fc.Args["input"].(ent.UpdateUserInput))
		}
		directive1 := func(ctx context.Context) (interface{}, error) {
			role, err := ec.unmarshalNRole2githubᚗcomᚋkcarrettoᚋrealmᚋtavernᚋinternalᚋgraphqlᚋmodelsᚐRole(ctx, "ADMIN")
			if err != nil {
				return nil, err
			}
			if ec.directives.RequireRole == nil {
				return nil, errors.New("directive requireRole is not implemented")
			}
			return ec.directives.RequireRole(ctx, nil, directive0, role)
		}

		tmp, err := directive1(rctx)
		if err != nil {
			return nil, graphql.ErrorOnPath(ctx, err)
		}
		if tmp == nil {
			return nil, nil
		}
		if data, ok := tmp.(*ent.User); ok {
			return data, nil
		}
		return nil, fmt.Errorf(`unexpected type %T from directive, should be *github.com/kcarretto/realm/tavern/internal/ent.User`, tmp)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*ent.User)
	fc.Result = res
	return ec.marshalOUser2ᚖgithubᚗcomᚋkcarrettoᚋrealmᚋtavernᚋinternalᚋentᚐUser(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_updateUser(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_User_id(ctx, field)
			case "name":
				return ec.fieldContext_User_name(ctx, field)
			case "photoURL":
				return ec.fieldContext_User_photoURL(ctx, field)
			case "isActivated":
				return ec.fieldContext_User_isActivated(ctx, field)
			case "isAdmin":
				return ec.fieldContext_User_isAdmin(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type User", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_updateUser_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var mutationImplementors = []string{"Mutation"}

func (ec *executionContext) _Mutation(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, mutationImplementors)
	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "Mutation",
	})

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		innerCtx := graphql.WithRootFieldContext(ctx, &graphql.RootFieldContext{
			Object: field.Name,
			Field:  field,
		})

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Mutation")
		case "createQuest":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_createQuest(ctx, field)
			})
		case "updateBeacon":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_updateBeacon(ctx, field)
			})
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "updateHost":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_updateHost(ctx, field)
			})
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "createTag":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_createTag(ctx, field)
			})
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "updateTag":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_updateTag(ctx, field)
			})
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "claimTasks":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_claimTasks(ctx, field)
			})
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "submitTaskResult":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_submitTaskResult(ctx, field)
			})
		case "createTome":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_createTome(ctx, field)
			})
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "updateUser":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_updateUser(ctx, field)
			})
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

// endregion ***************************** type.gotpl *****************************
