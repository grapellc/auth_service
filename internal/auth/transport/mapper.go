package transport

import (
	"github.com/google/uuid"
	"github.com/your-moon/grape-shared/entities"
	"github.com/your-moon/grape-shared/proto/authv1"
)

func AuthLogToProto(l *entities.AuthLogUUID) *authv1.AuthLogMsg {
	if l == nil {
		return nil
	}
	msg := &authv1.AuthLogMsg{
		Identifier:    l.Identifier,
		Action:        l.Action,
		Status:        l.Status,
		FailureReason: l.FailureReason,
		IpAddress:     l.IPAddress,
		UserAgent:     l.UserAgent,
	}
	if l.UserID != nil {
		s := l.UserID.String()
		msg.UserId = &s
	}
	return msg
}

func ProtoToAuthLog(msg *authv1.AuthLogMsg) *entities.AuthLogUUID {
	if msg == nil {
		return nil
	}
	l := &entities.AuthLogUUID{
		Identifier:    msg.Identifier,
		Action:        msg.Action,
		Status:        msg.Status,
		FailureReason: msg.FailureReason,
		IPAddress:     msg.IpAddress,
		UserAgent:     msg.UserAgent,
	}
	if msg.UserId != nil && *msg.UserId != "" {
		if id, err := uuid.Parse(*msg.UserId); err == nil {
			l.UserID = &id
		}
	}
	return l
}

func UserToProto(u *entities.UserUUID) *authv1.UserMsg {
	if u == nil {
		return nil
	}
	msg := &authv1.UserMsg{
		Id:              u.ID.String(),
		Email:           u.Email,
		Role:            u.Role,
		IsPhoneVerified: u.IsPhoneVerified,
		IsEmailVerified: u.IsEmailVerified,
	}
	if u.PhoneNumber != nil {
		msg.PhoneNumber = u.PhoneNumber
	}
	return msg
}

func ProtoToUser(msg *authv1.UserMsg) *entities.UserUUID {
	if msg == nil {
		return nil
	}
	u := &entities.UserUUID{
		Email:            msg.Email,
		Role:             msg.Role,
		IsPhoneVerified:  msg.IsPhoneVerified,
		IsEmailVerified:  msg.IsEmailVerified,
	}
	if msg.Id != "" {
		if id, err := uuid.Parse(msg.Id); err == nil {
			u.ID = id
		}
	}
	if msg.PhoneNumber != nil && *msg.PhoneNumber != "" {
		u.PhoneNumber = msg.PhoneNumber
	}
	return u
}
