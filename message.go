package ldapserver

import (
	"context"
	"fmt"

	ldap "github.com/cloudldap/goldap/message"
)

type Message struct {
	*ldap.LDAPMessage
	Client *client
	Cancel context.CancelFunc
	ctx    context.Context
}

func (m *Message) String() string {
	return fmt.Sprintf("MessageId=%d, %s", m.MessageID(), m.ProtocolOpName())
}

func (m *Message) Context() context.Context {
	return m.ctx
}

// Abandon cancel the context, to notify handler's user function to stop any
// running process
func (m *Message) Abandon() {
	m.Cancel()
}

func (m *Message) GetAbandonRequest() ldap.AbandonRequest {
	return m.ProtocolOp().(ldap.AbandonRequest)
}

func (m *Message) GetSearchRequest() ldap.SearchRequest {
	return m.ProtocolOp().(ldap.SearchRequest)
}

func (m *Message) GetBindRequest() ldap.BindRequest {
	return m.ProtocolOp().(ldap.BindRequest)
}

func (m *Message) GetAddRequest() ldap.AddRequest {
	return m.ProtocolOp().(ldap.AddRequest)
}

func (m *Message) GetDeleteRequest() ldap.DelRequest {
	return m.ProtocolOp().(ldap.DelRequest)
}

func (m *Message) GetModifyRequest() ldap.ModifyRequest {
	return m.ProtocolOp().(ldap.ModifyRequest)
}

func (m *Message) GetModifyDNRequest() ldap.ModifyDNRequest {
	return m.ProtocolOp().(ldap.ModifyDNRequest)
}

func (m *Message) GetCompareRequest() ldap.CompareRequest {
	return m.ProtocolOp().(ldap.CompareRequest)
}

func (m *Message) GetExtendedRequest() ldap.ExtendedRequest {
	return m.ProtocolOp().(ldap.ExtendedRequest)
}
