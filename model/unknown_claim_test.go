// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

func testUnknownClaims(t *testing.T) {
	t.Parallel()

	query := UnknownClaims(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testUnknownClaimsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	unknownClaim := &UnknownClaim{}
	if err = randomize.Struct(seed, unknownClaim, unknownClaimDBTypes, true, unknownClaimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UnknownClaim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = unknownClaim.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = unknownClaim.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := UnknownClaims(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUnknownClaimsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	unknownClaim := &UnknownClaim{}
	if err = randomize.Struct(seed, unknownClaim, unknownClaimDBTypes, true, unknownClaimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UnknownClaim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = unknownClaim.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = UnknownClaims(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := UnknownClaims(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUnknownClaimsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	unknownClaim := &UnknownClaim{}
	if err = randomize.Struct(seed, unknownClaim, unknownClaimDBTypes, true, unknownClaimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UnknownClaim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = unknownClaim.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := UnknownClaimSlice{unknownClaim}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := UnknownClaims(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testUnknownClaimsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	unknownClaim := &UnknownClaim{}
	if err = randomize.Struct(seed, unknownClaim, unknownClaimDBTypes, true, unknownClaimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UnknownClaim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = unknownClaim.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := UnknownClaimExists(tx, unknownClaim.ID)
	if err != nil {
		t.Errorf("Unable to check if UnknownClaim exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UnknownClaimExistsG to return true, but got false.")
	}
}
func testUnknownClaimsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	unknownClaim := &UnknownClaim{}
	if err = randomize.Struct(seed, unknownClaim, unknownClaimDBTypes, true, unknownClaimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UnknownClaim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = unknownClaim.Insert(tx); err != nil {
		t.Error(err)
	}

	unknownClaimFound, err := FindUnknownClaim(tx, unknownClaim.ID)
	if err != nil {
		t.Error(err)
	}

	if unknownClaimFound == nil {
		t.Error("want a record, got nil")
	}
}
func testUnknownClaimsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	unknownClaim := &UnknownClaim{}
	if err = randomize.Struct(seed, unknownClaim, unknownClaimDBTypes, true, unknownClaimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UnknownClaim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = unknownClaim.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = UnknownClaims(tx).Bind(unknownClaim); err != nil {
		t.Error(err)
	}
}

func testUnknownClaimsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	unknownClaim := &UnknownClaim{}
	if err = randomize.Struct(seed, unknownClaim, unknownClaimDBTypes, true, unknownClaimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UnknownClaim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = unknownClaim.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := UnknownClaims(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUnknownClaimsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	unknownClaimOne := &UnknownClaim{}
	unknownClaimTwo := &UnknownClaim{}
	if err = randomize.Struct(seed, unknownClaimOne, unknownClaimDBTypes, false, unknownClaimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UnknownClaim struct: %s", err)
	}
	if err = randomize.Struct(seed, unknownClaimTwo, unknownClaimDBTypes, false, unknownClaimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UnknownClaim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = unknownClaimOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = unknownClaimTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := UnknownClaims(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUnknownClaimsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	unknownClaimOne := &UnknownClaim{}
	unknownClaimTwo := &UnknownClaim{}
	if err = randomize.Struct(seed, unknownClaimOne, unknownClaimDBTypes, false, unknownClaimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UnknownClaim struct: %s", err)
	}
	if err = randomize.Struct(seed, unknownClaimTwo, unknownClaimDBTypes, false, unknownClaimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UnknownClaim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = unknownClaimOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = unknownClaimTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := UnknownClaims(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testUnknownClaimsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	unknownClaim := &UnknownClaim{}
	if err = randomize.Struct(seed, unknownClaim, unknownClaimDBTypes, true, unknownClaimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UnknownClaim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = unknownClaim.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := UnknownClaims(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUnknownClaimsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	unknownClaim := &UnknownClaim{}
	if err = randomize.Struct(seed, unknownClaim, unknownClaimDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UnknownClaim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = unknownClaim.Insert(tx, unknownClaimColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := UnknownClaims(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUnknownClaimToOneOutputUsingOutput(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local UnknownClaim
	var foreign Output

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, unknownClaimDBTypes, false, unknownClaimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UnknownClaim struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, outputDBTypes, false, outputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.OutputID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Output(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := UnknownClaimSlice{&local}
	if err = local.L.LoadOutput(tx, false, (*[]*UnknownClaim)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.Output == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Output = nil
	if err = local.L.LoadOutput(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Output == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testUnknownClaimToOneSetOpOutputUsingOutput(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a UnknownClaim
	var b, c Output

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, unknownClaimDBTypes, false, strmangle.SetComplement(unknownClaimPrimaryKeyColumns, unknownClaimColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, outputDBTypes, false, strmangle.SetComplement(outputPrimaryKeyColumns, outputColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, outputDBTypes, false, strmangle.SetComplement(outputPrimaryKeyColumns, outputColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Output{&b, &c} {
		err = a.SetOutput(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Output != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UnknownClaims[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.OutputID != x.ID {
			t.Error("foreign key was wrong value", a.OutputID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.OutputID))
		reflect.Indirect(reflect.ValueOf(&a.OutputID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.OutputID != x.ID {
			t.Error("foreign key was wrong value", a.OutputID, x.ID)
		}
	}
}
func testUnknownClaimsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	unknownClaim := &UnknownClaim{}
	if err = randomize.Struct(seed, unknownClaim, unknownClaimDBTypes, true, unknownClaimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UnknownClaim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = unknownClaim.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = unknownClaim.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testUnknownClaimsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	unknownClaim := &UnknownClaim{}
	if err = randomize.Struct(seed, unknownClaim, unknownClaimDBTypes, true, unknownClaimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UnknownClaim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = unknownClaim.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := UnknownClaimSlice{unknownClaim}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testUnknownClaimsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	unknownClaim := &UnknownClaim{}
	if err = randomize.Struct(seed, unknownClaim, unknownClaimDBTypes, true, unknownClaimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UnknownClaim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = unknownClaim.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := UnknownClaims(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	unknownClaimDBTypes = map[string]string{`BlockHash`: `varchar`, `ClaimID`: `varchar`, `ID`: `bigint`, `IsUpdate`: `tinyint`, `Name`: `varchar`, `OutputID`: `bigint`, `TransactionHash`: `varchar`, `ValueAsHex`: `mediumtext`, `ValueAsJSON`: `mediumtext`, `Vout`: `int`}
	_                   = bytes.MinRead
)

func testUnknownClaimsUpdate(t *testing.T) {
	t.Parallel()

	if len(unknownClaimColumns) == len(unknownClaimPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	unknownClaim := &UnknownClaim{}
	if err = randomize.Struct(seed, unknownClaim, unknownClaimDBTypes, true, unknownClaimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UnknownClaim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = unknownClaim.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := UnknownClaims(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, unknownClaim, unknownClaimDBTypes, true, unknownClaimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UnknownClaim struct: %s", err)
	}

	if err = unknownClaim.Update(tx); err != nil {
		t.Error(err)
	}
}

func testUnknownClaimsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(unknownClaimColumns) == len(unknownClaimPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	unknownClaim := &UnknownClaim{}
	if err = randomize.Struct(seed, unknownClaim, unknownClaimDBTypes, true, unknownClaimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UnknownClaim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = unknownClaim.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := UnknownClaims(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, unknownClaim, unknownClaimDBTypes, true, unknownClaimPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UnknownClaim struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(unknownClaimColumns, unknownClaimPrimaryKeyColumns) {
		fields = unknownClaimColumns
	} else {
		fields = strmangle.SetComplement(
			unknownClaimColumns,
			unknownClaimPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(unknownClaim))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := UnknownClaimSlice{unknownClaim}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testUnknownClaimsUpsert(t *testing.T) {
	t.Parallel()

	if len(unknownClaimColumns) == len(unknownClaimPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	unknownClaim := UnknownClaim{}
	if err = randomize.Struct(seed, &unknownClaim, unknownClaimDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UnknownClaim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = unknownClaim.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert UnknownClaim: %s", err)
	}

	count, err := UnknownClaims(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &unknownClaim, unknownClaimDBTypes, false, unknownClaimPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UnknownClaim struct: %s", err)
	}

	if err = unknownClaim.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert UnknownClaim: %s", err)
	}

	count, err = UnknownClaims(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
