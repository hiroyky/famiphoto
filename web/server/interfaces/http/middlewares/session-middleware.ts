import express from 'express'
import { newAuthUseCase } from '../../../di/registry'
import { SessionCookieName } from '../../../config'

export async function updateAccessToken (req: express.Request, res: express.Response, next: express.NextFunction) {
  if (!req.session || !req.session.access_token || !req.session.access_token_expires || !req.session.refresh_token) {
    return next()
  }
  const now = new Date()
  const usecase = newAuthUseCase()
  if (!usecase.isExpired(req.session.access_token_expires, now)) {
    return next()
  }

  try {
    const {
      accessToken,
      expireAt,
    } = await usecase.refreshToken(req.session.refresh_token, now)

    const sess = req.session
    sess.access_token = accessToken
    sess.access_token_expires = expireAt
    next()
  } catch (err) {
    res.clearCookie(SessionCookieName)
    next(err)
  }
}
