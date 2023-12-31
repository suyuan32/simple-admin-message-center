// Code generated by ent, DO NOT EDIT.

package emailprovider

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the emailprovider type in the database.
	Label = "email_provider"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAuthType holds the string denoting the auth_type field in the database.
	FieldAuthType = "auth_type"
	// FieldEmailAddr holds the string denoting the email_addr field in the database.
	FieldEmailAddr = "email_addr"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldHostName holds the string denoting the host_name field in the database.
	FieldHostName = "host_name"
	// FieldIdentify holds the string denoting the identify field in the database.
	FieldIdentify = "identify"
	// FieldSecret holds the string denoting the secret field in the database.
	FieldSecret = "secret"
	// FieldPort holds the string denoting the port field in the database.
	FieldPort = "port"
	// FieldTLS holds the string denoting the tls field in the database.
	FieldTLS = "tls"
	// FieldIsDefault holds the string denoting the is_default field in the database.
	FieldIsDefault = "is_default"
	// Table holds the table name of the emailprovider in the database.
	Table = "mcms_email_providers"
)

// Columns holds all SQL columns for emailprovider fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldAuthType,
	FieldEmailAddr,
	FieldPassword,
	FieldHostName,
	FieldIdentify,
	FieldSecret,
	FieldPort,
	FieldTLS,
	FieldIsDefault,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultTLS holds the default value on creation for the "tls" field.
	DefaultTLS bool
	// DefaultIsDefault holds the default value on creation for the "is_default" field.
	DefaultIsDefault bool
)

// OrderOption defines the ordering options for the EmailProvider queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByAuthType orders the results by the auth_type field.
func ByAuthType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthType, opts...).ToFunc()
}

// ByEmailAddr orders the results by the email_addr field.
func ByEmailAddr(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmailAddr, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByHostName orders the results by the host_name field.
func ByHostName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHostName, opts...).ToFunc()
}

// ByIdentify orders the results by the identify field.
func ByIdentify(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIdentify, opts...).ToFunc()
}

// BySecret orders the results by the secret field.
func BySecret(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSecret, opts...).ToFunc()
}

// ByPort orders the results by the port field.
func ByPort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPort, opts...).ToFunc()
}

// ByTLS orders the results by the tls field.
func ByTLS(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTLS, opts...).ToFunc()
}

// ByIsDefault orders the results by the is_default field.
func ByIsDefault(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsDefault, opts...).ToFunc()
}
