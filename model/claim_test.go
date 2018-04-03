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

func testClaims(t *testing.T) {
	t.Parallel()

	query := Claims(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testClaimsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claim := &Claim{}
	if err = randomize.Struct(seed, claim, claimDBTypes, true, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claim.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = claim.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Claims(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClaimsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claim := &Claim{}
	if err = randomize.Struct(seed, claim, claimDBTypes, true, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claim.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Claims(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Claims(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClaimsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claim := &Claim{}
	if err = randomize.Struct(seed, claim, claimDBTypes, true, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claim.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ClaimSlice{claim}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Claims(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testClaimsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claim := &Claim{}
	if err = randomize.Struct(seed, claim, claimDBTypes, true, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claim.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := ClaimExists(tx, claim.ID)
	if err != nil {
		t.Errorf("Unable to check if Claim exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ClaimExistsG to return true, but got false.")
	}
}
func testClaimsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claim := &Claim{}
	if err = randomize.Struct(seed, claim, claimDBTypes, true, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claim.Insert(tx); err != nil {
		t.Error(err)
	}

	claimFound, err := FindClaim(tx, claim.ID)
	if err != nil {
		t.Error(err)
	}

	if claimFound == nil {
		t.Error("want a record, got nil")
	}
}
func testClaimsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claim := &Claim{}
	if err = randomize.Struct(seed, claim, claimDBTypes, true, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claim.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Claims(tx).Bind(claim); err != nil {
		t.Error(err)
	}
}

func testClaimsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claim := &Claim{}
	if err = randomize.Struct(seed, claim, claimDBTypes, true, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claim.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Claims(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testClaimsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claimOne := &Claim{}
	claimTwo := &Claim{}
	if err = randomize.Struct(seed, claimOne, claimDBTypes, false, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}
	if err = randomize.Struct(seed, claimTwo, claimDBTypes, false, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claimOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = claimTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Claims(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testClaimsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	claimOne := &Claim{}
	claimTwo := &Claim{}
	if err = randomize.Struct(seed, claimOne, claimDBTypes, false, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}
	if err = randomize.Struct(seed, claimTwo, claimDBTypes, false, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claimOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = claimTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Claims(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testClaimsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claim := &Claim{}
	if err = randomize.Struct(seed, claim, claimDBTypes, true, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claim.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Claims(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClaimsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claim := &Claim{}
	if err = randomize.Struct(seed, claim, claimDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claim.Insert(tx, claimColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Claims(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClaimToManyPublisherClaims(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Claim
	var b, c Claim

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, claimDBTypes, true, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, claimDBTypes, false, claimColumnsWithDefault...)
	randomize.Struct(seed, &c, claimDBTypes, false, claimColumnsWithDefault...)

	b.PublisherID.Valid = true
	c.PublisherID.Valid = true
	b.PublisherID.String = a.ClaimID
	c.PublisherID.String = a.ClaimID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	claim, err := a.PublisherClaims(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range claim {
		if v.PublisherID.String == b.PublisherID.String {
			bFound = true
		}
		if v.PublisherID.String == c.PublisherID.String {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ClaimSlice{&a}
	if err = a.L.LoadPublisherClaims(tx, false, (*[]*Claim)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PublisherClaims); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.PublisherClaims = nil
	if err = a.L.LoadPublisherClaims(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PublisherClaims); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", claim)
	}
}

func testClaimToManySupportedClaimSupports(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Claim
	var b, c Support

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, claimDBTypes, true, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, supportDBTypes, false, supportColumnsWithDefault...)
	randomize.Struct(seed, &c, supportDBTypes, false, supportColumnsWithDefault...)

	b.SupportedClaimID = a.ClaimID
	c.SupportedClaimID = a.ClaimID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	support, err := a.SupportedClaimSupports(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range support {
		if v.SupportedClaimID == b.SupportedClaimID {
			bFound = true
		}
		if v.SupportedClaimID == c.SupportedClaimID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ClaimSlice{&a}
	if err = a.L.LoadSupportedClaimSupports(tx, false, (*[]*Claim)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.SupportedClaimSupports); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.SupportedClaimSupports = nil
	if err = a.L.LoadSupportedClaimSupports(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.SupportedClaimSupports); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", support)
	}
}

func testClaimToManyAddOpPublisherClaims(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Claim
	var b, c, d, e Claim

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, claimDBTypes, false, strmangle.SetComplement(claimPrimaryKeyColumns, claimColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Claim{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, claimDBTypes, false, strmangle.SetComplement(claimPrimaryKeyColumns, claimColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Claim{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPublisherClaims(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ClaimID != first.PublisherID.String {
			t.Error("foreign key was wrong value", a.ClaimID, first.PublisherID.String)
		}
		if a.ClaimID != second.PublisherID.String {
			t.Error("foreign key was wrong value", a.ClaimID, second.PublisherID.String)
		}

		if first.R.Publisher != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Publisher != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.PublisherClaims[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.PublisherClaims[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.PublisherClaims(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testClaimToManySetOpPublisherClaims(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Claim
	var b, c, d, e Claim

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, claimDBTypes, false, strmangle.SetComplement(claimPrimaryKeyColumns, claimColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Claim{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, claimDBTypes, false, strmangle.SetComplement(claimPrimaryKeyColumns, claimColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.SetPublisherClaims(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.PublisherClaims(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetPublisherClaims(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.PublisherClaims(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.PublisherID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.PublisherID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.ClaimID != d.PublisherID.String {
		t.Error("foreign key was wrong value", a.ClaimID, d.PublisherID.String)
	}
	if a.ClaimID != e.PublisherID.String {
		t.Error("foreign key was wrong value", a.ClaimID, e.PublisherID.String)
	}

	if b.R.Publisher != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Publisher != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Publisher != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Publisher != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.PublisherClaims[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.PublisherClaims[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testClaimToManyRemoveOpPublisherClaims(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Claim
	var b, c, d, e Claim

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, claimDBTypes, false, strmangle.SetComplement(claimPrimaryKeyColumns, claimColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Claim{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, claimDBTypes, false, strmangle.SetComplement(claimPrimaryKeyColumns, claimColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddPublisherClaims(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.PublisherClaims(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemovePublisherClaims(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.PublisherClaims(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.PublisherID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.PublisherID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Publisher != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Publisher != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Publisher != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Publisher != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.PublisherClaims) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.PublisherClaims[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.PublisherClaims[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testClaimToManyAddOpSupportedClaimSupports(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Claim
	var b, c, d, e Support

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, claimDBTypes, false, strmangle.SetComplement(claimPrimaryKeyColumns, claimColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Support{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, supportDBTypes, false, strmangle.SetComplement(supportPrimaryKeyColumns, supportColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Support{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddSupportedClaimSupports(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ClaimID != first.SupportedClaimID {
			t.Error("foreign key was wrong value", a.ClaimID, first.SupportedClaimID)
		}
		if a.ClaimID != second.SupportedClaimID {
			t.Error("foreign key was wrong value", a.ClaimID, second.SupportedClaimID)
		}

		if first.R.SupportedClaim != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.SupportedClaim != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.SupportedClaimSupports[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.SupportedClaimSupports[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.SupportedClaimSupports(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testClaimToOneTransactionUsingTransactionByHash(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Claim
	var foreign Transaction

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, claimDBTypes, true, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, transactionDBTypes, false, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	local.TransactionByHashID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.TransactionByHashID.String = foreign.Hash
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TransactionByHash(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.Hash != foreign.Hash {
		t.Errorf("want: %v, got %v", foreign.Hash, check.Hash)
	}

	slice := ClaimSlice{&local}
	if err = local.L.LoadTransactionByHash(tx, false, (*[]*Claim)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.TransactionByHash == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TransactionByHash = nil
	if err = local.L.LoadTransactionByHash(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TransactionByHash == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testClaimToOneClaimUsingPublisher(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Claim
	var foreign Claim

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, claimDBTypes, true, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, claimDBTypes, false, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	local.PublisherID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.PublisherID.String = foreign.ClaimID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Publisher(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ClaimID != foreign.ClaimID {
		t.Errorf("want: %v, got %v", foreign.ClaimID, check.ClaimID)
	}

	slice := ClaimSlice{&local}
	if err = local.L.LoadPublisher(tx, false, (*[]*Claim)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.Publisher == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Publisher = nil
	if err = local.L.LoadPublisher(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Publisher == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testClaimToOneSetOpTransactionUsingTransactionByHash(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Claim
	var b, c Transaction

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, claimDBTypes, false, strmangle.SetComplement(claimPrimaryKeyColumns, claimColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Transaction{&b, &c} {
		err = a.SetTransactionByHash(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TransactionByHash != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.TransactionByHashClaims[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.TransactionByHashID.String != x.Hash {
			t.Error("foreign key was wrong value", a.TransactionByHashID.String)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TransactionByHashID.String))
		reflect.Indirect(reflect.ValueOf(&a.TransactionByHashID.String)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.TransactionByHashID.String != x.Hash {
			t.Error("foreign key was wrong value", a.TransactionByHashID.String, x.Hash)
		}
	}
}

func testClaimToOneRemoveOpTransactionUsingTransactionByHash(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Claim
	var b Transaction

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, claimDBTypes, false, strmangle.SetComplement(claimPrimaryKeyColumns, claimColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetTransactionByHash(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveTransactionByHash(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.TransactionByHash(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.TransactionByHash != nil {
		t.Error("R struct entry should be nil")
	}

	if a.TransactionByHashID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.TransactionByHashClaims) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testClaimToOneSetOpClaimUsingPublisher(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Claim
	var b, c Claim

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, claimDBTypes, false, strmangle.SetComplement(claimPrimaryKeyColumns, claimColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, claimDBTypes, false, strmangle.SetComplement(claimPrimaryKeyColumns, claimColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, claimDBTypes, false, strmangle.SetComplement(claimPrimaryKeyColumns, claimColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Claim{&b, &c} {
		err = a.SetPublisher(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Publisher != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.PublisherClaims[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PublisherID.String != x.ClaimID {
			t.Error("foreign key was wrong value", a.PublisherID.String)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PublisherID.String))
		reflect.Indirect(reflect.ValueOf(&a.PublisherID.String)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PublisherID.String != x.ClaimID {
			t.Error("foreign key was wrong value", a.PublisherID.String, x.ClaimID)
		}
	}
}

func testClaimToOneRemoveOpClaimUsingPublisher(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Claim
	var b Claim

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, claimDBTypes, false, strmangle.SetComplement(claimPrimaryKeyColumns, claimColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, claimDBTypes, false, strmangle.SetComplement(claimPrimaryKeyColumns, claimColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetPublisher(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemovePublisher(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Publisher(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Publisher != nil {
		t.Error("R struct entry should be nil")
	}

	if a.PublisherID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.PublisherClaims) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testClaimsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claim := &Claim{}
	if err = randomize.Struct(seed, claim, claimDBTypes, true, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claim.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = claim.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testClaimsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claim := &Claim{}
	if err = randomize.Struct(seed, claim, claimDBTypes, true, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claim.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ClaimSlice{claim}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testClaimsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claim := &Claim{}
	if err = randomize.Struct(seed, claim, claimDBTypes, true, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claim.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Claims(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	claimDBTypes = map[string]string{`Author`: `varchar`, `BidState`: `varchar`, `Certificate`: `text`, `ClaimID`: `char`, `ClaimType`: `tinyint`, `ContentType`: `varchar`, `Created`: `datetime`, `Description`: `mediumtext`, `Fee`: `double`, `FeeCurrency`: `char`, `ID`: `bigint`, `IsFiltered`: `tinyint`, `IsNSFW`: `tinyint`, `Language`: `varchar`, `Modified`: `datetime`, `Name`: `varchar`, `PublisherID`: `char`, `PublisherSig`: `varchar`, `SDHash`: `varchar`, `ThumbnailURL`: `text`, `Title`: `text`, `TransactionByHashID`: `varchar`, `TransactionTime`: `bigint`, `ValueAsHex`: `mediumtext`, `ValueAsJSON`: `mediumtext`, `Version`: `varchar`, `Vout`: `int`}
	_            = bytes.MinRead
)

func testClaimsUpdate(t *testing.T) {
	t.Parallel()

	if len(claimColumns) == len(claimPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	claim := &Claim{}
	if err = randomize.Struct(seed, claim, claimDBTypes, true, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claim.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Claims(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, claim, claimDBTypes, true, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	if err = claim.Update(tx); err != nil {
		t.Error(err)
	}
}

func testClaimsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(claimColumns) == len(claimPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	claim := &Claim{}
	if err = randomize.Struct(seed, claim, claimDBTypes, true, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claim.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Claims(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, claim, claimDBTypes, true, claimPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(claimColumns, claimPrimaryKeyColumns) {
		fields = claimColumns
	} else {
		fields = strmangle.SetComplement(
			claimColumns,
			claimPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(claim))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := ClaimSlice{claim}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testClaimsUpsert(t *testing.T) {
	t.Parallel()

	if len(claimColumns) == len(claimPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	claim := Claim{}
	if err = randomize.Struct(seed, &claim, claimDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claim.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Claim: %s", err)
	}

	count, err := Claims(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &claim, claimDBTypes, false, claimPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	if err = claim.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Claim: %s", err)
	}

	count, err = Claims(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
