package service

import (
	"choirsearch/internal/model"
	"context"
	"fmt"
	"time"
)

type AuditEvent struct {
	At     time.Time
	Action string
	Entity string
	ID     int64
	Detail string
}
type AuditLog struct{ Events []AuditEvent }

func NewAudit() *AuditLog { return &AuditLog{Events: []AuditEvent{}} }
func (a *AuditLog) Add(action, entity string, id int64, detail string) {
	a.Events = append(a.Events, AuditEvent{At: time.Now().UTC(), Action: action, Entity: entity, ID: id, Detail: detail})
}
func (a *AuditLog) Len() int { return len(a.Events) }
func (a *AuditLog) Last() AuditEvent {
	if len(a.Events) == 0 {
		return AuditEvent{}
	}
	return a.Events[len(a.Events)-1]
}
func (a *AuditLog) ForEntity(entity string) []AuditEvent {
	out := []AuditEvent{}
	for _, x := range a.Events {
		if x.Entity == entity {
			out = append(out, x)
		}
	}
	return out
}
func (a *AuditLog) String() string {
	out := ""
	for _, x := range a.Events {
		out += fmt.Sprintf("%s %s %d\n", x.Action, x.Entity, x.ID)
	}
	return out
}
func (s *Service) AuditSong(ctx context.Context, a *AuditLog, x model.Song) error {
	if !x.IsValid() {
		return fmt.Errorf("invalid song")
	}
	if e := s.RegisterSong(ctx, x, nil); e != nil {
		return e
	}
	a.Add("create", "song", x.ID, x.Title)
	return nil
}
func (s *Service) AuditMember(ctx context.Context, a *AuditLog, x model.Member) error {
	if e := s.AddMember(ctx, x); e != nil {
		return e
	}
	a.Add("create", "member", x.ID, x.Name)
	return nil
}
func (s *Service) AuditRehearsal(ctx context.Context, a *AuditLog, x model.Rehearsal) error {
	if e := s.Schedule(ctx, x); e != nil {
		return e
	}
	a.Add("create", "rehearsal", x.ID, x.Name)
	return nil
}
