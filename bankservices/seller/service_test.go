package seller

import (
	"errors"
	"reflect"
	"testing"

	"github.com/luistm/testkit"

	"github.com/luistm/banksaurus/banklib"
	"github.com/luistm/banksaurus/banklib/seller"
	"github.com/luistm/banksaurus/bankservices"
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
			output:     banklib.ErrRepositoryUndefined,
			withMock:   false,
			mockInput:  nil,
			mockOutput: nil,
		},
		{
			name:       "Returns error if s is empty string",
			input:      "",
			output:     banklib.ErrBadInput,
			withMock:   false,
			mockInput:  nil,
			mockOutput: nil,
		},
		{
			name:       "Returns error on Sellers error",
			input:      s,
			output:     &banklib.ErrRepository{Msg: "test Error"},
			withMock:   true,
			mockInput:  seller.New(s, ""), // &Command{slug: s},
			mockOutput: errors.New("test Error"),
		},
		{
			name:       "Returns s entity created",
			input:      s,
			output:     nil,
			withMock:   true,
			mockInput:  seller.New(s, ""), //&Command{slug: s},
			mockOutput: nil,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		i := &Service{}
		var m *banklib.RepositoryMock
		if tc.withMock {
			m = new(banklib.RepositoryMock)
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
			sellerName: "Command Name",
			output:     banklib.ErrBadInput,
		},
		{
			name:       "Returns error if seller name is null",
			slug:       "Command Slug",
			sellerName: "",
			output:     banklib.ErrBadInput,
		},
		{
			name:       "Returns error if Sellers undefined",
			slug:       "Command Slug",
			sellerName: "Command Name",
			output:     banklib.ErrRepositoryUndefined,
		},
		{
			name:       "Returns error if Sellers fails",
			slug:       "Command Slug",
			sellerName: "Command Name",
			output:     &banklib.ErrRepository{Msg: "test Error"},
			withMock:   true,
			mockInput:  seller.New("Command Slug", "Command Name"),
			mockOutput: errors.New("test Error"),
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		i := &Service{}
		var m *banklib.RepositoryMock
		if tc.withMock {
			m = new(banklib.RepositoryMock)
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

	presenterMock := new(bankservices.PresenterMock)
	presenterMock.On(
		"Present",
		[]banklib.Entity{seller.New("", ""), seller.New("", "")},
	).Return(nil)

	testCases := []struct {
		name       string
		output     error
		withMock   bool
		mockOutput []interface{}
	}{
		{
			name:       "Returns error if Sellers is undefined",
			output:     banklib.ErrRepositoryUndefined,
			withMock:   false,
			mockOutput: nil,
		},
		{
			name:       "Returns error on Sellers error",
			output:     &banklib.ErrRepository{Msg: "test Error"},
			withMock:   true,
			mockOutput: []interface{}{[]banklib.Entity{}, errors.New("test Error")},
		},
		{
			name:       "Returns seller entities",
			output:     nil,
			withMock:   true,
			mockOutput: []interface{}{[]banklib.Entity{seller.New("", ""), seller.New("", "")}, nil},
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		i := Service{presenter: presenterMock}
		var repositoryMock *banklib.RepositoryMock
		if tc.withMock {
			repositoryMock = new(banklib.RepositoryMock)
			repositoryMock.On("GetAll").Return(tc.mockOutput...)
			i.repository = repositoryMock
		}

		err := i.GetAll()

		if tc.withMock {
			repositoryMock.AssertExpectations(t)
		}
		testkit.AssertEqual(t, tc.output, err)
	}

	repositoryMock := new(banklib.RepositoryMock)
	repositoryMock.On("GetAll").Return([]banklib.Entity{seller.New("", ""), seller.New("", "")}, nil)
	testCasesPresenter := []struct {
		name       string
		output     error
		withMock   bool
		mockInput  []banklib.Entity
		mockOutput error
	}{
		{
			name:   "Returns error if presenter is not defined",
			output: banklib.ErrPresenterUndefined,
		},
		{
			name:       "Handles presenter error",
			output:     &banklib.ErrPresenter{Msg: "test error"},
			withMock:   true,
			mockInput:  []banklib.Entity{seller.New("", ""), seller.New("", "")},
			mockOutput: errors.New("test error"),
		},
		{
			name:       "Handles presenter success",
			output:     nil,
			withMock:   true,
			mockInput:  []banklib.Entity{seller.New("", ""), seller.New("", "")},
			mockOutput: nil,
		},
	}

	for _, tc := range testCasesPresenter {
		t.Log(tc.name)
		i := Service{repository: repositoryMock}
		var presenterMock *bankservices.PresenterMock
		if tc.withMock {
			presenterMock = new(bankservices.PresenterMock)
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
