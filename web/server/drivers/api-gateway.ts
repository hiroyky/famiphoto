import { ApiDriver } from './api-driver'
import {
  AuthorizationCodeResponse,
  LoginRequest,
  PostOauthTokenRequest,
  PostOauthTokenResponse,
} from '~/server/types/api-types'

export class ApiGateway {
  constructor (
    private apiDriver: ApiDriver,
  ) { }

  public async login (req: LoginRequest): Promise<AuthorizationCodeResponse> {
    const conf = useRuntimeConfig()
    const res = await this.apiDriver.request('auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: basicAuthValue(conf.clientId, conf.clientSecret),
      },
      body: JSON.stringify({
        user_id: req.userId,
        password: req.password,
      }),
    })
    const body: any = await res.json()
    return {
      accessToken: body.access_token,
      refreshToken: body.refresh_token,
      expireIn: body.expires_in,
    }
  }

  public async postOauthToken (req: PostOauthTokenRequest): Promise<PostOauthTokenResponse> {
    const form = new URLSearchParams()
    form.append('grant_type', req.grantType)
    if (req.scope) {
      form.append('scope', req.scope)
    }
    if (req.code) {
      form.append('code', req.code)
    }
    if (req.redirectUri) {
      form.append('redirect_uri', req.redirectUri)
    }
    if (req.refreshToken) {
      form.append('refresh_token', req.refreshToken)
    }
    if (req.state) {
      form.append('state', req.state)
    }

    const conf = useRuntimeConfig()
    const res = await this.apiDriver.request('oauth/v2/token', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
        Authorization: basicAuthValue(conf.clientId, conf.clientSecret),
      },
      body: form,
    })

    const body: any = await res.json()
    return {
      accessToken: body.access_token,
      expireIn: body.expires_in,
      refreshToken: body.refresh_token,
    }
  }
}
