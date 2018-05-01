package bank

import (
	"errors"
	"reflect"
	"testing"

	"github.com/luistm/banksaurus/elib/testkit"

	"github.com/luistm/banksaurus/lib"
	"github.com/luistm/banksaurus/lib/seller"
)

func TestUnitInteractorCreate(t *testing.T) {

	var s = "TestDescription"

	testCases := []struct {
		name       string
		input      string
		output     error
		withMock   bool
		mockInput  *seller.Seller
		mockOutput error
	}{
		{
			name:       "Returns error if Sellers is not defined",
			input:      s,
			output:     lib.ErrRepositoryUndefined,
			withMock:   false,
			mockInput:  nil,
			mockOutput: nil,
		},
		{
			name:       "Returns error if s is empty string",
			input:      "",
			output:     lib.ErrBadInput,
			withMock:   false,
			mockInput:  nil,
			mockOutput: nil,
		},
		{
			name:       "Returns error on Sellers error",
			input:      s,
			output:     &lib.ErrRepository{Msg: "test Error"},
			withMock:   true,
			mockInput:  seller.New(s, ""), // &Seller{slug: s},
			mockOutput: errors.New("test Error"),
		},
		{
			name:       "Returns s entity created",
			input:      s,
			output:     nil,
			withMock:   true,
			mockInput:  seller.New(s, ""), //&Seller{slug: s},
			mockOutput: nil,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		i := &SellerInteractor{}
		var m *lib.RepositoryMock
		if tc.withMock {
			m = new(lib.RepositoryMock)
			m.On("Save", tc.mockInput).Return(tc.mockOutput)
			i.repository = m
		}

		err := i.Create(tc.input)

		if tc.withMock {
			m.AssertExpectations(t)
		}
		testkit.AssertEqual(t, tc.output, err)
	}
}

func TestUnitInteractorUpdate(t *testing.T) {

	testCases := []struct {
		name       string
		slug       string
		sellerName string
		output     error
		withMock   bool
		mockInput  *seller.Seller
		mockOutput error
	}{
		{
			name:       "Returns error if seller ID is null",
			slug:       "",
			sellerName: "Seller Name",
			output:     lib.ErrBadInput,
		},
		{
			name:       "Returns error if seller name is null",
			slug:       "Seller Slug",
			sellerName: "",
			output:     lib.ErrBadInput,
		},
		{
			name:       "Returns error if Sellers undefined",
			slug:       "Seller Slug",
			sellerName: "Seller Name",
			output:     lib.ErrRepositoryUndefined,
		},
		{
			name:       "Returns error if Sellers fails",
			slug:       "Seller Slug",
			sellerName: "Seller Name",
			output:     &lib.ErrRepository{Msg: "test Error"},
			withMock:   true,
			mockInput:  seller.New("Seller Slug", "Seller Name"),
			mockOutput: errors.New("test Error"),
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		i := &SellerInteractor{}
		var m *lib.RepositoryMock
		if tc.withMock {
			m = new(lib.RepositoryMock)
			m.On("Save", tc.mockInput).Return(tc.mockOutput)
			i.repository = m
		}

		err := i.Update(tc.slug, tc.sellerName)

		if tc.withMock {
			m.AssertExpectations(t)
		}
		if !reflect.DeepEqual(tc.output, err) {
			t.Errorf("Expected '%v', got '%v'", tc.output, err)
		}
	}
}

func TestUnitInteractorGetAll(t *testing.T) {

	presenterMock := new(PresenterMock)
	presenterMock.On(
		"Present",
		[]lib.Entity{seller.New("", ""), seller.New("", "")},
	).Return(nil)

	testCases := []struct {
		name       string
		output     error
		withMock   bool
		mockOutput []interface{}
	}{
		{
			name:       "Returns error if Sellers is undefined",
			output:     lib.ErrRepositoryUndefined,
			withMock:   false,
			mockOutput: nil,
		},
		{
			name:       "Returns error on Sellers error",
			output:     &lib.ErrRepository{Msg: "test Error"},
			withMock:   true,
			mockOutput: []interface{}{[]lib.Entity{}, errors.New("test Error")},
		},
		{
			name:       "Returns seller entities",
			output:     nil,
			withMock:   true,
			mockOutput: []interface{}{[]lib.Entity{seller.New("", ""), seller.New("", "")}, nil},
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		i := SellerInteractor{presenter: presenterMock}
		var repositoryMock *lib.RepositoryMock
		if tc.withMock {
			repositoryMock = new(lib.RepositoryMock)
			repositoryMock.On("GetAll").Return(tc.mockOutput...)
			i.repository = repositoryMock
		}

		err := i.GetAll()

		if tc.withMock {
			repositoryMock.AssertExpectations(t)
		}
		testkit.AssertEqual(t, tc.output, err)
	}

	repositoryMock := new(lib.RepositoryMock)
	repositoryMock.On("GetAll").Return([]lib.Entity{seller.New("", ""), seller.New("", "")}, nil)
	testCasesPresenter := []struct {
		name       string
		output     error
		withMock   bool
		mockInput  []lib.Entity
		mockOutput error
	}{
		{
			name:   "Returns error if presenter is not defined",
			output: lib.ErrPresenterUndefined,
		},
		{
			name:       "Handles presenter error",
			output:     &lib.ErrPresenter{Msg: "test error"},
			withMock:   true,
			mockInput:  []lib.Entity{seller.New("", ""), seller.New("", "")},
			mockOutput: errors.New("test error"),
		},
		{
			name:       "Handles presenter success",
			output:     nil,
			withMock:   true,
			mockInput:  []lib.Entity{seller.New("", ""), seller.New("", "")},
			mockOutput: nil,
		},
	}

	for _, tc := range testCasesPresenter {
		t.Log(tc.name)
		i := SellerInteractor{repository: repositoryMock}
		var presenterMock *PresenterMock
		if tc.withMock {
			presenterMock = new(PresenterMock)
			presenterMock.On("Present", tc.mockInput).Return(tc.mockOutput)
			i.presenter = presenterMock
		}

		err := i.GetAll()

		if tc.withMock {
			presenterMock.AssertExpectations(t)
		}
		testkit.AssertEqual(t, tc.output, err)
	}
}
