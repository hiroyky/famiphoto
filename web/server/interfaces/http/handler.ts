import express from "express";
import { checkSchema } from 'express-validator'
import expressValidation from "~/server/interfaces/http/middlewares/express-validation";
import {newAuthUseCase} from "~/server/di/registry";
import {updateAccessToken} from "~/server/interfaces/http/middlewares/session-middleware";
import {SessionCookieName} from "~/server/config";
import { ApiError } from "~/server/drivers/api-driver";
import {createProxyMiddleware} from "http-proxy-middleware";

const router = express.Router()
const nuxtConfig = useRuntimeConfig()

router.get('/api/status', (req, res) => {
    res.json({ status: 'ok' })
    res.end()
})

router.post(
    '/api/auth/login',
    express.urlencoded({ extended: true }),
    express.json(),
    checkSchema({
        userId: { in: ['body'] },
        password: { in: ['body'] },
    }),
    expressValidation,
    async (req: express.Request, res: express.Response) => {
        const { userId, password } = req.body

        const useCase = newAuthUseCase()
        try {
            const {
                accessToken,
                refreshToken,
                expireAt,
            } = await useCase.login(userId, password, new Date())
            req.session.access_token = accessToken
            req.session.access_token_expires = expireAt
            req.session.refresh_token = refreshToken
            res.sendStatus(200)
        } catch (err) {
            errorResponse(res, err)
        }
    },
)

router.post(
    '/api/auth/logout',
    updateAccessToken,
    (_: express.Request, res:express.Response) => {
        res.clearCookie(SessionCookieName)
        res.sendStatus(200)
    },
)

router.use(
    '/api/*',
    updateAccessToken,
    createProxyMiddleware({
        target: useRuntimeConfig().apiBaseUrl,
        changeOrigin: true,
        pathRewrite: {
            '^/api': '',
        },
        onProxyReq: (proxyReq, req, _) => {
            if (req.session && req.session.access_token) {
                proxyReq.setHeader('Authorization', `Bearer ${req.session.access_token}`)
            } else {
                proxyReq.setHeader('Authorization', basicAuthValue(nuxtConfig.clientId, nuxtConfig.clientSecret))
            }
        },
    }),
)

router.use(
    '/debug/graphql',
    updateAccessToken,
    createProxyMiddleware({
        target: useRuntimeConfig().apiBaseUrl,
        changeOrigin: true,
        onProxyReq: (proxyReq, req, _) => {
            if (req.session && req.session.access_token) {
                proxyReq.setHeader('Authorization', `Bearer ${req.session.access_token}`)
            } else {
                proxyReq.setHeader('Authorization', basicAuthValue(nuxtConfig.clientId, nuxtConfig.clientSecret))
            }
        },
    }),
)

function errorResponse (res: express.Response, err: any) {
    if (err instanceof ApiError) {
        if (err.httpStatus >= 500) {
            console.error(err)
        }
        res.status(err.httpStatus)
        res.send(JSON.stringify(err))
        return
    }

    console.log(err)
    res.status(500)
    res.send(JSON.stringify({ error: 'Internal server error' }))
}

export default router
