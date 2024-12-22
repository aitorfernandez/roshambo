import React, {
  useEffect,
} from 'react'
import PropTypes from 'prop-types'
import {
  gql,
  useMutation,
} from '@apollo/client'
import { useCookies } from 'react-cookie'

import { cookieOpts } from '../../helpers'

export function AuthenticationPage({
  history,
  match: { params: { id, token } },
}) {
  const [, setCookie] = useCookies(['id', 'jwt'])

  const [validateToken, { error }] = useMutation(AuthenticationPageValidateToken, {
    update: (_, { data }) => {
      const { id, jwt } = data.auth
      setCookie('id', id, cookieOpts)
      setCookie('jwt', jwt, cookieOpts)
    },
    onCompleted: () => {
      history.push('/play')
    },
  })

  useEffect(() => {
    async function runAuth() {
      await validateToken({
        variables: {
          input: {
            id,
            token,
          },
        },
      })
    }
    runAuth()
  }, [])

  if (error) {
    return null
  }

  return (
    <div>
      Validate access token...
    </div>
  )
}

AuthenticationPage.propTypes = {
  history: PropTypes.shape({
    push: PropTypes.func.isRequired,
  }),
  match: PropTypes.shape({
    params: PropTypes.shape({
      id: PropTypes.string.isRequired,
      token: PropTypes.string.isRequired,
    }),
  }),
}

const AuthenticationPageValidateToken = gql`
  mutation AuthenticationPageValidateToken(
    $input: ValidateTokenInput!
  ) {
    auth: validateToken(input: $input) {
      id
      jwt
    }
  }
`
