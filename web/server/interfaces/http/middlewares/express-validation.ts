import express from 'express'
import { validationResult } from 'express-validator'

export default function expressValidation (req: express.Request, res: express.Response, next: express.NextFunction) {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    res.status(400).json({ errors: errors.array() })
    return
  }
  next()
}
