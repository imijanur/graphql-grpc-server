// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testUserContacts(t *testing.T) {
	t.Parallel()

	query := UserContacts()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testUserContactsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserContact{}
	if err = randomize.Struct(seed, o, userContactDBTypes, true, userContactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserContact struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserContacts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserContactsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserContact{}
	if err = randomize.Struct(seed, o, userContactDBTypes, true, userContactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserContact struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := UserContacts().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserContacts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserContactsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserContact{}
	if err = randomize.Struct(seed, o, userContactDBTypes, true, userContactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserContact struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserContactSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserContacts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserContactsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserContact{}
	if err = randomize.Struct(seed, o, userContactDBTypes, true, userContactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserContact struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := UserContactExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if UserContact exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UserContactExists to return true, but got false.")
	}
}

func testUserContactsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserContact{}
	if err = randomize.Struct(seed, o, userContactDBTypes, true, userContactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserContact struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	userContactFound, err := FindUserContact(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if userContactFound == nil {
		t.Error("want a record, got nil")
	}
}

func testUserContactsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserContact{}
	if err = randomize.Struct(seed, o, userContactDBTypes, true, userContactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserContact struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = UserContacts().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testUserContactsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserContact{}
	if err = randomize.Struct(seed, o, userContactDBTypes, true, userContactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserContact struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := UserContacts().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUserContactsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userContactOne := &UserContact{}
	userContactTwo := &UserContact{}
	if err = randomize.Struct(seed, userContactOne, userContactDBTypes, false, userContactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserContact struct: %s", err)
	}
	if err = randomize.Struct(seed, userContactTwo, userContactDBTypes, false, userContactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserContact struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userContactOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userContactTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserContacts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUserContactsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	userContactOne := &UserContact{}
	userContactTwo := &UserContact{}
	if err = randomize.Struct(seed, userContactOne, userContactDBTypes, false, userContactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserContact struct: %s", err)
	}
	if err = randomize.Struct(seed, userContactTwo, userContactDBTypes, false, userContactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserContact struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userContactOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userContactTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserContacts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func userContactBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserContact) error {
	*o = UserContact{}
	return nil
}

func userContactAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserContact) error {
	*o = UserContact{}
	return nil
}

func userContactAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *UserContact) error {
	*o = UserContact{}
	return nil
}

func userContactBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserContact) error {
	*o = UserContact{}
	return nil
}

func userContactAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserContact) error {
	*o = UserContact{}
	return nil
}

func userContactBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserContact) error {
	*o = UserContact{}
	return nil
}

func userContactAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserContact) error {
	*o = UserContact{}
	return nil
}

func userContactBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserContact) error {
	*o = UserContact{}
	return nil
}

func userContactAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserContact) error {
	*o = UserContact{}
	return nil
}

func testUserContactsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &UserContact{}
	o := &UserContact{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, userContactDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UserContact object: %s", err)
	}

	AddUserContactHook(boil.BeforeInsertHook, userContactBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	userContactBeforeInsertHooks = []UserContactHook{}

	AddUserContactHook(boil.AfterInsertHook, userContactAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	userContactAfterInsertHooks = []UserContactHook{}

	AddUserContactHook(boil.AfterSelectHook, userContactAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	userContactAfterSelectHooks = []UserContactHook{}

	AddUserContactHook(boil.BeforeUpdateHook, userContactBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	userContactBeforeUpdateHooks = []UserContactHook{}

	AddUserContactHook(boil.AfterUpdateHook, userContactAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	userContactAfterUpdateHooks = []UserContactHook{}

	AddUserContactHook(boil.BeforeDeleteHook, userContactBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	userContactBeforeDeleteHooks = []UserContactHook{}

	AddUserContactHook(boil.AfterDeleteHook, userContactAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	userContactAfterDeleteHooks = []UserContactHook{}

	AddUserContactHook(boil.BeforeUpsertHook, userContactBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	userContactBeforeUpsertHooks = []UserContactHook{}

	AddUserContactHook(boil.AfterUpsertHook, userContactAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	userContactAfterUpsertHooks = []UserContactHook{}
}

func testUserContactsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserContact{}
	if err = randomize.Struct(seed, o, userContactDBTypes, true, userContactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserContact struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserContacts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserContactsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserContact{}
	if err = randomize.Struct(seed, o, userContactDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserContact struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(userContactColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := UserContacts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserContactToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local UserContact
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userContactDBTypes, false, userContactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserContact struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddUserHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *User) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := UserContactSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*UserContact)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testUserContactToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserContact
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userContactDBTypes, false, strmangle.SetComplement(userContactPrimaryKeyColumns, userContactColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserContacts[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testUserContactsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserContact{}
	if err = randomize.Struct(seed, o, userContactDBTypes, true, userContactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserContact struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUserContactsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserContact{}
	if err = randomize.Struct(seed, o, userContactDBTypes, true, userContactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserContact struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserContactSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUserContactsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserContact{}
	if err = randomize.Struct(seed, o, userContactDBTypes, true, userContactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserContact struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserContacts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	userContactDBTypes = map[string]string{`ID`: `int`, `FirstName`: `varchar`, `LastName`: `varchar`, `Phone`: `varchar`, `UserID`: `int`}
	_                  = bytes.MinRead
)

func testUserContactsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(userContactPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(userContactAllColumns) == len(userContactPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserContact{}
	if err = randomize.Struct(seed, o, userContactDBTypes, true, userContactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserContact struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserContacts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userContactDBTypes, true, userContactPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserContact struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testUserContactsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(userContactAllColumns) == len(userContactPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserContact{}
	if err = randomize.Struct(seed, o, userContactDBTypes, true, userContactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserContact struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserContacts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userContactDBTypes, true, userContactPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserContact struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(userContactAllColumns, userContactPrimaryKeyColumns) {
		fields = userContactAllColumns
	} else {
		fields = strmangle.SetComplement(
			userContactAllColumns,
			userContactPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := UserContactSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testUserContactsUpsert(t *testing.T) {
	t.Parallel()

	if len(userContactAllColumns) == len(userContactPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLUserContactUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := UserContact{}
	if err = randomize.Struct(seed, &o, userContactDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UserContact struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserContact: %s", err)
	}

	count, err := UserContacts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, userContactDBTypes, false, userContactPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserContact struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserContact: %s", err)
	}

	count, err = UserContacts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
