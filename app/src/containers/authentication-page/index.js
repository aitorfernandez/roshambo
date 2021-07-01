import React from 'react'
import PropTypes from 'prop-types'
import {
  gql,
  useMutation,
} from '@apollo/client'
import { useCookies } from 'react-cookie'

export function AuthenticationPage({
  history,
  match: { params: { id, token } },
}) {
  const [validateToken, { error }] = useMutation(AuthenticationPageValidateToken, {
    update: (_, { data }) => {
      console.log(data)
    },
    onCompleted: () => {

    },
  })

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
