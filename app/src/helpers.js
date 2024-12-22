export const cookieOpts = {
  // ref: https://security.stackexchange.com/questions/212809/what-is-the-most-secure-way-to-store-cross-subdomain-cookies
  path: '/',
  domain: 'localhost',
  sameSite: 'lax',
}

export const moves = ['rock', 'paper', 'scissors']
