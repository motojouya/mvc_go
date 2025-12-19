package core_test

// FIXME `stretchr/testify/assert` から脱却
import (
	"github.com/motojouya/ddd_go/domain/basic/core"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestInvalidArgumentError(t *testing.T) {
	var name = "TestInvalidArgumentError"
	var value = "invalid_value"
	var message = "This is a test invalid argument error"
	var httpStatus uint = 400

	var err = core.NewInvalidArgumentError(name, value, message)

	assert.Equal(t, name, err.Name)
	assert.Equal(t, value, err.Value)
	assert.Equal(t, message, err.Unwrap().Error())
	assert.Equal(t, message+", name: "+name+", value: "+value, err.Error())
	assert.Equal(t, httpStatus, err.HttpStatus())

	t.Logf("error: %s", err)
	t.Logf("error message: %s", err.Error())
}

func TestNewRangeError(t *testing.T) {
	var name = "TestRangeError"
	var value = 100
	var min = 10
	var max = 20
	var message = "This is a test range error"
	var httpStatus uint = 400

	var err = core.NewRangeError(name, value, min, max, message)

	assert.Equal(t, name, err.Name)
	assert.Equal(t, value, err.Value)
	assert.Equal(t, min, err.Min)
	assert.Equal(t, max, err.Max)
	assert.Equal(t, message, err.Unwrap().Error())
	assert.Equal(t, message+", name: "+name+", value: "+strconv.Itoa(value)+", min: "+strconv.Itoa(min)+", max: "+strconv.Itoa(max), err.Error())
	assert.Equal(t, httpStatus, err.HttpStatus())

	t.Logf("error: %s", err.Error())
	t.Logf("error.Name: %s", err.Name)
}

func TestNewAuthenticationError(t *testing.T) {
	var userIdentifier = "TestUserIdentifier"
	var message = "This is a test system config error"
	var httpStatus uint = 401

	var err = core.NewAuthenticationError(userIdentifier, message)

	assert.Equal(t, userIdentifier, err.UserIdentifier)
	assert.Equal(t, message, err.Unwrap().Error())
	assert.Equal(t, message+", userIdentifier: "+userIdentifier, err.Error())
	assert.Equal(t, httpStatus, err.HttpStatus())

	t.Logf("error: %s", err.Error())
	t.Logf("error.Name: %s", err.UserIdentifier)
}

func TestNewNotFoundError(t *testing.T) {
	var table = "TestUser"
	var key = "TestUserId"
	var value = "TestValue"
	var keys = map[string]string{key: value}
	var message = "This is a test range error"
	var httpStatus uint = 400

	var err = core.NewNotFoundError(table, keys, message)

	assert.Equal(t, table, err.Table)
	var val, exist = err.Keys[key]
	assert.True(t, exist)
	assert.Equal(t, value, val)
	assert.Equal(t, message, err.Unwrap().Error())
	assert.Equal(t, message+", table: "+table+", keys: {"+key+": "+value+", }", err.Error())
	assert.Equal(t, httpStatus, err.HttpStatus())

	t.Logf("error: %s", err.Error())
	t.Logf("error.Table: %s", err.Table)
}

func TestNewDuplicateError(t *testing.T) {
	var table = "TestUser"
	var key = "TestUserId"
	var value = "TestValue"
	var keys = map[string]string{key: value}
	var message = "This is a test range error"
	var httpStatus uint = 400

	var err = core.NewDuplicateError(table, keys, message)

	assert.Equal(t, table, err.Table)
	var val, exist = err.Keys[key]
	assert.True(t, exist)
	assert.Equal(t, value, val)
	assert.Equal(t, message, err.Unwrap().Error())
	assert.Equal(t, message+", table: "+table+", keys: {"+key+": "+value+", }", err.Error())
	assert.Equal(t, httpStatus, err.HttpStatus())

	t.Logf("error: %s", err.Error())
	t.Logf("error.Table: %s", err.Table)
}

func TestNewNilError(t *testing.T) {
	var name = "TestNilError"
	var message = "This is a test nil error"
	var httpStatus uint = 400

	var err = core.NewNilError(name, message)

	assert.Equal(t, name, err.Name)
	assert.Equal(t, message, err.Unwrap().Error())
	assert.Equal(t, message+", name: "+name, err.Error())
	assert.Equal(t, httpStatus, err.HttpStatus())

	t.Logf("error: %s", err.Error())
	t.Logf("error.Name: %s", err.Name)
	t.Logf("error.HttpStatus: %d", err.HttpStatus())
}

func TestNewSystemConfigError(t *testing.T) {
	var name = "TestSystemConfigError"
	var message = "This is a test system config error"
	var httpStatus uint = 500

	var err = core.NewSystemConfigError(name, message)

	assert.Equal(t, name, err.Name)
	assert.Equal(t, message, err.Unwrap().Error())
	assert.Equal(t, message+", name: "+name, err.Error())
	assert.Equal(t, httpStatus, err.HttpStatus())

	t.Logf("error: %s", err.Error())
	t.Logf("error.Name: %s", err.Name)
}

func TestNewPropertyError(t *testing.T) {
	var name = "TestNilError"
	var message = "This is a test nil error"
	var err = core.NewNilError(name, message)

	var prop = "TestPath"
	var httpStatus uint = 210

	var propertyError = core.NewPropertyError(prop, httpStatus, err)

	assert.Equal(t, prop, propertyError.Property)
	assert.Equal(t, httpStatus, propertyError.HttpStatusCode)
	assert.Equal(t, message+", name: "+name, propertyError.Unwrap().Error())
	assert.Equal(t, message+", name: "+name+", property: "+prop+", httpStatus: "+strconv.Itoa(int(httpStatus)), propertyError.Error())

	t.Logf("error: %s", propertyError.Error())
	t.Logf("error.Property: %s", propertyError.Property)
}

func TestCreatePropertyError(t *testing.T) {
	var name = "TestNilError"
	var message = "This is a test nil error"
	var err = core.NewNilError(name, message)

	var prop = "TestPath"
	var httpStatus uint = 400

	var propertyError = core.CreatePropertyError(prop, err)

	assert.Equal(t, prop, propertyError.Property)
	assert.Equal(t, httpStatus, propertyError.HttpStatusCode)
	assert.Equal(t, message+", name: "+name, propertyError.Unwrap().Error())
	assert.Equal(t, message+", name: "+name+", property: "+prop+", httpStatus: "+strconv.Itoa(int(httpStatus)), propertyError.Error())

	t.Logf("error: %s", propertyError.Error())
	t.Logf("error.Property: %s", propertyError.Property)
}

func TestPropertyErrorAdd(t *testing.T) {
	var name = "TestNilError"
	var message = "This is a test nil error"
	var err = core.NewNilError(name, message)

	var prop = "TestPath"
	var httpStatus uint = 210

	var propertyError = core.NewPropertyError(prop, httpStatus, err)
	var path = "additional"
	var added = propertyError.Add(path)

	assert.Equal(t, path+"."+prop, added.Property)
	assert.Equal(t, httpStatus, added.HttpStatusCode)
	assert.Equal(t, message+", name: "+name, added.Unwrap().Error())
	assert.Equal(t, message+", name: "+name+", property: "+path+"."+prop+", httpStatus: "+strconv.Itoa(int(httpStatus)), added.Error())

	t.Logf("error: %s", added.Error())
	t.Logf("error.Property: %s", added.Property)
}

func TestPropertyErrorChange(t *testing.T) {
	var name = "TestNilError"
	var message = "This is a test nil error"
	var err = core.NewNilError(name, message)

	var prop = "TestPath"
	var httpStatus uint = 210

	var propertyError = core.NewPropertyError(prop, httpStatus, err)
	var path = "additional"
	var changedStatus uint = 220
	var added = propertyError.Change(path, changedStatus)

	assert.Equal(t, path+"."+prop, added.Property)
	assert.Equal(t, changedStatus, added.HttpStatusCode)
	assert.Equal(t, message+", name: "+name, added.Unwrap().Error())
	assert.Equal(t, message+", name: "+name+", property: "+path+"."+prop+", httpStatus: "+strconv.Itoa(int(changedStatus)), added.Error())

	t.Logf("error: %s", added.Error())
	t.Logf("error.Property: %s", added.Property)
}

func TestAddPropertyError(t *testing.T) {
	var name = "TestNilError"
	var message = "This is a test nil error"
	var err = core.NewNilError(name, message)

	var prop = "TestPath"
	var propertyError = core.AddPropertyError(prop, err)

	var wrapPath = "additional"
	var httpStatus uint = 400
	var wrappedPropertyError = core.AddPropertyError(wrapPath, propertyError)

	assert.Equal(t, wrapPath+"."+prop, wrappedPropertyError.Property)
	assert.Equal(t, httpStatus, wrappedPropertyError.HttpStatusCode)
	assert.Equal(t, message+", name: "+name, wrappedPropertyError.Unwrap().Error())
	assert.Equal(t, message+", name: "+name+", property: "+wrapPath+"."+prop+", httpStatus: 400", wrappedPropertyError.Error())

	t.Logf("error: %s", wrappedPropertyError.Error())
	t.Logf("error.Property: %s", wrappedPropertyError.Property)
}

func TestChangePropertyError(t *testing.T) {
	var name = "TestNilError"
	var message = "This is a test nil error"
	var err = core.NewNilError(name, message)

	var prop = "TestPath"
	var httpStatus uint = 210
	var propertyError = core.ChangePropertyError(prop, err, httpStatus)

	var wrapPath = "additional"
	var wraphttpStatus uint = 210
	var wrappedPropertyError = core.ChangePropertyError(wrapPath, propertyError, wraphttpStatus)

	assert.Equal(t, wrapPath+"."+prop, wrappedPropertyError.Property)
	assert.Equal(t, wraphttpStatus, wrappedPropertyError.HttpStatusCode)
	assert.Equal(t, message+", name: "+name, wrappedPropertyError.Unwrap().Error())
	assert.Equal(t, message+", name: "+name+", property: "+wrapPath+"."+prop+", httpStatus: "+strconv.Itoa(int(wraphttpStatus)), wrappedPropertyError.Error())

	t.Logf("error: %s", wrappedPropertyError.Error())
	t.Logf("error.Property: %s", wrappedPropertyError.Property)
}

func TestAddPropertyErrorNil(t *testing.T) {
	defer func() {
		if rec := recover(); rec != nil {
			t.Logf("Recovered from panic: %v", rec)
		}
	}()

	var prop = "TestPath"
	var _ = core.AddPropertyError(prop, nil)

	t.Error("Expected panic for nil source error, but did not panic")
}

func TestChangePropertyErrorNil(t *testing.T) {
	defer func() {
		if rec := recover(); rec != nil {
			t.Logf("Recovered from panic: %v", rec)
		}
	}()

	var prop = "TestPath"
	var httpStatus uint = 210
	var _ = core.ChangePropertyError(prop, nil, httpStatus)

	t.Error("Expected panic for nil source error, but did not panic")
}
