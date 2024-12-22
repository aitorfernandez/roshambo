import React, {
  useRef,
} from 'react'
import PropTypes from 'prop-types'
import styled from 'styled-components'
import { gql } from '@apollo/client'

import { t } from '../../theme'
import {
  ButtonBase,
  ContentEditable,
  Flex,
  // Box,
  Icon,
} from '../../ui'

export function Profile({
  accountID,
  // avatar,
  setProfile,
  username,
}) {
  const avatarRef = useRef(null)

  async function onSubmit(value) {
    if (value === username) {
      return
    }
    await setProfile({
      variables: {
        input: {
          accountID,
          username: value,
        },
      },
    })
  }

  function onClick() {
    avatarRef.current.click()
  }

  function onChange() {
    const file = event.target.files[0]
    console.log(file)
  }

  return (
    <Flex
      align="center"
      direction="column"
      p={[2, 0, 2, 0]}
    >
      <Input
        accept="image/png, image/jpeg"
        onChange={onChange}
        ref={avatarRef}
        type="file"
      />
      <StyledButtonBase onClick={onClick}>
        <Icon
          fs={4}
          c="gray.7"
          icon="user-ninja"
        />
      </StyledButtonBase>
      <ContentEditable
        defaultValue={'Your username here...'}
        onEdit={onSubmit}
        value={username}
      />
    </Flex>
  )
}

Profile.propTypes = {
  accountID: PropTypes.string.isRequired,
  avatar: PropTypes.string,
  setProfile: PropTypes.func.isRequired,
  username: PropTypes.string,
}

Profile.fragments = {
  profile: gql`
    fragment ProfileFragment on Profile {
      accountID
      avatar
      username
    }
  `,
}

const Input = styled.input`
  display: none;
`

const StyledButtonBase = styled(ButtonBase)`
  background: ${t('color.white.5')};
  border-radius: 50%;
  height: 36px;
  margin-bottom: ${t('space.3')};
  vertical-align: middle;
  width: 36px;
`
