// package service

// import (
// 	"github.com/hellyab/techreview/entities"
// 	"github.com/hellyab/techreview/user"
// )

// // SessionServiceImpl implements user.SessionService interface
// type SessionServiceImpl struct {
// 	sessionRepo user.SessionRepository
// }

// // NewSessionService  returns a new SessionService object
// func NewSessionService(sessRepository user.SessionRepository) user.SessionService {
// 	return &SessionServiceImpl{sessionRepo: sessRepository}
// }

// // Session returns a given stored session
// func (ss *SessionServiceImpl) Session(sessionID string) (*entities.Session, []error) {
// 	sess, errs := ss.sessionRepo.Session(sessionID)
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return sess, errs
// }

// // StoreSession stores a given session
// func (ss *SessionServiceImpl) StoreSession(session *entities.Session) (*entities.Session, []error) {
// 	sess, errs := ss.sessionRepo.StoreSession(session)
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return sess, errs
// }

// // DeleteSession deletes a given session
// func (ss *SessionServiceImpl) DeleteSession(sessionID string) (*entities.Session, []error) {
// 	sess, errs := ss.sessionRepo.DeleteSession(sessionID)
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return sess, errs
// }
